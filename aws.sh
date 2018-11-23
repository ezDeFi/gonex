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
INSTANCE_TYPE=t2.large
declare -A IMAGE_ID
IMAGE_ID=(
	[us-east-1]=ami-0ac019f4fcb7cb7e6
	[us-east-2]=ami-0f65671a86f061fcd
	[us-west-1]=ami-063aa838bd7631e0b
	[us-west-2]=ami-0bbe6b35405ecebdb
	[ap-northeast-2]=ami-06e7b9c5e0c4dd014
	[ap-southeast-1]=ami-0c5199d385b432989
	[ap-southeast-2]=ami-07a3bd4944eb120a0
	[ca-central-1]=ami-0427e8367e3770df1
	[eu-central-1]=ami-0bdf93799014acdc4
	[eu-west-2]=ami-0b0a60c0a2bd40612
	[eu-west-3]=ami-08182c55a1c188dee
)
KEY_NAME=dvietha@gmail.com
BOOTNODE_REGION=us-east-1
ETHSTATS=nexty-testnet@198.13.40.85:80
CONTRACT_ADDR=cafecafecafecafecafecafecafecafecafecafe
EPOCH=60
SSH_USER=ubuntu

OUTPUT_TYPE=table

# Global Variables
BOOTNODE=
NETWORK_NAME=NextyLoadTest
NETWORK_ID=50913

# COMMAND SHORTCUTS
SSH="ssh -oStrictHostKeyChecking=no -o BatchMode=yes"
SCP="scp -oStrictHostKeyChecking=no -o BatchMode=yes"
PSCP="pscp -OStrictHostKeyChecking=no -OBatchMode=yes"
SSH_COPY_ID="ssh-copy-id -f"
GETH="./geth-linux-amd64 --verbosity=5 --pprof --pprofaddr 0.0.0.0 --syncmode full --networkid $NETWORK_ID --rpc --rpcapi db,eth,net,web3,personal --rpccorsdomain \"*\" --rpcaddr 0.0.0.0 --gasprice 0 --targetgaslimit 48000000000 --txpool.nolocals --txpool.pricelimit 0 --txpool.accountslots 64 --txpool.globalslots 16384 --txpool.accountqueue 128 --txpool.globalqueue 4096 --gcmode=archive"

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
				--region=$REGION\
				--key-name=$KEY_NAME\
				--tag-specifications="ResourceType=instance,Tags=[{Key=Name,Value=$BOOTNODE_NAME}]"\
				--query "Instances[].[InstanceId]"\
				--output=text | tr "\n" " " | awk '{$1=$1};1')
		aws ec2 wait instance-running --region $REGION --instance-ids $ID
	fi

	IP=`instance_ip $ID`
	ssh_ready $SSH_USER@$IP

	if ! $SSH $SSH_USER@$IP stat boot.key \> /dev/null 2\>\&1; then
		# remote boot.key not exist
		./build/bin/bootnode --genkey=boot.key
		$SCP build/bin/bootnode-linux-amd64 $SSH_USER@$IP:./
		$SCP boot.key $SSH_USER@$IP:./
	fi

	$SSH $SSH_USER@$IP "nohup yes | ./bootnode-linux-amd64 -nodekey=boot.key -verbosity 9 &>bootnode.log &"

	echo enode://`./build/bin/bootnode -nodekey=boot.key -writeaddress`@$IP:33333
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
	wait

	# ssh_key_copy $ALL_IPs
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
		$SSH $SSH_USER@${IPs[$i]} "./geth-linux-amd64 account new --password <(echo password)" >| /tmp/aws.sh/account/$i &
	done
	wait
	arr=()
	for i in "${!IPs[@]}"; do
		ACCOUNT=$(cat /tmp/aws.sh/account/$i)
		arr=(${arr[@]} ${ACCOUNT:10:40})
	done
	echo "${arr[@]}"
}

# load pre-fund account from keystore folder
function load_pre_fund_accounts {
	arr=(${CONTRACT_ADDR})
	for file in ./.gonex/keystore/*; do
		if [[ -f $file ]]; then
			filename=$(basename -- "$file")
			arr=(${arr[@]} ${filename:37:78})
		fi
	done
	echo "${arr[@]}"
}

function test_load_pre_fund_accounts {
	echo `load_pre_fund_accounts`
}

# generate the genesis json file
function generate_genesis {
	ACs=($@)
	PFACs=(`load_pre_fund_accounts`)

	(	echo 2
		echo 3
		echo $BLOCK_TIME
		for AC in "${ACs[@]}"; do
			echo $AC
		done
		echo
		for PFAC in "${PFACs[@]}"; do
			echo $PFAC
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
	CONTRACT_CODE="60806040526004361061004c576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063b688a36314610051578063d66d9e1914610068575b600080fd5b34801561005d57600080fd5b5061006661007f565b005b34801561007457600080fd5b5061007d6100e7565b005b60013390806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008090505b6001805490508110156102685760018181548110151561010957fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561025b5760018080805490500381548110151561017c57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166001828154811015156101b657fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060018080805490500381548110151561021357fe5b9060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556001805480919060019003610255919061026c565b50610269565b80806001019150506100ed565b5b50565b815481835581811115610293578183600052602060002091820191016102929190610298565b5b505050565b6102ba91905b808211156102b657600081600090555060010161029e565b5090565b905600a165627a7a723058201643c670757cdf1940e3f657b7ad97d5f76cfe82de027f5470c9b85424fd135a0029"

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
		(	$SSH $SSH_USER@$IP "./geth-linux-amd64 init *.json"
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
	$PSCP -h <(printf "%s\n" $IPs) -l ubuntu ./build/bin/geth-linux-amd64 /home/ubuntu/
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
