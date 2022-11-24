// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PaymentContract

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

// PaymentContractMetaData contains all meta data concerning the PaymentContract contract.
var PaymentContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeReceiver\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"PaymentContract__InvalidPaymentSubscriptionId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentContract__PaymentSubscriptionClosed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentContract__TransferFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"}],\"name\":\"NewPaymentSubscription\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"PaymentSubscriptionEnded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timeInterval\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_receivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_values\",\"type\":\"uint256[]\"}],\"name\":\"CreatePaymentSubscriptionEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"PayPaymentEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"payments\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fundsAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valuesSum\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"closed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561001057600080fd5b50604051610bde380380610bde83398101604081905261002f91610044565b336080526001600160a01b031660a052610074565b60006020828403121561005657600080fd5b81516001600160a01b038116811461006d57600080fd5b9392505050565b60805160a051610b386100a66000396000818161012601526106b301526000818160da01526103c40152610b386000f3fe6080604052600436106100555760003560e01c806348be6cd11461005a57806387d817891461006f5780638da5cb5b146100c8578063b3f0067414610114578063ddca3f4314610148578063ed660ff71461016b575b600080fd5b61006d6100683660046109bc565b61018b565b005b34801561007b57600080fd5b5061008f61008a366004610a3f565b61036b565b604080516001600160a01b03909616865260208601949094529284019190915260608301521515608082015260a0015b60405180910390f35b3480156100d457600080fd5b506100fc7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100bf565b34801561012057600080fd5b506100fc7f000000000000000000000000000000000000000000000000000000000000000081565b34801561015457600080fd5b5061015d600581565b6040519081526020016100bf565b34801561017757600080fd5b5061006d610186366004610a3f565b6103b9565b6000805b848110156101cf578383828181106101a9576101a9610a58565b90506020020135826101bb9190610a84565b9150806101c781610a9d565b91505061018f565b5060006040518060e00160405280336001600160a01b03168152602001348152602001888152602001878780806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250505090825250604080516020878102828101820190935287825292830192909188918891829185019084908082843760009201829052509385525050506020808301869052604092830182905284546001808201875595835291819020845160079093020180546001600160a01b0319166001600160a01b03909316929092178255838101519482019490945590820151600282015560608201518051929391926102d992600385019201906108bb565b50608082015180516102f5916004840191602090910190610920565b5060a0820151600582015560c0909101516006909101805460ff19169115159190911790556000547fd02fc200169694e30559f91d6cc10c6a31cf6597d45bfe0c12edee63f3f8f1de9061034b90600190610ab6565b60408051918252602082018a90520160405180910390a150505050505050565b6000818154811061037b57600080fd5b6000918252602090912060079091020180546001820154600283015460058401546006909401546001600160a01b0390931694509092909160ff1685565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146103ee57600080fd5b60005481106104105760405163b49a21ed60e01b815260040160405180910390fd5b600080828154811061042457610424610a58565b60009182526020918290206040805160e081018252600790930290910180546001600160a01b03168352600181015483850152600281015483830152600381018054835181870281018701909452808452939491936060860193928301828280156104b857602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161049a575b505050505081526020016004820180548060200260200160405190810160405280929190818152602001828054801561051057602002820191906000526020600020905b8154815260200190600101908083116104fc575b50505091835250506005820154602082015260069091015460ff16151560409091015260c0810151909150156105595760405163027dd98160e61b815260040160405180910390fd5b6000606460058360a0015161056e9190610ac9565b6105789190610ae0565b8260a001516105879190610a84565b905081602001518111156105a35761059e83610761565b505050565b80600084815481106105b7576105b7610a58565b906000526020600020906007020160010160008282546105d79190610ab6565b90915550600090505b8260600151518110156106ae5760008360600151828151811061060557610605610a58565b60200260200101516001600160a01b03168460800151838151811061062c5761062c610a58565b602002602001015160405160006040518083038185875af1925050503d8060008114610674576040519150601f19603f3d011682016040523d82523d6000602084013e610679565b606091505b505090508061069b57604051636207c02d60e01b815260040160405180910390fd5b50806106a681610a9d565b9150506105e0565b5060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316606460058560a001516106ee9190610ac9565b6106f89190610ae0565b604051600081818185875af1925050503d8060008114610734576040519150601f19603f3d011682016040523d82523d6000602084013e610739565b606091505b505090508061075b57604051636207c02d60e01b815260040160405180910390fd5b50505050565b600080828154811061077557610775610a58565b9060005260206000209060070201600101549050600080838154811061079d5761079d610a58565b9060005260206000209060070201600101819055506001600083815481106107c7576107c7610a58565b906000526020600020906007020160060160006101000a81548160ff021916908315150217905550600080838154811061080357610803610a58565b600091825260208220600790910201546040516001600160a01b039091169184919081818185875af1925050503d806000811461085c576040519150601f19603f3d011682016040523d82523d6000602084013e610861565b606091505b505090508061088357604051636207c02d60e01b815260040160405180910390fd5b6040518381527f90238c3bc8012cc678c7a9d36f8c10a9ad963187e0a36a28407f93843e1563479060200160405180910390a1505050565b828054828255906000526020600020908101928215610910579160200282015b8281111561091057825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906108db565b5061091c92915061095b565b5090565b828054828255906000526020600020908101928215610910579160200282015b82811115610910578251825591602001919060010190610940565b5b8082111561091c576000815560010161095c565b60008083601f84011261098257600080fd5b50813567ffffffffffffffff81111561099a57600080fd5b6020830191508360208260051b85010111156109b557600080fd5b9250929050565b600080600080600080608087890312156109d557600080fd5b8635955060208701359450604087013567ffffffffffffffff808211156109fb57600080fd5b610a078a838b01610970565b90965094506060890135915080821115610a2057600080fd5b50610a2d89828a01610970565b979a9699509497509295939492505050565b600060208284031215610a5157600080fd5b5035919050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b80820180821115610a9757610a97610a6e565b92915050565b600060018201610aaf57610aaf610a6e565b5060010190565b81810381811115610a9757610a97610a6e565b8082028115828204841417610a9757610a97610a6e565b600082610afd57634e487b7160e01b600052601260045260246000fd5b50049056fea26469706673582212203abda7b14a84c4baebf71067c43ac571347641d594115b1ff4106f2fde64ecc464736f6c63430008110033",
}

// PaymentContractABI is the input ABI used to generate the binding from.
// Deprecated: Use PaymentContractMetaData.ABI instead.
var PaymentContractABI = PaymentContractMetaData.ABI

// PaymentContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PaymentContractMetaData.Bin instead.
var PaymentContractBin = PaymentContractMetaData.Bin

// DeployPaymentContract deploys a new Ethereum contract, binding an instance of PaymentContract to it.
func DeployPaymentContract(auth *bind.TransactOpts, backend bind.ContractBackend, _feeReceiver common.Address) (common.Address, *types.Transaction, *PaymentContract, error) {
	parsed, err := PaymentContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PaymentContractBin), backend, _feeReceiver)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PaymentContract{PaymentContractCaller: PaymentContractCaller{contract: contract}, PaymentContractTransactor: PaymentContractTransactor{contract: contract}, PaymentContractFilterer: PaymentContractFilterer{contract: contract}}, nil
}

// PaymentContract is an auto generated Go binding around an Ethereum contract.
type PaymentContract struct {
	PaymentContractCaller     // Read-only binding to the contract
	PaymentContractTransactor // Write-only binding to the contract
	PaymentContractFilterer   // Log filterer for contract events
}

// PaymentContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type PaymentContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PaymentContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PaymentContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PaymentContractSession struct {
	Contract     *PaymentContract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PaymentContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PaymentContractCallerSession struct {
	Contract *PaymentContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// PaymentContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PaymentContractTransactorSession struct {
	Contract     *PaymentContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PaymentContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type PaymentContractRaw struct {
	Contract *PaymentContract // Generic contract binding to access the raw methods on
}

// PaymentContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PaymentContractCallerRaw struct {
	Contract *PaymentContractCaller // Generic read-only contract binding to access the raw methods on
}

// PaymentContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PaymentContractTransactorRaw struct {
	Contract *PaymentContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPaymentContract creates a new instance of PaymentContract, bound to a specific deployed contract.
func NewPaymentContract(address common.Address, backend bind.ContractBackend) (*PaymentContract, error) {
	contract, err := bindPaymentContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PaymentContract{PaymentContractCaller: PaymentContractCaller{contract: contract}, PaymentContractTransactor: PaymentContractTransactor{contract: contract}, PaymentContractFilterer: PaymentContractFilterer{contract: contract}}, nil
}

// NewPaymentContractCaller creates a new read-only instance of PaymentContract, bound to a specific deployed contract.
func NewPaymentContractCaller(address common.Address, caller bind.ContractCaller) (*PaymentContractCaller, error) {
	contract, err := bindPaymentContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentContractCaller{contract: contract}, nil
}

// NewPaymentContractTransactor creates a new write-only instance of PaymentContract, bound to a specific deployed contract.
func NewPaymentContractTransactor(address common.Address, transactor bind.ContractTransactor) (*PaymentContractTransactor, error) {
	contract, err := bindPaymentContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentContractTransactor{contract: contract}, nil
}

// NewPaymentContractFilterer creates a new log filterer instance of PaymentContract, bound to a specific deployed contract.
func NewPaymentContractFilterer(address common.Address, filterer bind.ContractFilterer) (*PaymentContractFilterer, error) {
	contract, err := bindPaymentContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PaymentContractFilterer{contract: contract}, nil
}

// bindPaymentContract binds a generic wrapper to an already deployed contract.
func bindPaymentContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PaymentContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PaymentContract *PaymentContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PaymentContract.Contract.PaymentContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PaymentContract *PaymentContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentContract.Contract.PaymentContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PaymentContract *PaymentContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PaymentContract.Contract.PaymentContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PaymentContract *PaymentContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PaymentContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PaymentContract *PaymentContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PaymentContract *PaymentContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PaymentContract.Contract.contract.Transact(opts, method, params...)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_PaymentContract *PaymentContractCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PaymentContract.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_PaymentContract *PaymentContractSession) Fee() (*big.Int, error) {
	return _PaymentContract.Contract.Fee(&_PaymentContract.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_PaymentContract *PaymentContractCallerSession) Fee() (*big.Int, error) {
	return _PaymentContract.Contract.Fee(&_PaymentContract.CallOpts)
}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_PaymentContract *PaymentContractCaller) FeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PaymentContract.contract.Call(opts, &out, "feeReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_PaymentContract *PaymentContractSession) FeeReceiver() (common.Address, error) {
	return _PaymentContract.Contract.FeeReceiver(&_PaymentContract.CallOpts)
}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_PaymentContract *PaymentContractCallerSession) FeeReceiver() (common.Address, error) {
	return _PaymentContract.Contract.FeeReceiver(&_PaymentContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PaymentContract *PaymentContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PaymentContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PaymentContract *PaymentContractSession) Owner() (common.Address, error) {
	return _PaymentContract.Contract.Owner(&_PaymentContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PaymentContract *PaymentContractCallerSession) Owner() (common.Address, error) {
	return _PaymentContract.Contract.Owner(&_PaymentContract.CallOpts)
}

// Payments is a free data retrieval call binding the contract method 0x87d81789.
//
// Solidity: function payments(uint256 ) view returns(address payer, uint256 fundsAmount, uint256 timeInterval, uint256 valuesSum, bool closed)
func (_PaymentContract *PaymentContractCaller) Payments(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Payer        common.Address
	FundsAmount  *big.Int
	TimeInterval *big.Int
	ValuesSum    *big.Int
	Closed       bool
}, error) {
	var out []interface{}
	err := _PaymentContract.contract.Call(opts, &out, "payments", arg0)

	outstruct := new(struct {
		Payer        common.Address
		FundsAmount  *big.Int
		TimeInterval *big.Int
		ValuesSum    *big.Int
		Closed       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Payer = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.FundsAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TimeInterval = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ValuesSum = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Closed = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// Payments is a free data retrieval call binding the contract method 0x87d81789.
//
// Solidity: function payments(uint256 ) view returns(address payer, uint256 fundsAmount, uint256 timeInterval, uint256 valuesSum, bool closed)
func (_PaymentContract *PaymentContractSession) Payments(arg0 *big.Int) (struct {
	Payer        common.Address
	FundsAmount  *big.Int
	TimeInterval *big.Int
	ValuesSum    *big.Int
	Closed       bool
}, error) {
	return _PaymentContract.Contract.Payments(&_PaymentContract.CallOpts, arg0)
}

// Payments is a free data retrieval call binding the contract method 0x87d81789.
//
// Solidity: function payments(uint256 ) view returns(address payer, uint256 fundsAmount, uint256 timeInterval, uint256 valuesSum, bool closed)
func (_PaymentContract *PaymentContractCallerSession) Payments(arg0 *big.Int) (struct {
	Payer        common.Address
	FundsAmount  *big.Int
	TimeInterval *big.Int
	ValuesSum    *big.Int
	Closed       bool
}, error) {
	return _PaymentContract.Contract.Payments(&_PaymentContract.CallOpts, arg0)
}

// CreatePaymentSubscriptionEth is a paid mutator transaction binding the contract method 0x48be6cd1.
//
// Solidity: function CreatePaymentSubscriptionEth(uint256 _startTime, uint256 _timeInterval, address[] _receivers, uint256[] _values) payable returns()
func (_PaymentContract *PaymentContractTransactor) CreatePaymentSubscriptionEth(opts *bind.TransactOpts, _startTime *big.Int, _timeInterval *big.Int, _receivers []common.Address, _values []*big.Int) (*types.Transaction, error) {
	return _PaymentContract.contract.Transact(opts, "CreatePaymentSubscriptionEth", _startTime, _timeInterval, _receivers, _values)
}

// CreatePaymentSubscriptionEth is a paid mutator transaction binding the contract method 0x48be6cd1.
//
// Solidity: function CreatePaymentSubscriptionEth(uint256 _startTime, uint256 _timeInterval, address[] _receivers, uint256[] _values) payable returns()
func (_PaymentContract *PaymentContractSession) CreatePaymentSubscriptionEth(_startTime *big.Int, _timeInterval *big.Int, _receivers []common.Address, _values []*big.Int) (*types.Transaction, error) {
	return _PaymentContract.Contract.CreatePaymentSubscriptionEth(&_PaymentContract.TransactOpts, _startTime, _timeInterval, _receivers, _values)
}

// CreatePaymentSubscriptionEth is a paid mutator transaction binding the contract method 0x48be6cd1.
//
// Solidity: function CreatePaymentSubscriptionEth(uint256 _startTime, uint256 _timeInterval, address[] _receivers, uint256[] _values) payable returns()
func (_PaymentContract *PaymentContractTransactorSession) CreatePaymentSubscriptionEth(_startTime *big.Int, _timeInterval *big.Int, _receivers []common.Address, _values []*big.Int) (*types.Transaction, error) {
	return _PaymentContract.Contract.CreatePaymentSubscriptionEth(&_PaymentContract.TransactOpts, _startTime, _timeInterval, _receivers, _values)
}

// PayPaymentEth is a paid mutator transaction binding the contract method 0xed660ff7.
//
// Solidity: function PayPaymentEth(uint256 index) returns()
func (_PaymentContract *PaymentContractTransactor) PayPaymentEth(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _PaymentContract.contract.Transact(opts, "PayPaymentEth", index)
}

// PayPaymentEth is a paid mutator transaction binding the contract method 0xed660ff7.
//
// Solidity: function PayPaymentEth(uint256 index) returns()
func (_PaymentContract *PaymentContractSession) PayPaymentEth(index *big.Int) (*types.Transaction, error) {
	return _PaymentContract.Contract.PayPaymentEth(&_PaymentContract.TransactOpts, index)
}

// PayPaymentEth is a paid mutator transaction binding the contract method 0xed660ff7.
//
// Solidity: function PayPaymentEth(uint256 index) returns()
func (_PaymentContract *PaymentContractTransactorSession) PayPaymentEth(index *big.Int) (*types.Transaction, error) {
	return _PaymentContract.Contract.PayPaymentEth(&_PaymentContract.TransactOpts, index)
}

// PaymentContractNewPaymentSubscriptionIterator is returned from FilterNewPaymentSubscription and is used to iterate over the raw logs and unpacked data for NewPaymentSubscription events raised by the PaymentContract contract.
type PaymentContractNewPaymentSubscriptionIterator struct {
	Event *PaymentContractNewPaymentSubscription // Event containing the contract specifics and raw log

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
func (it *PaymentContractNewPaymentSubscriptionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentContractNewPaymentSubscription)
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
		it.Event = new(PaymentContractNewPaymentSubscription)
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
func (it *PaymentContractNewPaymentSubscriptionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentContractNewPaymentSubscriptionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentContractNewPaymentSubscription represents a NewPaymentSubscription event raised by the PaymentContract contract.
type PaymentContractNewPaymentSubscription struct {
	Index     *big.Int
	StartTime *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewPaymentSubscription is a free log retrieval operation binding the contract event 0xd02fc200169694e30559f91d6cc10c6a31cf6597d45bfe0c12edee63f3f8f1de.
//
// Solidity: event NewPaymentSubscription(uint256 index, uint256 startTime)
func (_PaymentContract *PaymentContractFilterer) FilterNewPaymentSubscription(opts *bind.FilterOpts) (*PaymentContractNewPaymentSubscriptionIterator, error) {

	logs, sub, err := _PaymentContract.contract.FilterLogs(opts, "NewPaymentSubscription")
	if err != nil {
		return nil, err
	}
	return &PaymentContractNewPaymentSubscriptionIterator{contract: _PaymentContract.contract, event: "NewPaymentSubscription", logs: logs, sub: sub}, nil
}

// WatchNewPaymentSubscription is a free log subscription operation binding the contract event 0xd02fc200169694e30559f91d6cc10c6a31cf6597d45bfe0c12edee63f3f8f1de.
//
// Solidity: event NewPaymentSubscription(uint256 index, uint256 startTime)
func (_PaymentContract *PaymentContractFilterer) WatchNewPaymentSubscription(opts *bind.WatchOpts, sink chan<- *PaymentContractNewPaymentSubscription) (event.Subscription, error) {

	logs, sub, err := _PaymentContract.contract.WatchLogs(opts, "NewPaymentSubscription")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentContractNewPaymentSubscription)
				if err := _PaymentContract.contract.UnpackLog(event, "NewPaymentSubscription", log); err != nil {
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

// ParseNewPaymentSubscription is a log parse operation binding the contract event 0xd02fc200169694e30559f91d6cc10c6a31cf6597d45bfe0c12edee63f3f8f1de.
//
// Solidity: event NewPaymentSubscription(uint256 index, uint256 startTime)
func (_PaymentContract *PaymentContractFilterer) ParseNewPaymentSubscription(log types.Log) (*PaymentContractNewPaymentSubscription, error) {
	event := new(PaymentContractNewPaymentSubscription)
	if err := _PaymentContract.contract.UnpackLog(event, "NewPaymentSubscription", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentContractPaymentSubscriptionEndedIterator is returned from FilterPaymentSubscriptionEnded and is used to iterate over the raw logs and unpacked data for PaymentSubscriptionEnded events raised by the PaymentContract contract.
type PaymentContractPaymentSubscriptionEndedIterator struct {
	Event *PaymentContractPaymentSubscriptionEnded // Event containing the contract specifics and raw log

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
func (it *PaymentContractPaymentSubscriptionEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentContractPaymentSubscriptionEnded)
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
		it.Event = new(PaymentContractPaymentSubscriptionEnded)
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
func (it *PaymentContractPaymentSubscriptionEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentContractPaymentSubscriptionEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentContractPaymentSubscriptionEnded represents a PaymentSubscriptionEnded event raised by the PaymentContract contract.
type PaymentContractPaymentSubscriptionEnded struct {
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPaymentSubscriptionEnded is a free log retrieval operation binding the contract event 0x90238c3bc8012cc678c7a9d36f8c10a9ad963187e0a36a28407f93843e156347.
//
// Solidity: event PaymentSubscriptionEnded(uint256 index)
func (_PaymentContract *PaymentContractFilterer) FilterPaymentSubscriptionEnded(opts *bind.FilterOpts) (*PaymentContractPaymentSubscriptionEndedIterator, error) {

	logs, sub, err := _PaymentContract.contract.FilterLogs(opts, "PaymentSubscriptionEnded")
	if err != nil {
		return nil, err
	}
	return &PaymentContractPaymentSubscriptionEndedIterator{contract: _PaymentContract.contract, event: "PaymentSubscriptionEnded", logs: logs, sub: sub}, nil
}

// WatchPaymentSubscriptionEnded is a free log subscription operation binding the contract event 0x90238c3bc8012cc678c7a9d36f8c10a9ad963187e0a36a28407f93843e156347.
//
// Solidity: event PaymentSubscriptionEnded(uint256 index)
func (_PaymentContract *PaymentContractFilterer) WatchPaymentSubscriptionEnded(opts *bind.WatchOpts, sink chan<- *PaymentContractPaymentSubscriptionEnded) (event.Subscription, error) {

	logs, sub, err := _PaymentContract.contract.WatchLogs(opts, "PaymentSubscriptionEnded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentContractPaymentSubscriptionEnded)
				if err := _PaymentContract.contract.UnpackLog(event, "PaymentSubscriptionEnded", log); err != nil {
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

// ParsePaymentSubscriptionEnded is a log parse operation binding the contract event 0x90238c3bc8012cc678c7a9d36f8c10a9ad963187e0a36a28407f93843e156347.
//
// Solidity: event PaymentSubscriptionEnded(uint256 index)
func (_PaymentContract *PaymentContractFilterer) ParsePaymentSubscriptionEnded(log types.Log) (*PaymentContractPaymentSubscriptionEnded, error) {
	event := new(PaymentContractPaymentSubscriptionEnded)
	if err := _PaymentContract.contract.UnpackLog(event, "PaymentSubscriptionEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
