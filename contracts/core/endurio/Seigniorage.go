// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package endurio

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AbsorbableABI is the input ABI used to generate the binding from.
const AbsorbableABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"absorptionDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"absorptionExpiration\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"emptive\",\"type\":\"bool\"}],\"name\":\"Absorption\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Stop\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"Ask\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Bid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"index\",\"type\":\"bytes32\"}],\"name\":\"calcOrderID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"haveAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wantAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"assistingID\",\"type\":\"bytes32\"}],\"name\":\"findAssistingID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getOrder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"have\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"prevID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextID\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRemainToAbsorb\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"next\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"target\",\"type\":\"uint256\"}],\"name\":\"onBlockInitialized\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"prev\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"volatileToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stablizeToken\",\"type\":\"address\"}],\"name\":\"registerTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"}],\"name\":\"top\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AbsorbableFuncSigs maps the 4-byte function signature to its string representation.
var AbsorbableFuncSigs = map[string]string{
	"69c07d31": "Ask()",
	"6e6452cb": "Bid()",
	"f318722b": "calcOrderID(address,bytes32)",
	"43271d79": "cancel(bool,bytes32)",
	"ced4aac8": "findAssistingID(bool,address,uint256,uint256,bytes32)",
	"07c399a3": "getOrder(bool,bytes32)",
	"ee1a68c6": "getRemainToAbsorb()",
	"4ea09797": "next(bool,bytes32)",
	"be91d729": "onBlockInitialized(uint256)",
	"0d90b10a": "prev(bool,bytes32)",
	"aa1c259c": "registerTokens(address,address)",
	"8aa3f897": "top(bool)",
}

// AbsorbableBin is the compiled bytecode used for deploying new contracts.
var AbsorbableBin = "0x608060405262049d4060035560026003548161001757fe5b0460045534801561002757600080fd5b506040516119e13803806119e18339818101604052604081101561004a57600080fd5b508051602090910151801561005f5760038190555b600082116100705760028104610072565b815b600455505061195b806100866000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80638aa3f897116100715780638aa3f897146101ba578063aa1c259c146101d9578063be91d72914610207578063ced4aac814610224578063ee1a68c614610264578063f318722b14610287576100b4565b806307c399a3146100b95780630d90b10a1461011357806343271d791461014a5780634ea097971461017157806369c07d31146101965780636e6452cb146101b2575b600080fd5b6100de600480360360408110156100cf57600080fd5b508035151590602001356102b3565b604080516001600160a01b03909616865260208601949094528484019290925260608401526080830152519081900360a00190f35b6101386004803603604081101561012957600080fd5b508035151590602001356102fa565b60408051918252519081900360200190f35b61016f6004803603604081101561016057600080fd5b50803515159060200135610322565b005b6101386004803603604081101561018757600080fd5b508035151590602001356103a8565b61019e6103cc565b604080519115158252519081900360200190f35b61019e6103d1565b610138600480360360208110156101d057600080fd5b503515156103d6565b61016f600480360360408110156101ef57600080fd5b506001600160a01b03813581169160200135166103f7565b61016f6004803603602081101561021d57600080fd5b5035610572565b610138600480360360a081101561023a57600080fd5b5080351515906001600160a01b0360208201351690604081013590606081013590608001356106a4565b61026c61070c565b60408051921515835260208301919091528051918290030190f35b6101386004803603604081101561029d57600080fd5b506001600160a01b0381351690602001356107b6565b90151560009081526020818152604080832093835260029384019091529020805460018201549282015460038301546004909301546001600160a01b039092169490929190565b8115156000908152602081815260408083208484526002019091529020600301545b92915050565b8115156000908152602081815260408083208484526002810190925290912080546001600160a01b03163314610392576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037b93232b91036b0b5b2b960811b604482015290519081900360640190fd5b6103a2828463ffffffff6107c216565b50505050565b90151560009081526020818152604080832093835260029093019052206004015490565b600081565b600181565b80151560009081526020819052604081206103f0816108bc565b9392505050565b6001546001600160a01b031615610455576040805162461bcd60e51b815260206004820152601960248201527f566f6c6174696c65546f6b656e20616c72656164792073657400000000000000604482015290519081900360640190fd5b6002546001600160a01b0316156104b3576040805162461bcd60e51b815260206004820152601960248201527f537461626c697a65546f6b656e20616c72656164792073657400000000000000604482015290519081900360640190fd5b600180546001600160a01b038085166001600160a01b03199283161790925560028054928416929091169190911790556104ed82826108d5565b600254604080516318160ddd60e01b815290516000926001600160a01b0316916318160ddd916004808301926020929190829003018186803b15801561053257600080fd5b505afa158015610546573d6000803e3d6000fd5b505050506040513d602081101561055c57600080fd5b5051905061056d8180600080610949565b505050565b33156105b6576040805162461bcd60e51b815260206004820152600e60248201526d636f6e73656e737573206f6e6c7960901b604482015290519081900360640190fd5b6105c060056109f7565b156105cd576105cd610a10565b600254604080516318160ddd60e01b815290516000926001600160a01b0316916318160ddd916004808301926020929190829003018186803b15801561061257600080fd5b505afa158015610626573d6000803e3d6000fd5b505050506040513d602081101561063c57600080fd5b5051905081156106825761064e610a55565b15610665576106608282600080610949565b610682565b61066f8183610a67565b1561068257610682828260016000610949565b61069360058263ffffffff610b1516565b156106a0576106a0610b5b565b5050565b84151560009081526020819052604081206106bd6118f7565b506040805160a0810182526001600160a01b038816815260208101879052908101859052600060608201819052608082015261070082828663ffffffff610efe16565b98975050505050505050565b6007546000908190610723575060009050806107b2565b60016107ad600560020154600260009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b15801561077c57600080fd5b505afa158015610790573d6000803e3d6000fd5b505050506040513d60208110156107a657600080fd5b5051610f70565b915091505b9091565b60006103f08383610f8a565b6107ca6118f7565b50600081815260028084016020908152604092839020835160a08101855281546001600160a01b031681526001820154928101839052928101549383019390935260038301546060830152600490920154608082015290156108ac57825481516020808401516040805163a9059cbb60e01b81526001600160a01b039485166004820152602481019290925251929093169263a9059cbb92604480830193928290030181600087803b15801561087f57600080fd5b505af1158015610893573d6000803e3d6000fd5b505050506040513d60208110156108a957600080fd5b50505b61056d838363ffffffff61105716565b6000808052600282016020526040902060040154919050565b600080805260205261090e7fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5838363ffffffff6110b716565b600160009081526020526106a07fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d828463ffffffff6110b716565b6040805160a0810182526003544301808252602082018690529181018690526000606082018190528315156080909201829052600592909255600685905560078690556008805461ffff19166101009092029190911790556109ab8585610f70565b60408051828152602081018790528515158183015290519192507f0427b353dc7214e3d8c7f5039475a8e729f4d62922937381e304cd03becf66d2919081900360600190a15050505050565b6000610a0282611264565b801561031c57505054431190565b60006005819055600681905560078190556008805461ffff191690556040517fbedf0f4abfe86d4ffad593d9607fe70e83ea706033d44d24b3b6283cf3fc4f6b9190a1565b6000610a6160056109f7565b90505b90565b600082821415610a795750600061031c565b6006546007541415610a8d5750600161031c565b82821115610ae157600654600754848403911015610ac357600654600754036002818381610ab757fe5b0410159250505061031c565b600754600654036002828281610ad557fe5b0411159250505061031c565b600654600754838503911115610b0357600754600654036002818381610ab757fe5b600654600754036002828281610ad557fe5b6000610b2083611264565b8015610b315750600383015460ff16155b8015610b41575082600201548214155b80156103f057506103f0836001015483856002015461126a565b6000610b6561129a565b90506000806000808413610b7a576001610b7d565b60005b151581526020810191909152604001600090812060025481549193506001600160a01b039182169116149080610bc483610bb687611366565b86919063ffffffff61137c16565b6008549193509150610100900460ff16610be2575050505050610efc565b801580610bed575081155b15610bfc575050505050610efc565b600254604080516318160ddd60e01b815290516000926001600160a01b0316916318160ddd916004808301926020929190829003018186803b158015610c4157600080fd5b505afa158015610c55573d6000803e3d6000fd5b505050506040513d6020811015610c6b57600080fd5b50519050610c8060058263ffffffff610b1516565b610c8f57505050505050610efc565b60008085610c9e578484610ca1565b83855b6009548954604080516370a0823160e01b81526001600160a01b03938416600482018190529151959750939550939116916370a08231916024808301926020929190829003018186803b158015610cf757600080fd5b505afa158015610d0b573d6000803e3d6000fd5b505050506040513d6020811015610d2157600080fd5b5051831115610d3857505050505050505050610efc565b8754604080516337ed875b60e11b81526001600160a01b0384811660048301526024820187905291519190921691636fdb0eb691604480830192600092919082900301818387803b158015610d8c57600080fd5b505af1158015610da0573d6000803e3d6000fd5b505089546040805163117f5a5560e01b81526004810188905290516001600160a01b03909216935063117f5a55925060248082019260009290919082900301818387803b158015610df057600080fd5b505af1158015610e04573d6000803e3d6000fd5b50505060018901546040805163bdfde91160e01b81526004810186905290516001600160a01b03909216925063bdfde91191602480830192600092919082900301818387803b158015610e5657600080fd5b505af1158015610e6a573d6000803e3d6000fd5b5050505060018801546040805163a9059cbb60e01b81526001600160a01b038481166004830152602482018690529151919092169163a9059cbb9160448083019260209291908290030181600087803b158015610ec657600080fd5b505af1158015610eda573d6000803e3d6000fd5b505050506040513d6020811015610ef057600080fd5b50505050505050505050505b565b600081815260028401602052604081205b6004015460008181526002860160205260409020909250610f36848263ffffffff61162316565b15610f0f575b6003015460008181526002860160205260409020909250610f63848263ffffffff61162316565b610f3c5750909392505050565b6000818311610f84578282036000036103f0565b50900390565b60006002838360405160200180836001600160a01b03166001600160a01b031660601b8152601401828152602001925050506040516020818303038152906040526040518082805190602001908083835b60208310610ffa5780518252601f199092019160209182019101610fdb565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015611039573d6000803e3d6000fd5b5050506040513d602081101561104e57600080fd5b50519392505050565b6000818152600292830160205260408082206004808201805460038085018054885286882090940182905583549187529486209094019390935593835280546001600160a01b031916815560018101839055909301819055908190559055565b818360000160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550808360010160006101000a8154816001600160a01b0302191690836001600160a01b031602179055506040518060a00160405280306001600160a01b0316815260200160008152602001600081526020016000801b815260200160001960001b8152508360020160008060001b815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550602082015181600101556040820151816002015560608201518160030155608082015181600401559050506040518060a00160405280306001600160a01b0316815260200160008152602001600181526020016000801b815260200160001960001b81525083600201600060001960001b815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060208201518160010155604082015181600201556060820151816003015560808201518160040155905050505050565b54151590565b600082841115801561127c5750818311155b8061129257508284101580156112925750818310155b949350505050565b6000806112b1600560020154600560010154610f70565b600754600254604080516318160ddd60e01b8152905193945060009361130193926001600160a01b0316916318160ddd916004808301926020929190829003018186803b15801561077c57600080fd5b905061130f60008284611641565b61131e57600092505050610a64565b6000600454838161132b57fe5b6008549190059150610100900460ff161561134557600290055b61135160008284611641565b61135f57509150610a649050565b9250505090565b6000808213611378578160000361031c565b5090565b600080600061138a866108bc565b90505b600019811480159061139e57508382105b15611619576000818152600287016020526040812090866113c35781600201546113c9565b81600101545b90506000876113dc5782600101546113e2565b82600201545b90508487038083116114d857895460018501546040805163117f5a5560e01b81526004810192909252516001600160a01b039092169163117f5a559160248082019260009290919082900301818387803b15801561143f57600080fd5b505af1158015611453573d6000803e3d6000fd5b50505060018b015460028601546040805163bdfde91160e01b81526004810192909252516001600160a01b03909216925063bdfde91191602480830192600092919082900301818387803b1580156114aa57600080fd5b505af11580156114be573d6000803e3d6000fd5b5050505060048401546114d18b8761166c565b9450611608565b82818302816114e357fe5b0491508092508682880110156114fe575061161b9350505050565b60008961150b578261150d565b835b905060008a61151c578461151e565b835b8c546040805163117f5a5560e01b81526004810186905290519293506001600160a01b039091169163117f5a559160248082019260009290919082900301818387803b15801561156d57600080fd5b505af1158015611581573d6000803e3d6000fd5b50505060018d01546040805163bdfde91160e01b81526004810185905290516001600160a01b03909216925063bdfde91191602480830192600092919082900301818387803b1580156115d357600080fd5b505af11580156115e7573d6000803e3d6000fd5b5061160092508e9150899050848463ffffffff61172f16565b506000199550505b50949094019392909201915061138d565b505b935093915050565b60408201516001820154600290920154602090930151910291021190565b60008284131580156116535750818313155b8061129257508284121580156112925750501315919050565b6116746118f7565b50600081815260028084016020908152604092839020835160a08101855281546001600160a01b031681526001820154928101929092529182015492810183905260038201546060820152600490910154608082015290156108ac5760018301548151604080840151815163a9059cbb60e01b81526001600160a01b03938416600482015260248101919091529051919092169163a9059cbb9160448083019260209291908290030181600087803b15801561087f57600080fd5b60008381526002850160205260409020600181015483111561178e576040805162461bcd60e51b815260206004820152601360248201527250503a2066696c6c61626c65203e206861766560681b604482015290519081900360640190fd5b80600201548211156117dd576040805162461bcd60e51b815260206004820152601360248201527214140e88199a5b1b18589b19480f881dd85b9d606a1b604482015290519081900360640190fd5b600180820180548590039055600282018054849003905585015481546040805163a9059cbb60e01b81526001600160a01b039283166004820152602481018690529051919092169163a9059cbb9160448083019260209291908290030181600087803b15801561184c57600080fd5b505af1158015611860573d6000803e3d6000fd5b505050506040513d602081101561187657600080fd5b50506040805160a08101825282546001600160a01b031681526001830154602082015260028301549181019190915260038201546060820152600482015460808201526118c2906118de565b156118d7576118d7858563ffffffff6107c216565b5050505050565b600081602001516000148061031c575050604001511590565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091529056fea26469706673582212209b2e87fce192140e844d82c41e649380fa22a52c80446bd52d629e1358cda7f364736f6c63430006080033"

// DeployAbsorbable deploys a new Ethereum contract, binding an instance of Absorbable to it.
func DeployAbsorbable(auth *bind.TransactOpts, backend bind.ContractBackend, absorptionDuration *big.Int, absorptionExpiration *big.Int) (common.Address, *types.Transaction, *Absorbable, error) {
	parsed, err := abi.JSON(strings.NewReader(AbsorbableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AbsorbableBin), backend, absorptionDuration, absorptionExpiration)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Absorbable{AbsorbableCaller: AbsorbableCaller{contract: contract}, AbsorbableTransactor: AbsorbableTransactor{contract: contract}, AbsorbableFilterer: AbsorbableFilterer{contract: contract}}, nil
}

// Absorbable is an auto generated Go binding around an Ethereum contract.
type Absorbable struct {
	AbsorbableCaller     // Read-only binding to the contract
	AbsorbableTransactor // Write-only binding to the contract
	AbsorbableFilterer   // Log filterer for contract events
}

// AbsorbableCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbsorbableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbsorbableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbsorbableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbsorbableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbsorbableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbsorbableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbsorbableSession struct {
	Contract     *Absorbable       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbsorbableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbsorbableCallerSession struct {
	Contract *AbsorbableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AbsorbableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbsorbableTransactorSession struct {
	Contract     *AbsorbableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AbsorbableRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbsorbableRaw struct {
	Contract *Absorbable // Generic contract binding to access the raw methods on
}

// AbsorbableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbsorbableCallerRaw struct {
	Contract *AbsorbableCaller // Generic read-only contract binding to access the raw methods on
}

// AbsorbableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbsorbableTransactorRaw struct {
	Contract *AbsorbableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbsorbable creates a new instance of Absorbable, bound to a specific deployed contract.
func NewAbsorbable(address common.Address, backend bind.ContractBackend) (*Absorbable, error) {
	contract, err := bindAbsorbable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Absorbable{AbsorbableCaller: AbsorbableCaller{contract: contract}, AbsorbableTransactor: AbsorbableTransactor{contract: contract}, AbsorbableFilterer: AbsorbableFilterer{contract: contract}}, nil
}

// NewAbsorbableCaller creates a new read-only instance of Absorbable, bound to a specific deployed contract.
func NewAbsorbableCaller(address common.Address, caller bind.ContractCaller) (*AbsorbableCaller, error) {
	contract, err := bindAbsorbable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbsorbableCaller{contract: contract}, nil
}

// NewAbsorbableTransactor creates a new write-only instance of Absorbable, bound to a specific deployed contract.
func NewAbsorbableTransactor(address common.Address, transactor bind.ContractTransactor) (*AbsorbableTransactor, error) {
	contract, err := bindAbsorbable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbsorbableTransactor{contract: contract}, nil
}

// NewAbsorbableFilterer creates a new log filterer instance of Absorbable, bound to a specific deployed contract.
func NewAbsorbableFilterer(address common.Address, filterer bind.ContractFilterer) (*AbsorbableFilterer, error) {
	contract, err := bindAbsorbable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbsorbableFilterer{contract: contract}, nil
}

// bindAbsorbable binds a generic wrapper to an already deployed contract.
func bindAbsorbable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AbsorbableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Absorbable *AbsorbableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Absorbable.Contract.AbsorbableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Absorbable *AbsorbableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Absorbable.Contract.AbsorbableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Absorbable *AbsorbableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Absorbable.Contract.AbsorbableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Absorbable *AbsorbableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Absorbable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Absorbable *AbsorbableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Absorbable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Absorbable *AbsorbableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Absorbable.Contract.contract.Transact(opts, method, params...)
}

// Ask is a free data retrieval call binding the contract method 0x69c07d31.
//
// Solidity: function Ask() view returns(bool)
func (_Absorbable *AbsorbableCaller) Ask(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Absorbable.contract.Call(opts, out, "Ask")
	return *ret0, err
}

// Ask is a free data retrieval call binding the contract method 0x69c07d31.
//
// Solidity: function Ask() view returns(bool)
func (_Absorbable *AbsorbableSession) Ask() (bool, error) {
	return _Absorbable.Contract.Ask(&_Absorbable.CallOpts)
}

// Ask is a free data retrieval call binding the contract method 0x69c07d31.
//
// Solidity: function Ask() view returns(bool)
func (_Absorbable *AbsorbableCallerSession) Ask() (bool, error) {
	return _Absorbable.Contract.Ask(&_Absorbable.CallOpts)
}

// Bid is a free data retrieval call binding the contract method 0x6e6452cb.
//
// Solidity: function Bid() view returns(bool)
func (_Absorbable *AbsorbableCaller) Bid(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Absorbable.contract.Call(opts, out, "Bid")
	return *ret0, err
}

// Bid is a free data retrieval call binding the contract method 0x6e6452cb.
//
// Solidity: function Bid() view returns(bool)
func (_Absorbable *AbsorbableSession) Bid() (bool, error) {
	return _Absorbable.Contract.Bid(&_Absorbable.CallOpts)
}

// Bid is a free data retrieval call binding the contract method 0x6e6452cb.
//
// Solidity: function Bid() view returns(bool)
func (_Absorbable *AbsorbableCallerSession) Bid() (bool, error) {
	return _Absorbable.Contract.Bid(&_Absorbable.CallOpts)
}

// CalcOrderID is a free data retrieval call binding the contract method 0xf318722b.
//
// Solidity: function calcOrderID(address maker, bytes32 index) pure returns(bytes32)
func (_Absorbable *AbsorbableCaller) CalcOrderID(opts *bind.CallOpts, maker common.Address, index [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Absorbable.contract.Call(opts, out, "calcOrderID", maker, index)
	return *ret0, err
}

// CalcOrderID is a free data retrieval call binding the contract method 0xf318722b.
//
// Solidity: function calcOrderID(address maker, bytes32 index) pure returns(bytes32)
func (_Absorbable *AbsorbableSession) CalcOrderID(maker common.Address, index [32]byte) ([32]byte, error) {
	return _Absorbable.Contract.CalcOrderID(&_Absorbable.CallOpts, maker, index)
}

// CalcOrderID is a free data retrieval call binding the contract method 0xf318722b.
//
// Solidity: function calcOrderID(address maker, bytes32 index) pure returns(bytes32)
func (_Absorbable *AbsorbableCallerSession) CalcOrderID(maker common.Address, index [32]byte) ([32]byte, error) {
	return _Absorbable.Contract.CalcOrderID(&_Absorbable.CallOpts, maker, index)
}

// FindAssistingID is a free data retrieval call binding the contract method 0xced4aac8.
//
// Solidity: function findAssistingID(bool orderType, address maker, uint256 haveAmount, uint256 wantAmount, bytes32 assistingID) view returns(bytes32)
func (_Absorbable *AbsorbableCaller) FindAssistingID(opts *bind.CallOpts, orderType bool, maker common.Address, haveAmount *big.Int, wantAmount *big.Int, assistingID [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Absorbable.contract.Call(opts, out, "findAssistingID", orderType, maker, haveAmount, wantAmount, assistingID)
	return *ret0, err
}

// FindAssistingID is a free data retrieval call binding the contract method 0xced4aac8.
//
// Solidity: function findAssistingID(bool orderType, address maker, uint256 haveAmount, uint256 wantAmount, bytes32 assistingID) view returns(bytes32)
func (_Absorbable *AbsorbableSession) FindAssistingID(orderType bool, maker common.Address, haveAmount *big.Int, wantAmount *big.Int, assistingID [32]byte) ([32]byte, error) {
	return _Absorbable.Contract.FindAssistingID(&_Absorbable.CallOpts, orderType, maker, haveAmount, wantAmount, assistingID)
}

// FindAssistingID is a free data retrieval call binding the contract method 0xced4aac8.
//
// Solidity: function findAssistingID(bool orderType, address maker, uint256 haveAmount, uint256 wantAmount, bytes32 assistingID) view returns(bytes32)
func (_Absorbable *AbsorbableCallerSession) FindAssistingID(orderType bool, maker common.Address, haveAmount *big.Int, wantAmount *big.Int, assistingID [32]byte) ([32]byte, error) {
	return _Absorbable.Contract.FindAssistingID(&_Absorbable.CallOpts, orderType, maker, haveAmount, wantAmount, assistingID)
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) view returns(address maker, uint256 have, uint256 want, bytes32 prevID, bytes32 nextID)
func (_Absorbable *AbsorbableCaller) GetOrder(opts *bind.CallOpts, _orderType bool, _id [32]byte) (struct {
	Maker  common.Address
	Have   *big.Int
	Want   *big.Int
	PrevID [32]byte
	NextID [32]byte
}, error) {
	ret := new(struct {
		Maker  common.Address
		Have   *big.Int
		Want   *big.Int
		PrevID [32]byte
		NextID [32]byte
	})
	out := ret
	err := _Absorbable.contract.Call(opts, out, "getOrder", _orderType, _id)
	return *ret, err
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) view returns(address maker, uint256 have, uint256 want, bytes32 prevID, bytes32 nextID)
func (_Absorbable *AbsorbableSession) GetOrder(_orderType bool, _id [32]byte) (struct {
	Maker  common.Address
	Have   *big.Int
	Want   *big.Int
	PrevID [32]byte
	NextID [32]byte
}, error) {
	return _Absorbable.Contract.GetOrder(&_Absorbable.CallOpts, _orderType, _id)
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) view returns(address maker, uint256 have, uint256 want, bytes32 prevID, bytes32 nextID)
func (_Absorbable *AbsorbableCallerSession) GetOrder(_orderType bool, _id [32]byte) (struct {
	Maker  common.Address
	Have   *big.Int
	Want   *big.Int
	PrevID [32]byte
	NextID [32]byte
}, error) {
	return _Absorbable.Contract.GetOrder(&_Absorbable.CallOpts, _orderType, _id)
}

// GetRemainToAbsorb is a free data retrieval call binding the contract method 0xee1a68c6.
//
// Solidity: function getRemainToAbsorb() view returns(bool, int256)
func (_Absorbable *AbsorbableCaller) GetRemainToAbsorb(opts *bind.CallOpts) (bool, *big.Int, error) {
	var (
		ret0 = new(bool)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Absorbable.contract.Call(opts, out, "getRemainToAbsorb")
	return *ret0, *ret1, err
}

// GetRemainToAbsorb is a free data retrieval call binding the contract method 0xee1a68c6.
//
// Solidity: function getRemainToAbsorb() view returns(bool, int256)
func (_Absorbable *AbsorbableSession) GetRemainToAbsorb() (bool, *big.Int, error) {
	return _Absorbable.Contract.GetRemainToAbsorb(&_Absorbable.CallOpts)
}

// GetRemainToAbsorb is a free data retrieval call binding the contract method 0xee1a68c6.
//
// Solidity: function getRemainToAbsorb() view returns(bool, int256)
func (_Absorbable *AbsorbableCallerSession) GetRemainToAbsorb() (bool, *big.Int, error) {
	return _Absorbable.Contract.GetRemainToAbsorb(&_Absorbable.CallOpts)
}

// Next is a free data retrieval call binding the contract method 0x4ea09797.
//
// Solidity: function next(bool orderType, bytes32 id) view returns(bytes32)
func (_Absorbable *AbsorbableCaller) Next(opts *bind.CallOpts, orderType bool, id [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Absorbable.contract.Call(opts, out, "next", orderType, id)
	return *ret0, err
}

// Next is a free data retrieval call binding the contract method 0x4ea09797.
//
// Solidity: function next(bool orderType, bytes32 id) view returns(bytes32)
func (_Absorbable *AbsorbableSession) Next(orderType bool, id [32]byte) ([32]byte, error) {
	return _Absorbable.Contract.Next(&_Absorbable.CallOpts, orderType, id)
}

// Next is a free data retrieval call binding the contract method 0x4ea09797.
//
// Solidity: function next(bool orderType, bytes32 id) view returns(bytes32)
func (_Absorbable *AbsorbableCallerSession) Next(orderType bool, id [32]byte) ([32]byte, error) {
	return _Absorbable.Contract.Next(&_Absorbable.CallOpts, orderType, id)
}

// Prev is a free data retrieval call binding the contract method 0x0d90b10a.
//
// Solidity: function prev(bool orderType, bytes32 id) view returns(bytes32)
func (_Absorbable *AbsorbableCaller) Prev(opts *bind.CallOpts, orderType bool, id [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Absorbable.contract.Call(opts, out, "prev", orderType, id)
	return *ret0, err
}

// Prev is a free data retrieval call binding the contract method 0x0d90b10a.
//
// Solidity: function prev(bool orderType, bytes32 id) view returns(bytes32)
func (_Absorbable *AbsorbableSession) Prev(orderType bool, id [32]byte) ([32]byte, error) {
	return _Absorbable.Contract.Prev(&_Absorbable.CallOpts, orderType, id)
}

// Prev is a free data retrieval call binding the contract method 0x0d90b10a.
//
// Solidity: function prev(bool orderType, bytes32 id) view returns(bytes32)
func (_Absorbable *AbsorbableCallerSession) Prev(orderType bool, id [32]byte) ([32]byte, error) {
	return _Absorbable.Contract.Prev(&_Absorbable.CallOpts, orderType, id)
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool orderType) view returns(bytes32)
func (_Absorbable *AbsorbableCaller) Top(opts *bind.CallOpts, orderType bool) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Absorbable.contract.Call(opts, out, "top", orderType)
	return *ret0, err
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool orderType) view returns(bytes32)
func (_Absorbable *AbsorbableSession) Top(orderType bool) ([32]byte, error) {
	return _Absorbable.Contract.Top(&_Absorbable.CallOpts, orderType)
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool orderType) view returns(bytes32)
func (_Absorbable *AbsorbableCallerSession) Top(orderType bool) ([32]byte, error) {
	return _Absorbable.Contract.Top(&_Absorbable.CallOpts, orderType)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool orderType, bytes32 id) returns()
func (_Absorbable *AbsorbableTransactor) Cancel(opts *bind.TransactOpts, orderType bool, id [32]byte) (*types.Transaction, error) {
	return _Absorbable.contract.Transact(opts, "cancel", orderType, id)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool orderType, bytes32 id) returns()
func (_Absorbable *AbsorbableSession) Cancel(orderType bool, id [32]byte) (*types.Transaction, error) {
	return _Absorbable.Contract.Cancel(&_Absorbable.TransactOpts, orderType, id)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool orderType, bytes32 id) returns()
func (_Absorbable *AbsorbableTransactorSession) Cancel(orderType bool, id [32]byte) (*types.Transaction, error) {
	return _Absorbable.Contract.Cancel(&_Absorbable.TransactOpts, orderType, id)
}

// OnBlockInitialized is a paid mutator transaction binding the contract method 0xbe91d729.
//
// Solidity: function onBlockInitialized(uint256 target) returns()
func (_Absorbable *AbsorbableTransactor) OnBlockInitialized(opts *bind.TransactOpts, target *big.Int) (*types.Transaction, error) {
	return _Absorbable.contract.Transact(opts, "onBlockInitialized", target)
}

// OnBlockInitialized is a paid mutator transaction binding the contract method 0xbe91d729.
//
// Solidity: function onBlockInitialized(uint256 target) returns()
func (_Absorbable *AbsorbableSession) OnBlockInitialized(target *big.Int) (*types.Transaction, error) {
	return _Absorbable.Contract.OnBlockInitialized(&_Absorbable.TransactOpts, target)
}

// OnBlockInitialized is a paid mutator transaction binding the contract method 0xbe91d729.
//
// Solidity: function onBlockInitialized(uint256 target) returns()
func (_Absorbable *AbsorbableTransactorSession) OnBlockInitialized(target *big.Int) (*types.Transaction, error) {
	return _Absorbable.Contract.OnBlockInitialized(&_Absorbable.TransactOpts, target)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xaa1c259c.
//
// Solidity: function registerTokens(address volatileToken, address stablizeToken) returns()
func (_Absorbable *AbsorbableTransactor) RegisterTokens(opts *bind.TransactOpts, volatileToken common.Address, stablizeToken common.Address) (*types.Transaction, error) {
	return _Absorbable.contract.Transact(opts, "registerTokens", volatileToken, stablizeToken)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xaa1c259c.
//
// Solidity: function registerTokens(address volatileToken, address stablizeToken) returns()
func (_Absorbable *AbsorbableSession) RegisterTokens(volatileToken common.Address, stablizeToken common.Address) (*types.Transaction, error) {
	return _Absorbable.Contract.RegisterTokens(&_Absorbable.TransactOpts, volatileToken, stablizeToken)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xaa1c259c.
//
// Solidity: function registerTokens(address volatileToken, address stablizeToken) returns()
func (_Absorbable *AbsorbableTransactorSession) RegisterTokens(volatileToken common.Address, stablizeToken common.Address) (*types.Transaction, error) {
	return _Absorbable.Contract.RegisterTokens(&_Absorbable.TransactOpts, volatileToken, stablizeToken)
}

// AbsorbableAbsorptionIterator is returned from FilterAbsorption and is used to iterate over the raw logs and unpacked data for Absorption events raised by the Absorbable contract.
type AbsorbableAbsorptionIterator struct {
	Event *AbsorbableAbsorption // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbsorbableAbsorptionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbsorbableAbsorption)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbsorbableAbsorption)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbsorbableAbsorptionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbsorbableAbsorptionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbsorbableAbsorption represents a Absorption event raised by the Absorbable contract.
type AbsorbableAbsorption struct {
	Amount  *big.Int
	Supply  *big.Int
	Emptive bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAbsorption is a free log retrieval operation binding the contract event 0x0427b353dc7214e3d8c7f5039475a8e729f4d62922937381e304cd03becf66d2.
//
// Solidity: event Absorption(int256 amount, uint256 supply, bool emptive)
func (_Absorbable *AbsorbableFilterer) FilterAbsorption(opts *bind.FilterOpts) (*AbsorbableAbsorptionIterator, error) {

	logs, sub, err := _Absorbable.contract.FilterLogs(opts, "Absorption")
	if err != nil {
		return nil, err
	}
	return &AbsorbableAbsorptionIterator{contract: _Absorbable.contract, event: "Absorption", logs: logs, sub: sub}, nil
}

// WatchAbsorption is a free log subscription operation binding the contract event 0x0427b353dc7214e3d8c7f5039475a8e729f4d62922937381e304cd03becf66d2.
//
// Solidity: event Absorption(int256 amount, uint256 supply, bool emptive)
func (_Absorbable *AbsorbableFilterer) WatchAbsorption(opts *bind.WatchOpts, sink chan<- *AbsorbableAbsorption) (event.Subscription, error) {

	logs, sub, err := _Absorbable.contract.WatchLogs(opts, "Absorption")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbsorbableAbsorption)
				if err := _Absorbable.contract.UnpackLog(event, "Absorption", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAbsorption is a log parse operation binding the contract event 0x0427b353dc7214e3d8c7f5039475a8e729f4d62922937381e304cd03becf66d2.
//
// Solidity: event Absorption(int256 amount, uint256 supply, bool emptive)
func (_Absorbable *AbsorbableFilterer) ParseAbsorption(log types.Log) (*AbsorbableAbsorption, error) {
	event := new(AbsorbableAbsorption)
	if err := _Absorbable.contract.UnpackLog(event, "Absorption", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AbsorbableStopIterator is returned from FilterStop and is used to iterate over the raw logs and unpacked data for Stop events raised by the Absorbable contract.
type AbsorbableStopIterator struct {
	Event *AbsorbableStop // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbsorbableStopIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbsorbableStop)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbsorbableStop)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbsorbableStopIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbsorbableStopIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbsorbableStop represents a Stop event raised by the Absorbable contract.
type AbsorbableStop struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStop is a free log retrieval operation binding the contract event 0xbedf0f4abfe86d4ffad593d9607fe70e83ea706033d44d24b3b6283cf3fc4f6b.
//
// Solidity: event Stop()
func (_Absorbable *AbsorbableFilterer) FilterStop(opts *bind.FilterOpts) (*AbsorbableStopIterator, error) {

	logs, sub, err := _Absorbable.contract.FilterLogs(opts, "Stop")
	if err != nil {
		return nil, err
	}
	return &AbsorbableStopIterator{contract: _Absorbable.contract, event: "Stop", logs: logs, sub: sub}, nil
}

// WatchStop is a free log subscription operation binding the contract event 0xbedf0f4abfe86d4ffad593d9607fe70e83ea706033d44d24b3b6283cf3fc4f6b.
//
// Solidity: event Stop()
func (_Absorbable *AbsorbableFilterer) WatchStop(opts *bind.WatchOpts, sink chan<- *AbsorbableStop) (event.Subscription, error) {

	logs, sub, err := _Absorbable.contract.WatchLogs(opts, "Stop")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbsorbableStop)
				if err := _Absorbable.contract.UnpackLog(event, "Stop", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStop is a log parse operation binding the contract event 0xbedf0f4abfe86d4ffad593d9607fe70e83ea706033d44d24b3b6283cf3fc4f6b.
//
// Solidity: event Stop()
func (_Absorbable *AbsorbableFilterer) ParseStop(log types.Log) (*AbsorbableStop, error) {
	event := new(AbsorbableStop)
	if err := _Absorbable.contract.UnpackLog(event, "Stop", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ITokenABI is the input ABI used to generate the binding from.
const ITokenABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"dexBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"dexMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferToDex\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ITokenFuncSigs maps the 4-byte function signature to its string representation.
var ITokenFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"692058c2": "dex()",
	"117f5a55": "dexBurn(uint256)",
	"bdfde911": "dexMint(uint256)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"6fdb0eb6": "transferToDex(address,uint256)",
}

// IToken is an auto generated Go binding around an Ethereum contract.
type IToken struct {
	ITokenCaller     // Read-only binding to the contract
	ITokenTransactor // Write-only binding to the contract
	ITokenFilterer   // Log filterer for contract events
}

// ITokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITokenSession struct {
	Contract     *IToken           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITokenCallerSession struct {
	Contract *ITokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ITokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITokenTransactorSession struct {
	Contract     *ITokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITokenRaw struct {
	Contract *IToken // Generic contract binding to access the raw methods on
}

// ITokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITokenCallerRaw struct {
	Contract *ITokenCaller // Generic read-only contract binding to access the raw methods on
}

// ITokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITokenTransactorRaw struct {
	Contract *ITokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIToken creates a new instance of IToken, bound to a specific deployed contract.
func NewIToken(address common.Address, backend bind.ContractBackend) (*IToken, error) {
	contract, err := bindIToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IToken{ITokenCaller: ITokenCaller{contract: contract}, ITokenTransactor: ITokenTransactor{contract: contract}, ITokenFilterer: ITokenFilterer{contract: contract}}, nil
}

// NewITokenCaller creates a new read-only instance of IToken, bound to a specific deployed contract.
func NewITokenCaller(address common.Address, caller bind.ContractCaller) (*ITokenCaller, error) {
	contract, err := bindIToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenCaller{contract: contract}, nil
}

// NewITokenTransactor creates a new write-only instance of IToken, bound to a specific deployed contract.
func NewITokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ITokenTransactor, error) {
	contract, err := bindIToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenTransactor{contract: contract}, nil
}

// NewITokenFilterer creates a new log filterer instance of IToken, bound to a specific deployed contract.
func NewITokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ITokenFilterer, error) {
	contract, err := bindIToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITokenFilterer{contract: contract}, nil
}

// bindIToken binds a generic wrapper to an already deployed contract.
func bindIToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ITokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IToken *ITokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IToken.Contract.ITokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IToken *ITokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IToken.Contract.ITokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IToken *ITokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IToken.Contract.ITokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IToken *ITokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IToken *ITokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IToken *ITokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IToken *ITokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IToken.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IToken *ITokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IToken.Contract.Allowance(&_IToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IToken *ITokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IToken.Contract.Allowance(&_IToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_IToken *ITokenCaller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IToken.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_IToken *ITokenSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _IToken.Contract.BalanceOf(&_IToken.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_IToken *ITokenCallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _IToken.Contract.BalanceOf(&_IToken.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IToken *ITokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IToken *ITokenSession) TotalSupply() (*big.Int, error) {
	return _IToken.Contract.TotalSupply(&_IToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IToken *ITokenCallerSession) TotalSupply() (*big.Int, error) {
	return _IToken.Contract.TotalSupply(&_IToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IToken *ITokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IToken *ITokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.Approve(&_IToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IToken *ITokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.Approve(&_IToken.TransactOpts, spender, value)
}

// Dex is a paid mutator transaction binding the contract method 0x692058c2.
//
// Solidity: function dex() returns(address)
func (_IToken *ITokenTransactor) Dex(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "dex")
}

// Dex is a paid mutator transaction binding the contract method 0x692058c2.
//
// Solidity: function dex() returns(address)
func (_IToken *ITokenSession) Dex() (*types.Transaction, error) {
	return _IToken.Contract.Dex(&_IToken.TransactOpts)
}

// Dex is a paid mutator transaction binding the contract method 0x692058c2.
//
// Solidity: function dex() returns(address)
func (_IToken *ITokenTransactorSession) Dex() (*types.Transaction, error) {
	return _IToken.Contract.Dex(&_IToken.TransactOpts)
}

// DexBurn is a paid mutator transaction binding the contract method 0x117f5a55.
//
// Solidity: function dexBurn(uint256 value) returns()
func (_IToken *ITokenTransactor) DexBurn(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "dexBurn", value)
}

// DexBurn is a paid mutator transaction binding the contract method 0x117f5a55.
//
// Solidity: function dexBurn(uint256 value) returns()
func (_IToken *ITokenSession) DexBurn(value *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.DexBurn(&_IToken.TransactOpts, value)
}

// DexBurn is a paid mutator transaction binding the contract method 0x117f5a55.
//
// Solidity: function dexBurn(uint256 value) returns()
func (_IToken *ITokenTransactorSession) DexBurn(value *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.DexBurn(&_IToken.TransactOpts, value)
}

// DexMint is a paid mutator transaction binding the contract method 0xbdfde911.
//
// Solidity: function dexMint(uint256 value) returns()
func (_IToken *ITokenTransactor) DexMint(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "dexMint", value)
}

// DexMint is a paid mutator transaction binding the contract method 0xbdfde911.
//
// Solidity: function dexMint(uint256 value) returns()
func (_IToken *ITokenSession) DexMint(value *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.DexMint(&_IToken.TransactOpts, value)
}

// DexMint is a paid mutator transaction binding the contract method 0xbdfde911.
//
// Solidity: function dexMint(uint256 value) returns()
func (_IToken *ITokenTransactorSession) DexMint(value *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.DexMint(&_IToken.TransactOpts, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IToken *ITokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IToken *ITokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.Transfer(&_IToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IToken *ITokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.Transfer(&_IToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IToken *ITokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IToken *ITokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.TransferFrom(&_IToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IToken *ITokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.TransferFrom(&_IToken.TransactOpts, from, to, value)
}

// TransferToDex is a paid mutator transaction binding the contract method 0x6fdb0eb6.
//
// Solidity: function transferToDex(address from, uint256 value) returns()
func (_IToken *ITokenTransactor) TransferToDex(opts *bind.TransactOpts, from common.Address, value *big.Int) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "transferToDex", from, value)
}

// TransferToDex is a paid mutator transaction binding the contract method 0x6fdb0eb6.
//
// Solidity: function transferToDex(address from, uint256 value) returns()
func (_IToken *ITokenSession) TransferToDex(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.TransferToDex(&_IToken.TransactOpts, from, value)
}

// TransferToDex is a paid mutator transaction binding the contract method 0x6fdb0eb6.
//
// Solidity: function transferToDex(address from, uint256 value) returns()
func (_IToken *ITokenTransactorSession) TransferToDex(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.TransferToDex(&_IToken.TransactOpts, from, value)
}

// ITokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IToken contract.
type ITokenApprovalIterator struct {
	Event *ITokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ITokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITokenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ITokenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ITokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITokenApproval represents a Approval event raised by the IToken contract.
type ITokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IToken *ITokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ITokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ITokenApprovalIterator{contract: _IToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IToken *ITokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ITokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITokenApproval)
				if err := _IToken.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IToken *ITokenFilterer) ParseApproval(log types.Log) (*ITokenApproval, error) {
	event := new(ITokenApproval)
	if err := _IToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ITokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IToken contract.
type ITokenTransferIterator struct {
	Event *ITokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ITokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ITokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ITokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITokenTransfer represents a Transfer event raised by the IToken contract.
type ITokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IToken *ITokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ITokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ITokenTransferIterator{contract: _IToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IToken *ITokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ITokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITokenTransfer)
				if err := _IToken.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IToken *ITokenFilterer) ParseTransfer(log types.Log) (*ITokenTransfer, error) {
	event := new(ITokenTransfer)
	if err := _IToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OrderbookABI is the input ABI used to generate the binding from.
const OrderbookABI = "[{\"inputs\":[],\"name\":\"Ask\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Bid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"index\",\"type\":\"bytes32\"}],\"name\":\"calcOrderID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"haveAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wantAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"assistingID\",\"type\":\"bytes32\"}],\"name\":\"findAssistingID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getOrder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"have\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"prevID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextID\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"next\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"prev\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"volatileToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stablizeToken\",\"type\":\"address\"}],\"name\":\"registerTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"}],\"name\":\"top\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// OrderbookFuncSigs maps the 4-byte function signature to its string representation.
var OrderbookFuncSigs = map[string]string{
	"69c07d31": "Ask()",
	"6e6452cb": "Bid()",
	"f318722b": "calcOrderID(address,bytes32)",
	"43271d79": "cancel(bool,bytes32)",
	"ced4aac8": "findAssistingID(bool,address,uint256,uint256,bytes32)",
	"07c399a3": "getOrder(bool,bytes32)",
	"4ea09797": "next(bool,bytes32)",
	"0d90b10a": "prev(bool,bytes32)",
	"aa1c259c": "registerTokens(address,address)",
	"8aa3f897": "top(bool)",
}

// OrderbookBin is the compiled bytecode used for deploying new contracts.
var OrderbookBin = "0x608060405234801561001057600080fd5b5061096f806100206000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80636e6452cb116100665780636e6452cb1461019c5780638aa3f897146101a4578063aa1c259c146101c3578063ced4aac8146101f1578063f318722b146102315761009e565b806307c399a3146100a35780630d90b10a146100fd57806343271d79146101345780634ea097971461015b57806369c07d3114610180575b600080fd5b6100c8600480360360408110156100b957600080fd5b5080351515906020013561025d565b604080516001600160a01b03909616865260208601949094528484019290925260608401526080830152519081900360a00190f35b6101226004803603604081101561011357600080fd5b508035151590602001356102a4565b60408051918252519081900360200190f35b6101596004803603604081101561014a57600080fd5b508035151590602001356102c8565b005b6101226004803603604081101561017157600080fd5b5080351515906020013561034e565b610188610372565b604080519115158252519081900360200190f35b610188610377565b610122600480360360208110156101ba57600080fd5b5035151561037c565b610159600480360360408110156101d957600080fd5b506001600160a01b038135811691602001351661039d565b610122600480360360a081101561020757600080fd5b5080351515906001600160a01b036020820135169060408101359060608101359060800135610415565b6101226004803603604081101561024757600080fd5b506001600160a01b03813516906020013561047d565b90151560009081526020818152604080832093835260029384019091529020805460018201549282015460038301546004909301546001600160a01b039092169490929190565b90151560009081526020818152604080832093835260029093019052206003015490565b8115156000908152602081815260408083208484526002810190925290912080546001600160a01b03163314610338576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037b93232b91036b0b5b2b960811b604482015290519081900360640190fd5b610348828463ffffffff61048916565b50505050565b90151560009081526020818152604080832093835260029093019052206004015490565b600081565b600181565b801515600090815260208190526040812061039681610588565b9392505050565b60008080526020526103d67fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5838363ffffffff6105a116565b600160009081526020526104117fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d828463ffffffff6105a116565b5050565b841515600090815260208190526040812061042e61090b565b506040805160a0810182526001600160a01b038816815260208101879052908101859052600060608201819052608082015261047182828663ffffffff61074e16565b98975050505050505050565b600061039683836107c0565b61049161090b565b50600081815260028084016020908152604092839020835160a08101855281546001600160a01b0316815260018201549281018390529281015493830193909352600383015460608301526004909201546080820152901561057357825481516020808401516040805163a9059cbb60e01b81526001600160a01b039485166004820152602481019290925251929093169263a9059cbb92604480830193928290030181600087803b15801561054657600080fd5b505af115801561055a573d6000803e3d6000fd5b505050506040513d602081101561057057600080fd5b50505b610583838363ffffffff61088d16565b505050565b6000808052600282016020526040902060040154919050565b818360000160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550808360010160006101000a8154816001600160a01b0302191690836001600160a01b031602179055506040518060a00160405280306001600160a01b0316815260200160008152602001600081526020016000801b815260200160001960001b8152508360020160008060001b815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550602082015181600101556040820151816002015560608201518160030155608082015181600401559050506040518060a00160405280306001600160a01b0316815260200160008152602001600181526020016000801b815260200160001960001b81525083600201600060001960001b815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060208201518160010155604082015181600201556060820151816003015560808201518160040155905050505050565b600081815260028401602052604081205b6004015460008181526002860160205260409020909250610786848263ffffffff6108ed16565b1561075f575b60030154600081815260028601602052604090209092506107b3848263ffffffff6108ed16565b61078c5750909392505050565b60006002838360405160200180836001600160a01b03166001600160a01b031660601b8152601401828152602001925050506040516020818303038152906040526040518082805190602001908083835b602083106108305780518252601f199092019160209182019101610811565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa15801561086f573d6000803e3d6000fd5b5050506040513d602081101561088457600080fd5b50519392505050565b6000818152600292830160205260408082206004808201805460038085018054885286882090940182905583549187529486209094019390935593835280546001600160a01b031916815560018101839055909301819055908190559055565b60408201516001820154600290920154602090930151910291021190565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091529056fea2646970667358221220f7ab716a9394928cd4ad697268d0ccc87bbbdc20e195982aec8a9cd3376eb19964736f6c63430006080033"

// DeployOrderbook deploys a new Ethereum contract, binding an instance of Orderbook to it.
func DeployOrderbook(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Orderbook, error) {
	parsed, err := abi.JSON(strings.NewReader(OrderbookABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OrderbookBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Orderbook{OrderbookCaller: OrderbookCaller{contract: contract}, OrderbookTransactor: OrderbookTransactor{contract: contract}, OrderbookFilterer: OrderbookFilterer{contract: contract}}, nil
}

// Orderbook is an auto generated Go binding around an Ethereum contract.
type Orderbook struct {
	OrderbookCaller     // Read-only binding to the contract
	OrderbookTransactor // Write-only binding to the contract
	OrderbookFilterer   // Log filterer for contract events
}

// OrderbookCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrderbookCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderbookTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrderbookTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderbookFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrderbookFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderbookSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrderbookSession struct {
	Contract     *Orderbook        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrderbookCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrderbookCallerSession struct {
	Contract *OrderbookCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// OrderbookTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrderbookTransactorSession struct {
	Contract     *OrderbookTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OrderbookRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrderbookRaw struct {
	Contract *Orderbook // Generic contract binding to access the raw methods on
}

// OrderbookCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrderbookCallerRaw struct {
	Contract *OrderbookCaller // Generic read-only contract binding to access the raw methods on
}

// OrderbookTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrderbookTransactorRaw struct {
	Contract *OrderbookTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrderbook creates a new instance of Orderbook, bound to a specific deployed contract.
func NewOrderbook(address common.Address, backend bind.ContractBackend) (*Orderbook, error) {
	contract, err := bindOrderbook(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Orderbook{OrderbookCaller: OrderbookCaller{contract: contract}, OrderbookTransactor: OrderbookTransactor{contract: contract}, OrderbookFilterer: OrderbookFilterer{contract: contract}}, nil
}

// NewOrderbookCaller creates a new read-only instance of Orderbook, bound to a specific deployed contract.
func NewOrderbookCaller(address common.Address, caller bind.ContractCaller) (*OrderbookCaller, error) {
	contract, err := bindOrderbook(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrderbookCaller{contract: contract}, nil
}

// NewOrderbookTransactor creates a new write-only instance of Orderbook, bound to a specific deployed contract.
func NewOrderbookTransactor(address common.Address, transactor bind.ContractTransactor) (*OrderbookTransactor, error) {
	contract, err := bindOrderbook(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrderbookTransactor{contract: contract}, nil
}

// NewOrderbookFilterer creates a new log filterer instance of Orderbook, bound to a specific deployed contract.
func NewOrderbookFilterer(address common.Address, filterer bind.ContractFilterer) (*OrderbookFilterer, error) {
	contract, err := bindOrderbook(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrderbookFilterer{contract: contract}, nil
}

// bindOrderbook binds a generic wrapper to an already deployed contract.
func bindOrderbook(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OrderbookABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Orderbook *OrderbookRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Orderbook.Contract.OrderbookCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Orderbook *OrderbookRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Orderbook.Contract.OrderbookTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Orderbook *OrderbookRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Orderbook.Contract.OrderbookTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Orderbook *OrderbookCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Orderbook.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Orderbook *OrderbookTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Orderbook.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Orderbook *OrderbookTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Orderbook.Contract.contract.Transact(opts, method, params...)
}

// Ask is a free data retrieval call binding the contract method 0x69c07d31.
//
// Solidity: function Ask() view returns(bool)
func (_Orderbook *OrderbookCaller) Ask(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "Ask")
	return *ret0, err
}

// Ask is a free data retrieval call binding the contract method 0x69c07d31.
//
// Solidity: function Ask() view returns(bool)
func (_Orderbook *OrderbookSession) Ask() (bool, error) {
	return _Orderbook.Contract.Ask(&_Orderbook.CallOpts)
}

// Ask is a free data retrieval call binding the contract method 0x69c07d31.
//
// Solidity: function Ask() view returns(bool)
func (_Orderbook *OrderbookCallerSession) Ask() (bool, error) {
	return _Orderbook.Contract.Ask(&_Orderbook.CallOpts)
}

// Bid is a free data retrieval call binding the contract method 0x6e6452cb.
//
// Solidity: function Bid() view returns(bool)
func (_Orderbook *OrderbookCaller) Bid(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "Bid")
	return *ret0, err
}

// Bid is a free data retrieval call binding the contract method 0x6e6452cb.
//
// Solidity: function Bid() view returns(bool)
func (_Orderbook *OrderbookSession) Bid() (bool, error) {
	return _Orderbook.Contract.Bid(&_Orderbook.CallOpts)
}

// Bid is a free data retrieval call binding the contract method 0x6e6452cb.
//
// Solidity: function Bid() view returns(bool)
func (_Orderbook *OrderbookCallerSession) Bid() (bool, error) {
	return _Orderbook.Contract.Bid(&_Orderbook.CallOpts)
}

// CalcOrderID is a free data retrieval call binding the contract method 0xf318722b.
//
// Solidity: function calcOrderID(address maker, bytes32 index) pure returns(bytes32)
func (_Orderbook *OrderbookCaller) CalcOrderID(opts *bind.CallOpts, maker common.Address, index [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "calcOrderID", maker, index)
	return *ret0, err
}

// CalcOrderID is a free data retrieval call binding the contract method 0xf318722b.
//
// Solidity: function calcOrderID(address maker, bytes32 index) pure returns(bytes32)
func (_Orderbook *OrderbookSession) CalcOrderID(maker common.Address, index [32]byte) ([32]byte, error) {
	return _Orderbook.Contract.CalcOrderID(&_Orderbook.CallOpts, maker, index)
}

// CalcOrderID is a free data retrieval call binding the contract method 0xf318722b.
//
// Solidity: function calcOrderID(address maker, bytes32 index) pure returns(bytes32)
func (_Orderbook *OrderbookCallerSession) CalcOrderID(maker common.Address, index [32]byte) ([32]byte, error) {
	return _Orderbook.Contract.CalcOrderID(&_Orderbook.CallOpts, maker, index)
}

// FindAssistingID is a free data retrieval call binding the contract method 0xced4aac8.
//
// Solidity: function findAssistingID(bool orderType, address maker, uint256 haveAmount, uint256 wantAmount, bytes32 assistingID) view returns(bytes32)
func (_Orderbook *OrderbookCaller) FindAssistingID(opts *bind.CallOpts, orderType bool, maker common.Address, haveAmount *big.Int, wantAmount *big.Int, assistingID [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "findAssistingID", orderType, maker, haveAmount, wantAmount, assistingID)
	return *ret0, err
}

// FindAssistingID is a free data retrieval call binding the contract method 0xced4aac8.
//
// Solidity: function findAssistingID(bool orderType, address maker, uint256 haveAmount, uint256 wantAmount, bytes32 assistingID) view returns(bytes32)
func (_Orderbook *OrderbookSession) FindAssistingID(orderType bool, maker common.Address, haveAmount *big.Int, wantAmount *big.Int, assistingID [32]byte) ([32]byte, error) {
	return _Orderbook.Contract.FindAssistingID(&_Orderbook.CallOpts, orderType, maker, haveAmount, wantAmount, assistingID)
}

// FindAssistingID is a free data retrieval call binding the contract method 0xced4aac8.
//
// Solidity: function findAssistingID(bool orderType, address maker, uint256 haveAmount, uint256 wantAmount, bytes32 assistingID) view returns(bytes32)
func (_Orderbook *OrderbookCallerSession) FindAssistingID(orderType bool, maker common.Address, haveAmount *big.Int, wantAmount *big.Int, assistingID [32]byte) ([32]byte, error) {
	return _Orderbook.Contract.FindAssistingID(&_Orderbook.CallOpts, orderType, maker, haveAmount, wantAmount, assistingID)
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) view returns(address maker, uint256 have, uint256 want, bytes32 prevID, bytes32 nextID)
func (_Orderbook *OrderbookCaller) GetOrder(opts *bind.CallOpts, _orderType bool, _id [32]byte) (struct {
	Maker  common.Address
	Have   *big.Int
	Want   *big.Int
	PrevID [32]byte
	NextID [32]byte
}, error) {
	ret := new(struct {
		Maker  common.Address
		Have   *big.Int
		Want   *big.Int
		PrevID [32]byte
		NextID [32]byte
	})
	out := ret
	err := _Orderbook.contract.Call(opts, out, "getOrder", _orderType, _id)
	return *ret, err
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) view returns(address maker, uint256 have, uint256 want, bytes32 prevID, bytes32 nextID)
func (_Orderbook *OrderbookSession) GetOrder(_orderType bool, _id [32]byte) (struct {
	Maker  common.Address
	Have   *big.Int
	Want   *big.Int
	PrevID [32]byte
	NextID [32]byte
}, error) {
	return _Orderbook.Contract.GetOrder(&_Orderbook.CallOpts, _orderType, _id)
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) view returns(address maker, uint256 have, uint256 want, bytes32 prevID, bytes32 nextID)
func (_Orderbook *OrderbookCallerSession) GetOrder(_orderType bool, _id [32]byte) (struct {
	Maker  common.Address
	Have   *big.Int
	Want   *big.Int
	PrevID [32]byte
	NextID [32]byte
}, error) {
	return _Orderbook.Contract.GetOrder(&_Orderbook.CallOpts, _orderType, _id)
}

// Next is a free data retrieval call binding the contract method 0x4ea09797.
//
// Solidity: function next(bool orderType, bytes32 id) view returns(bytes32)
func (_Orderbook *OrderbookCaller) Next(opts *bind.CallOpts, orderType bool, id [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "next", orderType, id)
	return *ret0, err
}

// Next is a free data retrieval call binding the contract method 0x4ea09797.
//
// Solidity: function next(bool orderType, bytes32 id) view returns(bytes32)
func (_Orderbook *OrderbookSession) Next(orderType bool, id [32]byte) ([32]byte, error) {
	return _Orderbook.Contract.Next(&_Orderbook.CallOpts, orderType, id)
}

// Next is a free data retrieval call binding the contract method 0x4ea09797.
//
// Solidity: function next(bool orderType, bytes32 id) view returns(bytes32)
func (_Orderbook *OrderbookCallerSession) Next(orderType bool, id [32]byte) ([32]byte, error) {
	return _Orderbook.Contract.Next(&_Orderbook.CallOpts, orderType, id)
}

// Prev is a free data retrieval call binding the contract method 0x0d90b10a.
//
// Solidity: function prev(bool orderType, bytes32 id) view returns(bytes32)
func (_Orderbook *OrderbookCaller) Prev(opts *bind.CallOpts, orderType bool, id [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "prev", orderType, id)
	return *ret0, err
}

// Prev is a free data retrieval call binding the contract method 0x0d90b10a.
//
// Solidity: function prev(bool orderType, bytes32 id) view returns(bytes32)
func (_Orderbook *OrderbookSession) Prev(orderType bool, id [32]byte) ([32]byte, error) {
	return _Orderbook.Contract.Prev(&_Orderbook.CallOpts, orderType, id)
}

// Prev is a free data retrieval call binding the contract method 0x0d90b10a.
//
// Solidity: function prev(bool orderType, bytes32 id) view returns(bytes32)
func (_Orderbook *OrderbookCallerSession) Prev(orderType bool, id [32]byte) ([32]byte, error) {
	return _Orderbook.Contract.Prev(&_Orderbook.CallOpts, orderType, id)
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool orderType) view returns(bytes32)
func (_Orderbook *OrderbookCaller) Top(opts *bind.CallOpts, orderType bool) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "top", orderType)
	return *ret0, err
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool orderType) view returns(bytes32)
func (_Orderbook *OrderbookSession) Top(orderType bool) ([32]byte, error) {
	return _Orderbook.Contract.Top(&_Orderbook.CallOpts, orderType)
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool orderType) view returns(bytes32)
func (_Orderbook *OrderbookCallerSession) Top(orderType bool) ([32]byte, error) {
	return _Orderbook.Contract.Top(&_Orderbook.CallOpts, orderType)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool orderType, bytes32 id) returns()
func (_Orderbook *OrderbookTransactor) Cancel(opts *bind.TransactOpts, orderType bool, id [32]byte) (*types.Transaction, error) {
	return _Orderbook.contract.Transact(opts, "cancel", orderType, id)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool orderType, bytes32 id) returns()
func (_Orderbook *OrderbookSession) Cancel(orderType bool, id [32]byte) (*types.Transaction, error) {
	return _Orderbook.Contract.Cancel(&_Orderbook.TransactOpts, orderType, id)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool orderType, bytes32 id) returns()
func (_Orderbook *OrderbookTransactorSession) Cancel(orderType bool, id [32]byte) (*types.Transaction, error) {
	return _Orderbook.Contract.Cancel(&_Orderbook.TransactOpts, orderType, id)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xaa1c259c.
//
// Solidity: function registerTokens(address volatileToken, address stablizeToken) returns()
func (_Orderbook *OrderbookTransactor) RegisterTokens(opts *bind.TransactOpts, volatileToken common.Address, stablizeToken common.Address) (*types.Transaction, error) {
	return _Orderbook.contract.Transact(opts, "registerTokens", volatileToken, stablizeToken)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xaa1c259c.
//
// Solidity: function registerTokens(address volatileToken, address stablizeToken) returns()
func (_Orderbook *OrderbookSession) RegisterTokens(volatileToken common.Address, stablizeToken common.Address) (*types.Transaction, error) {
	return _Orderbook.Contract.RegisterTokens(&_Orderbook.TransactOpts, volatileToken, stablizeToken)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xaa1c259c.
//
// Solidity: function registerTokens(address volatileToken, address stablizeToken) returns()
func (_Orderbook *OrderbookTransactorSession) RegisterTokens(volatileToken common.Address, stablizeToken common.Address) (*types.Transaction, error) {
	return _Orderbook.Contract.RegisterTokens(&_Orderbook.TransactOpts, volatileToken, stablizeToken)
}

// PreemptivableABI is the input ABI used to generate the binding from.
const PreemptivableABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"absorptionPace\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"absorptionExpiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialSlashingPace\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialLockdownExpiration\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"emptive\",\"type\":\"bool\"}],\"name\":\"Absorption\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockdownExpiration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockNumber\",\"type\":\"uint256\"}],\"name\":\"Preemptive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockdownExpiration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashingRate\",\"type\":\"uint256\"}],\"name\":\"Propose\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"}],\"name\":\"Revoke\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Slash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Stop\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"}],\"name\":\"Unlock\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"Ask\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Bid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"index\",\"type\":\"bytes32\"}],\"name\":\"calcOrderID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"haveAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wantAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"assistingID\",\"type\":\"bytes32\"}],\"name\":\"findAssistingID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGlobalParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slashingRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockdownExpiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rank\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLockdown\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"slashingFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getOrder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"have\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"prevID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextID\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getProposal\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"slashingRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockdownExpiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProposalCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRemainToAbsorb\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"next\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"target\",\"type\":\"uint256\"}],\"name\":\"onBlockInitialized\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"prev\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"volatileToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stablizeToken\",\"type\":\"address\"}],\"name\":\"registerTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"}],\"name\":\"revoke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"}],\"name\":\"top\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"}],\"name\":\"totalVote\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"up\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PreemptivableFuncSigs maps the 4-byte function signature to its string representation.
var PreemptivableFuncSigs = map[string]string{
	"69c07d31": "Ask()",
	"6e6452cb": "Bid()",
	"f318722b": "calcOrderID(address,bytes32)",
	"43271d79": "cancel(bool,bytes32)",
	"ced4aac8": "findAssistingID(bool,address,uint256,uint256,bytes32)",
	"6b3222e6": "getGlobalParams()",
	"2a3ed595": "getLockdown()",
	"07c399a3": "getOrder(bool,bytes32)",
	"c7f758a8": "getProposal(uint256)",
	"c08cc02d": "getProposalCount()",
	"ee1a68c6": "getRemainToAbsorb()",
	"4ea09797": "next(bool,bytes32)",
	"be91d729": "onBlockInitialized(uint256)",
	"0d90b10a": "prev(bool,bytes32)",
	"aa1c259c": "registerTokens(address,address)",
	"74a8f103": "revoke(address)",
	"c0ee0b8a": "tokenFallback(address,uint256,bytes)",
	"8aa3f897": "top(bool)",
	"4def5645": "totalVote(address)",
	"bd041c4d": "vote(address,bool)",
}

// PreemptivableBin is the compiled bytecode used for deploying new contracts.
var PreemptivableBin = "0x608060405262049d406003556002600354816200001857fe5b0460045562093a80600e55670de0b6b3a7640000600f55600060105560006011553480156200004657600080fd5b5060405162003b3c38038062003b3c833981810160405260808110156200006c57600080fd5b508051602082015160408301516060909301519192909183838015620000925760038190555b60008211620000a55760028104620000a7565b815b60045550508015620000b957600e8190555b8115620000c657600f8290555b50505050613a6280620000da6000396000f3fe608060405234801561001057600080fd5b506004361061012c5760003560e01c80638aa3f897116100ad578063c0ee0b8a11610071578063c0ee0b8a14610354578063c7f758a8146103d9578063ced4aac814610433578063ee1a68c614610473578063f318722b146104965761012c565b80638aa3f897146102b4578063aa1c259c146102d3578063bd041c4d14610301578063be91d7291461032f578063c08cc02d1461034c5761012c565b80634ea09797116100f45780634ea097971461021757806369c07d311461023c5780636b3222e6146102585780636e6452cb1461028657806374a8f1031461028e5761012c565b806307c399a3146101315780630d90b10a1461018b5780632a3ed595146101c257806343271d79146101ca5780634def5645146101f1575b600080fd5b6101566004803603604081101561014757600080fd5b508035151590602001356104c2565b604080516001600160a01b03909616865260208601949094528484019290925260608401526080830152519081900360a00190f35b6101b0600480360360408110156101a157600080fd5b50803515159060200135610509565b60408051918252519081900360200190f35b610156610531565b6101ef600480360360408110156101e057600080fd5b50803515159060200135610553565b005b6101b06004803603602081101561020757600080fd5b50356001600160a01b03166105d9565b6101b06004803603604081101561022d57600080fd5b50803515159060200135610601565b610244610625565b604080519115158252519081900360200190f35b61026061062a565b604080519485526020850193909352838301919091526060830152519081900360800190f35b61024461063c565b6101ef600480360360208110156102a457600080fd5b50356001600160a01b0316610641565b6101b0600480360360208110156102ca57600080fd5b503515156107db565b6101ef600480360360408110156102e957600080fd5b506001600160a01b03813581169160200135166107f5565b6101ef6004803603604081101561031757600080fd5b506001600160a01b0381351690602001351515610970565b6101ef6004803603602081101561034557600080fd5b50356109ea565b6101b0610b6f565b6101ef6004803603606081101561036a57600080fd5b6001600160a01b038235169160208101359181019060608101604082013564010000000081111561039a57600080fd5b8201836020820111156103ac57600080fd5b803590602001918460018302840111640100000000831117156103ce57600080fd5b509092509050610b81565b6103f6600480360360208110156103ef57600080fd5b5035610ca4565b604080516001600160a01b0390971687526020870195909552858501939093526060850191909152608084015260a0830152519081900360c00190f35b6101b0600480360360a081101561044957600080fd5b5080351515906001600160a01b036020820135169060408101359060608101359060800135610cf8565b61047b610d60565b60408051921515835260208301919091528051918290030190f35b6101b0600480360360408110156104ac57600080fd5b506001600160a01b038135169060200135610e0a565b90151560009081526020818152604080832093835260029384019091529020805460018201549282015460038301546004909301546001600160a01b039092169490929190565b8115156000908152602081815260408083208484526002019091529020600301545b92915050565b600954600b54600a54600c54600d546001600160a01b03909416939091929394565b8115156000908152602081815260408083208484526002810190925290912080546001600160a01b031633146105c3576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037b93232b91036b0b5b2b960811b604482015290519081900360640190fd5b6105d3828463ffffffff610e1d16565b50505050565b6000806105ed60128463ffffffff610f1716565b90506105f881610f35565b9150505b919050565b90151560009081526020818152604080832093835260029093019052206004015490565b600081565b601154600f54600e5460105490919293565b600181565b600061065460128363ffffffff610f1716565b80549091506001600160a01b038381169116146106b8576040805162461bcd60e51b815260206004820152601e60248201527f6f6e6c79206d616b65722063616e207265766f6b652070726f706f73616c0000604482015290519081900360640190fd5b6015805460018101825560009190915260068201805490916003027f55f448fdea98c4d29eb340757ef0a66cd03dbb9538908a6a81d96026b71ec475019061070390829084906137dc565b5050600154825460028401546040805163a9059cbb60e01b81526001600160a01b03938416600482015260248101929092525191909216925063a9059cbb916044808201926020929091908290030181600087803b15801561076457600080fd5b505af1158015610778573d6000803e3d6000fd5b505050506040513d602081101561078e57600080fd5b506107a2905060128363ffffffff61101b16565b506040516001600160a01b038316907f9f77920c3de8baaa98d273e8aa75fae382aaa9f7f60f38979137853e5b73ea2c90600090a25050565b80151560009081526020819052604081206105f881611097565b6001546001600160a01b031615610853576040805162461bcd60e51b815260206004820152601960248201527f566f6c6174696c65546f6b656e20616c72656164792073657400000000000000604482015290519081900360640190fd5b6002546001600160a01b0316156108b1576040805162461bcd60e51b815260206004820152601960248201527f537461626c697a65546f6b656e20616c72656164792073657400000000000000604482015290519081900360640190fd5b600180546001600160a01b038085166001600160a01b03199283161790925560028054928416929091169190911790556108eb82826110b0565b600254604080516318160ddd60e01b815290516000926001600160a01b0316916318160ddd916004808301926020929190829003018186803b15801561093057600080fd5b505afa158015610944573d6000803e3d6000fd5b505050506040513d602081101561095a57600080fd5b5051905061096b8180600080611104565b505050565b61098160128363ffffffff6111b216565b6109c5576040805162461bcd60e51b815260206004820152601060248201526f1b9bc81cdd58da081c1c9bdc1bdcd85b60821b604482015290519081900360640190fd5b60006109d860128463ffffffff610f1716565b905061096b818363ffffffff6111d316565b3315610a2e576040805162461bcd60e51b815260206004820152600e60248201526d636f6e73656e737573206f6e6c7960901b604482015290519081900360640190fd5b60005b601554811015610a6657610a5e60158281548110610a4b57fe5b90600052602060002090600302016111e7565b600101610a31565b50610a7360156000613828565b610a7d6009611254565b15610a8a57610a8a611271565b610a92611374565b508015610b6357610aa360096113e8565b8015610ab65750600854610100900460ff165b15610b6357600254604080516318160ddd60e01b815290516000926001600160a01b0316916318160ddd916004808301926020929190829003018186803b158015610b0057600080fd5b505afa158015610b14573d6000803e3d6000fd5b505050506040513d6020811015610b2a57600080fd5b505190506000610b3a8383611404565b90508015801590610b4f5750610b4f8161141e565b6008805460ff191691151591909117905550505b610b6c81611548565b50565b6000610b7b6012611676565b90505b90565b608081148015610b9b57506001546001600160a01b031633145b15610c3e57610bb160128563ffffffff6111b216565b15610bfc576040805162461bcd60e51b8152602060048201526016602482015275185b1c9958591e481a185cc818481c1c9bdc1bdcd85b60521b604482015290519081900360640190fd5b60008060008085856080811015610c1257600080fd5b50803594506020810135935060400135915060009050610c35888589868661167a565b505050506105d3565b600080806060841415610c725784846060811015610c5b57600080fd5b508035935060208101359250604001359050610c8e565b84846040811015610c8257600080fd5b50803593506020013591505b610c9b87848885856118fc565b50505050505050565b6000808080808080610cbd60128963ffffffff61197d16565b805460028201546001830154600384015460048501546005909501546001600160a01b039094169d929c50909a509850919650945092505050565b8415156000908152602081905260408120610d11613849565b506040805160a0810182526001600160a01b0388168152602081018790529081018590526000606082018190526080820152610d5482828663ffffffff6119b216565b98975050505050505050565b6007546000908190610d7757506000905080610e06565b6001610e01600560020154600260009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b158015610dd057600080fd5b505afa158015610de4573d6000803e3d6000fd5b505050506040513d6020811015610dfa57600080fd5b5051611404565b915091505b9091565b6000610e168383611a24565b9392505050565b610e25613849565b50600081815260028084016020908152604092839020835160a08101855281546001600160a01b03168152600182015492810183905292810154938301939093526003830154606083015260049092015460808201529015610f0757825481516020808401516040805163a9059cbb60e01b81526001600160a01b039485166004820152602481019290925251929093169263a9059cbb92604480830193928290030181600087803b158015610eda57600080fd5b505af1158015610eee573d6000803e3d6000fd5b505050506040513d6020811015610f0457600080fd5b50505b61096b838363ffffffff611af116565b6001600160a01b031660009081526002919091016020526040902090565b600080805b610f4684600601611676565b81101561101457600080610f63600687018463ffffffff611b5116565b600154604080516370a0823160e01b81526001600160a01b03808616600483015291519496509294506000939116916370a08231916024808301926020929190829003018186803b158015610fb757600080fd5b505afa158015610fcb573d6000803e3d6000fd5b505050506040513d6020811015610fe157600080fd5b50516001600160a01b0384163101905081156110005793840193611006565b80850394505b505050806001019050610f3a565b5092915050565b6001600160a01b038116600090815260018301602052604081205480611078576040805162461bcd60e51b815260206004820152600d60248201526c1ad95e481b9bdd08195e1a5cdd609a1b604482015290519081900360640190fd5b61108d8460001983018563ffffffff611b9016565b5060019392505050565b6000808052600282016020526040902060040154919050565b60008080526020526110d7600080516020613a0d833981519152838363ffffffff611ce216565b600160009081526020526111006000805160206139cb833981519152828463ffffffff611ce216565b5050565b6040805160a0810182526003544301808252602082018690529181018690526000606082018190528315156080909201829052600592909255600685905560078690556008805461ffff19166101009092029190911790556111668585611404565b60408051828152602081018790528515158183015290519192507f0427b353dc7214e3d8c7f5039475a8e729f4d62922937381e304cd03becf66d2919081900360600190a15050505050565b6001600160a01b031660009081526001919091016020526040902054151590565b61096b60068301338363ffffffff611e8f16565b60005b815481101561124857600082600001828154811061120457fe5b60009182526020808320909101546001600160a01b03168252600185810182526040808420849055600287019092529120805460ff191690559190910190506111ea565b50610b6c816000613877565b600061125f82611fad565b801561052b5750506004015443101590565b61127b6009611fad565b61128457611372565b600b541561131557600154600954600b546040805163a9059cbb60e01b81526001600160a01b039384166004820152602481019290925251919092169163a9059cbb9160448083019260209291908290030181600087803b1580156112e857600080fd5b505af11580156112fc573d6000803e3d6000fd5b505050506040513d602081101561131257600080fd5b50505b6009546040516001600160a01b03909116907f0be774851955c26a1d6a32b13b020663a069006b4a3b643ff0b809d31826057290600090a2600980546001600160a01b03191690556000600a819055600b819055600c819055600d555b565b600061138060096113e8565b1561138d57506000610b7e565b600080611398611fbc565b90925090506001600160a01b0382166113b657600092505050610b7e565b60006113c960128463ffffffff610f1716565b90506113d58183612093565b6113de816120e4565b6001935050505090565b60006113f382611fad565b801561052b57505060040154431090565b600081831161141857828203600003610e16565b50900390565b6000611431600960010154600084612297565b61143d575060006105fc565b6000670de0b6b3a764000061145c611454856122c7565b600c546122dd565b8161146357fe5b0490508060096002015410156114785750600b545b600b805482900390556001546040805163117f5a5560e01b81526004810184905290516001600160a01b039092169163117f5a559160248082019260009290919082900301818387803b1580156114ce57600080fd5b505af11580156114e2573d6000803e3d6000fd5b50506009546040805185815290516001600160a01b0390921693507fa69f22d963cb7981f842db8c1aafcc93d915ba2a95dcf26dcc333a9c2a09be26925081900360200190a2600b5461153f57611537612312565b61153f611271565b50600192915050565b331561158c576040805162461bcd60e51b815260206004820152600e60248201526d636f6e73656e737573206f6e6c7960901b604482015290519081900360640190fd5b6115966005612357565b156115a3576115a3612312565b600254604080516318160ddd60e01b815290516000926001600160a01b0316916318160ddd916004808301926020929190829003018186803b1580156115e857600080fd5b505afa1580156115fc573d6000803e3d6000fd5b505050506040513d602081101561161257600080fd5b50519050811561165857611624612370565b1561163b576116368282600080611104565b611658565b611645818361237c565b1561165857611658828260016000611104565b61166960058263ffffffff61242a16565b1561110057611100612470565b5490565b60036011548161168657fe5b04601154038310156116cf576040805162461bcd60e51b815260206004820152600d60248201526c7374616b6520746f6f206c6f7760981b604482015290519081900360640190fd5b83611713576040805162461bcd60e51b815260206004820152600f60248201526e3d32b9379030b139b7b9383a34b7b760891b604482015290519081900360640190fd5b61171b613895565b8215611792576003600f548161172d57fe5b04600f5403831015611786576040805162461bcd60e51b815260206004820152601b60248201527f736c617368696e67207261746520706172616d20746f6f206c6f770000000000604482015290519081900360640190fd5b6060810183905261179b565b600f5460608201525b60006117a6866122c7565b6117b48684606001516122dd565b816117bb57fe5b049050600081116117fd5760405162461bcd60e51b81526004018080602001828103825260228152602001806139eb6022913960400191505060405180910390fd5b821561185e576003600e548161180f57fe5b04600e54038310156118525760405162461bcd60e51b81526004018080602001828103825260218152602001806139aa6021913960400191505060405180910390fd5b60808201839052611867565b600e5460808301525b6001600160a01b038716825260208201869052604082018590524360a083015261189860128363ffffffff61281216565b50606080830151608080850151604080518b8152602081018b9052808201949094529383015291516001600160a01b038a16927f56e25d1b63c01627fcd54936462c97aeb9a18352bf0ed161e8141a33cfd795ca928290030190a250505050505050565b6000611906613849565b61191287878787612966565b91509150600061192133612a5e565b905061193e8261193033612b41565b83919063ffffffff612bfb16565b61194782612cf6565b156119615761195c818363ffffffff612d0f16565b611973565b6119738184848763ffffffff612db116565b5050505050505050565b600080611990848463ffffffff612e8e16565b6001600160a01b03166000908152600285016020526040902091505092915050565b600081815260028401602052604081205b60040154600081815260028601602052604090209092506119ea848263ffffffff612ebb16565b156119c3575b6003015460008181526002860160205260409020909250611a17848263ffffffff612ebb16565b6119f05750909392505050565b60006002838360405160200180836001600160a01b03166001600160a01b031660601b8152601401828152602001925050506040516020818303038152906040526040518082805190602001908083835b60208310611a945780518252601f199092019160209182019101611a75565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015611ad3573d6000803e3d6000fd5b5050506040513d6020811015611ae857600080fd5b50519392505050565b6000818152600292830160205260408082206004808201805460038085018054885286882090940182905583549187529486209094019390935593835280546001600160a01b031916815560018101839055909301819055908190559055565b60008080611b65858563ffffffff612e8e16565b6001600160a01b038116600090815260028701602052604090205490935060ff169150509250929050565b6001600160a01b03811660009081526001808501602090815260408084208490556002808801909252832080546001600160a01b03191681559182018390558101829055600381018290556004810182905560058101829055906006820181611bf98282613877565b505084546000190184149150611cb0905057825483906000198101908110611c1d57fe5b60009182526020909120015483546001600160a01b0390911690849084908110611c4357fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555081600101836001016000856000018581548110611c8a57fe5b60009182526020808320909101546001600160a01b031683528201929092526040019020555b8254839080611cbb57fe5b600082815260209020810160001990810180546001600160a01b0319169055019055505050565b818360000160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550808360010160006101000a8154816001600160a01b0302191690836001600160a01b031602179055506040518060a00160405280306001600160a01b0316815260200160008152602001600081526020016000801b815260200160001960001b8152508360020160008060001b815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550602082015181600101556040820151816002015560608201518160030155608082015181600401559050506040518060a00160405280306001600160a01b0316815260200160008152602001600181526020016000801b815260200160001960001b81525083600201600060001960001b815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060208201518160010155604082015181600201556060820151816003015560808201518160040155905050505050565b6001600160a01b03821660009081526002840160209081526040808320805460ff19168515151790556001860190915281205480611f135750508254600180820185556000858152602080822090930180546001600160a01b0319166001600160a01b03871690811790915586549082528287019093526040902091909155610e16565b8454811180611f515750836001600160a01b0316856000016001830381548110611f3957fe5b6000918252602090912001546001600160a01b031614155b15611fa25750508254600180820185556000858152602080822090930180546001600160a01b0319166001600160a01b03871690811790915586549082528287019093526040902091909155610e16565b506000949350505050565b546001600160a01b0316151590565b6000806000600360105481611fcd57fe5b0460105403905060006004600e5481611fe257fe5b04905060008080805b611ff56012611676565b81101561207057600061200f60128363ffffffff61197d16565b9050600061201c82612ed9565b90508781121561202d575050612068565b858113156120475781549095506001600160a01b03169350845b8315801561205b5750816005015443038711155b1561206557600193505b50505b600101611feb565b5080612088575060009550859450610e069350505050565b509450925050509091565b6120a36011548360020154612f07565b601155600f5460038301546120b89190612f07565b600f55600e5460048301546120cd9190612f07565b600e556010546120dd9082612f07565b6010555050565b6120f0816006016111e7565b60006120ff82600101546122c7565b612111836002015484600301546122dd565b8161211857fe5b6040805160a08101825285546001600160a01b0390811680835260018801546020840181905260028901549484018590529590940460608301819052600488015443016080909301839052600980546001600160a01b031916909517909455600a94909455600b91909155600c829055600d5583549092506121a3916012911663ffffffff61101b16565b50600254604080516318160ddd60e01b815290516000926001600160a01b0316916318160ddd916004808301926020929190829003018186803b1580156121e957600080fd5b505afa1580156121fd573d6000803e3d6000fd5b505050506040513d602081101561221357600080fd5b5051600a54909150600090612229908390612f62565b90506122388183600180611104565b835460028501546003860154600d5460408051938452602084019290925282820152516001600160a01b03909216917f8427e4488966b7bd3193a4617993e5e6b9186f0c4b2c303cc6178f4e33b77d089181900360600190a250505050565b60008284131580156122a95750818313155b806122bf57508284121580156122bf5750818312155b949350505050565b60008082136122d9578160000361052b565b5090565b6000826122ec5750600061052b565b828202828482816122f957fe5b04141561230757905061052b565b506000199392505050565b60006005819055600681905560078190556008805461ffff191690556040517fbedf0f4abfe86d4ffad593d9607fe70e83ea706033d44d24b3b6283cf3fc4f6b9190a1565b600061236282612f7e565b801561052b57505054431190565b6000610b7b6005612357565b60008282141561238e5750600061052b565b60065460075414156123a25750600161052b565b828211156123f6576006546007548484039110156123d8576006546007540360028183816123cc57fe5b0410159250505061052b565b6007546006540360028282816123ea57fe5b0411159250505061052b565b600654600754838503911115612418576007546006540360028183816123cc57fe5b6006546007540360028282816123ea57fe5b600061243583612f7e565b80156124465750600383015460ff16155b8015612456575082600201548214155b8015610e165750610e168360010154838560020154612f84565b600061247a612faf565b9050600080600080841361248f576001612492565b60005b151581526020810191909152604001600090812060025481549193506001600160a01b0391821691161490806124d9836124cb876122c7565b86919063ffffffff61307b16565b6008549193509150610100900460ff166124f7575050505050611372565b801580612502575081155b15612511575050505050611372565b600254604080516318160ddd60e01b815290516000926001600160a01b0316916318160ddd916004808301926020929190829003018186803b15801561255657600080fd5b505afa15801561256a573d6000803e3d6000fd5b505050506040513d602081101561258057600080fd5b5051905061259560058263ffffffff61242a16565b6125a457505050505050611372565b600080856125b35784846125b6565b83855b6009548954604080516370a0823160e01b81526001600160a01b03938416600482018190529151959750939550939116916370a08231916024808301926020929190829003018186803b15801561260c57600080fd5b505afa158015612620573d6000803e3d6000fd5b505050506040513d602081101561263657600080fd5b505183111561264d57505050505050505050611372565b8754604080516337ed875b60e11b81526001600160a01b0384811660048301526024820187905291519190921691636fdb0eb691604480830192600092919082900301818387803b1580156126a157600080fd5b505af11580156126b5573d6000803e3d6000fd5b505089546040805163117f5a5560e01b81526004810188905290516001600160a01b03909216935063117f5a55925060248082019260009290919082900301818387803b15801561270557600080fd5b505af1158015612719573d6000803e3d6000fd5b50505060018901546040805163bdfde91160e01b81526004810186905290516001600160a01b03909216925063bdfde91191602480830192600092919082900301818387803b15801561276b57600080fd5b505af115801561277f573d6000803e3d6000fd5b5050505060018801546040805163a9059cbb60e01b81526001600160a01b038481166004830152602482018690529151919092169163a9059cbb9160448083019260209291908290030181600087803b1580156127db57600080fd5b505af11580156127ef573d6000803e3d6000fd5b505050506040513d602081101561280557600080fd5b5050505050505050505050565b80516001600160a01b03811660009081526001840160205260408120549091908015612885576040805162461bcd60e51b815260206004820152601c60248201527f6d616b657220616c72656164792068617320612070726f706f73616c00000000604482015290519081900360640190fd5b6001600160a01b03828116600090815260028781016020908152604092839020885181546001600160a01b03191695169490941784558781015160018501559187015190830155606086015160038301556080860151600483015560a0860151600583015560c086015180518051889493600685019261290b92849291909101906138e0565b5050875460018082018a5560008a8152602080822090930180546001600160a01b0319166001600160a01b039990991698891790558a54978152908a0182526040808220979097556002909901905250505093209392505050565b6000612970613849565b6000841180156129805750600083115b6129be576040805162461bcd60e51b815260206004820152600a6024820152691e995c9bc81a5b9c1d5d60b21b604482015290519081900360640190fd5b600160801b841080156129d45750600160801b83105b612a18576040805162461bcd60e51b815260206004820152601060248201526f1a5b9c1d5d081bdd995c881b1a5b5a5d60821b604482015290519081900360640190fd5b612a228686611a24565b6040805160a0810182526001600160a01b0398909816885260208801959095529386019290925250506000606084018190526080840152929050565b60008080526020819052600080516020613a0d833981519152546001600160a01b0383811691161415612aa857506000808052602052600080516020613a0d8339815191526105fc565b600160009081526020526000805160206139cb833981519152546001600160a01b0383811691161415612af45750600160009081526020526000805160206139cb8339815191526105fc565b6040805162461bcd60e51b815260206004820152601760248201527f6e6f206f7264657220626f6f6b20666f7220746f6b656e000000000000000000604482015290519081900360640190fd5b600080805260208190527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb6546001600160a01b0383811691161415612b9d57506000808052602052600080516020613a0d8339815191526105fc565b600160009081526020527fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7e546001600160a01b0383811691161415612af45750600160009081526020526000805160206139cb8339815191526105fc565b6000612c0682611097565b90505b60001981146105d357612c1b83612cf6565b15612c25576105d3565b60008181526002830160205260409020612c45848263ffffffff61332216565b612c4f57506105d3565b806001015484604001511015612cb75760008160010154826002015486604001510281612c7857fe5b049050612c968386604001518387613341909392919063ffffffff16565b6040850151612cb09087908790849063ffffffff6134e916565b50506105d3565b60028101546001820154612cd591879187919063ffffffff6134e916565b612ce5838363ffffffff61362316565b612cee83611097565b915050612c09565b600081602001516000148061052b575050604001511590565b602081015115612d9f57815481516020808401516040805163a9059cbb60e01b81526001600160a01b039485166004820152602481019290925251929093169263a9059cbb92604480830193928290030181600087803b158015612d7257600080fd5b505af1158015612d86573d6000803e3d6000fd5b505050506040513d6020811015612d9c57600080fd5b50505b60006020820181905260409091015250565b60008381526002850160205260409020612dca90611fad565b15612e11576040805162461bcd60e51b81526020600482015260126024820152716f7264657220696e6465782065786973747360701b604482015290519081900360640190fd5b6000838152600285810160209081526040808420865181546001600160a01b0319166001600160a01b0390911617815591860151600183015585015191810191909155606084015160038201556080840151600490910155612e748584846119b2565b9050612e8785858363ffffffff6136e616565b5050505050565b6000826000018281548110612e9f57fe5b6000918252602090912001546001600160a01b03169392505050565b60408201516001820154600290920154602090930151910291021190565b600080612ee583610f35565b905060008113612ef95760009150506105fc565b6105f8818460020154613721565b600082821415612f1857508161052b565b6000612f248484613779565b905083831080612f32575083155b15612f3e57905061052b565b83810360038504811115612f5a5750505060038204820161052b565b509392505050565b600080821215612f7957816000038303905061052b565b500190565b54151590565b6000828411158015612f965750818311155b806122bf57508284101580156122bf5750501115919050565b600080612fc6600560020154600560010154611404565b600754600254604080516318160ddd60e01b8152905193945060009361301693926001600160a01b0316916318160ddd916004808301926020929190829003018186803b158015610dd057600080fd5b905061302460008284612297565b61303357600092505050610b7e565b6000600454838161304057fe5b6008549190059150610100900460ff161561305a57600290055b61306660008284612297565b61307457509150610b7e9050565b9250505090565b600080600061308986611097565b90505b600019811480159061309d57508382105b15613318576000818152600287016020526040812090866130c25781600201546130c8565b81600101545b90506000876130db5782600101546130e1565b82600201545b90508487038083116131d757895460018501546040805163117f5a5560e01b81526004810192909252516001600160a01b039092169163117f5a559160248082019260009290919082900301818387803b15801561313e57600080fd5b505af1158015613152573d6000803e3d6000fd5b50505060018b015460028601546040805163bdfde91160e01b81526004810192909252516001600160a01b03909216925063bdfde91191602480830192600092919082900301818387803b1580156131a957600080fd5b505af11580156131bd573d6000803e3d6000fd5b5050505060048401546131d08b87613623565b9450613307565b82818302816131e257fe5b0491508092508682880110156131fd575061331a9350505050565b60008961320a578261320c565b835b905060008a61321b578461321d565b835b8c546040805163117f5a5560e01b81526004810186905290519293506001600160a01b039091169163117f5a559160248082019260009290919082900301818387803b15801561326c57600080fd5b505af1158015613280573d6000803e3d6000fd5b50505060018d01546040805163bdfde91160e01b81526004810185905290516001600160a01b03909216925063bdfde91191602480830192600092919082900301818387803b1580156132d257600080fd5b505af11580156132e6573d6000803e3d6000fd5b506132ff92508e9150899050848463ffffffff61334116565b506000199550505b50949094019392909201915061308c565b505b935093915050565b6040820151600282015460019092015460209093015191029102101590565b6000838152600285016020526040902060018101548311156133a0576040805162461bcd60e51b815260206004820152601360248201527250503a2066696c6c61626c65203e206861766560681b604482015290519081900360640190fd5b80600201548211156133ef576040805162461bcd60e51b815260206004820152601360248201527214140e88199a5b1b18589b19480f881dd85b9d606a1b604482015290519081900360640190fd5b600180820180548590039055600282018054849003905585015481546040805163a9059cbb60e01b81526001600160a01b039283166004820152602481018690529051919092169163a9059cbb9160448083019260209291908290030181600087803b15801561345e57600080fd5b505af1158015613472573d6000803e3d6000fd5b505050506040513d602081101561348857600080fd5b50506040805160a08101825282546001600160a01b031681526001830154602082015260028301549181019190915260038201546060820152600482015460808201526134d490612cf6565b15612e8757612e87858563ffffffff610e1d16565b8260200151821115613538576040805162461bcd60e51b815260206004820152601360248201527250503a2066696c6c61626c65203e206861766560681b604482015290519081900360640190fd5b8260400151811115613587576040805162461bcd60e51b815260206004820152601360248201527214140e88199a5b1b18589b19480f881dd85b9d606a1b604482015290519081900360640190fd5b60208084018051849003905260408085018051849003905260018601548551825163a9059cbb60e01b81526001600160a01b03918216600482015260248101869052925191169263a9059cbb92604480820193918290030181600087803b1580156135f157600080fd5b505af1158015613605573d6000803e3d6000fd5b505050506040513d602081101561361b57600080fd5b505050505050565b61362b613849565b50600081815260028084016020908152604092839020835160a08101855281546001600160a01b03168152600182015492810192909252918201549281018390526003820154606082015260049091015460808201529015610f075760018301548151604080840151815163a9059cbb60e01b81526001600160a01b03938416600482015260248101919091529051919092169163a9059cbb9160448083019260209291908290030181600087803b158015610eda57600080fd5b600081815260029093016020526040808420600490810180548587528387206003808201879055930181905586529185200183905592529055565b6000826137305750600061052b565b8282028284828161373d57fe5b05141561374b57905061052b565b613757846000856137b5565b156137695750600160ff1b905061052b565b506001600160ff1b039392505050565b600082820183811061378f5760011c905061052b565b600184811c9084901c018481106137a957915061052b9050565b50600019949350505050565b600082841280156137c557508183125b806122bf575082841380156122bf57505012919050565b82805482825590600052602060002090810192821561381c5760005260206000209182015b8281111561381c578254825591600101919060010190613801565b506122d9929150613935565b5080546000825560030290600052602060002090810190610b6c9190613959565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915290565b5080546000825590600052602060002090810190610b6c919061397c565b6040518060e0016040528060006001600160a01b0316815260200160008152602001600081526020016000815260200160008152602001600081526020016138db613996565b905290565b82805482825590600052602060002090810192821561381c579160200282015b8281111561381c57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190613900565b610b7e91905b808211156122d95780546001600160a01b031916815560010161393b565b610b7e91905b808211156122d95760006139738282613877565b5060030161395f565b610b7e91905b808211156122d95760008155600101613982565b604051806020016040528060608152509056fe6c6f636b646f776e206475726174696f6e20706172616d20746f6f2073686f7274ada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d736c617368696e6720666163746f722063616c63756c6174656420746f207a65726fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5a2646970667358221220844019c6778224f8bf2d123f734972175fcbdf63b1a4c26dd6e0b88bdd9f6ae764736f6c63430006080033"

// DeployPreemptivable deploys a new Ethereum contract, binding an instance of Preemptivable to it.
func DeployPreemptivable(auth *bind.TransactOpts, backend bind.ContractBackend, absorptionPace *big.Int, absorptionExpiration *big.Int, initialSlashingPace *big.Int, initialLockdownExpiration *big.Int) (common.Address, *types.Transaction, *Preemptivable, error) {
	parsed, err := abi.JSON(strings.NewReader(PreemptivableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PreemptivableBin), backend, absorptionPace, absorptionExpiration, initialSlashingPace, initialLockdownExpiration)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Preemptivable{PreemptivableCaller: PreemptivableCaller{contract: contract}, PreemptivableTransactor: PreemptivableTransactor{contract: contract}, PreemptivableFilterer: PreemptivableFilterer{contract: contract}}, nil
}

// Preemptivable is an auto generated Go binding around an Ethereum contract.
type Preemptivable struct {
	PreemptivableCaller     // Read-only binding to the contract
	PreemptivableTransactor // Write-only binding to the contract
	PreemptivableFilterer   // Log filterer for contract events
}

// PreemptivableCaller is an auto generated read-only Go binding around an Ethereum contract.
type PreemptivableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PreemptivableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PreemptivableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PreemptivableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PreemptivableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PreemptivableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PreemptivableSession struct {
	Contract     *Preemptivable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PreemptivableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PreemptivableCallerSession struct {
	Contract *PreemptivableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PreemptivableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PreemptivableTransactorSession struct {
	Contract     *PreemptivableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PreemptivableRaw is an auto generated low-level Go binding around an Ethereum contract.
type PreemptivableRaw struct {
	Contract *Preemptivable // Generic contract binding to access the raw methods on
}

// PreemptivableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PreemptivableCallerRaw struct {
	Contract *PreemptivableCaller // Generic read-only contract binding to access the raw methods on
}

// PreemptivableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PreemptivableTransactorRaw struct {
	Contract *PreemptivableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPreemptivable creates a new instance of Preemptivable, bound to a specific deployed contract.
func NewPreemptivable(address common.Address, backend bind.ContractBackend) (*Preemptivable, error) {
	contract, err := bindPreemptivable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Preemptivable{PreemptivableCaller: PreemptivableCaller{contract: contract}, PreemptivableTransactor: PreemptivableTransactor{contract: contract}, PreemptivableFilterer: PreemptivableFilterer{contract: contract}}, nil
}

// NewPreemptivableCaller creates a new read-only instance of Preemptivable, bound to a specific deployed contract.
func NewPreemptivableCaller(address common.Address, caller bind.ContractCaller) (*PreemptivableCaller, error) {
	contract, err := bindPreemptivable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PreemptivableCaller{contract: contract}, nil
}

// NewPreemptivableTransactor creates a new write-only instance of Preemptivable, bound to a specific deployed contract.
func NewPreemptivableTransactor(address common.Address, transactor bind.ContractTransactor) (*PreemptivableTransactor, error) {
	contract, err := bindPreemptivable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PreemptivableTransactor{contract: contract}, nil
}

// NewPreemptivableFilterer creates a new log filterer instance of Preemptivable, bound to a specific deployed contract.
func NewPreemptivableFilterer(address common.Address, filterer bind.ContractFilterer) (*PreemptivableFilterer, error) {
	contract, err := bindPreemptivable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PreemptivableFilterer{contract: contract}, nil
}

// bindPreemptivable binds a generic wrapper to an already deployed contract.
func bindPreemptivable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PreemptivableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Preemptivable *PreemptivableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Preemptivable.Contract.PreemptivableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Preemptivable *PreemptivableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Preemptivable.Contract.PreemptivableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Preemptivable *PreemptivableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Preemptivable.Contract.PreemptivableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Preemptivable *PreemptivableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Preemptivable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Preemptivable *PreemptivableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Preemptivable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Preemptivable *PreemptivableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Preemptivable.Contract.contract.Transact(opts, method, params...)
}

// Ask is a free data retrieval call binding the contract method 0x69c07d31.
//
// Solidity: function Ask() view returns(bool)
func (_Preemptivable *PreemptivableCaller) Ask(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Preemptivable.contract.Call(opts, out, "Ask")
	return *ret0, err
}

// Ask is a free data retrieval call binding the contract method 0x69c07d31.
//
// Solidity: function Ask() view returns(bool)
func (_Preemptivable *PreemptivableSession) Ask() (bool, error) {
	return _Preemptivable.Contract.Ask(&_Preemptivable.CallOpts)
}

// Ask is a free data retrieval call binding the contract method 0x69c07d31.
//
// Solidity: function Ask() view returns(bool)
func (_Preemptivable *PreemptivableCallerSession) Ask() (bool, error) {
	return _Preemptivable.Contract.Ask(&_Preemptivable.CallOpts)
}

// Bid is a free data retrieval call binding the contract method 0x6e6452cb.
//
// Solidity: function Bid() view returns(bool)
func (_Preemptivable *PreemptivableCaller) Bid(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Preemptivable.contract.Call(opts, out, "Bid")
	return *ret0, err
}

// Bid is a free data retrieval call binding the contract method 0x6e6452cb.
//
// Solidity: function Bid() view returns(bool)
func (_Preemptivable *PreemptivableSession) Bid() (bool, error) {
	return _Preemptivable.Contract.Bid(&_Preemptivable.CallOpts)
}

// Bid is a free data retrieval call binding the contract method 0x6e6452cb.
//
// Solidity: function Bid() view returns(bool)
func (_Preemptivable *PreemptivableCallerSession) Bid() (bool, error) {
	return _Preemptivable.Contract.Bid(&_Preemptivable.CallOpts)
}

// CalcOrderID is a free data retrieval call binding the contract method 0xf318722b.
//
// Solidity: function calcOrderID(address maker, bytes32 index) pure returns(bytes32)
func (_Preemptivable *PreemptivableCaller) CalcOrderID(opts *bind.CallOpts, maker common.Address, index [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Preemptivable.contract.Call(opts, out, "calcOrderID", maker, index)
	return *ret0, err
}

// CalcOrderID is a free data retrieval call binding the contract method 0xf318722b.
//
// Solidity: function calcOrderID(address maker, bytes32 index) pure returns(bytes32)
func (_Preemptivable *PreemptivableSession) CalcOrderID(maker common.Address, index [32]byte) ([32]byte, error) {
	return _Preemptivable.Contract.CalcOrderID(&_Preemptivable.CallOpts, maker, index)
}

// CalcOrderID is a free data retrieval call binding the contract method 0xf318722b.
//
// Solidity: function calcOrderID(address maker, bytes32 index) pure returns(bytes32)
func (_Preemptivable *PreemptivableCallerSession) CalcOrderID(maker common.Address, index [32]byte) ([32]byte, error) {
	return _Preemptivable.Contract.CalcOrderID(&_Preemptivable.CallOpts, maker, index)
}

// FindAssistingID is a free data retrieval call binding the contract method 0xced4aac8.
//
// Solidity: function findAssistingID(bool orderType, address maker, uint256 haveAmount, uint256 wantAmount, bytes32 assistingID) view returns(bytes32)
func (_Preemptivable *PreemptivableCaller) FindAssistingID(opts *bind.CallOpts, orderType bool, maker common.Address, haveAmount *big.Int, wantAmount *big.Int, assistingID [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Preemptivable.contract.Call(opts, out, "findAssistingID", orderType, maker, haveAmount, wantAmount, assistingID)
	return *ret0, err
}

// FindAssistingID is a free data retrieval call binding the contract method 0xced4aac8.
//
// Solidity: function findAssistingID(bool orderType, address maker, uint256 haveAmount, uint256 wantAmount, bytes32 assistingID) view returns(bytes32)
func (_Preemptivable *PreemptivableSession) FindAssistingID(orderType bool, maker common.Address, haveAmount *big.Int, wantAmount *big.Int, assistingID [32]byte) ([32]byte, error) {
	return _Preemptivable.Contract.FindAssistingID(&_Preemptivable.CallOpts, orderType, maker, haveAmount, wantAmount, assistingID)
}

// FindAssistingID is a free data retrieval call binding the contract method 0xced4aac8.
//
// Solidity: function findAssistingID(bool orderType, address maker, uint256 haveAmount, uint256 wantAmount, bytes32 assistingID) view returns(bytes32)
func (_Preemptivable *PreemptivableCallerSession) FindAssistingID(orderType bool, maker common.Address, haveAmount *big.Int, wantAmount *big.Int, assistingID [32]byte) ([32]byte, error) {
	return _Preemptivable.Contract.FindAssistingID(&_Preemptivable.CallOpts, orderType, maker, haveAmount, wantAmount, assistingID)
}

// GetGlobalParams is a free data retrieval call binding the contract method 0x6b3222e6.
//
// Solidity: function getGlobalParams() view returns(uint256 stake, uint256 slashingRate, uint256 lockdownExpiration, uint256 rank)
func (_Preemptivable *PreemptivableCaller) GetGlobalParams(opts *bind.CallOpts) (struct {
	Stake              *big.Int
	SlashingRate       *big.Int
	LockdownExpiration *big.Int
	Rank               *big.Int
}, error) {
	ret := new(struct {
		Stake              *big.Int
		SlashingRate       *big.Int
		LockdownExpiration *big.Int
		Rank               *big.Int
	})
	out := ret
	err := _Preemptivable.contract.Call(opts, out, "getGlobalParams")
	return *ret, err
}

// GetGlobalParams is a free data retrieval call binding the contract method 0x6b3222e6.
//
// Solidity: function getGlobalParams() view returns(uint256 stake, uint256 slashingRate, uint256 lockdownExpiration, uint256 rank)
func (_Preemptivable *PreemptivableSession) GetGlobalParams() (struct {
	Stake              *big.Int
	SlashingRate       *big.Int
	LockdownExpiration *big.Int
	Rank               *big.Int
}, error) {
	return _Preemptivable.Contract.GetGlobalParams(&_Preemptivable.CallOpts)
}

// GetGlobalParams is a free data retrieval call binding the contract method 0x6b3222e6.
//
// Solidity: function getGlobalParams() view returns(uint256 stake, uint256 slashingRate, uint256 lockdownExpiration, uint256 rank)
func (_Preemptivable *PreemptivableCallerSession) GetGlobalParams() (struct {
	Stake              *big.Int
	SlashingRate       *big.Int
	LockdownExpiration *big.Int
	Rank               *big.Int
}, error) {
	return _Preemptivable.Contract.GetGlobalParams(&_Preemptivable.CallOpts)
}

// GetLockdown is a free data retrieval call binding the contract method 0x2a3ed595.
//
// Solidity: function getLockdown() view returns(address maker, uint256 stake, int256 amount, uint256 slashingFactor, uint256 unlockNumber)
func (_Preemptivable *PreemptivableCaller) GetLockdown(opts *bind.CallOpts) (struct {
	Maker          common.Address
	Stake          *big.Int
	Amount         *big.Int
	SlashingFactor *big.Int
	UnlockNumber   *big.Int
}, error) {
	ret := new(struct {
		Maker          common.Address
		Stake          *big.Int
		Amount         *big.Int
		SlashingFactor *big.Int
		UnlockNumber   *big.Int
	})
	out := ret
	err := _Preemptivable.contract.Call(opts, out, "getLockdown")
	return *ret, err
}

// GetLockdown is a free data retrieval call binding the contract method 0x2a3ed595.
//
// Solidity: function getLockdown() view returns(address maker, uint256 stake, int256 amount, uint256 slashingFactor, uint256 unlockNumber)
func (_Preemptivable *PreemptivableSession) GetLockdown() (struct {
	Maker          common.Address
	Stake          *big.Int
	Amount         *big.Int
	SlashingFactor *big.Int
	UnlockNumber   *big.Int
}, error) {
	return _Preemptivable.Contract.GetLockdown(&_Preemptivable.CallOpts)
}

// GetLockdown is a free data retrieval call binding the contract method 0x2a3ed595.
//
// Solidity: function getLockdown() view returns(address maker, uint256 stake, int256 amount, uint256 slashingFactor, uint256 unlockNumber)
func (_Preemptivable *PreemptivableCallerSession) GetLockdown() (struct {
	Maker          common.Address
	Stake          *big.Int
	Amount         *big.Int
	SlashingFactor *big.Int
	UnlockNumber   *big.Int
}, error) {
	return _Preemptivable.Contract.GetLockdown(&_Preemptivable.CallOpts)
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) view returns(address maker, uint256 have, uint256 want, bytes32 prevID, bytes32 nextID)
func (_Preemptivable *PreemptivableCaller) GetOrder(opts *bind.CallOpts, _orderType bool, _id [32]byte) (struct {
	Maker  common.Address
	Have   *big.Int
	Want   *big.Int
	PrevID [32]byte
	NextID [32]byte
}, error) {
	ret := new(struct {
		Maker  common.Address
		Have   *big.Int
		Want   *big.Int
		PrevID [32]byte
		NextID [32]byte
	})
	out := ret
	err := _Preemptivable.contract.Call(opts, out, "getOrder", _orderType, _id)
	return *ret, err
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) view returns(address maker, uint256 have, uint256 want, bytes32 prevID, bytes32 nextID)
func (_Preemptivable *PreemptivableSession) GetOrder(_orderType bool, _id [32]byte) (struct {
	Maker  common.Address
	Have   *big.Int
	Want   *big.Int
	PrevID [32]byte
	NextID [32]byte
}, error) {
	return _Preemptivable.Contract.GetOrder(&_Preemptivable.CallOpts, _orderType, _id)
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) view returns(address maker, uint256 have, uint256 want, bytes32 prevID, bytes32 nextID)
func (_Preemptivable *PreemptivableCallerSession) GetOrder(_orderType bool, _id [32]byte) (struct {
	Maker  common.Address
	Have   *big.Int
	Want   *big.Int
	PrevID [32]byte
	NextID [32]byte
}, error) {
	return _Preemptivable.Contract.GetOrder(&_Preemptivable.CallOpts, _orderType, _id)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 idx) view returns(address maker, uint256 stake, int256 amount, uint256 slashingRate, uint256 lockdownExpiration, uint256 number)
func (_Preemptivable *PreemptivableCaller) GetProposal(opts *bind.CallOpts, idx *big.Int) (struct {
	Maker              common.Address
	Stake              *big.Int
	Amount             *big.Int
	SlashingRate       *big.Int
	LockdownExpiration *big.Int
	Number             *big.Int
}, error) {
	ret := new(struct {
		Maker              common.Address
		Stake              *big.Int
		Amount             *big.Int
		SlashingRate       *big.Int
		LockdownExpiration *big.Int
		Number             *big.Int
	})
	out := ret
	err := _Preemptivable.contract.Call(opts, out, "getProposal", idx)
	return *ret, err
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 idx) view returns(address maker, uint256 stake, int256 amount, uint256 slashingRate, uint256 lockdownExpiration, uint256 number)
func (_Preemptivable *PreemptivableSession) GetProposal(idx *big.Int) (struct {
	Maker              common.Address
	Stake              *big.Int
	Amount             *big.Int
	SlashingRate       *big.Int
	LockdownExpiration *big.Int
	Number             *big.Int
}, error) {
	return _Preemptivable.Contract.GetProposal(&_Preemptivable.CallOpts, idx)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 idx) view returns(address maker, uint256 stake, int256 amount, uint256 slashingRate, uint256 lockdownExpiration, uint256 number)
func (_Preemptivable *PreemptivableCallerSession) GetProposal(idx *big.Int) (struct {
	Maker              common.Address
	Stake              *big.Int
	Amount             *big.Int
	SlashingRate       *big.Int
	LockdownExpiration *big.Int
	Number             *big.Int
}, error) {
	return _Preemptivable.Contract.GetProposal(&_Preemptivable.CallOpts, idx)
}

// GetProposalCount is a free data retrieval call binding the contract method 0xc08cc02d.
//
// Solidity: function getProposalCount() view returns(uint256)
func (_Preemptivable *PreemptivableCaller) GetProposalCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Preemptivable.contract.Call(opts, out, "getProposalCount")
	return *ret0, err
}

// GetProposalCount is a free data retrieval call binding the contract method 0xc08cc02d.
//
// Solidity: function getProposalCount() view returns(uint256)
func (_Preemptivable *PreemptivableSession) GetProposalCount() (*big.Int, error) {
	return _Preemptivable.Contract.GetProposalCount(&_Preemptivable.CallOpts)
}

// GetProposalCount is a free data retrieval call binding the contract method 0xc08cc02d.
//
// Solidity: function getProposalCount() view returns(uint256)
func (_Preemptivable *PreemptivableCallerSession) GetProposalCount() (*big.Int, error) {
	return _Preemptivable.Contract.GetProposalCount(&_Preemptivable.CallOpts)
}

// GetRemainToAbsorb is a free data retrieval call binding the contract method 0xee1a68c6.
//
// Solidity: function getRemainToAbsorb() view returns(bool, int256)
func (_Preemptivable *PreemptivableCaller) GetRemainToAbsorb(opts *bind.CallOpts) (bool, *big.Int, error) {
	var (
		ret0 = new(bool)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Preemptivable.contract.Call(opts, out, "getRemainToAbsorb")
	return *ret0, *ret1, err
}

// GetRemainToAbsorb is a free data retrieval call binding the contract method 0xee1a68c6.
//
// Solidity: function getRemainToAbsorb() view returns(bool, int256)
func (_Preemptivable *PreemptivableSession) GetRemainToAbsorb() (bool, *big.Int, error) {
	return _Preemptivable.Contract.GetRemainToAbsorb(&_Preemptivable.CallOpts)
}

// GetRemainToAbsorb is a free data retrieval call binding the contract method 0xee1a68c6.
//
// Solidity: function getRemainToAbsorb() view returns(bool, int256)
func (_Preemptivable *PreemptivableCallerSession) GetRemainToAbsorb() (bool, *big.Int, error) {
	return _Preemptivable.Contract.GetRemainToAbsorb(&_Preemptivable.CallOpts)
}

// Next is a free data retrieval call binding the contract method 0x4ea09797.
//
// Solidity: function next(bool orderType, bytes32 id) view returns(bytes32)
func (_Preemptivable *PreemptivableCaller) Next(opts *bind.CallOpts, orderType bool, id [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Preemptivable.contract.Call(opts, out, "next", orderType, id)
	return *ret0, err
}

// Next is a free data retrieval call binding the contract method 0x4ea09797.
//
// Solidity: function next(bool orderType, bytes32 id) view returns(bytes32)
func (_Preemptivable *PreemptivableSession) Next(orderType bool, id [32]byte) ([32]byte, error) {
	return _Preemptivable.Contract.Next(&_Preemptivable.CallOpts, orderType, id)
}

// Next is a free data retrieval call binding the contract method 0x4ea09797.
//
// Solidity: function next(bool orderType, bytes32 id) view returns(bytes32)
func (_Preemptivable *PreemptivableCallerSession) Next(orderType bool, id [32]byte) ([32]byte, error) {
	return _Preemptivable.Contract.Next(&_Preemptivable.CallOpts, orderType, id)
}

// Prev is a free data retrieval call binding the contract method 0x0d90b10a.
//
// Solidity: function prev(bool orderType, bytes32 id) view returns(bytes32)
func (_Preemptivable *PreemptivableCaller) Prev(opts *bind.CallOpts, orderType bool, id [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Preemptivable.contract.Call(opts, out, "prev", orderType, id)
	return *ret0, err
}

// Prev is a free data retrieval call binding the contract method 0x0d90b10a.
//
// Solidity: function prev(bool orderType, bytes32 id) view returns(bytes32)
func (_Preemptivable *PreemptivableSession) Prev(orderType bool, id [32]byte) ([32]byte, error) {
	return _Preemptivable.Contract.Prev(&_Preemptivable.CallOpts, orderType, id)
}

// Prev is a free data retrieval call binding the contract method 0x0d90b10a.
//
// Solidity: function prev(bool orderType, bytes32 id) view returns(bytes32)
func (_Preemptivable *PreemptivableCallerSession) Prev(orderType bool, id [32]byte) ([32]byte, error) {
	return _Preemptivable.Contract.Prev(&_Preemptivable.CallOpts, orderType, id)
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool orderType) view returns(bytes32)
func (_Preemptivable *PreemptivableCaller) Top(opts *bind.CallOpts, orderType bool) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Preemptivable.contract.Call(opts, out, "top", orderType)
	return *ret0, err
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool orderType) view returns(bytes32)
func (_Preemptivable *PreemptivableSession) Top(orderType bool) ([32]byte, error) {
	return _Preemptivable.Contract.Top(&_Preemptivable.CallOpts, orderType)
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool orderType) view returns(bytes32)
func (_Preemptivable *PreemptivableCallerSession) Top(orderType bool) ([32]byte, error) {
	return _Preemptivable.Contract.Top(&_Preemptivable.CallOpts, orderType)
}

// TotalVote is a free data retrieval call binding the contract method 0x4def5645.
//
// Solidity: function totalVote(address maker) view returns(int256)
func (_Preemptivable *PreemptivableCaller) TotalVote(opts *bind.CallOpts, maker common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Preemptivable.contract.Call(opts, out, "totalVote", maker)
	return *ret0, err
}

// TotalVote is a free data retrieval call binding the contract method 0x4def5645.
//
// Solidity: function totalVote(address maker) view returns(int256)
func (_Preemptivable *PreemptivableSession) TotalVote(maker common.Address) (*big.Int, error) {
	return _Preemptivable.Contract.TotalVote(&_Preemptivable.CallOpts, maker)
}

// TotalVote is a free data retrieval call binding the contract method 0x4def5645.
//
// Solidity: function totalVote(address maker) view returns(int256)
func (_Preemptivable *PreemptivableCallerSession) TotalVote(maker common.Address) (*big.Int, error) {
	return _Preemptivable.Contract.TotalVote(&_Preemptivable.CallOpts, maker)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool orderType, bytes32 id) returns()
func (_Preemptivable *PreemptivableTransactor) Cancel(opts *bind.TransactOpts, orderType bool, id [32]byte) (*types.Transaction, error) {
	return _Preemptivable.contract.Transact(opts, "cancel", orderType, id)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool orderType, bytes32 id) returns()
func (_Preemptivable *PreemptivableSession) Cancel(orderType bool, id [32]byte) (*types.Transaction, error) {
	return _Preemptivable.Contract.Cancel(&_Preemptivable.TransactOpts, orderType, id)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool orderType, bytes32 id) returns()
func (_Preemptivable *PreemptivableTransactorSession) Cancel(orderType bool, id [32]byte) (*types.Transaction, error) {
	return _Preemptivable.Contract.Cancel(&_Preemptivable.TransactOpts, orderType, id)
}

// OnBlockInitialized is a paid mutator transaction binding the contract method 0xbe91d729.
//
// Solidity: function onBlockInitialized(uint256 target) returns()
func (_Preemptivable *PreemptivableTransactor) OnBlockInitialized(opts *bind.TransactOpts, target *big.Int) (*types.Transaction, error) {
	return _Preemptivable.contract.Transact(opts, "onBlockInitialized", target)
}

// OnBlockInitialized is a paid mutator transaction binding the contract method 0xbe91d729.
//
// Solidity: function onBlockInitialized(uint256 target) returns()
func (_Preemptivable *PreemptivableSession) OnBlockInitialized(target *big.Int) (*types.Transaction, error) {
	return _Preemptivable.Contract.OnBlockInitialized(&_Preemptivable.TransactOpts, target)
}

// OnBlockInitialized is a paid mutator transaction binding the contract method 0xbe91d729.
//
// Solidity: function onBlockInitialized(uint256 target) returns()
func (_Preemptivable *PreemptivableTransactorSession) OnBlockInitialized(target *big.Int) (*types.Transaction, error) {
	return _Preemptivable.Contract.OnBlockInitialized(&_Preemptivable.TransactOpts, target)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xaa1c259c.
//
// Solidity: function registerTokens(address volatileToken, address stablizeToken) returns()
func (_Preemptivable *PreemptivableTransactor) RegisterTokens(opts *bind.TransactOpts, volatileToken common.Address, stablizeToken common.Address) (*types.Transaction, error) {
	return _Preemptivable.contract.Transact(opts, "registerTokens", volatileToken, stablizeToken)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xaa1c259c.
//
// Solidity: function registerTokens(address volatileToken, address stablizeToken) returns()
func (_Preemptivable *PreemptivableSession) RegisterTokens(volatileToken common.Address, stablizeToken common.Address) (*types.Transaction, error) {
	return _Preemptivable.Contract.RegisterTokens(&_Preemptivable.TransactOpts, volatileToken, stablizeToken)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xaa1c259c.
//
// Solidity: function registerTokens(address volatileToken, address stablizeToken) returns()
func (_Preemptivable *PreemptivableTransactorSession) RegisterTokens(volatileToken common.Address, stablizeToken common.Address) (*types.Transaction, error) {
	return _Preemptivable.Contract.RegisterTokens(&_Preemptivable.TransactOpts, volatileToken, stablizeToken)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address maker) returns()
func (_Preemptivable *PreemptivableTransactor) Revoke(opts *bind.TransactOpts, maker common.Address) (*types.Transaction, error) {
	return _Preemptivable.contract.Transact(opts, "revoke", maker)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address maker) returns()
func (_Preemptivable *PreemptivableSession) Revoke(maker common.Address) (*types.Transaction, error) {
	return _Preemptivable.Contract.Revoke(&_Preemptivable.TransactOpts, maker)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address maker) returns()
func (_Preemptivable *PreemptivableTransactorSession) Revoke(maker common.Address) (*types.Transaction, error) {
	return _Preemptivable.Contract.Revoke(&_Preemptivable.TransactOpts, maker)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address maker, uint256 value, bytes data) returns()
func (_Preemptivable *PreemptivableTransactor) TokenFallback(opts *bind.TransactOpts, maker common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Preemptivable.contract.Transact(opts, "tokenFallback", maker, value, data)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address maker, uint256 value, bytes data) returns()
func (_Preemptivable *PreemptivableSession) TokenFallback(maker common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Preemptivable.Contract.TokenFallback(&_Preemptivable.TransactOpts, maker, value, data)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address maker, uint256 value, bytes data) returns()
func (_Preemptivable *PreemptivableTransactorSession) TokenFallback(maker common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Preemptivable.Contract.TokenFallback(&_Preemptivable.TransactOpts, maker, value, data)
}

// Vote is a paid mutator transaction binding the contract method 0xbd041c4d.
//
// Solidity: function vote(address maker, bool up) returns()
func (_Preemptivable *PreemptivableTransactor) Vote(opts *bind.TransactOpts, maker common.Address, up bool) (*types.Transaction, error) {
	return _Preemptivable.contract.Transact(opts, "vote", maker, up)
}

// Vote is a paid mutator transaction binding the contract method 0xbd041c4d.
//
// Solidity: function vote(address maker, bool up) returns()
func (_Preemptivable *PreemptivableSession) Vote(maker common.Address, up bool) (*types.Transaction, error) {
	return _Preemptivable.Contract.Vote(&_Preemptivable.TransactOpts, maker, up)
}

// Vote is a paid mutator transaction binding the contract method 0xbd041c4d.
//
// Solidity: function vote(address maker, bool up) returns()
func (_Preemptivable *PreemptivableTransactorSession) Vote(maker common.Address, up bool) (*types.Transaction, error) {
	return _Preemptivable.Contract.Vote(&_Preemptivable.TransactOpts, maker, up)
}

// PreemptivableAbsorptionIterator is returned from FilterAbsorption and is used to iterate over the raw logs and unpacked data for Absorption events raised by the Preemptivable contract.
type PreemptivableAbsorptionIterator struct {
	Event *PreemptivableAbsorption // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PreemptivableAbsorptionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PreemptivableAbsorption)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PreemptivableAbsorption)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PreemptivableAbsorptionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PreemptivableAbsorptionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PreemptivableAbsorption represents a Absorption event raised by the Preemptivable contract.
type PreemptivableAbsorption struct {
	Amount  *big.Int
	Supply  *big.Int
	Emptive bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAbsorption is a free log retrieval operation binding the contract event 0x0427b353dc7214e3d8c7f5039475a8e729f4d62922937381e304cd03becf66d2.
//
// Solidity: event Absorption(int256 amount, uint256 supply, bool emptive)
func (_Preemptivable *PreemptivableFilterer) FilterAbsorption(opts *bind.FilterOpts) (*PreemptivableAbsorptionIterator, error) {

	logs, sub, err := _Preemptivable.contract.FilterLogs(opts, "Absorption")
	if err != nil {
		return nil, err
	}
	return &PreemptivableAbsorptionIterator{contract: _Preemptivable.contract, event: "Absorption", logs: logs, sub: sub}, nil
}

// WatchAbsorption is a free log subscription operation binding the contract event 0x0427b353dc7214e3d8c7f5039475a8e729f4d62922937381e304cd03becf66d2.
//
// Solidity: event Absorption(int256 amount, uint256 supply, bool emptive)
func (_Preemptivable *PreemptivableFilterer) WatchAbsorption(opts *bind.WatchOpts, sink chan<- *PreemptivableAbsorption) (event.Subscription, error) {

	logs, sub, err := _Preemptivable.contract.WatchLogs(opts, "Absorption")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PreemptivableAbsorption)
				if err := _Preemptivable.contract.UnpackLog(event, "Absorption", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAbsorption is a log parse operation binding the contract event 0x0427b353dc7214e3d8c7f5039475a8e729f4d62922937381e304cd03becf66d2.
//
// Solidity: event Absorption(int256 amount, uint256 supply, bool emptive)
func (_Preemptivable *PreemptivableFilterer) ParseAbsorption(log types.Log) (*PreemptivableAbsorption, error) {
	event := new(PreemptivableAbsorption)
	if err := _Preemptivable.contract.UnpackLog(event, "Absorption", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PreemptivablePreemptiveIterator is returned from FilterPreemptive and is used to iterate over the raw logs and unpacked data for Preemptive events raised by the Preemptivable contract.
type PreemptivablePreemptiveIterator struct {
	Event *PreemptivablePreemptive // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PreemptivablePreemptiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PreemptivablePreemptive)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PreemptivablePreemptive)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PreemptivablePreemptiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PreemptivablePreemptiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PreemptivablePreemptive represents a Preemptive event raised by the Preemptivable contract.
type PreemptivablePreemptive struct {
	Maker              common.Address
	Stake              *big.Int
	LockdownExpiration *big.Int
	UnlockNumber       *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPreemptive is a free log retrieval operation binding the contract event 0x8427e4488966b7bd3193a4617993e5e6b9186f0c4b2c303cc6178f4e33b77d08.
//
// Solidity: event Preemptive(address indexed maker, uint256 stake, uint256 lockdownExpiration, uint256 unlockNumber)
func (_Preemptivable *PreemptivableFilterer) FilterPreemptive(opts *bind.FilterOpts, maker []common.Address) (*PreemptivablePreemptiveIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Preemptivable.contract.FilterLogs(opts, "Preemptive", makerRule)
	if err != nil {
		return nil, err
	}
	return &PreemptivablePreemptiveIterator{contract: _Preemptivable.contract, event: "Preemptive", logs: logs, sub: sub}, nil
}

// WatchPreemptive is a free log subscription operation binding the contract event 0x8427e4488966b7bd3193a4617993e5e6b9186f0c4b2c303cc6178f4e33b77d08.
//
// Solidity: event Preemptive(address indexed maker, uint256 stake, uint256 lockdownExpiration, uint256 unlockNumber)
func (_Preemptivable *PreemptivableFilterer) WatchPreemptive(opts *bind.WatchOpts, sink chan<- *PreemptivablePreemptive, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Preemptivable.contract.WatchLogs(opts, "Preemptive", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PreemptivablePreemptive)
				if err := _Preemptivable.contract.UnpackLog(event, "Preemptive", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePreemptive is a log parse operation binding the contract event 0x8427e4488966b7bd3193a4617993e5e6b9186f0c4b2c303cc6178f4e33b77d08.
//
// Solidity: event Preemptive(address indexed maker, uint256 stake, uint256 lockdownExpiration, uint256 unlockNumber)
func (_Preemptivable *PreemptivableFilterer) ParsePreemptive(log types.Log) (*PreemptivablePreemptive, error) {
	event := new(PreemptivablePreemptive)
	if err := _Preemptivable.contract.UnpackLog(event, "Preemptive", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PreemptivableProposeIterator is returned from FilterPropose and is used to iterate over the raw logs and unpacked data for Propose events raised by the Preemptivable contract.
type PreemptivableProposeIterator struct {
	Event *PreemptivablePropose // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PreemptivableProposeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PreemptivablePropose)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PreemptivablePropose)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PreemptivableProposeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PreemptivableProposeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PreemptivablePropose represents a Propose event raised by the Preemptivable contract.
type PreemptivablePropose struct {
	Maker              common.Address
	Amount             *big.Int
	Stake              *big.Int
	LockdownExpiration *big.Int
	SlashingRate       *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPropose is a free log retrieval operation binding the contract event 0x56e25d1b63c01627fcd54936462c97aeb9a18352bf0ed161e8141a33cfd795ca.
//
// Solidity: event Propose(address indexed maker, int256 amount, uint256 stake, uint256 lockdownExpiration, uint256 slashingRate)
func (_Preemptivable *PreemptivableFilterer) FilterPropose(opts *bind.FilterOpts, maker []common.Address) (*PreemptivableProposeIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Preemptivable.contract.FilterLogs(opts, "Propose", makerRule)
	if err != nil {
		return nil, err
	}
	return &PreemptivableProposeIterator{contract: _Preemptivable.contract, event: "Propose", logs: logs, sub: sub}, nil
}

// WatchPropose is a free log subscription operation binding the contract event 0x56e25d1b63c01627fcd54936462c97aeb9a18352bf0ed161e8141a33cfd795ca.
//
// Solidity: event Propose(address indexed maker, int256 amount, uint256 stake, uint256 lockdownExpiration, uint256 slashingRate)
func (_Preemptivable *PreemptivableFilterer) WatchPropose(opts *bind.WatchOpts, sink chan<- *PreemptivablePropose, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Preemptivable.contract.WatchLogs(opts, "Propose", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PreemptivablePropose)
				if err := _Preemptivable.contract.UnpackLog(event, "Propose", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePropose is a log parse operation binding the contract event 0x56e25d1b63c01627fcd54936462c97aeb9a18352bf0ed161e8141a33cfd795ca.
//
// Solidity: event Propose(address indexed maker, int256 amount, uint256 stake, uint256 lockdownExpiration, uint256 slashingRate)
func (_Preemptivable *PreemptivableFilterer) ParsePropose(log types.Log) (*PreemptivablePropose, error) {
	event := new(PreemptivablePropose)
	if err := _Preemptivable.contract.UnpackLog(event, "Propose", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PreemptivableRevokeIterator is returned from FilterRevoke and is used to iterate over the raw logs and unpacked data for Revoke events raised by the Preemptivable contract.
type PreemptivableRevokeIterator struct {
	Event *PreemptivableRevoke // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PreemptivableRevokeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PreemptivableRevoke)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PreemptivableRevoke)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PreemptivableRevokeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PreemptivableRevokeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PreemptivableRevoke represents a Revoke event raised by the Preemptivable contract.
type PreemptivableRevoke struct {
	Maker common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRevoke is a free log retrieval operation binding the contract event 0x9f77920c3de8baaa98d273e8aa75fae382aaa9f7f60f38979137853e5b73ea2c.
//
// Solidity: event Revoke(address indexed maker)
func (_Preemptivable *PreemptivableFilterer) FilterRevoke(opts *bind.FilterOpts, maker []common.Address) (*PreemptivableRevokeIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Preemptivable.contract.FilterLogs(opts, "Revoke", makerRule)
	if err != nil {
		return nil, err
	}
	return &PreemptivableRevokeIterator{contract: _Preemptivable.contract, event: "Revoke", logs: logs, sub: sub}, nil
}

// WatchRevoke is a free log subscription operation binding the contract event 0x9f77920c3de8baaa98d273e8aa75fae382aaa9f7f60f38979137853e5b73ea2c.
//
// Solidity: event Revoke(address indexed maker)
func (_Preemptivable *PreemptivableFilterer) WatchRevoke(opts *bind.WatchOpts, sink chan<- *PreemptivableRevoke, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Preemptivable.contract.WatchLogs(opts, "Revoke", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PreemptivableRevoke)
				if err := _Preemptivable.contract.UnpackLog(event, "Revoke", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRevoke is a log parse operation binding the contract event 0x9f77920c3de8baaa98d273e8aa75fae382aaa9f7f60f38979137853e5b73ea2c.
//
// Solidity: event Revoke(address indexed maker)
func (_Preemptivable *PreemptivableFilterer) ParseRevoke(log types.Log) (*PreemptivableRevoke, error) {
	event := new(PreemptivableRevoke)
	if err := _Preemptivable.contract.UnpackLog(event, "Revoke", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PreemptivableSlashIterator is returned from FilterSlash and is used to iterate over the raw logs and unpacked data for Slash events raised by the Preemptivable contract.
type PreemptivableSlashIterator struct {
	Event *PreemptivableSlash // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PreemptivableSlashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PreemptivableSlash)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PreemptivableSlash)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PreemptivableSlashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PreemptivableSlashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PreemptivableSlash represents a Slash event raised by the Preemptivable contract.
type PreemptivableSlash struct {
	Maker  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSlash is a free log retrieval operation binding the contract event 0xa69f22d963cb7981f842db8c1aafcc93d915ba2a95dcf26dcc333a9c2a09be26.
//
// Solidity: event Slash(address indexed maker, uint256 amount)
func (_Preemptivable *PreemptivableFilterer) FilterSlash(opts *bind.FilterOpts, maker []common.Address) (*PreemptivableSlashIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Preemptivable.contract.FilterLogs(opts, "Slash", makerRule)
	if err != nil {
		return nil, err
	}
	return &PreemptivableSlashIterator{contract: _Preemptivable.contract, event: "Slash", logs: logs, sub: sub}, nil
}

// WatchSlash is a free log subscription operation binding the contract event 0xa69f22d963cb7981f842db8c1aafcc93d915ba2a95dcf26dcc333a9c2a09be26.
//
// Solidity: event Slash(address indexed maker, uint256 amount)
func (_Preemptivable *PreemptivableFilterer) WatchSlash(opts *bind.WatchOpts, sink chan<- *PreemptivableSlash, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Preemptivable.contract.WatchLogs(opts, "Slash", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PreemptivableSlash)
				if err := _Preemptivable.contract.UnpackLog(event, "Slash", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSlash is a log parse operation binding the contract event 0xa69f22d963cb7981f842db8c1aafcc93d915ba2a95dcf26dcc333a9c2a09be26.
//
// Solidity: event Slash(address indexed maker, uint256 amount)
func (_Preemptivable *PreemptivableFilterer) ParseSlash(log types.Log) (*PreemptivableSlash, error) {
	event := new(PreemptivableSlash)
	if err := _Preemptivable.contract.UnpackLog(event, "Slash", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PreemptivableStopIterator is returned from FilterStop and is used to iterate over the raw logs and unpacked data for Stop events raised by the Preemptivable contract.
type PreemptivableStopIterator struct {
	Event *PreemptivableStop // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PreemptivableStopIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PreemptivableStop)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PreemptivableStop)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PreemptivableStopIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PreemptivableStopIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PreemptivableStop represents a Stop event raised by the Preemptivable contract.
type PreemptivableStop struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStop is a free log retrieval operation binding the contract event 0xbedf0f4abfe86d4ffad593d9607fe70e83ea706033d44d24b3b6283cf3fc4f6b.
//
// Solidity: event Stop()
func (_Preemptivable *PreemptivableFilterer) FilterStop(opts *bind.FilterOpts) (*PreemptivableStopIterator, error) {

	logs, sub, err := _Preemptivable.contract.FilterLogs(opts, "Stop")
	if err != nil {
		return nil, err
	}
	return &PreemptivableStopIterator{contract: _Preemptivable.contract, event: "Stop", logs: logs, sub: sub}, nil
}

// WatchStop is a free log subscription operation binding the contract event 0xbedf0f4abfe86d4ffad593d9607fe70e83ea706033d44d24b3b6283cf3fc4f6b.
//
// Solidity: event Stop()
func (_Preemptivable *PreemptivableFilterer) WatchStop(opts *bind.WatchOpts, sink chan<- *PreemptivableStop) (event.Subscription, error) {

	logs, sub, err := _Preemptivable.contract.WatchLogs(opts, "Stop")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PreemptivableStop)
				if err := _Preemptivable.contract.UnpackLog(event, "Stop", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStop is a log parse operation binding the contract event 0xbedf0f4abfe86d4ffad593d9607fe70e83ea706033d44d24b3b6283cf3fc4f6b.
//
// Solidity: event Stop()
func (_Preemptivable *PreemptivableFilterer) ParseStop(log types.Log) (*PreemptivableStop, error) {
	event := new(PreemptivableStop)
	if err := _Preemptivable.contract.UnpackLog(event, "Stop", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PreemptivableUnlockIterator is returned from FilterUnlock and is used to iterate over the raw logs and unpacked data for Unlock events raised by the Preemptivable contract.
type PreemptivableUnlockIterator struct {
	Event *PreemptivableUnlock // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PreemptivableUnlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PreemptivableUnlock)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PreemptivableUnlock)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PreemptivableUnlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PreemptivableUnlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PreemptivableUnlock represents a Unlock event raised by the Preemptivable contract.
type PreemptivableUnlock struct {
	Maker common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnlock is a free log retrieval operation binding the contract event 0x0be774851955c26a1d6a32b13b020663a069006b4a3b643ff0b809d318260572.
//
// Solidity: event Unlock(address indexed maker)
func (_Preemptivable *PreemptivableFilterer) FilterUnlock(opts *bind.FilterOpts, maker []common.Address) (*PreemptivableUnlockIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Preemptivable.contract.FilterLogs(opts, "Unlock", makerRule)
	if err != nil {
		return nil, err
	}
	return &PreemptivableUnlockIterator{contract: _Preemptivable.contract, event: "Unlock", logs: logs, sub: sub}, nil
}

// WatchUnlock is a free log subscription operation binding the contract event 0x0be774851955c26a1d6a32b13b020663a069006b4a3b643ff0b809d318260572.
//
// Solidity: event Unlock(address indexed maker)
func (_Preemptivable *PreemptivableFilterer) WatchUnlock(opts *bind.WatchOpts, sink chan<- *PreemptivableUnlock, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Preemptivable.contract.WatchLogs(opts, "Unlock", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PreemptivableUnlock)
				if err := _Preemptivable.contract.UnpackLog(event, "Unlock", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnlock is a log parse operation binding the contract event 0x0be774851955c26a1d6a32b13b020663a069006b4a3b643ff0b809d318260572.
//
// Solidity: event Unlock(address indexed maker)
func (_Preemptivable *PreemptivableFilterer) ParseUnlock(log types.Log) (*PreemptivableUnlock, error) {
	event := new(PreemptivableUnlock)
	if err := _Preemptivable.contract.UnpackLog(event, "Unlock", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SeigniorageABI is the input ABI used to generate the binding from.
const SeigniorageABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"absorptionDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"absorptionExpiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialSlashingPace\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialLockdownExpiration\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"emptive\",\"type\":\"bool\"}],\"name\":\"Absorption\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockdownExpiration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockNumber\",\"type\":\"uint256\"}],\"name\":\"Preemptive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockdownExpiration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashingRate\",\"type\":\"uint256\"}],\"name\":\"Propose\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"}],\"name\":\"Revoke\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Slash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Stop\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"}],\"name\":\"Unlock\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"Ask\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Bid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"index\",\"type\":\"bytes32\"}],\"name\":\"calcOrderID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"haveAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wantAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"assistingID\",\"type\":\"bytes32\"}],\"name\":\"findAssistingID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGlobalParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slashingRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockdownExpiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rank\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLockdown\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"slashingFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getOrder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"have\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"prevID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextID\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getProposal\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"slashingRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockdownExpiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProposalCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRemainToAbsorb\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"next\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"target\",\"type\":\"uint256\"}],\"name\":\"onBlockInitialized\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"prev\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"volatileToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stablizeToken\",\"type\":\"address\"}],\"name\":\"registerTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"}],\"name\":\"revoke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"}],\"name\":\"top\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"}],\"name\":\"totalVote\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"up\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SeigniorageFuncSigs maps the 4-byte function signature to its string representation.
var SeigniorageFuncSigs = map[string]string{
	"69c07d31": "Ask()",
	"6e6452cb": "Bid()",
	"f318722b": "calcOrderID(address,bytes32)",
	"43271d79": "cancel(bool,bytes32)",
	"ced4aac8": "findAssistingID(bool,address,uint256,uint256,bytes32)",
	"6b3222e6": "getGlobalParams()",
	"2a3ed595": "getLockdown()",
	"07c399a3": "getOrder(bool,bytes32)",
	"c7f758a8": "getProposal(uint256)",
	"c08cc02d": "getProposalCount()",
	"ee1a68c6": "getRemainToAbsorb()",
	"4ea09797": "next(bool,bytes32)",
	"be91d729": "onBlockInitialized(uint256)",
	"0d90b10a": "prev(bool,bytes32)",
	"aa1c259c": "registerTokens(address,address)",
	"74a8f103": "revoke(address)",
	"c0ee0b8a": "tokenFallback(address,uint256,bytes)",
	"8aa3f897": "top(bool)",
	"4def5645": "totalVote(address)",
	"bd041c4d": "vote(address,bool)",
}

// SeigniorageBin is the compiled bytecode used for deploying new contracts.
var SeigniorageBin = "0x608060405262049d406003556002600354816200001857fe5b0460045562093a80600e55670de0b6b3a7640000600f55600060105560006011553480156200004657600080fd5b5060405162003b4438038062003b44833981810160405260808110156200006c57600080fd5b50805160208201516040830151606090930151919290918383838383838015620000965760038190555b60008211620000a95760028104620000ab565b815b60045550508015620000bd57600e8190555b8115620000ca57600f8290555b5050505050505050613a6280620000e26000396000f3fe608060405234801561001057600080fd5b506004361061012c5760003560e01c80638aa3f897116100ad578063c0ee0b8a11610071578063c0ee0b8a14610354578063c7f758a8146103d9578063ced4aac814610433578063ee1a68c614610473578063f318722b146104965761012c565b80638aa3f897146102b4578063aa1c259c146102d3578063bd041c4d14610301578063be91d7291461032f578063c08cc02d1461034c5761012c565b80634ea09797116100f45780634ea097971461021757806369c07d311461023c5780636b3222e6146102585780636e6452cb1461028657806374a8f1031461028e5761012c565b806307c399a3146101315780630d90b10a1461018b5780632a3ed595146101c257806343271d79146101ca5780634def5645146101f1575b600080fd5b6101566004803603604081101561014757600080fd5b508035151590602001356104c2565b604080516001600160a01b03909616865260208601949094528484019290925260608401526080830152519081900360a00190f35b6101b0600480360360408110156101a157600080fd5b50803515159060200135610509565b60408051918252519081900360200190f35b610156610531565b6101ef600480360360408110156101e057600080fd5b50803515159060200135610553565b005b6101b06004803603602081101561020757600080fd5b50356001600160a01b03166105d9565b6101b06004803603604081101561022d57600080fd5b50803515159060200135610601565b610244610625565b604080519115158252519081900360200190f35b61026061062a565b604080519485526020850193909352838301919091526060830152519081900360800190f35b61024461063c565b6101ef600480360360208110156102a457600080fd5b50356001600160a01b0316610641565b6101b0600480360360208110156102ca57600080fd5b503515156107db565b6101ef600480360360408110156102e957600080fd5b506001600160a01b03813581169160200135166107f5565b6101ef6004803603604081101561031757600080fd5b506001600160a01b0381351690602001351515610970565b6101ef6004803603602081101561034557600080fd5b50356109ea565b6101b0610b6f565b6101ef6004803603606081101561036a57600080fd5b6001600160a01b038235169160208101359181019060608101604082013564010000000081111561039a57600080fd5b8201836020820111156103ac57600080fd5b803590602001918460018302840111640100000000831117156103ce57600080fd5b509092509050610b81565b6103f6600480360360208110156103ef57600080fd5b5035610ca4565b604080516001600160a01b0390971687526020870195909552858501939093526060850191909152608084015260a0830152519081900360c00190f35b6101b0600480360360a081101561044957600080fd5b5080351515906001600160a01b036020820135169060408101359060608101359060800135610cf8565b61047b610d60565b60408051921515835260208301919091528051918290030190f35b6101b0600480360360408110156104ac57600080fd5b506001600160a01b038135169060200135610e0a565b90151560009081526020818152604080832093835260029384019091529020805460018201549282015460038301546004909301546001600160a01b039092169490929190565b8115156000908152602081815260408083208484526002019091529020600301545b92915050565b600954600b54600a54600c54600d546001600160a01b03909416939091929394565b8115156000908152602081815260408083208484526002810190925290912080546001600160a01b031633146105c3576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037b93232b91036b0b5b2b960811b604482015290519081900360640190fd5b6105d3828463ffffffff610e1d16565b50505050565b6000806105ed60128463ffffffff610f1716565b90506105f881610f35565b9150505b919050565b90151560009081526020818152604080832093835260029093019052206004015490565b600081565b601154600f54600e5460105490919293565b600181565b600061065460128363ffffffff610f1716565b80549091506001600160a01b038381169116146106b8576040805162461bcd60e51b815260206004820152601e60248201527f6f6e6c79206d616b65722063616e207265766f6b652070726f706f73616c0000604482015290519081900360640190fd5b6015805460018101825560009190915260068201805490916003027f55f448fdea98c4d29eb340757ef0a66cd03dbb9538908a6a81d96026b71ec475019061070390829084906137dc565b5050600154825460028401546040805163a9059cbb60e01b81526001600160a01b03938416600482015260248101929092525191909216925063a9059cbb916044808201926020929091908290030181600087803b15801561076457600080fd5b505af1158015610778573d6000803e3d6000fd5b505050506040513d602081101561078e57600080fd5b506107a2905060128363ffffffff61101b16565b506040516001600160a01b038316907f9f77920c3de8baaa98d273e8aa75fae382aaa9f7f60f38979137853e5b73ea2c90600090a25050565b80151560009081526020819052604081206105f881611097565b6001546001600160a01b031615610853576040805162461bcd60e51b815260206004820152601960248201527f566f6c6174696c65546f6b656e20616c72656164792073657400000000000000604482015290519081900360640190fd5b6002546001600160a01b0316156108b1576040805162461bcd60e51b815260206004820152601960248201527f537461626c697a65546f6b656e20616c72656164792073657400000000000000604482015290519081900360640190fd5b600180546001600160a01b038085166001600160a01b03199283161790925560028054928416929091169190911790556108eb82826110b0565b600254604080516318160ddd60e01b815290516000926001600160a01b0316916318160ddd916004808301926020929190829003018186803b15801561093057600080fd5b505afa158015610944573d6000803e3d6000fd5b505050506040513d602081101561095a57600080fd5b5051905061096b8180600080611104565b505050565b61098160128363ffffffff6111b216565b6109c5576040805162461bcd60e51b815260206004820152601060248201526f1b9bc81cdd58da081c1c9bdc1bdcd85b60821b604482015290519081900360640190fd5b60006109d860128463ffffffff610f1716565b905061096b818363ffffffff6111d316565b3315610a2e576040805162461bcd60e51b815260206004820152600e60248201526d636f6e73656e737573206f6e6c7960901b604482015290519081900360640190fd5b60005b601554811015610a6657610a5e60158281548110610a4b57fe5b90600052602060002090600302016111e7565b600101610a31565b50610a7360156000613828565b610a7d6009611254565b15610a8a57610a8a611271565b610a92611374565b508015610b6357610aa360096113e8565b8015610ab65750600854610100900460ff165b15610b6357600254604080516318160ddd60e01b815290516000926001600160a01b0316916318160ddd916004808301926020929190829003018186803b158015610b0057600080fd5b505afa158015610b14573d6000803e3d6000fd5b505050506040513d6020811015610b2a57600080fd5b505190506000610b3a8383611404565b90508015801590610b4f5750610b4f8161141e565b6008805460ff191691151591909117905550505b610b6c81611548565b50565b6000610b7b6012611676565b90505b90565b608081148015610b9b57506001546001600160a01b031633145b15610c3e57610bb160128563ffffffff6111b216565b15610bfc576040805162461bcd60e51b8152602060048201526016602482015275185b1c9958591e481a185cc818481c1c9bdc1bdcd85b60521b604482015290519081900360640190fd5b60008060008085856080811015610c1257600080fd5b50803594506020810135935060400135915060009050610c35888589868661167a565b505050506105d3565b600080806060841415610c725784846060811015610c5b57600080fd5b508035935060208101359250604001359050610c8e565b84846040811015610c8257600080fd5b50803593506020013591505b610c9b87848885856118fc565b50505050505050565b6000808080808080610cbd60128963ffffffff61197d16565b805460028201546001830154600384015460048501546005909501546001600160a01b039094169d929c50909a509850919650945092505050565b8415156000908152602081905260408120610d11613849565b506040805160a0810182526001600160a01b0388168152602081018790529081018590526000606082018190526080820152610d5482828663ffffffff6119b216565b98975050505050505050565b6007546000908190610d7757506000905080610e06565b6001610e01600560020154600260009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b158015610dd057600080fd5b505afa158015610de4573d6000803e3d6000fd5b505050506040513d6020811015610dfa57600080fd5b5051611404565b915091505b9091565b6000610e168383611a24565b9392505050565b610e25613849565b50600081815260028084016020908152604092839020835160a08101855281546001600160a01b03168152600182015492810183905292810154938301939093526003830154606083015260049092015460808201529015610f0757825481516020808401516040805163a9059cbb60e01b81526001600160a01b039485166004820152602481019290925251929093169263a9059cbb92604480830193928290030181600087803b158015610eda57600080fd5b505af1158015610eee573d6000803e3d6000fd5b505050506040513d6020811015610f0457600080fd5b50505b61096b838363ffffffff611af116565b6001600160a01b031660009081526002919091016020526040902090565b600080805b610f4684600601611676565b81101561101457600080610f63600687018463ffffffff611b5116565b600154604080516370a0823160e01b81526001600160a01b03808616600483015291519496509294506000939116916370a08231916024808301926020929190829003018186803b158015610fb757600080fd5b505afa158015610fcb573d6000803e3d6000fd5b505050506040513d6020811015610fe157600080fd5b50516001600160a01b0384163101905081156110005793840193611006565b80850394505b505050806001019050610f3a565b5092915050565b6001600160a01b038116600090815260018301602052604081205480611078576040805162461bcd60e51b815260206004820152600d60248201526c1ad95e481b9bdd08195e1a5cdd609a1b604482015290519081900360640190fd5b61108d8460001983018563ffffffff611b9016565b5060019392505050565b6000808052600282016020526040902060040154919050565b60008080526020526110d7600080516020613a0d833981519152838363ffffffff611ce216565b600160009081526020526111006000805160206139cb833981519152828463ffffffff611ce216565b5050565b6040805160a0810182526003544301808252602082018690529181018690526000606082018190528315156080909201829052600592909255600685905560078690556008805461ffff19166101009092029190911790556111668585611404565b60408051828152602081018790528515158183015290519192507f0427b353dc7214e3d8c7f5039475a8e729f4d62922937381e304cd03becf66d2919081900360600190a15050505050565b6001600160a01b031660009081526001919091016020526040902054151590565b61096b60068301338363ffffffff611e8f16565b60005b815481101561124857600082600001828154811061120457fe5b60009182526020808320909101546001600160a01b03168252600185810182526040808420849055600287019092529120805460ff191690559190910190506111ea565b50610b6c816000613877565b600061125f82611fad565b801561052b5750506004015443101590565b61127b6009611fad565b61128457611372565b600b541561131557600154600954600b546040805163a9059cbb60e01b81526001600160a01b039384166004820152602481019290925251919092169163a9059cbb9160448083019260209291908290030181600087803b1580156112e857600080fd5b505af11580156112fc573d6000803e3d6000fd5b505050506040513d602081101561131257600080fd5b50505b6009546040516001600160a01b03909116907f0be774851955c26a1d6a32b13b020663a069006b4a3b643ff0b809d31826057290600090a2600980546001600160a01b03191690556000600a819055600b819055600c819055600d555b565b600061138060096113e8565b1561138d57506000610b7e565b600080611398611fbc565b90925090506001600160a01b0382166113b657600092505050610b7e565b60006113c960128463ffffffff610f1716565b90506113d58183612093565b6113de816120e4565b6001935050505090565b60006113f382611fad565b801561052b57505060040154431090565b600081831161141857828203600003610e16565b50900390565b6000611431600960010154600084612297565b61143d575060006105fc565b6000670de0b6b3a764000061145c611454856122c7565b600c546122dd565b8161146357fe5b0490508060096002015410156114785750600b545b600b805482900390556001546040805163117f5a5560e01b81526004810184905290516001600160a01b039092169163117f5a559160248082019260009290919082900301818387803b1580156114ce57600080fd5b505af11580156114e2573d6000803e3d6000fd5b50506009546040805185815290516001600160a01b0390921693507fa69f22d963cb7981f842db8c1aafcc93d915ba2a95dcf26dcc333a9c2a09be26925081900360200190a2600b5461153f57611537612312565b61153f611271565b50600192915050565b331561158c576040805162461bcd60e51b815260206004820152600e60248201526d636f6e73656e737573206f6e6c7960901b604482015290519081900360640190fd5b6115966005612357565b156115a3576115a3612312565b600254604080516318160ddd60e01b815290516000926001600160a01b0316916318160ddd916004808301926020929190829003018186803b1580156115e857600080fd5b505afa1580156115fc573d6000803e3d6000fd5b505050506040513d602081101561161257600080fd5b50519050811561165857611624612370565b1561163b576116368282600080611104565b611658565b611645818361237c565b1561165857611658828260016000611104565b61166960058263ffffffff61242a16565b1561110057611100612470565b5490565b60036011548161168657fe5b04601154038310156116cf576040805162461bcd60e51b815260206004820152600d60248201526c7374616b6520746f6f206c6f7760981b604482015290519081900360640190fd5b83611713576040805162461bcd60e51b815260206004820152600f60248201526e3d32b9379030b139b7b9383a34b7b760891b604482015290519081900360640190fd5b61171b613895565b8215611792576003600f548161172d57fe5b04600f5403831015611786576040805162461bcd60e51b815260206004820152601b60248201527f736c617368696e67207261746520706172616d20746f6f206c6f770000000000604482015290519081900360640190fd5b6060810183905261179b565b600f5460608201525b60006117a6866122c7565b6117b48684606001516122dd565b816117bb57fe5b049050600081116117fd5760405162461bcd60e51b81526004018080602001828103825260228152602001806139eb6022913960400191505060405180910390fd5b821561185e576003600e548161180f57fe5b04600e54038310156118525760405162461bcd60e51b81526004018080602001828103825260218152602001806139aa6021913960400191505060405180910390fd5b60808201839052611867565b600e5460808301525b6001600160a01b038716825260208201869052604082018590524360a083015261189860128363ffffffff61281216565b50606080830151608080850151604080518b8152602081018b9052808201949094529383015291516001600160a01b038a16927f56e25d1b63c01627fcd54936462c97aeb9a18352bf0ed161e8141a33cfd795ca928290030190a250505050505050565b6000611906613849565b61191287878787612966565b91509150600061192133612a5e565b905061193e8261193033612b41565b83919063ffffffff612bfb16565b61194782612cf6565b156119615761195c818363ffffffff612d0f16565b611973565b6119738184848763ffffffff612db116565b5050505050505050565b600080611990848463ffffffff612e8e16565b6001600160a01b03166000908152600285016020526040902091505092915050565b600081815260028401602052604081205b60040154600081815260028601602052604090209092506119ea848263ffffffff612ebb16565b156119c3575b6003015460008181526002860160205260409020909250611a17848263ffffffff612ebb16565b6119f05750909392505050565b60006002838360405160200180836001600160a01b03166001600160a01b031660601b8152601401828152602001925050506040516020818303038152906040526040518082805190602001908083835b60208310611a945780518252601f199092019160209182019101611a75565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015611ad3573d6000803e3d6000fd5b5050506040513d6020811015611ae857600080fd5b50519392505050565b6000818152600292830160205260408082206004808201805460038085018054885286882090940182905583549187529486209094019390935593835280546001600160a01b031916815560018101839055909301819055908190559055565b60008080611b65858563ffffffff612e8e16565b6001600160a01b038116600090815260028701602052604090205490935060ff169150509250929050565b6001600160a01b03811660009081526001808501602090815260408084208490556002808801909252832080546001600160a01b03191681559182018390558101829055600381018290556004810182905560058101829055906006820181611bf98282613877565b505084546000190184149150611cb0905057825483906000198101908110611c1d57fe5b60009182526020909120015483546001600160a01b0390911690849084908110611c4357fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555081600101836001016000856000018581548110611c8a57fe5b60009182526020808320909101546001600160a01b031683528201929092526040019020555b8254839080611cbb57fe5b600082815260209020810160001990810180546001600160a01b0319169055019055505050565b818360000160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550808360010160006101000a8154816001600160a01b0302191690836001600160a01b031602179055506040518060a00160405280306001600160a01b0316815260200160008152602001600081526020016000801b815260200160001960001b8152508360020160008060001b815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550602082015181600101556040820151816002015560608201518160030155608082015181600401559050506040518060a00160405280306001600160a01b0316815260200160008152602001600181526020016000801b815260200160001960001b81525083600201600060001960001b815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060208201518160010155604082015181600201556060820151816003015560808201518160040155905050505050565b6001600160a01b03821660009081526002840160209081526040808320805460ff19168515151790556001860190915281205480611f135750508254600180820185556000858152602080822090930180546001600160a01b0319166001600160a01b03871690811790915586549082528287019093526040902091909155610e16565b8454811180611f515750836001600160a01b0316856000016001830381548110611f3957fe5b6000918252602090912001546001600160a01b031614155b15611fa25750508254600180820185556000858152602080822090930180546001600160a01b0319166001600160a01b03871690811790915586549082528287019093526040902091909155610e16565b506000949350505050565b546001600160a01b0316151590565b6000806000600360105481611fcd57fe5b0460105403905060006004600e5481611fe257fe5b04905060008080805b611ff56012611676565b81101561207057600061200f60128363ffffffff61197d16565b9050600061201c82612ed9565b90508781121561202d575050612068565b858113156120475781549095506001600160a01b03169350845b8315801561205b5750816005015443038711155b1561206557600193505b50505b600101611feb565b5080612088575060009550859450610e069350505050565b509450925050509091565b6120a36011548360020154612f07565b601155600f5460038301546120b89190612f07565b600f55600e5460048301546120cd9190612f07565b600e556010546120dd9082612f07565b6010555050565b6120f0816006016111e7565b60006120ff82600101546122c7565b612111836002015484600301546122dd565b8161211857fe5b6040805160a08101825285546001600160a01b0390811680835260018801546020840181905260028901549484018590529590940460608301819052600488015443016080909301839052600980546001600160a01b031916909517909455600a94909455600b91909155600c829055600d5583549092506121a3916012911663ffffffff61101b16565b50600254604080516318160ddd60e01b815290516000926001600160a01b0316916318160ddd916004808301926020929190829003018186803b1580156121e957600080fd5b505afa1580156121fd573d6000803e3d6000fd5b505050506040513d602081101561221357600080fd5b5051600a54909150600090612229908390612f62565b90506122388183600180611104565b835460028501546003860154600d5460408051938452602084019290925282820152516001600160a01b03909216917f8427e4488966b7bd3193a4617993e5e6b9186f0c4b2c303cc6178f4e33b77d089181900360600190a250505050565b60008284131580156122a95750818313155b806122bf57508284121580156122bf5750818312155b949350505050565b60008082136122d9578160000361052b565b5090565b6000826122ec5750600061052b565b828202828482816122f957fe5b04141561230757905061052b565b506000199392505050565b60006005819055600681905560078190556008805461ffff191690556040517fbedf0f4abfe86d4ffad593d9607fe70e83ea706033d44d24b3b6283cf3fc4f6b9190a1565b600061236282612f7e565b801561052b57505054431190565b6000610b7b6005612357565b60008282141561238e5750600061052b565b60065460075414156123a25750600161052b565b828211156123f6576006546007548484039110156123d8576006546007540360028183816123cc57fe5b0410159250505061052b565b6007546006540360028282816123ea57fe5b0411159250505061052b565b600654600754838503911115612418576007546006540360028183816123cc57fe5b6006546007540360028282816123ea57fe5b600061243583612f7e565b80156124465750600383015460ff16155b8015612456575082600201548214155b8015610e165750610e168360010154838560020154612f84565b600061247a612faf565b9050600080600080841361248f576001612492565b60005b151581526020810191909152604001600090812060025481549193506001600160a01b0391821691161490806124d9836124cb876122c7565b86919063ffffffff61307b16565b6008549193509150610100900460ff166124f7575050505050611372565b801580612502575081155b15612511575050505050611372565b600254604080516318160ddd60e01b815290516000926001600160a01b0316916318160ddd916004808301926020929190829003018186803b15801561255657600080fd5b505afa15801561256a573d6000803e3d6000fd5b505050506040513d602081101561258057600080fd5b5051905061259560058263ffffffff61242a16565b6125a457505050505050611372565b600080856125b35784846125b6565b83855b6009548954604080516370a0823160e01b81526001600160a01b03938416600482018190529151959750939550939116916370a08231916024808301926020929190829003018186803b15801561260c57600080fd5b505afa158015612620573d6000803e3d6000fd5b505050506040513d602081101561263657600080fd5b505183111561264d57505050505050505050611372565b8754604080516337ed875b60e11b81526001600160a01b0384811660048301526024820187905291519190921691636fdb0eb691604480830192600092919082900301818387803b1580156126a157600080fd5b505af11580156126b5573d6000803e3d6000fd5b505089546040805163117f5a5560e01b81526004810188905290516001600160a01b03909216935063117f5a55925060248082019260009290919082900301818387803b15801561270557600080fd5b505af1158015612719573d6000803e3d6000fd5b50505060018901546040805163bdfde91160e01b81526004810186905290516001600160a01b03909216925063bdfde91191602480830192600092919082900301818387803b15801561276b57600080fd5b505af115801561277f573d6000803e3d6000fd5b5050505060018801546040805163a9059cbb60e01b81526001600160a01b038481166004830152602482018690529151919092169163a9059cbb9160448083019260209291908290030181600087803b1580156127db57600080fd5b505af11580156127ef573d6000803e3d6000fd5b505050506040513d602081101561280557600080fd5b5050505050505050505050565b80516001600160a01b03811660009081526001840160205260408120549091908015612885576040805162461bcd60e51b815260206004820152601c60248201527f6d616b657220616c72656164792068617320612070726f706f73616c00000000604482015290519081900360640190fd5b6001600160a01b03828116600090815260028781016020908152604092839020885181546001600160a01b03191695169490941784558781015160018501559187015190830155606086015160038301556080860151600483015560a0860151600583015560c086015180518051889493600685019261290b92849291909101906138e0565b5050875460018082018a5560008a8152602080822090930180546001600160a01b0319166001600160a01b039990991698891790558a54978152908a0182526040808220979097556002909901905250505093209392505050565b6000612970613849565b6000841180156129805750600083115b6129be576040805162461bcd60e51b815260206004820152600a6024820152691e995c9bc81a5b9c1d5d60b21b604482015290519081900360640190fd5b600160801b841080156129d45750600160801b83105b612a18576040805162461bcd60e51b815260206004820152601060248201526f1a5b9c1d5d081bdd995c881b1a5b5a5d60821b604482015290519081900360640190fd5b612a228686611a24565b6040805160a0810182526001600160a01b0398909816885260208801959095529386019290925250506000606084018190526080840152929050565b60008080526020819052600080516020613a0d833981519152546001600160a01b0383811691161415612aa857506000808052602052600080516020613a0d8339815191526105fc565b600160009081526020526000805160206139cb833981519152546001600160a01b0383811691161415612af45750600160009081526020526000805160206139cb8339815191526105fc565b6040805162461bcd60e51b815260206004820152601760248201527f6e6f206f7264657220626f6f6b20666f7220746f6b656e000000000000000000604482015290519081900360640190fd5b600080805260208190527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb6546001600160a01b0383811691161415612b9d57506000808052602052600080516020613a0d8339815191526105fc565b600160009081526020527fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7e546001600160a01b0383811691161415612af45750600160009081526020526000805160206139cb8339815191526105fc565b6000612c0682611097565b90505b60001981146105d357612c1b83612cf6565b15612c25576105d3565b60008181526002830160205260409020612c45848263ffffffff61332216565b612c4f57506105d3565b806001015484604001511015612cb75760008160010154826002015486604001510281612c7857fe5b049050612c968386604001518387613341909392919063ffffffff16565b6040850151612cb09087908790849063ffffffff6134e916565b50506105d3565b60028101546001820154612cd591879187919063ffffffff6134e916565b612ce5838363ffffffff61362316565b612cee83611097565b915050612c09565b600081602001516000148061052b575050604001511590565b602081015115612d9f57815481516020808401516040805163a9059cbb60e01b81526001600160a01b039485166004820152602481019290925251929093169263a9059cbb92604480830193928290030181600087803b158015612d7257600080fd5b505af1158015612d86573d6000803e3d6000fd5b505050506040513d6020811015612d9c57600080fd5b50505b60006020820181905260409091015250565b60008381526002850160205260409020612dca90611fad565b15612e11576040805162461bcd60e51b81526020600482015260126024820152716f7264657220696e6465782065786973747360701b604482015290519081900360640190fd5b6000838152600285810160209081526040808420865181546001600160a01b0319166001600160a01b0390911617815591860151600183015585015191810191909155606084015160038201556080840151600490910155612e748584846119b2565b9050612e8785858363ffffffff6136e616565b5050505050565b6000826000018281548110612e9f57fe5b6000918252602090912001546001600160a01b03169392505050565b60408201516001820154600290920154602090930151910291021190565b600080612ee583610f35565b905060008113612ef95760009150506105fc565b6105f8818460020154613721565b600082821415612f1857508161052b565b6000612f248484613779565b905083831080612f32575083155b15612f3e57905061052b565b83810360038504811115612f5a5750505060038204820161052b565b509392505050565b600080821215612f7957816000038303905061052b565b500190565b54151590565b6000828411158015612f965750818311155b806122bf57508284101580156122bf5750501115919050565b600080612fc6600560020154600560010154611404565b600754600254604080516318160ddd60e01b8152905193945060009361301693926001600160a01b0316916318160ddd916004808301926020929190829003018186803b158015610dd057600080fd5b905061302460008284612297565b61303357600092505050610b7e565b6000600454838161304057fe5b6008549190059150610100900460ff161561305a57600290055b61306660008284612297565b61307457509150610b7e9050565b9250505090565b600080600061308986611097565b90505b600019811480159061309d57508382105b15613318576000818152600287016020526040812090866130c25781600201546130c8565b81600101545b90506000876130db5782600101546130e1565b82600201545b90508487038083116131d757895460018501546040805163117f5a5560e01b81526004810192909252516001600160a01b039092169163117f5a559160248082019260009290919082900301818387803b15801561313e57600080fd5b505af1158015613152573d6000803e3d6000fd5b50505060018b015460028601546040805163bdfde91160e01b81526004810192909252516001600160a01b03909216925063bdfde91191602480830192600092919082900301818387803b1580156131a957600080fd5b505af11580156131bd573d6000803e3d6000fd5b5050505060048401546131d08b87613623565b9450613307565b82818302816131e257fe5b0491508092508682880110156131fd575061331a9350505050565b60008961320a578261320c565b835b905060008a61321b578461321d565b835b8c546040805163117f5a5560e01b81526004810186905290519293506001600160a01b039091169163117f5a559160248082019260009290919082900301818387803b15801561326c57600080fd5b505af1158015613280573d6000803e3d6000fd5b50505060018d01546040805163bdfde91160e01b81526004810185905290516001600160a01b03909216925063bdfde91191602480830192600092919082900301818387803b1580156132d257600080fd5b505af11580156132e6573d6000803e3d6000fd5b506132ff92508e9150899050848463ffffffff61334116565b506000199550505b50949094019392909201915061308c565b505b935093915050565b6040820151600282015460019092015460209093015191029102101590565b6000838152600285016020526040902060018101548311156133a0576040805162461bcd60e51b815260206004820152601360248201527250503a2066696c6c61626c65203e206861766560681b604482015290519081900360640190fd5b80600201548211156133ef576040805162461bcd60e51b815260206004820152601360248201527214140e88199a5b1b18589b19480f881dd85b9d606a1b604482015290519081900360640190fd5b600180820180548590039055600282018054849003905585015481546040805163a9059cbb60e01b81526001600160a01b039283166004820152602481018690529051919092169163a9059cbb9160448083019260209291908290030181600087803b15801561345e57600080fd5b505af1158015613472573d6000803e3d6000fd5b505050506040513d602081101561348857600080fd5b50506040805160a08101825282546001600160a01b031681526001830154602082015260028301549181019190915260038201546060820152600482015460808201526134d490612cf6565b15612e8757612e87858563ffffffff610e1d16565b8260200151821115613538576040805162461bcd60e51b815260206004820152601360248201527250503a2066696c6c61626c65203e206861766560681b604482015290519081900360640190fd5b8260400151811115613587576040805162461bcd60e51b815260206004820152601360248201527214140e88199a5b1b18589b19480f881dd85b9d606a1b604482015290519081900360640190fd5b60208084018051849003905260408085018051849003905260018601548551825163a9059cbb60e01b81526001600160a01b03918216600482015260248101869052925191169263a9059cbb92604480820193918290030181600087803b1580156135f157600080fd5b505af1158015613605573d6000803e3d6000fd5b505050506040513d602081101561361b57600080fd5b505050505050565b61362b613849565b50600081815260028084016020908152604092839020835160a08101855281546001600160a01b03168152600182015492810192909252918201549281018390526003820154606082015260049091015460808201529015610f075760018301548151604080840151815163a9059cbb60e01b81526001600160a01b03938416600482015260248101919091529051919092169163a9059cbb9160448083019260209291908290030181600087803b158015610eda57600080fd5b600081815260029093016020526040808420600490810180548587528387206003808201879055930181905586529185200183905592529055565b6000826137305750600061052b565b8282028284828161373d57fe5b05141561374b57905061052b565b613757846000856137b5565b156137695750600160ff1b905061052b565b506001600160ff1b039392505050565b600082820183811061378f5760011c905061052b565b600184811c9084901c018481106137a957915061052b9050565b50600019949350505050565b600082841280156137c557508183125b806122bf575082841380156122bf57505012919050565b82805482825590600052602060002090810192821561381c5760005260206000209182015b8281111561381c578254825591600101919060010190613801565b506122d9929150613935565b5080546000825560030290600052602060002090810190610b6c9190613959565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915290565b5080546000825590600052602060002090810190610b6c919061397c565b6040518060e0016040528060006001600160a01b0316815260200160008152602001600081526020016000815260200160008152602001600081526020016138db613996565b905290565b82805482825590600052602060002090810192821561381c579160200282015b8281111561381c57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190613900565b610b7e91905b808211156122d95780546001600160a01b031916815560010161393b565b610b7e91905b808211156122d95760006139738282613877565b5060030161395f565b610b7e91905b808211156122d95760008155600101613982565b604051806020016040528060608152509056fe6c6f636b646f776e206475726174696f6e20706172616d20746f6f2073686f7274ada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d736c617368696e6720666163746f722063616c63756c6174656420746f207a65726fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5a2646970667358221220b3185f0d537d81c15e179bdbbea55c34f1f4e0cb67b77810af9fe4bc5f36198764736f6c63430006080033"

// DeploySeigniorage deploys a new Ethereum contract, binding an instance of Seigniorage to it.
func DeploySeigniorage(auth *bind.TransactOpts, backend bind.ContractBackend, absorptionDuration *big.Int, absorptionExpiration *big.Int, initialSlashingPace *big.Int, initialLockdownExpiration *big.Int) (common.Address, *types.Transaction, *Seigniorage, error) {
	parsed, err := abi.JSON(strings.NewReader(SeigniorageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SeigniorageBin), backend, absorptionDuration, absorptionExpiration, initialSlashingPace, initialLockdownExpiration)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Seigniorage{SeigniorageCaller: SeigniorageCaller{contract: contract}, SeigniorageTransactor: SeigniorageTransactor{contract: contract}, SeigniorageFilterer: SeigniorageFilterer{contract: contract}}, nil
}

// Seigniorage is an auto generated Go binding around an Ethereum contract.
type Seigniorage struct {
	SeigniorageCaller     // Read-only binding to the contract
	SeigniorageTransactor // Write-only binding to the contract
	SeigniorageFilterer   // Log filterer for contract events
}

// SeigniorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type SeigniorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SeigniorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SeigniorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SeigniorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SeigniorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SeigniorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SeigniorageSession struct {
	Contract     *Seigniorage      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SeigniorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SeigniorageCallerSession struct {
	Contract *SeigniorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SeigniorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SeigniorageTransactorSession struct {
	Contract     *SeigniorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SeigniorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type SeigniorageRaw struct {
	Contract *Seigniorage // Generic contract binding to access the raw methods on
}

// SeigniorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SeigniorageCallerRaw struct {
	Contract *SeigniorageCaller // Generic read-only contract binding to access the raw methods on
}

// SeigniorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SeigniorageTransactorRaw struct {
	Contract *SeigniorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSeigniorage creates a new instance of Seigniorage, bound to a specific deployed contract.
func NewSeigniorage(address common.Address, backend bind.ContractBackend) (*Seigniorage, error) {
	contract, err := bindSeigniorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Seigniorage{SeigniorageCaller: SeigniorageCaller{contract: contract}, SeigniorageTransactor: SeigniorageTransactor{contract: contract}, SeigniorageFilterer: SeigniorageFilterer{contract: contract}}, nil
}

// NewSeigniorageCaller creates a new read-only instance of Seigniorage, bound to a specific deployed contract.
func NewSeigniorageCaller(address common.Address, caller bind.ContractCaller) (*SeigniorageCaller, error) {
	contract, err := bindSeigniorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SeigniorageCaller{contract: contract}, nil
}

// NewSeigniorageTransactor creates a new write-only instance of Seigniorage, bound to a specific deployed contract.
func NewSeigniorageTransactor(address common.Address, transactor bind.ContractTransactor) (*SeigniorageTransactor, error) {
	contract, err := bindSeigniorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SeigniorageTransactor{contract: contract}, nil
}

// NewSeigniorageFilterer creates a new log filterer instance of Seigniorage, bound to a specific deployed contract.
func NewSeigniorageFilterer(address common.Address, filterer bind.ContractFilterer) (*SeigniorageFilterer, error) {
	contract, err := bindSeigniorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SeigniorageFilterer{contract: contract}, nil
}

// bindSeigniorage binds a generic wrapper to an already deployed contract.
func bindSeigniorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SeigniorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Seigniorage *SeigniorageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Seigniorage.Contract.SeigniorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Seigniorage *SeigniorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Seigniorage.Contract.SeigniorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Seigniorage *SeigniorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Seigniorage.Contract.SeigniorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Seigniorage *SeigniorageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Seigniorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Seigniorage *SeigniorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Seigniorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Seigniorage *SeigniorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Seigniorage.Contract.contract.Transact(opts, method, params...)
}

// Ask is a free data retrieval call binding the contract method 0x69c07d31.
//
// Solidity: function Ask() view returns(bool)
func (_Seigniorage *SeigniorageCaller) Ask(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Seigniorage.contract.Call(opts, out, "Ask")
	return *ret0, err
}

// Ask is a free data retrieval call binding the contract method 0x69c07d31.
//
// Solidity: function Ask() view returns(bool)
func (_Seigniorage *SeigniorageSession) Ask() (bool, error) {
	return _Seigniorage.Contract.Ask(&_Seigniorage.CallOpts)
}

// Ask is a free data retrieval call binding the contract method 0x69c07d31.
//
// Solidity: function Ask() view returns(bool)
func (_Seigniorage *SeigniorageCallerSession) Ask() (bool, error) {
	return _Seigniorage.Contract.Ask(&_Seigniorage.CallOpts)
}

// Bid is a free data retrieval call binding the contract method 0x6e6452cb.
//
// Solidity: function Bid() view returns(bool)
func (_Seigniorage *SeigniorageCaller) Bid(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Seigniorage.contract.Call(opts, out, "Bid")
	return *ret0, err
}

// Bid is a free data retrieval call binding the contract method 0x6e6452cb.
//
// Solidity: function Bid() view returns(bool)
func (_Seigniorage *SeigniorageSession) Bid() (bool, error) {
	return _Seigniorage.Contract.Bid(&_Seigniorage.CallOpts)
}

// Bid is a free data retrieval call binding the contract method 0x6e6452cb.
//
// Solidity: function Bid() view returns(bool)
func (_Seigniorage *SeigniorageCallerSession) Bid() (bool, error) {
	return _Seigniorage.Contract.Bid(&_Seigniorage.CallOpts)
}

// CalcOrderID is a free data retrieval call binding the contract method 0xf318722b.
//
// Solidity: function calcOrderID(address maker, bytes32 index) pure returns(bytes32)
func (_Seigniorage *SeigniorageCaller) CalcOrderID(opts *bind.CallOpts, maker common.Address, index [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Seigniorage.contract.Call(opts, out, "calcOrderID", maker, index)
	return *ret0, err
}

// CalcOrderID is a free data retrieval call binding the contract method 0xf318722b.
//
// Solidity: function calcOrderID(address maker, bytes32 index) pure returns(bytes32)
func (_Seigniorage *SeigniorageSession) CalcOrderID(maker common.Address, index [32]byte) ([32]byte, error) {
	return _Seigniorage.Contract.CalcOrderID(&_Seigniorage.CallOpts, maker, index)
}

// CalcOrderID is a free data retrieval call binding the contract method 0xf318722b.
//
// Solidity: function calcOrderID(address maker, bytes32 index) pure returns(bytes32)
func (_Seigniorage *SeigniorageCallerSession) CalcOrderID(maker common.Address, index [32]byte) ([32]byte, error) {
	return _Seigniorage.Contract.CalcOrderID(&_Seigniorage.CallOpts, maker, index)
}

// FindAssistingID is a free data retrieval call binding the contract method 0xced4aac8.
//
// Solidity: function findAssistingID(bool orderType, address maker, uint256 haveAmount, uint256 wantAmount, bytes32 assistingID) view returns(bytes32)
func (_Seigniorage *SeigniorageCaller) FindAssistingID(opts *bind.CallOpts, orderType bool, maker common.Address, haveAmount *big.Int, wantAmount *big.Int, assistingID [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Seigniorage.contract.Call(opts, out, "findAssistingID", orderType, maker, haveAmount, wantAmount, assistingID)
	return *ret0, err
}

// FindAssistingID is a free data retrieval call binding the contract method 0xced4aac8.
//
// Solidity: function findAssistingID(bool orderType, address maker, uint256 haveAmount, uint256 wantAmount, bytes32 assistingID) view returns(bytes32)
func (_Seigniorage *SeigniorageSession) FindAssistingID(orderType bool, maker common.Address, haveAmount *big.Int, wantAmount *big.Int, assistingID [32]byte) ([32]byte, error) {
	return _Seigniorage.Contract.FindAssistingID(&_Seigniorage.CallOpts, orderType, maker, haveAmount, wantAmount, assistingID)
}

// FindAssistingID is a free data retrieval call binding the contract method 0xced4aac8.
//
// Solidity: function findAssistingID(bool orderType, address maker, uint256 haveAmount, uint256 wantAmount, bytes32 assistingID) view returns(bytes32)
func (_Seigniorage *SeigniorageCallerSession) FindAssistingID(orderType bool, maker common.Address, haveAmount *big.Int, wantAmount *big.Int, assistingID [32]byte) ([32]byte, error) {
	return _Seigniorage.Contract.FindAssistingID(&_Seigniorage.CallOpts, orderType, maker, haveAmount, wantAmount, assistingID)
}

// GetGlobalParams is a free data retrieval call binding the contract method 0x6b3222e6.
//
// Solidity: function getGlobalParams() view returns(uint256 stake, uint256 slashingRate, uint256 lockdownExpiration, uint256 rank)
func (_Seigniorage *SeigniorageCaller) GetGlobalParams(opts *bind.CallOpts) (struct {
	Stake              *big.Int
	SlashingRate       *big.Int
	LockdownExpiration *big.Int
	Rank               *big.Int
}, error) {
	ret := new(struct {
		Stake              *big.Int
		SlashingRate       *big.Int
		LockdownExpiration *big.Int
		Rank               *big.Int
	})
	out := ret
	err := _Seigniorage.contract.Call(opts, out, "getGlobalParams")
	return *ret, err
}

// GetGlobalParams is a free data retrieval call binding the contract method 0x6b3222e6.
//
// Solidity: function getGlobalParams() view returns(uint256 stake, uint256 slashingRate, uint256 lockdownExpiration, uint256 rank)
func (_Seigniorage *SeigniorageSession) GetGlobalParams() (struct {
	Stake              *big.Int
	SlashingRate       *big.Int
	LockdownExpiration *big.Int
	Rank               *big.Int
}, error) {
	return _Seigniorage.Contract.GetGlobalParams(&_Seigniorage.CallOpts)
}

// GetGlobalParams is a free data retrieval call binding the contract method 0x6b3222e6.
//
// Solidity: function getGlobalParams() view returns(uint256 stake, uint256 slashingRate, uint256 lockdownExpiration, uint256 rank)
func (_Seigniorage *SeigniorageCallerSession) GetGlobalParams() (struct {
	Stake              *big.Int
	SlashingRate       *big.Int
	LockdownExpiration *big.Int
	Rank               *big.Int
}, error) {
	return _Seigniorage.Contract.GetGlobalParams(&_Seigniorage.CallOpts)
}

// GetLockdown is a free data retrieval call binding the contract method 0x2a3ed595.
//
// Solidity: function getLockdown() view returns(address maker, uint256 stake, int256 amount, uint256 slashingFactor, uint256 unlockNumber)
func (_Seigniorage *SeigniorageCaller) GetLockdown(opts *bind.CallOpts) (struct {
	Maker          common.Address
	Stake          *big.Int
	Amount         *big.Int
	SlashingFactor *big.Int
	UnlockNumber   *big.Int
}, error) {
	ret := new(struct {
		Maker          common.Address
		Stake          *big.Int
		Amount         *big.Int
		SlashingFactor *big.Int
		UnlockNumber   *big.Int
	})
	out := ret
	err := _Seigniorage.contract.Call(opts, out, "getLockdown")
	return *ret, err
}

// GetLockdown is a free data retrieval call binding the contract method 0x2a3ed595.
//
// Solidity: function getLockdown() view returns(address maker, uint256 stake, int256 amount, uint256 slashingFactor, uint256 unlockNumber)
func (_Seigniorage *SeigniorageSession) GetLockdown() (struct {
	Maker          common.Address
	Stake          *big.Int
	Amount         *big.Int
	SlashingFactor *big.Int
	UnlockNumber   *big.Int
}, error) {
	return _Seigniorage.Contract.GetLockdown(&_Seigniorage.CallOpts)
}

// GetLockdown is a free data retrieval call binding the contract method 0x2a3ed595.
//
// Solidity: function getLockdown() view returns(address maker, uint256 stake, int256 amount, uint256 slashingFactor, uint256 unlockNumber)
func (_Seigniorage *SeigniorageCallerSession) GetLockdown() (struct {
	Maker          common.Address
	Stake          *big.Int
	Amount         *big.Int
	SlashingFactor *big.Int
	UnlockNumber   *big.Int
}, error) {
	return _Seigniorage.Contract.GetLockdown(&_Seigniorage.CallOpts)
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) view returns(address maker, uint256 have, uint256 want, bytes32 prevID, bytes32 nextID)
func (_Seigniorage *SeigniorageCaller) GetOrder(opts *bind.CallOpts, _orderType bool, _id [32]byte) (struct {
	Maker  common.Address
	Have   *big.Int
	Want   *big.Int
	PrevID [32]byte
	NextID [32]byte
}, error) {
	ret := new(struct {
		Maker  common.Address
		Have   *big.Int
		Want   *big.Int
		PrevID [32]byte
		NextID [32]byte
	})
	out := ret
	err := _Seigniorage.contract.Call(opts, out, "getOrder", _orderType, _id)
	return *ret, err
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) view returns(address maker, uint256 have, uint256 want, bytes32 prevID, bytes32 nextID)
func (_Seigniorage *SeigniorageSession) GetOrder(_orderType bool, _id [32]byte) (struct {
	Maker  common.Address
	Have   *big.Int
	Want   *big.Int
	PrevID [32]byte
	NextID [32]byte
}, error) {
	return _Seigniorage.Contract.GetOrder(&_Seigniorage.CallOpts, _orderType, _id)
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) view returns(address maker, uint256 have, uint256 want, bytes32 prevID, bytes32 nextID)
func (_Seigniorage *SeigniorageCallerSession) GetOrder(_orderType bool, _id [32]byte) (struct {
	Maker  common.Address
	Have   *big.Int
	Want   *big.Int
	PrevID [32]byte
	NextID [32]byte
}, error) {
	return _Seigniorage.Contract.GetOrder(&_Seigniorage.CallOpts, _orderType, _id)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 idx) view returns(address maker, uint256 stake, int256 amount, uint256 slashingRate, uint256 lockdownExpiration, uint256 number)
func (_Seigniorage *SeigniorageCaller) GetProposal(opts *bind.CallOpts, idx *big.Int) (struct {
	Maker              common.Address
	Stake              *big.Int
	Amount             *big.Int
	SlashingRate       *big.Int
	LockdownExpiration *big.Int
	Number             *big.Int
}, error) {
	ret := new(struct {
		Maker              common.Address
		Stake              *big.Int
		Amount             *big.Int
		SlashingRate       *big.Int
		LockdownExpiration *big.Int
		Number             *big.Int
	})
	out := ret
	err := _Seigniorage.contract.Call(opts, out, "getProposal", idx)
	return *ret, err
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 idx) view returns(address maker, uint256 stake, int256 amount, uint256 slashingRate, uint256 lockdownExpiration, uint256 number)
func (_Seigniorage *SeigniorageSession) GetProposal(idx *big.Int) (struct {
	Maker              common.Address
	Stake              *big.Int
	Amount             *big.Int
	SlashingRate       *big.Int
	LockdownExpiration *big.Int
	Number             *big.Int
}, error) {
	return _Seigniorage.Contract.GetProposal(&_Seigniorage.CallOpts, idx)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 idx) view returns(address maker, uint256 stake, int256 amount, uint256 slashingRate, uint256 lockdownExpiration, uint256 number)
func (_Seigniorage *SeigniorageCallerSession) GetProposal(idx *big.Int) (struct {
	Maker              common.Address
	Stake              *big.Int
	Amount             *big.Int
	SlashingRate       *big.Int
	LockdownExpiration *big.Int
	Number             *big.Int
}, error) {
	return _Seigniorage.Contract.GetProposal(&_Seigniorage.CallOpts, idx)
}

// GetProposalCount is a free data retrieval call binding the contract method 0xc08cc02d.
//
// Solidity: function getProposalCount() view returns(uint256)
func (_Seigniorage *SeigniorageCaller) GetProposalCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Seigniorage.contract.Call(opts, out, "getProposalCount")
	return *ret0, err
}

// GetProposalCount is a free data retrieval call binding the contract method 0xc08cc02d.
//
// Solidity: function getProposalCount() view returns(uint256)
func (_Seigniorage *SeigniorageSession) GetProposalCount() (*big.Int, error) {
	return _Seigniorage.Contract.GetProposalCount(&_Seigniorage.CallOpts)
}

// GetProposalCount is a free data retrieval call binding the contract method 0xc08cc02d.
//
// Solidity: function getProposalCount() view returns(uint256)
func (_Seigniorage *SeigniorageCallerSession) GetProposalCount() (*big.Int, error) {
	return _Seigniorage.Contract.GetProposalCount(&_Seigniorage.CallOpts)
}

// GetRemainToAbsorb is a free data retrieval call binding the contract method 0xee1a68c6.
//
// Solidity: function getRemainToAbsorb() view returns(bool, int256)
func (_Seigniorage *SeigniorageCaller) GetRemainToAbsorb(opts *bind.CallOpts) (bool, *big.Int, error) {
	var (
		ret0 = new(bool)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Seigniorage.contract.Call(opts, out, "getRemainToAbsorb")
	return *ret0, *ret1, err
}

// GetRemainToAbsorb is a free data retrieval call binding the contract method 0xee1a68c6.
//
// Solidity: function getRemainToAbsorb() view returns(bool, int256)
func (_Seigniorage *SeigniorageSession) GetRemainToAbsorb() (bool, *big.Int, error) {
	return _Seigniorage.Contract.GetRemainToAbsorb(&_Seigniorage.CallOpts)
}

// GetRemainToAbsorb is a free data retrieval call binding the contract method 0xee1a68c6.
//
// Solidity: function getRemainToAbsorb() view returns(bool, int256)
func (_Seigniorage *SeigniorageCallerSession) GetRemainToAbsorb() (bool, *big.Int, error) {
	return _Seigniorage.Contract.GetRemainToAbsorb(&_Seigniorage.CallOpts)
}

// Next is a free data retrieval call binding the contract method 0x4ea09797.
//
// Solidity: function next(bool orderType, bytes32 id) view returns(bytes32)
func (_Seigniorage *SeigniorageCaller) Next(opts *bind.CallOpts, orderType bool, id [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Seigniorage.contract.Call(opts, out, "next", orderType, id)
	return *ret0, err
}

// Next is a free data retrieval call binding the contract method 0x4ea09797.
//
// Solidity: function next(bool orderType, bytes32 id) view returns(bytes32)
func (_Seigniorage *SeigniorageSession) Next(orderType bool, id [32]byte) ([32]byte, error) {
	return _Seigniorage.Contract.Next(&_Seigniorage.CallOpts, orderType, id)
}

// Next is a free data retrieval call binding the contract method 0x4ea09797.
//
// Solidity: function next(bool orderType, bytes32 id) view returns(bytes32)
func (_Seigniorage *SeigniorageCallerSession) Next(orderType bool, id [32]byte) ([32]byte, error) {
	return _Seigniorage.Contract.Next(&_Seigniorage.CallOpts, orderType, id)
}

// Prev is a free data retrieval call binding the contract method 0x0d90b10a.
//
// Solidity: function prev(bool orderType, bytes32 id) view returns(bytes32)
func (_Seigniorage *SeigniorageCaller) Prev(opts *bind.CallOpts, orderType bool, id [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Seigniorage.contract.Call(opts, out, "prev", orderType, id)
	return *ret0, err
}

// Prev is a free data retrieval call binding the contract method 0x0d90b10a.
//
// Solidity: function prev(bool orderType, bytes32 id) view returns(bytes32)
func (_Seigniorage *SeigniorageSession) Prev(orderType bool, id [32]byte) ([32]byte, error) {
	return _Seigniorage.Contract.Prev(&_Seigniorage.CallOpts, orderType, id)
}

// Prev is a free data retrieval call binding the contract method 0x0d90b10a.
//
// Solidity: function prev(bool orderType, bytes32 id) view returns(bytes32)
func (_Seigniorage *SeigniorageCallerSession) Prev(orderType bool, id [32]byte) ([32]byte, error) {
	return _Seigniorage.Contract.Prev(&_Seigniorage.CallOpts, orderType, id)
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool orderType) view returns(bytes32)
func (_Seigniorage *SeigniorageCaller) Top(opts *bind.CallOpts, orderType bool) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Seigniorage.contract.Call(opts, out, "top", orderType)
	return *ret0, err
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool orderType) view returns(bytes32)
func (_Seigniorage *SeigniorageSession) Top(orderType bool) ([32]byte, error) {
	return _Seigniorage.Contract.Top(&_Seigniorage.CallOpts, orderType)
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool orderType) view returns(bytes32)
func (_Seigniorage *SeigniorageCallerSession) Top(orderType bool) ([32]byte, error) {
	return _Seigniorage.Contract.Top(&_Seigniorage.CallOpts, orderType)
}

// TotalVote is a free data retrieval call binding the contract method 0x4def5645.
//
// Solidity: function totalVote(address maker) view returns(int256)
func (_Seigniorage *SeigniorageCaller) TotalVote(opts *bind.CallOpts, maker common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Seigniorage.contract.Call(opts, out, "totalVote", maker)
	return *ret0, err
}

// TotalVote is a free data retrieval call binding the contract method 0x4def5645.
//
// Solidity: function totalVote(address maker) view returns(int256)
func (_Seigniorage *SeigniorageSession) TotalVote(maker common.Address) (*big.Int, error) {
	return _Seigniorage.Contract.TotalVote(&_Seigniorage.CallOpts, maker)
}

// TotalVote is a free data retrieval call binding the contract method 0x4def5645.
//
// Solidity: function totalVote(address maker) view returns(int256)
func (_Seigniorage *SeigniorageCallerSession) TotalVote(maker common.Address) (*big.Int, error) {
	return _Seigniorage.Contract.TotalVote(&_Seigniorage.CallOpts, maker)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool orderType, bytes32 id) returns()
func (_Seigniorage *SeigniorageTransactor) Cancel(opts *bind.TransactOpts, orderType bool, id [32]byte) (*types.Transaction, error) {
	return _Seigniorage.contract.Transact(opts, "cancel", orderType, id)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool orderType, bytes32 id) returns()
func (_Seigniorage *SeigniorageSession) Cancel(orderType bool, id [32]byte) (*types.Transaction, error) {
	return _Seigniorage.Contract.Cancel(&_Seigniorage.TransactOpts, orderType, id)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool orderType, bytes32 id) returns()
func (_Seigniorage *SeigniorageTransactorSession) Cancel(orderType bool, id [32]byte) (*types.Transaction, error) {
	return _Seigniorage.Contract.Cancel(&_Seigniorage.TransactOpts, orderType, id)
}

// OnBlockInitialized is a paid mutator transaction binding the contract method 0xbe91d729.
//
// Solidity: function onBlockInitialized(uint256 target) returns()
func (_Seigniorage *SeigniorageTransactor) OnBlockInitialized(opts *bind.TransactOpts, target *big.Int) (*types.Transaction, error) {
	return _Seigniorage.contract.Transact(opts, "onBlockInitialized", target)
}

// OnBlockInitialized is a paid mutator transaction binding the contract method 0xbe91d729.
//
// Solidity: function onBlockInitialized(uint256 target) returns()
func (_Seigniorage *SeigniorageSession) OnBlockInitialized(target *big.Int) (*types.Transaction, error) {
	return _Seigniorage.Contract.OnBlockInitialized(&_Seigniorage.TransactOpts, target)
}

// OnBlockInitialized is a paid mutator transaction binding the contract method 0xbe91d729.
//
// Solidity: function onBlockInitialized(uint256 target) returns()
func (_Seigniorage *SeigniorageTransactorSession) OnBlockInitialized(target *big.Int) (*types.Transaction, error) {
	return _Seigniorage.Contract.OnBlockInitialized(&_Seigniorage.TransactOpts, target)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xaa1c259c.
//
// Solidity: function registerTokens(address volatileToken, address stablizeToken) returns()
func (_Seigniorage *SeigniorageTransactor) RegisterTokens(opts *bind.TransactOpts, volatileToken common.Address, stablizeToken common.Address) (*types.Transaction, error) {
	return _Seigniorage.contract.Transact(opts, "registerTokens", volatileToken, stablizeToken)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xaa1c259c.
//
// Solidity: function registerTokens(address volatileToken, address stablizeToken) returns()
func (_Seigniorage *SeigniorageSession) RegisterTokens(volatileToken common.Address, stablizeToken common.Address) (*types.Transaction, error) {
	return _Seigniorage.Contract.RegisterTokens(&_Seigniorage.TransactOpts, volatileToken, stablizeToken)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xaa1c259c.
//
// Solidity: function registerTokens(address volatileToken, address stablizeToken) returns()
func (_Seigniorage *SeigniorageTransactorSession) RegisterTokens(volatileToken common.Address, stablizeToken common.Address) (*types.Transaction, error) {
	return _Seigniorage.Contract.RegisterTokens(&_Seigniorage.TransactOpts, volatileToken, stablizeToken)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address maker) returns()
func (_Seigniorage *SeigniorageTransactor) Revoke(opts *bind.TransactOpts, maker common.Address) (*types.Transaction, error) {
	return _Seigniorage.contract.Transact(opts, "revoke", maker)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address maker) returns()
func (_Seigniorage *SeigniorageSession) Revoke(maker common.Address) (*types.Transaction, error) {
	return _Seigniorage.Contract.Revoke(&_Seigniorage.TransactOpts, maker)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address maker) returns()
func (_Seigniorage *SeigniorageTransactorSession) Revoke(maker common.Address) (*types.Transaction, error) {
	return _Seigniorage.Contract.Revoke(&_Seigniorage.TransactOpts, maker)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address maker, uint256 value, bytes data) returns()
func (_Seigniorage *SeigniorageTransactor) TokenFallback(opts *bind.TransactOpts, maker common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Seigniorage.contract.Transact(opts, "tokenFallback", maker, value, data)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address maker, uint256 value, bytes data) returns()
func (_Seigniorage *SeigniorageSession) TokenFallback(maker common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Seigniorage.Contract.TokenFallback(&_Seigniorage.TransactOpts, maker, value, data)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address maker, uint256 value, bytes data) returns()
func (_Seigniorage *SeigniorageTransactorSession) TokenFallback(maker common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Seigniorage.Contract.TokenFallback(&_Seigniorage.TransactOpts, maker, value, data)
}

// Vote is a paid mutator transaction binding the contract method 0xbd041c4d.
//
// Solidity: function vote(address maker, bool up) returns()
func (_Seigniorage *SeigniorageTransactor) Vote(opts *bind.TransactOpts, maker common.Address, up bool) (*types.Transaction, error) {
	return _Seigniorage.contract.Transact(opts, "vote", maker, up)
}

// Vote is a paid mutator transaction binding the contract method 0xbd041c4d.
//
// Solidity: function vote(address maker, bool up) returns()
func (_Seigniorage *SeigniorageSession) Vote(maker common.Address, up bool) (*types.Transaction, error) {
	return _Seigniorage.Contract.Vote(&_Seigniorage.TransactOpts, maker, up)
}

// Vote is a paid mutator transaction binding the contract method 0xbd041c4d.
//
// Solidity: function vote(address maker, bool up) returns()
func (_Seigniorage *SeigniorageTransactorSession) Vote(maker common.Address, up bool) (*types.Transaction, error) {
	return _Seigniorage.Contract.Vote(&_Seigniorage.TransactOpts, maker, up)
}

// SeigniorageAbsorptionIterator is returned from FilterAbsorption and is used to iterate over the raw logs and unpacked data for Absorption events raised by the Seigniorage contract.
type SeigniorageAbsorptionIterator struct {
	Event *SeigniorageAbsorption // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SeigniorageAbsorptionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeigniorageAbsorption)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SeigniorageAbsorption)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SeigniorageAbsorptionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeigniorageAbsorptionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeigniorageAbsorption represents a Absorption event raised by the Seigniorage contract.
type SeigniorageAbsorption struct {
	Amount  *big.Int
	Supply  *big.Int
	Emptive bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAbsorption is a free log retrieval operation binding the contract event 0x0427b353dc7214e3d8c7f5039475a8e729f4d62922937381e304cd03becf66d2.
//
// Solidity: event Absorption(int256 amount, uint256 supply, bool emptive)
func (_Seigniorage *SeigniorageFilterer) FilterAbsorption(opts *bind.FilterOpts) (*SeigniorageAbsorptionIterator, error) {

	logs, sub, err := _Seigniorage.contract.FilterLogs(opts, "Absorption")
	if err != nil {
		return nil, err
	}
	return &SeigniorageAbsorptionIterator{contract: _Seigniorage.contract, event: "Absorption", logs: logs, sub: sub}, nil
}

// WatchAbsorption is a free log subscription operation binding the contract event 0x0427b353dc7214e3d8c7f5039475a8e729f4d62922937381e304cd03becf66d2.
//
// Solidity: event Absorption(int256 amount, uint256 supply, bool emptive)
func (_Seigniorage *SeigniorageFilterer) WatchAbsorption(opts *bind.WatchOpts, sink chan<- *SeigniorageAbsorption) (event.Subscription, error) {

	logs, sub, err := _Seigniorage.contract.WatchLogs(opts, "Absorption")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeigniorageAbsorption)
				if err := _Seigniorage.contract.UnpackLog(event, "Absorption", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAbsorption is a log parse operation binding the contract event 0x0427b353dc7214e3d8c7f5039475a8e729f4d62922937381e304cd03becf66d2.
//
// Solidity: event Absorption(int256 amount, uint256 supply, bool emptive)
func (_Seigniorage *SeigniorageFilterer) ParseAbsorption(log types.Log) (*SeigniorageAbsorption, error) {
	event := new(SeigniorageAbsorption)
	if err := _Seigniorage.contract.UnpackLog(event, "Absorption", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SeignioragePreemptiveIterator is returned from FilterPreemptive and is used to iterate over the raw logs and unpacked data for Preemptive events raised by the Seigniorage contract.
type SeignioragePreemptiveIterator struct {
	Event *SeignioragePreemptive // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SeignioragePreemptiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeignioragePreemptive)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SeignioragePreemptive)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SeignioragePreemptiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeignioragePreemptiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeignioragePreemptive represents a Preemptive event raised by the Seigniorage contract.
type SeignioragePreemptive struct {
	Maker              common.Address
	Stake              *big.Int
	LockdownExpiration *big.Int
	UnlockNumber       *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPreemptive is a free log retrieval operation binding the contract event 0x8427e4488966b7bd3193a4617993e5e6b9186f0c4b2c303cc6178f4e33b77d08.
//
// Solidity: event Preemptive(address indexed maker, uint256 stake, uint256 lockdownExpiration, uint256 unlockNumber)
func (_Seigniorage *SeigniorageFilterer) FilterPreemptive(opts *bind.FilterOpts, maker []common.Address) (*SeignioragePreemptiveIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Seigniorage.contract.FilterLogs(opts, "Preemptive", makerRule)
	if err != nil {
		return nil, err
	}
	return &SeignioragePreemptiveIterator{contract: _Seigniorage.contract, event: "Preemptive", logs: logs, sub: sub}, nil
}

// WatchPreemptive is a free log subscription operation binding the contract event 0x8427e4488966b7bd3193a4617993e5e6b9186f0c4b2c303cc6178f4e33b77d08.
//
// Solidity: event Preemptive(address indexed maker, uint256 stake, uint256 lockdownExpiration, uint256 unlockNumber)
func (_Seigniorage *SeigniorageFilterer) WatchPreemptive(opts *bind.WatchOpts, sink chan<- *SeignioragePreemptive, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Seigniorage.contract.WatchLogs(opts, "Preemptive", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeignioragePreemptive)
				if err := _Seigniorage.contract.UnpackLog(event, "Preemptive", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePreemptive is a log parse operation binding the contract event 0x8427e4488966b7bd3193a4617993e5e6b9186f0c4b2c303cc6178f4e33b77d08.
//
// Solidity: event Preemptive(address indexed maker, uint256 stake, uint256 lockdownExpiration, uint256 unlockNumber)
func (_Seigniorage *SeigniorageFilterer) ParsePreemptive(log types.Log) (*SeignioragePreemptive, error) {
	event := new(SeignioragePreemptive)
	if err := _Seigniorage.contract.UnpackLog(event, "Preemptive", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SeigniorageProposeIterator is returned from FilterPropose and is used to iterate over the raw logs and unpacked data for Propose events raised by the Seigniorage contract.
type SeigniorageProposeIterator struct {
	Event *SeignioragePropose // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SeigniorageProposeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeignioragePropose)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SeignioragePropose)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SeigniorageProposeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeigniorageProposeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeignioragePropose represents a Propose event raised by the Seigniorage contract.
type SeignioragePropose struct {
	Maker              common.Address
	Amount             *big.Int
	Stake              *big.Int
	LockdownExpiration *big.Int
	SlashingRate       *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPropose is a free log retrieval operation binding the contract event 0x56e25d1b63c01627fcd54936462c97aeb9a18352bf0ed161e8141a33cfd795ca.
//
// Solidity: event Propose(address indexed maker, int256 amount, uint256 stake, uint256 lockdownExpiration, uint256 slashingRate)
func (_Seigniorage *SeigniorageFilterer) FilterPropose(opts *bind.FilterOpts, maker []common.Address) (*SeigniorageProposeIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Seigniorage.contract.FilterLogs(opts, "Propose", makerRule)
	if err != nil {
		return nil, err
	}
	return &SeigniorageProposeIterator{contract: _Seigniorage.contract, event: "Propose", logs: logs, sub: sub}, nil
}

// WatchPropose is a free log subscription operation binding the contract event 0x56e25d1b63c01627fcd54936462c97aeb9a18352bf0ed161e8141a33cfd795ca.
//
// Solidity: event Propose(address indexed maker, int256 amount, uint256 stake, uint256 lockdownExpiration, uint256 slashingRate)
func (_Seigniorage *SeigniorageFilterer) WatchPropose(opts *bind.WatchOpts, sink chan<- *SeignioragePropose, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Seigniorage.contract.WatchLogs(opts, "Propose", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeignioragePropose)
				if err := _Seigniorage.contract.UnpackLog(event, "Propose", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePropose is a log parse operation binding the contract event 0x56e25d1b63c01627fcd54936462c97aeb9a18352bf0ed161e8141a33cfd795ca.
//
// Solidity: event Propose(address indexed maker, int256 amount, uint256 stake, uint256 lockdownExpiration, uint256 slashingRate)
func (_Seigniorage *SeigniorageFilterer) ParsePropose(log types.Log) (*SeignioragePropose, error) {
	event := new(SeignioragePropose)
	if err := _Seigniorage.contract.UnpackLog(event, "Propose", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SeigniorageRevokeIterator is returned from FilterRevoke and is used to iterate over the raw logs and unpacked data for Revoke events raised by the Seigniorage contract.
type SeigniorageRevokeIterator struct {
	Event *SeigniorageRevoke // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SeigniorageRevokeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeigniorageRevoke)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SeigniorageRevoke)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SeigniorageRevokeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeigniorageRevokeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeigniorageRevoke represents a Revoke event raised by the Seigniorage contract.
type SeigniorageRevoke struct {
	Maker common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRevoke is a free log retrieval operation binding the contract event 0x9f77920c3de8baaa98d273e8aa75fae382aaa9f7f60f38979137853e5b73ea2c.
//
// Solidity: event Revoke(address indexed maker)
func (_Seigniorage *SeigniorageFilterer) FilterRevoke(opts *bind.FilterOpts, maker []common.Address) (*SeigniorageRevokeIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Seigniorage.contract.FilterLogs(opts, "Revoke", makerRule)
	if err != nil {
		return nil, err
	}
	return &SeigniorageRevokeIterator{contract: _Seigniorage.contract, event: "Revoke", logs: logs, sub: sub}, nil
}

// WatchRevoke is a free log subscription operation binding the contract event 0x9f77920c3de8baaa98d273e8aa75fae382aaa9f7f60f38979137853e5b73ea2c.
//
// Solidity: event Revoke(address indexed maker)
func (_Seigniorage *SeigniorageFilterer) WatchRevoke(opts *bind.WatchOpts, sink chan<- *SeigniorageRevoke, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Seigniorage.contract.WatchLogs(opts, "Revoke", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeigniorageRevoke)
				if err := _Seigniorage.contract.UnpackLog(event, "Revoke", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRevoke is a log parse operation binding the contract event 0x9f77920c3de8baaa98d273e8aa75fae382aaa9f7f60f38979137853e5b73ea2c.
//
// Solidity: event Revoke(address indexed maker)
func (_Seigniorage *SeigniorageFilterer) ParseRevoke(log types.Log) (*SeigniorageRevoke, error) {
	event := new(SeigniorageRevoke)
	if err := _Seigniorage.contract.UnpackLog(event, "Revoke", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SeigniorageSlashIterator is returned from FilterSlash and is used to iterate over the raw logs and unpacked data for Slash events raised by the Seigniorage contract.
type SeigniorageSlashIterator struct {
	Event *SeigniorageSlash // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SeigniorageSlashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeigniorageSlash)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SeigniorageSlash)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SeigniorageSlashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeigniorageSlashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeigniorageSlash represents a Slash event raised by the Seigniorage contract.
type SeigniorageSlash struct {
	Maker  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSlash is a free log retrieval operation binding the contract event 0xa69f22d963cb7981f842db8c1aafcc93d915ba2a95dcf26dcc333a9c2a09be26.
//
// Solidity: event Slash(address indexed maker, uint256 amount)
func (_Seigniorage *SeigniorageFilterer) FilterSlash(opts *bind.FilterOpts, maker []common.Address) (*SeigniorageSlashIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Seigniorage.contract.FilterLogs(opts, "Slash", makerRule)
	if err != nil {
		return nil, err
	}
	return &SeigniorageSlashIterator{contract: _Seigniorage.contract, event: "Slash", logs: logs, sub: sub}, nil
}

// WatchSlash is a free log subscription operation binding the contract event 0xa69f22d963cb7981f842db8c1aafcc93d915ba2a95dcf26dcc333a9c2a09be26.
//
// Solidity: event Slash(address indexed maker, uint256 amount)
func (_Seigniorage *SeigniorageFilterer) WatchSlash(opts *bind.WatchOpts, sink chan<- *SeigniorageSlash, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Seigniorage.contract.WatchLogs(opts, "Slash", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeigniorageSlash)
				if err := _Seigniorage.contract.UnpackLog(event, "Slash", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSlash is a log parse operation binding the contract event 0xa69f22d963cb7981f842db8c1aafcc93d915ba2a95dcf26dcc333a9c2a09be26.
//
// Solidity: event Slash(address indexed maker, uint256 amount)
func (_Seigniorage *SeigniorageFilterer) ParseSlash(log types.Log) (*SeigniorageSlash, error) {
	event := new(SeigniorageSlash)
	if err := _Seigniorage.contract.UnpackLog(event, "Slash", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SeigniorageStopIterator is returned from FilterStop and is used to iterate over the raw logs and unpacked data for Stop events raised by the Seigniorage contract.
type SeigniorageStopIterator struct {
	Event *SeigniorageStop // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SeigniorageStopIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeigniorageStop)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SeigniorageStop)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SeigniorageStopIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeigniorageStopIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeigniorageStop represents a Stop event raised by the Seigniorage contract.
type SeigniorageStop struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStop is a free log retrieval operation binding the contract event 0xbedf0f4abfe86d4ffad593d9607fe70e83ea706033d44d24b3b6283cf3fc4f6b.
//
// Solidity: event Stop()
func (_Seigniorage *SeigniorageFilterer) FilterStop(opts *bind.FilterOpts) (*SeigniorageStopIterator, error) {

	logs, sub, err := _Seigniorage.contract.FilterLogs(opts, "Stop")
	if err != nil {
		return nil, err
	}
	return &SeigniorageStopIterator{contract: _Seigniorage.contract, event: "Stop", logs: logs, sub: sub}, nil
}

// WatchStop is a free log subscription operation binding the contract event 0xbedf0f4abfe86d4ffad593d9607fe70e83ea706033d44d24b3b6283cf3fc4f6b.
//
// Solidity: event Stop()
func (_Seigniorage *SeigniorageFilterer) WatchStop(opts *bind.WatchOpts, sink chan<- *SeigniorageStop) (event.Subscription, error) {

	logs, sub, err := _Seigniorage.contract.WatchLogs(opts, "Stop")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeigniorageStop)
				if err := _Seigniorage.contract.UnpackLog(event, "Stop", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStop is a log parse operation binding the contract event 0xbedf0f4abfe86d4ffad593d9607fe70e83ea706033d44d24b3b6283cf3fc4f6b.
//
// Solidity: event Stop()
func (_Seigniorage *SeigniorageFilterer) ParseStop(log types.Log) (*SeigniorageStop, error) {
	event := new(SeigniorageStop)
	if err := _Seigniorage.contract.UnpackLog(event, "Stop", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SeigniorageUnlockIterator is returned from FilterUnlock and is used to iterate over the raw logs and unpacked data for Unlock events raised by the Seigniorage contract.
type SeigniorageUnlockIterator struct {
	Event *SeigniorageUnlock // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SeigniorageUnlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeigniorageUnlock)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SeigniorageUnlock)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SeigniorageUnlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeigniorageUnlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeigniorageUnlock represents a Unlock event raised by the Seigniorage contract.
type SeigniorageUnlock struct {
	Maker common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnlock is a free log retrieval operation binding the contract event 0x0be774851955c26a1d6a32b13b020663a069006b4a3b643ff0b809d318260572.
//
// Solidity: event Unlock(address indexed maker)
func (_Seigniorage *SeigniorageFilterer) FilterUnlock(opts *bind.FilterOpts, maker []common.Address) (*SeigniorageUnlockIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Seigniorage.contract.FilterLogs(opts, "Unlock", makerRule)
	if err != nil {
		return nil, err
	}
	return &SeigniorageUnlockIterator{contract: _Seigniorage.contract, event: "Unlock", logs: logs, sub: sub}, nil
}

// WatchUnlock is a free log subscription operation binding the contract event 0x0be774851955c26a1d6a32b13b020663a069006b4a3b643ff0b809d318260572.
//
// Solidity: event Unlock(address indexed maker)
func (_Seigniorage *SeigniorageFilterer) WatchUnlock(opts *bind.WatchOpts, sink chan<- *SeigniorageUnlock, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Seigniorage.contract.WatchLogs(opts, "Unlock", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeigniorageUnlock)
				if err := _Seigniorage.contract.UnpackLog(event, "Unlock", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnlock is a log parse operation binding the contract event 0x0be774851955c26a1d6a32b13b020663a069006b4a3b643ff0b809d318260572.
//
// Solidity: event Unlock(address indexed maker)
func (_Seigniorage *SeigniorageFilterer) ParseUnlock(log types.Log) (*SeigniorageUnlock, error) {
	event := new(SeigniorageUnlock)
	if err := _Seigniorage.contract.UnpackLog(event, "Unlock", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AbsnABI is the input ABI used to generate the binding from.
const AbsnABI = "[]"

// AbsnBin is the compiled bytecode used for deploying new contracts.
var AbsnBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220199fe36af36bc62fe7710fcc9151fe3739efe1881b960511232fee57b7312e8864736f6c63430006080033"

// DeployAbsn deploys a new Ethereum contract, binding an instance of Absn to it.
func DeployAbsn(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Absn, error) {
	parsed, err := abi.JSON(strings.NewReader(AbsnABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AbsnBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Absn{AbsnCaller: AbsnCaller{contract: contract}, AbsnTransactor: AbsnTransactor{contract: contract}, AbsnFilterer: AbsnFilterer{contract: contract}}, nil
}

// Absn is an auto generated Go binding around an Ethereum contract.
type Absn struct {
	AbsnCaller     // Read-only binding to the contract
	AbsnTransactor // Write-only binding to the contract
	AbsnFilterer   // Log filterer for contract events
}

// AbsnCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbsnCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbsnTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbsnTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbsnFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbsnFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbsnSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbsnSession struct {
	Contract     *Absn             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbsnCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbsnCallerSession struct {
	Contract *AbsnCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AbsnTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbsnTransactorSession struct {
	Contract     *AbsnTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbsnRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbsnRaw struct {
	Contract *Absn // Generic contract binding to access the raw methods on
}

// AbsnCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbsnCallerRaw struct {
	Contract *AbsnCaller // Generic read-only contract binding to access the raw methods on
}

// AbsnTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbsnTransactorRaw struct {
	Contract *AbsnTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbsn creates a new instance of Absn, bound to a specific deployed contract.
func NewAbsn(address common.Address, backend bind.ContractBackend) (*Absn, error) {
	contract, err := bindAbsn(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Absn{AbsnCaller: AbsnCaller{contract: contract}, AbsnTransactor: AbsnTransactor{contract: contract}, AbsnFilterer: AbsnFilterer{contract: contract}}, nil
}

// NewAbsnCaller creates a new read-only instance of Absn, bound to a specific deployed contract.
func NewAbsnCaller(address common.Address, caller bind.ContractCaller) (*AbsnCaller, error) {
	contract, err := bindAbsn(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbsnCaller{contract: contract}, nil
}

// NewAbsnTransactor creates a new write-only instance of Absn, bound to a specific deployed contract.
func NewAbsnTransactor(address common.Address, transactor bind.ContractTransactor) (*AbsnTransactor, error) {
	contract, err := bindAbsn(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbsnTransactor{contract: contract}, nil
}

// NewAbsnFilterer creates a new log filterer instance of Absn, bound to a specific deployed contract.
func NewAbsnFilterer(address common.Address, filterer bind.ContractFilterer) (*AbsnFilterer, error) {
	contract, err := bindAbsn(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbsnFilterer{contract: contract}, nil
}

// bindAbsn binds a generic wrapper to an already deployed contract.
func bindAbsn(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AbsnABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Absn *AbsnRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Absn.Contract.AbsnCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Absn *AbsnRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Absn.Contract.AbsnTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Absn *AbsnRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Absn.Contract.AbsnTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Absn *AbsnCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Absn.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Absn *AbsnTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Absn.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Absn *AbsnTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Absn.Contract.contract.Transact(opts, method, params...)
}

// DexABI is the input ABI used to generate the binding from.
const DexABI = "[]"

// DexBin is the compiled bytecode used for deploying new contracts.
var DexBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122007a6b7284a81e04215184c8f4fd4b576a1b770bd7e39ace27421165be8d2112064736f6c63430006080033"

// DeployDex deploys a new Ethereum contract, binding an instance of Dex to it.
func DeployDex(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Dex, error) {
	parsed, err := abi.JSON(strings.NewReader(DexABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DexBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Dex{DexCaller: DexCaller{contract: contract}, DexTransactor: DexTransactor{contract: contract}, DexFilterer: DexFilterer{contract: contract}}, nil
}

// Dex is an auto generated Go binding around an Ethereum contract.
type Dex struct {
	DexCaller     // Read-only binding to the contract
	DexTransactor // Write-only binding to the contract
	DexFilterer   // Log filterer for contract events
}

// DexCaller is an auto generated read-only Go binding around an Ethereum contract.
type DexCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DexTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DexTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DexFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DexFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DexSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DexSession struct {
	Contract     *Dex              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DexCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DexCallerSession struct {
	Contract *DexCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DexTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DexTransactorSession struct {
	Contract     *DexTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DexRaw is an auto generated low-level Go binding around an Ethereum contract.
type DexRaw struct {
	Contract *Dex // Generic contract binding to access the raw methods on
}

// DexCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DexCallerRaw struct {
	Contract *DexCaller // Generic read-only contract binding to access the raw methods on
}

// DexTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DexTransactorRaw struct {
	Contract *DexTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDex creates a new instance of Dex, bound to a specific deployed contract.
func NewDex(address common.Address, backend bind.ContractBackend) (*Dex, error) {
	contract, err := bindDex(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dex{DexCaller: DexCaller{contract: contract}, DexTransactor: DexTransactor{contract: contract}, DexFilterer: DexFilterer{contract: contract}}, nil
}

// NewDexCaller creates a new read-only instance of Dex, bound to a specific deployed contract.
func NewDexCaller(address common.Address, caller bind.ContractCaller) (*DexCaller, error) {
	contract, err := bindDex(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DexCaller{contract: contract}, nil
}

// NewDexTransactor creates a new write-only instance of Dex, bound to a specific deployed contract.
func NewDexTransactor(address common.Address, transactor bind.ContractTransactor) (*DexTransactor, error) {
	contract, err := bindDex(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DexTransactor{contract: contract}, nil
}

// NewDexFilterer creates a new log filterer instance of Dex, bound to a specific deployed contract.
func NewDexFilterer(address common.Address, filterer bind.ContractFilterer) (*DexFilterer, error) {
	contract, err := bindDex(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DexFilterer{contract: contract}, nil
}

// bindDex binds a generic wrapper to an already deployed contract.
func bindDex(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DexABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dex *DexRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dex.Contract.DexCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dex *DexRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dex.Contract.DexTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dex *DexRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dex.Contract.DexTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dex *DexCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dex.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dex *DexTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dex.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dex *DexTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dex.Contract.contract.Transact(opts, method, params...)
}

// MapABI is the input ABI used to generate the binding from.
const MapABI = "[]"

// MapBin is the compiled bytecode used for deploying new contracts.
var MapBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122090a4d42a8f58babfe7ad99341a836756a489d4f397b48002ff2b0c31e7540e7c64736f6c63430006080033"

// DeployMap deploys a new Ethereum contract, binding an instance of Map to it.
func DeployMap(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Map, error) {
	parsed, err := abi.JSON(strings.NewReader(MapABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MapBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Map{MapCaller: MapCaller{contract: contract}, MapTransactor: MapTransactor{contract: contract}, MapFilterer: MapFilterer{contract: contract}}, nil
}

// Map is an auto generated Go binding around an Ethereum contract.
type Map struct {
	MapCaller     // Read-only binding to the contract
	MapTransactor // Write-only binding to the contract
	MapFilterer   // Log filterer for contract events
}

// MapCaller is an auto generated read-only Go binding around an Ethereum contract.
type MapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MapSession struct {
	Contract     *Map              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MapCallerSession struct {
	Contract *MapCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MapTransactorSession struct {
	Contract     *MapTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MapRaw is an auto generated low-level Go binding around an Ethereum contract.
type MapRaw struct {
	Contract *Map // Generic contract binding to access the raw methods on
}

// MapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MapCallerRaw struct {
	Contract *MapCaller // Generic read-only contract binding to access the raw methods on
}

// MapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MapTransactorRaw struct {
	Contract *MapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMap creates a new instance of Map, bound to a specific deployed contract.
func NewMap(address common.Address, backend bind.ContractBackend) (*Map, error) {
	contract, err := bindMap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Map{MapCaller: MapCaller{contract: contract}, MapTransactor: MapTransactor{contract: contract}, MapFilterer: MapFilterer{contract: contract}}, nil
}

// NewMapCaller creates a new read-only instance of Map, bound to a specific deployed contract.
func NewMapCaller(address common.Address, caller bind.ContractCaller) (*MapCaller, error) {
	contract, err := bindMap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MapCaller{contract: contract}, nil
}

// NewMapTransactor creates a new write-only instance of Map, bound to a specific deployed contract.
func NewMapTransactor(address common.Address, transactor bind.ContractTransactor) (*MapTransactor, error) {
	contract, err := bindMap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MapTransactor{contract: contract}, nil
}

// NewMapFilterer creates a new log filterer instance of Map, bound to a specific deployed contract.
func NewMapFilterer(address common.Address, filterer bind.ContractFilterer) (*MapFilterer, error) {
	contract, err := bindMap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MapFilterer{contract: contract}, nil
}

// bindMap binds a generic wrapper to an already deployed contract.
func bindMap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Map *MapRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Map.Contract.MapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Map *MapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Map.Contract.MapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Map *MapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Map.Contract.MapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Map *MapCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Map.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Map *MapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Map.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Map *MapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Map.Contract.contract.Transact(opts, method, params...)
}

// UtilABI is the input ABI used to generate the binding from.
const UtilABI = "[]"

// UtilBin is the compiled bytecode used for deploying new contracts.
var UtilBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220108c578570969fe4e2bc0f2974f856ea5e0cbaea5c9b0a895c68635a8d69693364736f6c63430006080033"

// DeployUtil deploys a new Ethereum contract, binding an instance of Util to it.
func DeployUtil(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Util, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UtilBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Util{UtilCaller: UtilCaller{contract: contract}, UtilTransactor: UtilTransactor{contract: contract}, UtilFilterer: UtilFilterer{contract: contract}}, nil
}

// Util is an auto generated Go binding around an Ethereum contract.
type Util struct {
	UtilCaller     // Read-only binding to the contract
	UtilTransactor // Write-only binding to the contract
	UtilFilterer   // Log filterer for contract events
}

// UtilCaller is an auto generated read-only Go binding around an Ethereum contract.
type UtilCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UtilTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UtilFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UtilSession struct {
	Contract     *Util             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UtilCallerSession struct {
	Contract *UtilCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UtilTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UtilTransactorSession struct {
	Contract     *UtilTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilRaw is an auto generated low-level Go binding around an Ethereum contract.
type UtilRaw struct {
	Contract *Util // Generic contract binding to access the raw methods on
}

// UtilCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UtilCallerRaw struct {
	Contract *UtilCaller // Generic read-only contract binding to access the raw methods on
}

// UtilTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UtilTransactorRaw struct {
	Contract *UtilTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUtil creates a new instance of Util, bound to a specific deployed contract.
func NewUtil(address common.Address, backend bind.ContractBackend) (*Util, error) {
	contract, err := bindUtil(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Util{UtilCaller: UtilCaller{contract: contract}, UtilTransactor: UtilTransactor{contract: contract}, UtilFilterer: UtilFilterer{contract: contract}}, nil
}

// NewUtilCaller creates a new read-only instance of Util, bound to a specific deployed contract.
func NewUtilCaller(address common.Address, caller bind.ContractCaller) (*UtilCaller, error) {
	contract, err := bindUtil(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UtilCaller{contract: contract}, nil
}

// NewUtilTransactor creates a new write-only instance of Util, bound to a specific deployed contract.
func NewUtilTransactor(address common.Address, transactor bind.ContractTransactor) (*UtilTransactor, error) {
	contract, err := bindUtil(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UtilTransactor{contract: contract}, nil
}

// NewUtilFilterer creates a new log filterer instance of Util, bound to a specific deployed contract.
func NewUtilFilterer(address common.Address, filterer bind.ContractFilterer) (*UtilFilterer, error) {
	contract, err := bindUtil(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UtilFilterer{contract: contract}, nil
}

// bindUtil binds a generic wrapper to an already deployed contract.
func bindUtil(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Util *UtilRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Util.Contract.UtilCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Util *UtilRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Util.Contract.UtilTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Util *UtilRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Util.Contract.UtilTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Util *UtilCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Util.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Util *UtilTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Util.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Util *UtilTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Util.Contract.contract.Transact(opts, method, params...)
}
