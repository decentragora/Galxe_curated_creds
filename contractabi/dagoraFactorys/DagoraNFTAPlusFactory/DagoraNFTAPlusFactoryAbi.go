// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dagoranftaplusfactory

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

// DagoranftaplusfactoryMetaData contains all meta data concerning the Dagoranftaplusfactory contract.
var DagoranftaplusfactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DagoraFactory__ExpiredMembership\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DagoraFactory__FailedToCreateContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"tier\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"neededTier\",\"type\":\"uint8\"}],\"name\":\"DagoraFactory__InvalidTier\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DagoraFactory__NotDAgoraMembershipsOwnerOrDelegate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DagoraFactory__TokenCreationPaused\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newContractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ownerOF\",\"type\":\"address\"}],\"name\":\"NFTAPlusCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"contractsDeployed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"baseURI_\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"_bulkBuyLimit\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_maxAllowListAmount\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_mintCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_presaleMintCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"createNFTAPlus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dAgoraMembershipsAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserContracts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dAgoraMembershipsAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minNFTAPlusTier\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_minNFTAPlusTier\",\"type\":\"uint8\"}],\"name\":\"setMinNFTAPlusTier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"togglePaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userContracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DagoranftaplusfactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use DagoranftaplusfactoryMetaData.ABI instead.
var DagoranftaplusfactoryABI = DagoranftaplusfactoryMetaData.ABI

// Dagoranftaplusfactory is an auto generated Go binding around an Ethereum contract.
type Dagoranftaplusfactory struct {
	DagoranftaplusfactoryCaller     // Read-only binding to the contract
	DagoranftaplusfactoryTransactor // Write-only binding to the contract
	DagoranftaplusfactoryFilterer   // Log filterer for contract events
}

// DagoranftaplusfactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type DagoranftaplusfactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DagoranftaplusfactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DagoranftaplusfactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DagoranftaplusfactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DagoranftaplusfactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DagoranftaplusfactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DagoranftaplusfactorySession struct {
	Contract     *Dagoranftaplusfactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DagoranftaplusfactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DagoranftaplusfactoryCallerSession struct {
	Contract *DagoranftaplusfactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// DagoranftaplusfactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DagoranftaplusfactoryTransactorSession struct {
	Contract     *DagoranftaplusfactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// DagoranftaplusfactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type DagoranftaplusfactoryRaw struct {
	Contract *Dagoranftaplusfactory // Generic contract binding to access the raw methods on
}

// DagoranftaplusfactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DagoranftaplusfactoryCallerRaw struct {
	Contract *DagoranftaplusfactoryCaller // Generic read-only contract binding to access the raw methods on
}

// DagoranftaplusfactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DagoranftaplusfactoryTransactorRaw struct {
	Contract *DagoranftaplusfactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDagoranftaplusfactory creates a new instance of Dagoranftaplusfactory, bound to a specific deployed contract.
func NewDagoranftaplusfactory(address common.Address, backend bind.ContractBackend) (*Dagoranftaplusfactory, error) {
	contract, err := bindDagoranftaplusfactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dagoranftaplusfactory{DagoranftaplusfactoryCaller: DagoranftaplusfactoryCaller{contract: contract}, DagoranftaplusfactoryTransactor: DagoranftaplusfactoryTransactor{contract: contract}, DagoranftaplusfactoryFilterer: DagoranftaplusfactoryFilterer{contract: contract}}, nil
}

// NewDagoranftaplusfactoryCaller creates a new read-only instance of Dagoranftaplusfactory, bound to a specific deployed contract.
func NewDagoranftaplusfactoryCaller(address common.Address, caller bind.ContractCaller) (*DagoranftaplusfactoryCaller, error) {
	contract, err := bindDagoranftaplusfactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DagoranftaplusfactoryCaller{contract: contract}, nil
}

// NewDagoranftaplusfactoryTransactor creates a new write-only instance of Dagoranftaplusfactory, bound to a specific deployed contract.
func NewDagoranftaplusfactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*DagoranftaplusfactoryTransactor, error) {
	contract, err := bindDagoranftaplusfactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DagoranftaplusfactoryTransactor{contract: contract}, nil
}

// NewDagoranftaplusfactoryFilterer creates a new log filterer instance of Dagoranftaplusfactory, bound to a specific deployed contract.
func NewDagoranftaplusfactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*DagoranftaplusfactoryFilterer, error) {
	contract, err := bindDagoranftaplusfactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DagoranftaplusfactoryFilterer{contract: contract}, nil
}

// bindDagoranftaplusfactory binds a generic wrapper to an already deployed contract.
func bindDagoranftaplusfactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DagoranftaplusfactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dagoranftaplusfactory *DagoranftaplusfactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dagoranftaplusfactory.Contract.DagoranftaplusfactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dagoranftaplusfactory *DagoranftaplusfactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dagoranftaplusfactory.Contract.DagoranftaplusfactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dagoranftaplusfactory *DagoranftaplusfactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dagoranftaplusfactory.Contract.DagoranftaplusfactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dagoranftaplusfactory *DagoranftaplusfactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dagoranftaplusfactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dagoranftaplusfactory *DagoranftaplusfactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dagoranftaplusfactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dagoranftaplusfactory *DagoranftaplusfactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dagoranftaplusfactory.Contract.contract.Transact(opts, method, params...)
}

// ContractsDeployed is a free data retrieval call binding the contract method 0xf0a761e4.
//
// Solidity: function contractsDeployed() view returns(uint256)
func (_Dagoranftaplusfactory *DagoranftaplusfactoryCaller) ContractsDeployed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dagoranftaplusfactory.contract.Call(opts, &out, "contractsDeployed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ContractsDeployed is a free data retrieval call binding the contract method 0xf0a761e4.
//
// Solidity: function contractsDeployed() view returns(uint256)
func (_Dagoranftaplusfactory *DagoranftaplusfactorySession) ContractsDeployed() (*big.Int, error) {
	return _Dagoranftaplusfactory.Contract.ContractsDeployed(&_Dagoranftaplusfactory.CallOpts)
}

// ContractsDeployed is a free data retrieval call binding the contract method 0xf0a761e4.
//
// Solidity: function contractsDeployed() view returns(uint256)
func (_Dagoranftaplusfactory *DagoranftaplusfactoryCallerSession) ContractsDeployed() (*big.Int, error) {
	return _Dagoranftaplusfactory.Contract.ContractsDeployed(&_Dagoranftaplusfactory.CallOpts)
}

// DAgoraMembershipsAddress is a free data retrieval call binding the contract method 0x514fd254.
//
// Solidity: function dAgoraMembershipsAddress() view returns(address)
func (_Dagoranftaplusfactory *DagoranftaplusfactoryCaller) DAgoraMembershipsAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dagoranftaplusfactory.contract.Call(opts, &out, "dAgoraMembershipsAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DAgoraMembershipsAddress is a free data retrieval call binding the contract method 0x514fd254.
//
// Solidity: function dAgoraMembershipsAddress() view returns(address)
func (_Dagoranftaplusfactory *DagoranftaplusfactorySession) DAgoraMembershipsAddress() (common.Address, error) {
	return _Dagoranftaplusfactory.Contract.DAgoraMembershipsAddress(&_Dagoranftaplusfactory.CallOpts)
}

// DAgoraMembershipsAddress is a free data retrieval call binding the contract method 0x514fd254.
//
// Solidity: function dAgoraMembershipsAddress() view returns(address)
func (_Dagoranftaplusfactory *DagoranftaplusfactoryCallerSession) DAgoraMembershipsAddress() (common.Address, error) {
	return _Dagoranftaplusfactory.Contract.DAgoraMembershipsAddress(&_Dagoranftaplusfactory.CallOpts)
}

// GetUserContracts is a free data retrieval call binding the contract method 0x34c925f0.
//
// Solidity: function getUserContracts(address _user) view returns(address[])
func (_Dagoranftaplusfactory *DagoranftaplusfactoryCaller) GetUserContracts(opts *bind.CallOpts, _user common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Dagoranftaplusfactory.contract.Call(opts, &out, "getUserContracts", _user)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetUserContracts is a free data retrieval call binding the contract method 0x34c925f0.
//
// Solidity: function getUserContracts(address _user) view returns(address[])
func (_Dagoranftaplusfactory *DagoranftaplusfactorySession) GetUserContracts(_user common.Address) ([]common.Address, error) {
	return _Dagoranftaplusfactory.Contract.GetUserContracts(&_Dagoranftaplusfactory.CallOpts, _user)
}

// GetUserContracts is a free data retrieval call binding the contract method 0x34c925f0.
//
// Solidity: function getUserContracts(address _user) view returns(address[])
func (_Dagoranftaplusfactory *DagoranftaplusfactoryCallerSession) GetUserContracts(_user common.Address) ([]common.Address, error) {
	return _Dagoranftaplusfactory.Contract.GetUserContracts(&_Dagoranftaplusfactory.CallOpts, _user)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Dagoranftaplusfactory *DagoranftaplusfactoryCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Dagoranftaplusfactory.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Dagoranftaplusfactory *DagoranftaplusfactorySession) IsPaused() (bool, error) {
	return _Dagoranftaplusfactory.Contract.IsPaused(&_Dagoranftaplusfactory.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Dagoranftaplusfactory *DagoranftaplusfactoryCallerSession) IsPaused() (bool, error) {
	return _Dagoranftaplusfactory.Contract.IsPaused(&_Dagoranftaplusfactory.CallOpts)
}

// MinNFTAPlusTier is a free data retrieval call binding the contract method 0xc369a5fa.
//
// Solidity: function minNFTAPlusTier() view returns(uint8)
func (_Dagoranftaplusfactory *DagoranftaplusfactoryCaller) MinNFTAPlusTier(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Dagoranftaplusfactory.contract.Call(opts, &out, "minNFTAPlusTier")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MinNFTAPlusTier is a free data retrieval call binding the contract method 0xc369a5fa.
//
// Solidity: function minNFTAPlusTier() view returns(uint8)
func (_Dagoranftaplusfactory *DagoranftaplusfactorySession) MinNFTAPlusTier() (uint8, error) {
	return _Dagoranftaplusfactory.Contract.MinNFTAPlusTier(&_Dagoranftaplusfactory.CallOpts)
}

// MinNFTAPlusTier is a free data retrieval call binding the contract method 0xc369a5fa.
//
// Solidity: function minNFTAPlusTier() view returns(uint8)
func (_Dagoranftaplusfactory *DagoranftaplusfactoryCallerSession) MinNFTAPlusTier() (uint8, error) {
	return _Dagoranftaplusfactory.Contract.MinNFTAPlusTier(&_Dagoranftaplusfactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dagoranftaplusfactory *DagoranftaplusfactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dagoranftaplusfactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dagoranftaplusfactory *DagoranftaplusfactorySession) Owner() (common.Address, error) {
	return _Dagoranftaplusfactory.Contract.Owner(&_Dagoranftaplusfactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dagoranftaplusfactory *DagoranftaplusfactoryCallerSession) Owner() (common.Address, error) {
	return _Dagoranftaplusfactory.Contract.Owner(&_Dagoranftaplusfactory.CallOpts)
}

// UserContracts is a free data retrieval call binding the contract method 0xb789003d.
//
// Solidity: function userContracts(address , uint256 ) view returns(address)
func (_Dagoranftaplusfactory *DagoranftaplusfactoryCaller) UserContracts(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Dagoranftaplusfactory.contract.Call(opts, &out, "userContracts", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserContracts is a free data retrieval call binding the contract method 0xb789003d.
//
// Solidity: function userContracts(address , uint256 ) view returns(address)
func (_Dagoranftaplusfactory *DagoranftaplusfactorySession) UserContracts(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Dagoranftaplusfactory.Contract.UserContracts(&_Dagoranftaplusfactory.CallOpts, arg0, arg1)
}

// UserContracts is a free data retrieval call binding the contract method 0xb789003d.
//
// Solidity: function userContracts(address , uint256 ) view returns(address)
func (_Dagoranftaplusfactory *DagoranftaplusfactoryCallerSession) UserContracts(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Dagoranftaplusfactory.Contract.UserContracts(&_Dagoranftaplusfactory.CallOpts, arg0, arg1)
}

// CreateNFTAPlus is a paid mutator transaction binding the contract method 0xfef5e645.
//
// Solidity: function createNFTAPlus(string name_, string symbol_, string baseURI_, uint16 _bulkBuyLimit, uint16 _maxAllowListAmount, uint256 _mintCost, uint256 _presaleMintCost, uint256 _maxSupply, address _newOwner, bytes32 _merkleRoot, uint256 _id) returns()
func (_Dagoranftaplusfactory *DagoranftaplusfactoryTransactor) CreateNFTAPlus(opts *bind.TransactOpts, name_ string, symbol_ string, baseURI_ string, _bulkBuyLimit uint16, _maxAllowListAmount uint16, _mintCost *big.Int, _presaleMintCost *big.Int, _maxSupply *big.Int, _newOwner common.Address, _merkleRoot [32]byte, _id *big.Int) (*types.Transaction, error) {
	return _Dagoranftaplusfactory.contract.Transact(opts, "createNFTAPlus", name_, symbol_, baseURI_, _bulkBuyLimit, _maxAllowListAmount, _mintCost, _presaleMintCost, _maxSupply, _newOwner, _merkleRoot, _id)
}

// CreateNFTAPlus is a paid mutator transaction binding the contract method 0xfef5e645.
//
// Solidity: function createNFTAPlus(string name_, string symbol_, string baseURI_, uint16 _bulkBuyLimit, uint16 _maxAllowListAmount, uint256 _mintCost, uint256 _presaleMintCost, uint256 _maxSupply, address _newOwner, bytes32 _merkleRoot, uint256 _id) returns()
func (_Dagoranftaplusfactory *DagoranftaplusfactorySession) CreateNFTAPlus(name_ string, symbol_ string, baseURI_ string, _bulkBuyLimit uint16, _maxAllowListAmount uint16, _mintCost *big.Int, _presaleMintCost *big.Int, _maxSupply *big.Int, _newOwner common.Address, _merkleRoot [32]byte, _id *big.Int) (*types.Transaction, error) {
	return _Dagoranftaplusfactory.Contract.CreateNFTAPlus(&_Dagoranftaplusfactory.TransactOpts, name_, symbol_, baseURI_, _bulkBuyLimit, _maxAllowListAmount, _mintCost, _presaleMintCost, _maxSupply, _newOwner, _merkleRoot, _id)
}

// CreateNFTAPlus is a paid mutator transaction binding the contract method 0xfef5e645.
//
// Solidity: function createNFTAPlus(string name_, string symbol_, string baseURI_, uint16 _bulkBuyLimit, uint16 _maxAllowListAmount, uint256 _mintCost, uint256 _presaleMintCost, uint256 _maxSupply, address _newOwner, bytes32 _merkleRoot, uint256 _id) returns()
func (_Dagoranftaplusfactory *DagoranftaplusfactoryTransactorSession) CreateNFTAPlus(name_ string, symbol_ string, baseURI_ string, _bulkBuyLimit uint16, _maxAllowListAmount uint16, _mintCost *big.Int, _presaleMintCost *big.Int, _maxSupply *big.Int, _newOwner common.Address, _merkleRoot [32]byte, _id *big.Int) (*types.Transaction, error) {
	return _Dagoranftaplusfactory.Contract.CreateNFTAPlus(&_Dagoranftaplusfactory.TransactOpts, name_, symbol_, baseURI_, _bulkBuyLimit, _maxAllowListAmount, _mintCost, _presaleMintCost, _maxSupply, _newOwner, _merkleRoot, _id)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _dAgoraMembershipsAddress) returns()
func (_Dagoranftaplusfactory *DagoranftaplusfactoryTransactor) Initialize(opts *bind.TransactOpts, _dAgoraMembershipsAddress common.Address) (*types.Transaction, error) {
	return _Dagoranftaplusfactory.contract.Transact(opts, "initialize", _dAgoraMembershipsAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _dAgoraMembershipsAddress) returns()
func (_Dagoranftaplusfactory *DagoranftaplusfactorySession) Initialize(_dAgoraMembershipsAddress common.Address) (*types.Transaction, error) {
	return _Dagoranftaplusfactory.Contract.Initialize(&_Dagoranftaplusfactory.TransactOpts, _dAgoraMembershipsAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _dAgoraMembershipsAddress) returns()
func (_Dagoranftaplusfactory *DagoranftaplusfactoryTransactorSession) Initialize(_dAgoraMembershipsAddress common.Address) (*types.Transaction, error) {
	return _Dagoranftaplusfactory.Contract.Initialize(&_Dagoranftaplusfactory.TransactOpts, _dAgoraMembershipsAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dagoranftaplusfactory *DagoranftaplusfactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dagoranftaplusfactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dagoranftaplusfactory *DagoranftaplusfactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _Dagoranftaplusfactory.Contract.RenounceOwnership(&_Dagoranftaplusfactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dagoranftaplusfactory *DagoranftaplusfactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dagoranftaplusfactory.Contract.RenounceOwnership(&_Dagoranftaplusfactory.TransactOpts)
}

// SetMinNFTAPlusTier is a paid mutator transaction binding the contract method 0x18fdf2af.
//
// Solidity: function setMinNFTAPlusTier(uint8 _minNFTAPlusTier) returns()
func (_Dagoranftaplusfactory *DagoranftaplusfactoryTransactor) SetMinNFTAPlusTier(opts *bind.TransactOpts, _minNFTAPlusTier uint8) (*types.Transaction, error) {
	return _Dagoranftaplusfactory.contract.Transact(opts, "setMinNFTAPlusTier", _minNFTAPlusTier)
}

// SetMinNFTAPlusTier is a paid mutator transaction binding the contract method 0x18fdf2af.
//
// Solidity: function setMinNFTAPlusTier(uint8 _minNFTAPlusTier) returns()
func (_Dagoranftaplusfactory *DagoranftaplusfactorySession) SetMinNFTAPlusTier(_minNFTAPlusTier uint8) (*types.Transaction, error) {
	return _Dagoranftaplusfactory.Contract.SetMinNFTAPlusTier(&_Dagoranftaplusfactory.TransactOpts, _minNFTAPlusTier)
}

// SetMinNFTAPlusTier is a paid mutator transaction binding the contract method 0x18fdf2af.
//
// Solidity: function setMinNFTAPlusTier(uint8 _minNFTAPlusTier) returns()
func (_Dagoranftaplusfactory *DagoranftaplusfactoryTransactorSession) SetMinNFTAPlusTier(_minNFTAPlusTier uint8) (*types.Transaction, error) {
	return _Dagoranftaplusfactory.Contract.SetMinNFTAPlusTier(&_Dagoranftaplusfactory.TransactOpts, _minNFTAPlusTier)
}

// TogglePaused is a paid mutator transaction binding the contract method 0x36566f06.
//
// Solidity: function togglePaused() returns()
func (_Dagoranftaplusfactory *DagoranftaplusfactoryTransactor) TogglePaused(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dagoranftaplusfactory.contract.Transact(opts, "togglePaused")
}

// TogglePaused is a paid mutator transaction binding the contract method 0x36566f06.
//
// Solidity: function togglePaused() returns()
func (_Dagoranftaplusfactory *DagoranftaplusfactorySession) TogglePaused() (*types.Transaction, error) {
	return _Dagoranftaplusfactory.Contract.TogglePaused(&_Dagoranftaplusfactory.TransactOpts)
}

// TogglePaused is a paid mutator transaction binding the contract method 0x36566f06.
//
// Solidity: function togglePaused() returns()
func (_Dagoranftaplusfactory *DagoranftaplusfactoryTransactorSession) TogglePaused() (*types.Transaction, error) {
	return _Dagoranftaplusfactory.Contract.TogglePaused(&_Dagoranftaplusfactory.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dagoranftaplusfactory *DagoranftaplusfactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Dagoranftaplusfactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dagoranftaplusfactory *DagoranftaplusfactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dagoranftaplusfactory.Contract.TransferOwnership(&_Dagoranftaplusfactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dagoranftaplusfactory *DagoranftaplusfactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dagoranftaplusfactory.Contract.TransferOwnership(&_Dagoranftaplusfactory.TransactOpts, newOwner)
}

// DagoranftaplusfactoryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Dagoranftaplusfactory contract.
type DagoranftaplusfactoryInitializedIterator struct {
	Event *DagoranftaplusfactoryInitialized // Event containing the contract specifics and raw log

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
func (it *DagoranftaplusfactoryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DagoranftaplusfactoryInitialized)
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
		it.Event = new(DagoranftaplusfactoryInitialized)
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
func (it *DagoranftaplusfactoryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DagoranftaplusfactoryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DagoranftaplusfactoryInitialized represents a Initialized event raised by the Dagoranftaplusfactory contract.
type DagoranftaplusfactoryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Dagoranftaplusfactory *DagoranftaplusfactoryFilterer) FilterInitialized(opts *bind.FilterOpts) (*DagoranftaplusfactoryInitializedIterator, error) {

	logs, sub, err := _Dagoranftaplusfactory.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DagoranftaplusfactoryInitializedIterator{contract: _Dagoranftaplusfactory.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Dagoranftaplusfactory *DagoranftaplusfactoryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DagoranftaplusfactoryInitialized) (event.Subscription, error) {

	logs, sub, err := _Dagoranftaplusfactory.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DagoranftaplusfactoryInitialized)
				if err := _Dagoranftaplusfactory.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Dagoranftaplusfactory *DagoranftaplusfactoryFilterer) ParseInitialized(log types.Log) (*DagoranftaplusfactoryInitialized, error) {
	event := new(DagoranftaplusfactoryInitialized)
	if err := _Dagoranftaplusfactory.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DagoranftaplusfactoryNFTAPlusCreatedIterator is returned from FilterNFTAPlusCreated and is used to iterate over the raw logs and unpacked data for NFTAPlusCreated events raised by the Dagoranftaplusfactory contract.
type DagoranftaplusfactoryNFTAPlusCreatedIterator struct {
	Event *DagoranftaplusfactoryNFTAPlusCreated // Event containing the contract specifics and raw log

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
func (it *DagoranftaplusfactoryNFTAPlusCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DagoranftaplusfactoryNFTAPlusCreated)
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
		it.Event = new(DagoranftaplusfactoryNFTAPlusCreated)
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
func (it *DagoranftaplusfactoryNFTAPlusCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DagoranftaplusfactoryNFTAPlusCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DagoranftaplusfactoryNFTAPlusCreated represents a NFTAPlusCreated event raised by the Dagoranftaplusfactory contract.
type DagoranftaplusfactoryNFTAPlusCreated struct {
	NewContractAddress common.Address
	OwnerOF            common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNFTAPlusCreated is a free log retrieval operation binding the contract event 0x5a9d698669ede1fc3b7cd396725110acc39654752309fd13fcea2b70269ca3b9.
//
// Solidity: event NFTAPlusCreated(address newContractAddress, address ownerOF)
func (_Dagoranftaplusfactory *DagoranftaplusfactoryFilterer) FilterNFTAPlusCreated(opts *bind.FilterOpts) (*DagoranftaplusfactoryNFTAPlusCreatedIterator, error) {

	logs, sub, err := _Dagoranftaplusfactory.contract.FilterLogs(opts, "NFTAPlusCreated")
	if err != nil {
		return nil, err
	}
	return &DagoranftaplusfactoryNFTAPlusCreatedIterator{contract: _Dagoranftaplusfactory.contract, event: "NFTAPlusCreated", logs: logs, sub: sub}, nil
}

// WatchNFTAPlusCreated is a free log subscription operation binding the contract event 0x5a9d698669ede1fc3b7cd396725110acc39654752309fd13fcea2b70269ca3b9.
//
// Solidity: event NFTAPlusCreated(address newContractAddress, address ownerOF)
func (_Dagoranftaplusfactory *DagoranftaplusfactoryFilterer) WatchNFTAPlusCreated(opts *bind.WatchOpts, sink chan<- *DagoranftaplusfactoryNFTAPlusCreated) (event.Subscription, error) {

	logs, sub, err := _Dagoranftaplusfactory.contract.WatchLogs(opts, "NFTAPlusCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DagoranftaplusfactoryNFTAPlusCreated)
				if err := _Dagoranftaplusfactory.contract.UnpackLog(event, "NFTAPlusCreated", log); err != nil {
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

// ParseNFTAPlusCreated is a log parse operation binding the contract event 0x5a9d698669ede1fc3b7cd396725110acc39654752309fd13fcea2b70269ca3b9.
//
// Solidity: event NFTAPlusCreated(address newContractAddress, address ownerOF)
func (_Dagoranftaplusfactory *DagoranftaplusfactoryFilterer) ParseNFTAPlusCreated(log types.Log) (*DagoranftaplusfactoryNFTAPlusCreated, error) {
	event := new(DagoranftaplusfactoryNFTAPlusCreated)
	if err := _Dagoranftaplusfactory.contract.UnpackLog(event, "NFTAPlusCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DagoranftaplusfactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Dagoranftaplusfactory contract.
type DagoranftaplusfactoryOwnershipTransferredIterator struct {
	Event *DagoranftaplusfactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DagoranftaplusfactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DagoranftaplusfactoryOwnershipTransferred)
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
		it.Event = new(DagoranftaplusfactoryOwnershipTransferred)
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
func (it *DagoranftaplusfactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DagoranftaplusfactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DagoranftaplusfactoryOwnershipTransferred represents a OwnershipTransferred event raised by the Dagoranftaplusfactory contract.
type DagoranftaplusfactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dagoranftaplusfactory *DagoranftaplusfactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DagoranftaplusfactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dagoranftaplusfactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DagoranftaplusfactoryOwnershipTransferredIterator{contract: _Dagoranftaplusfactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dagoranftaplusfactory *DagoranftaplusfactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DagoranftaplusfactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dagoranftaplusfactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DagoranftaplusfactoryOwnershipTransferred)
				if err := _Dagoranftaplusfactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Dagoranftaplusfactory *DagoranftaplusfactoryFilterer) ParseOwnershipTransferred(log types.Log) (*DagoranftaplusfactoryOwnershipTransferred, error) {
	event := new(DagoranftaplusfactoryOwnershipTransferred)
	if err := _Dagoranftaplusfactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
