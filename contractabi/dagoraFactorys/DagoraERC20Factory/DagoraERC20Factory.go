// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dagoraerc20factory

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

// Dagoraerc20factoryMetaData contains all meta data concerning the Dagoraerc20factory contract.
var Dagoraerc20factoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DagoraFactory__ExpiredMembership\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DagoraFactory__FailedToCreateContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"tier\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"neededTier\",\"type\":\"uint8\"}],\"name\":\"DagoraFactory__InvalidTier\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DagoraFactory__NotDAgoraMembershipsOwnerOrDelegate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DagoraFactory__TokenCreationPaused\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newContractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ownerOF\",\"type\":\"address\"}],\"name\":\"DagoraERC20Created\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"contractsDeployed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_initialSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"createDagoraERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dAgoraMembershipsAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserContracts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dAgoraMembershipsAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minERC20Tier\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_minTier\",\"type\":\"uint8\"}],\"name\":\"setMinERC20Tier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"togglePaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userContracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Dagoraerc20factoryABI is the input ABI used to generate the binding from.
// Deprecated: Use Dagoraerc20factoryMetaData.ABI instead.
var Dagoraerc20factoryABI = Dagoraerc20factoryMetaData.ABI

// Dagoraerc20factory is an auto generated Go binding around an Ethereum contract.
type Dagoraerc20factory struct {
	Dagoraerc20factoryCaller     // Read-only binding to the contract
	Dagoraerc20factoryTransactor // Write-only binding to the contract
	Dagoraerc20factoryFilterer   // Log filterer for contract events
}

// Dagoraerc20factoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type Dagoraerc20factoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Dagoraerc20factoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Dagoraerc20factoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Dagoraerc20factoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Dagoraerc20factoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Dagoraerc20factorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Dagoraerc20factorySession struct {
	Contract     *Dagoraerc20factory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// Dagoraerc20factoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Dagoraerc20factoryCallerSession struct {
	Contract *Dagoraerc20factoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// Dagoraerc20factoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Dagoraerc20factoryTransactorSession struct {
	Contract     *Dagoraerc20factoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// Dagoraerc20factoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type Dagoraerc20factoryRaw struct {
	Contract *Dagoraerc20factory // Generic contract binding to access the raw methods on
}

// Dagoraerc20factoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Dagoraerc20factoryCallerRaw struct {
	Contract *Dagoraerc20factoryCaller // Generic read-only contract binding to access the raw methods on
}

// Dagoraerc20factoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Dagoraerc20factoryTransactorRaw struct {
	Contract *Dagoraerc20factoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDagoraerc20factory creates a new instance of Dagoraerc20factory, bound to a specific deployed contract.
func NewDagoraerc20factory(address common.Address, backend bind.ContractBackend) (*Dagoraerc20factory, error) {
	contract, err := bindDagoraerc20factory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dagoraerc20factory{Dagoraerc20factoryCaller: Dagoraerc20factoryCaller{contract: contract}, Dagoraerc20factoryTransactor: Dagoraerc20factoryTransactor{contract: contract}, Dagoraerc20factoryFilterer: Dagoraerc20factoryFilterer{contract: contract}}, nil
}

// NewDagoraerc20factoryCaller creates a new read-only instance of Dagoraerc20factory, bound to a specific deployed contract.
func NewDagoraerc20factoryCaller(address common.Address, caller bind.ContractCaller) (*Dagoraerc20factoryCaller, error) {
	contract, err := bindDagoraerc20factory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Dagoraerc20factoryCaller{contract: contract}, nil
}

// NewDagoraerc20factoryTransactor creates a new write-only instance of Dagoraerc20factory, bound to a specific deployed contract.
func NewDagoraerc20factoryTransactor(address common.Address, transactor bind.ContractTransactor) (*Dagoraerc20factoryTransactor, error) {
	contract, err := bindDagoraerc20factory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Dagoraerc20factoryTransactor{contract: contract}, nil
}

// NewDagoraerc20factoryFilterer creates a new log filterer instance of Dagoraerc20factory, bound to a specific deployed contract.
func NewDagoraerc20factoryFilterer(address common.Address, filterer bind.ContractFilterer) (*Dagoraerc20factoryFilterer, error) {
	contract, err := bindDagoraerc20factory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Dagoraerc20factoryFilterer{contract: contract}, nil
}

// bindDagoraerc20factory binds a generic wrapper to an already deployed contract.
func bindDagoraerc20factory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Dagoraerc20factoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dagoraerc20factory *Dagoraerc20factoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dagoraerc20factory.Contract.Dagoraerc20factoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dagoraerc20factory *Dagoraerc20factoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dagoraerc20factory.Contract.Dagoraerc20factoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dagoraerc20factory *Dagoraerc20factoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dagoraerc20factory.Contract.Dagoraerc20factoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dagoraerc20factory *Dagoraerc20factoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dagoraerc20factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dagoraerc20factory *Dagoraerc20factoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dagoraerc20factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dagoraerc20factory *Dagoraerc20factoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dagoraerc20factory.Contract.contract.Transact(opts, method, params...)
}

// ContractsDeployed is a free data retrieval call binding the contract method 0xf0a761e4.
//
// Solidity: function contractsDeployed() view returns(uint256)
func (_Dagoraerc20factory *Dagoraerc20factoryCaller) ContractsDeployed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dagoraerc20factory.contract.Call(opts, &out, "contractsDeployed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ContractsDeployed is a free data retrieval call binding the contract method 0xf0a761e4.
//
// Solidity: function contractsDeployed() view returns(uint256)
func (_Dagoraerc20factory *Dagoraerc20factorySession) ContractsDeployed() (*big.Int, error) {
	return _Dagoraerc20factory.Contract.ContractsDeployed(&_Dagoraerc20factory.CallOpts)
}

// ContractsDeployed is a free data retrieval call binding the contract method 0xf0a761e4.
//
// Solidity: function contractsDeployed() view returns(uint256)
func (_Dagoraerc20factory *Dagoraerc20factoryCallerSession) ContractsDeployed() (*big.Int, error) {
	return _Dagoraerc20factory.Contract.ContractsDeployed(&_Dagoraerc20factory.CallOpts)
}

// DAgoraMembershipsAddress is a free data retrieval call binding the contract method 0x514fd254.
//
// Solidity: function dAgoraMembershipsAddress() view returns(address)
func (_Dagoraerc20factory *Dagoraerc20factoryCaller) DAgoraMembershipsAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dagoraerc20factory.contract.Call(opts, &out, "dAgoraMembershipsAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DAgoraMembershipsAddress is a free data retrieval call binding the contract method 0x514fd254.
//
// Solidity: function dAgoraMembershipsAddress() view returns(address)
func (_Dagoraerc20factory *Dagoraerc20factorySession) DAgoraMembershipsAddress() (common.Address, error) {
	return _Dagoraerc20factory.Contract.DAgoraMembershipsAddress(&_Dagoraerc20factory.CallOpts)
}

// DAgoraMembershipsAddress is a free data retrieval call binding the contract method 0x514fd254.
//
// Solidity: function dAgoraMembershipsAddress() view returns(address)
func (_Dagoraerc20factory *Dagoraerc20factoryCallerSession) DAgoraMembershipsAddress() (common.Address, error) {
	return _Dagoraerc20factory.Contract.DAgoraMembershipsAddress(&_Dagoraerc20factory.CallOpts)
}

// GetUserContracts is a free data retrieval call binding the contract method 0x34c925f0.
//
// Solidity: function getUserContracts(address _user) view returns(address[])
func (_Dagoraerc20factory *Dagoraerc20factoryCaller) GetUserContracts(opts *bind.CallOpts, _user common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Dagoraerc20factory.contract.Call(opts, &out, "getUserContracts", _user)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetUserContracts is a free data retrieval call binding the contract method 0x34c925f0.
//
// Solidity: function getUserContracts(address _user) view returns(address[])
func (_Dagoraerc20factory *Dagoraerc20factorySession) GetUserContracts(_user common.Address) ([]common.Address, error) {
	return _Dagoraerc20factory.Contract.GetUserContracts(&_Dagoraerc20factory.CallOpts, _user)
}

// GetUserContracts is a free data retrieval call binding the contract method 0x34c925f0.
//
// Solidity: function getUserContracts(address _user) view returns(address[])
func (_Dagoraerc20factory *Dagoraerc20factoryCallerSession) GetUserContracts(_user common.Address) ([]common.Address, error) {
	return _Dagoraerc20factory.Contract.GetUserContracts(&_Dagoraerc20factory.CallOpts, _user)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Dagoraerc20factory *Dagoraerc20factoryCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Dagoraerc20factory.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Dagoraerc20factory *Dagoraerc20factorySession) IsPaused() (bool, error) {
	return _Dagoraerc20factory.Contract.IsPaused(&_Dagoraerc20factory.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Dagoraerc20factory *Dagoraerc20factoryCallerSession) IsPaused() (bool, error) {
	return _Dagoraerc20factory.Contract.IsPaused(&_Dagoraerc20factory.CallOpts)
}

// MinERC20Tier is a free data retrieval call binding the contract method 0x45044572.
//
// Solidity: function minERC20Tier() view returns(uint8)
func (_Dagoraerc20factory *Dagoraerc20factoryCaller) MinERC20Tier(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Dagoraerc20factory.contract.Call(opts, &out, "minERC20Tier")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MinERC20Tier is a free data retrieval call binding the contract method 0x45044572.
//
// Solidity: function minERC20Tier() view returns(uint8)
func (_Dagoraerc20factory *Dagoraerc20factorySession) MinERC20Tier() (uint8, error) {
	return _Dagoraerc20factory.Contract.MinERC20Tier(&_Dagoraerc20factory.CallOpts)
}

// MinERC20Tier is a free data retrieval call binding the contract method 0x45044572.
//
// Solidity: function minERC20Tier() view returns(uint8)
func (_Dagoraerc20factory *Dagoraerc20factoryCallerSession) MinERC20Tier() (uint8, error) {
	return _Dagoraerc20factory.Contract.MinERC20Tier(&_Dagoraerc20factory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dagoraerc20factory *Dagoraerc20factoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dagoraerc20factory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dagoraerc20factory *Dagoraerc20factorySession) Owner() (common.Address, error) {
	return _Dagoraerc20factory.Contract.Owner(&_Dagoraerc20factory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dagoraerc20factory *Dagoraerc20factoryCallerSession) Owner() (common.Address, error) {
	return _Dagoraerc20factory.Contract.Owner(&_Dagoraerc20factory.CallOpts)
}

// UserContracts is a free data retrieval call binding the contract method 0xb789003d.
//
// Solidity: function userContracts(address , uint256 ) view returns(address)
func (_Dagoraerc20factory *Dagoraerc20factoryCaller) UserContracts(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Dagoraerc20factory.contract.Call(opts, &out, "userContracts", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserContracts is a free data retrieval call binding the contract method 0xb789003d.
//
// Solidity: function userContracts(address , uint256 ) view returns(address)
func (_Dagoraerc20factory *Dagoraerc20factorySession) UserContracts(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Dagoraerc20factory.Contract.UserContracts(&_Dagoraerc20factory.CallOpts, arg0, arg1)
}

// UserContracts is a free data retrieval call binding the contract method 0xb789003d.
//
// Solidity: function userContracts(address , uint256 ) view returns(address)
func (_Dagoraerc20factory *Dagoraerc20factoryCallerSession) UserContracts(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Dagoraerc20factory.Contract.UserContracts(&_Dagoraerc20factory.CallOpts, arg0, arg1)
}

// CreateDagoraERC20 is a paid mutator transaction binding the contract method 0xe5ece4df.
//
// Solidity: function createDagoraERC20(string name_, string symbol_, address _newOwner, uint256 _initialSupply, uint256 _maxSupply, uint256 _id) returns()
func (_Dagoraerc20factory *Dagoraerc20factoryTransactor) CreateDagoraERC20(opts *bind.TransactOpts, name_ string, symbol_ string, _newOwner common.Address, _initialSupply *big.Int, _maxSupply *big.Int, _id *big.Int) (*types.Transaction, error) {
	return _Dagoraerc20factory.contract.Transact(opts, "createDagoraERC20", name_, symbol_, _newOwner, _initialSupply, _maxSupply, _id)
}

// CreateDagoraERC20 is a paid mutator transaction binding the contract method 0xe5ece4df.
//
// Solidity: function createDagoraERC20(string name_, string symbol_, address _newOwner, uint256 _initialSupply, uint256 _maxSupply, uint256 _id) returns()
func (_Dagoraerc20factory *Dagoraerc20factorySession) CreateDagoraERC20(name_ string, symbol_ string, _newOwner common.Address, _initialSupply *big.Int, _maxSupply *big.Int, _id *big.Int) (*types.Transaction, error) {
	return _Dagoraerc20factory.Contract.CreateDagoraERC20(&_Dagoraerc20factory.TransactOpts, name_, symbol_, _newOwner, _initialSupply, _maxSupply, _id)
}

// CreateDagoraERC20 is a paid mutator transaction binding the contract method 0xe5ece4df.
//
// Solidity: function createDagoraERC20(string name_, string symbol_, address _newOwner, uint256 _initialSupply, uint256 _maxSupply, uint256 _id) returns()
func (_Dagoraerc20factory *Dagoraerc20factoryTransactorSession) CreateDagoraERC20(name_ string, symbol_ string, _newOwner common.Address, _initialSupply *big.Int, _maxSupply *big.Int, _id *big.Int) (*types.Transaction, error) {
	return _Dagoraerc20factory.Contract.CreateDagoraERC20(&_Dagoraerc20factory.TransactOpts, name_, symbol_, _newOwner, _initialSupply, _maxSupply, _id)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _dAgoraMembershipsAddress) returns()
func (_Dagoraerc20factory *Dagoraerc20factoryTransactor) Initialize(opts *bind.TransactOpts, _dAgoraMembershipsAddress common.Address) (*types.Transaction, error) {
	return _Dagoraerc20factory.contract.Transact(opts, "initialize", _dAgoraMembershipsAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _dAgoraMembershipsAddress) returns()
func (_Dagoraerc20factory *Dagoraerc20factorySession) Initialize(_dAgoraMembershipsAddress common.Address) (*types.Transaction, error) {
	return _Dagoraerc20factory.Contract.Initialize(&_Dagoraerc20factory.TransactOpts, _dAgoraMembershipsAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _dAgoraMembershipsAddress) returns()
func (_Dagoraerc20factory *Dagoraerc20factoryTransactorSession) Initialize(_dAgoraMembershipsAddress common.Address) (*types.Transaction, error) {
	return _Dagoraerc20factory.Contract.Initialize(&_Dagoraerc20factory.TransactOpts, _dAgoraMembershipsAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dagoraerc20factory *Dagoraerc20factoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dagoraerc20factory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dagoraerc20factory *Dagoraerc20factorySession) RenounceOwnership() (*types.Transaction, error) {
	return _Dagoraerc20factory.Contract.RenounceOwnership(&_Dagoraerc20factory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dagoraerc20factory *Dagoraerc20factoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dagoraerc20factory.Contract.RenounceOwnership(&_Dagoraerc20factory.TransactOpts)
}

// SetMinERC20Tier is a paid mutator transaction binding the contract method 0x78d8d424.
//
// Solidity: function setMinERC20Tier(uint8 _minTier) returns()
func (_Dagoraerc20factory *Dagoraerc20factoryTransactor) SetMinERC20Tier(opts *bind.TransactOpts, _minTier uint8) (*types.Transaction, error) {
	return _Dagoraerc20factory.contract.Transact(opts, "setMinERC20Tier", _minTier)
}

// SetMinERC20Tier is a paid mutator transaction binding the contract method 0x78d8d424.
//
// Solidity: function setMinERC20Tier(uint8 _minTier) returns()
func (_Dagoraerc20factory *Dagoraerc20factorySession) SetMinERC20Tier(_minTier uint8) (*types.Transaction, error) {
	return _Dagoraerc20factory.Contract.SetMinERC20Tier(&_Dagoraerc20factory.TransactOpts, _minTier)
}

// SetMinERC20Tier is a paid mutator transaction binding the contract method 0x78d8d424.
//
// Solidity: function setMinERC20Tier(uint8 _minTier) returns()
func (_Dagoraerc20factory *Dagoraerc20factoryTransactorSession) SetMinERC20Tier(_minTier uint8) (*types.Transaction, error) {
	return _Dagoraerc20factory.Contract.SetMinERC20Tier(&_Dagoraerc20factory.TransactOpts, _minTier)
}

// TogglePaused is a paid mutator transaction binding the contract method 0x36566f06.
//
// Solidity: function togglePaused() returns()
func (_Dagoraerc20factory *Dagoraerc20factoryTransactor) TogglePaused(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dagoraerc20factory.contract.Transact(opts, "togglePaused")
}

// TogglePaused is a paid mutator transaction binding the contract method 0x36566f06.
//
// Solidity: function togglePaused() returns()
func (_Dagoraerc20factory *Dagoraerc20factorySession) TogglePaused() (*types.Transaction, error) {
	return _Dagoraerc20factory.Contract.TogglePaused(&_Dagoraerc20factory.TransactOpts)
}

// TogglePaused is a paid mutator transaction binding the contract method 0x36566f06.
//
// Solidity: function togglePaused() returns()
func (_Dagoraerc20factory *Dagoraerc20factoryTransactorSession) TogglePaused() (*types.Transaction, error) {
	return _Dagoraerc20factory.Contract.TogglePaused(&_Dagoraerc20factory.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dagoraerc20factory *Dagoraerc20factoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Dagoraerc20factory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dagoraerc20factory *Dagoraerc20factorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dagoraerc20factory.Contract.TransferOwnership(&_Dagoraerc20factory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dagoraerc20factory *Dagoraerc20factoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dagoraerc20factory.Contract.TransferOwnership(&_Dagoraerc20factory.TransactOpts, newOwner)
}

// Dagoraerc20factoryDagoraERC20CreatedIterator is returned from FilterDagoraERC20Created and is used to iterate over the raw logs and unpacked data for DagoraERC20Created events raised by the Dagoraerc20factory contract.
type Dagoraerc20factoryDagoraERC20CreatedIterator struct {
	Event *Dagoraerc20factoryDagoraERC20Created // Event containing the contract specifics and raw log

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
func (it *Dagoraerc20factoryDagoraERC20CreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Dagoraerc20factoryDagoraERC20Created)
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
		it.Event = new(Dagoraerc20factoryDagoraERC20Created)
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
func (it *Dagoraerc20factoryDagoraERC20CreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Dagoraerc20factoryDagoraERC20CreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Dagoraerc20factoryDagoraERC20Created represents a DagoraERC20Created event raised by the Dagoraerc20factory contract.
type Dagoraerc20factoryDagoraERC20Created struct {
	NewContractAddress common.Address
	OwnerOF            common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDagoraERC20Created is a free log retrieval operation binding the contract event 0xa7b789ad2073113a69c2629d4bd8b42c3135bdae140c4b231ede7721156a7404.
//
// Solidity: event DagoraERC20Created(address newContractAddress, address ownerOF)
func (_Dagoraerc20factory *Dagoraerc20factoryFilterer) FilterDagoraERC20Created(opts *bind.FilterOpts) (*Dagoraerc20factoryDagoraERC20CreatedIterator, error) {

	logs, sub, err := _Dagoraerc20factory.contract.FilterLogs(opts, "DagoraERC20Created")
	if err != nil {
		return nil, err
	}
	return &Dagoraerc20factoryDagoraERC20CreatedIterator{contract: _Dagoraerc20factory.contract, event: "DagoraERC20Created", logs: logs, sub: sub}, nil
}

// WatchDagoraERC20Created is a free log subscription operation binding the contract event 0xa7b789ad2073113a69c2629d4bd8b42c3135bdae140c4b231ede7721156a7404.
//
// Solidity: event DagoraERC20Created(address newContractAddress, address ownerOF)
func (_Dagoraerc20factory *Dagoraerc20factoryFilterer) WatchDagoraERC20Created(opts *bind.WatchOpts, sink chan<- *Dagoraerc20factoryDagoraERC20Created) (event.Subscription, error) {

	logs, sub, err := _Dagoraerc20factory.contract.WatchLogs(opts, "DagoraERC20Created")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Dagoraerc20factoryDagoraERC20Created)
				if err := _Dagoraerc20factory.contract.UnpackLog(event, "DagoraERC20Created", log); err != nil {
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

// ParseDagoraERC20Created is a log parse operation binding the contract event 0xa7b789ad2073113a69c2629d4bd8b42c3135bdae140c4b231ede7721156a7404.
//
// Solidity: event DagoraERC20Created(address newContractAddress, address ownerOF)
func (_Dagoraerc20factory *Dagoraerc20factoryFilterer) ParseDagoraERC20Created(log types.Log) (*Dagoraerc20factoryDagoraERC20Created, error) {
	event := new(Dagoraerc20factoryDagoraERC20Created)
	if err := _Dagoraerc20factory.contract.UnpackLog(event, "DagoraERC20Created", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Dagoraerc20factoryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Dagoraerc20factory contract.
type Dagoraerc20factoryInitializedIterator struct {
	Event *Dagoraerc20factoryInitialized // Event containing the contract specifics and raw log

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
func (it *Dagoraerc20factoryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Dagoraerc20factoryInitialized)
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
		it.Event = new(Dagoraerc20factoryInitialized)
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
func (it *Dagoraerc20factoryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Dagoraerc20factoryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Dagoraerc20factoryInitialized represents a Initialized event raised by the Dagoraerc20factory contract.
type Dagoraerc20factoryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Dagoraerc20factory *Dagoraerc20factoryFilterer) FilterInitialized(opts *bind.FilterOpts) (*Dagoraerc20factoryInitializedIterator, error) {

	logs, sub, err := _Dagoraerc20factory.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &Dagoraerc20factoryInitializedIterator{contract: _Dagoraerc20factory.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Dagoraerc20factory *Dagoraerc20factoryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *Dagoraerc20factoryInitialized) (event.Subscription, error) {

	logs, sub, err := _Dagoraerc20factory.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Dagoraerc20factoryInitialized)
				if err := _Dagoraerc20factory.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Dagoraerc20factory *Dagoraerc20factoryFilterer) ParseInitialized(log types.Log) (*Dagoraerc20factoryInitialized, error) {
	event := new(Dagoraerc20factoryInitialized)
	if err := _Dagoraerc20factory.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Dagoraerc20factoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Dagoraerc20factory contract.
type Dagoraerc20factoryOwnershipTransferredIterator struct {
	Event *Dagoraerc20factoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Dagoraerc20factoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Dagoraerc20factoryOwnershipTransferred)
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
		it.Event = new(Dagoraerc20factoryOwnershipTransferred)
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
func (it *Dagoraerc20factoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Dagoraerc20factoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Dagoraerc20factoryOwnershipTransferred represents a OwnershipTransferred event raised by the Dagoraerc20factory contract.
type Dagoraerc20factoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dagoraerc20factory *Dagoraerc20factoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Dagoraerc20factoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dagoraerc20factory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Dagoraerc20factoryOwnershipTransferredIterator{contract: _Dagoraerc20factory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dagoraerc20factory *Dagoraerc20factoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Dagoraerc20factoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dagoraerc20factory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Dagoraerc20factoryOwnershipTransferred)
				if err := _Dagoraerc20factory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Dagoraerc20factory *Dagoraerc20factoryFilterer) ParseOwnershipTransferred(log types.Log) (*Dagoraerc20factoryOwnershipTransferred, error) {
	event := new(Dagoraerc20factoryOwnershipTransferred)
	if err := _Dagoraerc20factory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
