// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dagorasimplefntfactory

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

// DagorasimplefntfactoryMetaData contains all meta data concerning the Dagorasimplefntfactory contract.
var DagorasimplefntfactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DagoraFactory__ExpiredMembership\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DagoraFactory__FailedToCreateContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"tier\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"neededTier\",\"type\":\"uint8\"}],\"name\":\"DagoraFactory__InvalidTier\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DagoraFactory__NotDAgoraMembershipsOwnerOrDelegate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DagoraFactory__TokenCreationPaused\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newContractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ownerOF\",\"type\":\"address\"}],\"name\":\"SimpleNFTACreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"contractsDeployed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"baseURI_\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"_bulkBuyLimit\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_mintCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"createSimpleNFTA\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dAgoraMembershipsAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserContracts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dAgoraMembershipsAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minSimpleNFTATier\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_minSimpleNFTATier\",\"type\":\"uint8\"}],\"name\":\"setMinSimpleNFTATier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"togglePaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userContracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DagorasimplefntfactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use DagorasimplefntfactoryMetaData.ABI instead.
var DagorasimplefntfactoryABI = DagorasimplefntfactoryMetaData.ABI

// Dagorasimplefntfactory is an auto generated Go binding around an Ethereum contract.
type Dagorasimplefntfactory struct {
	DagorasimplefntfactoryCaller     // Read-only binding to the contract
	DagorasimplefntfactoryTransactor // Write-only binding to the contract
	DagorasimplefntfactoryFilterer   // Log filterer for contract events
}

// DagorasimplefntfactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type DagorasimplefntfactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DagorasimplefntfactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DagorasimplefntfactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DagorasimplefntfactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DagorasimplefntfactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DagorasimplefntfactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DagorasimplefntfactorySession struct {
	Contract     *Dagorasimplefntfactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DagorasimplefntfactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DagorasimplefntfactoryCallerSession struct {
	Contract *DagorasimplefntfactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// DagorasimplefntfactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DagorasimplefntfactoryTransactorSession struct {
	Contract     *DagorasimplefntfactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// DagorasimplefntfactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type DagorasimplefntfactoryRaw struct {
	Contract *Dagorasimplefntfactory // Generic contract binding to access the raw methods on
}

// DagorasimplefntfactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DagorasimplefntfactoryCallerRaw struct {
	Contract *DagorasimplefntfactoryCaller // Generic read-only contract binding to access the raw methods on
}

// DagorasimplefntfactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DagorasimplefntfactoryTransactorRaw struct {
	Contract *DagorasimplefntfactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDagorasimplefntfactory creates a new instance of Dagorasimplefntfactory, bound to a specific deployed contract.
func NewDagorasimplefntfactory(address common.Address, backend bind.ContractBackend) (*Dagorasimplefntfactory, error) {
	contract, err := bindDagorasimplefntfactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dagorasimplefntfactory{DagorasimplefntfactoryCaller: DagorasimplefntfactoryCaller{contract: contract}, DagorasimplefntfactoryTransactor: DagorasimplefntfactoryTransactor{contract: contract}, DagorasimplefntfactoryFilterer: DagorasimplefntfactoryFilterer{contract: contract}}, nil
}

// NewDagorasimplefntfactoryCaller creates a new read-only instance of Dagorasimplefntfactory, bound to a specific deployed contract.
func NewDagorasimplefntfactoryCaller(address common.Address, caller bind.ContractCaller) (*DagorasimplefntfactoryCaller, error) {
	contract, err := bindDagorasimplefntfactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DagorasimplefntfactoryCaller{contract: contract}, nil
}

// NewDagorasimplefntfactoryTransactor creates a new write-only instance of Dagorasimplefntfactory, bound to a specific deployed contract.
func NewDagorasimplefntfactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*DagorasimplefntfactoryTransactor, error) {
	contract, err := bindDagorasimplefntfactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DagorasimplefntfactoryTransactor{contract: contract}, nil
}

// NewDagorasimplefntfactoryFilterer creates a new log filterer instance of Dagorasimplefntfactory, bound to a specific deployed contract.
func NewDagorasimplefntfactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*DagorasimplefntfactoryFilterer, error) {
	contract, err := bindDagorasimplefntfactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DagorasimplefntfactoryFilterer{contract: contract}, nil
}

// bindDagorasimplefntfactory binds a generic wrapper to an already deployed contract.
func bindDagorasimplefntfactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DagorasimplefntfactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dagorasimplefntfactory *DagorasimplefntfactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dagorasimplefntfactory.Contract.DagorasimplefntfactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dagorasimplefntfactory *DagorasimplefntfactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dagorasimplefntfactory.Contract.DagorasimplefntfactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dagorasimplefntfactory *DagorasimplefntfactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dagorasimplefntfactory.Contract.DagorasimplefntfactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dagorasimplefntfactory *DagorasimplefntfactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dagorasimplefntfactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dagorasimplefntfactory *DagorasimplefntfactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dagorasimplefntfactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dagorasimplefntfactory *DagorasimplefntfactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dagorasimplefntfactory.Contract.contract.Transact(opts, method, params...)
}

// ContractsDeployed is a free data retrieval call binding the contract method 0xf0a761e4.
//
// Solidity: function contractsDeployed() view returns(uint256)
func (_Dagorasimplefntfactory *DagorasimplefntfactoryCaller) ContractsDeployed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dagorasimplefntfactory.contract.Call(opts, &out, "contractsDeployed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ContractsDeployed is a free data retrieval call binding the contract method 0xf0a761e4.
//
// Solidity: function contractsDeployed() view returns(uint256)
func (_Dagorasimplefntfactory *DagorasimplefntfactorySession) ContractsDeployed() (*big.Int, error) {
	return _Dagorasimplefntfactory.Contract.ContractsDeployed(&_Dagorasimplefntfactory.CallOpts)
}

// ContractsDeployed is a free data retrieval call binding the contract method 0xf0a761e4.
//
// Solidity: function contractsDeployed() view returns(uint256)
func (_Dagorasimplefntfactory *DagorasimplefntfactoryCallerSession) ContractsDeployed() (*big.Int, error) {
	return _Dagorasimplefntfactory.Contract.ContractsDeployed(&_Dagorasimplefntfactory.CallOpts)
}

// DAgoraMembershipsAddress is a free data retrieval call binding the contract method 0x514fd254.
//
// Solidity: function dAgoraMembershipsAddress() view returns(address)
func (_Dagorasimplefntfactory *DagorasimplefntfactoryCaller) DAgoraMembershipsAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dagorasimplefntfactory.contract.Call(opts, &out, "dAgoraMembershipsAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DAgoraMembershipsAddress is a free data retrieval call binding the contract method 0x514fd254.
//
// Solidity: function dAgoraMembershipsAddress() view returns(address)
func (_Dagorasimplefntfactory *DagorasimplefntfactorySession) DAgoraMembershipsAddress() (common.Address, error) {
	return _Dagorasimplefntfactory.Contract.DAgoraMembershipsAddress(&_Dagorasimplefntfactory.CallOpts)
}

// DAgoraMembershipsAddress is a free data retrieval call binding the contract method 0x514fd254.
//
// Solidity: function dAgoraMembershipsAddress() view returns(address)
func (_Dagorasimplefntfactory *DagorasimplefntfactoryCallerSession) DAgoraMembershipsAddress() (common.Address, error) {
	return _Dagorasimplefntfactory.Contract.DAgoraMembershipsAddress(&_Dagorasimplefntfactory.CallOpts)
}

// GetUserContracts is a free data retrieval call binding the contract method 0x34c925f0.
//
// Solidity: function getUserContracts(address _user) view returns(address[])
func (_Dagorasimplefntfactory *DagorasimplefntfactoryCaller) GetUserContracts(opts *bind.CallOpts, _user common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Dagorasimplefntfactory.contract.Call(opts, &out, "getUserContracts", _user)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetUserContracts is a free data retrieval call binding the contract method 0x34c925f0.
//
// Solidity: function getUserContracts(address _user) view returns(address[])
func (_Dagorasimplefntfactory *DagorasimplefntfactorySession) GetUserContracts(_user common.Address) ([]common.Address, error) {
	return _Dagorasimplefntfactory.Contract.GetUserContracts(&_Dagorasimplefntfactory.CallOpts, _user)
}

// GetUserContracts is a free data retrieval call binding the contract method 0x34c925f0.
//
// Solidity: function getUserContracts(address _user) view returns(address[])
func (_Dagorasimplefntfactory *DagorasimplefntfactoryCallerSession) GetUserContracts(_user common.Address) ([]common.Address, error) {
	return _Dagorasimplefntfactory.Contract.GetUserContracts(&_Dagorasimplefntfactory.CallOpts, _user)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Dagorasimplefntfactory *DagorasimplefntfactoryCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Dagorasimplefntfactory.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Dagorasimplefntfactory *DagorasimplefntfactorySession) IsPaused() (bool, error) {
	return _Dagorasimplefntfactory.Contract.IsPaused(&_Dagorasimplefntfactory.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Dagorasimplefntfactory *DagorasimplefntfactoryCallerSession) IsPaused() (bool, error) {
	return _Dagorasimplefntfactory.Contract.IsPaused(&_Dagorasimplefntfactory.CallOpts)
}

// MinSimpleNFTATier is a free data retrieval call binding the contract method 0x4d23cca1.
//
// Solidity: function minSimpleNFTATier() view returns(uint8)
func (_Dagorasimplefntfactory *DagorasimplefntfactoryCaller) MinSimpleNFTATier(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Dagorasimplefntfactory.contract.Call(opts, &out, "minSimpleNFTATier")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MinSimpleNFTATier is a free data retrieval call binding the contract method 0x4d23cca1.
//
// Solidity: function minSimpleNFTATier() view returns(uint8)
func (_Dagorasimplefntfactory *DagorasimplefntfactorySession) MinSimpleNFTATier() (uint8, error) {
	return _Dagorasimplefntfactory.Contract.MinSimpleNFTATier(&_Dagorasimplefntfactory.CallOpts)
}

// MinSimpleNFTATier is a free data retrieval call binding the contract method 0x4d23cca1.
//
// Solidity: function minSimpleNFTATier() view returns(uint8)
func (_Dagorasimplefntfactory *DagorasimplefntfactoryCallerSession) MinSimpleNFTATier() (uint8, error) {
	return _Dagorasimplefntfactory.Contract.MinSimpleNFTATier(&_Dagorasimplefntfactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dagorasimplefntfactory *DagorasimplefntfactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dagorasimplefntfactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dagorasimplefntfactory *DagorasimplefntfactorySession) Owner() (common.Address, error) {
	return _Dagorasimplefntfactory.Contract.Owner(&_Dagorasimplefntfactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dagorasimplefntfactory *DagorasimplefntfactoryCallerSession) Owner() (common.Address, error) {
	return _Dagorasimplefntfactory.Contract.Owner(&_Dagorasimplefntfactory.CallOpts)
}

// UserContracts is a free data retrieval call binding the contract method 0xb789003d.
//
// Solidity: function userContracts(address , uint256 ) view returns(address)
func (_Dagorasimplefntfactory *DagorasimplefntfactoryCaller) UserContracts(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Dagorasimplefntfactory.contract.Call(opts, &out, "userContracts", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserContracts is a free data retrieval call binding the contract method 0xb789003d.
//
// Solidity: function userContracts(address , uint256 ) view returns(address)
func (_Dagorasimplefntfactory *DagorasimplefntfactorySession) UserContracts(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Dagorasimplefntfactory.Contract.UserContracts(&_Dagorasimplefntfactory.CallOpts, arg0, arg1)
}

// UserContracts is a free data retrieval call binding the contract method 0xb789003d.
//
// Solidity: function userContracts(address , uint256 ) view returns(address)
func (_Dagorasimplefntfactory *DagorasimplefntfactoryCallerSession) UserContracts(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Dagorasimplefntfactory.Contract.UserContracts(&_Dagorasimplefntfactory.CallOpts, arg0, arg1)
}

// CreateSimpleNFTA is a paid mutator transaction binding the contract method 0x7db5f539.
//
// Solidity: function createSimpleNFTA(string name_, string symbol_, string baseURI_, uint16 _bulkBuyLimit, uint256 _mintCost, uint256 _maxSupply, address _newOwner, uint256 _id) returns()
func (_Dagorasimplefntfactory *DagorasimplefntfactoryTransactor) CreateSimpleNFTA(opts *bind.TransactOpts, name_ string, symbol_ string, baseURI_ string, _bulkBuyLimit uint16, _mintCost *big.Int, _maxSupply *big.Int, _newOwner common.Address, _id *big.Int) (*types.Transaction, error) {
	return _Dagorasimplefntfactory.contract.Transact(opts, "createSimpleNFTA", name_, symbol_, baseURI_, _bulkBuyLimit, _mintCost, _maxSupply, _newOwner, _id)
}

// CreateSimpleNFTA is a paid mutator transaction binding the contract method 0x7db5f539.
//
// Solidity: function createSimpleNFTA(string name_, string symbol_, string baseURI_, uint16 _bulkBuyLimit, uint256 _mintCost, uint256 _maxSupply, address _newOwner, uint256 _id) returns()
func (_Dagorasimplefntfactory *DagorasimplefntfactorySession) CreateSimpleNFTA(name_ string, symbol_ string, baseURI_ string, _bulkBuyLimit uint16, _mintCost *big.Int, _maxSupply *big.Int, _newOwner common.Address, _id *big.Int) (*types.Transaction, error) {
	return _Dagorasimplefntfactory.Contract.CreateSimpleNFTA(&_Dagorasimplefntfactory.TransactOpts, name_, symbol_, baseURI_, _bulkBuyLimit, _mintCost, _maxSupply, _newOwner, _id)
}

// CreateSimpleNFTA is a paid mutator transaction binding the contract method 0x7db5f539.
//
// Solidity: function createSimpleNFTA(string name_, string symbol_, string baseURI_, uint16 _bulkBuyLimit, uint256 _mintCost, uint256 _maxSupply, address _newOwner, uint256 _id) returns()
func (_Dagorasimplefntfactory *DagorasimplefntfactoryTransactorSession) CreateSimpleNFTA(name_ string, symbol_ string, baseURI_ string, _bulkBuyLimit uint16, _mintCost *big.Int, _maxSupply *big.Int, _newOwner common.Address, _id *big.Int) (*types.Transaction, error) {
	return _Dagorasimplefntfactory.Contract.CreateSimpleNFTA(&_Dagorasimplefntfactory.TransactOpts, name_, symbol_, baseURI_, _bulkBuyLimit, _mintCost, _maxSupply, _newOwner, _id)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _dAgoraMembershipsAddress) returns()
func (_Dagorasimplefntfactory *DagorasimplefntfactoryTransactor) Initialize(opts *bind.TransactOpts, _dAgoraMembershipsAddress common.Address) (*types.Transaction, error) {
	return _Dagorasimplefntfactory.contract.Transact(opts, "initialize", _dAgoraMembershipsAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _dAgoraMembershipsAddress) returns()
func (_Dagorasimplefntfactory *DagorasimplefntfactorySession) Initialize(_dAgoraMembershipsAddress common.Address) (*types.Transaction, error) {
	return _Dagorasimplefntfactory.Contract.Initialize(&_Dagorasimplefntfactory.TransactOpts, _dAgoraMembershipsAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _dAgoraMembershipsAddress) returns()
func (_Dagorasimplefntfactory *DagorasimplefntfactoryTransactorSession) Initialize(_dAgoraMembershipsAddress common.Address) (*types.Transaction, error) {
	return _Dagorasimplefntfactory.Contract.Initialize(&_Dagorasimplefntfactory.TransactOpts, _dAgoraMembershipsAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dagorasimplefntfactory *DagorasimplefntfactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dagorasimplefntfactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dagorasimplefntfactory *DagorasimplefntfactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _Dagorasimplefntfactory.Contract.RenounceOwnership(&_Dagorasimplefntfactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dagorasimplefntfactory *DagorasimplefntfactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dagorasimplefntfactory.Contract.RenounceOwnership(&_Dagorasimplefntfactory.TransactOpts)
}

// SetMinSimpleNFTATier is a paid mutator transaction binding the contract method 0x8d6f04f6.
//
// Solidity: function setMinSimpleNFTATier(uint8 _minSimpleNFTATier) returns()
func (_Dagorasimplefntfactory *DagorasimplefntfactoryTransactor) SetMinSimpleNFTATier(opts *bind.TransactOpts, _minSimpleNFTATier uint8) (*types.Transaction, error) {
	return _Dagorasimplefntfactory.contract.Transact(opts, "setMinSimpleNFTATier", _minSimpleNFTATier)
}

// SetMinSimpleNFTATier is a paid mutator transaction binding the contract method 0x8d6f04f6.
//
// Solidity: function setMinSimpleNFTATier(uint8 _minSimpleNFTATier) returns()
func (_Dagorasimplefntfactory *DagorasimplefntfactorySession) SetMinSimpleNFTATier(_minSimpleNFTATier uint8) (*types.Transaction, error) {
	return _Dagorasimplefntfactory.Contract.SetMinSimpleNFTATier(&_Dagorasimplefntfactory.TransactOpts, _minSimpleNFTATier)
}

// SetMinSimpleNFTATier is a paid mutator transaction binding the contract method 0x8d6f04f6.
//
// Solidity: function setMinSimpleNFTATier(uint8 _minSimpleNFTATier) returns()
func (_Dagorasimplefntfactory *DagorasimplefntfactoryTransactorSession) SetMinSimpleNFTATier(_minSimpleNFTATier uint8) (*types.Transaction, error) {
	return _Dagorasimplefntfactory.Contract.SetMinSimpleNFTATier(&_Dagorasimplefntfactory.TransactOpts, _minSimpleNFTATier)
}

// TogglePaused is a paid mutator transaction binding the contract method 0x36566f06.
//
// Solidity: function togglePaused() returns()
func (_Dagorasimplefntfactory *DagorasimplefntfactoryTransactor) TogglePaused(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dagorasimplefntfactory.contract.Transact(opts, "togglePaused")
}

// TogglePaused is a paid mutator transaction binding the contract method 0x36566f06.
//
// Solidity: function togglePaused() returns()
func (_Dagorasimplefntfactory *DagorasimplefntfactorySession) TogglePaused() (*types.Transaction, error) {
	return _Dagorasimplefntfactory.Contract.TogglePaused(&_Dagorasimplefntfactory.TransactOpts)
}

// TogglePaused is a paid mutator transaction binding the contract method 0x36566f06.
//
// Solidity: function togglePaused() returns()
func (_Dagorasimplefntfactory *DagorasimplefntfactoryTransactorSession) TogglePaused() (*types.Transaction, error) {
	return _Dagorasimplefntfactory.Contract.TogglePaused(&_Dagorasimplefntfactory.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dagorasimplefntfactory *DagorasimplefntfactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Dagorasimplefntfactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dagorasimplefntfactory *DagorasimplefntfactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dagorasimplefntfactory.Contract.TransferOwnership(&_Dagorasimplefntfactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dagorasimplefntfactory *DagorasimplefntfactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dagorasimplefntfactory.Contract.TransferOwnership(&_Dagorasimplefntfactory.TransactOpts, newOwner)
}

// DagorasimplefntfactoryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Dagorasimplefntfactory contract.
type DagorasimplefntfactoryInitializedIterator struct {
	Event *DagorasimplefntfactoryInitialized // Event containing the contract specifics and raw log

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
func (it *DagorasimplefntfactoryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DagorasimplefntfactoryInitialized)
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
		it.Event = new(DagorasimplefntfactoryInitialized)
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
func (it *DagorasimplefntfactoryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DagorasimplefntfactoryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DagorasimplefntfactoryInitialized represents a Initialized event raised by the Dagorasimplefntfactory contract.
type DagorasimplefntfactoryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Dagorasimplefntfactory *DagorasimplefntfactoryFilterer) FilterInitialized(opts *bind.FilterOpts) (*DagorasimplefntfactoryInitializedIterator, error) {

	logs, sub, err := _Dagorasimplefntfactory.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DagorasimplefntfactoryInitializedIterator{contract: _Dagorasimplefntfactory.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Dagorasimplefntfactory *DagorasimplefntfactoryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DagorasimplefntfactoryInitialized) (event.Subscription, error) {

	logs, sub, err := _Dagorasimplefntfactory.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DagorasimplefntfactoryInitialized)
				if err := _Dagorasimplefntfactory.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Dagorasimplefntfactory *DagorasimplefntfactoryFilterer) ParseInitialized(log types.Log) (*DagorasimplefntfactoryInitialized, error) {
	event := new(DagorasimplefntfactoryInitialized)
	if err := _Dagorasimplefntfactory.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DagorasimplefntfactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Dagorasimplefntfactory contract.
type DagorasimplefntfactoryOwnershipTransferredIterator struct {
	Event *DagorasimplefntfactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DagorasimplefntfactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DagorasimplefntfactoryOwnershipTransferred)
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
		it.Event = new(DagorasimplefntfactoryOwnershipTransferred)
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
func (it *DagorasimplefntfactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DagorasimplefntfactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DagorasimplefntfactoryOwnershipTransferred represents a OwnershipTransferred event raised by the Dagorasimplefntfactory contract.
type DagorasimplefntfactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dagorasimplefntfactory *DagorasimplefntfactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DagorasimplefntfactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dagorasimplefntfactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DagorasimplefntfactoryOwnershipTransferredIterator{contract: _Dagorasimplefntfactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dagorasimplefntfactory *DagorasimplefntfactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DagorasimplefntfactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dagorasimplefntfactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DagorasimplefntfactoryOwnershipTransferred)
				if err := _Dagorasimplefntfactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Dagorasimplefntfactory *DagorasimplefntfactoryFilterer) ParseOwnershipTransferred(log types.Log) (*DagorasimplefntfactoryOwnershipTransferred, error) {
	event := new(DagorasimplefntfactoryOwnershipTransferred)
	if err := _Dagorasimplefntfactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DagorasimplefntfactorySimpleNFTACreatedIterator is returned from FilterSimpleNFTACreated and is used to iterate over the raw logs and unpacked data for SimpleNFTACreated events raised by the Dagorasimplefntfactory contract.
type DagorasimplefntfactorySimpleNFTACreatedIterator struct {
	Event *DagorasimplefntfactorySimpleNFTACreated // Event containing the contract specifics and raw log

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
func (it *DagorasimplefntfactorySimpleNFTACreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DagorasimplefntfactorySimpleNFTACreated)
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
		it.Event = new(DagorasimplefntfactorySimpleNFTACreated)
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
func (it *DagorasimplefntfactorySimpleNFTACreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DagorasimplefntfactorySimpleNFTACreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DagorasimplefntfactorySimpleNFTACreated represents a SimpleNFTACreated event raised by the Dagorasimplefntfactory contract.
type DagorasimplefntfactorySimpleNFTACreated struct {
	NewContractAddress common.Address
	OwnerOF            common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSimpleNFTACreated is a free log retrieval operation binding the contract event 0x663ce7590af72a41b27067b95f2b102b72df0c912cc5e44f40a26a617e26865d.
//
// Solidity: event SimpleNFTACreated(address newContractAddress, address ownerOF)
func (_Dagorasimplefntfactory *DagorasimplefntfactoryFilterer) FilterSimpleNFTACreated(opts *bind.FilterOpts) (*DagorasimplefntfactorySimpleNFTACreatedIterator, error) {

	logs, sub, err := _Dagorasimplefntfactory.contract.FilterLogs(opts, "SimpleNFTACreated")
	if err != nil {
		return nil, err
	}
	return &DagorasimplefntfactorySimpleNFTACreatedIterator{contract: _Dagorasimplefntfactory.contract, event: "SimpleNFTACreated", logs: logs, sub: sub}, nil
}

// WatchSimpleNFTACreated is a free log subscription operation binding the contract event 0x663ce7590af72a41b27067b95f2b102b72df0c912cc5e44f40a26a617e26865d.
//
// Solidity: event SimpleNFTACreated(address newContractAddress, address ownerOF)
func (_Dagorasimplefntfactory *DagorasimplefntfactoryFilterer) WatchSimpleNFTACreated(opts *bind.WatchOpts, sink chan<- *DagorasimplefntfactorySimpleNFTACreated) (event.Subscription, error) {

	logs, sub, err := _Dagorasimplefntfactory.contract.WatchLogs(opts, "SimpleNFTACreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DagorasimplefntfactorySimpleNFTACreated)
				if err := _Dagorasimplefntfactory.contract.UnpackLog(event, "SimpleNFTACreated", log); err != nil {
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

// ParseSimpleNFTACreated is a log parse operation binding the contract event 0x663ce7590af72a41b27067b95f2b102b72df0c912cc5e44f40a26a617e26865d.
//
// Solidity: event SimpleNFTACreated(address newContractAddress, address ownerOF)
func (_Dagorasimplefntfactory *DagorasimplefntfactoryFilterer) ParseSimpleNFTACreated(log types.Log) (*DagorasimplefntfactorySimpleNFTACreated, error) {
	event := new(DagorasimplefntfactorySimpleNFTACreated)
	if err := _Dagorasimplefntfactory.contract.UnpackLog(event, "SimpleNFTACreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
