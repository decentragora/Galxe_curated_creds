// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dagorapowerplusnft

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

// PowerPlusNFTParams is an auto generated low-level Go binding around an user-defined struct.
type PowerPlusNFTParams struct {
	Name               string
	Symbol             string
	BaseURI            string
	BulkBuyLimit       uint16
	MaxAllowListAmount uint16
	RoyaltyBps         *big.Int
	MintPrice          *big.Int
	PresaleMintCost    *big.Int
	MaxSupply          *big.Int
	RoyaltyRecipient   common.Address
	NewOwner           common.Address
	MerkleRoot         [32]byte
}

// DagorapowerplusnftMetaData contains all meta data concerning the Dagorapowerplusnft contract.
var DagorapowerplusnftMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DagoraFactory__ExpiredMembership\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DagoraFactory__FailedToCreateContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"tier\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"neededTier\",\"type\":\"uint8\"}],\"name\":\"DagoraFactory__InvalidTier\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DagoraFactory__NotDAgoraMembershipsOwnerOrDelegate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DagoraFactory__TokenCreationPaused\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newContractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ownerOF\",\"type\":\"address\"}],\"name\":\"PowerPlusNFTACreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"contractsDeployed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"baseURI_\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"_bulkBuyLimit\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_maxAllowListAmount\",\"type\":\"uint16\"},{\"internalType\":\"uint96\",\"name\":\"_royaltyBps\",\"type\":\"uint96\"},{\"internalType\":\"uint256\",\"name\":\"_mintPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_presaleMintCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_royaltyRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structPowerPlusNFT.Params\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"createPowerPlusNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dAgoraMembershipsAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserContracts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dAgoraMembershipsAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minPowerNFTATier\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_minPowerNFTATier\",\"type\":\"uint8\"}],\"name\":\"setMinPowerNFTATier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"togglePaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userContracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DagorapowerplusnftABI is the input ABI used to generate the binding from.
// Deprecated: Use DagorapowerplusnftMetaData.ABI instead.
var DagorapowerplusnftABI = DagorapowerplusnftMetaData.ABI

// Dagorapowerplusnft is an auto generated Go binding around an Ethereum contract.
type Dagorapowerplusnft struct {
	DagorapowerplusnftCaller     // Read-only binding to the contract
	DagorapowerplusnftTransactor // Write-only binding to the contract
	DagorapowerplusnftFilterer   // Log filterer for contract events
}

// DagorapowerplusnftCaller is an auto generated read-only Go binding around an Ethereum contract.
type DagorapowerplusnftCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DagorapowerplusnftTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DagorapowerplusnftTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DagorapowerplusnftFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DagorapowerplusnftFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DagorapowerplusnftSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DagorapowerplusnftSession struct {
	Contract     *Dagorapowerplusnft // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DagorapowerplusnftCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DagorapowerplusnftCallerSession struct {
	Contract *DagorapowerplusnftCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// DagorapowerplusnftTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DagorapowerplusnftTransactorSession struct {
	Contract     *DagorapowerplusnftTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// DagorapowerplusnftRaw is an auto generated low-level Go binding around an Ethereum contract.
type DagorapowerplusnftRaw struct {
	Contract *Dagorapowerplusnft // Generic contract binding to access the raw methods on
}

// DagorapowerplusnftCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DagorapowerplusnftCallerRaw struct {
	Contract *DagorapowerplusnftCaller // Generic read-only contract binding to access the raw methods on
}

// DagorapowerplusnftTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DagorapowerplusnftTransactorRaw struct {
	Contract *DagorapowerplusnftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDagorapowerplusnft creates a new instance of Dagorapowerplusnft, bound to a specific deployed contract.
func NewDagorapowerplusnft(address common.Address, backend bind.ContractBackend) (*Dagorapowerplusnft, error) {
	contract, err := bindDagorapowerplusnft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dagorapowerplusnft{DagorapowerplusnftCaller: DagorapowerplusnftCaller{contract: contract}, DagorapowerplusnftTransactor: DagorapowerplusnftTransactor{contract: contract}, DagorapowerplusnftFilterer: DagorapowerplusnftFilterer{contract: contract}}, nil
}

// NewDagorapowerplusnftCaller creates a new read-only instance of Dagorapowerplusnft, bound to a specific deployed contract.
func NewDagorapowerplusnftCaller(address common.Address, caller bind.ContractCaller) (*DagorapowerplusnftCaller, error) {
	contract, err := bindDagorapowerplusnft(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DagorapowerplusnftCaller{contract: contract}, nil
}

// NewDagorapowerplusnftTransactor creates a new write-only instance of Dagorapowerplusnft, bound to a specific deployed contract.
func NewDagorapowerplusnftTransactor(address common.Address, transactor bind.ContractTransactor) (*DagorapowerplusnftTransactor, error) {
	contract, err := bindDagorapowerplusnft(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DagorapowerplusnftTransactor{contract: contract}, nil
}

// NewDagorapowerplusnftFilterer creates a new log filterer instance of Dagorapowerplusnft, bound to a specific deployed contract.
func NewDagorapowerplusnftFilterer(address common.Address, filterer bind.ContractFilterer) (*DagorapowerplusnftFilterer, error) {
	contract, err := bindDagorapowerplusnft(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DagorapowerplusnftFilterer{contract: contract}, nil
}

// bindDagorapowerplusnft binds a generic wrapper to an already deployed contract.
func bindDagorapowerplusnft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DagorapowerplusnftMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dagorapowerplusnft *DagorapowerplusnftRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dagorapowerplusnft.Contract.DagorapowerplusnftCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dagorapowerplusnft *DagorapowerplusnftRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dagorapowerplusnft.Contract.DagorapowerplusnftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dagorapowerplusnft *DagorapowerplusnftRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dagorapowerplusnft.Contract.DagorapowerplusnftTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dagorapowerplusnft *DagorapowerplusnftCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dagorapowerplusnft.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dagorapowerplusnft *DagorapowerplusnftTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dagorapowerplusnft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dagorapowerplusnft *DagorapowerplusnftTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dagorapowerplusnft.Contract.contract.Transact(opts, method, params...)
}

// ContractsDeployed is a free data retrieval call binding the contract method 0xf0a761e4.
//
// Solidity: function contractsDeployed() view returns(uint256)
func (_Dagorapowerplusnft *DagorapowerplusnftCaller) ContractsDeployed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dagorapowerplusnft.contract.Call(opts, &out, "contractsDeployed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ContractsDeployed is a free data retrieval call binding the contract method 0xf0a761e4.
//
// Solidity: function contractsDeployed() view returns(uint256)
func (_Dagorapowerplusnft *DagorapowerplusnftSession) ContractsDeployed() (*big.Int, error) {
	return _Dagorapowerplusnft.Contract.ContractsDeployed(&_Dagorapowerplusnft.CallOpts)
}

// ContractsDeployed is a free data retrieval call binding the contract method 0xf0a761e4.
//
// Solidity: function contractsDeployed() view returns(uint256)
func (_Dagorapowerplusnft *DagorapowerplusnftCallerSession) ContractsDeployed() (*big.Int, error) {
	return _Dagorapowerplusnft.Contract.ContractsDeployed(&_Dagorapowerplusnft.CallOpts)
}

// DAgoraMembershipsAddress is a free data retrieval call binding the contract method 0x514fd254.
//
// Solidity: function dAgoraMembershipsAddress() view returns(address)
func (_Dagorapowerplusnft *DagorapowerplusnftCaller) DAgoraMembershipsAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dagorapowerplusnft.contract.Call(opts, &out, "dAgoraMembershipsAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DAgoraMembershipsAddress is a free data retrieval call binding the contract method 0x514fd254.
//
// Solidity: function dAgoraMembershipsAddress() view returns(address)
func (_Dagorapowerplusnft *DagorapowerplusnftSession) DAgoraMembershipsAddress() (common.Address, error) {
	return _Dagorapowerplusnft.Contract.DAgoraMembershipsAddress(&_Dagorapowerplusnft.CallOpts)
}

// DAgoraMembershipsAddress is a free data retrieval call binding the contract method 0x514fd254.
//
// Solidity: function dAgoraMembershipsAddress() view returns(address)
func (_Dagorapowerplusnft *DagorapowerplusnftCallerSession) DAgoraMembershipsAddress() (common.Address, error) {
	return _Dagorapowerplusnft.Contract.DAgoraMembershipsAddress(&_Dagorapowerplusnft.CallOpts)
}

// GetUserContracts is a free data retrieval call binding the contract method 0x34c925f0.
//
// Solidity: function getUserContracts(address _user) view returns(address[])
func (_Dagorapowerplusnft *DagorapowerplusnftCaller) GetUserContracts(opts *bind.CallOpts, _user common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Dagorapowerplusnft.contract.Call(opts, &out, "getUserContracts", _user)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetUserContracts is a free data retrieval call binding the contract method 0x34c925f0.
//
// Solidity: function getUserContracts(address _user) view returns(address[])
func (_Dagorapowerplusnft *DagorapowerplusnftSession) GetUserContracts(_user common.Address) ([]common.Address, error) {
	return _Dagorapowerplusnft.Contract.GetUserContracts(&_Dagorapowerplusnft.CallOpts, _user)
}

// GetUserContracts is a free data retrieval call binding the contract method 0x34c925f0.
//
// Solidity: function getUserContracts(address _user) view returns(address[])
func (_Dagorapowerplusnft *DagorapowerplusnftCallerSession) GetUserContracts(_user common.Address) ([]common.Address, error) {
	return _Dagorapowerplusnft.Contract.GetUserContracts(&_Dagorapowerplusnft.CallOpts, _user)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Dagorapowerplusnft *DagorapowerplusnftCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Dagorapowerplusnft.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Dagorapowerplusnft *DagorapowerplusnftSession) IsPaused() (bool, error) {
	return _Dagorapowerplusnft.Contract.IsPaused(&_Dagorapowerplusnft.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Dagorapowerplusnft *DagorapowerplusnftCallerSession) IsPaused() (bool, error) {
	return _Dagorapowerplusnft.Contract.IsPaused(&_Dagorapowerplusnft.CallOpts)
}

// MinPowerNFTATier is a free data retrieval call binding the contract method 0xfdd17a4f.
//
// Solidity: function minPowerNFTATier() view returns(uint8)
func (_Dagorapowerplusnft *DagorapowerplusnftCaller) MinPowerNFTATier(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Dagorapowerplusnft.contract.Call(opts, &out, "minPowerNFTATier")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MinPowerNFTATier is a free data retrieval call binding the contract method 0xfdd17a4f.
//
// Solidity: function minPowerNFTATier() view returns(uint8)
func (_Dagorapowerplusnft *DagorapowerplusnftSession) MinPowerNFTATier() (uint8, error) {
	return _Dagorapowerplusnft.Contract.MinPowerNFTATier(&_Dagorapowerplusnft.CallOpts)
}

// MinPowerNFTATier is a free data retrieval call binding the contract method 0xfdd17a4f.
//
// Solidity: function minPowerNFTATier() view returns(uint8)
func (_Dagorapowerplusnft *DagorapowerplusnftCallerSession) MinPowerNFTATier() (uint8, error) {
	return _Dagorapowerplusnft.Contract.MinPowerNFTATier(&_Dagorapowerplusnft.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dagorapowerplusnft *DagorapowerplusnftCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dagorapowerplusnft.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dagorapowerplusnft *DagorapowerplusnftSession) Owner() (common.Address, error) {
	return _Dagorapowerplusnft.Contract.Owner(&_Dagorapowerplusnft.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dagorapowerplusnft *DagorapowerplusnftCallerSession) Owner() (common.Address, error) {
	return _Dagorapowerplusnft.Contract.Owner(&_Dagorapowerplusnft.CallOpts)
}

// UserContracts is a free data retrieval call binding the contract method 0xb789003d.
//
// Solidity: function userContracts(address , uint256 ) view returns(address)
func (_Dagorapowerplusnft *DagorapowerplusnftCaller) UserContracts(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Dagorapowerplusnft.contract.Call(opts, &out, "userContracts", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserContracts is a free data retrieval call binding the contract method 0xb789003d.
//
// Solidity: function userContracts(address , uint256 ) view returns(address)
func (_Dagorapowerplusnft *DagorapowerplusnftSession) UserContracts(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Dagorapowerplusnft.Contract.UserContracts(&_Dagorapowerplusnft.CallOpts, arg0, arg1)
}

// UserContracts is a free data retrieval call binding the contract method 0xb789003d.
//
// Solidity: function userContracts(address , uint256 ) view returns(address)
func (_Dagorapowerplusnft *DagorapowerplusnftCallerSession) UserContracts(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Dagorapowerplusnft.Contract.UserContracts(&_Dagorapowerplusnft.CallOpts, arg0, arg1)
}

// CreatePowerPlusNFT is a paid mutator transaction binding the contract method 0x80ccd727.
//
// Solidity: function createPowerPlusNFT((string,string,string,uint16,uint16,uint96,uint256,uint256,uint256,address,address,bytes32) params, uint256 _id) returns()
func (_Dagorapowerplusnft *DagorapowerplusnftTransactor) CreatePowerPlusNFT(opts *bind.TransactOpts, params PowerPlusNFTParams, _id *big.Int) (*types.Transaction, error) {
	return _Dagorapowerplusnft.contract.Transact(opts, "createPowerPlusNFT", params, _id)
}

// CreatePowerPlusNFT is a paid mutator transaction binding the contract method 0x80ccd727.
//
// Solidity: function createPowerPlusNFT((string,string,string,uint16,uint16,uint96,uint256,uint256,uint256,address,address,bytes32) params, uint256 _id) returns()
func (_Dagorapowerplusnft *DagorapowerplusnftSession) CreatePowerPlusNFT(params PowerPlusNFTParams, _id *big.Int) (*types.Transaction, error) {
	return _Dagorapowerplusnft.Contract.CreatePowerPlusNFT(&_Dagorapowerplusnft.TransactOpts, params, _id)
}

// CreatePowerPlusNFT is a paid mutator transaction binding the contract method 0x80ccd727.
//
// Solidity: function createPowerPlusNFT((string,string,string,uint16,uint16,uint96,uint256,uint256,uint256,address,address,bytes32) params, uint256 _id) returns()
func (_Dagorapowerplusnft *DagorapowerplusnftTransactorSession) CreatePowerPlusNFT(params PowerPlusNFTParams, _id *big.Int) (*types.Transaction, error) {
	return _Dagorapowerplusnft.Contract.CreatePowerPlusNFT(&_Dagorapowerplusnft.TransactOpts, params, _id)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _dAgoraMembershipsAddress) returns()
func (_Dagorapowerplusnft *DagorapowerplusnftTransactor) Initialize(opts *bind.TransactOpts, _dAgoraMembershipsAddress common.Address) (*types.Transaction, error) {
	return _Dagorapowerplusnft.contract.Transact(opts, "initialize", _dAgoraMembershipsAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _dAgoraMembershipsAddress) returns()
func (_Dagorapowerplusnft *DagorapowerplusnftSession) Initialize(_dAgoraMembershipsAddress common.Address) (*types.Transaction, error) {
	return _Dagorapowerplusnft.Contract.Initialize(&_Dagorapowerplusnft.TransactOpts, _dAgoraMembershipsAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _dAgoraMembershipsAddress) returns()
func (_Dagorapowerplusnft *DagorapowerplusnftTransactorSession) Initialize(_dAgoraMembershipsAddress common.Address) (*types.Transaction, error) {
	return _Dagorapowerplusnft.Contract.Initialize(&_Dagorapowerplusnft.TransactOpts, _dAgoraMembershipsAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dagorapowerplusnft *DagorapowerplusnftTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dagorapowerplusnft.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dagorapowerplusnft *DagorapowerplusnftSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dagorapowerplusnft.Contract.RenounceOwnership(&_Dagorapowerplusnft.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dagorapowerplusnft *DagorapowerplusnftTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dagorapowerplusnft.Contract.RenounceOwnership(&_Dagorapowerplusnft.TransactOpts)
}

// SetMinPowerNFTATier is a paid mutator transaction binding the contract method 0x04bad475.
//
// Solidity: function setMinPowerNFTATier(uint8 _minPowerNFTATier) returns()
func (_Dagorapowerplusnft *DagorapowerplusnftTransactor) SetMinPowerNFTATier(opts *bind.TransactOpts, _minPowerNFTATier uint8) (*types.Transaction, error) {
	return _Dagorapowerplusnft.contract.Transact(opts, "setMinPowerNFTATier", _minPowerNFTATier)
}

// SetMinPowerNFTATier is a paid mutator transaction binding the contract method 0x04bad475.
//
// Solidity: function setMinPowerNFTATier(uint8 _minPowerNFTATier) returns()
func (_Dagorapowerplusnft *DagorapowerplusnftSession) SetMinPowerNFTATier(_minPowerNFTATier uint8) (*types.Transaction, error) {
	return _Dagorapowerplusnft.Contract.SetMinPowerNFTATier(&_Dagorapowerplusnft.TransactOpts, _minPowerNFTATier)
}

// SetMinPowerNFTATier is a paid mutator transaction binding the contract method 0x04bad475.
//
// Solidity: function setMinPowerNFTATier(uint8 _minPowerNFTATier) returns()
func (_Dagorapowerplusnft *DagorapowerplusnftTransactorSession) SetMinPowerNFTATier(_minPowerNFTATier uint8) (*types.Transaction, error) {
	return _Dagorapowerplusnft.Contract.SetMinPowerNFTATier(&_Dagorapowerplusnft.TransactOpts, _minPowerNFTATier)
}

// TogglePaused is a paid mutator transaction binding the contract method 0x36566f06.
//
// Solidity: function togglePaused() returns()
func (_Dagorapowerplusnft *DagorapowerplusnftTransactor) TogglePaused(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dagorapowerplusnft.contract.Transact(opts, "togglePaused")
}

// TogglePaused is a paid mutator transaction binding the contract method 0x36566f06.
//
// Solidity: function togglePaused() returns()
func (_Dagorapowerplusnft *DagorapowerplusnftSession) TogglePaused() (*types.Transaction, error) {
	return _Dagorapowerplusnft.Contract.TogglePaused(&_Dagorapowerplusnft.TransactOpts)
}

// TogglePaused is a paid mutator transaction binding the contract method 0x36566f06.
//
// Solidity: function togglePaused() returns()
func (_Dagorapowerplusnft *DagorapowerplusnftTransactorSession) TogglePaused() (*types.Transaction, error) {
	return _Dagorapowerplusnft.Contract.TogglePaused(&_Dagorapowerplusnft.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dagorapowerplusnft *DagorapowerplusnftTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Dagorapowerplusnft.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dagorapowerplusnft *DagorapowerplusnftSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dagorapowerplusnft.Contract.TransferOwnership(&_Dagorapowerplusnft.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dagorapowerplusnft *DagorapowerplusnftTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dagorapowerplusnft.Contract.TransferOwnership(&_Dagorapowerplusnft.TransactOpts, newOwner)
}

// DagorapowerplusnftInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Dagorapowerplusnft contract.
type DagorapowerplusnftInitializedIterator struct {
	Event *DagorapowerplusnftInitialized // Event containing the contract specifics and raw log

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
func (it *DagorapowerplusnftInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DagorapowerplusnftInitialized)
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
		it.Event = new(DagorapowerplusnftInitialized)
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
func (it *DagorapowerplusnftInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DagorapowerplusnftInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DagorapowerplusnftInitialized represents a Initialized event raised by the Dagorapowerplusnft contract.
type DagorapowerplusnftInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Dagorapowerplusnft *DagorapowerplusnftFilterer) FilterInitialized(opts *bind.FilterOpts) (*DagorapowerplusnftInitializedIterator, error) {

	logs, sub, err := _Dagorapowerplusnft.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DagorapowerplusnftInitializedIterator{contract: _Dagorapowerplusnft.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Dagorapowerplusnft *DagorapowerplusnftFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DagorapowerplusnftInitialized) (event.Subscription, error) {

	logs, sub, err := _Dagorapowerplusnft.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DagorapowerplusnftInitialized)
				if err := _Dagorapowerplusnft.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Dagorapowerplusnft *DagorapowerplusnftFilterer) ParseInitialized(log types.Log) (*DagorapowerplusnftInitialized, error) {
	event := new(DagorapowerplusnftInitialized)
	if err := _Dagorapowerplusnft.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DagorapowerplusnftOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Dagorapowerplusnft contract.
type DagorapowerplusnftOwnershipTransferredIterator struct {
	Event *DagorapowerplusnftOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DagorapowerplusnftOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DagorapowerplusnftOwnershipTransferred)
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
		it.Event = new(DagorapowerplusnftOwnershipTransferred)
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
func (it *DagorapowerplusnftOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DagorapowerplusnftOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DagorapowerplusnftOwnershipTransferred represents a OwnershipTransferred event raised by the Dagorapowerplusnft contract.
type DagorapowerplusnftOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dagorapowerplusnft *DagorapowerplusnftFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DagorapowerplusnftOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dagorapowerplusnft.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DagorapowerplusnftOwnershipTransferredIterator{contract: _Dagorapowerplusnft.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dagorapowerplusnft *DagorapowerplusnftFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DagorapowerplusnftOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dagorapowerplusnft.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DagorapowerplusnftOwnershipTransferred)
				if err := _Dagorapowerplusnft.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Dagorapowerplusnft *DagorapowerplusnftFilterer) ParseOwnershipTransferred(log types.Log) (*DagorapowerplusnftOwnershipTransferred, error) {
	event := new(DagorapowerplusnftOwnershipTransferred)
	if err := _Dagorapowerplusnft.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DagorapowerplusnftPowerPlusNFTACreatedIterator is returned from FilterPowerPlusNFTACreated and is used to iterate over the raw logs and unpacked data for PowerPlusNFTACreated events raised by the Dagorapowerplusnft contract.
type DagorapowerplusnftPowerPlusNFTACreatedIterator struct {
	Event *DagorapowerplusnftPowerPlusNFTACreated // Event containing the contract specifics and raw log

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
func (it *DagorapowerplusnftPowerPlusNFTACreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DagorapowerplusnftPowerPlusNFTACreated)
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
		it.Event = new(DagorapowerplusnftPowerPlusNFTACreated)
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
func (it *DagorapowerplusnftPowerPlusNFTACreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DagorapowerplusnftPowerPlusNFTACreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DagorapowerplusnftPowerPlusNFTACreated represents a PowerPlusNFTACreated event raised by the Dagorapowerplusnft contract.
type DagorapowerplusnftPowerPlusNFTACreated struct {
	NewContractAddress common.Address
	OwnerOF            common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPowerPlusNFTACreated is a free log retrieval operation binding the contract event 0x6f309ab3b1bd382131425f1f01e10b4ff38eb6a5a0002c0c2013fccfdacc352a.
//
// Solidity: event PowerPlusNFTACreated(address newContractAddress, address ownerOF)
func (_Dagorapowerplusnft *DagorapowerplusnftFilterer) FilterPowerPlusNFTACreated(opts *bind.FilterOpts) (*DagorapowerplusnftPowerPlusNFTACreatedIterator, error) {

	logs, sub, err := _Dagorapowerplusnft.contract.FilterLogs(opts, "PowerPlusNFTACreated")
	if err != nil {
		return nil, err
	}
	return &DagorapowerplusnftPowerPlusNFTACreatedIterator{contract: _Dagorapowerplusnft.contract, event: "PowerPlusNFTACreated", logs: logs, sub: sub}, nil
}

// WatchPowerPlusNFTACreated is a free log subscription operation binding the contract event 0x6f309ab3b1bd382131425f1f01e10b4ff38eb6a5a0002c0c2013fccfdacc352a.
//
// Solidity: event PowerPlusNFTACreated(address newContractAddress, address ownerOF)
func (_Dagorapowerplusnft *DagorapowerplusnftFilterer) WatchPowerPlusNFTACreated(opts *bind.WatchOpts, sink chan<- *DagorapowerplusnftPowerPlusNFTACreated) (event.Subscription, error) {

	logs, sub, err := _Dagorapowerplusnft.contract.WatchLogs(opts, "PowerPlusNFTACreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DagorapowerplusnftPowerPlusNFTACreated)
				if err := _Dagorapowerplusnft.contract.UnpackLog(event, "PowerPlusNFTACreated", log); err != nil {
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

// ParsePowerPlusNFTACreated is a log parse operation binding the contract event 0x6f309ab3b1bd382131425f1f01e10b4ff38eb6a5a0002c0c2013fccfdacc352a.
//
// Solidity: event PowerPlusNFTACreated(address newContractAddress, address ownerOF)
func (_Dagorapowerplusnft *DagorapowerplusnftFilterer) ParsePowerPlusNFTACreated(log types.Log) (*DagorapowerplusnftPowerPlusNFTACreated, error) {
	event := new(DagorapowerplusnftPowerPlusNFTACreated)
	if err := _Dagorapowerplusnft.contract.UnpackLog(event, "PowerPlusNFTACreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
