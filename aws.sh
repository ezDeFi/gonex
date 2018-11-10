#!/bin/bash

# A POSIX variable
OPTIND=1         # Reset in case getopts has been used previously in the shell.

# Initialize our own variables:

while getopts "h?frdw" opt; do
    case "$opt" in
    h|\?)
        echo "$(basename ""$0"") [-h|-?] command"
        exit 0
        ;;
    esac
done

shift $((OPTIND-1))

[ "${1:-}" = "--" ] && shift

# CONFIG
BLOCK_TIME=2
BOOTNODE_NAME=BootNode
INSTANCE_NAME=LoadTest
INSTANCE_TYPE=t2.micro
declare -A IMAGE_ID
IMAGE_ID=(
	[ap-southeast-1]=ami-0c5199d385b432989
	[us-east-1]=ami-0ac019f4fcb7cb7e6
	[eu-central-1]=ami-0bdf93799014acdc4
	[sa-east-1]=ami-03c6239555bb12112
	[us-west-1]=ami-063aa838bd7631e0b
	[ca-central-1]=ami-0427e8367e3770df1
	[ap-northeast-1]=ami-07ad4b1c3af1ea214
	[eu-west-1]=ami-0ab7944c6328200be
	[ap-south-1]=ami-0d773a3b7bb2bb1c1
	[us-east-2]=ami-0f65671a86f061fcd
	[us-west-2]=ami-0bbe6b35405ecebdb
	[ap-northeast-2]=ami-06e7b9c5e0c4dd014
	[ap-southeast-2]=ami-07a3bd4944eb120a0
	[eu-west-2]=ami-0b0a60c0a2bd40612
	[eu-west-3]=ami-08182c55a1c188dee
)
KEY_NAME=zergity@gmail.com
BOOTNODE_REGION=ap-southeast-1
ETHSTATS=nexty-testnet@198.13.40.85:80
CONTRACT_ADDR=00000000000000000000000000000000000000ff
EPOCH=20
SSH_USER=ubuntu

OUTPUT_TYPE=table

# Global Variables
BOOTNODE=
NETWORK_NAME=NextyLoadTest
NETWORK_ID=50913

# COMMAND SHORTCUTS
SSH="ssh -oStrictHostKeyChecking=no -o BatchMode=yes"
SCP="scp -oStrictHostKeyChecking=no -o BatchMode=yes"
PSCP="parallel-scp -OStrictHostKeyChecking=no -OBatchMode=yes"
SSH_COPY_ID="ssh-copy-id -f"
GETH="./geth --verbosity=5 --syncmode full --networkid $NETWORK_ID --rpc --rpcapi db,eth,net,web3,personal --rpccorsdomain \"*\" --rpcaddr 0.0.0.0 --gasprice 0 --targetgaslimit 42000000 -gcmode=archive"

function trim {
	awk '{$1=$1};1'
}

function instance_ids {
	aws ec2 describe-instances\
			--region=$REGION\
			--filters Name=tag:$1,Values=$2 Name=instance-state-name,Values=running,stopped\
			--query="Reservations[].Instances[].[InstanceId]"\
			--output=text | tr "\n" " " | awk '{$1=$1};1'
}

function instance_state {
	aws ec2 describe-instance-status\
			--region=$REGION\
			--output=$OUTPUT_TYPE\
			--instance-id $1\
			--query "InstanceStatuses[].InstanceState[].[Name]"\
			--output=text | tr "\n" " " | awk '{$1=$1};1'
}

function instance_ip {
	aws ec2 describe-instances\
			--region=$REGION\
			--instance-ids $@\
			--query "Reservations[].Instances[].[PublicIpAddress]"\
			--output=text | tr "\n" " " | awk '{$1=$1};1'
}

function ssh_ready {
	# Probe SSH connection until it's avalable 
	X_READY=''
	while [ ! $X_READY ]; do
		sleep 1s
		set +e
		OUT=$(ssh -o ConnectTimeout=1 -o StrictHostKeyChecking=no -o BatchMode=yes $@ exit &>/dev/null )
		[[ $? = 0 ]] && X_READY='ready'
		set -e
	done 
}

function bootnode {
	REGION=$BOOTNODE_REGION
	ID=`instance_ids Name $BOOTNODE_NAME | awk {'print $NF'}`

	if [ -n "$ID" ]; then
		STATE=`instance_state $ID`
		if [ "$STATE" != "running" ]; then
			aws ec2 start-instances --instance-ids $ID &>/dev/null
			aws ec2 wait instance-running --instance-ids $ID
		fi
	fi

	if [ -z "$ID" ]; then
		ID=$(aws ec2 run-instances\
				--image-id=${IMAGE_ID[$REGION]}\
				--instance-type=$INSTANCE_TYPE\
				--key-name=$KEY_NAME\
				--tag-specifications="ResourceType=instance,Tags=[{Key=Name,Value=$BOOTNODE_NAME}]"\
				--query "Instances[].[InstanceId]"\
				--output=text | tr "\n" " " | awk '{$1=$1};1')
		aws ec2 wait instance-running --instance-ids $ID
	fi

	if [ -f "bin/bootnode" ]; then
		# workaround for not-up-to-date code
		cp bin/bootnode build/bin/
	fi

	IP=`instance_ip $ID`
	ssh_ready $SSH_USER@$IP

	if ! $SSH $SSH_USER@$IP stat boot.key \> /dev/null 2\>\&1; then
		# remote boot.key not exist
		strip -s build/bin/bootnode
		build/bin/bootnode --genkey=boot.key
		$SCP build/bin/bootnode $SSH_USER@$IP:./
		$SCP boot.key $SSH_USER@$IP:./
	fi

	$SSH $SSH_USER@$IP "nohup yes | ./bootnode -nodekey=boot.key -verbosity 9 &>bootnode.log &"

	echo enode://`build/bin/bootnode -nodekey=boot.key -writeaddress`@$IP:33333
}

function load {
	COUNT=${1:-1}
	BOOTNODE=`bootnode`

	rm -rf /tmp/aws.sh/ips
	mkdir -p /tmp/aws.sh/ips
	for REGION in "${!IMAGE_ID[@]}"; do
		(	IDs=`launch_instance $COUNT`
			aws ec2 wait instance-running --instance-ids $IDs --region=$REGION
			IPs=`instance_ip $IDs`
			LAST_IP=${IPs##* }
			ssh_ready $SSH_USER@$LAST_IP
			echo $IPs >| /tmp/aws.sh/ips/$REGION
		) &
	done
	wait

	ALL_IPs=`cat /tmp/aws.sh/ips/* | tr "\n" " "`

	deploy $ALL_IPs &&\
	start $ALL_IPs | tr "\n" " " | awk '{$1=$1};1'
	ssh_key_copy $ALL_IPs
}

function ssh_key_copy {
	for pk in ./sshpubkeys/*.pub; do
		for IP in $@; do
			$SSH_COPY_ID -i $pk $SSH_USER@$IP &>/dev/null &
		done
	done
	wait
}

function create_account {
	IPs=($@)
	rm -rf /tmp/aws.sh/account
	mkdir -p /tmp/aws.sh/account
	for i in "${!IPs[@]}"; do
		$SSH $SSH_USER@${IPs[$i]} "./geth account new --password <(echo password)" >| /tmp/aws.sh/account/$i &
	done
	wait
	for i in "${!IPs[@]}"; do
		ACCOUNT=$(cat /tmp/aws.sh/account/$i)
		printf "%s " ${ACCOUNT:10:40}
	done
}

function generate_genesis {
	ACs=($@)

	(	echo 2
		echo 3
		echo $BLOCK_TIME
		for AC in "${ACs[@]}"; do
			echo $AC
		done
		echo
		for AC in "${ACs[@]}"; do
			echo $AC
		done
		echo
		echo $NETWORK_ID
		echo 2
		echo 2
		echo
	) >| /tmp/puppeth.json

	GENESIS_JSON=$NETWORK_NAME.json
	rm *.json ~/.puppeth/* /tmp/$GENESIS_JSON
	./build/bin/puppeth --network=$NETWORK_NAME < /tmp/puppeth.json >/dev/null

	if [ ! -f "$GENESIS_JSON" ]; then
		>&2 echo "Unable to create genesis file with Puppeth"
		exit -1
	fi

	echo $GENESIS_JSON
}

function init_smartcontract {
	ACs=($@)
	CODE=
	# Nexty Smart Contract
	for AC in "${ACs[@]}"; do
		CODE="$CODE signers[0x$AC] = 0x$AC; sealers.push(0x$AC);"
	done
	sed "s/CODE_HERE/$CODE/g" nexty.sol >| /tmp/nexty.sol

	solc --bin-runtime /tmp/nexty.sol 2>/dev/null | tail -n1
}

function init_genesis {
	ACs=($@)

	GENESIS_JSON=`generate_genesis ${ACs[@]}`
	CONTRACT_CODE=`init_smartcontract ${ACs[@]}`

	mv ./$GENESIS_JSON /tmp/$GENESIS_JSON.template
	while IFS='' read -r line || [[ -n "$line" ]]; do
		if [[ "$line" == *"\"epoch\": "* ]]; then
			echo \"epoch\": $EPOCH
			continue
		fi
		echo $line
		if [[ "$line" == *"$CONTRACT_ADDR"* ]]; then
			echo \"code\": \"0x$CONTRACT_CODE\",
			echo \"storage\": \{
				idx=3086617846
				for AC in "${ACs[@]}"; do
					echo \"0xb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2$(printf "%08x" $idx)\": \"0x$AC\",
					((idx++))
				done
				echo \"0x0000000000000000000000000000000000000000000000000000000000000001\": \"$(printf "%04x" ${#ACs[@]})\"
			echo \},
		fi
	done < /tmp/$GENESIS_JSON.template >| ./$GENESIS_JSON

	echo $GENESIS_JSON
}

function geth_start {
	IPs=($@)
	for IP in "${IPs[@]}"; do
		(	$SSH $SSH_USER@$IP "./geth init *.json"
			$SSH $SSH_USER@$IP "nohup $GETH --bootnodes $BOOTNODE --mine --unlock 0 --password <(echo password) --ethstats $IP:$ETHSTATS &>./geth.log &"
		) &
	done
	wait
}

function start {
	IPs=($@)
	ACs=(`create_account ${IPs[@]}`)
	
	GENESIS_JSON=`init_genesis ${ACs[@]}`

	$PSCP -h <(printf "%s\n" $@) -l $SSH_USER $GENESIS_JSON /home/ubuntu/

	mv $GENESIS_JSON /tmp/

	geth_start ${IPs[@]}
}

# restart InstantName
function restart {
	for REGION in "${!IMAGE_ID[@]}"; do
		(	IDs=`instance_ids Name ${1:-$INSTANCE_NAME}`
			geth_restart `instance_ip $IDs`
		) &
	done
	wait
}

function deploy {
	IPs="$@"
	strip -s build/bin/geth
	$PSCP -h <(printf "%s\n" $IPs) -l ubuntu ./build/bin/geth /home/ubuntu/
}

function launch_instance {
	COUNT=${1:-1}
	aws ec2 run-instances\
			--region=$REGION\
			--image-id=${IMAGE_ID[$REGION]}\
			--instance-type=$INSTANCE_TYPE\
			--key-name=$KEY_NAME\
			--tag-specifications="ResourceType=instance,Tags=[{Key=Name,Value=$INSTANCE_NAME}]"\
			--count=$COUNT\
			--query "Instances[].[InstanceId]"\
			--output=text | tr "\n" " " | awk '{$1=$1};1'
}

function terminate {
	if [ "$1" = "bn" ]; then
		REGION=$BOOTNODE_REGION
		_terminate $BOOTNODE_NAME
	elif [ "$1" = "all" ]; then
		for REGION in "${!IMAGE_ID[@]}"; do
			_terminate &
		done
		wait
	fi
}

function _terminate {
	IDs=`instance_ids Name ${1:-$INSTANCE_NAME}`

	if [ -z "$IDs" ]; then
		return
	fi

	stop `instance_ip $IDs`

	aws ec2 terminate-instances\
			--region=$REGION\
			--instance-ids $IDs >/dev/null
}

function stop {
	for IP in $@; do
		$SSH $SSH_USER@$IP killall -q --signal SIGINT geth &
	done
	wait
}

"$@"
