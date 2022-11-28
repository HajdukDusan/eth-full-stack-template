// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package StupidContract

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StupidContractMetaData contains all meta data concerning the StupidContract contract.
var StupidContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_entryFee\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"StupidContract__FeeNotEnough\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StupidContract__NotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StupidContract__TransferFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"StupidEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_message\",\"type\":\"string\"}],\"name\":\"AddToRegistry\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MessageEntryFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PullStupidFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"StupidContractDescription\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stupidRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b5060405162000a8238038062000a8283398101604081905262000034916200006a565b3360805260a082905260016200004b8282620001d9565b505050620002a5565b634e487b7160e01b600052604160045260246000fd5b600080604083850312156200007e57600080fd5b8251602080850151919350906001600160401b0380821115620000a057600080fd5b818601915086601f830112620000b557600080fd5b815181811115620000ca57620000ca62000054565b604051601f8201601f19908116603f01168101908382118183101715620000f557620000f562000054565b8160405282815289868487010111156200010e57600080fd5b600093505b8284101562000132578484018601518185018701529285019262000113565b60008684830101528096505050505050509250929050565b600181811c908216806200015f57607f821691505b6020821081036200018057634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115620001d457600081815260208120601f850160051c81016020861015620001af5750805b601f850160051c820191505b81811015620001d057828155600101620001bb565b5050505b505050565b81516001600160401b03811115620001f557620001f562000054565b6200020d816200020684546200014a565b8462000186565b602080601f8311600181146200024557600084156200022c5750858301515b600019600386901b1c1916600185901b178555620001d0565b600085815260208120601f198616915b82811015620002765788860151825594840194600190910190840162000255565b5085821015620002955787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60805160a0516107a3620002df6000396000818160a301526102be015260008181610107015281816103fb015261043d01526107a36000f3fe6080604052600436106100555760003560e01c8063045917f31461005a578063692c4ae2146100915780636d34aafc146100d35780638da5cb5b146100f5578063c6a06fc314610141578063d1fb08c514610156575b600080fd5b34801561006657600080fd5b5061007a6100753660046104d0565b61016b565b60405161008892919061052f565b60405180910390f35b34801561009d57600080fd5b506100c57f000000000000000000000000000000000000000000000000000000000000000081565b604051908152602001610088565b3480156100df57600080fd5b506100e861022d565b604051610088919061055b565b34801561010157600080fd5b506101297f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610088565b61015461014f366004610575565b6102bb565b005b34801561016257600080fd5b506101546103f0565b6000818154811061017b57600080fd5b6000918252602090912060029091020180546001820180546001600160a01b039092169350906101aa906105e7565b80601f01602080910402602001604051908101604052809291908181526020018280546101d6906105e7565b80156102235780601f106101f857610100808354040283529160200191610223565b820191906000526020600020905b81548152906001019060200180831161020657829003601f168201915b5050505050905082565b6001805461023a906105e7565b80601f0160208091040260200160405190810160405280929190818152602001828054610266906105e7565b80156102b35780601f10610288576101008083540402835291602001916102b3565b820191906000526020600020905b81548152906001019060200180831161029657829003601f168201915b505050505081565b347f00000000000000000000000000000000000000000000000000000000000000008110156102fd57604051638daa29d960e01b815260040160405180910390fd5b60006040518060400160405280336001600160a01b0316815260200185858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920182905250939094525050835460018082018655948252602091829020845160029092020180546001600160a01b0319166001600160a01b039092169190911781559083015192939092908301915061039f9082610686565b50506000543391507f7a4b508fdb79d1a00f12fc7c4d08d8b1a843a27f8576d2ed06f0691dc2b76d7e906103d590600190610746565b604080519182524260208301520160405180910390a2505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146104395760405163a7bb3f6f60e01b815260040160405180910390fd5b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03164760405160006040518083038185875af1925050503d80600081146104a6576040519150601f19603f3d011682016040523d82523d6000602084013e6104ab565b606091505b50509050806104cd57604051631d5bb8d360e21b815260040160405180910390fd5b50565b6000602082840312156104e257600080fd5b5035919050565b6000815180845260005b8181101561050f576020818501810151868301820152016104f3565b506000602082860101526020601f19601f83011685010191505092915050565b6001600160a01b0383168152604060208201819052600090610553908301846104e9565b949350505050565b60208152600061056e60208301846104e9565b9392505050565b6000806020838503121561058857600080fd5b823567ffffffffffffffff808211156105a057600080fd5b818501915085601f8301126105b457600080fd5b8135818111156105c357600080fd5b8660208285010111156105d557600080fd5b60209290920196919550909350505050565b600181811c908216806105fb57607f821691505b60208210810361061b57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fd5b601f82111561068157600081815260208120601f850160051c8101602086101561065e5750805b601f850160051c820191505b8181101561067d5782815560010161066a565b5050505b505050565b815167ffffffffffffffff8111156106a0576106a0610621565b6106b4816106ae84546105e7565b84610637565b602080601f8311600181146106e957600084156106d15750858301515b600019600386901b1c1916600185901b17855561067d565b600085815260208120601f198616915b82811015610718578886015182559484019460019091019084016106f9565b50858210156107365787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b8181038181111561076757634e487b7160e01b600052601160045260246000fd5b9291505056fea2646970667358221220eac69d834fbf6a360d331fd506b720b50ed96ac7a305386c14d08a3a5787262c64736f6c63430008110033",
}

// StupidContractABI is the input ABI used to generate the binding from.
// Deprecated: Use StupidContractMetaData.ABI instead.
var StupidContractABI = StupidContractMetaData.ABI

// StupidContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StupidContractMetaData.Bin instead.
var StupidContractBin = StupidContractMetaData.Bin

// DeployStupidContract deploys a new Ethereum contract, binding an instance of StupidContract to it.
func DeployStupidContract(auth *bind.TransactOpts, backend bind.ContractBackend, _entryFee *big.Int, _description string) (common.Address, *types.Transaction, *StupidContract, error) {
	parsed, err := StupidContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StupidContractBin), backend, _entryFee, _description)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StupidContract{StupidContractCaller: StupidContractCaller{contract: contract}, StupidContractTransactor: StupidContractTransactor{contract: contract}, StupidContractFilterer: StupidContractFilterer{contract: contract}}, nil
}

// StupidContract is an auto generated Go binding around an Ethereum contract.
type StupidContract struct {
	StupidContractCaller     // Read-only binding to the contract
	StupidContractTransactor // Write-only binding to the contract
	StupidContractFilterer   // Log filterer for contract events
}

// StupidContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type StupidContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StupidContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StupidContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StupidContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StupidContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StupidContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StupidContractSession struct {
	Contract     *StupidContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StupidContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StupidContractCallerSession struct {
	Contract *StupidContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// StupidContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StupidContractTransactorSession struct {
	Contract     *StupidContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// StupidContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type StupidContractRaw struct {
	Contract *StupidContract // Generic contract binding to access the raw methods on
}

// StupidContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StupidContractCallerRaw struct {
	Contract *StupidContractCaller // Generic read-only contract binding to access the raw methods on
}

// StupidContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StupidContractTransactorRaw struct {
	Contract *StupidContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStupidContract creates a new instance of StupidContract, bound to a specific deployed contract.
func NewStupidContract(address common.Address, backend bind.ContractBackend) (*StupidContract, error) {
	contract, err := bindStupidContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StupidContract{StupidContractCaller: StupidContractCaller{contract: contract}, StupidContractTransactor: StupidContractTransactor{contract: contract}, StupidContractFilterer: StupidContractFilterer{contract: contract}}, nil
}

// NewStupidContractCaller creates a new read-only instance of StupidContract, bound to a specific deployed contract.
func NewStupidContractCaller(address common.Address, caller bind.ContractCaller) (*StupidContractCaller, error) {
	contract, err := bindStupidContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StupidContractCaller{contract: contract}, nil
}

// NewStupidContractTransactor creates a new write-only instance of StupidContract, bound to a specific deployed contract.
func NewStupidContractTransactor(address common.Address, transactor bind.ContractTransactor) (*StupidContractTransactor, error) {
	contract, err := bindStupidContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StupidContractTransactor{contract: contract}, nil
}

// NewStupidContractFilterer creates a new log filterer instance of StupidContract, bound to a specific deployed contract.
func NewStupidContractFilterer(address common.Address, filterer bind.ContractFilterer) (*StupidContractFilterer, error) {
	contract, err := bindStupidContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StupidContractFilterer{contract: contract}, nil
}

// bindStupidContract binds a generic wrapper to an already deployed contract.
func bindStupidContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StupidContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StupidContract *StupidContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StupidContract.Contract.StupidContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StupidContract *StupidContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StupidContract.Contract.StupidContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StupidContract *StupidContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StupidContract.Contract.StupidContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StupidContract *StupidContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StupidContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StupidContract *StupidContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StupidContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StupidContract *StupidContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StupidContract.Contract.contract.Transact(opts, method, params...)
}

// MessageEntryFee is a free data retrieval call binding the contract method 0x692c4ae2.
//
// Solidity: function MessageEntryFee() view returns(uint256)
func (_StupidContract *StupidContractCaller) MessageEntryFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StupidContract.contract.Call(opts, &out, "MessageEntryFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MessageEntryFee is a free data retrieval call binding the contract method 0x692c4ae2.
//
// Solidity: function MessageEntryFee() view returns(uint256)
func (_StupidContract *StupidContractSession) MessageEntryFee() (*big.Int, error) {
	return _StupidContract.Contract.MessageEntryFee(&_StupidContract.CallOpts)
}

// MessageEntryFee is a free data retrieval call binding the contract method 0x692c4ae2.
//
// Solidity: function MessageEntryFee() view returns(uint256)
func (_StupidContract *StupidContractCallerSession) MessageEntryFee() (*big.Int, error) {
	return _StupidContract.Contract.MessageEntryFee(&_StupidContract.CallOpts)
}

// StupidContractDescription is a free data retrieval call binding the contract method 0x6d34aafc.
//
// Solidity: function StupidContractDescription() view returns(string)
func (_StupidContract *StupidContractCaller) StupidContractDescription(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StupidContract.contract.Call(opts, &out, "StupidContractDescription")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// StupidContractDescription is a free data retrieval call binding the contract method 0x6d34aafc.
//
// Solidity: function StupidContractDescription() view returns(string)
func (_StupidContract *StupidContractSession) StupidContractDescription() (string, error) {
	return _StupidContract.Contract.StupidContractDescription(&_StupidContract.CallOpts)
}

// StupidContractDescription is a free data retrieval call binding the contract method 0x6d34aafc.
//
// Solidity: function StupidContractDescription() view returns(string)
func (_StupidContract *StupidContractCallerSession) StupidContractDescription() (string, error) {
	return _StupidContract.Contract.StupidContractDescription(&_StupidContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StupidContract *StupidContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StupidContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StupidContract *StupidContractSession) Owner() (common.Address, error) {
	return _StupidContract.Contract.Owner(&_StupidContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StupidContract *StupidContractCallerSession) Owner() (common.Address, error) {
	return _StupidContract.Contract.Owner(&_StupidContract.CallOpts)
}

// StupidRegistry is a free data retrieval call binding the contract method 0x045917f3.
//
// Solidity: function stupidRegistry(uint256 ) view returns(address sender, string message)
func (_StupidContract *StupidContractCaller) StupidRegistry(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Sender  common.Address
	Message string
}, error) {
	var out []interface{}
	err := _StupidContract.contract.Call(opts, &out, "stupidRegistry", arg0)

	outstruct := new(struct {
		Sender  common.Address
		Message string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Sender = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Message = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// StupidRegistry is a free data retrieval call binding the contract method 0x045917f3.
//
// Solidity: function stupidRegistry(uint256 ) view returns(address sender, string message)
func (_StupidContract *StupidContractSession) StupidRegistry(arg0 *big.Int) (struct {
	Sender  common.Address
	Message string
}, error) {
	return _StupidContract.Contract.StupidRegistry(&_StupidContract.CallOpts, arg0)
}

// StupidRegistry is a free data retrieval call binding the contract method 0x045917f3.
//
// Solidity: function stupidRegistry(uint256 ) view returns(address sender, string message)
func (_StupidContract *StupidContractCallerSession) StupidRegistry(arg0 *big.Int) (struct {
	Sender  common.Address
	Message string
}, error) {
	return _StupidContract.Contract.StupidRegistry(&_StupidContract.CallOpts, arg0)
}

// AddToRegistry is a paid mutator transaction binding the contract method 0xc6a06fc3.
//
// Solidity: function AddToRegistry(string _message) payable returns()
func (_StupidContract *StupidContractTransactor) AddToRegistry(opts *bind.TransactOpts, _message string) (*types.Transaction, error) {
	return _StupidContract.contract.Transact(opts, "AddToRegistry", _message)
}

// AddToRegistry is a paid mutator transaction binding the contract method 0xc6a06fc3.
//
// Solidity: function AddToRegistry(string _message) payable returns()
func (_StupidContract *StupidContractSession) AddToRegistry(_message string) (*types.Transaction, error) {
	return _StupidContract.Contract.AddToRegistry(&_StupidContract.TransactOpts, _message)
}

// AddToRegistry is a paid mutator transaction binding the contract method 0xc6a06fc3.
//
// Solidity: function AddToRegistry(string _message) payable returns()
func (_StupidContract *StupidContractTransactorSession) AddToRegistry(_message string) (*types.Transaction, error) {
	return _StupidContract.Contract.AddToRegistry(&_StupidContract.TransactOpts, _message)
}

// PullStupidFees is a paid mutator transaction binding the contract method 0xd1fb08c5.
//
// Solidity: function PullStupidFees() returns()
func (_StupidContract *StupidContractTransactor) PullStupidFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StupidContract.contract.Transact(opts, "PullStupidFees")
}

// PullStupidFees is a paid mutator transaction binding the contract method 0xd1fb08c5.
//
// Solidity: function PullStupidFees() returns()
func (_StupidContract *StupidContractSession) PullStupidFees() (*types.Transaction, error) {
	return _StupidContract.Contract.PullStupidFees(&_StupidContract.TransactOpts)
}

// PullStupidFees is a paid mutator transaction binding the contract method 0xd1fb08c5.
//
// Solidity: function PullStupidFees() returns()
func (_StupidContract *StupidContractTransactorSession) PullStupidFees() (*types.Transaction, error) {
	return _StupidContract.Contract.PullStupidFees(&_StupidContract.TransactOpts)
}

// StupidContractStupidEventIterator is returned from FilterStupidEvent and is used to iterate over the raw logs and unpacked data for StupidEvent events raised by the StupidContract contract.
type StupidContractStupidEventIterator struct {
	Event *StupidContractStupidEvent // Event containing the contract specifics and raw log

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
func (it *StupidContractStupidEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StupidContractStupidEvent)
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
		it.Event = new(StupidContractStupidEvent)
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
func (it *StupidContractStupidEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StupidContractStupidEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StupidContractStupidEvent represents a StupidEvent event raised by the StupidContract contract.
type StupidContractStupidEvent struct {
	Index     *big.Int
	Sender    common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStupidEvent is a free log retrieval operation binding the contract event 0x7a4b508fdb79d1a00f12fc7c4d08d8b1a843a27f8576d2ed06f0691dc2b76d7e.
//
// Solidity: event StupidEvent(uint256 index, address indexed sender, uint256 timestamp)
func (_StupidContract *StupidContractFilterer) FilterStupidEvent(opts *bind.FilterOpts, sender []common.Address) (*StupidContractStupidEventIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StupidContract.contract.FilterLogs(opts, "StupidEvent", senderRule)
	if err != nil {
		return nil, err
	}
	return &StupidContractStupidEventIterator{contract: _StupidContract.contract, event: "StupidEvent", logs: logs, sub: sub}, nil
}

// WatchStupidEvent is a free log subscription operation binding the contract event 0x7a4b508fdb79d1a00f12fc7c4d08d8b1a843a27f8576d2ed06f0691dc2b76d7e.
//
// Solidity: event StupidEvent(uint256 index, address indexed sender, uint256 timestamp)
func (_StupidContract *StupidContractFilterer) WatchStupidEvent(opts *bind.WatchOpts, sink chan<- *StupidContractStupidEvent, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StupidContract.contract.WatchLogs(opts, "StupidEvent", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StupidContractStupidEvent)
				if err := _StupidContract.contract.UnpackLog(event, "StupidEvent", log); err != nil {
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

// ParseStupidEvent is a log parse operation binding the contract event 0x7a4b508fdb79d1a00f12fc7c4d08d8b1a843a27f8576d2ed06f0691dc2b76d7e.
//
// Solidity: event StupidEvent(uint256 index, address indexed sender, uint256 timestamp)
func (_StupidContract *StupidContractFilterer) ParseStupidEvent(log types.Log) (*StupidContractStupidEvent, error) {
	event := new(StupidContractStupidEvent)
	if err := _StupidContract.contract.UnpackLog(event, "StupidEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
