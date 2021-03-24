// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// EthereumDIDRegistryABI is the input ABI used to generate the binding from.
const EthereumDIDRegistryABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"bytes32\"},{\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"revokeAttribute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"owners\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"delegates\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"sigV\",\"type\":\"uint8\"},{\"name\":\"sigR\",\"type\":\"bytes32\"},{\"name\":\"sigS\",\"type\":\"bytes32\"},{\"name\":\"name\",\"type\":\"bytes32\"},{\"name\":\"value\",\"type\":\"bytes\"},{\"name\":\"validity\",\"type\":\"uint256\"}],\"name\":\"setAttributeSigned\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"sigV\",\"type\":\"uint8\"},{\"name\":\"sigR\",\"type\":\"bytes32\"},{\"name\":\"sigS\",\"type\":\"bytes32\"},{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwnerSigned\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"delegateType\",\"type\":\"bytes32\"},{\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"validDelegate\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonce\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"bytes32\"},{\"name\":\"value\",\"type\":\"bytes\"},{\"name\":\"validity\",\"type\":\"uint256\"}],\"name\":\"setAttribute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"delegateType\",\"type\":\"bytes32\"},{\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"revokeDelegate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"}],\"name\":\"identityOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"sigV\",\"type\":\"uint8\"},{\"name\":\"sigR\",\"type\":\"bytes32\"},{\"name\":\"sigS\",\"type\":\"bytes32\"},{\"name\":\"delegateType\",\"type\":\"bytes32\"},{\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"revokeDelegateSigned\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"sigV\",\"type\":\"uint8\"},{\"name\":\"sigR\",\"type\":\"bytes32\"},{\"name\":\"sigS\",\"type\":\"bytes32\"},{\"name\":\"delegateType\",\"type\":\"bytes32\"},{\"name\":\"delegate\",\"type\":\"address\"},{\"name\":\"validity\",\"type\":\"uint256\"}],\"name\":\"addDelegateSigned\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"delegateType\",\"type\":\"bytes32\"},{\"name\":\"delegate\",\"type\":\"address\"},{\"name\":\"validity\",\"type\":\"uint256\"}],\"name\":\"addDelegate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"sigV\",\"type\":\"uint8\"},{\"name\":\"sigR\",\"type\":\"bytes32\"},{\"name\":\"sigS\",\"type\":\"bytes32\"},{\"name\":\"name\",\"type\":\"bytes32\"},{\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"revokeAttributeSigned\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"changed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"previousChange\",\"type\":\"uint256\"}],\"name\":\"DIDOwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"delegateType\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"validTo\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"previousChange\",\"type\":\"uint256\"}],\"name\":\"DIDDelegateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"validTo\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"previousChange\",\"type\":\"uint256\"}],\"name\":\"DIDAttributeChanged\",\"type\":\"event\"}]"

// EthereumDIDRegistryFuncSigs maps the 4-byte function signature to its string representation.
var EthereumDIDRegistryFuncSigs = map[string]string{
	"a7068d66": "addDelegate(address,bytes32,address,uint256)",
	"9c2c1b2b": "addDelegateSigned(address,uint8,bytes32,bytes32,bytes32,address,uint256)",
	"f00d4b5d": "changeOwner(address,address)",
	"240cf1fa": "changeOwnerSigned(address,uint8,bytes32,bytes32,address)",
	"f96d0f9f": "changed(address)",
	"0d44625b": "delegates(address,bytes32,address)",
	"8733d4e8": "identityOwner(address)",
	"70ae92d2": "nonce(address)",
	"022914a7": "owners(address)",
	"00c023da": "revokeAttribute(address,bytes32,bytes)",
	"e476af5c": "revokeAttributeSigned(address,uint8,bytes32,bytes32,bytes32,bytes)",
	"80b29f7c": "revokeDelegate(address,bytes32,address)",
	"93072684": "revokeDelegateSigned(address,uint8,bytes32,bytes32,bytes32,address)",
	"7ad4b0a4": "setAttribute(address,bytes32,bytes,uint256)",
	"123b5e98": "setAttributeSigned(address,uint8,bytes32,bytes32,bytes32,bytes,uint256)",
	"622b2a3c": "validDelegate(address,bytes32,address)",
}

// EthereumDIDRegistryBin is the compiled bytecode used for deploying new contracts.
var EthereumDIDRegistryBin = "0x608060405234801561001057600080fd5b5061108c806100206000396000f3006080604052600436106100e45763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041662c023da81146100e9578063022914a7146101545780630d44625b14610191578063123b5e98146101ce578063240cf1fa14610248578063622b2a3c1461027e57806370ae92d2146102bd5780637ad4b0a4146102de57806380b29f7c146103495780638733d4e81461037457806393072684146103955780639c2c1b2b146103cf578063a7068d661461040c578063e476af5c1461043a578063f00d4b5d146104b2578063f96d0f9f146104d9575b600080fd5b3480156100f557600080fd5b50604080516020600460443581810135601f8101849004840285018401909552848452610152948235600160a060020a03169460248035953695946064949201919081908401838280828437509497506104fa9650505050505050565b005b34801561016057600080fd5b50610175600160a060020a036004351661050b565b60408051600160a060020a039092168252519081900360200190f35b34801561019d57600080fd5b506101bc600160a060020a036004358116906024359060443516610526565b60408051918252519081900360200190f35b3480156101da57600080fd5b50604080516020600460a43581810135601f8101849004840285018401909552848452610152948235600160a060020a0316946024803560ff16956044359560643595608435953695929460c494920191819084018382808284375094975050933594506105499350505050565b34801561025457600080fd5b50610152600160a060020a0360043581169060ff6024351690604435906064359060843516610684565b34801561028a57600080fd5b506102a9600160a060020a036004358116906024359060443516610757565b604080519115158252519081900360200190f35b3480156102c957600080fd5b506101bc600160a060020a036004351661079b565b3480156102ea57600080fd5b50604080516020600460443581810135601f8101849004840285018401909552848452610152948235600160a060020a031694602480359536959460649492019190819084018382808284375094975050933594506107ad9350505050565b34801561035557600080fd5b50610152600160a060020a0360043581169060243590604435166107c0565b34801561038057600080fd5b50610175600160a060020a03600435166107cc565b3480156103a157600080fd5b50610152600160a060020a0360043581169060ff602435169060443590606435906084359060a43516610801565b3480156103db57600080fd5b50610152600160a060020a0360043581169060ff602435169060443590606435906084359060a4351660c4356108dd565b34801561041857600080fd5b50610152600160a060020a0360043581169060243590604435166064356109b8565b34801561044657600080fd5b50604080516020600460a43581810135601f8101849004840285018401909552848452610152948235600160a060020a0316946024803560ff16956044359560643595608435953695929460c49492019181908401838280828437509497506109c59650505050505050565b3480156104be57600080fd5b50610152600160a060020a0360043581169060243516610aee565b3480156104e557600080fd5b506101bc600160a060020a0360043516610afd565b61050683338484610b0f565b505050565b600060208190529081526040902054600160a060020a031681565b600160209081526000938452604080852082529284528284209052825290205481565b600060f860020a60190281306003826105618d6107cc565b600160a060020a0390811682526020808301939093526040918201600020549151600160f860020a03198088168252861660018201526c010000000000000000000000008583168102600283015260168201849052918f1690910260368201527f7365744174747269627574650000000000000000000000000000000000000000604a820152605681018a9052885191928e928b928b928b9260768301918501908083835b602083106106255780518252601f199092019160209182019101610606565b6001836020036101000a038019825116818451168082178552505050505050905001828152602001985050505050505050506040518091039020905061067a886106728a8a8a8a87610c22565b868686610cd4565b5050505050505050565b600060f860020a601902813060038261069c8b6107cc565b600160a060020a03908116825260208201929092526040908101600020548151600160f860020a031996871681529490951660018501526c01000000000000000000000000928216830260028501526016840194909452898116820260368401527f6368616e67654f776e6572000000000000000000000000000000000000000000604a84015285160260558201529051908190036069019020905061074f866107498188888887610c22565b84610ded565b505050505050565b600160a060020a0392831660009081526001602090815260408083208151958652815195869003830190952083529381528382209290941681529252902054421090565b60036020526000908152604090205481565b6107ba8433858585610cd4565b50505050565b61050683338484610eab565b600160a060020a0380821660009081526020819052604081205490911680156107f7578091506107fb565b8291505b50919050565b600060f860020a60190281306003826108198c6107cc565b600160a060020a03908116825260208201929092526040908101600020548151600160f860020a031996871681529490951660018501526c010000000000000000000000009282168302600285015260168401949094528a8116820260368401527f7265766f6b6544656c6567617465000000000000000000000000000000000000604a840152605883018790528516026078820152905190819003608c01902090506108d4876108cd8189898987610c22565b8585610eab565b50505050505050565b600060f860020a60190281306003826108f58d6107cc565b600160a060020a03908116825260208201929092526040908101600020548151600160f860020a031996871681529490951660018501526c010000000000000000000000009282168302600285015260168401949094528b8116820260368401527f61646444656c6567617465000000000000000000000000000000000000000000604a8401526055830188905286160260758201526089810184905290519081900360a9019020905061067a886109b0818a8a8a87610c22565b868686610f84565b6107ba8433858585610f84565b600060f860020a60190281306003826109dd8c6107cc565b600160a060020a0390811682526020808301939093526040918201600020549151600160f860020a03198088168252861660018201526c010000000000000000000000008583168102600283015260168201849052918e1690910260368201527f7265766f6b654174747269627574650000000000000000000000000000000000604a82015260598101899052875191928d928a928a92909160798301918401908083835b60208310610aa15780518252601f199092019160209182019101610a82565b6001836020036101000a038019825116818451168082178552505050505050905001975050505050505050604051809103902090506108d487610ae78989898987610c22565b8585610b0f565b610af9823383610ded565b5050565b60026020526000908152604090205481565b8383610b1a826107cc565b600160a060020a03828116911614610b3157600080fd5b600160a060020a038616600081815260026020908152604080832054815189815291820184905260608201819052608082840181815289519184019190915288517f18ab6b2ae3d64306c00ce663125f2bd680e441a098de1635bd7ad8b0d44965e4958b958b95919493919260a0840191870190808383895b83811015610bc2578181015183820152602001610baa565b50505050905090810190601f168015610bef5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a2505050600160a060020a0390921660009081526002602052604090204390555050565b604080516000808252602080830180855285905260ff881683850152606083018790526080830186905292519092839260019260a08083019392601f19830192908190039091019086865af1158015610c7f573d6000803e3d6000fd5b505050602060405103519050610c94876107cc565b600160a060020a03828116911614610cab57600080fd5b600160a060020a0381166000908152600360205260409020805460010190559695505050505050565b8484610cdf826107cc565b600160a060020a03828116911614610cf657600080fd5b600160a060020a03871660008181526002602090815260408083205481518a81524289019281018390526060810182905260808185018181528b51918301919091528a517f18ab6b2ae3d64306c00ce663125f2bd680e441a098de1635bd7ad8b0d44965e4968d968d969594939260a0850192918801918190849084905b83811015610d8c578181015183820152602001610d74565b50505050905090810190601f168015610db95780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a2505050600160a060020a039093166000908152600260205260409020439055505050565b8282610df8826107cc565b600160a060020a03828116911614610e0f57600080fd5b600160a060020a03858116600081815260208181526040808320805473ffffffffffffffffffffffffffffffffffffffff19169589169586179055600282529182902054825194855290840152805191927f38a5a6e68f30ed1ab45860a4afb34bcb2fc00f22ca462d249b8a8d40cda6f7a3929081900390910190a250505050600160a060020a03166000908152600260205260409020439055565b8383610eb6826107cc565b600160a060020a03828116911614610ecd57600080fd5b600160a060020a03808716600081815260016020908152604080832081518a815282519081900384018120855290835281842095891680855295835281842042908190558585526002845293829020548a8252928101959095528481019290925260608401525190917f5a5084339536bcab65f20799fcc58724588145ca054bd2be626174b27ba156f7919081900360800190a2505050600160a060020a0390921660009081526002602052604090204390555050565b8484610f8f826107cc565b600160a060020a03828116911614610fa657600080fd5b600160a060020a03808816600081815260016020908152604080832081518b8152825190819003840181208552908352818420958a16808552958352818420428a01908190558585526002845293829020548b8252928101959095528481019290925260608401525190917f5a5084339536bcab65f20799fcc58724588145ca054bd2be626174b27ba156f7919081900360800190a2505050600160a060020a0390931660009081526002602052604090204390555050505600a165627a7a72305820effa580be1b266b6306fe295c8e9a7365ba2c87d2b1feb057dd32d1d281b69b30029"

// DeployEthereumDIDRegistry deploys a new Ethereum contract, binding an instance of EthereumDIDRegistry to it.
func DeployEthereumDIDRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EthereumDIDRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(EthereumDIDRegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EthereumDIDRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthereumDIDRegistry{EthereumDIDRegistryCaller: EthereumDIDRegistryCaller{contract: contract}, EthereumDIDRegistryTransactor: EthereumDIDRegistryTransactor{contract: contract}, EthereumDIDRegistryFilterer: EthereumDIDRegistryFilterer{contract: contract}}, nil
}

// EthereumDIDRegistry is an auto generated Go binding around an Ethereum contract.
type EthereumDIDRegistry struct {
	EthereumDIDRegistryCaller     // Read-only binding to the contract
	EthereumDIDRegistryTransactor // Write-only binding to the contract
	EthereumDIDRegistryFilterer   // Log filterer for contract events
}

// EthereumDIDRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthereumDIDRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthereumDIDRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthereumDIDRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthereumDIDRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthereumDIDRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthereumDIDRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthereumDIDRegistrySession struct {
	Contract     *EthereumDIDRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// EthereumDIDRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthereumDIDRegistryCallerSession struct {
	Contract *EthereumDIDRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// EthereumDIDRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthereumDIDRegistryTransactorSession struct {
	Contract     *EthereumDIDRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// EthereumDIDRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthereumDIDRegistryRaw struct {
	Contract *EthereumDIDRegistry // Generic contract binding to access the raw methods on
}

// EthereumDIDRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthereumDIDRegistryCallerRaw struct {
	Contract *EthereumDIDRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// EthereumDIDRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthereumDIDRegistryTransactorRaw struct {
	Contract *EthereumDIDRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthereumDIDRegistry creates a new instance of EthereumDIDRegistry, bound to a specific deployed contract.
func NewEthereumDIDRegistry(address common.Address, backend bind.ContractBackend) (*EthereumDIDRegistry, error) {
	contract, err := bindEthereumDIDRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthereumDIDRegistry{EthereumDIDRegistryCaller: EthereumDIDRegistryCaller{contract: contract}, EthereumDIDRegistryTransactor: EthereumDIDRegistryTransactor{contract: contract}, EthereumDIDRegistryFilterer: EthereumDIDRegistryFilterer{contract: contract}}, nil
}

// NewEthereumDIDRegistryCaller creates a new read-only instance of EthereumDIDRegistry, bound to a specific deployed contract.
func NewEthereumDIDRegistryCaller(address common.Address, caller bind.ContractCaller) (*EthereumDIDRegistryCaller, error) {
	contract, err := bindEthereumDIDRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthereumDIDRegistryCaller{contract: contract}, nil
}

// NewEthereumDIDRegistryTransactor creates a new write-only instance of EthereumDIDRegistry, bound to a specific deployed contract.
func NewEthereumDIDRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*EthereumDIDRegistryTransactor, error) {
	contract, err := bindEthereumDIDRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthereumDIDRegistryTransactor{contract: contract}, nil
}

// NewEthereumDIDRegistryFilterer creates a new log filterer instance of EthereumDIDRegistry, bound to a specific deployed contract.
func NewEthereumDIDRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*EthereumDIDRegistryFilterer, error) {
	contract, err := bindEthereumDIDRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthereumDIDRegistryFilterer{contract: contract}, nil
}

// bindEthereumDIDRegistry binds a generic wrapper to an already deployed contract.
func bindEthereumDIDRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthereumDIDRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthereumDIDRegistry *EthereumDIDRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthereumDIDRegistry.Contract.EthereumDIDRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthereumDIDRegistry *EthereumDIDRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthereumDIDRegistry.Contract.EthereumDIDRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthereumDIDRegistry *EthereumDIDRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthereumDIDRegistry.Contract.EthereumDIDRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthereumDIDRegistry *EthereumDIDRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthereumDIDRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthereumDIDRegistry *EthereumDIDRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthereumDIDRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthereumDIDRegistry *EthereumDIDRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthereumDIDRegistry.Contract.contract.Transact(opts, method, params...)
}

// Changed is a free data retrieval call binding the contract method 0xf96d0f9f.
//
// Solidity: function changed(address ) view returns(uint256)
func (_EthereumDIDRegistry *EthereumDIDRegistryCaller) Changed(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EthereumDIDRegistry.contract.Call(opts, &out, "changed", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Changed is a free data retrieval call binding the contract method 0xf96d0f9f.
//
// Solidity: function changed(address ) view returns(uint256)
func (_EthereumDIDRegistry *EthereumDIDRegistrySession) Changed(arg0 common.Address) (*big.Int, error) {
	return _EthereumDIDRegistry.Contract.Changed(&_EthereumDIDRegistry.CallOpts, arg0)
}

// Changed is a free data retrieval call binding the contract method 0xf96d0f9f.
//
// Solidity: function changed(address ) view returns(uint256)
func (_EthereumDIDRegistry *EthereumDIDRegistryCallerSession) Changed(arg0 common.Address) (*big.Int, error) {
	return _EthereumDIDRegistry.Contract.Changed(&_EthereumDIDRegistry.CallOpts, arg0)
}

// Delegates is a free data retrieval call binding the contract method 0x0d44625b.
//
// Solidity: function delegates(address , bytes32 , address ) view returns(uint256)
func (_EthereumDIDRegistry *EthereumDIDRegistryCaller) Delegates(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte, arg2 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EthereumDIDRegistry.contract.Call(opts, &out, "delegates", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Delegates is a free data retrieval call binding the contract method 0x0d44625b.
//
// Solidity: function delegates(address , bytes32 , address ) view returns(uint256)
func (_EthereumDIDRegistry *EthereumDIDRegistrySession) Delegates(arg0 common.Address, arg1 [32]byte, arg2 common.Address) (*big.Int, error) {
	return _EthereumDIDRegistry.Contract.Delegates(&_EthereumDIDRegistry.CallOpts, arg0, arg1, arg2)
}

// Delegates is a free data retrieval call binding the contract method 0x0d44625b.
//
// Solidity: function delegates(address , bytes32 , address ) view returns(uint256)
func (_EthereumDIDRegistry *EthereumDIDRegistryCallerSession) Delegates(arg0 common.Address, arg1 [32]byte, arg2 common.Address) (*big.Int, error) {
	return _EthereumDIDRegistry.Contract.Delegates(&_EthereumDIDRegistry.CallOpts, arg0, arg1, arg2)
}

// IdentityOwner is a free data retrieval call binding the contract method 0x8733d4e8.
//
// Solidity: function identityOwner(address identity) view returns(address)
func (_EthereumDIDRegistry *EthereumDIDRegistryCaller) IdentityOwner(opts *bind.CallOpts, identity common.Address) (common.Address, error) {
	var out []interface{}
	err := _EthereumDIDRegistry.contract.Call(opts, &out, "identityOwner", identity)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IdentityOwner is a free data retrieval call binding the contract method 0x8733d4e8.
//
// Solidity: function identityOwner(address identity) view returns(address)
func (_EthereumDIDRegistry *EthereumDIDRegistrySession) IdentityOwner(identity common.Address) (common.Address, error) {
	return _EthereumDIDRegistry.Contract.IdentityOwner(&_EthereumDIDRegistry.CallOpts, identity)
}

// IdentityOwner is a free data retrieval call binding the contract method 0x8733d4e8.
//
// Solidity: function identityOwner(address identity) view returns(address)
func (_EthereumDIDRegistry *EthereumDIDRegistryCallerSession) IdentityOwner(identity common.Address) (common.Address, error) {
	return _EthereumDIDRegistry.Contract.IdentityOwner(&_EthereumDIDRegistry.CallOpts, identity)
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) view returns(uint256)
func (_EthereumDIDRegistry *EthereumDIDRegistryCaller) Nonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EthereumDIDRegistry.contract.Call(opts, &out, "nonce", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) view returns(uint256)
func (_EthereumDIDRegistry *EthereumDIDRegistrySession) Nonce(arg0 common.Address) (*big.Int, error) {
	return _EthereumDIDRegistry.Contract.Nonce(&_EthereumDIDRegistry.CallOpts, arg0)
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) view returns(uint256)
func (_EthereumDIDRegistry *EthereumDIDRegistryCallerSession) Nonce(arg0 common.Address) (*big.Int, error) {
	return _EthereumDIDRegistry.Contract.Nonce(&_EthereumDIDRegistry.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(address)
func (_EthereumDIDRegistry *EthereumDIDRegistryCaller) Owners(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _EthereumDIDRegistry.contract.Call(opts, &out, "owners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(address)
func (_EthereumDIDRegistry *EthereumDIDRegistrySession) Owners(arg0 common.Address) (common.Address, error) {
	return _EthereumDIDRegistry.Contract.Owners(&_EthereumDIDRegistry.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(address)
func (_EthereumDIDRegistry *EthereumDIDRegistryCallerSession) Owners(arg0 common.Address) (common.Address, error) {
	return _EthereumDIDRegistry.Contract.Owners(&_EthereumDIDRegistry.CallOpts, arg0)
}

// ValidDelegate is a free data retrieval call binding the contract method 0x622b2a3c.
//
// Solidity: function validDelegate(address identity, bytes32 delegateType, address delegate) view returns(bool)
func (_EthereumDIDRegistry *EthereumDIDRegistryCaller) ValidDelegate(opts *bind.CallOpts, identity common.Address, delegateType [32]byte, delegate common.Address) (bool, error) {
	var out []interface{}
	err := _EthereumDIDRegistry.contract.Call(opts, &out, "validDelegate", identity, delegateType, delegate)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidDelegate is a free data retrieval call binding the contract method 0x622b2a3c.
//
// Solidity: function validDelegate(address identity, bytes32 delegateType, address delegate) view returns(bool)
func (_EthereumDIDRegistry *EthereumDIDRegistrySession) ValidDelegate(identity common.Address, delegateType [32]byte, delegate common.Address) (bool, error) {
	return _EthereumDIDRegistry.Contract.ValidDelegate(&_EthereumDIDRegistry.CallOpts, identity, delegateType, delegate)
}

// ValidDelegate is a free data retrieval call binding the contract method 0x622b2a3c.
//
// Solidity: function validDelegate(address identity, bytes32 delegateType, address delegate) view returns(bool)
func (_EthereumDIDRegistry *EthereumDIDRegistryCallerSession) ValidDelegate(identity common.Address, delegateType [32]byte, delegate common.Address) (bool, error) {
	return _EthereumDIDRegistry.Contract.ValidDelegate(&_EthereumDIDRegistry.CallOpts, identity, delegateType, delegate)
}

// AddDelegate is a paid mutator transaction binding the contract method 0xa7068d66.
//
// Solidity: function addDelegate(address identity, bytes32 delegateType, address delegate, uint256 validity) returns()
func (_EthereumDIDRegistry *EthereumDIDRegistryTransactor) AddDelegate(opts *bind.TransactOpts, identity common.Address, delegateType [32]byte, delegate common.Address, validity *big.Int) (*types.Transaction, error) {
	return _EthereumDIDRegistry.contract.Transact(opts, "addDelegate", identity, delegateType, delegate, validity)
}

// AddDelegate is a paid mutator transaction binding the contract method 0xa7068d66.
//
// Solidity: function addDelegate(address identity, bytes32 delegateType, address delegate, uint256 validity) returns()
func (_EthereumDIDRegistry *EthereumDIDRegistrySession) AddDelegate(identity common.Address, delegateType [32]byte, delegate common.Address, validity *big.Int) (*types.Transaction, error) {
	return _EthereumDIDRegistry.Contract.AddDelegate(&_EthereumDIDRegistry.TransactOpts, identity, delegateType, delegate, validity)
}

// AddDelegate is a paid mutator transaction binding the contract method 0xa7068d66.
//
// Solidity: function addDelegate(address identity, bytes32 delegateType, address delegate, uint256 validity) returns()
func (_EthereumDIDRegistry *EthereumDIDRegistryTransactorSession) AddDelegate(identity common.Address, delegateType [32]byte, delegate common.Address, validity *big.Int) (*types.Transaction, error) {
	return _EthereumDIDRegistry.Contract.AddDelegate(&_EthereumDIDRegistry.TransactOpts, identity, delegateType, delegate, validity)
}

// AddDelegateSigned is a paid mutator transaction binding the contract method 0x9c2c1b2b.
//
// Solidity: function addDelegateSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 delegateType, address delegate, uint256 validity) returns()
func (_EthereumDIDRegistry *EthereumDIDRegistryTransactor) AddDelegateSigned(opts *bind.TransactOpts, identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, delegateType [32]byte, delegate common.Address, validity *big.Int) (*types.Transaction, error) {
	return _EthereumDIDRegistry.contract.Transact(opts, "addDelegateSigned", identity, sigV, sigR, sigS, delegateType, delegate, validity)
}

// AddDelegateSigned is a paid mutator transaction binding the contract method 0x9c2c1b2b.
//
// Solidity: function addDelegateSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 delegateType, address delegate, uint256 validity) returns()
func (_EthereumDIDRegistry *EthereumDIDRegistrySession) AddDelegateSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, delegateType [32]byte, delegate common.Address, validity *big.Int) (*types.Transaction, error) {
	return _EthereumDIDRegistry.Contract.AddDelegateSigned(&_EthereumDIDRegistry.TransactOpts, identity, sigV, sigR, sigS, delegateType, delegate, validity)
}

// AddDelegateSigned is a paid mutator transaction binding the contract method 0x9c2c1b2b.
//
// Solidity: function addDelegateSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 delegateType, address delegate, uint256 validity) returns()
func (_EthereumDIDRegistry *EthereumDIDRegistryTransactorSession) AddDelegateSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, delegateType [32]byte, delegate common.Address, validity *big.Int) (*types.Transaction, error) {
	return _EthereumDIDRegistry.Contract.AddDelegateSigned(&_EthereumDIDRegistry.TransactOpts, identity, sigV, sigR, sigS, delegateType, delegate, validity)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xf00d4b5d.
//
// Solidity: function changeOwner(address identity, address newOwner) returns()
func (_EthereumDIDRegistry *EthereumDIDRegistryTransactor) ChangeOwner(opts *bind.TransactOpts, identity common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _EthereumDIDRegistry.contract.Transact(opts, "changeOwner", identity, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xf00d4b5d.
//
// Solidity: function changeOwner(address identity, address newOwner) returns()
func (_EthereumDIDRegistry *EthereumDIDRegistrySession) ChangeOwner(identity common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _EthereumDIDRegistry.Contract.ChangeOwner(&_EthereumDIDRegistry.TransactOpts, identity, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xf00d4b5d.
//
// Solidity: function changeOwner(address identity, address newOwner) returns()
func (_EthereumDIDRegistry *EthereumDIDRegistryTransactorSession) ChangeOwner(identity common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _EthereumDIDRegistry.Contract.ChangeOwner(&_EthereumDIDRegistry.TransactOpts, identity, newOwner)
}

// ChangeOwnerSigned is a paid mutator transaction binding the contract method 0x240cf1fa.
//
// Solidity: function changeOwnerSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, address newOwner) returns()
func (_EthereumDIDRegistry *EthereumDIDRegistryTransactor) ChangeOwnerSigned(opts *bind.TransactOpts, identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _EthereumDIDRegistry.contract.Transact(opts, "changeOwnerSigned", identity, sigV, sigR, sigS, newOwner)
}

// ChangeOwnerSigned is a paid mutator transaction binding the contract method 0x240cf1fa.
//
// Solidity: function changeOwnerSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, address newOwner) returns()
func (_EthereumDIDRegistry *EthereumDIDRegistrySession) ChangeOwnerSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _EthereumDIDRegistry.Contract.ChangeOwnerSigned(&_EthereumDIDRegistry.TransactOpts, identity, sigV, sigR, sigS, newOwner)
}

// ChangeOwnerSigned is a paid mutator transaction binding the contract method 0x240cf1fa.
//
// Solidity: function changeOwnerSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, address newOwner) returns()
func (_EthereumDIDRegistry *EthereumDIDRegistryTransactorSession) ChangeOwnerSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _EthereumDIDRegistry.Contract.ChangeOwnerSigned(&_EthereumDIDRegistry.TransactOpts, identity, sigV, sigR, sigS, newOwner)
}

// RevokeAttribute is a paid mutator transaction binding the contract method 0x00c023da.
//
// Solidity: function revokeAttribute(address identity, bytes32 name, bytes value) returns()
func (_EthereumDIDRegistry *EthereumDIDRegistryTransactor) RevokeAttribute(opts *bind.TransactOpts, identity common.Address, name [32]byte, value []byte) (*types.Transaction, error) {
	return _EthereumDIDRegistry.contract.Transact(opts, "revokeAttribute", identity, name, value)
}

// RevokeAttribute is a paid mutator transaction binding the contract method 0x00c023da.
//
// Solidity: function revokeAttribute(address identity, bytes32 name, bytes value) returns()
func (_EthereumDIDRegistry *EthereumDIDRegistrySession) RevokeAttribute(identity common.Address, name [32]byte, value []byte) (*types.Transaction, error) {
	return _EthereumDIDRegistry.Contract.RevokeAttribute(&_EthereumDIDRegistry.TransactOpts, identity, name, value)
}

// RevokeAttribute is a paid mutator transaction binding the contract method 0x00c023da.
//
// Solidity: function revokeAttribute(address identity, bytes32 name, bytes value) returns()
func (_EthereumDIDRegistry *EthereumDIDRegistryTransactorSession) RevokeAttribute(identity common.Address, name [32]byte, value []byte) (*types.Transaction, error) {
	return _EthereumDIDRegistry.Contract.RevokeAttribute(&_EthereumDIDRegistry.TransactOpts, identity, name, value)
}

// RevokeAttributeSigned is a paid mutator transaction binding the contract method 0xe476af5c.
//
// Solidity: function revokeAttributeSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 name, bytes value) returns()
func (_EthereumDIDRegistry *EthereumDIDRegistryTransactor) RevokeAttributeSigned(opts *bind.TransactOpts, identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, name [32]byte, value []byte) (*types.Transaction, error) {
	return _EthereumDIDRegistry.contract.Transact(opts, "revokeAttributeSigned", identity, sigV, sigR, sigS, name, value)
}

// RevokeAttributeSigned is a paid mutator transaction binding the contract method 0xe476af5c.
//
// Solidity: function revokeAttributeSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 name, bytes value) returns()
func (_EthereumDIDRegistry *EthereumDIDRegistrySession) RevokeAttributeSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, name [32]byte, value []byte) (*types.Transaction, error) {
	return _EthereumDIDRegistry.Contract.RevokeAttributeSigned(&_EthereumDIDRegistry.TransactOpts, identity, sigV, sigR, sigS, name, value)
}

// RevokeAttributeSigned is a paid mutator transaction binding the contract method 0xe476af5c.
//
// Solidity: function revokeAttributeSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 name, bytes value) returns()
func (_EthereumDIDRegistry *EthereumDIDRegistryTransactorSession) RevokeAttributeSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, name [32]byte, value []byte) (*types.Transaction, error) {
	return _EthereumDIDRegistry.Contract.RevokeAttributeSigned(&_EthereumDIDRegistry.TransactOpts, identity, sigV, sigR, sigS, name, value)
}

// RevokeDelegate is a paid mutator transaction binding the contract method 0x80b29f7c.
//
// Solidity: function revokeDelegate(address identity, bytes32 delegateType, address delegate) returns()
func (_EthereumDIDRegistry *EthereumDIDRegistryTransactor) RevokeDelegate(opts *bind.TransactOpts, identity common.Address, delegateType [32]byte, delegate common.Address) (*types.Transaction, error) {
	return _EthereumDIDRegistry.contract.Transact(opts, "revokeDelegate", identity, delegateType, delegate)
}

// RevokeDelegate is a paid mutator transaction binding the contract method 0x80b29f7c.
//
// Solidity: function revokeDelegate(address identity, bytes32 delegateType, address delegate) returns()
func (_EthereumDIDRegistry *EthereumDIDRegistrySession) RevokeDelegate(identity common.Address, delegateType [32]byte, delegate common.Address) (*types.Transaction, error) {
	return _EthereumDIDRegistry.Contract.RevokeDelegate(&_EthereumDIDRegistry.TransactOpts, identity, delegateType, delegate)
}

// RevokeDelegate is a paid mutator transaction binding the contract method 0x80b29f7c.
//
// Solidity: function revokeDelegate(address identity, bytes32 delegateType, address delegate) returns()
func (_EthereumDIDRegistry *EthereumDIDRegistryTransactorSession) RevokeDelegate(identity common.Address, delegateType [32]byte, delegate common.Address) (*types.Transaction, error) {
	return _EthereumDIDRegistry.Contract.RevokeDelegate(&_EthereumDIDRegistry.TransactOpts, identity, delegateType, delegate)
}

// RevokeDelegateSigned is a paid mutator transaction binding the contract method 0x93072684.
//
// Solidity: function revokeDelegateSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 delegateType, address delegate) returns()
func (_EthereumDIDRegistry *EthereumDIDRegistryTransactor) RevokeDelegateSigned(opts *bind.TransactOpts, identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, delegateType [32]byte, delegate common.Address) (*types.Transaction, error) {
	return _EthereumDIDRegistry.contract.Transact(opts, "revokeDelegateSigned", identity, sigV, sigR, sigS, delegateType, delegate)
}

// RevokeDelegateSigned is a paid mutator transaction binding the contract method 0x93072684.
//
// Solidity: function revokeDelegateSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 delegateType, address delegate) returns()
func (_EthereumDIDRegistry *EthereumDIDRegistrySession) RevokeDelegateSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, delegateType [32]byte, delegate common.Address) (*types.Transaction, error) {
	return _EthereumDIDRegistry.Contract.RevokeDelegateSigned(&_EthereumDIDRegistry.TransactOpts, identity, sigV, sigR, sigS, delegateType, delegate)
}

// RevokeDelegateSigned is a paid mutator transaction binding the contract method 0x93072684.
//
// Solidity: function revokeDelegateSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 delegateType, address delegate) returns()
func (_EthereumDIDRegistry *EthereumDIDRegistryTransactorSession) RevokeDelegateSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, delegateType [32]byte, delegate common.Address) (*types.Transaction, error) {
	return _EthereumDIDRegistry.Contract.RevokeDelegateSigned(&_EthereumDIDRegistry.TransactOpts, identity, sigV, sigR, sigS, delegateType, delegate)
}

// SetAttribute is a paid mutator transaction binding the contract method 0x7ad4b0a4.
//
// Solidity: function setAttribute(address identity, bytes32 name, bytes value, uint256 validity) returns()
func (_EthereumDIDRegistry *EthereumDIDRegistryTransactor) SetAttribute(opts *bind.TransactOpts, identity common.Address, name [32]byte, value []byte, validity *big.Int) (*types.Transaction, error) {
	return _EthereumDIDRegistry.contract.Transact(opts, "setAttribute", identity, name, value, validity)
}

// SetAttribute is a paid mutator transaction binding the contract method 0x7ad4b0a4.
//
// Solidity: function setAttribute(address identity, bytes32 name, bytes value, uint256 validity) returns()
func (_EthereumDIDRegistry *EthereumDIDRegistrySession) SetAttribute(identity common.Address, name [32]byte, value []byte, validity *big.Int) (*types.Transaction, error) {
	return _EthereumDIDRegistry.Contract.SetAttribute(&_EthereumDIDRegistry.TransactOpts, identity, name, value, validity)
}

// SetAttribute is a paid mutator transaction binding the contract method 0x7ad4b0a4.
//
// Solidity: function setAttribute(address identity, bytes32 name, bytes value, uint256 validity) returns()
func (_EthereumDIDRegistry *EthereumDIDRegistryTransactorSession) SetAttribute(identity common.Address, name [32]byte, value []byte, validity *big.Int) (*types.Transaction, error) {
	return _EthereumDIDRegistry.Contract.SetAttribute(&_EthereumDIDRegistry.TransactOpts, identity, name, value, validity)
}

// SetAttributeSigned is a paid mutator transaction binding the contract method 0x123b5e98.
//
// Solidity: function setAttributeSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 name, bytes value, uint256 validity) returns()
func (_EthereumDIDRegistry *EthereumDIDRegistryTransactor) SetAttributeSigned(opts *bind.TransactOpts, identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, name [32]byte, value []byte, validity *big.Int) (*types.Transaction, error) {
	return _EthereumDIDRegistry.contract.Transact(opts, "setAttributeSigned", identity, sigV, sigR, sigS, name, value, validity)
}

// SetAttributeSigned is a paid mutator transaction binding the contract method 0x123b5e98.
//
// Solidity: function setAttributeSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 name, bytes value, uint256 validity) returns()
func (_EthereumDIDRegistry *EthereumDIDRegistrySession) SetAttributeSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, name [32]byte, value []byte, validity *big.Int) (*types.Transaction, error) {
	return _EthereumDIDRegistry.Contract.SetAttributeSigned(&_EthereumDIDRegistry.TransactOpts, identity, sigV, sigR, sigS, name, value, validity)
}

// SetAttributeSigned is a paid mutator transaction binding the contract method 0x123b5e98.
//
// Solidity: function setAttributeSigned(address identity, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 name, bytes value, uint256 validity) returns()
func (_EthereumDIDRegistry *EthereumDIDRegistryTransactorSession) SetAttributeSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, name [32]byte, value []byte, validity *big.Int) (*types.Transaction, error) {
	return _EthereumDIDRegistry.Contract.SetAttributeSigned(&_EthereumDIDRegistry.TransactOpts, identity, sigV, sigR, sigS, name, value, validity)
}

// EthereumDIDRegistryDIDAttributeChangedIterator is returned from FilterDIDAttributeChanged and is used to iterate over the raw logs and unpacked data for DIDAttributeChanged events raised by the EthereumDIDRegistry contract.
type EthereumDIDRegistryDIDAttributeChangedIterator struct {
	Event *EthereumDIDRegistryDIDAttributeChanged // Event containing the contract specifics and raw log

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
func (it *EthereumDIDRegistryDIDAttributeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumDIDRegistryDIDAttributeChanged)
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
		it.Event = new(EthereumDIDRegistryDIDAttributeChanged)
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
func (it *EthereumDIDRegistryDIDAttributeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumDIDRegistryDIDAttributeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumDIDRegistryDIDAttributeChanged represents a DIDAttributeChanged event raised by the EthereumDIDRegistry contract.
type EthereumDIDRegistryDIDAttributeChanged struct {
	Identity       common.Address
	Name           [32]byte
	Value          []byte
	ValidTo        *big.Int
	PreviousChange *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDIDAttributeChanged is a free log retrieval operation binding the contract event 0x18ab6b2ae3d64306c00ce663125f2bd680e441a098de1635bd7ad8b0d44965e4.
//
// Solidity: event DIDAttributeChanged(address indexed identity, bytes32 name, bytes value, uint256 validTo, uint256 previousChange)
func (_EthereumDIDRegistry *EthereumDIDRegistryFilterer) FilterDIDAttributeChanged(opts *bind.FilterOpts, identity []common.Address) (*EthereumDIDRegistryDIDAttributeChangedIterator, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _EthereumDIDRegistry.contract.FilterLogs(opts, "DIDAttributeChanged", identityRule)
	if err != nil {
		return nil, err
	}
	return &EthereumDIDRegistryDIDAttributeChangedIterator{contract: _EthereumDIDRegistry.contract, event: "DIDAttributeChanged", logs: logs, sub: sub}, nil
}

// WatchDIDAttributeChanged is a free log subscription operation binding the contract event 0x18ab6b2ae3d64306c00ce663125f2bd680e441a098de1635bd7ad8b0d44965e4.
//
// Solidity: event DIDAttributeChanged(address indexed identity, bytes32 name, bytes value, uint256 validTo, uint256 previousChange)
func (_EthereumDIDRegistry *EthereumDIDRegistryFilterer) WatchDIDAttributeChanged(opts *bind.WatchOpts, sink chan<- *EthereumDIDRegistryDIDAttributeChanged, identity []common.Address) (event.Subscription, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _EthereumDIDRegistry.contract.WatchLogs(opts, "DIDAttributeChanged", identityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumDIDRegistryDIDAttributeChanged)
				if err := _EthereumDIDRegistry.contract.UnpackLog(event, "DIDAttributeChanged", log); err != nil {
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

// ParseDIDAttributeChanged is a log parse operation binding the contract event 0x18ab6b2ae3d64306c00ce663125f2bd680e441a098de1635bd7ad8b0d44965e4.
//
// Solidity: event DIDAttributeChanged(address indexed identity, bytes32 name, bytes value, uint256 validTo, uint256 previousChange)
func (_EthereumDIDRegistry *EthereumDIDRegistryFilterer) ParseDIDAttributeChanged(log types.Log) (*EthereumDIDRegistryDIDAttributeChanged, error) {
	event := new(EthereumDIDRegistryDIDAttributeChanged)
	if err := _EthereumDIDRegistry.contract.UnpackLog(event, "DIDAttributeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumDIDRegistryDIDDelegateChangedIterator is returned from FilterDIDDelegateChanged and is used to iterate over the raw logs and unpacked data for DIDDelegateChanged events raised by the EthereumDIDRegistry contract.
type EthereumDIDRegistryDIDDelegateChangedIterator struct {
	Event *EthereumDIDRegistryDIDDelegateChanged // Event containing the contract specifics and raw log

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
func (it *EthereumDIDRegistryDIDDelegateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumDIDRegistryDIDDelegateChanged)
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
		it.Event = new(EthereumDIDRegistryDIDDelegateChanged)
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
func (it *EthereumDIDRegistryDIDDelegateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumDIDRegistryDIDDelegateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumDIDRegistryDIDDelegateChanged represents a DIDDelegateChanged event raised by the EthereumDIDRegistry contract.
type EthereumDIDRegistryDIDDelegateChanged struct {
	Identity       common.Address
	DelegateType   [32]byte
	Delegate       common.Address
	ValidTo        *big.Int
	PreviousChange *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDIDDelegateChanged is a free log retrieval operation binding the contract event 0x5a5084339536bcab65f20799fcc58724588145ca054bd2be626174b27ba156f7.
//
// Solidity: event DIDDelegateChanged(address indexed identity, bytes32 delegateType, address delegate, uint256 validTo, uint256 previousChange)
func (_EthereumDIDRegistry *EthereumDIDRegistryFilterer) FilterDIDDelegateChanged(opts *bind.FilterOpts, identity []common.Address) (*EthereumDIDRegistryDIDDelegateChangedIterator, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _EthereumDIDRegistry.contract.FilterLogs(opts, "DIDDelegateChanged", identityRule)
	if err != nil {
		return nil, err
	}
	return &EthereumDIDRegistryDIDDelegateChangedIterator{contract: _EthereumDIDRegistry.contract, event: "DIDDelegateChanged", logs: logs, sub: sub}, nil
}

// WatchDIDDelegateChanged is a free log subscription operation binding the contract event 0x5a5084339536bcab65f20799fcc58724588145ca054bd2be626174b27ba156f7.
//
// Solidity: event DIDDelegateChanged(address indexed identity, bytes32 delegateType, address delegate, uint256 validTo, uint256 previousChange)
func (_EthereumDIDRegistry *EthereumDIDRegistryFilterer) WatchDIDDelegateChanged(opts *bind.WatchOpts, sink chan<- *EthereumDIDRegistryDIDDelegateChanged, identity []common.Address) (event.Subscription, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _EthereumDIDRegistry.contract.WatchLogs(opts, "DIDDelegateChanged", identityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumDIDRegistryDIDDelegateChanged)
				if err := _EthereumDIDRegistry.contract.UnpackLog(event, "DIDDelegateChanged", log); err != nil {
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

// ParseDIDDelegateChanged is a log parse operation binding the contract event 0x5a5084339536bcab65f20799fcc58724588145ca054bd2be626174b27ba156f7.
//
// Solidity: event DIDDelegateChanged(address indexed identity, bytes32 delegateType, address delegate, uint256 validTo, uint256 previousChange)
func (_EthereumDIDRegistry *EthereumDIDRegistryFilterer) ParseDIDDelegateChanged(log types.Log) (*EthereumDIDRegistryDIDDelegateChanged, error) {
	event := new(EthereumDIDRegistryDIDDelegateChanged)
	if err := _EthereumDIDRegistry.contract.UnpackLog(event, "DIDDelegateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumDIDRegistryDIDOwnerChangedIterator is returned from FilterDIDOwnerChanged and is used to iterate over the raw logs and unpacked data for DIDOwnerChanged events raised by the EthereumDIDRegistry contract.
type EthereumDIDRegistryDIDOwnerChangedIterator struct {
	Event *EthereumDIDRegistryDIDOwnerChanged // Event containing the contract specifics and raw log

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
func (it *EthereumDIDRegistryDIDOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumDIDRegistryDIDOwnerChanged)
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
		it.Event = new(EthereumDIDRegistryDIDOwnerChanged)
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
func (it *EthereumDIDRegistryDIDOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumDIDRegistryDIDOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumDIDRegistryDIDOwnerChanged represents a DIDOwnerChanged event raised by the EthereumDIDRegistry contract.
type EthereumDIDRegistryDIDOwnerChanged struct {
	Identity       common.Address
	Owner          common.Address
	PreviousChange *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDIDOwnerChanged is a free log retrieval operation binding the contract event 0x38a5a6e68f30ed1ab45860a4afb34bcb2fc00f22ca462d249b8a8d40cda6f7a3.
//
// Solidity: event DIDOwnerChanged(address indexed identity, address owner, uint256 previousChange)
func (_EthereumDIDRegistry *EthereumDIDRegistryFilterer) FilterDIDOwnerChanged(opts *bind.FilterOpts, identity []common.Address) (*EthereumDIDRegistryDIDOwnerChangedIterator, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _EthereumDIDRegistry.contract.FilterLogs(opts, "DIDOwnerChanged", identityRule)
	if err != nil {
		return nil, err
	}
	return &EthereumDIDRegistryDIDOwnerChangedIterator{contract: _EthereumDIDRegistry.contract, event: "DIDOwnerChanged", logs: logs, sub: sub}, nil
}

// WatchDIDOwnerChanged is a free log subscription operation binding the contract event 0x38a5a6e68f30ed1ab45860a4afb34bcb2fc00f22ca462d249b8a8d40cda6f7a3.
//
// Solidity: event DIDOwnerChanged(address indexed identity, address owner, uint256 previousChange)
func (_EthereumDIDRegistry *EthereumDIDRegistryFilterer) WatchDIDOwnerChanged(opts *bind.WatchOpts, sink chan<- *EthereumDIDRegistryDIDOwnerChanged, identity []common.Address) (event.Subscription, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _EthereumDIDRegistry.contract.WatchLogs(opts, "DIDOwnerChanged", identityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumDIDRegistryDIDOwnerChanged)
				if err := _EthereumDIDRegistry.contract.UnpackLog(event, "DIDOwnerChanged", log); err != nil {
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

// ParseDIDOwnerChanged is a log parse operation binding the contract event 0x38a5a6e68f30ed1ab45860a4afb34bcb2fc00f22ca462d249b8a8d40cda6f7a3.
//
// Solidity: event DIDOwnerChanged(address indexed identity, address owner, uint256 previousChange)
func (_EthereumDIDRegistry *EthereumDIDRegistryFilterer) ParseDIDOwnerChanged(log types.Log) (*EthereumDIDRegistryDIDOwnerChanged, error) {
	event := new(EthereumDIDRegistryDIDOwnerChanged)
	if err := _EthereumDIDRegistry.contract.UnpackLog(event, "DIDOwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
