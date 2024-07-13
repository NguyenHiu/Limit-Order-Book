// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lOB

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

// LOBOrder is an auto generated low-level Go binding around an user-defined struct.
type LOBOrder struct {
	OrderID [16]byte
	Price   *big.Int
	Amount  *big.Int
	Side    bool
}

// LOBMetaData contains all meta data concerning the LOB contract.
var LOBMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MyMatchEvent\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes16\",\"name\":\"orderID\",\"type\":\"bytes16\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"side\",\"type\":\"bool\"}],\"internalType\":\"structLOB.Order\",\"name\":\"newOrder\",\"type\":\"tuple\"}],\"name\":\"addOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506115798061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610029575f3560e01c80639cfa8aa61461002d575b5f80fd5b61004760048036038101906100429190611357565b610049565b005b5f60801b6fffffffffffffffffffffffffffffffff1916815f01516fffffffffffffffffffffffffffffffff1916036100bd577f4d85d09fd9972adb50adbf062a43da84abfebb9efd2340d5dd19e02d4f4d2683815f015182604001516040516100b49291906113a0565b60405180910390a15b8060600151156100d5576100d0816100ea565b6100df565b6100de81610694565b5b6100e7610c48565b50565b5f808054905003610184575f81908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f820151815f015f6101000a8154816fffffffffffffffffffffffffffffffff021916908360801c021790555060208201518160010155604082015181600201556060820151816003015f6101000a81548160ff0219169083151502179055505050610690565b60015f80549050036103965780602001515f80815481106101a8576101a76113c7565b5b905f5260205f209060040201600101541015610306575f805f815481106101d2576101d16113c7565b5b905f5260205f209060040201908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f82015f9054906101000a900460801b815f015f6101000a8154816fffffffffffffffffffffffffffffffff021916908360801c02179055506001820154816001015560028201548160020155600382015f9054906101000a900460ff16816003015f6101000a81548160ff0219169083151502179055505050805f8081548110610292576102916113c7565b5b905f5260205f2090600402015f820151815f015f6101000a8154816fffffffffffffffffffffffffffffffff021916908360801c021790555060208201518160010155604082015181600201556060820151816003015f6101000a81548160ff021916908315150217905550905050610391565b5f81908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f820151815f015f6101000a8154816fffffffffffffffffffffffffffffffff021916908360801c021790555060208201518160010155604082015181600201556060820151816003015f6101000a81548160ff02191690831515021790555050505b61068f565b5f5b5f805490508160ff1610156106035781602001515f8260ff16815481106103c2576103c16113c7565b5b905f5260205f2090600402016001015410156105f0575f82908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f820151815f015f6101000a8154816fffffffffffffffffffffffffffffffff021916908360801c021790555060208201518160010155604082015181600201556060820151816003015f6101000a81548160ff02191690831515021790555050505f60015f80549050610474919061142d565b90505b8160ff168160ff161115610563575f600182610493919061142d565b60ff16815481106104a7576104a66113c7565b5b905f5260205f2090600402015f8260ff16815481106104c9576104c86113c7565b5b905f5260205f2090600402015f82015f9054906101000a900460801b815f015f6101000a8154816fffffffffffffffffffffffffffffffff021916908360801c02179055506001820154816001015560028201548160020155600382015f9054906101000a900460ff16816003015f6101000a81548160ff021916908315150217905550905050808061055b90611461565b915050610477565b50815f8260ff168154811061057b5761057a6113c7565b5b905f5260205f2090600402015f820151815f015f6101000a8154816fffffffffffffffffffffffffffffffff021916908360801c021790555060208201518160010155604082015181600201556060820151816003015f6101000a81548160ff02191690831515021790555090505050610691565b80806105fb90611488565b915050610398565b505f81908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f820151815f015f6101000a8154816fffffffffffffffffffffffffffffffff021916908360801c021790555060208201518160010155604082015181600201556060820151816003015f6101000a81548160ff02191690831515021790555050505b5b5b50565b5f6001805490500361073057600181908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f820151815f015f6101000a8154816fffffffffffffffffffffffffffffffff021916908360801c021790555060208201518160010155604082015181600201556060820151816003015f6101000a81548160ff0219169083151502179055505050610c44565b600180805490500361094657806020015160015f81548110610755576107546113c7565b5b905f5260205f2090600402016001015411156108b5576001805f815481106107805761077f6113c7565b5b905f5260205f209060040201908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f82015f9054906101000a900460801b815f015f6101000a8154816fffffffffffffffffffffffffffffffff021916908360801c02179055506001820154816001015560028201548160020155600382015f9054906101000a900460ff16816003015f6101000a81548160ff02191690831515021790555050508060015f81548110610841576108406113c7565b5b905f5260205f2090600402015f820151815f015f6101000a8154816fffffffffffffffffffffffffffffffff021916908360801c021790555060208201518160010155604082015181600201556060820151816003015f6101000a81548160ff021916908315150217905550905050610941565b600181908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f820151815f015f6101000a8154816fffffffffffffffffffffffffffffffff021916908360801c021790555060208201518160010155604082015181600201556060820151816003015f6101000a81548160ff02191690831515021790555050505b610c43565b5f5b6001805490508160ff161015610bb757816020015160018260ff1681548110610974576109736113c7565b5b905f5260205f209060040201600101541115610ba457600182908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f820151815f015f6101000a8154816fffffffffffffffffffffffffffffffff021916908360801c021790555060208201518160010155604082015181600201556060820151816003015f6101000a81548160ff02191690831515021790555050505f60018080549050610a27919061142d565b90505b8160ff168160ff161115610b175760018082610a46919061142d565b60ff1681548110610a5a57610a596113c7565b5b905f5260205f20906004020160018260ff1681548110610a7d57610a7c6113c7565b5b905f5260205f2090600402015f82015f9054906101000a900460801b815f015f6101000a8154816fffffffffffffffffffffffffffffffff021916908360801c02179055506001820154816001015560028201548160020155600382015f9054906101000a900460ff16816003015f6101000a81548160ff0219169083151502179055509050508080610b0f90611461565b915050610a2a565b50815f8260ff1681548110610b2f57610b2e6113c7565b5b905f5260205f2090600402015f820151815f015f6101000a8154816fffffffffffffffffffffffffffffffff021916908360801c021790555060208201518160010155604082015181600201556060820151816003015f6101000a81548160ff02191690831515021790555090505050610c45565b8080610baf90611488565b915050610948565b505f81908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f820151815f015f6101000a8154816fffffffffffffffffffffffffffffffff021916908360801c021790555060208201518160010155604082015181600201556060820151816003015f6101000a81548160ff02191690831515021790555050505b5b5b50565b5b610c51611119565b15611117575f805f81548110610c6a57610c696113c7565b5b905f5260205f20906004020160020154905060015f81548110610c9057610c8f6113c7565b5b905f5260205f20906004020160020154811115610cce5760015f81548110610cbb57610cba6113c7565b5b905f5260205f2090600402016002015490505b7f4d85d09fd9972adb50adbf062a43da84abfebb9efd2340d5dd19e02d4f4d26835f8081548110610d0257610d016113c7565b5b905f5260205f2090600402015f015f9054906101000a900460801b82604051610d2c9291906113a0565b60405180910390a17f4d85d09fd9972adb50adbf062a43da84abfebb9efd2340d5dd19e02d4f4d268360015f81548110610d6957610d686113c7565b5b905f5260205f2090600402015f015f9054906101000a900460801b82604051610d939291906113a0565b60405180910390a1805f8081548110610daf57610dae6113c7565b5b905f5260205f2090600402016002015f828254610dcc91906114b0565b925050819055508060015f81548110610de857610de76113c7565b5b905f5260205f2090600402016002015f828254610e0591906114b0565b925050819055505f805f81548110610e2057610e1f6113c7565b5b905f5260205f2090600402016002015403610f8d575f5b60015f80549050610e4891906114b0565b811015610f21575f600182610e5d91906114e3565b81548110610e6e57610e6d6113c7565b5b905f5260205f2090600402015f8281548110610e8d57610e8c6113c7565b5b905f5260205f2090600402015f82015f9054906101000a900460801b815f015f6101000a8154816fffffffffffffffffffffffffffffffff021916908360801c02179055506001820154816001015560028201548160020155600382015f9054906101000a900460ff16816003015f6101000a81548160ff0219169083151502179055509050508080600101915050610e37565b505f805480610f3357610f32611516565b5b600190038181905f5260205f2090600402015f8082015f6101000a8154906fffffffffffffffffffffffffffffffff0219169055600182015f9055600282015f9055600382015f6101000a81549060ff0219169055505090555b5f60015f81548110610fa257610fa16113c7565b5b905f5260205f2090600402016002015403611111575f5b60018080549050610fca91906114b0565b8110156110a45760018082610fdf91906114e3565b81548110610ff057610fef6113c7565b5b905f5260205f209060040201600182815481106110105761100f6113c7565b5b905f5260205f2090600402015f82015f9054906101000a900460801b815f015f6101000a8154816fffffffffffffffffffffffffffffffff021916908360801c02179055506001820154816001015560028201548160020155600382015f9054906101000a900460ff16816003015f6101000a81548160ff0219169083151502179055509050508080600101915050610fb9565b5060018054806110b7576110b6611516565b5b600190038181905f5260205f2090600402015f8082015f6101000a8154906fffffffffffffffffffffffffffffffff0219169055600182015f9055600282015f9055600382015f6101000a81549060ff0219169055505090555b50610c49565b565b5f805f80549050148061113057505f600180549050145b1561113d575f9050611189565b60015f81548110611151576111506113c7565b5b905f5260205f209060040201600101545f8081548110611174576111736113c7565b5b905f5260205f20906004020160010154101590505b90565b5f604051905090565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6111e38261119d565b810181811067ffffffffffffffff82111715611202576112016111ad565b5b80604052505050565b5f61121461118c565b905061122082826111da565b919050565b5f7fffffffffffffffffffffffffffffffff0000000000000000000000000000000082169050919050565b61125981611225565b8114611263575f80fd5b50565b5f8135905061127481611250565b92915050565b5f819050919050565b61128c8161127a565b8114611296575f80fd5b50565b5f813590506112a781611283565b92915050565b5f8115159050919050565b6112c1816112ad565b81146112cb575f80fd5b50565b5f813590506112dc816112b8565b92915050565b5f608082840312156112f7576112f6611199565b5b611301608061120b565b90505f61131084828501611266565b5f83015250602061132384828501611299565b602083015250604061133784828501611299565b604083015250606061134b848285016112ce565b60608301525092915050565b5f6080828403121561136c5761136b611195565b5b5f611379848285016112e2565b91505092915050565b61138b81611225565b82525050565b61139a8161127a565b82525050565b5f6040820190506113b35f830185611382565b6113c06020830184611391565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f60ff82169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611437826113f4565b9150611442836113f4565b9250828203905060ff81111561145b5761145a611400565b5b92915050565b5f61146b826113f4565b91505f820361147d5761147c611400565b5b600182039050919050565b5f611492826113f4565b915060ff82036114a5576114a4611400565b5b600182019050919050565b5f6114ba8261127a565b91506114c58361127a565b92508282039050818111156114dd576114dc611400565b5b92915050565b5f6114ed8261127a565b91506114f88361127a565b92508282019050808211156115105761150f611400565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffdfea2646970667358221220d864e0d4092661901a1a49b3194ad204eb7ee88e9542238d4d0ae9348014245264736f6c63430008180033",
}

// LOBABI is the input ABI used to generate the binding from.
// Deprecated: Use LOBMetaData.ABI instead.
var LOBABI = LOBMetaData.ABI

// LOBBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LOBMetaData.Bin instead.
var LOBBin = LOBMetaData.Bin

// DeployLOB deploys a new Ethereum contract, binding an instance of LOB to it.
func DeployLOB(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LOB, error) {
	parsed, err := LOBMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LOBBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LOB{LOBCaller: LOBCaller{contract: contract}, LOBTransactor: LOBTransactor{contract: contract}, LOBFilterer: LOBFilterer{contract: contract}}, nil
}

// LOB is an auto generated Go binding around an Ethereum contract.
type LOB struct {
	LOBCaller     // Read-only binding to the contract
	LOBTransactor // Write-only binding to the contract
	LOBFilterer   // Log filterer for contract events
}

// LOBCaller is an auto generated read-only Go binding around an Ethereum contract.
type LOBCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LOBTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LOBTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LOBFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LOBFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LOBSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LOBSession struct {
	Contract     *LOB              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LOBCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LOBCallerSession struct {
	Contract *LOBCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LOBTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LOBTransactorSession struct {
	Contract     *LOBTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LOBRaw is an auto generated low-level Go binding around an Ethereum contract.
type LOBRaw struct {
	Contract *LOB // Generic contract binding to access the raw methods on
}

// LOBCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LOBCallerRaw struct {
	Contract *LOBCaller // Generic read-only contract binding to access the raw methods on
}

// LOBTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LOBTransactorRaw struct {
	Contract *LOBTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLOB creates a new instance of LOB, bound to a specific deployed contract.
func NewLOB(address common.Address, backend bind.ContractBackend) (*LOB, error) {
	contract, err := bindLOB(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LOB{LOBCaller: LOBCaller{contract: contract}, LOBTransactor: LOBTransactor{contract: contract}, LOBFilterer: LOBFilterer{contract: contract}}, nil
}

// NewLOBCaller creates a new read-only instance of LOB, bound to a specific deployed contract.
func NewLOBCaller(address common.Address, caller bind.ContractCaller) (*LOBCaller, error) {
	contract, err := bindLOB(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LOBCaller{contract: contract}, nil
}

// NewLOBTransactor creates a new write-only instance of LOB, bound to a specific deployed contract.
func NewLOBTransactor(address common.Address, transactor bind.ContractTransactor) (*LOBTransactor, error) {
	contract, err := bindLOB(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LOBTransactor{contract: contract}, nil
}

// NewLOBFilterer creates a new log filterer instance of LOB, bound to a specific deployed contract.
func NewLOBFilterer(address common.Address, filterer bind.ContractFilterer) (*LOBFilterer, error) {
	contract, err := bindLOB(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LOBFilterer{contract: contract}, nil
}

// bindLOB binds a generic wrapper to an already deployed contract.
func bindLOB(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LOBMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LOB *LOBRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LOB.Contract.LOBCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LOB *LOBRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LOB.Contract.LOBTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LOB *LOBRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LOB.Contract.LOBTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LOB *LOBCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LOB.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LOB *LOBTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LOB.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LOB *LOBTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LOB.Contract.contract.Transact(opts, method, params...)
}

// AddOrder is a paid mutator transaction binding the contract method 0x9cfa8aa6.
//
// Solidity: function addOrder((bytes16,uint256,uint256,bool) newOrder) returns()
func (_LOB *LOBTransactor) AddOrder(opts *bind.TransactOpts, newOrder LOBOrder) (*types.Transaction, error) {
	return _LOB.contract.Transact(opts, "addOrder", newOrder)
}

// AddOrder is a paid mutator transaction binding the contract method 0x9cfa8aa6.
//
// Solidity: function addOrder((bytes16,uint256,uint256,bool) newOrder) returns()
func (_LOB *LOBSession) AddOrder(newOrder LOBOrder) (*types.Transaction, error) {
	return _LOB.Contract.AddOrder(&_LOB.TransactOpts, newOrder)
}

// AddOrder is a paid mutator transaction binding the contract method 0x9cfa8aa6.
//
// Solidity: function addOrder((bytes16,uint256,uint256,bool) newOrder) returns()
func (_LOB *LOBTransactorSession) AddOrder(newOrder LOBOrder) (*types.Transaction, error) {
	return _LOB.Contract.AddOrder(&_LOB.TransactOpts, newOrder)
}

// LOBMyMatchEventIterator is returned from FilterMyMatchEvent and is used to iterate over the raw logs and unpacked data for MyMatchEvent events raised by the LOB contract.
type LOBMyMatchEventIterator struct {
	Event *LOBMyMatchEvent // Event containing the contract specifics and raw log

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
func (it *LOBMyMatchEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LOBMyMatchEvent)
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
		it.Event = new(LOBMyMatchEvent)
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
func (it *LOBMyMatchEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LOBMyMatchEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LOBMyMatchEvent represents a MyMatchEvent event raised by the LOB contract.
type LOBMyMatchEvent struct {
	Arg0 [16]byte
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMyMatchEvent is a free log retrieval operation binding the contract event 0x4d85d09fd9972adb50adbf062a43da84abfebb9efd2340d5dd19e02d4f4d2683.
//
// Solidity: event MyMatchEvent(bytes16 arg0, uint256 arg1)
func (_LOB *LOBFilterer) FilterMyMatchEvent(opts *bind.FilterOpts) (*LOBMyMatchEventIterator, error) {

	logs, sub, err := _LOB.contract.FilterLogs(opts, "MyMatchEvent")
	if err != nil {
		return nil, err
	}
	return &LOBMyMatchEventIterator{contract: _LOB.contract, event: "MyMatchEvent", logs: logs, sub: sub}, nil
}

// WatchMyMatchEvent is a free log subscription operation binding the contract event 0x4d85d09fd9972adb50adbf062a43da84abfebb9efd2340d5dd19e02d4f4d2683.
//
// Solidity: event MyMatchEvent(bytes16 arg0, uint256 arg1)
func (_LOB *LOBFilterer) WatchMyMatchEvent(opts *bind.WatchOpts, sink chan<- *LOBMyMatchEvent) (event.Subscription, error) {

	logs, sub, err := _LOB.contract.WatchLogs(opts, "MyMatchEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LOBMyMatchEvent)
				if err := _LOB.contract.UnpackLog(event, "MyMatchEvent", log); err != nil {
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

// ParseMyMatchEvent is a log parse operation binding the contract event 0x4d85d09fd9972adb50adbf062a43da84abfebb9efd2340d5dd19e02d4f4d2683.
//
// Solidity: event MyMatchEvent(bytes16 arg0, uint256 arg1)
func (_LOB *LOBFilterer) ParseMyMatchEvent(log types.Log) (*LOBMyMatchEvent, error) {
	event := new(LOBMyMatchEvent)
	if err := _LOB.contract.UnpackLog(event, "MyMatchEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
