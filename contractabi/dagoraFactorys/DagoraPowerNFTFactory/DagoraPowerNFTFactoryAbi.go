// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dagorapowernft

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

// DagorapowernftMetaData contains all meta data concerning the Dagorapowernft contract.
var DagorapowernftMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DagoraFactory__ExpiredMembership\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DagoraFactory__FailedToCreateContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"tier\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"neededTier\",\"type\":\"uint8\"}],\"name\":\"DagoraFactory__InvalidTier\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DagoraFactory__NotDAgoraMembershipsOwnerOrDelegate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DagoraFactory__TokenCreationPaused\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newContractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ownerOF\",\"type\":\"address\"}],\"name\":\"PowerNFTACreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"contractsDeployed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"baseURI_\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"_bulkBuyLimit\",\"type\":\"uint16\"},{\"internalType\":\"uint96\",\"name\":\"_royaltyBps\",\"type\":\"uint96\"},{\"internalType\":\"uint256\",\"name\":\"_mintCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_royaltyRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"createPowerNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dAgoraMembershipsAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserContracts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dAgoraMembershipsAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minPowerNFTATier\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"powerNFtAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_minPowerNFTATier\",\"type\":\"uint8\"}],\"name\":\"setMinPowerNFTATier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"togglePaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userContracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DagorapowernftABI is the input ABI used to generate the binding from.
// Deprecated: Use DagorapowernftMetaData.ABI instead.
var DagorapowernftABI = DagorapowernftMetaData.ABI

// Dagorapowernft is an auto generated Go binding around an Ethereum contract.
type Dagorapowernft struct {
	DagorapowernftCaller     // Read-only binding to the contract
	DagorapowernftTransactor // Write-only binding to the contract
	DagorapowernftFilterer   // Log filterer for contract events
}

// DagorapowernftCaller is an auto generated read-only Go binding around an Ethereum contract.
type DagorapowernftCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DagorapowernftTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DagorapowernftTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DagorapowernftFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DagorapowernftFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DagorapowernftSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DagorapowernftSession struct {
	Contract     *Dagorapowernft   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DagorapowernftCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DagorapowernftCallerSession struct {
	Contract *DagorapowernftCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// DagorapowernftTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DagorapowernftTransactorSession struct {
	Contract     *DagorapowernftTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// DagorapowernftRaw is an auto generated low-level Go binding around an Ethereum contract.
type DagorapowernftRaw struct {
	Contract *Dagorapowernft // Generic contract binding to access the raw methods on
}

// DagorapowernftCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DagorapowernftCallerRaw struct {
	Contract *DagorapowernftCaller // Generic read-only contract binding to access the raw methods on
}

// DagorapowernftTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DagorapowernftTransactorRaw struct {
	Contract *DagorapowernftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDagorapowernft creates a new instance of Dagorapowernft, bound to a specific deployed contract.
func NewDagorapowernft(address common.Address, backend bind.ContractBackend) (*Dagorapowernft, error) {
	contract, err := bindDagorapowernft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dagorapowernft{DagorapowernftCaller: DagorapowernftCaller{contract: contract}, DagorapowernftTransactor: DagorapowernftTransactor{contract: contract}, DagorapowernftFilterer: DagorapowernftFilterer{contract: contract}}, nil
}

// NewDagorapowernftCaller creates a new read-only instance of Dagorapowernft, bound to a specific deployed contract.
func NewDagorapowernftCaller(address common.Address, caller bind.ContractCaller) (*DagorapowernftCaller, error) {
	contract, err := bindDagorapowernft(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DagorapowernftCaller{contract: contract}, nil
}

// NewDagorapowernftTransactor creates a new write-only instance of Dagorapowernft, bound to a specific deployed contract.
func NewDagorapowernftTransactor(address common.Address, transactor bind.ContractTransactor) (*DagorapowernftTransactor, error) {
	contract, err := bindDagorapowernft(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DagorapowernftTransactor{contract: contract}, nil
}

// NewDagorapowernftFilterer creates a new log filterer instance of Dagorapowernft, bound to a specific deployed contract.
func NewDagorapowernftFilterer(address common.Address, filterer bind.ContractFilterer) (*DagorapowernftFilterer, error) {
	contract, err := bindDagorapowernft(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DagorapowernftFilterer{contract: contract}, nil
}

// bindDagorapowernft binds a generic wrapper to an already deployed contract.
func bindDagorapowernft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DagorapowernftMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dagorapowernft *DagorapowernftRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dagorapowernft.Contract.DagorapowernftCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dagorapowernft *DagorapowernftRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dagorapowernft.Contract.DagorapowernftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dagorapowernft *DagorapowernftRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dagorapowernft.Contract.DagorapowernftTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dagorapowernft *DagorapowernftCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dagorapowernft.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dagorapowernft *DagorapowernftTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dagorapowernft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dagorapowernft *DagorapowernftTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dagorapowernft.Contract.contract.Transact(opts, method, params...)
}

// ContractsDeployed is a free data retrieval call binding the contract method 0xf0a761e4.
//
// Solidity: function contractsDeployed() view returns(uint256)
func (_Dagorapowernft *DagorapowernftCaller) ContractsDeployed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dagorapowernft.contract.Call(opts, &out, "contractsDeployed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ContractsDeployed is a free data retrieval call binding the contract method 0xf0a761e4.
//
// Solidity: function contractsDeployed() view returns(uint256)
func (_Dagorapowernft *DagorapowernftSession) ContractsDeployed() (*big.Int, error) {
	return _Dagorapowernft.Contract.ContractsDeployed(&_Dagorapowernft.CallOpts)
}

// ContractsDeployed is a free data retrieval call binding the contract method 0xf0a761e4.
//
// Solidity: function contractsDeployed() view returns(uint256)
func (_Dagorapowernft *DagorapowernftCallerSession) ContractsDeployed() (*big.Int, error) {
	return _Dagorapowernft.Contract.ContractsDeployed(&_Dagorapowernft.CallOpts)
}

// DAgoraMembershipsAddress is a free data retrieval call binding the contract method 0x514fd254.
//
// Solidity: function dAgoraMembershipsAddress() view returns(address)
func (_Dagorapowernft *DagorapowernftCaller) DAgoraMembershipsAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dagorapowernft.contract.Call(opts, &out, "dAgoraMembershipsAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DAgoraMembershipsAddress is a free data retrieval call binding the contract method 0x514fd254.
//
// Solidity: function dAgoraMembershipsAddress() view returns(address)
func (_Dagorapowernft *DagorapowernftSession) DAgoraMembershipsAddress() (common.Address, error) {
	return _Dagorapowernft.Contract.DAgoraMembershipsAddress(&_Dagorapowernft.CallOpts)
}

// DAgoraMembershipsAddress is a free data retrieval call binding the contract method 0x514fd254.
//
// Solidity: function dAgoraMembershipsAddress() view returns(address)
func (_Dagorapowernft *DagorapowernftCallerSession) DAgoraMembershipsAddress() (common.Address, error) {
	return _Dagorapowernft.Contract.DAgoraMembershipsAddress(&_Dagorapowernft.CallOpts)
}

// GetUserContracts is a free data retrieval call binding the contract method 0x34c925f0.
//
// Solidity: function getUserContracts(address _user) view returns(address[])
func (_Dagorapowernft *DagorapowernftCaller) GetUserContracts(opts *bind.CallOpts, _user common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Dagorapowernft.contract.Call(opts, &out, "getUserContracts", _user)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetUserContracts is a free data retrieval call binding the contract method 0x34c925f0.
//
// Solidity: function getUserContracts(address _user) view returns(address[])
func (_Dagorapowernft *DagorapowernftSession) GetUserContracts(_user common.Address) ([]common.Address, error) {
	return _Dagorapowernft.Contract.GetUserContracts(&_Dagorapowernft.CallOpts, _user)
}

// GetUserContracts is a free data retrieval call binding the contract method 0x34c925f0.
//
// Solidity: function getUserContracts(address _user) view returns(address[])
func (_Dagorapowernft *DagorapowernftCallerSession) GetUserContracts(_user common.Address) ([]common.Address, error) {
	return _Dagorapowernft.Contract.GetUserContracts(&_Dagorapowernft.CallOpts, _user)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Dagorapowernft *DagorapowernftCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Dagorapowernft.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Dagorapowernft *DagorapowernftSession) IsPaused() (bool, error) {
	return _Dagorapowernft.Contract.IsPaused(&_Dagorapowernft.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Dagorapowernft *DagorapowernftCallerSession) IsPaused() (bool, error) {
	return _Dagorapowernft.Contract.IsPaused(&_Dagorapowernft.CallOpts)
}

// MinPowerNFTATier is a free data retrieval call binding the contract method 0xfdd17a4f.
//
// Solidity: function minPowerNFTATier() view returns(uint8)
func (_Dagorapowernft *DagorapowernftCaller) MinPowerNFTATier(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Dagorapowernft.contract.Call(opts, &out, "minPowerNFTATier")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MinPowerNFTATier is a free data retrieval call binding the contract method 0xfdd17a4f.
//
// Solidity: function minPowerNFTATier() view returns(uint8)
func (_Dagorapowernft *DagorapowernftSession) MinPowerNFTATier() (uint8, error) {
	return _Dagorapowernft.Contract.MinPowerNFTATier(&_Dagorapowernft.CallOpts)
}

// MinPowerNFTATier is a free data retrieval call binding the contract method 0xfdd17a4f.
//
// Solidity: function minPowerNFTATier() view returns(uint8)
func (_Dagorapowernft *DagorapowernftCallerSession) MinPowerNFTATier() (uint8, error) {
	return _Dagorapowernft.Contract.MinPowerNFTATier(&_Dagorapowernft.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dagorapowernft *DagorapowernftCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dagorapowernft.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dagorapowernft *DagorapowernftSession) Owner() (common.Address, error) {
	return _Dagorapowernft.Contract.Owner(&_Dagorapowernft.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dagorapowernft *DagorapowernftCallerSession) Owner() (common.Address, error) {
	return _Dagorapowernft.Contract.Owner(&_Dagorapowernft.CallOpts)
}

// PowerNFtAddress is a free data retrieval call binding the contract method 0xee3e08b3.
//
// Solidity: function powerNFtAddress() view returns(address)
func (_Dagorapowernft *DagorapowernftCaller) PowerNFtAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dagorapowernft.contract.Call(opts, &out, "powerNFtAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PowerNFtAddress is a free data retrieval call binding the contract method 0xee3e08b3.
//
// Solidity: function powerNFtAddress() view returns(address)
func (_Dagorapowernft *DagorapowernftSession) PowerNFtAddress() (common.Address, error) {
	return _Dagorapowernft.Contract.PowerNFtAddress(&_Dagorapowernft.CallOpts)
}

// PowerNFtAddress is a free data retrieval call binding the contract method 0xee3e08b3.
//
// Solidity: function powerNFtAddress() view returns(address)
func (_Dagorapowernft *DagorapowernftCallerSession) PowerNFtAddress() (common.Address, error) {
	return _Dagorapowernft.Contract.PowerNFtAddress(&_Dagorapowernft.CallOpts)
}

// UserContracts is a free data retrieval call binding the contract method 0xb789003d.
//
// Solidity: function userContracts(address , uint256 ) view returns(address)
func (_Dagorapowernft *DagorapowernftCaller) UserContracts(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Dagorapowernft.contract.Call(opts, &out, "userContracts", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserContracts is a free data retrieval call binding the contract method 0xb789003d.
//
// Solidity: function userContracts(address , uint256 ) view returns(address)
func (_Dagorapowernft *DagorapowernftSession) UserContracts(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Dagorapowernft.Contract.UserContracts(&_Dagorapowernft.CallOpts, arg0, arg1)
}

// UserContracts is a free data retrieval call binding the contract method 0xb789003d.
//
// Solidity: function userContracts(address , uint256 ) view returns(address)
func (_Dagorapowernft *DagorapowernftCallerSession) UserContracts(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Dagorapowernft.Contract.UserContracts(&_Dagorapowernft.CallOpts, arg0, arg1)
}

// CreatePowerNFT is a paid mutator transaction binding the contract method 0xf4aef580.
//
// Solidity: function createPowerNFT(string name_, string symbol_, string baseURI_, uint16 _bulkBuyLimit, uint96 _royaltyBps, uint256 _mintCost, uint256 _maxSupply, address _royaltyRecipient, address _newOwner, uint256 _id) returns()
func (_Dagorapowernft *DagorapowernftTransactor) CreatePowerNFT(opts *bind.TransactOpts, name_ string, symbol_ string, baseURI_ string, _bulkBuyLimit uint16, _royaltyBps *big.Int, _mintCost *big.Int, _maxSupply *big.Int, _royaltyRecipient common.Address, _newOwner common.Address, _id *big.Int) (*types.Transaction, error) {
	return _Dagorapowernft.contract.Transact(opts, "createPowerNFT", name_, symbol_, baseURI_, _bulkBuyLimit, _royaltyBps, _mintCost, _maxSupply, _royaltyRecipient, _newOwner, _id)
}

// CreatePowerNFT is a paid mutator transaction binding the contract method 0xf4aef580.
//
// Solidity: function createPowerNFT(string name_, string symbol_, string baseURI_, uint16 _bulkBuyLimit, uint96 _royaltyBps, uint256 _mintCost, uint256 _maxSupply, address _royaltyRecipient, address _newOwner, uint256 _id) returns()
func (_Dagorapowernft *DagorapowernftSession) CreatePowerNFT(name_ string, symbol_ string, baseURI_ string, _bulkBuyLimit uint16, _royaltyBps *big.Int, _mintCost *big.Int, _maxSupply *big.Int, _royaltyRecipient common.Address, _newOwner common.Address, _id *big.Int) (*types.Transaction, error) {
	return _Dagorapowernft.Contract.CreatePowerNFT(&_Dagorapowernft.TransactOpts, name_, symbol_, baseURI_, _bulkBuyLimit, _royaltyBps, _mintCost, _maxSupply, _royaltyRecipient, _newOwner, _id)
}

// CreatePowerNFT is a paid mutator transaction binding the contract method 0xf4aef580.
//
// Solidity: function createPowerNFT(string name_, string symbol_, string baseURI_, uint16 _bulkBuyLimit, uint96 _royaltyBps, uint256 _mintCost, uint256 _maxSupply, address _royaltyRecipient, address _newOwner, uint256 _id) returns()
func (_Dagorapowernft *DagorapowernftTransactorSession) CreatePowerNFT(name_ string, symbol_ string, baseURI_ string, _bulkBuyLimit uint16, _royaltyBps *big.Int, _mintCost *big.Int, _maxSupply *big.Int, _royaltyRecipient common.Address, _newOwner common.Address, _id *big.Int) (*types.Transaction, error) {
	return _Dagorapowernft.Contract.CreatePowerNFT(&_Dagorapowernft.TransactOpts, name_, symbol_, baseURI_, _bulkBuyLimit, _royaltyBps, _mintCost, _maxSupply, _royaltyRecipient, _newOwner, _id)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _dAgoraMembershipsAddress) returns()
func (_Dagorapowernft *DagorapowernftTransactor) Initialize(opts *bind.TransactOpts, _dAgoraMembershipsAddress common.Address) (*types.Transaction, error) {
	return _Dagorapowernft.contract.Transact(opts, "initialize", _dAgoraMembershipsAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _dAgoraMembershipsAddress) returns()
func (_Dagorapowernft *DagorapowernftSession) Initialize(_dAgoraMembershipsAddress common.Address) (*types.Transaction, error) {
	return _Dagorapowernft.Contract.Initialize(&_Dagorapowernft.TransactOpts, _dAgoraMembershipsAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _dAgoraMembershipsAddress) returns()
func (_Dagorapowernft *DagorapowernftTransactorSession) Initialize(_dAgoraMembershipsAddress common.Address) (*types.Transaction, error) {
	return _Dagorapowernft.Contract.Initialize(&_Dagorapowernft.TransactOpts, _dAgoraMembershipsAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dagorapowernft *DagorapowernftTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dagorapowernft.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dagorapowernft *DagorapowernftSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dagorapowernft.Contract.RenounceOwnership(&_Dagorapowernft.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dagorapowernft *DagorapowernftTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dagorapowernft.Contract.RenounceOwnership(&_Dagorapowernft.TransactOpts)
}

// SetMinPowerNFTATier is a paid mutator transaction binding the contract method 0x04bad475.
//
// Solidity: function setMinPowerNFTATier(uint8 _minPowerNFTATier) returns()
func (_Dagorapowernft *DagorapowernftTransactor) SetMinPowerNFTATier(opts *bind.TransactOpts, _minPowerNFTATier uint8) (*types.Transaction, error) {
	return _Dagorapowernft.contract.Transact(opts, "setMinPowerNFTATier", _minPowerNFTATier)
}

// SetMinPowerNFTATier is a paid mutator transaction binding the contract method 0x04bad475.
//
// Solidity: function setMinPowerNFTATier(uint8 _minPowerNFTATier) returns()
func (_Dagorapowernft *DagorapowernftSession) SetMinPowerNFTATier(_minPowerNFTATier uint8) (*types.Transaction, error) {
	return _Dagorapowernft.Contract.SetMinPowerNFTATier(&_Dagorapowernft.TransactOpts, _minPowerNFTATier)
}

// SetMinPowerNFTATier is a paid mutator transaction binding the contract method 0x04bad475.
//
// Solidity: function setMinPowerNFTATier(uint8 _minPowerNFTATier) returns()
func (_Dagorapowernft *DagorapowernftTransactorSession) SetMinPowerNFTATier(_minPowerNFTATier uint8) (*types.Transaction, error) {
	return _Dagorapowernft.Contract.SetMinPowerNFTATier(&_Dagorapowernft.TransactOpts, _minPowerNFTATier)
}

// TogglePaused is a paid mutator transaction binding the contract method 0x36566f06.
//
// Solidity: function togglePaused() returns()
func (_Dagorapowernft *DagorapowernftTransactor) TogglePaused(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dagorapowernft.contract.Transact(opts, "togglePaused")
}

// TogglePaused is a paid mutator transaction binding the contract method 0x36566f06.
//
// Solidity: function togglePaused() returns()
func (_Dagorapowernft *DagorapowernftSession) TogglePaused() (*types.Transaction, error) {
	return _Dagorapowernft.Contract.TogglePaused(&_Dagorapowernft.TransactOpts)
}

// TogglePaused is a paid mutator transaction binding the contract method 0x36566f06.
//
// Solidity: function togglePaused() returns()
func (_Dagorapowernft *DagorapowernftTransactorSession) TogglePaused() (*types.Transaction, error) {
	return _Dagorapowernft.Contract.TogglePaused(&_Dagorapowernft.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dagorapowernft *DagorapowernftTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Dagorapowernft.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dagorapowernft *DagorapowernftSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dagorapowernft.Contract.TransferOwnership(&_Dagorapowernft.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dagorapowernft *DagorapowernftTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dagorapowernft.Contract.TransferOwnership(&_Dagorapowernft.TransactOpts, newOwner)
}

// DagorapowernftInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Dagorapowernft contract.
type DagorapowernftInitializedIterator struct {
	Event *DagorapowernftInitialized // Event containing the contract specifics and raw log

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
func (it *DagorapowernftInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DagorapowernftInitialized)
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
		it.Event = new(DagorapowernftInitialized)
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
func (it *DagorapowernftInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DagorapowernftInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DagorapowernftInitialized represents a Initialized event raised by the Dagorapowernft contract.
type DagorapowernftInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Dagorapowernft *DagorapowernftFilterer) FilterInitialized(opts *bind.FilterOpts) (*DagorapowernftInitializedIterator, error) {

	logs, sub, err := _Dagorapowernft.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DagorapowernftInitializedIterator{contract: _Dagorapowernft.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Dagorapowernft *DagorapowernftFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DagorapowernftInitialized) (event.Subscription, error) {

	logs, sub, err := _Dagorapowernft.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DagorapowernftInitialized)
				if err := _Dagorapowernft.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Dagorapowernft *DagorapowernftFilterer) ParseInitialized(log types.Log) (*DagorapowernftInitialized, error) {
	event := new(DagorapowernftInitialized)
	if err := _Dagorapowernft.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DagorapowernftOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Dagorapowernft contract.
type DagorapowernftOwnershipTransferredIterator struct {
	Event *DagorapowernftOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DagorapowernftOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DagorapowernftOwnershipTransferred)
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
		it.Event = new(DagorapowernftOwnershipTransferred)
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
func (it *DagorapowernftOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DagorapowernftOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DagorapowernftOwnershipTransferred represents a OwnershipTransferred event raised by the Dagorapowernft contract.
type DagorapowernftOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dagorapowernft *DagorapowernftFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DagorapowernftOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dagorapowernft.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DagorapowernftOwnershipTransferredIterator{contract: _Dagorapowernft.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dagorapowernft *DagorapowernftFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DagorapowernftOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dagorapowernft.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DagorapowernftOwnershipTransferred)
				if err := _Dagorapowernft.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Dagorapowernft *DagorapowernftFilterer) ParseOwnershipTransferred(log types.Log) (*DagorapowernftOwnershipTransferred, error) {
	event := new(DagorapowernftOwnershipTransferred)
	if err := _Dagorapowernft.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DagorapowernftPowerNFTACreatedIterator is returned from FilterPowerNFTACreated and is used to iterate over the raw logs and unpacked data for PowerNFTACreated events raised by the Dagorapowernft contract.
type DagorapowernftPowerNFTACreatedIterator struct {
	Event *DagorapowernftPowerNFTACreated // Event containing the contract specifics and raw log

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
func (it *DagorapowernftPowerNFTACreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DagorapowernftPowerNFTACreated)
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
		it.Event = new(DagorapowernftPowerNFTACreated)
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
func (it *DagorapowernftPowerNFTACreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DagorapowernftPowerNFTACreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DagorapowernftPowerNFTACreated represents a PowerNFTACreated event raised by the Dagorapowernft contract.
type DagorapowernftPowerNFTACreated struct {
	NewContractAddress common.Address
	OwnerOF            common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPowerNFTACreated is a free log retrieval operation binding the contract event 0x96c23507d82fce99fdde9536b40a916988652da91d270af7dbd9968b35358f78.
//
// Solidity: event PowerNFTACreated(address newContractAddress, address ownerOF)
func (_Dagorapowernft *DagorapowernftFilterer) FilterPowerNFTACreated(opts *bind.FilterOpts) (*DagorapowernftPowerNFTACreatedIterator, error) {

	logs, sub, err := _Dagorapowernft.contract.FilterLogs(opts, "PowerNFTACreated")
	if err != nil {
		return nil, err
	}
	return &DagorapowernftPowerNFTACreatedIterator{contract: _Dagorapowernft.contract, event: "PowerNFTACreated", logs: logs, sub: sub}, nil
}

// WatchPowerNFTACreated is a free log subscription operation binding the contract event 0x96c23507d82fce99fdde9536b40a916988652da91d270af7dbd9968b35358f78.
//
// Solidity: event PowerNFTACreated(address newContractAddress, address ownerOF)
func (_Dagorapowernft *DagorapowernftFilterer) WatchPowerNFTACreated(opts *bind.WatchOpts, sink chan<- *DagorapowernftPowerNFTACreated) (event.Subscription, error) {

	logs, sub, err := _Dagorapowernft.contract.WatchLogs(opts, "PowerNFTACreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DagorapowernftPowerNFTACreated)
				if err := _Dagorapowernft.contract.UnpackLog(event, "PowerNFTACreated", log); err != nil {
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

// ParsePowerNFTACreated is a log parse operation binding the contract event 0x96c23507d82fce99fdde9536b40a916988652da91d270af7dbd9968b35358f78.
//
// Solidity: event PowerNFTACreated(address newContractAddress, address ownerOF)
func (_Dagorapowernft *DagorapowernftFilterer) ParsePowerNFTACreated(log types.Log) (*DagorapowernftPowerNFTACreated, error) {
	event := new(DagorapowernftPowerNFTACreated)
	if err := _Dagorapowernft.contract.UnpackLog(event, "PowerNFTACreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
