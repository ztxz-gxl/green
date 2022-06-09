// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth

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

// AbiABI is the input ABI used to generate the binding from.
const AbiABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ip\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"time\",\"type\":\"bytes32\"}],\"name\":\"ADDBANK\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ip\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_limitTime\",\"type\":\"uint256\"}],\"name\":\"VOTE\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"eventType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"VoteEvt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ip\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"initializationEVT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ip\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"takePart\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ip\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AddERC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_year\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_interest\",\"type\":\"uint256\"}],\"name\":\"Bank\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ip\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_pName\",\"type\":\"string\"}],\"name\":\"createVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_operation\",\"type\":\"string\"}],\"name\":\"Deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Exchange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"MASSID\",\"type\":\"uint256\"}],\"name\":\"finish\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"GetInterest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"Initialization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_year\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_neutralizeTotal\",\"type\":\"uint256\"}],\"name\":\"ReleaseNeutralize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ip\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_reason\",\"type\":\"string\"}],\"name\":\"Report\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ip\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SubERC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_fun\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_NeutralizedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_personalAmount\",\"type\":\"uint256\"}],\"name\":\"TakeAPart\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_operation\",\"type\":\"string\"}],\"name\":\"TakeOutBank\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pId\",\"type\":\"uint256\"}],\"name\":\"Voting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"BankAdminList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"ip\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"operationAmount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"operation\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"operationTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BankAdminOperationNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"BankVoucherHash\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"partakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interest\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"takeOut\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentYear\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializationSTU\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"messageID\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"ip\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"fun\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"stu\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"NeutralizeYear\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"publisher\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"neutralizeTotal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"neutralizePersonals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"peopleNUM\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recordTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"UpToStandard\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"CurrentQuota\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"newPersonal\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"personalName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"personalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NeutralizedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"registrationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NumberOfReported\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NUM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RVNUM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UserNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"ip\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"vName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"chairperson\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limitTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voterAddr\",\"type\":\"address\"}],\"name\":\"VotingT_F\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Abi is an auto generated Go binding around an Ethereum contract.
type Abi struct {
	AbiCaller     // Read-only binding to the contract
	AbiTransactor // Write-only binding to the contract
	AbiFilterer   // Log filterer for contract events
}

// AbiCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbiSession struct {
	Contract     *Abi              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbiCallerSession struct {
	Contract *AbiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AbiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbiTransactorSession struct {
	Contract     *AbiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbiRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbiRaw struct {
	Contract *Abi // Generic contract binding to access the raw methods on
}

// AbiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbiCallerRaw struct {
	Contract *AbiCaller // Generic read-only contract binding to access the raw methods on
}

// AbiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbiTransactorRaw struct {
	Contract *AbiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbi creates a new instance of Abi, bound to a specific deployed contract.
func NewAbi(address common.Address, backend bind.ContractBackend) (*Abi, error) {
	contract, err := bindAbi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Abi{AbiCaller: AbiCaller{contract: contract}, AbiTransactor: AbiTransactor{contract: contract}, AbiFilterer: AbiFilterer{contract: contract}}, nil
}

// NewAbiCaller creates a new read-only instance of Abi, bound to a specific deployed contract.
func NewAbiCaller(address common.Address, caller bind.ContractCaller) (*AbiCaller, error) {
	contract, err := bindAbi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbiCaller{contract: contract}, nil
}

// NewAbiTransactor creates a new write-only instance of Abi, bound to a specific deployed contract.
func NewAbiTransactor(address common.Address, transactor bind.ContractTransactor) (*AbiTransactor, error) {
	contract, err := bindAbi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbiTransactor{contract: contract}, nil
}

// NewAbiFilterer creates a new log filterer instance of Abi, bound to a specific deployed contract.
func NewAbiFilterer(address common.Address, filterer bind.ContractFilterer) (*AbiFilterer, error) {
	contract, err := bindAbi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbiFilterer{contract: contract}, nil
}

// bindAbi binds a generic wrapper to an already deployed contract.
func bindAbi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AbiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abi *AbiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abi.Contract.AbiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abi *AbiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.Contract.AbiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abi *AbiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abi.Contract.AbiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abi *AbiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abi *AbiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abi *AbiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abi.Contract.contract.Transact(opts, method, params...)
}

// BankAdminList is a free data retrieval call binding the contract method 0x4a3915d2.
//
// Solidity: function BankAdminList(uint256 ) view returns(address ip, uint256 operationAmount, string operation, uint256 operationTime)
func (_Abi *AbiCaller) BankAdminList(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Ip              common.Address
	OperationAmount *big.Int
	Operation       string
	OperationTime   *big.Int
}, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "BankAdminList", arg0)

	outstruct := new(struct {
		Ip              common.Address
		OperationAmount *big.Int
		Operation       string
		OperationTime   *big.Int
	})

	outstruct.Ip = out[0].(common.Address)
	outstruct.OperationAmount = out[1].(*big.Int)
	outstruct.Operation = out[2].(string)
	outstruct.OperationTime = out[3].(*big.Int)

	return *outstruct, err

}

// BankAdminList is a free data retrieval call binding the contract method 0x4a3915d2.
//
// Solidity: function BankAdminList(uint256 ) view returns(address ip, uint256 operationAmount, string operation, uint256 operationTime)
func (_Abi *AbiSession) BankAdminList(arg0 *big.Int) (struct {
	Ip              common.Address
	OperationAmount *big.Int
	Operation       string
	OperationTime   *big.Int
}, error) {
	return _Abi.Contract.BankAdminList(&_Abi.CallOpts, arg0)
}

// BankAdminList is a free data retrieval call binding the contract method 0x4a3915d2.
//
// Solidity: function BankAdminList(uint256 ) view returns(address ip, uint256 operationAmount, string operation, uint256 operationTime)
func (_Abi *AbiCallerSession) BankAdminList(arg0 *big.Int) (struct {
	Ip              common.Address
	OperationAmount *big.Int
	Operation       string
	OperationTime   *big.Int
}, error) {
	return _Abi.Contract.BankAdminList(&_Abi.CallOpts, arg0)
}

// BankAdminOperationNum is a free data retrieval call binding the contract method 0x4f5f63c5.
//
// Solidity: function BankAdminOperationNum() view returns(uint256)
func (_Abi *AbiCaller) BankAdminOperationNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "BankAdminOperationNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BankAdminOperationNum is a free data retrieval call binding the contract method 0x4f5f63c5.
//
// Solidity: function BankAdminOperationNum() view returns(uint256)
func (_Abi *AbiSession) BankAdminOperationNum() (*big.Int, error) {
	return _Abi.Contract.BankAdminOperationNum(&_Abi.CallOpts)
}

// BankAdminOperationNum is a free data retrieval call binding the contract method 0x4f5f63c5.
//
// Solidity: function BankAdminOperationNum() view returns(uint256)
func (_Abi *AbiCallerSession) BankAdminOperationNum() (*big.Int, error) {
	return _Abi.Contract.BankAdminOperationNum(&_Abi.CallOpts)
}

// BankVoucherHash is a free data retrieval call binding the contract method 0x06d9f503.
//
// Solidity: function BankVoucherHash(bytes32 ) view returns(address participant, uint256 partakeAmount, uint256 deadline, uint256 interest, bool takeOut)
func (_Abi *AbiCaller) BankVoucherHash(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Participant   common.Address
	PartakeAmount *big.Int
	Deadline      *big.Int
	Interest      *big.Int
	TakeOut       bool
}, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "BankVoucherHash", arg0)

	outstruct := new(struct {
		Participant   common.Address
		PartakeAmount *big.Int
		Deadline      *big.Int
		Interest      *big.Int
		TakeOut       bool
	})

	outstruct.Participant = out[0].(common.Address)
	outstruct.PartakeAmount = out[1].(*big.Int)
	outstruct.Deadline = out[2].(*big.Int)
	outstruct.Interest = out[3].(*big.Int)
	outstruct.TakeOut = out[4].(bool)

	return *outstruct, err

}

// BankVoucherHash is a free data retrieval call binding the contract method 0x06d9f503.
//
// Solidity: function BankVoucherHash(bytes32 ) view returns(address participant, uint256 partakeAmount, uint256 deadline, uint256 interest, bool takeOut)
func (_Abi *AbiSession) BankVoucherHash(arg0 [32]byte) (struct {
	Participant   common.Address
	PartakeAmount *big.Int
	Deadline      *big.Int
	Interest      *big.Int
	TakeOut       bool
}, error) {
	return _Abi.Contract.BankVoucherHash(&_Abi.CallOpts, arg0)
}

// BankVoucherHash is a free data retrieval call binding the contract method 0x06d9f503.
//
// Solidity: function BankVoucherHash(bytes32 ) view returns(address participant, uint256 partakeAmount, uint256 deadline, uint256 interest, bool takeOut)
func (_Abi *AbiCallerSession) BankVoucherHash(arg0 [32]byte) (struct {
	Participant   common.Address
	PartakeAmount *big.Int
	Deadline      *big.Int
	Interest      *big.Int
	TakeOut       bool
}, error) {
	return _Abi.Contract.BankVoucherHash(&_Abi.CallOpts, arg0)
}

// NUM is a free data retrieval call binding the contract method 0x897cb91e.
//
// Solidity: function NUM() view returns(uint256)
func (_Abi *AbiCaller) NUM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "NUM")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NUM is a free data retrieval call binding the contract method 0x897cb91e.
//
// Solidity: function NUM() view returns(uint256)
func (_Abi *AbiSession) NUM() (*big.Int, error) {
	return _Abi.Contract.NUM(&_Abi.CallOpts)
}

// NUM is a free data retrieval call binding the contract method 0x897cb91e.
//
// Solidity: function NUM() view returns(uint256)
func (_Abi *AbiCallerSession) NUM() (*big.Int, error) {
	return _Abi.Contract.NUM(&_Abi.CallOpts)
}

// NeutralizeYear is a free data retrieval call binding the contract method 0xd65afc7c.
//
// Solidity: function NeutralizeYear(uint256 ) view returns(address publisher, uint256 neutralizeTotal, uint256 neutralizePersonals, uint256 peopleNUM, uint256 recordTime, bool UpToStandard, uint256 CurrentQuota)
func (_Abi *AbiCaller) NeutralizeYear(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Publisher           common.Address
	NeutralizeTotal     *big.Int
	NeutralizePersonals *big.Int
	PeopleNUM           *big.Int
	RecordTime          *big.Int
	UpToStandard        bool
	CurrentQuota        *big.Int
}, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "NeutralizeYear", arg0)

	outstruct := new(struct {
		Publisher           common.Address
		NeutralizeTotal     *big.Int
		NeutralizePersonals *big.Int
		PeopleNUM           *big.Int
		RecordTime          *big.Int
		UpToStandard        bool
		CurrentQuota        *big.Int
	})

	outstruct.Publisher = out[0].(common.Address)
	outstruct.NeutralizeTotal = out[1].(*big.Int)
	outstruct.NeutralizePersonals = out[2].(*big.Int)
	outstruct.PeopleNUM = out[3].(*big.Int)
	outstruct.RecordTime = out[4].(*big.Int)
	outstruct.UpToStandard = out[5].(bool)
	outstruct.CurrentQuota = out[6].(*big.Int)

	return *outstruct, err

}

// NeutralizeYear is a free data retrieval call binding the contract method 0xd65afc7c.
//
// Solidity: function NeutralizeYear(uint256 ) view returns(address publisher, uint256 neutralizeTotal, uint256 neutralizePersonals, uint256 peopleNUM, uint256 recordTime, bool UpToStandard, uint256 CurrentQuota)
func (_Abi *AbiSession) NeutralizeYear(arg0 *big.Int) (struct {
	Publisher           common.Address
	NeutralizeTotal     *big.Int
	NeutralizePersonals *big.Int
	PeopleNUM           *big.Int
	RecordTime          *big.Int
	UpToStandard        bool
	CurrentQuota        *big.Int
}, error) {
	return _Abi.Contract.NeutralizeYear(&_Abi.CallOpts, arg0)
}

// NeutralizeYear is a free data retrieval call binding the contract method 0xd65afc7c.
//
// Solidity: function NeutralizeYear(uint256 ) view returns(address publisher, uint256 neutralizeTotal, uint256 neutralizePersonals, uint256 peopleNUM, uint256 recordTime, bool UpToStandard, uint256 CurrentQuota)
func (_Abi *AbiCallerSession) NeutralizeYear(arg0 *big.Int) (struct {
	Publisher           common.Address
	NeutralizeTotal     *big.Int
	NeutralizePersonals *big.Int
	PeopleNUM           *big.Int
	RecordTime          *big.Int
	UpToStandard        bool
	CurrentQuota        *big.Int
}, error) {
	return _Abi.Contract.NeutralizeYear(&_Abi.CallOpts, arg0)
}

// RVNUM is a free data retrieval call binding the contract method 0x82718101.
//
// Solidity: function RVNUM() view returns(uint256)
func (_Abi *AbiCaller) RVNUM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "RVNUM")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RVNUM is a free data retrieval call binding the contract method 0x82718101.
//
// Solidity: function RVNUM() view returns(uint256)
func (_Abi *AbiSession) RVNUM() (*big.Int, error) {
	return _Abi.Contract.RVNUM(&_Abi.CallOpts)
}

// RVNUM is a free data retrieval call binding the contract method 0x82718101.
//
// Solidity: function RVNUM() view returns(uint256)
func (_Abi *AbiCallerSession) RVNUM() (*big.Int, error) {
	return _Abi.Contract.RVNUM(&_Abi.CallOpts)
}

// UserNum is a free data retrieval call binding the contract method 0x7747ae60.
//
// Solidity: function UserNum() view returns(uint256)
func (_Abi *AbiCaller) UserNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "UserNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserNum is a free data retrieval call binding the contract method 0x7747ae60.
//
// Solidity: function UserNum() view returns(uint256)
func (_Abi *AbiSession) UserNum() (*big.Int, error) {
	return _Abi.Contract.UserNum(&_Abi.CallOpts)
}

// UserNum is a free data retrieval call binding the contract method 0x7747ae60.
//
// Solidity: function UserNum() view returns(uint256)
func (_Abi *AbiCallerSession) UserNum() (*big.Int, error) {
	return _Abi.Contract.UserNum(&_Abi.CallOpts)
}

// VotingTF is a free data retrieval call binding the contract method 0x76e27f59.
//
// Solidity: function VotingT_F(uint256 pId, address voterAddr) view returns(uint256)
func (_Abi *AbiCaller) VotingTF(opts *bind.CallOpts, pId *big.Int, voterAddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "VotingT_F", pId, voterAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingTF is a free data retrieval call binding the contract method 0x76e27f59.
//
// Solidity: function VotingT_F(uint256 pId, address voterAddr) view returns(uint256)
func (_Abi *AbiSession) VotingTF(pId *big.Int, voterAddr common.Address) (*big.Int, error) {
	return _Abi.Contract.VotingTF(&_Abi.CallOpts, pId, voterAddr)
}

// VotingTF is a free data retrieval call binding the contract method 0x76e27f59.
//
// Solidity: function VotingT_F(uint256 pId, address voterAddr) view returns(uint256)
func (_Abi *AbiCallerSession) VotingTF(pId *big.Int, voterAddr common.Address) (*big.Int, error) {
	return _Abi.Contract.VotingTF(&_Abi.CallOpts, pId, voterAddr)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Abi *AbiCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Abi *AbiSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Abi.Contract.Allowance(&_Abi.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Abi *AbiCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Abi.Contract.Allowance(&_Abi.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Abi *AbiCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Abi *AbiSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Abi.Contract.BalanceOf(&_Abi.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Abi *AbiCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Abi.Contract.BalanceOf(&_Abi.CallOpts, account)
}

// CurrentYear is a free data retrieval call binding the contract method 0x0b5a006b.
//
// Solidity: function currentYear() view returns(uint256)
func (_Abi *AbiCaller) CurrentYear(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "currentYear")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentYear is a free data retrieval call binding the contract method 0x0b5a006b.
//
// Solidity: function currentYear() view returns(uint256)
func (_Abi *AbiSession) CurrentYear() (*big.Int, error) {
	return _Abi.Contract.CurrentYear(&_Abi.CallOpts)
}

// CurrentYear is a free data retrieval call binding the contract method 0x0b5a006b.
//
// Solidity: function currentYear() view returns(uint256)
func (_Abi *AbiCallerSession) CurrentYear() (*big.Int, error) {
	return _Abi.Contract.CurrentYear(&_Abi.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Abi *AbiCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Abi *AbiSession) Decimals() (uint8, error) {
	return _Abi.Contract.Decimals(&_Abi.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Abi *AbiCallerSession) Decimals() (uint8, error) {
	return _Abi.Contract.Decimals(&_Abi.CallOpts)
}

// InitializationSTU is a free data retrieval call binding the contract method 0x796fc5cd.
//
// Solidity: function initializationSTU() view returns(bool)
func (_Abi *AbiCaller) InitializationSTU(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "initializationSTU")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InitializationSTU is a free data retrieval call binding the contract method 0x796fc5cd.
//
// Solidity: function initializationSTU() view returns(bool)
func (_Abi *AbiSession) InitializationSTU() (bool, error) {
	return _Abi.Contract.InitializationSTU(&_Abi.CallOpts)
}

// InitializationSTU is a free data retrieval call binding the contract method 0x796fc5cd.
//
// Solidity: function initializationSTU() view returns(bool)
func (_Abi *AbiCallerSession) InitializationSTU() (bool, error) {
	return _Abi.Contract.InitializationSTU(&_Abi.CallOpts)
}

// MessageID is a free data retrieval call binding the contract method 0x2b6aeaad.
//
// Solidity: function messageID(uint256 ) view returns(address ip, string fun, uint256 time, bool stu)
func (_Abi *AbiCaller) MessageID(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Ip   common.Address
	Fun  string
	Time *big.Int
	Stu  bool
}, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "messageID", arg0)

	outstruct := new(struct {
		Ip   common.Address
		Fun  string
		Time *big.Int
		Stu  bool
	})

	outstruct.Ip = out[0].(common.Address)
	outstruct.Fun = out[1].(string)
	outstruct.Time = out[2].(*big.Int)
	outstruct.Stu = out[3].(bool)

	return *outstruct, err

}

// MessageID is a free data retrieval call binding the contract method 0x2b6aeaad.
//
// Solidity: function messageID(uint256 ) view returns(address ip, string fun, uint256 time, bool stu)
func (_Abi *AbiSession) MessageID(arg0 *big.Int) (struct {
	Ip   common.Address
	Fun  string
	Time *big.Int
	Stu  bool
}, error) {
	return _Abi.Contract.MessageID(&_Abi.CallOpts, arg0)
}

// MessageID is a free data retrieval call binding the contract method 0x2b6aeaad.
//
// Solidity: function messageID(uint256 ) view returns(address ip, string fun, uint256 time, bool stu)
func (_Abi *AbiCallerSession) MessageID(arg0 *big.Int) (struct {
	Ip   common.Address
	Fun  string
	Time *big.Int
	Stu  bool
}, error) {
	return _Abi.Contract.MessageID(&_Abi.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Abi *AbiCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Abi *AbiSession) Name() (string, error) {
	return _Abi.Contract.Name(&_Abi.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Abi *AbiCallerSession) Name() (string, error) {
	return _Abi.Contract.Name(&_Abi.CallOpts)
}

// NewPersonal is a free data retrieval call binding the contract method 0x23515fed.
//
// Solidity: function newPersonal(address ) view returns(string personalName, uint256 personalAmount, uint256 NeutralizedAmount, uint256 registrationTime, uint256 NumberOfReported)
func (_Abi *AbiCaller) NewPersonal(opts *bind.CallOpts, arg0 common.Address) (struct {
	PersonalName      string
	PersonalAmount    *big.Int
	NeutralizedAmount *big.Int
	RegistrationTime  *big.Int
	NumberOfReported  *big.Int
}, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "newPersonal", arg0)

	outstruct := new(struct {
		PersonalName      string
		PersonalAmount    *big.Int
		NeutralizedAmount *big.Int
		RegistrationTime  *big.Int
		NumberOfReported  *big.Int
	})

	outstruct.PersonalName = out[0].(string)
	outstruct.PersonalAmount = out[1].(*big.Int)
	outstruct.NeutralizedAmount = out[2].(*big.Int)
	outstruct.RegistrationTime = out[3].(*big.Int)
	outstruct.NumberOfReported = out[4].(*big.Int)

	return *outstruct, err

}

// NewPersonal is a free data retrieval call binding the contract method 0x23515fed.
//
// Solidity: function newPersonal(address ) view returns(string personalName, uint256 personalAmount, uint256 NeutralizedAmount, uint256 registrationTime, uint256 NumberOfReported)
func (_Abi *AbiSession) NewPersonal(arg0 common.Address) (struct {
	PersonalName      string
	PersonalAmount    *big.Int
	NeutralizedAmount *big.Int
	RegistrationTime  *big.Int
	NumberOfReported  *big.Int
}, error) {
	return _Abi.Contract.NewPersonal(&_Abi.CallOpts, arg0)
}

// NewPersonal is a free data retrieval call binding the contract method 0x23515fed.
//
// Solidity: function newPersonal(address ) view returns(string personalName, uint256 personalAmount, uint256 NeutralizedAmount, uint256 registrationTime, uint256 NumberOfReported)
func (_Abi *AbiCallerSession) NewPersonal(arg0 common.Address) (struct {
	PersonalName      string
	PersonalAmount    *big.Int
	NeutralizedAmount *big.Int
	RegistrationTime  *big.Int
	NumberOfReported  *big.Int
}, error) {
	return _Abi.Contract.NewPersonal(&_Abi.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Abi *AbiCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Abi *AbiSession) Symbol() (string, error) {
	return _Abi.Contract.Symbol(&_Abi.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Abi *AbiCallerSession) Symbol() (string, error) {
	return _Abi.Contract.Symbol(&_Abi.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Abi *AbiCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Abi *AbiSession) TotalSupply() (*big.Int, error) {
	return _Abi.Contract.TotalSupply(&_Abi.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Abi *AbiCallerSession) TotalSupply() (*big.Int, error) {
	return _Abi.Contract.TotalSupply(&_Abi.CallOpts)
}

// Vote is a free data retrieval call binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 ) view returns(address ip, string vName, address chairperson, uint256 voteCount, uint256 limitTime)
func (_Abi *AbiCaller) Vote(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Ip          common.Address
	VName       string
	Chairperson common.Address
	VoteCount   *big.Int
	LimitTime   *big.Int
}, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "vote", arg0)

	outstruct := new(struct {
		Ip          common.Address
		VName       string
		Chairperson common.Address
		VoteCount   *big.Int
		LimitTime   *big.Int
	})

	outstruct.Ip = out[0].(common.Address)
	outstruct.VName = out[1].(string)
	outstruct.Chairperson = out[2].(common.Address)
	outstruct.VoteCount = out[3].(*big.Int)
	outstruct.LimitTime = out[4].(*big.Int)

	return *outstruct, err

}

// Vote is a free data retrieval call binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 ) view returns(address ip, string vName, address chairperson, uint256 voteCount, uint256 limitTime)
func (_Abi *AbiSession) Vote(arg0 *big.Int) (struct {
	Ip          common.Address
	VName       string
	Chairperson common.Address
	VoteCount   *big.Int
	LimitTime   *big.Int
}, error) {
	return _Abi.Contract.Vote(&_Abi.CallOpts, arg0)
}

// Vote is a free data retrieval call binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 ) view returns(address ip, string vName, address chairperson, uint256 voteCount, uint256 limitTime)
func (_Abi *AbiCallerSession) Vote(arg0 *big.Int) (struct {
	Ip          common.Address
	VName       string
	Chairperson common.Address
	VoteCount   *big.Int
	LimitTime   *big.Int
}, error) {
	return _Abi.Contract.Vote(&_Abi.CallOpts, arg0)
}

// AddERC is a paid mutator transaction binding the contract method 0xf3b9113a.
//
// Solidity: function AddERC(address ip, uint256 amount) returns()
func (_Abi *AbiTransactor) AddERC(opts *bind.TransactOpts, ip common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "AddERC", ip, amount)
}

// AddERC is a paid mutator transaction binding the contract method 0xf3b9113a.
//
// Solidity: function AddERC(address ip, uint256 amount) returns()
func (_Abi *AbiSession) AddERC(ip common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.AddERC(&_Abi.TransactOpts, ip, amount)
}

// AddERC is a paid mutator transaction binding the contract method 0xf3b9113a.
//
// Solidity: function AddERC(address ip, uint256 amount) returns()
func (_Abi *AbiTransactorSession) AddERC(ip common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.AddERC(&_Abi.TransactOpts, ip, amount)
}

// Bank is a paid mutator transaction binding the contract method 0x9108ebeb.
//
// Solidity: function Bank(uint256 _amount, uint256 _year, uint256 _interest) returns()
func (_Abi *AbiTransactor) Bank(opts *bind.TransactOpts, _amount *big.Int, _year *big.Int, _interest *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "Bank", _amount, _year, _interest)
}

// Bank is a paid mutator transaction binding the contract method 0x9108ebeb.
//
// Solidity: function Bank(uint256 _amount, uint256 _year, uint256 _interest) returns()
func (_Abi *AbiSession) Bank(_amount *big.Int, _year *big.Int, _interest *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.Bank(&_Abi.TransactOpts, _amount, _year, _interest)
}

// Bank is a paid mutator transaction binding the contract method 0x9108ebeb.
//
// Solidity: function Bank(uint256 _amount, uint256 _year, uint256 _interest) returns()
func (_Abi *AbiTransactorSession) Bank(_amount *big.Int, _year *big.Int, _interest *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.Bank(&_Abi.TransactOpts, _amount, _year, _interest)
}

// Deposit is a paid mutator transaction binding the contract method 0x02411f5a.
//
// Solidity: function Deposit(uint256 _amount, string _operation) returns()
func (_Abi *AbiTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int, _operation string) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "Deposit", _amount, _operation)
}

// Deposit is a paid mutator transaction binding the contract method 0x02411f5a.
//
// Solidity: function Deposit(uint256 _amount, string _operation) returns()
func (_Abi *AbiSession) Deposit(_amount *big.Int, _operation string) (*types.Transaction, error) {
	return _Abi.Contract.Deposit(&_Abi.TransactOpts, _amount, _operation)
}

// Deposit is a paid mutator transaction binding the contract method 0x02411f5a.
//
// Solidity: function Deposit(uint256 _amount, string _operation) returns()
func (_Abi *AbiTransactorSession) Deposit(_amount *big.Int, _operation string) (*types.Transaction, error) {
	return _Abi.Contract.Deposit(&_Abi.TransactOpts, _amount, _operation)
}

// Exchange is a paid mutator transaction binding the contract method 0x11d92a08.
//
// Solidity: function Exchange(uint256 _amount) returns()
func (_Abi *AbiTransactor) Exchange(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "Exchange", _amount)
}

// Exchange is a paid mutator transaction binding the contract method 0x11d92a08.
//
// Solidity: function Exchange(uint256 _amount) returns()
func (_Abi *AbiSession) Exchange(_amount *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.Exchange(&_Abi.TransactOpts, _amount)
}

// Exchange is a paid mutator transaction binding the contract method 0x11d92a08.
//
// Solidity: function Exchange(uint256 _amount) returns()
func (_Abi *AbiTransactorSession) Exchange(_amount *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.Exchange(&_Abi.TransactOpts, _amount)
}

// GetInterest is a paid mutator transaction binding the contract method 0x2c173950.
//
// Solidity: function GetInterest(bytes32 _hash) returns()
func (_Abi *AbiTransactor) GetInterest(opts *bind.TransactOpts, _hash [32]byte) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "GetInterest", _hash)
}

// GetInterest is a paid mutator transaction binding the contract method 0x2c173950.
//
// Solidity: function GetInterest(bytes32 _hash) returns()
func (_Abi *AbiSession) GetInterest(_hash [32]byte) (*types.Transaction, error) {
	return _Abi.Contract.GetInterest(&_Abi.TransactOpts, _hash)
}

// GetInterest is a paid mutator transaction binding the contract method 0x2c173950.
//
// Solidity: function GetInterest(bytes32 _hash) returns()
func (_Abi *AbiTransactorSession) GetInterest(_hash [32]byte) (*types.Transaction, error) {
	return _Abi.Contract.GetInterest(&_Abi.TransactOpts, _hash)
}

// Initialization is a paid mutator transaction binding the contract method 0x860c7f94.
//
// Solidity: function Initialization(string _name) returns()
func (_Abi *AbiTransactor) Initialization(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "Initialization", _name)
}

// Initialization is a paid mutator transaction binding the contract method 0x860c7f94.
//
// Solidity: function Initialization(string _name) returns()
func (_Abi *AbiSession) Initialization(_name string) (*types.Transaction, error) {
	return _Abi.Contract.Initialization(&_Abi.TransactOpts, _name)
}

// Initialization is a paid mutator transaction binding the contract method 0x860c7f94.
//
// Solidity: function Initialization(string _name) returns()
func (_Abi *AbiTransactorSession) Initialization(_name string) (*types.Transaction, error) {
	return _Abi.Contract.Initialization(&_Abi.TransactOpts, _name)
}

// ReleaseNeutralize is a paid mutator transaction binding the contract method 0xc9cc5ee0.
//
// Solidity: function ReleaseNeutralize(uint256 _year, uint256 _neutralizeTotal) returns()
func (_Abi *AbiTransactor) ReleaseNeutralize(opts *bind.TransactOpts, _year *big.Int, _neutralizeTotal *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "ReleaseNeutralize", _year, _neutralizeTotal)
}

// ReleaseNeutralize is a paid mutator transaction binding the contract method 0xc9cc5ee0.
//
// Solidity: function ReleaseNeutralize(uint256 _year, uint256 _neutralizeTotal) returns()
func (_Abi *AbiSession) ReleaseNeutralize(_year *big.Int, _neutralizeTotal *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.ReleaseNeutralize(&_Abi.TransactOpts, _year, _neutralizeTotal)
}

// ReleaseNeutralize is a paid mutator transaction binding the contract method 0xc9cc5ee0.
//
// Solidity: function ReleaseNeutralize(uint256 _year, uint256 _neutralizeTotal) returns()
func (_Abi *AbiTransactorSession) ReleaseNeutralize(_year *big.Int, _neutralizeTotal *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.ReleaseNeutralize(&_Abi.TransactOpts, _year, _neutralizeTotal)
}

// Report is a paid mutator transaction binding the contract method 0x039d1422.
//
// Solidity: function Report(address _ip, string _reason) returns()
func (_Abi *AbiTransactor) Report(opts *bind.TransactOpts, _ip common.Address, _reason string) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "Report", _ip, _reason)
}

// Report is a paid mutator transaction binding the contract method 0x039d1422.
//
// Solidity: function Report(address _ip, string _reason) returns()
func (_Abi *AbiSession) Report(_ip common.Address, _reason string) (*types.Transaction, error) {
	return _Abi.Contract.Report(&_Abi.TransactOpts, _ip, _reason)
}

// Report is a paid mutator transaction binding the contract method 0x039d1422.
//
// Solidity: function Report(address _ip, string _reason) returns()
func (_Abi *AbiTransactorSession) Report(_ip common.Address, _reason string) (*types.Transaction, error) {
	return _Abi.Contract.Report(&_Abi.TransactOpts, _ip, _reason)
}

// SubERC is a paid mutator transaction binding the contract method 0xadf39022.
//
// Solidity: function SubERC(address ip, uint256 amount) returns()
func (_Abi *AbiTransactor) SubERC(opts *bind.TransactOpts, ip common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "SubERC", ip, amount)
}

// SubERC is a paid mutator transaction binding the contract method 0xadf39022.
//
// Solidity: function SubERC(address ip, uint256 amount) returns()
func (_Abi *AbiSession) SubERC(ip common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.SubERC(&_Abi.TransactOpts, ip, amount)
}

// SubERC is a paid mutator transaction binding the contract method 0xadf39022.
//
// Solidity: function SubERC(address ip, uint256 amount) returns()
func (_Abi *AbiTransactorSession) SubERC(ip common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.SubERC(&_Abi.TransactOpts, ip, amount)
}

// TakeAPart is a paid mutator transaction binding the contract method 0x1226d16f.
//
// Solidity: function TakeAPart(string _fun, uint256 _NeutralizedAmount, uint256 _personalAmount) returns()
func (_Abi *AbiTransactor) TakeAPart(opts *bind.TransactOpts, _fun string, _NeutralizedAmount *big.Int, _personalAmount *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "TakeAPart", _fun, _NeutralizedAmount, _personalAmount)
}

// TakeAPart is a paid mutator transaction binding the contract method 0x1226d16f.
//
// Solidity: function TakeAPart(string _fun, uint256 _NeutralizedAmount, uint256 _personalAmount) returns()
func (_Abi *AbiSession) TakeAPart(_fun string, _NeutralizedAmount *big.Int, _personalAmount *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.TakeAPart(&_Abi.TransactOpts, _fun, _NeutralizedAmount, _personalAmount)
}

// TakeAPart is a paid mutator transaction binding the contract method 0x1226d16f.
//
// Solidity: function TakeAPart(string _fun, uint256 _NeutralizedAmount, uint256 _personalAmount) returns()
func (_Abi *AbiTransactorSession) TakeAPart(_fun string, _NeutralizedAmount *big.Int, _personalAmount *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.TakeAPart(&_Abi.TransactOpts, _fun, _NeutralizedAmount, _personalAmount)
}

// TakeOutBank is a paid mutator transaction binding the contract method 0x1cf8c5b6.
//
// Solidity: function TakeOutBank(uint256 _amount, string _operation) returns()
func (_Abi *AbiTransactor) TakeOutBank(opts *bind.TransactOpts, _amount *big.Int, _operation string) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "TakeOutBank", _amount, _operation)
}

// TakeOutBank is a paid mutator transaction binding the contract method 0x1cf8c5b6.
//
// Solidity: function TakeOutBank(uint256 _amount, string _operation) returns()
func (_Abi *AbiSession) TakeOutBank(_amount *big.Int, _operation string) (*types.Transaction, error) {
	return _Abi.Contract.TakeOutBank(&_Abi.TransactOpts, _amount, _operation)
}

// TakeOutBank is a paid mutator transaction binding the contract method 0x1cf8c5b6.
//
// Solidity: function TakeOutBank(uint256 _amount, string _operation) returns()
func (_Abi *AbiTransactorSession) TakeOutBank(_amount *big.Int, _operation string) (*types.Transaction, error) {
	return _Abi.Contract.TakeOutBank(&_Abi.TransactOpts, _amount, _operation)
}

// Voting is a paid mutator transaction binding the contract method 0x14ad08ba.
//
// Solidity: function Voting(uint256 pId) returns()
func (_Abi *AbiTransactor) Voting(opts *bind.TransactOpts, pId *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "Voting", pId)
}

// Voting is a paid mutator transaction binding the contract method 0x14ad08ba.
//
// Solidity: function Voting(uint256 pId) returns()
func (_Abi *AbiSession) Voting(pId *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.Voting(&_Abi.TransactOpts, pId)
}

// Voting is a paid mutator transaction binding the contract method 0x14ad08ba.
//
// Solidity: function Voting(uint256 pId) returns()
func (_Abi *AbiTransactorSession) Voting(pId *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.Voting(&_Abi.TransactOpts, pId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Abi *AbiTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Abi *AbiSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.Approve(&_Abi.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Abi *AbiTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.Approve(&_Abi.TransactOpts, spender, amount)
}

// CreateVote is a paid mutator transaction binding the contract method 0x96c29a31.
//
// Solidity: function createVote(address ip, string _pName) returns()
func (_Abi *AbiTransactor) CreateVote(opts *bind.TransactOpts, ip common.Address, _pName string) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "createVote", ip, _pName)
}

// CreateVote is a paid mutator transaction binding the contract method 0x96c29a31.
//
// Solidity: function createVote(address ip, string _pName) returns()
func (_Abi *AbiSession) CreateVote(ip common.Address, _pName string) (*types.Transaction, error) {
	return _Abi.Contract.CreateVote(&_Abi.TransactOpts, ip, _pName)
}

// CreateVote is a paid mutator transaction binding the contract method 0x96c29a31.
//
// Solidity: function createVote(address ip, string _pName) returns()
func (_Abi *AbiTransactorSession) CreateVote(ip common.Address, _pName string) (*types.Transaction, error) {
	return _Abi.Contract.CreateVote(&_Abi.TransactOpts, ip, _pName)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Abi *AbiTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Abi *AbiSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.DecreaseAllowance(&_Abi.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Abi *AbiTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.DecreaseAllowance(&_Abi.TransactOpts, spender, subtractedValue)
}

// Finish is a paid mutator transaction binding the contract method 0x10290dd1.
//
// Solidity: function finish(uint256 pId, uint256 _reward, uint256 MASSID) returns()
func (_Abi *AbiTransactor) Finish(opts *bind.TransactOpts, pId *big.Int, _reward *big.Int, MASSID *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "finish", pId, _reward, MASSID)
}

// Finish is a paid mutator transaction binding the contract method 0x10290dd1.
//
// Solidity: function finish(uint256 pId, uint256 _reward, uint256 MASSID) returns()
func (_Abi *AbiSession) Finish(pId *big.Int, _reward *big.Int, MASSID *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.Finish(&_Abi.TransactOpts, pId, _reward, MASSID)
}

// Finish is a paid mutator transaction binding the contract method 0x10290dd1.
//
// Solidity: function finish(uint256 pId, uint256 _reward, uint256 MASSID) returns()
func (_Abi *AbiTransactorSession) Finish(pId *big.Int, _reward *big.Int, MASSID *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.Finish(&_Abi.TransactOpts, pId, _reward, MASSID)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Abi *AbiTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Abi *AbiSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.IncreaseAllowance(&_Abi.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Abi *AbiTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.IncreaseAllowance(&_Abi.TransactOpts, spender, addedValue)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Abi *AbiTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Abi *AbiSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.TransferFrom(&_Abi.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Abi *AbiTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.TransferFrom(&_Abi.TransactOpts, sender, recipient, amount)
}

// Transfes is a paid mutator transaction binding the contract method 0x853aa34f.
//
// Solidity: function transfes(address recipient, uint256 amount) returns(bool)
func (_Abi *AbiTransactor) Transfes(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "transfes", recipient, amount)
}

// Transfes is a paid mutator transaction binding the contract method 0x853aa34f.
//
// Solidity: function transfes(address recipient, uint256 amount) returns(bool)
func (_Abi *AbiSession) Transfes(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.Transfes(&_Abi.TransactOpts, recipient, amount)
}

// Transfes is a paid mutator transaction binding the contract method 0x853aa34f.
//
// Solidity: function transfes(address recipient, uint256 amount) returns(bool)
func (_Abi *AbiTransactorSession) Transfes(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.Transfes(&_Abi.TransactOpts, recipient, amount)
}

// AbiADDBANKIterator is returned from FilterADDBANK and is used to iterate over the raw logs and unpacked data for ADDBANK events raised by the Abi contract.
type AbiADDBANKIterator struct {
	Event *AbiADDBANK // Event containing the contract specifics and raw log

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
func (it *AbiADDBANKIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiADDBANK)
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
		it.Event = new(AbiADDBANK)
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
func (it *AbiADDBANKIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiADDBANKIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiADDBANK represents a ADDBANK event raised by the Abi contract.
type AbiADDBANK struct {
	Ip   common.Address
	Hash [32]byte
	Time [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterADDBANK is a free log retrieval operation binding the contract event 0x7beeef7388159dfa39729d6f420a34aba3e238b62c496c0aa3342a86ab63d7e8.
//
// Solidity: event ADDBANK(address ip, bytes32 hash, bytes32 time)
func (_Abi *AbiFilterer) FilterADDBANK(opts *bind.FilterOpts) (*AbiADDBANKIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "ADDBANK")
	if err != nil {
		return nil, err
	}
	return &AbiADDBANKIterator{contract: _Abi.contract, event: "ADDBANK", logs: logs, sub: sub}, nil
}

// WatchADDBANK is a free log subscription operation binding the contract event 0x7beeef7388159dfa39729d6f420a34aba3e238b62c496c0aa3342a86ab63d7e8.
//
// Solidity: event ADDBANK(address ip, bytes32 hash, bytes32 time)
func (_Abi *AbiFilterer) WatchADDBANK(opts *bind.WatchOpts, sink chan<- *AbiADDBANK) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "ADDBANK")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiADDBANK)
				if err := _Abi.contract.UnpackLog(event, "ADDBANK", log); err != nil {
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

// ParseADDBANK is a log parse operation binding the contract event 0x7beeef7388159dfa39729d6f420a34aba3e238b62c496c0aa3342a86ab63d7e8.
//
// Solidity: event ADDBANK(address ip, bytes32 hash, bytes32 time)
func (_Abi *AbiFilterer) ParseADDBANK(log types.Log) (*AbiADDBANK, error) {
	event := new(AbiADDBANK)
	if err := _Abi.contract.UnpackLog(event, "ADDBANK", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AbiApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Abi contract.
type AbiApprovalIterator struct {
	Event *AbiApproval // Event containing the contract specifics and raw log

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
func (it *AbiApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiApproval)
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
		it.Event = new(AbiApproval)
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
func (it *AbiApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiApproval represents a Approval event raised by the Abi contract.
type AbiApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Abi *AbiFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*AbiApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &AbiApprovalIterator{contract: _Abi.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Abi *AbiFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AbiApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiApproval)
				if err := _Abi.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Abi *AbiFilterer) ParseApproval(log types.Log) (*AbiApproval, error) {
	event := new(AbiApproval)
	if err := _Abi.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AbiTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Abi contract.
type AbiTransferIterator struct {
	Event *AbiTransfer // Event containing the contract specifics and raw log

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
func (it *AbiTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiTransfer)
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
		it.Event = new(AbiTransfer)
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
func (it *AbiTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiTransfer represents a Transfer event raised by the Abi contract.
type AbiTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Abi *AbiFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AbiTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AbiTransferIterator{contract: _Abi.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Abi *AbiFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AbiTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiTransfer)
				if err := _Abi.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Abi *AbiFilterer) ParseTransfer(log types.Log) (*AbiTransfer, error) {
	event := new(AbiTransfer)
	if err := _Abi.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AbiVOTEIterator is returned from FilterVOTE and is used to iterate over the raw logs and unpacked data for VOTE events raised by the Abi contract.
type AbiVOTEIterator struct {
	Event *AbiVOTE // Event containing the contract specifics and raw log

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
func (it *AbiVOTEIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiVOTE)
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
		it.Event = new(AbiVOTE)
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
func (it *AbiVOTEIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiVOTEIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiVOTE represents a VOTE event raised by the Abi contract.
type AbiVOTE struct {
	Ip         common.Address
	ProposalId *big.Int
	LimitTime  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVOTE is a free log retrieval operation binding the contract event 0x37d18319ced642c4dc26af530833449cfdc6fdc8929c0db00cdf772f7c140ae7.
//
// Solidity: event VOTE(address ip, uint256 _proposalId, uint256 _limitTime)
func (_Abi *AbiFilterer) FilterVOTE(opts *bind.FilterOpts) (*AbiVOTEIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "VOTE")
	if err != nil {
		return nil, err
	}
	return &AbiVOTEIterator{contract: _Abi.contract, event: "VOTE", logs: logs, sub: sub}, nil
}

// WatchVOTE is a free log subscription operation binding the contract event 0x37d18319ced642c4dc26af530833449cfdc6fdc8929c0db00cdf772f7c140ae7.
//
// Solidity: event VOTE(address ip, uint256 _proposalId, uint256 _limitTime)
func (_Abi *AbiFilterer) WatchVOTE(opts *bind.WatchOpts, sink chan<- *AbiVOTE) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "VOTE")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiVOTE)
				if err := _Abi.contract.UnpackLog(event, "VOTE", log); err != nil {
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

// ParseVOTE is a log parse operation binding the contract event 0x37d18319ced642c4dc26af530833449cfdc6fdc8929c0db00cdf772f7c140ae7.
//
// Solidity: event VOTE(address ip, uint256 _proposalId, uint256 _limitTime)
func (_Abi *AbiFilterer) ParseVOTE(log types.Log) (*AbiVOTE, error) {
	event := new(AbiVOTE)
	if err := _Abi.contract.UnpackLog(event, "VOTE", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AbiVoteEvtIterator is returned from FilterVoteEvt and is used to iterate over the raw logs and unpacked data for VoteEvt events raised by the Abi contract.
type AbiVoteEvtIterator struct {
	Event *AbiVoteEvt // Event containing the contract specifics and raw log

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
func (it *AbiVoteEvtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiVoteEvt)
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
		it.Event = new(AbiVoteEvt)
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
func (it *AbiVoteEvtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiVoteEvtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiVoteEvt represents a VoteEvt event raised by the Abi contract.
type AbiVoteEvt struct {
	EventType common.Hash
	Voter     common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVoteEvt is a free log retrieval operation binding the contract event 0x6e1d6794e13b69f31305f22ea91ab03481b376fdf800c35bb36b6b582c25f193.
//
// Solidity: event VoteEvt(string indexed eventType, address _voter, uint256 timestamp)
func (_Abi *AbiFilterer) FilterVoteEvt(opts *bind.FilterOpts, eventType []string) (*AbiVoteEvtIterator, error) {

	var eventTypeRule []interface{}
	for _, eventTypeItem := range eventType {
		eventTypeRule = append(eventTypeRule, eventTypeItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "VoteEvt", eventTypeRule)
	if err != nil {
		return nil, err
	}
	return &AbiVoteEvtIterator{contract: _Abi.contract, event: "VoteEvt", logs: logs, sub: sub}, nil
}

// WatchVoteEvt is a free log subscription operation binding the contract event 0x6e1d6794e13b69f31305f22ea91ab03481b376fdf800c35bb36b6b582c25f193.
//
// Solidity: event VoteEvt(string indexed eventType, address _voter, uint256 timestamp)
func (_Abi *AbiFilterer) WatchVoteEvt(opts *bind.WatchOpts, sink chan<- *AbiVoteEvt, eventType []string) (event.Subscription, error) {

	var eventTypeRule []interface{}
	for _, eventTypeItem := range eventType {
		eventTypeRule = append(eventTypeRule, eventTypeItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "VoteEvt", eventTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiVoteEvt)
				if err := _Abi.contract.UnpackLog(event, "VoteEvt", log); err != nil {
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

// ParseVoteEvt is a log parse operation binding the contract event 0x6e1d6794e13b69f31305f22ea91ab03481b376fdf800c35bb36b6b582c25f193.
//
// Solidity: event VoteEvt(string indexed eventType, address _voter, uint256 timestamp)
func (_Abi *AbiFilterer) ParseVoteEvt(log types.Log) (*AbiVoteEvt, error) {
	event := new(AbiVoteEvt)
	if err := _Abi.contract.UnpackLog(event, "VoteEvt", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AbiInitializationEVTIterator is returned from FilterInitializationEVT and is used to iterate over the raw logs and unpacked data for InitializationEVT events raised by the Abi contract.
type AbiInitializationEVTIterator struct {
	Event *AbiInitializationEVT // Event containing the contract specifics and raw log

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
func (it *AbiInitializationEVTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiInitializationEVT)
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
		it.Event = new(AbiInitializationEVT)
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
func (it *AbiInitializationEVTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiInitializationEVTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiInitializationEVT represents a InitializationEVT event raised by the Abi contract.
type AbiInitializationEVT struct {
	Ip   common.Address
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterInitializationEVT is a free log retrieval operation binding the contract event 0xab31f6da07e7e79e1a8e2f9634461f3a58980269e23b501f490716b791478a62.
//
// Solidity: event initializationEVT(address ip, uint256 time)
func (_Abi *AbiFilterer) FilterInitializationEVT(opts *bind.FilterOpts) (*AbiInitializationEVTIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "initializationEVT")
	if err != nil {
		return nil, err
	}
	return &AbiInitializationEVTIterator{contract: _Abi.contract, event: "initializationEVT", logs: logs, sub: sub}, nil
}

// WatchInitializationEVT is a free log subscription operation binding the contract event 0xab31f6da07e7e79e1a8e2f9634461f3a58980269e23b501f490716b791478a62.
//
// Solidity: event initializationEVT(address ip, uint256 time)
func (_Abi *AbiFilterer) WatchInitializationEVT(opts *bind.WatchOpts, sink chan<- *AbiInitializationEVT) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "initializationEVT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiInitializationEVT)
				if err := _Abi.contract.UnpackLog(event, "initializationEVT", log); err != nil {
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

// ParseInitializationEVT is a log parse operation binding the contract event 0xab31f6da07e7e79e1a8e2f9634461f3a58980269e23b501f490716b791478a62.
//
// Solidity: event initializationEVT(address ip, uint256 time)
func (_Abi *AbiFilterer) ParseInitializationEVT(log types.Log) (*AbiInitializationEVT, error) {
	event := new(AbiInitializationEVT)
	if err := _Abi.contract.UnpackLog(event, "initializationEVT", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AbiTakePartIterator is returned from FilterTakePart and is used to iterate over the raw logs and unpacked data for TakePart events raised by the Abi contract.
type AbiTakePartIterator struct {
	Event *AbiTakePart // Event containing the contract specifics and raw log

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
func (it *AbiTakePartIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiTakePart)
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
		it.Event = new(AbiTakePart)
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
func (it *AbiTakePartIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiTakePartIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiTakePart represents a TakePart event raised by the Abi contract.
type AbiTakePart struct {
	Ip   common.Address
	Id   *big.Int
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTakePart is a free log retrieval operation binding the contract event 0x328aecc2cf9bbc8f7b548423be007d35e29000d267b5b90f36b76d6e196715b9.
//
// Solidity: event takePart(address ip, uint256 id, uint256 time)
func (_Abi *AbiFilterer) FilterTakePart(opts *bind.FilterOpts) (*AbiTakePartIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "takePart")
	if err != nil {
		return nil, err
	}
	return &AbiTakePartIterator{contract: _Abi.contract, event: "takePart", logs: logs, sub: sub}, nil
}

// WatchTakePart is a free log subscription operation binding the contract event 0x328aecc2cf9bbc8f7b548423be007d35e29000d267b5b90f36b76d6e196715b9.
//
// Solidity: event takePart(address ip, uint256 id, uint256 time)
func (_Abi *AbiFilterer) WatchTakePart(opts *bind.WatchOpts, sink chan<- *AbiTakePart) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "takePart")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiTakePart)
				if err := _Abi.contract.UnpackLog(event, "takePart", log); err != nil {
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

// ParseTakePart is a log parse operation binding the contract event 0x328aecc2cf9bbc8f7b548423be007d35e29000d267b5b90f36b76d6e196715b9.
//
// Solidity: event takePart(address ip, uint256 id, uint256 time)
func (_Abi *AbiFilterer) ParseTakePart(log types.Log) (*AbiTakePart, error) {
	event := new(AbiTakePart)
	if err := _Abi.contract.UnpackLog(event, "takePart", log); err != nil {
		return nil, err
	}
	return event, nil
}
