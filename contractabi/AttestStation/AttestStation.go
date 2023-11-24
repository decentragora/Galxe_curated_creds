// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package atteststation

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

// AtteststationMetaData contains all meta data concerning the Atteststation contract.
var AtteststationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"about\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"}],\"name\":\"AttestationCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_about\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_val\",\"type\":\"bytes\"}],\"name\":\"attest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"attestations\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AtteststationABI is the input ABI used to generate the binding from.
// Deprecated: Use AtteststationMetaData.ABI instead.
var AtteststationABI = AtteststationMetaData.ABI

// Atteststation is an auto generated Go binding around an Ethereum contract.
type Atteststation struct {
	AtteststationCaller     // Read-only binding to the contract
	AtteststationTransactor // Write-only binding to the contract
	AtteststationFilterer   // Log filterer for contract events
}

// AtteststationCaller is an auto generated read-only Go binding around an Ethereum contract.
type AtteststationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AtteststationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AtteststationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AtteststationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AtteststationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AtteststationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AtteststationSession struct {
	Contract     *Atteststation    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AtteststationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AtteststationCallerSession struct {
	Contract *AtteststationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AtteststationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AtteststationTransactorSession struct {
	Contract     *AtteststationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AtteststationRaw is an auto generated low-level Go binding around an Ethereum contract.
type AtteststationRaw struct {
	Contract *Atteststation // Generic contract binding to access the raw methods on
}

// AtteststationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AtteststationCallerRaw struct {
	Contract *AtteststationCaller // Generic read-only contract binding to access the raw methods on
}

// AtteststationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AtteststationTransactorRaw struct {
	Contract *AtteststationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAtteststation creates a new instance of Atteststation, bound to a specific deployed contract.
func NewAtteststation(address common.Address, backend bind.ContractBackend) (*Atteststation, error) {
	contract, err := bindAtteststation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Atteststation{AtteststationCaller: AtteststationCaller{contract: contract}, AtteststationTransactor: AtteststationTransactor{contract: contract}, AtteststationFilterer: AtteststationFilterer{contract: contract}}, nil
}

// NewAtteststationCaller creates a new read-only instance of Atteststation, bound to a specific deployed contract.
func NewAtteststationCaller(address common.Address, caller bind.ContractCaller) (*AtteststationCaller, error) {
	contract, err := bindAtteststation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AtteststationCaller{contract: contract}, nil
}

// NewAtteststationTransactor creates a new write-only instance of Atteststation, bound to a specific deployed contract.
func NewAtteststationTransactor(address common.Address, transactor bind.ContractTransactor) (*AtteststationTransactor, error) {
	contract, err := bindAtteststation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AtteststationTransactor{contract: contract}, nil
}

// NewAtteststationFilterer creates a new log filterer instance of Atteststation, bound to a specific deployed contract.
func NewAtteststationFilterer(address common.Address, filterer bind.ContractFilterer) (*AtteststationFilterer, error) {
	contract, err := bindAtteststation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AtteststationFilterer{contract: contract}, nil
}

// bindAtteststation binds a generic wrapper to an already deployed contract.
func bindAtteststation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AtteststationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Atteststation *AtteststationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Atteststation.Contract.AtteststationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Atteststation *AtteststationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Atteststation.Contract.AtteststationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Atteststation *AtteststationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Atteststation.Contract.AtteststationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Atteststation *AtteststationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Atteststation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Atteststation *AtteststationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Atteststation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Atteststation *AtteststationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Atteststation.Contract.contract.Transact(opts, method, params...)
}

// Attestations is a free data retrieval call binding the contract method 0x29b42cb5.
//
// Solidity: function attestations(address , address , bytes32 ) view returns(bytes)
func (_Atteststation *AtteststationCaller) Attestations(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Atteststation.contract.Call(opts, &out, "attestations", arg0, arg1, arg2)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Attestations is a free data retrieval call binding the contract method 0x29b42cb5.
//
// Solidity: function attestations(address , address , bytes32 ) view returns(bytes)
func (_Atteststation *AtteststationSession) Attestations(arg0 common.Address, arg1 common.Address, arg2 [32]byte) ([]byte, error) {
	return _Atteststation.Contract.Attestations(&_Atteststation.CallOpts, arg0, arg1, arg2)
}

// Attestations is a free data retrieval call binding the contract method 0x29b42cb5.
//
// Solidity: function attestations(address , address , bytes32 ) view returns(bytes)
func (_Atteststation *AtteststationCallerSession) Attestations(arg0 common.Address, arg1 common.Address, arg2 [32]byte) ([]byte, error) {
	return _Atteststation.Contract.Attestations(&_Atteststation.CallOpts, arg0, arg1, arg2)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Atteststation *AtteststationCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Atteststation.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Atteststation *AtteststationSession) Version() (string, error) {
	return _Atteststation.Contract.Version(&_Atteststation.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Atteststation *AtteststationCallerSession) Version() (string, error) {
	return _Atteststation.Contract.Version(&_Atteststation.CallOpts)
}

// Attest is a paid mutator transaction binding the contract method 0x702b9dee.
//
// Solidity: function attest(address _about, bytes32 _key, bytes _val) returns()
func (_Atteststation *AtteststationTransactor) Attest(opts *bind.TransactOpts, _about common.Address, _key [32]byte, _val []byte) (*types.Transaction, error) {
	return _Atteststation.contract.Transact(opts, "attest", _about, _key, _val)
}

// Attest is a paid mutator transaction binding the contract method 0x702b9dee.
//
// Solidity: function attest(address _about, bytes32 _key, bytes _val) returns()
func (_Atteststation *AtteststationSession) Attest(_about common.Address, _key [32]byte, _val []byte) (*types.Transaction, error) {
	return _Atteststation.Contract.Attest(&_Atteststation.TransactOpts, _about, _key, _val)
}

// Attest is a paid mutator transaction binding the contract method 0x702b9dee.
//
// Solidity: function attest(address _about, bytes32 _key, bytes _val) returns()
func (_Atteststation *AtteststationTransactorSession) Attest(_about common.Address, _key [32]byte, _val []byte) (*types.Transaction, error) {
	return _Atteststation.Contract.Attest(&_Atteststation.TransactOpts, _about, _key, _val)
}

// AtteststationAttestationCreatedIterator is returned from FilterAttestationCreated and is used to iterate over the raw logs and unpacked data for AttestationCreated events raised by the Atteststation contract.
type AtteststationAttestationCreatedIterator struct {
	Event *AtteststationAttestationCreated // Event containing the contract specifics and raw log

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
func (it *AtteststationAttestationCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtteststationAttestationCreated)
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
		it.Event = new(AtteststationAttestationCreated)
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
func (it *AtteststationAttestationCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtteststationAttestationCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtteststationAttestationCreated represents a AttestationCreated event raised by the Atteststation contract.
type AtteststationAttestationCreated struct {
	Creator common.Address
	About   common.Address
	Key     [32]byte
	Val     []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAttestationCreated is a free log retrieval operation binding the contract event 0x28710dfecab43d1e29e02aa56b2e1e610c0bae19135c9cf7a83a1adb6df96d85.
//
// Solidity: event AttestationCreated(address indexed creator, address indexed about, bytes32 indexed key, bytes val)
func (_Atteststation *AtteststationFilterer) FilterAttestationCreated(opts *bind.FilterOpts, creator []common.Address, about []common.Address, key [][32]byte) (*AtteststationAttestationCreatedIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var aboutRule []interface{}
	for _, aboutItem := range about {
		aboutRule = append(aboutRule, aboutItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _Atteststation.contract.FilterLogs(opts, "AttestationCreated", creatorRule, aboutRule, keyRule)
	if err != nil {
		return nil, err
	}
	return &AtteststationAttestationCreatedIterator{contract: _Atteststation.contract, event: "AttestationCreated", logs: logs, sub: sub}, nil
}

// WatchAttestationCreated is a free log subscription operation binding the contract event 0x28710dfecab43d1e29e02aa56b2e1e610c0bae19135c9cf7a83a1adb6df96d85.
//
// Solidity: event AttestationCreated(address indexed creator, address indexed about, bytes32 indexed key, bytes val)
func (_Atteststation *AtteststationFilterer) WatchAttestationCreated(opts *bind.WatchOpts, sink chan<- *AtteststationAttestationCreated, creator []common.Address, about []common.Address, key [][32]byte) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var aboutRule []interface{}
	for _, aboutItem := range about {
		aboutRule = append(aboutRule, aboutItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _Atteststation.contract.WatchLogs(opts, "AttestationCreated", creatorRule, aboutRule, keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtteststationAttestationCreated)
				if err := _Atteststation.contract.UnpackLog(event, "AttestationCreated", log); err != nil {
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

// ParseAttestationCreated is a log parse operation binding the contract event 0x28710dfecab43d1e29e02aa56b2e1e610c0bae19135c9cf7a83a1adb6df96d85.
//
// Solidity: event AttestationCreated(address indexed creator, address indexed about, bytes32 indexed key, bytes val)
func (_Atteststation *AtteststationFilterer) ParseAttestationCreated(log types.Log) (*AtteststationAttestationCreated, error) {
	event := new(AtteststationAttestationCreated)
	if err := _Atteststation.contract.UnpackLog(event, "AttestationCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
