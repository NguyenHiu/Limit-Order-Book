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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MatchTimestamp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MyMatchEvent\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes16\",\"name\":\"orderID\",\"type\":\"bytes16\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"side\",\"type\":\"bool\"}],\"internalType\":\"structLOB.Order\",\"name\":\"newOrder\",\"type\":\"tuple\"}],\"name\":\"addOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b505f6001819055506116f3806100245f395ff3fe608060405234801561000f575f80fd5b5060043610610029575f3560e01c80639cfa8aa61461002d575b5f80fd5b610047600480360381019061004291906114b8565b610049565b005b6001805f82825461005a9190611510565b925050819055506001545f80835f01516fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff191681526020019081526020015f20819055505f60801b6fffffffffffffffffffffffffffffffff1916815f01516fffffffffffffffffffffffffffffffff191603610115577f4d85d09fd9972adb50adbf062a43da84abfebb9efd2340d5dd19e02d4f4d2683815f0151826040015160405161010c929190611561565b60405180910390a15b80606001511561012d5761012881610142565b610137565b610136816106fb565b5b61013f610cb4565b50565b5f600280549050036101de57600281908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f820151815f015f6101000a8154816fffffffffffffffffffffffffffffffff021916908360801c021790555060208201518160010155604082015181600201556060820151816003015f6101000a81548160ff02191690831515021790555050506106f7565b6001600280549050036103f557806020015160025f8154811061020457610203611588565b5b905f5260205f209060040201600101541015610364576002805f8154811061022f5761022e611588565b5b905f5260205f209060040201908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f82015f9054906101000a900460801b815f015f6101000a8154816fffffffffffffffffffffffffffffffff021916908360801c02179055506001820154816001015560028201548160020155600382015f9054906101000a900460ff16816003015f6101000a81548160ff02191690831515021790555050508060025f815481106102f0576102ef611588565b5b905f5260205f2090600402015f820151815f015f6101000a8154816fffffffffffffffffffffffffffffffff021916908360801c021790555060208201518160010155604082015181600201556060820151816003015f6101000a81548160ff0219169083151502179055509050506103f0565b600281908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f820151815f015f6101000a8154816fffffffffffffffffffffffffffffffff021916908360801c021790555060208201518160010155604082015181600201556060820151816003015f6101000a81548160ff02191690831515021790555050505b6106f6565b5f5b6002805490508160ff16101561066957816020015160028260ff168154811061042357610422611588565b5b905f5260205f20906004020160010154101561065657600282908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f820151815f015f6101000a8154816fffffffffffffffffffffffffffffffff021916908360801c021790555060208201518160010155604082015181600201556060820151816003015f6101000a81548160ff02191690831515021790555050505f60016002805490506104d791906115c1565b90505b8160ff168160ff1611156105c85760026001826104f791906115c1565b60ff168154811061050b5761050a611588565b5b905f5260205f20906004020160028260ff168154811061052e5761052d611588565b5b905f5260205f2090600402015f82015f9054906101000a900460801b815f015f6101000a8154816fffffffffffffffffffffffffffffffff021916908360801c02179055506001820154816001015560028201548160020155600382015f9054906101000a900460ff16816003015f6101000a81548160ff02191690831515021790555090505080806105c0906115f5565b9150506104da565b508160028260ff16815481106105e1576105e0611588565b5b905f5260205f2090600402015f820151815f015f6101000a8154816fffffffffffffffffffffffffffffffff021916908360801c021790555060208201518160010155604082015181600201556060820151816003015f6101000a81548160ff021916908315150217905550905050506106f8565b80806106619061161c565b9150506103f7565b50600281908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f820151815f015f6101000a8154816fffffffffffffffffffffffffffffffff021916908360801c021790555060208201518160010155604082015181600201556060820151816003015f6101000a81548160ff02191690831515021790555050505b5b5b50565b5f6003805490500361079757600381908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f820151815f015f6101000a8154816fffffffffffffffffffffffffffffffff021916908360801c021790555060208201518160010155604082015181600201556060820151816003015f6101000a81548160ff0219169083151502179055505050610cb0565b6001600380549050036109ae57806020015160035f815481106107bd576107bc611588565b5b905f5260205f20906004020160010154111561091d576003805f815481106107e8576107e7611588565b5b905f5260205f209060040201908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f82015f9054906101000a900460801b815f015f6101000a8154816fffffffffffffffffffffffffffffffff021916908360801c02179055506001820154816001015560028201548160020155600382015f9054906101000a900460ff16816003015f6101000a81548160ff02191690831515021790555050508060035f815481106108a9576108a8611588565b5b905f5260205f2090600402015f820151815f015f6101000a8154816fffffffffffffffffffffffffffffffff021916908360801c021790555060208201518160010155604082015181600201556060820151816003015f6101000a81548160ff0219169083151502179055509050506109a9565b600381908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f820151815f015f6101000a8154816fffffffffffffffffffffffffffffffff021916908360801c021790555060208201518160010155604082015181600201556060820151816003015f6101000a81548160ff02191690831515021790555050505b610caf565b5f5b6003805490508160ff161015610c2257816020015160038260ff16815481106109dc576109db611588565b5b905f5260205f209060040201600101541115610c0f57600382908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f820151815f015f6101000a8154816fffffffffffffffffffffffffffffffff021916908360801c021790555060208201518160010155604082015181600201556060820151816003015f6101000a81548160ff02191690831515021790555050505f6001600380549050610a9091906115c1565b90505b8160ff168160ff161115610b81576003600182610ab091906115c1565b60ff1681548110610ac457610ac3611588565b5b905f5260205f20906004020160038260ff1681548110610ae757610ae6611588565b5b905f5260205f2090600402015f82015f9054906101000a900460801b815f015f6101000a8154816fffffffffffffffffffffffffffffffff021916908360801c02179055506001820154816001015560028201548160020155600382015f9054906101000a900460ff16816003015f6101000a81548160ff0219169083151502179055509050508080610b79906115f5565b915050610a93565b508160028260ff1681548110610b9a57610b99611588565b5b905f5260205f2090600402015f820151815f015f6101000a8154816fffffffffffffffffffffffffffffffff021916908360801c021790555060208201518160010155604082015181600201556060820151816003015f6101000a81548160ff02191690831515021790555090505050610cb1565b8080610c1a9061161c565b9150506109b0565b50600281908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f820151815f015f6101000a8154816fffffffffffffffffffffffffffffffff021916908360801c021790555060208201518160010155604082015181600201556060820151816003015f6101000a81548160ff02191690831515021790555050505b5b5b50565b5b610cbd611278565b15611276575f60025f81548110610cd757610cd6611588565b5b905f5260205f20906004020160020154905060035f81548110610cfd57610cfc611588565b5b905f5260205f20906004020160020154811115610d3b5760035f81548110610d2857610d27611588565b5b905f5260205f2090600402016002015490505b7f4d85d09fd9972adb50adbf062a43da84abfebb9efd2340d5dd19e02d4f4d268360025f81548110610d7057610d6f611588565b5b905f5260205f2090600402015f015f9054906101000a900460801b82604051610d9a929190611561565b60405180910390a17fa0520e308970c9638737ab0138fbd19df951462a486a129799a46d3ac357ec6a5f8060025f81548110610dd957610dd8611588565b5b905f5260205f2090600402015f015f9054906101000a900460801b6fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff191681526020019081526020015f2054600154610e359190611644565b604051610e429190611677565b60405180910390a17fa0520e308970c9638737ab0138fbd19df951462a486a129799a46d3ac357ec6a5f8060035f81548110610e8157610e80611588565b5b905f5260205f2090600402015f015f9054906101000a900460801b6fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff191681526020019081526020015f2054600154610edd9190611644565b604051610eea9190611677565b60405180910390a18060025f81548110610f0757610f06611588565b5b905f5260205f2090600402016002015f828254610f249190611644565b925050819055508060035f81548110610f4057610f3f611588565b5b905f5260205f2090600402016002015f828254610f5d9190611644565b925050819055505f60025f81548110610f7957610f78611588565b5b905f5260205f20906004020160020154036110ea575f5b6001600280549050610fa29190611644565b81101561107d576002600182610fb89190611510565b81548110610fc957610fc8611588565b5b905f5260205f20906004020160028281548110610fe957610fe8611588565b5b905f5260205f2090600402015f82015f9054906101000a900460801b815f015f6101000a8154816fffffffffffffffffffffffffffffffff021916908360801c02179055506001820154816001015560028201548160020155600382015f9054906101000a900460ff16816003015f6101000a81548160ff0219169083151502179055509050508080600101915050610f90565b5060028054806110905761108f611690565b5b600190038181905f5260205f2090600402015f8082015f6101000a8154906fffffffffffffffffffffffffffffffff0219169055600182015f9055600282015f9055600382015f6101000a81549060ff0219169055505090555b5f60035f815481106110ff576110fe611588565b5b905f5260205f2090600402016002015403611270575f5b60016003805490506111289190611644565b81101561120357600360018261113e9190611510565b8154811061114f5761114e611588565b5b905f5260205f2090600402016003828154811061116f5761116e611588565b5b905f5260205f2090600402015f82015f9054906101000a900460801b815f015f6101000a8154816fffffffffffffffffffffffffffffffff021916908360801c02179055506001820154816001015560028201548160020155600382015f9054906101000a900460ff16816003015f6101000a81548160ff0219169083151502179055509050508080600101915050611116565b50600380548061121657611215611690565b5b600190038181905f5260205f2090600402015f8082015f6101000a8154906fffffffffffffffffffffffffffffffff0219169055600182015f9055600282015f9055600382015f6101000a81549060ff0219169055505090555b50610cb5565b565b5f80600280549050148061129057505f600380549050145b1561129d575f90506112ea565b60035f815481106112b1576112b0611588565b5b905f5260205f2090600402016001015460025f815481106112d5576112d4611588565b5b905f5260205f20906004020160010154101590505b90565b5f604051905090565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b611344826112fe565b810181811067ffffffffffffffff821117156113635761136261130e565b5b80604052505050565b5f6113756112ed565b9050611381828261133b565b919050565b5f7fffffffffffffffffffffffffffffffff0000000000000000000000000000000082169050919050565b6113ba81611386565b81146113c4575f80fd5b50565b5f813590506113d5816113b1565b92915050565b5f819050919050565b6113ed816113db565b81146113f7575f80fd5b50565b5f81359050611408816113e4565b92915050565b5f8115159050919050565b6114228161140e565b811461142c575f80fd5b50565b5f8135905061143d81611419565b92915050565b5f60808284031215611458576114576112fa565b5b611462608061136c565b90505f611471848285016113c7565b5f830152506020611484848285016113fa565b6020830152506040611498848285016113fa565b60408301525060606114ac8482850161142f565b60608301525092915050565b5f608082840312156114cd576114cc6112f6565b5b5f6114da84828501611443565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61151a826113db565b9150611525836113db565b925082820190508082111561153d5761153c6114e3565b5b92915050565b61154c81611386565b82525050565b61155b816113db565b82525050565b5f6040820190506115745f830185611543565b6115816020830184611552565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f60ff82169050919050565b5f6115cb826115b5565b91506115d6836115b5565b9250828203905060ff8111156115ef576115ee6114e3565b5b92915050565b5f6115ff826115b5565b91505f8203611611576116106114e3565b5b600182039050919050565b5f611626826115b5565b915060ff8203611639576116386114e3565b5b600182019050919050565b5f61164e826113db565b9150611659836113db565b9250828203905081811115611671576116706114e3565b5b92915050565b5f60208201905061168a5f830184611552565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffdfea26469706673582212203339196aff396f41a40a765215a1dfefddc7c340eabc1d27e04e817f5957921564736f6c63430008180033",
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

// LOBMatchTimestampIterator is returned from FilterMatchTimestamp and is used to iterate over the raw logs and unpacked data for MatchTimestamp events raised by the LOB contract.
type LOBMatchTimestampIterator struct {
	Event *LOBMatchTimestamp // Event containing the contract specifics and raw log

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
func (it *LOBMatchTimestampIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LOBMatchTimestamp)
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
		it.Event = new(LOBMatchTimestamp)
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
func (it *LOBMatchTimestampIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LOBMatchTimestampIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LOBMatchTimestamp represents a MatchTimestamp event raised by the LOB contract.
type LOBMatchTimestamp struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMatchTimestamp is a free log retrieval operation binding the contract event 0xa0520e308970c9638737ab0138fbd19df951462a486a129799a46d3ac357ec6a.
//
// Solidity: event MatchTimestamp(uint256 arg0)
func (_LOB *LOBFilterer) FilterMatchTimestamp(opts *bind.FilterOpts) (*LOBMatchTimestampIterator, error) {

	logs, sub, err := _LOB.contract.FilterLogs(opts, "MatchTimestamp")
	if err != nil {
		return nil, err
	}
	return &LOBMatchTimestampIterator{contract: _LOB.contract, event: "MatchTimestamp", logs: logs, sub: sub}, nil
}

// WatchMatchTimestamp is a free log subscription operation binding the contract event 0xa0520e308970c9638737ab0138fbd19df951462a486a129799a46d3ac357ec6a.
//
// Solidity: event MatchTimestamp(uint256 arg0)
func (_LOB *LOBFilterer) WatchMatchTimestamp(opts *bind.WatchOpts, sink chan<- *LOBMatchTimestamp) (event.Subscription, error) {

	logs, sub, err := _LOB.contract.WatchLogs(opts, "MatchTimestamp")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LOBMatchTimestamp)
				if err := _LOB.contract.UnpackLog(event, "MatchTimestamp", log); err != nil {
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

// ParseMatchTimestamp is a log parse operation binding the contract event 0xa0520e308970c9638737ab0138fbd19df951462a486a129799a46d3ac357ec6a.
//
// Solidity: event MatchTimestamp(uint256 arg0)
func (_LOB *LOBFilterer) ParseMatchTimestamp(log types.Log) (*LOBMatchTimestamp, error) {
	event := new(LOBMatchTimestamp)
	if err := _LOB.contract.UnpackLog(event, "MatchTimestamp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
