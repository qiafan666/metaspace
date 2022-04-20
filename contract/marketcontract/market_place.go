// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package marketcontract

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

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"MatchTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"feeToAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_paymentErc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_saltNonce\",\"type\":\"uint256\"}],\"name\":\"getMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_paymentErc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_saltNonce\",\"type\":\"uint256\"}],\"name\":\"getMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftaddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"saltnonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"ignoreSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[3]\",\"name\":\"addresses\",\"type\":\"address[3]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"saltnonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"matchTransaction\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[3]\",\"name\":\"addresses\",\"type\":\"address[3]\"},{\"internalType\":\"uint256[3]\",\"name\":\"values\",\"type\":\"uint256[3]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"matchTransaction\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"paymentTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_removedPaymentTokens\",\"type\":\"address[]\"}],\"name\":\"removePaymentTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToAddress\",\"type\":\"address\"}],\"name\":\"setFeeToAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_paymentTokens\",\"type\":\"address[]\"}],\"name\":\"setPaymentTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_transactionFeePerMillion\",\"type\":\"uint256\"}],\"name\":\"setTransactionFeePerMillion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transactionFeePerMillion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"usedSignatures\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// FeeToAddress is a free data retrieval call binding the contract method 0x083d80f9.
//
// Solidity: function feeToAddress() view returns(address)
func (_Contracts *ContractsCaller) FeeToAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "feeToAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeToAddress is a free data retrieval call binding the contract method 0x083d80f9.
//
// Solidity: function feeToAddress() view returns(address)
func (_Contracts *ContractsSession) FeeToAddress() (common.Address, error) {
	return _Contracts.Contract.FeeToAddress(&_Contracts.CallOpts)
}

// FeeToAddress is a free data retrieval call binding the contract method 0x083d80f9.
//
// Solidity: function feeToAddress() view returns(address)
func (_Contracts *ContractsCallerSession) FeeToAddress() (common.Address, error) {
	return _Contracts.Contract.FeeToAddress(&_Contracts.CallOpts)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x1e5192e5.
//
// Solidity: function getMessageHash(address _nftAddress, uint256 _tokenId, address _paymentErc20, uint256 _price, uint256 _saltNonce) pure returns(bytes32)
func (_Contracts *ContractsCaller) GetMessageHash(opts *bind.CallOpts, _nftAddress common.Address, _tokenId *big.Int, _paymentErc20 common.Address, _price *big.Int, _saltNonce *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getMessageHash", _nftAddress, _tokenId, _paymentErc20, _price, _saltNonce)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMessageHash is a free data retrieval call binding the contract method 0x1e5192e5.
//
// Solidity: function getMessageHash(address _nftAddress, uint256 _tokenId, address _paymentErc20, uint256 _price, uint256 _saltNonce) pure returns(bytes32)
func (_Contracts *ContractsSession) GetMessageHash(_nftAddress common.Address, _tokenId *big.Int, _paymentErc20 common.Address, _price *big.Int, _saltNonce *big.Int) ([32]byte, error) {
	return _Contracts.Contract.GetMessageHash(&_Contracts.CallOpts, _nftAddress, _tokenId, _paymentErc20, _price, _saltNonce)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x1e5192e5.
//
// Solidity: function getMessageHash(address _nftAddress, uint256 _tokenId, address _paymentErc20, uint256 _price, uint256 _saltNonce) pure returns(bytes32)
func (_Contracts *ContractsCallerSession) GetMessageHash(_nftAddress common.Address, _tokenId *big.Int, _paymentErc20 common.Address, _price *big.Int, _saltNonce *big.Int) ([32]byte, error) {
	return _Contracts.Contract.GetMessageHash(&_Contracts.CallOpts, _nftAddress, _tokenId, _paymentErc20, _price, _saltNonce)
}

// GetMessageHash0 is a free data retrieval call binding the contract method 0x9d76b8ad.
//
// Solidity: function getMessageHash(address _nftAddress, uint256[] _tokenIds, address _paymentErc20, uint256 _price, uint256 _saltNonce) pure returns(bytes32)
func (_Contracts *ContractsCaller) GetMessageHash0(opts *bind.CallOpts, _nftAddress common.Address, _tokenIds []*big.Int, _paymentErc20 common.Address, _price *big.Int, _saltNonce *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getMessageHash0", _nftAddress, _tokenIds, _paymentErc20, _price, _saltNonce)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMessageHash0 is a free data retrieval call binding the contract method 0x9d76b8ad.
//
// Solidity: function getMessageHash(address _nftAddress, uint256[] _tokenIds, address _paymentErc20, uint256 _price, uint256 _saltNonce) pure returns(bytes32)
func (_Contracts *ContractsSession) GetMessageHash0(_nftAddress common.Address, _tokenIds []*big.Int, _paymentErc20 common.Address, _price *big.Int, _saltNonce *big.Int) ([32]byte, error) {
	return _Contracts.Contract.GetMessageHash0(&_Contracts.CallOpts, _nftAddress, _tokenIds, _paymentErc20, _price, _saltNonce)
}

// GetMessageHash0 is a free data retrieval call binding the contract method 0x9d76b8ad.
//
// Solidity: function getMessageHash(address _nftAddress, uint256[] _tokenIds, address _paymentErc20, uint256 _price, uint256 _saltNonce) pure returns(bytes32)
func (_Contracts *ContractsCallerSession) GetMessageHash0(_nftAddress common.Address, _tokenIds []*big.Int, _paymentErc20 common.Address, _price *big.Int, _saltNonce *big.Int) ([32]byte, error) {
	return _Contracts.Contract.GetMessageHash0(&_Contracts.CallOpts, _nftAddress, _tokenIds, _paymentErc20, _price, _saltNonce)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCallerSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// PaymentTokens is a free data retrieval call binding the contract method 0xc3b88b42.
//
// Solidity: function paymentTokens(address ) view returns(bool)
func (_Contracts *ContractsCaller) PaymentTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "paymentTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PaymentTokens is a free data retrieval call binding the contract method 0xc3b88b42.
//
// Solidity: function paymentTokens(address ) view returns(bool)
func (_Contracts *ContractsSession) PaymentTokens(arg0 common.Address) (bool, error) {
	return _Contracts.Contract.PaymentTokens(&_Contracts.CallOpts, arg0)
}

// PaymentTokens is a free data retrieval call binding the contract method 0xc3b88b42.
//
// Solidity: function paymentTokens(address ) view returns(bool)
func (_Contracts *ContractsCallerSession) PaymentTokens(arg0 common.Address) (bool, error) {
	return _Contracts.Contract.PaymentTokens(&_Contracts.CallOpts, arg0)
}

// TransactionFeePerMillion is a free data retrieval call binding the contract method 0x2b08f1af.
//
// Solidity: function transactionFeePerMillion() view returns(uint256)
func (_Contracts *ContractsCaller) TransactionFeePerMillion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "transactionFeePerMillion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransactionFeePerMillion is a free data retrieval call binding the contract method 0x2b08f1af.
//
// Solidity: function transactionFeePerMillion() view returns(uint256)
func (_Contracts *ContractsSession) TransactionFeePerMillion() (*big.Int, error) {
	return _Contracts.Contract.TransactionFeePerMillion(&_Contracts.CallOpts)
}

// TransactionFeePerMillion is a free data retrieval call binding the contract method 0x2b08f1af.
//
// Solidity: function transactionFeePerMillion() view returns(uint256)
func (_Contracts *ContractsCallerSession) TransactionFeePerMillion() (*big.Int, error) {
	return _Contracts.Contract.TransactionFeePerMillion(&_Contracts.CallOpts)
}

// UsedSignatures is a free data retrieval call binding the contract method 0xe949580e.
//
// Solidity: function usedSignatures(bytes ) view returns(bool)
func (_Contracts *ContractsCaller) UsedSignatures(opts *bind.CallOpts, arg0 []byte) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "usedSignatures", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UsedSignatures is a free data retrieval call binding the contract method 0xe949580e.
//
// Solidity: function usedSignatures(bytes ) view returns(bool)
func (_Contracts *ContractsSession) UsedSignatures(arg0 []byte) (bool, error) {
	return _Contracts.Contract.UsedSignatures(&_Contracts.CallOpts, arg0)
}

// UsedSignatures is a free data retrieval call binding the contract method 0xe949580e.
//
// Solidity: function usedSignatures(bytes ) view returns(bool)
func (_Contracts *ContractsCallerSession) UsedSignatures(arg0 []byte) (bool, error) {
	return _Contracts.Contract.UsedSignatures(&_Contracts.CallOpts, arg0)
}

// IgnoreSignature is a paid mutator transaction binding the contract method 0x4b290720.
//
// Solidity: function ignoreSignature(address nftaddress, address paymentToken, uint256[] tokenIds, uint256 price, uint256 saltnonce, bytes signature) returns()
func (_Contracts *ContractsTransactor) IgnoreSignature(opts *bind.TransactOpts, nftaddress common.Address, paymentToken common.Address, tokenIds []*big.Int, price *big.Int, saltnonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "ignoreSignature", nftaddress, paymentToken, tokenIds, price, saltnonce, signature)
}

// IgnoreSignature is a paid mutator transaction binding the contract method 0x4b290720.
//
// Solidity: function ignoreSignature(address nftaddress, address paymentToken, uint256[] tokenIds, uint256 price, uint256 saltnonce, bytes signature) returns()
func (_Contracts *ContractsSession) IgnoreSignature(nftaddress common.Address, paymentToken common.Address, tokenIds []*big.Int, price *big.Int, saltnonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _Contracts.Contract.IgnoreSignature(&_Contracts.TransactOpts, nftaddress, paymentToken, tokenIds, price, saltnonce, signature)
}

// IgnoreSignature is a paid mutator transaction binding the contract method 0x4b290720.
//
// Solidity: function ignoreSignature(address nftaddress, address paymentToken, uint256[] tokenIds, uint256 price, uint256 saltnonce, bytes signature) returns()
func (_Contracts *ContractsTransactorSession) IgnoreSignature(nftaddress common.Address, paymentToken common.Address, tokenIds []*big.Int, price *big.Int, saltnonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _Contracts.Contract.IgnoreSignature(&_Contracts.TransactOpts, nftaddress, paymentToken, tokenIds, price, saltnonce, signature)
}

// MatchTransaction is a paid mutator transaction binding the contract method 0x0bde414e.
//
// Solidity: function matchTransaction(address[3] addresses, uint256[] tokenIds, uint256 price, uint256 saltnonce, bytes signature) returns(bool)
func (_Contracts *ContractsTransactor) MatchTransaction(opts *bind.TransactOpts, addresses [3]common.Address, tokenIds []*big.Int, price *big.Int, saltnonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "matchTransaction", addresses, tokenIds, price, saltnonce, signature)
}

// MatchTransaction is a paid mutator transaction binding the contract method 0x0bde414e.
//
// Solidity: function matchTransaction(address[3] addresses, uint256[] tokenIds, uint256 price, uint256 saltnonce, bytes signature) returns(bool)
func (_Contracts *ContractsSession) MatchTransaction(addresses [3]common.Address, tokenIds []*big.Int, price *big.Int, saltnonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _Contracts.Contract.MatchTransaction(&_Contracts.TransactOpts, addresses, tokenIds, price, saltnonce, signature)
}

// MatchTransaction is a paid mutator transaction binding the contract method 0x0bde414e.
//
// Solidity: function matchTransaction(address[3] addresses, uint256[] tokenIds, uint256 price, uint256 saltnonce, bytes signature) returns(bool)
func (_Contracts *ContractsTransactorSession) MatchTransaction(addresses [3]common.Address, tokenIds []*big.Int, price *big.Int, saltnonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _Contracts.Contract.MatchTransaction(&_Contracts.TransactOpts, addresses, tokenIds, price, saltnonce, signature)
}

// MatchTransaction0 is a paid mutator transaction binding the contract method 0xe8e8e872.
//
// Solidity: function matchTransaction(address[3] addresses, uint256[3] values, bytes signature) returns(bool)
func (_Contracts *ContractsTransactor) MatchTransaction0(opts *bind.TransactOpts, addresses [3]common.Address, values [3]*big.Int, signature []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "matchTransaction0", addresses, values, signature)
}

// MatchTransaction0 is a paid mutator transaction binding the contract method 0xe8e8e872.
//
// Solidity: function matchTransaction(address[3] addresses, uint256[3] values, bytes signature) returns(bool)
func (_Contracts *ContractsSession) MatchTransaction0(addresses [3]common.Address, values [3]*big.Int, signature []byte) (*types.Transaction, error) {
	return _Contracts.Contract.MatchTransaction0(&_Contracts.TransactOpts, addresses, values, signature)
}

// MatchTransaction0 is a paid mutator transaction binding the contract method 0xe8e8e872.
//
// Solidity: function matchTransaction(address[3] addresses, uint256[3] values, bytes signature) returns(bool)
func (_Contracts *ContractsTransactorSession) MatchTransaction0(addresses [3]common.Address, values [3]*big.Int, signature []byte) (*types.Transaction, error) {
	return _Contracts.Contract.MatchTransaction0(&_Contracts.TransactOpts, addresses, values, signature)
}

// RemovePaymentTokens is a paid mutator transaction binding the contract method 0x64e60ef4.
//
// Solidity: function removePaymentTokens(address[] _removedPaymentTokens) returns()
func (_Contracts *ContractsTransactor) RemovePaymentTokens(opts *bind.TransactOpts, _removedPaymentTokens []common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "removePaymentTokens", _removedPaymentTokens)
}

// RemovePaymentTokens is a paid mutator transaction binding the contract method 0x64e60ef4.
//
// Solidity: function removePaymentTokens(address[] _removedPaymentTokens) returns()
func (_Contracts *ContractsSession) RemovePaymentTokens(_removedPaymentTokens []common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RemovePaymentTokens(&_Contracts.TransactOpts, _removedPaymentTokens)
}

// RemovePaymentTokens is a paid mutator transaction binding the contract method 0x64e60ef4.
//
// Solidity: function removePaymentTokens(address[] _removedPaymentTokens) returns()
func (_Contracts *ContractsTransactorSession) RemovePaymentTokens(_removedPaymentTokens []common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RemovePaymentTokens(&_Contracts.TransactOpts, _removedPaymentTokens)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contracts *ContractsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contracts *ContractsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contracts.Contract.RenounceOwnership(&_Contracts.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contracts *ContractsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contracts.Contract.RenounceOwnership(&_Contracts.TransactOpts)
}

// SetFeeToAddress is a paid mutator transaction binding the contract method 0x580bb9a5.
//
// Solidity: function setFeeToAddress(address _feeToAddress) returns()
func (_Contracts *ContractsTransactor) SetFeeToAddress(opts *bind.TransactOpts, _feeToAddress common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setFeeToAddress", _feeToAddress)
}

// SetFeeToAddress is a paid mutator transaction binding the contract method 0x580bb9a5.
//
// Solidity: function setFeeToAddress(address _feeToAddress) returns()
func (_Contracts *ContractsSession) SetFeeToAddress(_feeToAddress common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetFeeToAddress(&_Contracts.TransactOpts, _feeToAddress)
}

// SetFeeToAddress is a paid mutator transaction binding the contract method 0x580bb9a5.
//
// Solidity: function setFeeToAddress(address _feeToAddress) returns()
func (_Contracts *ContractsTransactorSession) SetFeeToAddress(_feeToAddress common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetFeeToAddress(&_Contracts.TransactOpts, _feeToAddress)
}

// SetPaymentTokens is a paid mutator transaction binding the contract method 0xb88dccac.
//
// Solidity: function setPaymentTokens(address[] _paymentTokens) returns()
func (_Contracts *ContractsTransactor) SetPaymentTokens(opts *bind.TransactOpts, _paymentTokens []common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setPaymentTokens", _paymentTokens)
}

// SetPaymentTokens is a paid mutator transaction binding the contract method 0xb88dccac.
//
// Solidity: function setPaymentTokens(address[] _paymentTokens) returns()
func (_Contracts *ContractsSession) SetPaymentTokens(_paymentTokens []common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetPaymentTokens(&_Contracts.TransactOpts, _paymentTokens)
}

// SetPaymentTokens is a paid mutator transaction binding the contract method 0xb88dccac.
//
// Solidity: function setPaymentTokens(address[] _paymentTokens) returns()
func (_Contracts *ContractsTransactorSession) SetPaymentTokens(_paymentTokens []common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetPaymentTokens(&_Contracts.TransactOpts, _paymentTokens)
}

// SetTransactionFeePerMillion is a paid mutator transaction binding the contract method 0x5952fee5.
//
// Solidity: function setTransactionFeePerMillion(uint256 _transactionFeePerMillion) returns()
func (_Contracts *ContractsTransactor) SetTransactionFeePerMillion(opts *bind.TransactOpts, _transactionFeePerMillion *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setTransactionFeePerMillion", _transactionFeePerMillion)
}

// SetTransactionFeePerMillion is a paid mutator transaction binding the contract method 0x5952fee5.
//
// Solidity: function setTransactionFeePerMillion(uint256 _transactionFeePerMillion) returns()
func (_Contracts *ContractsSession) SetTransactionFeePerMillion(_transactionFeePerMillion *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetTransactionFeePerMillion(&_Contracts.TransactOpts, _transactionFeePerMillion)
}

// SetTransactionFeePerMillion is a paid mutator transaction binding the contract method 0x5952fee5.
//
// Solidity: function setTransactionFeePerMillion(uint256 _transactionFeePerMillion) returns()
func (_Contracts *ContractsTransactorSession) SetTransactionFeePerMillion(_transactionFeePerMillion *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetTransactionFeePerMillion(&_Contracts.TransactOpts, _transactionFeePerMillion)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.TransferOwnership(&_Contracts.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.TransferOwnership(&_Contracts.TransactOpts, newOwner)
}

// ContractsMatchTransactionIterator is returned from FilterMatchTransaction and is used to iterate over the raw logs and unpacked data for MatchTransaction events raised by the Contracts contract.
type ContractsMatchTransactionIterator struct {
	Event *ContractsMatchTransaction // Event containing the contract specifics and raw log

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
func (it *ContractsMatchTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsMatchTransaction)
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
		it.Event = new(ContractsMatchTransaction)
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
func (it *ContractsMatchTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsMatchTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsMatchTransaction represents a MatchTransaction event raised by the Contracts contract.
type ContractsMatchTransaction struct {
	TokenIds        []*big.Int
	ContractAddress common.Address
	Price           *big.Int
	PaymentToken    common.Address
	Seller          common.Address
	Buyer           common.Address
	Fee             *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMatchTransaction is a free log retrieval operation binding the contract event 0x48f0b1d9e60d7fe2a9d34381733069d2f7ec8a3720c8f2432f48d615192846b1.
//
// Solidity: event MatchTransaction(uint256[] indexed tokenIds, address contractAddress, uint256 price, address paymentToken, address seller, address buyer, uint256 fee)
func (_Contracts *ContractsFilterer) FilterMatchTransaction(opts *bind.FilterOpts, tokenIds [][]*big.Int) (*ContractsMatchTransactionIterator, error) {

	var tokenIdsRule []interface{}
	for _, tokenIdsItem := range tokenIds {
		tokenIdsRule = append(tokenIdsRule, tokenIdsItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "MatchTransaction", tokenIdsRule)
	if err != nil {
		return nil, err
	}
	return &ContractsMatchTransactionIterator{contract: _Contracts.contract, event: "MatchTransaction", logs: logs, sub: sub}, nil
}

// WatchMatchTransaction is a free log subscription operation binding the contract event 0x48f0b1d9e60d7fe2a9d34381733069d2f7ec8a3720c8f2432f48d615192846b1.
//
// Solidity: event MatchTransaction(uint256[] indexed tokenIds, address contractAddress, uint256 price, address paymentToken, address seller, address buyer, uint256 fee)
func (_Contracts *ContractsFilterer) WatchMatchTransaction(opts *bind.WatchOpts, sink chan<- *ContractsMatchTransaction, tokenIds [][]*big.Int) (event.Subscription, error) {

	var tokenIdsRule []interface{}
	for _, tokenIdsItem := range tokenIds {
		tokenIdsRule = append(tokenIdsRule, tokenIdsItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "MatchTransaction", tokenIdsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsMatchTransaction)
				if err := _Contracts.contract.UnpackLog(event, "MatchTransaction", log); err != nil {
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

// ParseMatchTransaction is a log parse operation binding the contract event 0x48f0b1d9e60d7fe2a9d34381733069d2f7ec8a3720c8f2432f48d615192846b1.
//
// Solidity: event MatchTransaction(uint256[] indexed tokenIds, address contractAddress, uint256 price, address paymentToken, address seller, address buyer, uint256 fee)
func (_Contracts *ContractsFilterer) ParseMatchTransaction(log types.Log) (*ContractsMatchTransaction, error) {
	event := new(ContractsMatchTransaction)
	if err := _Contracts.contract.UnpackLog(event, "MatchTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contracts contract.
type ContractsOwnershipTransferredIterator struct {
	Event *ContractsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsOwnershipTransferred)
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
		it.Event = new(ContractsOwnershipTransferred)
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
func (it *ContractsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsOwnershipTransferred represents a OwnershipTransferred event raised by the Contracts contract.
type ContractsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contracts *ContractsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsOwnershipTransferredIterator{contract: _Contracts.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contracts *ContractsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsOwnershipTransferred)
				if err := _Contracts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contracts *ContractsFilterer) ParseOwnershipTransferred(log types.Log) (*ContractsOwnershipTransferred, error) {
	event := new(ContractsOwnershipTransferred)
	if err := _Contracts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
