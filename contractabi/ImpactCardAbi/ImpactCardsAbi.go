// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package impactcardsabi

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
	_ = abi.ConvertType
)

// ImpactcardsabiMetaData contains all meta data concerning the Impactcardsabi contract.
var ImpactcardsabiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"baseExtension\",\"type\":\"string\"}],\"name\":\"BaseExtensionChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"}],\"name\":\"BaseURIChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"limit\",\"type\":\"uint16\"}],\"name\":\"BulkBuyLimitSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"CardMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"CardsBatchMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"}],\"name\":\"ChangedContractPausedState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"MintPriceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[2]\",\"name\":\"payees\",\"type\":\"address[2]\"}],\"name\":\"PayeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"season\",\"type\":\"uint256\"}],\"name\":\"SeasonChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[2]\",\"name\":\"shares\",\"type\":\"uint256[2]\"}],\"name\":\"SharesSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_CARDS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_SUPPLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_baseExtension\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bulkBuyLimit\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentSeason\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getAccumulatedFunds\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getPayees\",\"outputs\":[{\"internalType\":\"address[2]\",\"name\":\"\",\"type\":\"address[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getShares\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"isMintable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenId\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"}],\"name\":\"mintBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextSeason\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"payeeIndex\",\"type\":\"uint8\"}],\"name\":\"release\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"releaseToAllPayees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"seasonCardCounts\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newExtension\",\"type\":\"string\"}],\"name\":\"setBaseExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newuri\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"limit\",\"type\":\"uint16\"}],\"name\":\"setBulkBuyLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setMintPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address[2]\",\"name\":\"payees\",\"type\":\"address[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"shares\",\"type\":\"uint256[2]\"}],\"name\":\"setPayees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"togglePaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"}],\"name\":\"totalReleasedToPayee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ImpactcardsabiABI is the input ABI used to generate the binding from.
// Deprecated: Use ImpactcardsabiMetaData.ABI instead.
var ImpactcardsabiABI = ImpactcardsabiMetaData.ABI

// Impactcardsabi is an auto generated Go binding around an Ethereum contract.
type Impactcardsabi struct {
	ImpactcardsabiCaller     // Read-only binding to the contract
	ImpactcardsabiTransactor // Write-only binding to the contract
	ImpactcardsabiFilterer   // Log filterer for contract events
}

// ImpactcardsabiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ImpactcardsabiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ImpactcardsabiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ImpactcardsabiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ImpactcardsabiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ImpactcardsabiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ImpactcardsabiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ImpactcardsabiSession struct {
	Contract     *Impactcardsabi   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ImpactcardsabiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ImpactcardsabiCallerSession struct {
	Contract *ImpactcardsabiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ImpactcardsabiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ImpactcardsabiTransactorSession struct {
	Contract     *ImpactcardsabiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ImpactcardsabiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ImpactcardsabiRaw struct {
	Contract *Impactcardsabi // Generic contract binding to access the raw methods on
}

// ImpactcardsabiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ImpactcardsabiCallerRaw struct {
	Contract *ImpactcardsabiCaller // Generic read-only contract binding to access the raw methods on
}

// ImpactcardsabiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ImpactcardsabiTransactorRaw struct {
	Contract *ImpactcardsabiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewImpactcardsabi creates a new instance of Impactcardsabi, bound to a specific deployed contract.
func NewImpactcardsabi(address common.Address, backend bind.ContractBackend) (*Impactcardsabi, error) {
	contract, err := bindImpactcardsabi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Impactcardsabi{ImpactcardsabiCaller: ImpactcardsabiCaller{contract: contract}, ImpactcardsabiTransactor: ImpactcardsabiTransactor{contract: contract}, ImpactcardsabiFilterer: ImpactcardsabiFilterer{contract: contract}}, nil
}

// NewImpactcardsabiCaller creates a new read-only instance of Impactcardsabi, bound to a specific deployed contract.
func NewImpactcardsabiCaller(address common.Address, caller bind.ContractCaller) (*ImpactcardsabiCaller, error) {
	contract, err := bindImpactcardsabi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ImpactcardsabiCaller{contract: contract}, nil
}

// NewImpactcardsabiTransactor creates a new write-only instance of Impactcardsabi, bound to a specific deployed contract.
func NewImpactcardsabiTransactor(address common.Address, transactor bind.ContractTransactor) (*ImpactcardsabiTransactor, error) {
	contract, err := bindImpactcardsabi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ImpactcardsabiTransactor{contract: contract}, nil
}

// NewImpactcardsabiFilterer creates a new log filterer instance of Impactcardsabi, bound to a specific deployed contract.
func NewImpactcardsabiFilterer(address common.Address, filterer bind.ContractFilterer) (*ImpactcardsabiFilterer, error) {
	contract, err := bindImpactcardsabi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ImpactcardsabiFilterer{contract: contract}, nil
}

// bindImpactcardsabi binds a generic wrapper to an already deployed contract.
func bindImpactcardsabi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ImpactcardsabiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Impactcardsabi *ImpactcardsabiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Impactcardsabi.Contract.ImpactcardsabiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Impactcardsabi *ImpactcardsabiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Impactcardsabi.Contract.ImpactcardsabiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Impactcardsabi *ImpactcardsabiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Impactcardsabi.Contract.ImpactcardsabiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Impactcardsabi *ImpactcardsabiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Impactcardsabi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Impactcardsabi *ImpactcardsabiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Impactcardsabi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Impactcardsabi *ImpactcardsabiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Impactcardsabi.Contract.contract.Transact(opts, method, params...)
}

// MAXCARDS is a free data retrieval call binding the contract method 0x7e744a76.
//
// Solidity: function MAX_CARDS() view returns(uint256)
func (_Impactcardsabi *ImpactcardsabiCaller) MAXCARDS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Impactcardsabi.contract.Call(opts, &out, "MAX_CARDS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXCARDS is a free data retrieval call binding the contract method 0x7e744a76.
//
// Solidity: function MAX_CARDS() view returns(uint256)
func (_Impactcardsabi *ImpactcardsabiSession) MAXCARDS() (*big.Int, error) {
	return _Impactcardsabi.Contract.MAXCARDS(&_Impactcardsabi.CallOpts)
}

// MAXCARDS is a free data retrieval call binding the contract method 0x7e744a76.
//
// Solidity: function MAX_CARDS() view returns(uint256)
func (_Impactcardsabi *ImpactcardsabiCallerSession) MAXCARDS() (*big.Int, error) {
	return _Impactcardsabi.Contract.MAXCARDS(&_Impactcardsabi.CallOpts)
}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint256)
func (_Impactcardsabi *ImpactcardsabiCaller) MAXSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Impactcardsabi.contract.Call(opts, &out, "MAX_SUPPLY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint256)
func (_Impactcardsabi *ImpactcardsabiSession) MAXSUPPLY() (*big.Int, error) {
	return _Impactcardsabi.Contract.MAXSUPPLY(&_Impactcardsabi.CallOpts)
}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint256)
func (_Impactcardsabi *ImpactcardsabiCallerSession) MAXSUPPLY() (*big.Int, error) {
	return _Impactcardsabi.Contract.MAXSUPPLY(&_Impactcardsabi.CallOpts)
}

// BaseExtension is a free data retrieval call binding the contract method 0x2672c902.
//
// Solidity: function _baseExtension() view returns(string)
func (_Impactcardsabi *ImpactcardsabiCaller) BaseExtension(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Impactcardsabi.contract.Call(opts, &out, "_baseExtension")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseExtension is a free data retrieval call binding the contract method 0x2672c902.
//
// Solidity: function _baseExtension() view returns(string)
func (_Impactcardsabi *ImpactcardsabiSession) BaseExtension() (string, error) {
	return _Impactcardsabi.Contract.BaseExtension(&_Impactcardsabi.CallOpts)
}

// BaseExtension is a free data retrieval call binding the contract method 0x2672c902.
//
// Solidity: function _baseExtension() view returns(string)
func (_Impactcardsabi *ImpactcardsabiCallerSession) BaseExtension() (string, error) {
	return _Impactcardsabi.Contract.BaseExtension(&_Impactcardsabi.CallOpts)
}

// Uri is a free data retrieval call binding the contract method 0x0dccc9ad.
//
// Solidity: function _uri() view returns(string)
func (_Impactcardsabi *ImpactcardsabiCaller) Uri(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Impactcardsabi.contract.Call(opts, &out, "_uri")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0dccc9ad.
//
// Solidity: function _uri() view returns(string)
func (_Impactcardsabi *ImpactcardsabiSession) Uri() (string, error) {
	return _Impactcardsabi.Contract.Uri(&_Impactcardsabi.CallOpts)
}

// Uri is a free data retrieval call binding the contract method 0x0dccc9ad.
//
// Solidity: function _uri() view returns(string)
func (_Impactcardsabi *ImpactcardsabiCallerSession) Uri() (string, error) {
	return _Impactcardsabi.Contract.Uri(&_Impactcardsabi.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Impactcardsabi *ImpactcardsabiCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Impactcardsabi.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Impactcardsabi *ImpactcardsabiSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Impactcardsabi.Contract.BalanceOf(&_Impactcardsabi.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Impactcardsabi *ImpactcardsabiCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Impactcardsabi.Contract.BalanceOf(&_Impactcardsabi.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Impactcardsabi *ImpactcardsabiCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Impactcardsabi.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Impactcardsabi *ImpactcardsabiSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Impactcardsabi.Contract.BalanceOfBatch(&_Impactcardsabi.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Impactcardsabi *ImpactcardsabiCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Impactcardsabi.Contract.BalanceOfBatch(&_Impactcardsabi.CallOpts, accounts, ids)
}

// BulkBuyLimit is a free data retrieval call binding the contract method 0xa49bccca.
//
// Solidity: function bulkBuyLimit() view returns(uint16)
func (_Impactcardsabi *ImpactcardsabiCaller) BulkBuyLimit(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Impactcardsabi.contract.Call(opts, &out, "bulkBuyLimit")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BulkBuyLimit is a free data retrieval call binding the contract method 0xa49bccca.
//
// Solidity: function bulkBuyLimit() view returns(uint16)
func (_Impactcardsabi *ImpactcardsabiSession) BulkBuyLimit() (uint16, error) {
	return _Impactcardsabi.Contract.BulkBuyLimit(&_Impactcardsabi.CallOpts)
}

// BulkBuyLimit is a free data retrieval call binding the contract method 0xa49bccca.
//
// Solidity: function bulkBuyLimit() view returns(uint16)
func (_Impactcardsabi *ImpactcardsabiCallerSession) BulkBuyLimit() (uint16, error) {
	return _Impactcardsabi.Contract.BulkBuyLimit(&_Impactcardsabi.CallOpts)
}

// CurrentSeason is a free data retrieval call binding the contract method 0xbcb39621.
//
// Solidity: function currentSeason() view returns(uint256)
func (_Impactcardsabi *ImpactcardsabiCaller) CurrentSeason(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Impactcardsabi.contract.Call(opts, &out, "currentSeason")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentSeason is a free data retrieval call binding the contract method 0xbcb39621.
//
// Solidity: function currentSeason() view returns(uint256)
func (_Impactcardsabi *ImpactcardsabiSession) CurrentSeason() (*big.Int, error) {
	return _Impactcardsabi.Contract.CurrentSeason(&_Impactcardsabi.CallOpts)
}

// CurrentSeason is a free data retrieval call binding the contract method 0xbcb39621.
//
// Solidity: function currentSeason() view returns(uint256)
func (_Impactcardsabi *ImpactcardsabiCallerSession) CurrentSeason() (*big.Int, error) {
	return _Impactcardsabi.Contract.CurrentSeason(&_Impactcardsabi.CallOpts)
}

// GetAccumulatedFunds is a free data retrieval call binding the contract method 0x72c410a6.
//
// Solidity: function getAccumulatedFunds(uint256 tokenId) view returns(uint256[2])
func (_Impactcardsabi *ImpactcardsabiCaller) GetAccumulatedFunds(opts *bind.CallOpts, tokenId *big.Int) ([2]*big.Int, error) {
	var out []interface{}
	err := _Impactcardsabi.contract.Call(opts, &out, "getAccumulatedFunds", tokenId)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetAccumulatedFunds is a free data retrieval call binding the contract method 0x72c410a6.
//
// Solidity: function getAccumulatedFunds(uint256 tokenId) view returns(uint256[2])
func (_Impactcardsabi *ImpactcardsabiSession) GetAccumulatedFunds(tokenId *big.Int) ([2]*big.Int, error) {
	return _Impactcardsabi.Contract.GetAccumulatedFunds(&_Impactcardsabi.CallOpts, tokenId)
}

// GetAccumulatedFunds is a free data retrieval call binding the contract method 0x72c410a6.
//
// Solidity: function getAccumulatedFunds(uint256 tokenId) view returns(uint256[2])
func (_Impactcardsabi *ImpactcardsabiCallerSession) GetAccumulatedFunds(tokenId *big.Int) ([2]*big.Int, error) {
	return _Impactcardsabi.Contract.GetAccumulatedFunds(&_Impactcardsabi.CallOpts, tokenId)
}

// GetPayees is a free data retrieval call binding the contract method 0x6c02c75c.
//
// Solidity: function getPayees(uint256 tokenId) view returns(address[2])
func (_Impactcardsabi *ImpactcardsabiCaller) GetPayees(opts *bind.CallOpts, tokenId *big.Int) ([2]common.Address, error) {
	var out []interface{}
	err := _Impactcardsabi.contract.Call(opts, &out, "getPayees", tokenId)

	if err != nil {
		return *new([2]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([2]common.Address)).(*[2]common.Address)

	return out0, err

}

// GetPayees is a free data retrieval call binding the contract method 0x6c02c75c.
//
// Solidity: function getPayees(uint256 tokenId) view returns(address[2])
func (_Impactcardsabi *ImpactcardsabiSession) GetPayees(tokenId *big.Int) ([2]common.Address, error) {
	return _Impactcardsabi.Contract.GetPayees(&_Impactcardsabi.CallOpts, tokenId)
}

// GetPayees is a free data retrieval call binding the contract method 0x6c02c75c.
//
// Solidity: function getPayees(uint256 tokenId) view returns(address[2])
func (_Impactcardsabi *ImpactcardsabiCallerSession) GetPayees(tokenId *big.Int) ([2]common.Address, error) {
	return _Impactcardsabi.Contract.GetPayees(&_Impactcardsabi.CallOpts, tokenId)
}

// GetShares is a free data retrieval call binding the contract method 0x6f7267b7.
//
// Solidity: function getShares(uint256 tokenId) view returns(uint256[2])
func (_Impactcardsabi *ImpactcardsabiCaller) GetShares(opts *bind.CallOpts, tokenId *big.Int) ([2]*big.Int, error) {
	var out []interface{}
	err := _Impactcardsabi.contract.Call(opts, &out, "getShares", tokenId)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetShares is a free data retrieval call binding the contract method 0x6f7267b7.
//
// Solidity: function getShares(uint256 tokenId) view returns(uint256[2])
func (_Impactcardsabi *ImpactcardsabiSession) GetShares(tokenId *big.Int) ([2]*big.Int, error) {
	return _Impactcardsabi.Contract.GetShares(&_Impactcardsabi.CallOpts, tokenId)
}

// GetShares is a free data retrieval call binding the contract method 0x6f7267b7.
//
// Solidity: function getShares(uint256 tokenId) view returns(uint256[2])
func (_Impactcardsabi *ImpactcardsabiCallerSession) GetShares(tokenId *big.Int) ([2]*big.Int, error) {
	return _Impactcardsabi.Contract.GetShares(&_Impactcardsabi.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Impactcardsabi *ImpactcardsabiCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Impactcardsabi.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Impactcardsabi *ImpactcardsabiSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Impactcardsabi.Contract.IsApprovedForAll(&_Impactcardsabi.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Impactcardsabi *ImpactcardsabiCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Impactcardsabi.Contract.IsApprovedForAll(&_Impactcardsabi.CallOpts, account, operator)
}

// IsMintable is a free data retrieval call binding the contract method 0x5b31cb36.
//
// Solidity: function isMintable(uint256 tokenId) view returns(bool)
func (_Impactcardsabi *ImpactcardsabiCaller) IsMintable(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Impactcardsabi.contract.Call(opts, &out, "isMintable", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMintable is a free data retrieval call binding the contract method 0x5b31cb36.
//
// Solidity: function isMintable(uint256 tokenId) view returns(bool)
func (_Impactcardsabi *ImpactcardsabiSession) IsMintable(tokenId *big.Int) (bool, error) {
	return _Impactcardsabi.Contract.IsMintable(&_Impactcardsabi.CallOpts, tokenId)
}

// IsMintable is a free data retrieval call binding the contract method 0x5b31cb36.
//
// Solidity: function isMintable(uint256 tokenId) view returns(bool)
func (_Impactcardsabi *ImpactcardsabiCallerSession) IsMintable(tokenId *big.Int) (bool, error) {
	return _Impactcardsabi.Contract.IsMintable(&_Impactcardsabi.CallOpts, tokenId)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Impactcardsabi *ImpactcardsabiCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Impactcardsabi.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Impactcardsabi *ImpactcardsabiSession) IsPaused() (bool, error) {
	return _Impactcardsabi.Contract.IsPaused(&_Impactcardsabi.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Impactcardsabi *ImpactcardsabiCallerSession) IsPaused() (bool, error) {
	return _Impactcardsabi.Contract.IsPaused(&_Impactcardsabi.CallOpts)
}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_Impactcardsabi *ImpactcardsabiCaller) MintPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Impactcardsabi.contract.Call(opts, &out, "mintPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_Impactcardsabi *ImpactcardsabiSession) MintPrice() (*big.Int, error) {
	return _Impactcardsabi.Contract.MintPrice(&_Impactcardsabi.CallOpts)
}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_Impactcardsabi *ImpactcardsabiCallerSession) MintPrice() (*big.Int, error) {
	return _Impactcardsabi.Contract.MintPrice(&_Impactcardsabi.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Impactcardsabi *ImpactcardsabiCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Impactcardsabi.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Impactcardsabi *ImpactcardsabiSession) Name() (string, error) {
	return _Impactcardsabi.Contract.Name(&_Impactcardsabi.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Impactcardsabi *ImpactcardsabiCallerSession) Name() (string, error) {
	return _Impactcardsabi.Contract.Name(&_Impactcardsabi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Impactcardsabi *ImpactcardsabiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Impactcardsabi.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Impactcardsabi *ImpactcardsabiSession) Owner() (common.Address, error) {
	return _Impactcardsabi.Contract.Owner(&_Impactcardsabi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Impactcardsabi *ImpactcardsabiCallerSession) Owner() (common.Address, error) {
	return _Impactcardsabi.Contract.Owner(&_Impactcardsabi.CallOpts)
}

// SeasonCardCounts is a free data retrieval call binding the contract method 0xb0fae03b.
//
// Solidity: function seasonCardCounts(uint256 ) view returns(uint8)
func (_Impactcardsabi *ImpactcardsabiCaller) SeasonCardCounts(opts *bind.CallOpts, arg0 *big.Int) (uint8, error) {
	var out []interface{}
	err := _Impactcardsabi.contract.Call(opts, &out, "seasonCardCounts", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SeasonCardCounts is a free data retrieval call binding the contract method 0xb0fae03b.
//
// Solidity: function seasonCardCounts(uint256 ) view returns(uint8)
func (_Impactcardsabi *ImpactcardsabiSession) SeasonCardCounts(arg0 *big.Int) (uint8, error) {
	return _Impactcardsabi.Contract.SeasonCardCounts(&_Impactcardsabi.CallOpts, arg0)
}

// SeasonCardCounts is a free data retrieval call binding the contract method 0xb0fae03b.
//
// Solidity: function seasonCardCounts(uint256 ) view returns(uint8)
func (_Impactcardsabi *ImpactcardsabiCallerSession) SeasonCardCounts(arg0 *big.Int) (uint8, error) {
	return _Impactcardsabi.Contract.SeasonCardCounts(&_Impactcardsabi.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Impactcardsabi *ImpactcardsabiCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Impactcardsabi.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Impactcardsabi *ImpactcardsabiSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Impactcardsabi.Contract.SupportsInterface(&_Impactcardsabi.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Impactcardsabi *ImpactcardsabiCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Impactcardsabi.Contract.SupportsInterface(&_Impactcardsabi.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Impactcardsabi *ImpactcardsabiCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Impactcardsabi.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Impactcardsabi *ImpactcardsabiSession) Symbol() (string, error) {
	return _Impactcardsabi.Contract.Symbol(&_Impactcardsabi.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Impactcardsabi *ImpactcardsabiCallerSession) Symbol() (string, error) {
	return _Impactcardsabi.Contract.Symbol(&_Impactcardsabi.CallOpts)
}

// TotalReleasedToPayee is a free data retrieval call binding the contract method 0x16f44232.
//
// Solidity: function totalReleasedToPayee(address payee) view returns(uint256)
func (_Impactcardsabi *ImpactcardsabiCaller) TotalReleasedToPayee(opts *bind.CallOpts, payee common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Impactcardsabi.contract.Call(opts, &out, "totalReleasedToPayee", payee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalReleasedToPayee is a free data retrieval call binding the contract method 0x16f44232.
//
// Solidity: function totalReleasedToPayee(address payee) view returns(uint256)
func (_Impactcardsabi *ImpactcardsabiSession) TotalReleasedToPayee(payee common.Address) (*big.Int, error) {
	return _Impactcardsabi.Contract.TotalReleasedToPayee(&_Impactcardsabi.CallOpts, payee)
}

// TotalReleasedToPayee is a free data retrieval call binding the contract method 0x16f44232.
//
// Solidity: function totalReleasedToPayee(address payee) view returns(uint256)
func (_Impactcardsabi *ImpactcardsabiCallerSession) TotalReleasedToPayee(payee common.Address) (*big.Int, error) {
	return _Impactcardsabi.Contract.TotalReleasedToPayee(&_Impactcardsabi.CallOpts, payee)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 tokenId) view returns(uint256)
func (_Impactcardsabi *ImpactcardsabiCaller) TotalSupply(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Impactcardsabi.contract.Call(opts, &out, "totalSupply", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 tokenId) view returns(uint256)
func (_Impactcardsabi *ImpactcardsabiSession) TotalSupply(tokenId *big.Int) (*big.Int, error) {
	return _Impactcardsabi.Contract.TotalSupply(&_Impactcardsabi.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 tokenId) view returns(uint256)
func (_Impactcardsabi *ImpactcardsabiCallerSession) TotalSupply(tokenId *big.Int) (*big.Int, error) {
	return _Impactcardsabi.Contract.TotalSupply(&_Impactcardsabi.CallOpts, tokenId)
}

// UriAlias is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_Impactcardsabi *ImpactcardsabiCaller) UriAlias(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Impactcardsabi.contract.Call(opts, &out, "uri", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UriAlias is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_Impactcardsabi *ImpactcardsabiSession) UriAlias(tokenId *big.Int) (string, error) {
	return _Impactcardsabi.Contract.UriAlias(&_Impactcardsabi.CallOpts, tokenId)
}

// UriAlias is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_Impactcardsabi *ImpactcardsabiCallerSession) UriAlias(tokenId *big.Int) (string, error) {
	return _Impactcardsabi.Contract.UriAlias(&_Impactcardsabi.CallOpts, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x1b2ef1ca.
//
// Solidity: function mint(uint256 tokenId, uint256 amount) payable returns()
func (_Impactcardsabi *ImpactcardsabiTransactor) Mint(opts *bind.TransactOpts, tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Impactcardsabi.contract.Transact(opts, "mint", tokenId, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x1b2ef1ca.
//
// Solidity: function mint(uint256 tokenId, uint256 amount) payable returns()
func (_Impactcardsabi *ImpactcardsabiSession) Mint(tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Impactcardsabi.Contract.Mint(&_Impactcardsabi.TransactOpts, tokenId, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x1b2ef1ca.
//
// Solidity: function mint(uint256 tokenId, uint256 amount) payable returns()
func (_Impactcardsabi *ImpactcardsabiTransactorSession) Mint(tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Impactcardsabi.Contract.Mint(&_Impactcardsabi.TransactOpts, tokenId, amount)
}

// MintBatch is a paid mutator transaction binding the contract method 0xd351cfdc.
//
// Solidity: function mintBatch(uint256[] tokenId, uint256[] amount) payable returns()
func (_Impactcardsabi *ImpactcardsabiTransactor) MintBatch(opts *bind.TransactOpts, tokenId []*big.Int, amount []*big.Int) (*types.Transaction, error) {
	return _Impactcardsabi.contract.Transact(opts, "mintBatch", tokenId, amount)
}

// MintBatch is a paid mutator transaction binding the contract method 0xd351cfdc.
//
// Solidity: function mintBatch(uint256[] tokenId, uint256[] amount) payable returns()
func (_Impactcardsabi *ImpactcardsabiSession) MintBatch(tokenId []*big.Int, amount []*big.Int) (*types.Transaction, error) {
	return _Impactcardsabi.Contract.MintBatch(&_Impactcardsabi.TransactOpts, tokenId, amount)
}

// MintBatch is a paid mutator transaction binding the contract method 0xd351cfdc.
//
// Solidity: function mintBatch(uint256[] tokenId, uint256[] amount) payable returns()
func (_Impactcardsabi *ImpactcardsabiTransactorSession) MintBatch(tokenId []*big.Int, amount []*big.Int) (*types.Transaction, error) {
	return _Impactcardsabi.Contract.MintBatch(&_Impactcardsabi.TransactOpts, tokenId, amount)
}

// NextSeason is a paid mutator transaction binding the contract method 0xbc734f0f.
//
// Solidity: function nextSeason() returns()
func (_Impactcardsabi *ImpactcardsabiTransactor) NextSeason(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Impactcardsabi.contract.Transact(opts, "nextSeason")
}

// NextSeason is a paid mutator transaction binding the contract method 0xbc734f0f.
//
// Solidity: function nextSeason() returns()
func (_Impactcardsabi *ImpactcardsabiSession) NextSeason() (*types.Transaction, error) {
	return _Impactcardsabi.Contract.NextSeason(&_Impactcardsabi.TransactOpts)
}

// NextSeason is a paid mutator transaction binding the contract method 0xbc734f0f.
//
// Solidity: function nextSeason() returns()
func (_Impactcardsabi *ImpactcardsabiTransactorSession) NextSeason() (*types.Transaction, error) {
	return _Impactcardsabi.Contract.NextSeason(&_Impactcardsabi.TransactOpts)
}

// Release is a paid mutator transaction binding the contract method 0x0eae0f0e.
//
// Solidity: function release(uint256 tokenId, uint8 payeeIndex) returns()
func (_Impactcardsabi *ImpactcardsabiTransactor) Release(opts *bind.TransactOpts, tokenId *big.Int, payeeIndex uint8) (*types.Transaction, error) {
	return _Impactcardsabi.contract.Transact(opts, "release", tokenId, payeeIndex)
}

// Release is a paid mutator transaction binding the contract method 0x0eae0f0e.
//
// Solidity: function release(uint256 tokenId, uint8 payeeIndex) returns()
func (_Impactcardsabi *ImpactcardsabiSession) Release(tokenId *big.Int, payeeIndex uint8) (*types.Transaction, error) {
	return _Impactcardsabi.Contract.Release(&_Impactcardsabi.TransactOpts, tokenId, payeeIndex)
}

// Release is a paid mutator transaction binding the contract method 0x0eae0f0e.
//
// Solidity: function release(uint256 tokenId, uint8 payeeIndex) returns()
func (_Impactcardsabi *ImpactcardsabiTransactorSession) Release(tokenId *big.Int, payeeIndex uint8) (*types.Transaction, error) {
	return _Impactcardsabi.Contract.Release(&_Impactcardsabi.TransactOpts, tokenId, payeeIndex)
}

// ReleaseToAllPayees is a paid mutator transaction binding the contract method 0xa67b8bf3.
//
// Solidity: function releaseToAllPayees() returns()
func (_Impactcardsabi *ImpactcardsabiTransactor) ReleaseToAllPayees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Impactcardsabi.contract.Transact(opts, "releaseToAllPayees")
}

// ReleaseToAllPayees is a paid mutator transaction binding the contract method 0xa67b8bf3.
//
// Solidity: function releaseToAllPayees() returns()
func (_Impactcardsabi *ImpactcardsabiSession) ReleaseToAllPayees() (*types.Transaction, error) {
	return _Impactcardsabi.Contract.ReleaseToAllPayees(&_Impactcardsabi.TransactOpts)
}

// ReleaseToAllPayees is a paid mutator transaction binding the contract method 0xa67b8bf3.
//
// Solidity: function releaseToAllPayees() returns()
func (_Impactcardsabi *ImpactcardsabiTransactorSession) ReleaseToAllPayees() (*types.Transaction, error) {
	return _Impactcardsabi.Contract.ReleaseToAllPayees(&_Impactcardsabi.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Impactcardsabi *ImpactcardsabiTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Impactcardsabi.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Impactcardsabi *ImpactcardsabiSession) RenounceOwnership() (*types.Transaction, error) {
	return _Impactcardsabi.Contract.RenounceOwnership(&_Impactcardsabi.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Impactcardsabi *ImpactcardsabiTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Impactcardsabi.Contract.RenounceOwnership(&_Impactcardsabi.TransactOpts)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Impactcardsabi *ImpactcardsabiTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Impactcardsabi.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Impactcardsabi *ImpactcardsabiSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Impactcardsabi.Contract.SafeBatchTransferFrom(&_Impactcardsabi.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Impactcardsabi *ImpactcardsabiTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Impactcardsabi.Contract.SafeBatchTransferFrom(&_Impactcardsabi.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Impactcardsabi *ImpactcardsabiTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Impactcardsabi.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Impactcardsabi *ImpactcardsabiSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Impactcardsabi.Contract.SafeTransferFrom(&_Impactcardsabi.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Impactcardsabi *ImpactcardsabiTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Impactcardsabi.Contract.SafeTransferFrom(&_Impactcardsabi.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Impactcardsabi *ImpactcardsabiTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Impactcardsabi.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Impactcardsabi *ImpactcardsabiSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Impactcardsabi.Contract.SetApprovalForAll(&_Impactcardsabi.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Impactcardsabi *ImpactcardsabiTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Impactcardsabi.Contract.SetApprovalForAll(&_Impactcardsabi.TransactOpts, operator, approved)
}

// SetBaseExtension is a paid mutator transaction binding the contract method 0xda3ef23f.
//
// Solidity: function setBaseExtension(string newExtension) returns()
func (_Impactcardsabi *ImpactcardsabiTransactor) SetBaseExtension(opts *bind.TransactOpts, newExtension string) (*types.Transaction, error) {
	return _Impactcardsabi.contract.Transact(opts, "setBaseExtension", newExtension)
}

// SetBaseExtension is a paid mutator transaction binding the contract method 0xda3ef23f.
//
// Solidity: function setBaseExtension(string newExtension) returns()
func (_Impactcardsabi *ImpactcardsabiSession) SetBaseExtension(newExtension string) (*types.Transaction, error) {
	return _Impactcardsabi.Contract.SetBaseExtension(&_Impactcardsabi.TransactOpts, newExtension)
}

// SetBaseExtension is a paid mutator transaction binding the contract method 0xda3ef23f.
//
// Solidity: function setBaseExtension(string newExtension) returns()
func (_Impactcardsabi *ImpactcardsabiTransactorSession) SetBaseExtension(newExtension string) (*types.Transaction, error) {
	return _Impactcardsabi.Contract.SetBaseExtension(&_Impactcardsabi.TransactOpts, newExtension)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string newuri) returns()
func (_Impactcardsabi *ImpactcardsabiTransactor) SetBaseURI(opts *bind.TransactOpts, newuri string) (*types.Transaction, error) {
	return _Impactcardsabi.contract.Transact(opts, "setBaseURI", newuri)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string newuri) returns()
func (_Impactcardsabi *ImpactcardsabiSession) SetBaseURI(newuri string) (*types.Transaction, error) {
	return _Impactcardsabi.Contract.SetBaseURI(&_Impactcardsabi.TransactOpts, newuri)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string newuri) returns()
func (_Impactcardsabi *ImpactcardsabiTransactorSession) SetBaseURI(newuri string) (*types.Transaction, error) {
	return _Impactcardsabi.Contract.SetBaseURI(&_Impactcardsabi.TransactOpts, newuri)
}

// SetBulkBuyLimit is a paid mutator transaction binding the contract method 0x14b27590.
//
// Solidity: function setBulkBuyLimit(uint16 limit) returns()
func (_Impactcardsabi *ImpactcardsabiTransactor) SetBulkBuyLimit(opts *bind.TransactOpts, limit uint16) (*types.Transaction, error) {
	return _Impactcardsabi.contract.Transact(opts, "setBulkBuyLimit", limit)
}

// SetBulkBuyLimit is a paid mutator transaction binding the contract method 0x14b27590.
//
// Solidity: function setBulkBuyLimit(uint16 limit) returns()
func (_Impactcardsabi *ImpactcardsabiSession) SetBulkBuyLimit(limit uint16) (*types.Transaction, error) {
	return _Impactcardsabi.Contract.SetBulkBuyLimit(&_Impactcardsabi.TransactOpts, limit)
}

// SetBulkBuyLimit is a paid mutator transaction binding the contract method 0x14b27590.
//
// Solidity: function setBulkBuyLimit(uint16 limit) returns()
func (_Impactcardsabi *ImpactcardsabiTransactorSession) SetBulkBuyLimit(limit uint16) (*types.Transaction, error) {
	return _Impactcardsabi.Contract.SetBulkBuyLimit(&_Impactcardsabi.TransactOpts, limit)
}

// SetMintPrice is a paid mutator transaction binding the contract method 0xf4a0a528.
//
// Solidity: function setMintPrice(uint256 price) returns()
func (_Impactcardsabi *ImpactcardsabiTransactor) SetMintPrice(opts *bind.TransactOpts, price *big.Int) (*types.Transaction, error) {
	return _Impactcardsabi.contract.Transact(opts, "setMintPrice", price)
}

// SetMintPrice is a paid mutator transaction binding the contract method 0xf4a0a528.
//
// Solidity: function setMintPrice(uint256 price) returns()
func (_Impactcardsabi *ImpactcardsabiSession) SetMintPrice(price *big.Int) (*types.Transaction, error) {
	return _Impactcardsabi.Contract.SetMintPrice(&_Impactcardsabi.TransactOpts, price)
}

// SetMintPrice is a paid mutator transaction binding the contract method 0xf4a0a528.
//
// Solidity: function setMintPrice(uint256 price) returns()
func (_Impactcardsabi *ImpactcardsabiTransactorSession) SetMintPrice(price *big.Int) (*types.Transaction, error) {
	return _Impactcardsabi.Contract.SetMintPrice(&_Impactcardsabi.TransactOpts, price)
}

// SetPayees is a paid mutator transaction binding the contract method 0x1cff6268.
//
// Solidity: function setPayees(uint256 tokenId, address[2] payees, uint256[2] shares) returns()
func (_Impactcardsabi *ImpactcardsabiTransactor) SetPayees(opts *bind.TransactOpts, tokenId *big.Int, payees [2]common.Address, shares [2]*big.Int) (*types.Transaction, error) {
	return _Impactcardsabi.contract.Transact(opts, "setPayees", tokenId, payees, shares)
}

// SetPayees is a paid mutator transaction binding the contract method 0x1cff6268.
//
// Solidity: function setPayees(uint256 tokenId, address[2] payees, uint256[2] shares) returns()
func (_Impactcardsabi *ImpactcardsabiSession) SetPayees(tokenId *big.Int, payees [2]common.Address, shares [2]*big.Int) (*types.Transaction, error) {
	return _Impactcardsabi.Contract.SetPayees(&_Impactcardsabi.TransactOpts, tokenId, payees, shares)
}

// SetPayees is a paid mutator transaction binding the contract method 0x1cff6268.
//
// Solidity: function setPayees(uint256 tokenId, address[2] payees, uint256[2] shares) returns()
func (_Impactcardsabi *ImpactcardsabiTransactorSession) SetPayees(tokenId *big.Int, payees [2]common.Address, shares [2]*big.Int) (*types.Transaction, error) {
	return _Impactcardsabi.Contract.SetPayees(&_Impactcardsabi.TransactOpts, tokenId, payees, shares)
}

// TogglePaused is a paid mutator transaction binding the contract method 0x36566f06.
//
// Solidity: function togglePaused() returns()
func (_Impactcardsabi *ImpactcardsabiTransactor) TogglePaused(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Impactcardsabi.contract.Transact(opts, "togglePaused")
}

// TogglePaused is a paid mutator transaction binding the contract method 0x36566f06.
//
// Solidity: function togglePaused() returns()
func (_Impactcardsabi *ImpactcardsabiSession) TogglePaused() (*types.Transaction, error) {
	return _Impactcardsabi.Contract.TogglePaused(&_Impactcardsabi.TransactOpts)
}

// TogglePaused is a paid mutator transaction binding the contract method 0x36566f06.
//
// Solidity: function togglePaused() returns()
func (_Impactcardsabi *ImpactcardsabiTransactorSession) TogglePaused() (*types.Transaction, error) {
	return _Impactcardsabi.Contract.TogglePaused(&_Impactcardsabi.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Impactcardsabi *ImpactcardsabiTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Impactcardsabi.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Impactcardsabi *ImpactcardsabiSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Impactcardsabi.Contract.TransferOwnership(&_Impactcardsabi.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Impactcardsabi *ImpactcardsabiTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Impactcardsabi.Contract.TransferOwnership(&_Impactcardsabi.TransactOpts, newOwner)
}

// ImpactcardsabiApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Impactcardsabi contract.
type ImpactcardsabiApprovalForAllIterator struct {
	Event *ImpactcardsabiApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ImpactcardsabiApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImpactcardsabiApprovalForAll)
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
		it.Event = new(ImpactcardsabiApprovalForAll)
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
func (it *ImpactcardsabiApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImpactcardsabiApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImpactcardsabiApprovalForAll represents a ApprovalForAll event raised by the Impactcardsabi contract.
type ImpactcardsabiApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Impactcardsabi *ImpactcardsabiFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*ImpactcardsabiApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Impactcardsabi.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ImpactcardsabiApprovalForAllIterator{contract: _Impactcardsabi.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Impactcardsabi *ImpactcardsabiFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ImpactcardsabiApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Impactcardsabi.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImpactcardsabiApprovalForAll)
				if err := _Impactcardsabi.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Impactcardsabi *ImpactcardsabiFilterer) ParseApprovalForAll(log types.Log) (*ImpactcardsabiApprovalForAll, error) {
	event := new(ImpactcardsabiApprovalForAll)
	if err := _Impactcardsabi.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ImpactcardsabiBaseExtensionChangedIterator is returned from FilterBaseExtensionChanged and is used to iterate over the raw logs and unpacked data for BaseExtensionChanged events raised by the Impactcardsabi contract.
type ImpactcardsabiBaseExtensionChangedIterator struct {
	Event *ImpactcardsabiBaseExtensionChanged // Event containing the contract specifics and raw log

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
func (it *ImpactcardsabiBaseExtensionChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImpactcardsabiBaseExtensionChanged)
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
		it.Event = new(ImpactcardsabiBaseExtensionChanged)
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
func (it *ImpactcardsabiBaseExtensionChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImpactcardsabiBaseExtensionChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImpactcardsabiBaseExtensionChanged represents a BaseExtensionChanged event raised by the Impactcardsabi contract.
type ImpactcardsabiBaseExtensionChanged struct {
	BaseExtension string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBaseExtensionChanged is a free log retrieval operation binding the contract event 0x8731128c6185faae57b7561c5f15b0fd4ef267565ef90a6980c4e4b81a25e420.
//
// Solidity: event BaseExtensionChanged(string baseExtension)
func (_Impactcardsabi *ImpactcardsabiFilterer) FilterBaseExtensionChanged(opts *bind.FilterOpts) (*ImpactcardsabiBaseExtensionChangedIterator, error) {

	logs, sub, err := _Impactcardsabi.contract.FilterLogs(opts, "BaseExtensionChanged")
	if err != nil {
		return nil, err
	}
	return &ImpactcardsabiBaseExtensionChangedIterator{contract: _Impactcardsabi.contract, event: "BaseExtensionChanged", logs: logs, sub: sub}, nil
}

// WatchBaseExtensionChanged is a free log subscription operation binding the contract event 0x8731128c6185faae57b7561c5f15b0fd4ef267565ef90a6980c4e4b81a25e420.
//
// Solidity: event BaseExtensionChanged(string baseExtension)
func (_Impactcardsabi *ImpactcardsabiFilterer) WatchBaseExtensionChanged(opts *bind.WatchOpts, sink chan<- *ImpactcardsabiBaseExtensionChanged) (event.Subscription, error) {

	logs, sub, err := _Impactcardsabi.contract.WatchLogs(opts, "BaseExtensionChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImpactcardsabiBaseExtensionChanged)
				if err := _Impactcardsabi.contract.UnpackLog(event, "BaseExtensionChanged", log); err != nil {
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

// ParseBaseExtensionChanged is a log parse operation binding the contract event 0x8731128c6185faae57b7561c5f15b0fd4ef267565ef90a6980c4e4b81a25e420.
//
// Solidity: event BaseExtensionChanged(string baseExtension)
func (_Impactcardsabi *ImpactcardsabiFilterer) ParseBaseExtensionChanged(log types.Log) (*ImpactcardsabiBaseExtensionChanged, error) {
	event := new(ImpactcardsabiBaseExtensionChanged)
	if err := _Impactcardsabi.contract.UnpackLog(event, "BaseExtensionChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ImpactcardsabiBaseURIChangedIterator is returned from FilterBaseURIChanged and is used to iterate over the raw logs and unpacked data for BaseURIChanged events raised by the Impactcardsabi contract.
type ImpactcardsabiBaseURIChangedIterator struct {
	Event *ImpactcardsabiBaseURIChanged // Event containing the contract specifics and raw log

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
func (it *ImpactcardsabiBaseURIChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImpactcardsabiBaseURIChanged)
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
		it.Event = new(ImpactcardsabiBaseURIChanged)
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
func (it *ImpactcardsabiBaseURIChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImpactcardsabiBaseURIChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImpactcardsabiBaseURIChanged represents a BaseURIChanged event raised by the Impactcardsabi contract.
type ImpactcardsabiBaseURIChanged struct {
	BaseURI string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBaseURIChanged is a free log retrieval operation binding the contract event 0x5411e8ebf1636d9e83d5fc4900bf80cbac82e8790da2a4c94db4895e889eedf6.
//
// Solidity: event BaseURIChanged(string baseURI)
func (_Impactcardsabi *ImpactcardsabiFilterer) FilterBaseURIChanged(opts *bind.FilterOpts) (*ImpactcardsabiBaseURIChangedIterator, error) {

	logs, sub, err := _Impactcardsabi.contract.FilterLogs(opts, "BaseURIChanged")
	if err != nil {
		return nil, err
	}
	return &ImpactcardsabiBaseURIChangedIterator{contract: _Impactcardsabi.contract, event: "BaseURIChanged", logs: logs, sub: sub}, nil
}

// WatchBaseURIChanged is a free log subscription operation binding the contract event 0x5411e8ebf1636d9e83d5fc4900bf80cbac82e8790da2a4c94db4895e889eedf6.
//
// Solidity: event BaseURIChanged(string baseURI)
func (_Impactcardsabi *ImpactcardsabiFilterer) WatchBaseURIChanged(opts *bind.WatchOpts, sink chan<- *ImpactcardsabiBaseURIChanged) (event.Subscription, error) {

	logs, sub, err := _Impactcardsabi.contract.WatchLogs(opts, "BaseURIChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImpactcardsabiBaseURIChanged)
				if err := _Impactcardsabi.contract.UnpackLog(event, "BaseURIChanged", log); err != nil {
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

// ParseBaseURIChanged is a log parse operation binding the contract event 0x5411e8ebf1636d9e83d5fc4900bf80cbac82e8790da2a4c94db4895e889eedf6.
//
// Solidity: event BaseURIChanged(string baseURI)
func (_Impactcardsabi *ImpactcardsabiFilterer) ParseBaseURIChanged(log types.Log) (*ImpactcardsabiBaseURIChanged, error) {
	event := new(ImpactcardsabiBaseURIChanged)
	if err := _Impactcardsabi.contract.UnpackLog(event, "BaseURIChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ImpactcardsabiBulkBuyLimitSetIterator is returned from FilterBulkBuyLimitSet and is used to iterate over the raw logs and unpacked data for BulkBuyLimitSet events raised by the Impactcardsabi contract.
type ImpactcardsabiBulkBuyLimitSetIterator struct {
	Event *ImpactcardsabiBulkBuyLimitSet // Event containing the contract specifics and raw log

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
func (it *ImpactcardsabiBulkBuyLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImpactcardsabiBulkBuyLimitSet)
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
		it.Event = new(ImpactcardsabiBulkBuyLimitSet)
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
func (it *ImpactcardsabiBulkBuyLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImpactcardsabiBulkBuyLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImpactcardsabiBulkBuyLimitSet represents a BulkBuyLimitSet event raised by the Impactcardsabi contract.
type ImpactcardsabiBulkBuyLimitSet struct {
	Limit uint16
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBulkBuyLimitSet is a free log retrieval operation binding the contract event 0x5b6913e01c266ca9d82244c7ad7ff446888a2e08a9a6d3554386fbf22fae36e1.
//
// Solidity: event BulkBuyLimitSet(uint16 limit)
func (_Impactcardsabi *ImpactcardsabiFilterer) FilterBulkBuyLimitSet(opts *bind.FilterOpts) (*ImpactcardsabiBulkBuyLimitSetIterator, error) {

	logs, sub, err := _Impactcardsabi.contract.FilterLogs(opts, "BulkBuyLimitSet")
	if err != nil {
		return nil, err
	}
	return &ImpactcardsabiBulkBuyLimitSetIterator{contract: _Impactcardsabi.contract, event: "BulkBuyLimitSet", logs: logs, sub: sub}, nil
}

// WatchBulkBuyLimitSet is a free log subscription operation binding the contract event 0x5b6913e01c266ca9d82244c7ad7ff446888a2e08a9a6d3554386fbf22fae36e1.
//
// Solidity: event BulkBuyLimitSet(uint16 limit)
func (_Impactcardsabi *ImpactcardsabiFilterer) WatchBulkBuyLimitSet(opts *bind.WatchOpts, sink chan<- *ImpactcardsabiBulkBuyLimitSet) (event.Subscription, error) {

	logs, sub, err := _Impactcardsabi.contract.WatchLogs(opts, "BulkBuyLimitSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImpactcardsabiBulkBuyLimitSet)
				if err := _Impactcardsabi.contract.UnpackLog(event, "BulkBuyLimitSet", log); err != nil {
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

// ParseBulkBuyLimitSet is a log parse operation binding the contract event 0x5b6913e01c266ca9d82244c7ad7ff446888a2e08a9a6d3554386fbf22fae36e1.
//
// Solidity: event BulkBuyLimitSet(uint16 limit)
func (_Impactcardsabi *ImpactcardsabiFilterer) ParseBulkBuyLimitSet(log types.Log) (*ImpactcardsabiBulkBuyLimitSet, error) {
	event := new(ImpactcardsabiBulkBuyLimitSet)
	if err := _Impactcardsabi.contract.UnpackLog(event, "BulkBuyLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ImpactcardsabiCardMintedIterator is returned from FilterCardMinted and is used to iterate over the raw logs and unpacked data for CardMinted events raised by the Impactcardsabi contract.
type ImpactcardsabiCardMintedIterator struct {
	Event *ImpactcardsabiCardMinted // Event containing the contract specifics and raw log

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
func (it *ImpactcardsabiCardMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImpactcardsabiCardMinted)
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
		it.Event = new(ImpactcardsabiCardMinted)
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
func (it *ImpactcardsabiCardMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImpactcardsabiCardMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImpactcardsabiCardMinted represents a CardMinted event raised by the Impactcardsabi contract.
type ImpactcardsabiCardMinted struct {
	TokenId *big.Int
	Amount  *big.Int
	Minter  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCardMinted is a free log retrieval operation binding the contract event 0x11bf33de2604234483e348d4c7ed39423cdb90c24a9ded01acf4e647ff533575.
//
// Solidity: event CardMinted(uint256 tokenId, uint256 amount, address minter)
func (_Impactcardsabi *ImpactcardsabiFilterer) FilterCardMinted(opts *bind.FilterOpts) (*ImpactcardsabiCardMintedIterator, error) {

	logs, sub, err := _Impactcardsabi.contract.FilterLogs(opts, "CardMinted")
	if err != nil {
		return nil, err
	}
	return &ImpactcardsabiCardMintedIterator{contract: _Impactcardsabi.contract, event: "CardMinted", logs: logs, sub: sub}, nil
}

// WatchCardMinted is a free log subscription operation binding the contract event 0x11bf33de2604234483e348d4c7ed39423cdb90c24a9ded01acf4e647ff533575.
//
// Solidity: event CardMinted(uint256 tokenId, uint256 amount, address minter)
func (_Impactcardsabi *ImpactcardsabiFilterer) WatchCardMinted(opts *bind.WatchOpts, sink chan<- *ImpactcardsabiCardMinted) (event.Subscription, error) {

	logs, sub, err := _Impactcardsabi.contract.WatchLogs(opts, "CardMinted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImpactcardsabiCardMinted)
				if err := _Impactcardsabi.contract.UnpackLog(event, "CardMinted", log); err != nil {
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

// ParseCardMinted is a log parse operation binding the contract event 0x11bf33de2604234483e348d4c7ed39423cdb90c24a9ded01acf4e647ff533575.
//
// Solidity: event CardMinted(uint256 tokenId, uint256 amount, address minter)
func (_Impactcardsabi *ImpactcardsabiFilterer) ParseCardMinted(log types.Log) (*ImpactcardsabiCardMinted, error) {
	event := new(ImpactcardsabiCardMinted)
	if err := _Impactcardsabi.contract.UnpackLog(event, "CardMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ImpactcardsabiCardsBatchMintedIterator is returned from FilterCardsBatchMinted and is used to iterate over the raw logs and unpacked data for CardsBatchMinted events raised by the Impactcardsabi contract.
type ImpactcardsabiCardsBatchMintedIterator struct {
	Event *ImpactcardsabiCardsBatchMinted // Event containing the contract specifics and raw log

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
func (it *ImpactcardsabiCardsBatchMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImpactcardsabiCardsBatchMinted)
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
		it.Event = new(ImpactcardsabiCardsBatchMinted)
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
func (it *ImpactcardsabiCardsBatchMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImpactcardsabiCardsBatchMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImpactcardsabiCardsBatchMinted represents a CardsBatchMinted event raised by the Impactcardsabi contract.
type ImpactcardsabiCardsBatchMinted struct {
	TokenIds []*big.Int
	Amounts  []*big.Int
	Minter   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCardsBatchMinted is a free log retrieval operation binding the contract event 0x355f71da8f0d3b720224957780562eb1a15c004f16d64a89d66ce075ca29a6e1.
//
// Solidity: event CardsBatchMinted(uint256[] tokenIds, uint256[] amounts, address minter)
func (_Impactcardsabi *ImpactcardsabiFilterer) FilterCardsBatchMinted(opts *bind.FilterOpts) (*ImpactcardsabiCardsBatchMintedIterator, error) {

	logs, sub, err := _Impactcardsabi.contract.FilterLogs(opts, "CardsBatchMinted")
	if err != nil {
		return nil, err
	}
	return &ImpactcardsabiCardsBatchMintedIterator{contract: _Impactcardsabi.contract, event: "CardsBatchMinted", logs: logs, sub: sub}, nil
}

// WatchCardsBatchMinted is a free log subscription operation binding the contract event 0x355f71da8f0d3b720224957780562eb1a15c004f16d64a89d66ce075ca29a6e1.
//
// Solidity: event CardsBatchMinted(uint256[] tokenIds, uint256[] amounts, address minter)
func (_Impactcardsabi *ImpactcardsabiFilterer) WatchCardsBatchMinted(opts *bind.WatchOpts, sink chan<- *ImpactcardsabiCardsBatchMinted) (event.Subscription, error) {

	logs, sub, err := _Impactcardsabi.contract.WatchLogs(opts, "CardsBatchMinted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImpactcardsabiCardsBatchMinted)
				if err := _Impactcardsabi.contract.UnpackLog(event, "CardsBatchMinted", log); err != nil {
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

// ParseCardsBatchMinted is a log parse operation binding the contract event 0x355f71da8f0d3b720224957780562eb1a15c004f16d64a89d66ce075ca29a6e1.
//
// Solidity: event CardsBatchMinted(uint256[] tokenIds, uint256[] amounts, address minter)
func (_Impactcardsabi *ImpactcardsabiFilterer) ParseCardsBatchMinted(log types.Log) (*ImpactcardsabiCardsBatchMinted, error) {
	event := new(ImpactcardsabiCardsBatchMinted)
	if err := _Impactcardsabi.contract.UnpackLog(event, "CardsBatchMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ImpactcardsabiChangedContractPausedStateIterator is returned from FilterChangedContractPausedState and is used to iterate over the raw logs and unpacked data for ChangedContractPausedState events raised by the Impactcardsabi contract.
type ImpactcardsabiChangedContractPausedStateIterator struct {
	Event *ImpactcardsabiChangedContractPausedState // Event containing the contract specifics and raw log

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
func (it *ImpactcardsabiChangedContractPausedStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImpactcardsabiChangedContractPausedState)
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
		it.Event = new(ImpactcardsabiChangedContractPausedState)
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
func (it *ImpactcardsabiChangedContractPausedStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImpactcardsabiChangedContractPausedStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImpactcardsabiChangedContractPausedState represents a ChangedContractPausedState event raised by the Impactcardsabi contract.
type ImpactcardsabiChangedContractPausedState struct {
	IsPaused bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChangedContractPausedState is a free log retrieval operation binding the contract event 0x5560278b5edfde1816c415de6402bf68981d1f9a503ffbdad45fbb77a3b4e844.
//
// Solidity: event ChangedContractPausedState(bool isPaused)
func (_Impactcardsabi *ImpactcardsabiFilterer) FilterChangedContractPausedState(opts *bind.FilterOpts) (*ImpactcardsabiChangedContractPausedStateIterator, error) {

	logs, sub, err := _Impactcardsabi.contract.FilterLogs(opts, "ChangedContractPausedState")
	if err != nil {
		return nil, err
	}
	return &ImpactcardsabiChangedContractPausedStateIterator{contract: _Impactcardsabi.contract, event: "ChangedContractPausedState", logs: logs, sub: sub}, nil
}

// WatchChangedContractPausedState is a free log subscription operation binding the contract event 0x5560278b5edfde1816c415de6402bf68981d1f9a503ffbdad45fbb77a3b4e844.
//
// Solidity: event ChangedContractPausedState(bool isPaused)
func (_Impactcardsabi *ImpactcardsabiFilterer) WatchChangedContractPausedState(opts *bind.WatchOpts, sink chan<- *ImpactcardsabiChangedContractPausedState) (event.Subscription, error) {

	logs, sub, err := _Impactcardsabi.contract.WatchLogs(opts, "ChangedContractPausedState")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImpactcardsabiChangedContractPausedState)
				if err := _Impactcardsabi.contract.UnpackLog(event, "ChangedContractPausedState", log); err != nil {
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

// ParseChangedContractPausedState is a log parse operation binding the contract event 0x5560278b5edfde1816c415de6402bf68981d1f9a503ffbdad45fbb77a3b4e844.
//
// Solidity: event ChangedContractPausedState(bool isPaused)
func (_Impactcardsabi *ImpactcardsabiFilterer) ParseChangedContractPausedState(log types.Log) (*ImpactcardsabiChangedContractPausedState, error) {
	event := new(ImpactcardsabiChangedContractPausedState)
	if err := _Impactcardsabi.contract.UnpackLog(event, "ChangedContractPausedState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ImpactcardsabiFundsReleasedIterator is returned from FilterFundsReleased and is used to iterate over the raw logs and unpacked data for FundsReleased events raised by the Impactcardsabi contract.
type ImpactcardsabiFundsReleasedIterator struct {
	Event *ImpactcardsabiFundsReleased // Event containing the contract specifics and raw log

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
func (it *ImpactcardsabiFundsReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImpactcardsabiFundsReleased)
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
		it.Event = new(ImpactcardsabiFundsReleased)
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
func (it *ImpactcardsabiFundsReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImpactcardsabiFundsReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImpactcardsabiFundsReleased represents a FundsReleased event raised by the Impactcardsabi contract.
type ImpactcardsabiFundsReleased struct {
	TokenId *big.Int
	Payee   common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFundsReleased is a free log retrieval operation binding the contract event 0x6e3c6096795c8298a218b2cfb8bde42726ff7c9a3d27b4d3ba41ab7f74feb5fb.
//
// Solidity: event FundsReleased(uint256 tokenId, address payee, uint256 amount)
func (_Impactcardsabi *ImpactcardsabiFilterer) FilterFundsReleased(opts *bind.FilterOpts) (*ImpactcardsabiFundsReleasedIterator, error) {

	logs, sub, err := _Impactcardsabi.contract.FilterLogs(opts, "FundsReleased")
	if err != nil {
		return nil, err
	}
	return &ImpactcardsabiFundsReleasedIterator{contract: _Impactcardsabi.contract, event: "FundsReleased", logs: logs, sub: sub}, nil
}

// WatchFundsReleased is a free log subscription operation binding the contract event 0x6e3c6096795c8298a218b2cfb8bde42726ff7c9a3d27b4d3ba41ab7f74feb5fb.
//
// Solidity: event FundsReleased(uint256 tokenId, address payee, uint256 amount)
func (_Impactcardsabi *ImpactcardsabiFilterer) WatchFundsReleased(opts *bind.WatchOpts, sink chan<- *ImpactcardsabiFundsReleased) (event.Subscription, error) {

	logs, sub, err := _Impactcardsabi.contract.WatchLogs(opts, "FundsReleased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImpactcardsabiFundsReleased)
				if err := _Impactcardsabi.contract.UnpackLog(event, "FundsReleased", log); err != nil {
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

// ParseFundsReleased is a log parse operation binding the contract event 0x6e3c6096795c8298a218b2cfb8bde42726ff7c9a3d27b4d3ba41ab7f74feb5fb.
//
// Solidity: event FundsReleased(uint256 tokenId, address payee, uint256 amount)
func (_Impactcardsabi *ImpactcardsabiFilterer) ParseFundsReleased(log types.Log) (*ImpactcardsabiFundsReleased, error) {
	event := new(ImpactcardsabiFundsReleased)
	if err := _Impactcardsabi.contract.UnpackLog(event, "FundsReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ImpactcardsabiMintPriceSetIterator is returned from FilterMintPriceSet and is used to iterate over the raw logs and unpacked data for MintPriceSet events raised by the Impactcardsabi contract.
type ImpactcardsabiMintPriceSetIterator struct {
	Event *ImpactcardsabiMintPriceSet // Event containing the contract specifics and raw log

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
func (it *ImpactcardsabiMintPriceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImpactcardsabiMintPriceSet)
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
		it.Event = new(ImpactcardsabiMintPriceSet)
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
func (it *ImpactcardsabiMintPriceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImpactcardsabiMintPriceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImpactcardsabiMintPriceSet represents a MintPriceSet event raised by the Impactcardsabi contract.
type ImpactcardsabiMintPriceSet struct {
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMintPriceSet is a free log retrieval operation binding the contract event 0x680d48f3ac056f34ba048f33f9724cf2b4023626c5baa2b6bc554ec47fbfa101.
//
// Solidity: event MintPriceSet(uint256 price)
func (_Impactcardsabi *ImpactcardsabiFilterer) FilterMintPriceSet(opts *bind.FilterOpts) (*ImpactcardsabiMintPriceSetIterator, error) {

	logs, sub, err := _Impactcardsabi.contract.FilterLogs(opts, "MintPriceSet")
	if err != nil {
		return nil, err
	}
	return &ImpactcardsabiMintPriceSetIterator{contract: _Impactcardsabi.contract, event: "MintPriceSet", logs: logs, sub: sub}, nil
}

// WatchMintPriceSet is a free log subscription operation binding the contract event 0x680d48f3ac056f34ba048f33f9724cf2b4023626c5baa2b6bc554ec47fbfa101.
//
// Solidity: event MintPriceSet(uint256 price)
func (_Impactcardsabi *ImpactcardsabiFilterer) WatchMintPriceSet(opts *bind.WatchOpts, sink chan<- *ImpactcardsabiMintPriceSet) (event.Subscription, error) {

	logs, sub, err := _Impactcardsabi.contract.WatchLogs(opts, "MintPriceSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImpactcardsabiMintPriceSet)
				if err := _Impactcardsabi.contract.UnpackLog(event, "MintPriceSet", log); err != nil {
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

// ParseMintPriceSet is a log parse operation binding the contract event 0x680d48f3ac056f34ba048f33f9724cf2b4023626c5baa2b6bc554ec47fbfa101.
//
// Solidity: event MintPriceSet(uint256 price)
func (_Impactcardsabi *ImpactcardsabiFilterer) ParseMintPriceSet(log types.Log) (*ImpactcardsabiMintPriceSet, error) {
	event := new(ImpactcardsabiMintPriceSet)
	if err := _Impactcardsabi.contract.UnpackLog(event, "MintPriceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ImpactcardsabiOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Impactcardsabi contract.
type ImpactcardsabiOwnershipTransferredIterator struct {
	Event *ImpactcardsabiOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ImpactcardsabiOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImpactcardsabiOwnershipTransferred)
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
		it.Event = new(ImpactcardsabiOwnershipTransferred)
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
func (it *ImpactcardsabiOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImpactcardsabiOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImpactcardsabiOwnershipTransferred represents a OwnershipTransferred event raised by the Impactcardsabi contract.
type ImpactcardsabiOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Impactcardsabi *ImpactcardsabiFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ImpactcardsabiOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Impactcardsabi.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ImpactcardsabiOwnershipTransferredIterator{contract: _Impactcardsabi.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Impactcardsabi *ImpactcardsabiFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ImpactcardsabiOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Impactcardsabi.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImpactcardsabiOwnershipTransferred)
				if err := _Impactcardsabi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Impactcardsabi *ImpactcardsabiFilterer) ParseOwnershipTransferred(log types.Log) (*ImpactcardsabiOwnershipTransferred, error) {
	event := new(ImpactcardsabiOwnershipTransferred)
	if err := _Impactcardsabi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ImpactcardsabiPayeeSetIterator is returned from FilterPayeeSet and is used to iterate over the raw logs and unpacked data for PayeeSet events raised by the Impactcardsabi contract.
type ImpactcardsabiPayeeSetIterator struct {
	Event *ImpactcardsabiPayeeSet // Event containing the contract specifics and raw log

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
func (it *ImpactcardsabiPayeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImpactcardsabiPayeeSet)
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
		it.Event = new(ImpactcardsabiPayeeSet)
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
func (it *ImpactcardsabiPayeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImpactcardsabiPayeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImpactcardsabiPayeeSet represents a PayeeSet event raised by the Impactcardsabi contract.
type ImpactcardsabiPayeeSet struct {
	TokenId *big.Int
	Payees  [2]common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPayeeSet is a free log retrieval operation binding the contract event 0x120ed5e733cc0dd8ace773ccaa200600da5cd3282021a6770067f439de80615e.
//
// Solidity: event PayeeSet(uint256 tokenId, address[2] payees)
func (_Impactcardsabi *ImpactcardsabiFilterer) FilterPayeeSet(opts *bind.FilterOpts) (*ImpactcardsabiPayeeSetIterator, error) {

	logs, sub, err := _Impactcardsabi.contract.FilterLogs(opts, "PayeeSet")
	if err != nil {
		return nil, err
	}
	return &ImpactcardsabiPayeeSetIterator{contract: _Impactcardsabi.contract, event: "PayeeSet", logs: logs, sub: sub}, nil
}

// WatchPayeeSet is a free log subscription operation binding the contract event 0x120ed5e733cc0dd8ace773ccaa200600da5cd3282021a6770067f439de80615e.
//
// Solidity: event PayeeSet(uint256 tokenId, address[2] payees)
func (_Impactcardsabi *ImpactcardsabiFilterer) WatchPayeeSet(opts *bind.WatchOpts, sink chan<- *ImpactcardsabiPayeeSet) (event.Subscription, error) {

	logs, sub, err := _Impactcardsabi.contract.WatchLogs(opts, "PayeeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImpactcardsabiPayeeSet)
				if err := _Impactcardsabi.contract.UnpackLog(event, "PayeeSet", log); err != nil {
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

// ParsePayeeSet is a log parse operation binding the contract event 0x120ed5e733cc0dd8ace773ccaa200600da5cd3282021a6770067f439de80615e.
//
// Solidity: event PayeeSet(uint256 tokenId, address[2] payees)
func (_Impactcardsabi *ImpactcardsabiFilterer) ParsePayeeSet(log types.Log) (*ImpactcardsabiPayeeSet, error) {
	event := new(ImpactcardsabiPayeeSet)
	if err := _Impactcardsabi.contract.UnpackLog(event, "PayeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ImpactcardsabiSeasonChangedIterator is returned from FilterSeasonChanged and is used to iterate over the raw logs and unpacked data for SeasonChanged events raised by the Impactcardsabi contract.
type ImpactcardsabiSeasonChangedIterator struct {
	Event *ImpactcardsabiSeasonChanged // Event containing the contract specifics and raw log

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
func (it *ImpactcardsabiSeasonChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImpactcardsabiSeasonChanged)
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
		it.Event = new(ImpactcardsabiSeasonChanged)
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
func (it *ImpactcardsabiSeasonChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImpactcardsabiSeasonChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImpactcardsabiSeasonChanged represents a SeasonChanged event raised by the Impactcardsabi contract.
type ImpactcardsabiSeasonChanged struct {
	Season *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSeasonChanged is a free log retrieval operation binding the contract event 0x6a9286bf980e56e1b373424a6139713d6ab9ef9254f1da993e5669711bfef8b8.
//
// Solidity: event SeasonChanged(uint256 season)
func (_Impactcardsabi *ImpactcardsabiFilterer) FilterSeasonChanged(opts *bind.FilterOpts) (*ImpactcardsabiSeasonChangedIterator, error) {

	logs, sub, err := _Impactcardsabi.contract.FilterLogs(opts, "SeasonChanged")
	if err != nil {
		return nil, err
	}
	return &ImpactcardsabiSeasonChangedIterator{contract: _Impactcardsabi.contract, event: "SeasonChanged", logs: logs, sub: sub}, nil
}

// WatchSeasonChanged is a free log subscription operation binding the contract event 0x6a9286bf980e56e1b373424a6139713d6ab9ef9254f1da993e5669711bfef8b8.
//
// Solidity: event SeasonChanged(uint256 season)
func (_Impactcardsabi *ImpactcardsabiFilterer) WatchSeasonChanged(opts *bind.WatchOpts, sink chan<- *ImpactcardsabiSeasonChanged) (event.Subscription, error) {

	logs, sub, err := _Impactcardsabi.contract.WatchLogs(opts, "SeasonChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImpactcardsabiSeasonChanged)
				if err := _Impactcardsabi.contract.UnpackLog(event, "SeasonChanged", log); err != nil {
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

// ParseSeasonChanged is a log parse operation binding the contract event 0x6a9286bf980e56e1b373424a6139713d6ab9ef9254f1da993e5669711bfef8b8.
//
// Solidity: event SeasonChanged(uint256 season)
func (_Impactcardsabi *ImpactcardsabiFilterer) ParseSeasonChanged(log types.Log) (*ImpactcardsabiSeasonChanged, error) {
	event := new(ImpactcardsabiSeasonChanged)
	if err := _Impactcardsabi.contract.UnpackLog(event, "SeasonChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ImpactcardsabiSharesSetIterator is returned from FilterSharesSet and is used to iterate over the raw logs and unpacked data for SharesSet events raised by the Impactcardsabi contract.
type ImpactcardsabiSharesSetIterator struct {
	Event *ImpactcardsabiSharesSet // Event containing the contract specifics and raw log

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
func (it *ImpactcardsabiSharesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImpactcardsabiSharesSet)
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
		it.Event = new(ImpactcardsabiSharesSet)
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
func (it *ImpactcardsabiSharesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImpactcardsabiSharesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImpactcardsabiSharesSet represents a SharesSet event raised by the Impactcardsabi contract.
type ImpactcardsabiSharesSet struct {
	TokenId *big.Int
	Shares  [2]*big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSharesSet is a free log retrieval operation binding the contract event 0xde702a3afda93461efb92f10ed3cdd0d0df3d292af2fc9dd64628e5f31a1901d.
//
// Solidity: event SharesSet(uint256 tokenId, uint256[2] shares)
func (_Impactcardsabi *ImpactcardsabiFilterer) FilterSharesSet(opts *bind.FilterOpts) (*ImpactcardsabiSharesSetIterator, error) {

	logs, sub, err := _Impactcardsabi.contract.FilterLogs(opts, "SharesSet")
	if err != nil {
		return nil, err
	}
	return &ImpactcardsabiSharesSetIterator{contract: _Impactcardsabi.contract, event: "SharesSet", logs: logs, sub: sub}, nil
}

// WatchSharesSet is a free log subscription operation binding the contract event 0xde702a3afda93461efb92f10ed3cdd0d0df3d292af2fc9dd64628e5f31a1901d.
//
// Solidity: event SharesSet(uint256 tokenId, uint256[2] shares)
func (_Impactcardsabi *ImpactcardsabiFilterer) WatchSharesSet(opts *bind.WatchOpts, sink chan<- *ImpactcardsabiSharesSet) (event.Subscription, error) {

	logs, sub, err := _Impactcardsabi.contract.WatchLogs(opts, "SharesSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImpactcardsabiSharesSet)
				if err := _Impactcardsabi.contract.UnpackLog(event, "SharesSet", log); err != nil {
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

// ParseSharesSet is a log parse operation binding the contract event 0xde702a3afda93461efb92f10ed3cdd0d0df3d292af2fc9dd64628e5f31a1901d.
//
// Solidity: event SharesSet(uint256 tokenId, uint256[2] shares)
func (_Impactcardsabi *ImpactcardsabiFilterer) ParseSharesSet(log types.Log) (*ImpactcardsabiSharesSet, error) {
	event := new(ImpactcardsabiSharesSet)
	if err := _Impactcardsabi.contract.UnpackLog(event, "SharesSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ImpactcardsabiTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the Impactcardsabi contract.
type ImpactcardsabiTransferBatchIterator struct {
	Event *ImpactcardsabiTransferBatch // Event containing the contract specifics and raw log

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
func (it *ImpactcardsabiTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImpactcardsabiTransferBatch)
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
		it.Event = new(ImpactcardsabiTransferBatch)
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
func (it *ImpactcardsabiTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImpactcardsabiTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImpactcardsabiTransferBatch represents a TransferBatch event raised by the Impactcardsabi contract.
type ImpactcardsabiTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Impactcardsabi *ImpactcardsabiFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ImpactcardsabiTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Impactcardsabi.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ImpactcardsabiTransferBatchIterator{contract: _Impactcardsabi.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Impactcardsabi *ImpactcardsabiFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *ImpactcardsabiTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Impactcardsabi.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImpactcardsabiTransferBatch)
				if err := _Impactcardsabi.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Impactcardsabi *ImpactcardsabiFilterer) ParseTransferBatch(log types.Log) (*ImpactcardsabiTransferBatch, error) {
	event := new(ImpactcardsabiTransferBatch)
	if err := _Impactcardsabi.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ImpactcardsabiTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the Impactcardsabi contract.
type ImpactcardsabiTransferSingleIterator struct {
	Event *ImpactcardsabiTransferSingle // Event containing the contract specifics and raw log

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
func (it *ImpactcardsabiTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImpactcardsabiTransferSingle)
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
		it.Event = new(ImpactcardsabiTransferSingle)
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
func (it *ImpactcardsabiTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImpactcardsabiTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImpactcardsabiTransferSingle represents a TransferSingle event raised by the Impactcardsabi contract.
type ImpactcardsabiTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Impactcardsabi *ImpactcardsabiFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ImpactcardsabiTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Impactcardsabi.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ImpactcardsabiTransferSingleIterator{contract: _Impactcardsabi.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Impactcardsabi *ImpactcardsabiFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *ImpactcardsabiTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Impactcardsabi.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImpactcardsabiTransferSingle)
				if err := _Impactcardsabi.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Impactcardsabi *ImpactcardsabiFilterer) ParseTransferSingle(log types.Log) (*ImpactcardsabiTransferSingle, error) {
	event := new(ImpactcardsabiTransferSingle)
	if err := _Impactcardsabi.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ImpactcardsabiURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the Impactcardsabi contract.
type ImpactcardsabiURIIterator struct {
	Event *ImpactcardsabiURI // Event containing the contract specifics and raw log

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
func (it *ImpactcardsabiURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImpactcardsabiURI)
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
		it.Event = new(ImpactcardsabiURI)
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
func (it *ImpactcardsabiURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImpactcardsabiURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImpactcardsabiURI represents a URI event raised by the Impactcardsabi contract.
type ImpactcardsabiURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Impactcardsabi *ImpactcardsabiFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*ImpactcardsabiURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Impactcardsabi.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &ImpactcardsabiURIIterator{contract: _Impactcardsabi.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Impactcardsabi *ImpactcardsabiFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *ImpactcardsabiURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Impactcardsabi.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImpactcardsabiURI)
				if err := _Impactcardsabi.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Impactcardsabi *ImpactcardsabiFilterer) ParseURI(log types.Log) (*ImpactcardsabiURI, error) {
	event := new(ImpactcardsabiURI)
	if err := _Impactcardsabi.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
