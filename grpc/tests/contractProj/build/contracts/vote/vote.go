// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vote

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
)

// VoteMetaData contains all meta data concerning the Vote contract.
var VoteMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proposalNames\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"chairperson\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"voted\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"vote\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"giveRightToVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposal\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"winningProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"winningProposal_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"winnerName\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"winnerName_\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200152538038062001525833981810160405281019062000037919062000342565b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060018060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000018190555060005b81518110156200017657600260405180604001604052808484815181106200010f576200010e62000393565b5b60200260200101518152602001600081525090806001815401808255809150506001900390600052602060002090600202016000909190919091506000820151816000015560208201518160010155505080806200016d90620003fb565b915050620000e2565b505062000448565b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b620001e28262000197565b810181811067ffffffffffffffff82111715620002045762000203620001a8565b5b80604052505050565b6000620002196200017e565b9050620002278282620001d7565b919050565b600067ffffffffffffffff8211156200024a5762000249620001a8565b5b602082029050602081019050919050565b600080fd5b6000819050919050565b620002758162000260565b81146200028157600080fd5b50565b60008151905062000295816200026a565b92915050565b6000620002b2620002ac846200022c565b6200020d565b90508083825260208201905060208402830185811115620002d857620002d76200025b565b5b835b81811015620003055780620002f0888262000284565b845260208401935050602081019050620002da565b5050509392505050565b600082601f83011262000327576200032662000192565b5b8151620003398482602086016200029b565b91505092915050565b6000602082840312156200035b576200035a62000188565b5b600082015167ffffffffffffffff8111156200037c576200037b6200018d565b5b6200038a848285016200030f565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000819050919050565b60006200040882620003f1565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036200043d576200043c620003c2565b5b600182019050919050565b6110cd80620004586000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063609ff1bd1161005b578063609ff1bd146101145780639e7b8d6114610132578063a3ec138d1461014e578063e2ba53f01461018157610088565b80630121b93f1461008d578063013cf08b146100a95780632e4176cf146100da5780635c19a95c146100f8575b600080fd5b6100a760048036038101906100a29190610a3b565b61019f565b005b6100c360048036038101906100be9190610a3b565b6102e5565b6040516100d1929190610a90565b60405180910390f35b6100e2610319565b6040516100ef9190610afa565b60405180910390f35b610112600480360381019061010d9190610b41565b61033d565b005b61011c610730565b6040516101299190610b6e565b60405180910390f35b61014c60048036038101906101479190610b41565b6107b8565b005b61016860048036038101906101639190610b41565b61096f565b6040516101789493929190610ba4565b60405180910390f35b6101896109cc565b6040516101969190610be9565b60405180910390f35b6000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000816000015403610229576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161022090610c61565b60405180910390fd5b8060010160009054906101000a900460ff161561027b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161027290610ccd565b60405180910390fd5b60018160010160006101000a81548160ff0219169083151502179055508181600201819055508060000154600283815481106102ba576102b9610ced565b5b906000526020600020906002020160010160008282546102da9190610d4b565b925050819055505050565b600281815481106102f557600080fd5b90600052602060002090600202016000915090508060000154908060010154905082565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060008160000154036103c7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103be90610ded565b60405180910390fd5b8060010160009054906101000a900460ff1615610419576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161041090610e59565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610487576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161047e90610ec5565b60405180910390fd5b5b600073ffffffffffffffffffffffffffffffffffffffff16600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146105f657600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691503373ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036105f1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105e890610f31565b60405180910390fd5b610488565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060018160000154101561064b57600080fd5b60018260010160006101000a81548160ff021916908315150217905550828260010160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060010160009054906101000a900460ff161561070b57816000015460028260020154815481106106df576106de610ced565b5b906000526020600020906002020160010160008282546106ff9190610d4b565b9250508190555061072b565b81600001548160000160008282546107239190610d4b565b925050819055505b505050565b6000806000905060005b6002805490508110156107b357816002828154811061075c5761075b610ced565b5b90600052602060002090600202016001015411156107a0576002818154811061078857610787610ced565b5b90600052602060002090600202016001015491508092505b80806107ab90610f51565b91505061073a565b505090565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610846576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161083d9061100b565b60405180910390fd5b600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff16156108d6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108cd90611077565b60405180910390fd5b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001541461092557600080fd5b60018060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000018190555050565b60016020528060005260406000206000915090508060000154908060010160009054906101000a900460ff16908060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020154905084565b600060026109d8610730565b815481106109e9576109e8610ced565b5b906000526020600020906002020160000154905090565b600080fd5b6000819050919050565b610a1881610a05565b8114610a2357600080fd5b50565b600081359050610a3581610a0f565b92915050565b600060208284031215610a5157610a50610a00565b5b6000610a5f84828501610a26565b91505092915050565b6000819050919050565b610a7b81610a68565b82525050565b610a8a81610a05565b82525050565b6000604082019050610aa56000830185610a72565b610ab26020830184610a81565b9392505050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610ae482610ab9565b9050919050565b610af481610ad9565b82525050565b6000602082019050610b0f6000830184610aeb565b92915050565b610b1e81610ad9565b8114610b2957600080fd5b50565b600081359050610b3b81610b15565b92915050565b600060208284031215610b5757610b56610a00565b5b6000610b6584828501610b2c565b91505092915050565b6000602082019050610b836000830184610a81565b92915050565b60008115159050919050565b610b9e81610b89565b82525050565b6000608082019050610bb96000830187610a81565b610bc66020830186610b95565b610bd36040830185610aeb565b610be06060830184610a81565b95945050505050565b6000602082019050610bfe6000830184610a72565b92915050565b600082825260208201905092915050565b7f486173206e6f20726967687420746f20766f7465000000000000000000000000600082015250565b6000610c4b601483610c04565b9150610c5682610c15565b602082019050919050565b60006020820190508181036000830152610c7a81610c3e565b9050919050565b7f416c726561647920766f7465642e000000000000000000000000000000000000600082015250565b6000610cb7600e83610c04565b9150610cc282610c81565b602082019050919050565b60006020820190508181036000830152610ce681610caa565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610d5682610a05565b9150610d6183610a05565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115610d9657610d95610d1c565b5b828201905092915050565b7f596f752068617665206e6f20726967687420746f20766f746500000000000000600082015250565b6000610dd7601983610c04565b9150610de282610da1565b602082019050919050565b60006020820190508181036000830152610e0681610dca565b9050919050565b7f596f7520616c726561647920766f7465642e0000000000000000000000000000600082015250565b6000610e43601283610c04565b9150610e4e82610e0d565b602082019050919050565b60006020820190508181036000830152610e7281610e36565b9050919050565b7f53656c662d64656c65676174696f6e20697320646973616c6c6f7765642e0000600082015250565b6000610eaf601e83610c04565b9150610eba82610e79565b602082019050919050565b60006020820190508181036000830152610ede81610ea2565b9050919050565b7f466f756e64206c6f6f7020696e2064656c65676174696f6e2e00000000000000600082015250565b6000610f1b601983610c04565b9150610f2682610ee5565b602082019050919050565b60006020820190508181036000830152610f4a81610f0e565b9050919050565b6000610f5c82610a05565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610f8e57610f8d610d1c565b5b600182019050919050565b7f4f6e6c79206368616972706572736f6e2063616e20676976652072696768742060008201527f746f20766f74652e000000000000000000000000000000000000000000000000602082015250565b6000610ff5602883610c04565b915061100082610f99565b604082019050919050565b6000602082019050818103600083015261102481610fe8565b9050919050565b7f54686520766f74657220616c726561647920766f7465642e0000000000000000600082015250565b6000611061601883610c04565b915061106c8261102b565b602082019050919050565b6000602082019050818103600083015261109081611054565b905091905056fea26469706673582212204b7abaebe72ebf8cef5f650088a86536beaf01afe7069f3ef43d70d02098699264736f6c634300080f0033",
}

// VoteABI is the input ABI used to generate the binding from.
// Deprecated: Use VoteMetaData.ABI instead.
var VoteABI = VoteMetaData.ABI

// VoteBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VoteMetaData.Bin instead.
var VoteBin = VoteMetaData.Bin

// DeployVote deploys a new Ethereum contract, binding an instance of Vote to it.
func DeployVote(auth *bind.TransactOpts, backend bind.ContractBackend, proposalNames [][32]byte) (common.Address, *types.Transaction, *Vote, error) {
	parsed, err := VoteMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VoteBin), backend, proposalNames)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Vote{VoteCaller: VoteCaller{contract: contract}, VoteTransactor: VoteTransactor{contract: contract}, VoteFilterer: VoteFilterer{contract: contract}}, nil
}

// Vote is an auto generated Go binding around an Ethereum contract.
type Vote struct {
	VoteCaller     // Read-only binding to the contract
	VoteTransactor // Write-only binding to the contract
	VoteFilterer   // Log filterer for contract events
}

// VoteCaller is an auto generated read-only Go binding around an Ethereum contract.
type VoteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VoteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VoteFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VoteSession struct {
	Contract     *Vote             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VoteCallerSession struct {
	Contract *VoteCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VoteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VoteTransactorSession struct {
	Contract     *VoteTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoteRaw is an auto generated low-level Go binding around an Ethereum contract.
type VoteRaw struct {
	Contract *Vote // Generic contract binding to access the raw methods on
}

// VoteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VoteCallerRaw struct {
	Contract *VoteCaller // Generic read-only contract binding to access the raw methods on
}

// VoteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VoteTransactorRaw struct {
	Contract *VoteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVote creates a new instance of Vote, bound to a specific deployed contract.
func NewVote(address common.Address, backend bind.ContractBackend) (*Vote, error) {
	contract, err := bindVote(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vote{VoteCaller: VoteCaller{contract: contract}, VoteTransactor: VoteTransactor{contract: contract}, VoteFilterer: VoteFilterer{contract: contract}}, nil
}

// NewVoteCaller creates a new read-only instance of Vote, bound to a specific deployed contract.
func NewVoteCaller(address common.Address, caller bind.ContractCaller) (*VoteCaller, error) {
	contract, err := bindVote(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VoteCaller{contract: contract}, nil
}

// NewVoteTransactor creates a new write-only instance of Vote, bound to a specific deployed contract.
func NewVoteTransactor(address common.Address, transactor bind.ContractTransactor) (*VoteTransactor, error) {
	contract, err := bindVote(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VoteTransactor{contract: contract}, nil
}

// NewVoteFilterer creates a new log filterer instance of Vote, bound to a specific deployed contract.
func NewVoteFilterer(address common.Address, filterer bind.ContractFilterer) (*VoteFilterer, error) {
	contract, err := bindVote(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VoteFilterer{contract: contract}, nil
}

// bindVote binds a generic wrapper to an already deployed contract.
func bindVote(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VoteABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vote *VoteRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vote.Contract.VoteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vote *VoteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vote.Contract.VoteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vote *VoteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vote.Contract.VoteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vote *VoteCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vote.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vote *VoteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vote.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vote *VoteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vote.Contract.contract.Transact(opts, method, params...)
}

// Chairperson is a free data retrieval call binding the contract method 0x2e4176cf.
//
// Solidity: function chairperson() view returns(address)
func (_Vote *VoteCaller) Chairperson(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "chairperson")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Chairperson is a free data retrieval call binding the contract method 0x2e4176cf.
//
// Solidity: function chairperson() view returns(address)
func (_Vote *VoteSession) Chairperson() (common.Address, error) {
	return _Vote.Contract.Chairperson(&_Vote.CallOpts)
}

// Chairperson is a free data retrieval call binding the contract method 0x2e4176cf.
//
// Solidity: function chairperson() view returns(address)
func (_Vote *VoteCallerSession) Chairperson() (common.Address, error) {
	return _Vote.Contract.Chairperson(&_Vote.CallOpts)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(bytes32 name, uint256 voteCount)
func (_Vote *VoteCaller) Proposals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name      [32]byte
	VoteCount *big.Int
}, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "proposals", arg0)

	outstruct := new(struct {
		Name      [32]byte
		VoteCount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.VoteCount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(bytes32 name, uint256 voteCount)
func (_Vote *VoteSession) Proposals(arg0 *big.Int) (struct {
	Name      [32]byte
	VoteCount *big.Int
}, error) {
	return _Vote.Contract.Proposals(&_Vote.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(bytes32 name, uint256 voteCount)
func (_Vote *VoteCallerSession) Proposals(arg0 *big.Int) (struct {
	Name      [32]byte
	VoteCount *big.Int
}, error) {
	return _Vote.Contract.Proposals(&_Vote.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(uint256 weight, bool voted, address delegate, uint256 vote)
func (_Vote *VoteCaller) Voters(opts *bind.CallOpts, arg0 common.Address) (struct {
	Weight   *big.Int
	Voted    bool
	Delegate common.Address
	Vote     *big.Int
}, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "voters", arg0)

	outstruct := new(struct {
		Weight   *big.Int
		Voted    bool
		Delegate common.Address
		Vote     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Weight = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Voted = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Delegate = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Vote = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(uint256 weight, bool voted, address delegate, uint256 vote)
func (_Vote *VoteSession) Voters(arg0 common.Address) (struct {
	Weight   *big.Int
	Voted    bool
	Delegate common.Address
	Vote     *big.Int
}, error) {
	return _Vote.Contract.Voters(&_Vote.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(uint256 weight, bool voted, address delegate, uint256 vote)
func (_Vote *VoteCallerSession) Voters(arg0 common.Address) (struct {
	Weight   *big.Int
	Voted    bool
	Delegate common.Address
	Vote     *big.Int
}, error) {
	return _Vote.Contract.Voters(&_Vote.CallOpts, arg0)
}

// WinnerName is a free data retrieval call binding the contract method 0xe2ba53f0.
//
// Solidity: function winnerName() view returns(bytes32 winnerName_)
func (_Vote *VoteCaller) WinnerName(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "winnerName")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WinnerName is a free data retrieval call binding the contract method 0xe2ba53f0.
//
// Solidity: function winnerName() view returns(bytes32 winnerName_)
func (_Vote *VoteSession) WinnerName() ([32]byte, error) {
	return _Vote.Contract.WinnerName(&_Vote.CallOpts)
}

// WinnerName is a free data retrieval call binding the contract method 0xe2ba53f0.
//
// Solidity: function winnerName() view returns(bytes32 winnerName_)
func (_Vote *VoteCallerSession) WinnerName() ([32]byte, error) {
	return _Vote.Contract.WinnerName(&_Vote.CallOpts)
}

// WinningProposal is a free data retrieval call binding the contract method 0x609ff1bd.
//
// Solidity: function winningProposal() view returns(uint256 winningProposal_)
func (_Vote *VoteCaller) WinningProposal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "winningProposal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WinningProposal is a free data retrieval call binding the contract method 0x609ff1bd.
//
// Solidity: function winningProposal() view returns(uint256 winningProposal_)
func (_Vote *VoteSession) WinningProposal() (*big.Int, error) {
	return _Vote.Contract.WinningProposal(&_Vote.CallOpts)
}

// WinningProposal is a free data retrieval call binding the contract method 0x609ff1bd.
//
// Solidity: function winningProposal() view returns(uint256 winningProposal_)
func (_Vote *VoteCallerSession) WinningProposal() (*big.Int, error) {
	return _Vote.Contract.WinningProposal(&_Vote.CallOpts)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address to) returns()
func (_Vote *VoteTransactor) Delegate(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "delegate", to)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address to) returns()
func (_Vote *VoteSession) Delegate(to common.Address) (*types.Transaction, error) {
	return _Vote.Contract.Delegate(&_Vote.TransactOpts, to)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address to) returns()
func (_Vote *VoteTransactorSession) Delegate(to common.Address) (*types.Transaction, error) {
	return _Vote.Contract.Delegate(&_Vote.TransactOpts, to)
}

// GiveRightToVote is a paid mutator transaction binding the contract method 0x9e7b8d61.
//
// Solidity: function giveRightToVote(address voter) returns()
func (_Vote *VoteTransactor) GiveRightToVote(opts *bind.TransactOpts, voter common.Address) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "giveRightToVote", voter)
}

// GiveRightToVote is a paid mutator transaction binding the contract method 0x9e7b8d61.
//
// Solidity: function giveRightToVote(address voter) returns()
func (_Vote *VoteSession) GiveRightToVote(voter common.Address) (*types.Transaction, error) {
	return _Vote.Contract.GiveRightToVote(&_Vote.TransactOpts, voter)
}

// GiveRightToVote is a paid mutator transaction binding the contract method 0x9e7b8d61.
//
// Solidity: function giveRightToVote(address voter) returns()
func (_Vote *VoteTransactorSession) GiveRightToVote(voter common.Address) (*types.Transaction, error) {
	return _Vote.Contract.GiveRightToVote(&_Vote.TransactOpts, voter)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 proposal) returns()
func (_Vote *VoteTransactor) Vote(opts *bind.TransactOpts, proposal *big.Int) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "vote", proposal)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 proposal) returns()
func (_Vote *VoteSession) Vote(proposal *big.Int) (*types.Transaction, error) {
	return _Vote.Contract.Vote(&_Vote.TransactOpts, proposal)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 proposal) returns()
func (_Vote *VoteTransactorSession) Vote(proposal *big.Int) (*types.Transaction, error) {
	return _Vote.Contract.Vote(&_Vote.TransactOpts, proposal)
}
