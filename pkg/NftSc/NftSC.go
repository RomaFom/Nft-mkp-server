// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nft_api

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

// NftApiMetaData contains all meta data concerning the NftApi contract.
var NftApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_tokenURI\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040518060400160405280600881526020016711105c1c0813919560c21b815250604051806040016040528060048152602001630444150560e41b81525081600090816200006191906200011e565b5060016200007082826200011e565b505050620001ea565b634e487b7160e01b600052604160045260246000fd5b600181811c90821680620000a457607f821691505b602082108103620000c557634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200011957600081815260208120601f850160051c81016020861015620000f45750805b601f850160051c820191505b81811015620001155782815560010162000100565b5050505b505050565b81516001600160401b038111156200013a576200013a62000079565b62000152816200014b84546200008f565b84620000cb565b602080601f8311600181146200018a5760008415620001715750858301515b600019600386901b1c1916600185901b17855562000115565b600085815260208120601f198616915b82811015620001bb578886015182559484019460019091019084016200019a565b5085821015620001da5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b61164a80620001fa6000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c806395d89b4111610097578063c41a360a11610066578063c41a360a14610213578063c87b56dd14610226578063d85d3d2714610239578063e985e9c51461024c57600080fd5b806395d89b41146101dc5780639f181b5e146101e4578063a22cb465146101ed578063b88d4fde1461020057600080fd5b806323b872dd116100d357806323b872dd1461018257806342842e0e146101955780636352211e146101a857806370a08231146101bb57600080fd5b806301ffc9a71461010557806306fdde031461012d578063081812fc14610142578063095ea7b31461016d575b600080fd5b610118610113366004611004565b61025f565b60405190151581526020015b60405180910390f35b6101356102b1565b6040516101249190611071565b610155610150366004611084565b610343565b6040516001600160a01b039091168152602001610124565b61018061017b3660046110b9565b61036a565b005b6101806101903660046110e3565b610484565b6101806101a33660046110e3565b6104b5565b6101556101b6366004611084565b6104d0565b6101ce6101c936600461111f565b610530565b604051908152602001610124565b6101356105b6565b6101ce60075481565b6101806101fb36600461113a565b6105c5565b61018061020e366004611202565b6105d4565b610155610221366004611084565b61060c565b610135610234366004611084565b610617565b6101ce61024736600461127e565b610727565b61011861025a3660046112c7565b61075d565b60006001600160e01b031982166380ac58cd60e01b148061029057506001600160e01b03198216635b5e139f60e01b145b806102ab57506301ffc9a760e01b6001600160e01b03198316145b92915050565b6060600080546102c0906112fa565b80601f01602080910402602001604051908101604052809291908181526020018280546102ec906112fa565b80156103395780601f1061030e57610100808354040283529160200191610339565b820191906000526020600020905b81548152906001019060200180831161031c57829003601f168201915b5050505050905090565b600061034e8261078b565b506000908152600460205260409020546001600160a01b031690565b6000610375826104d0565b9050806001600160a01b0316836001600160a01b0316036103e75760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656044820152603960f91b60648201526084015b60405180910390fd5b336001600160a01b03821614806104035750610403813361075d565b6104755760405162461bcd60e51b815260206004820152603e60248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60448201527f6b656e206f776e6572206e6f7220617070726f76656420666f7220616c6c000060648201526084016103de565b61047f83836107ed565b505050565b61048e338261085b565b6104aa5760405162461bcd60e51b81526004016103de90611334565b61047f8383836108b9565b61047f838383604051806020016040528060008152506105d4565b6000818152600260205260408120546001600160a01b0316806102ab5760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b60448201526064016103de565b60006001600160a01b03821661059a5760405162461bcd60e51b815260206004820152602960248201527f4552433732313a2061646472657373207a65726f206973206e6f7420612076616044820152683634b21037bbb732b960b91b60648201526084016103de565b506001600160a01b031660009081526003602052604090205490565b6060600180546102c0906112fa565b6105d0338383610a55565b5050565b6105de338361085b565b6105fa5760405162461bcd60e51b81526004016103de90611334565b61060684848484610b23565b50505050565b60006102ab826104d0565b60606106228261078b565b6000828152600660205260408120805461063b906112fa565b80601f0160208091040260200160405190810160405280929190818152602001828054610667906112fa565b80156106b45780601f10610689576101008083540402835291602001916106b4565b820191906000526020600020905b81548152906001019060200180831161069757829003601f168201915b5050505050905060006106d260408051602081019091526000815290565b905080516000036106e4575092915050565b8151156107165780826040516020016106fe929190611382565b60405160208183030381529060405292505050919050565b61071f84610b56565b949350505050565b6007805460009182610738836113c7565b919050555061074933600754610bca565b61075560075483610be4565b505060075490565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205460ff1690565b6000818152600260205260409020546001600160a01b03166107ea5760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b60448201526064016103de565b50565b600081815260046020526040902080546001600160a01b0319166001600160a01b0384169081179091558190610822826104d0565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b600080610867836104d0565b9050806001600160a01b0316846001600160a01b0316148061088e575061088e818561075d565b8061071f5750836001600160a01b03166108a784610343565b6001600160a01b031614949350505050565b826001600160a01b03166108cc826104d0565b6001600160a01b0316146109305760405162461bcd60e51b815260206004820152602560248201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060448201526437bbb732b960d91b60648201526084016103de565b6001600160a01b0382166109925760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b60648201526084016103de565b61099d6000826107ed565b6001600160a01b03831660009081526003602052604081208054600192906109c69084906113e0565b90915550506001600160a01b03821660009081526003602052604081208054600192906109f49084906113f3565b909155505060008181526002602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b816001600160a01b0316836001600160a01b031603610ab65760405162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c65720000000000000060448201526064016103de565b6001600160a01b03838116600081815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b610b2e8484846108b9565b610b3a84848484610c77565b6106065760405162461bcd60e51b81526004016103de90611406565b6060610b618261078b565b6000610b7860408051602081019091526000815290565b90506000815111610b985760405180602001604052806000815250610bc3565b80610ba284610d78565b604051602001610bb3929190611382565b6040516020818303038152906040525b9392505050565b6105d0828260405180602001604052806000815250610e79565b6000828152600260205260409020546001600160a01b0316610c5f5760405162461bcd60e51b815260206004820152602e60248201527f45524337323155524953746f726167653a2055524920736574206f66206e6f6e60448201526d32bc34b9ba32b73a103a37b5b2b760911b60648201526084016103de565b600082815260066020526040902061047f82826114a6565b60006001600160a01b0384163b15610d6d57604051630a85bd0160e11b81526001600160a01b0385169063150b7a0290610cbb903390899088908890600401611566565b6020604051808303816000875af1925050508015610cf6575060408051601f3d908101601f19168201909252610cf3918101906115a3565b60015b610d53573d808015610d24576040519150601f19603f3d011682016040523d82523d6000602084013e610d29565b606091505b508051600003610d4b5760405162461bcd60e51b81526004016103de90611406565b805181602001fd5b6001600160e01b031916630a85bd0160e11b14905061071f565b506001949350505050565b606081600003610d9f5750506040805180820190915260018152600360fc1b602082015290565b8160005b8115610dc95780610db3816113c7565b9150610dc29050600a836115d6565b9150610da3565b60008167ffffffffffffffff811115610de457610de4611176565b6040519080825280601f01601f191660200182016040528015610e0e576020820181803683370190505b5090505b841561071f57610e236001836113e0565b9150610e30600a866115ea565b610e3b9060306113f3565b60f81b818381518110610e5057610e506115fe565b60200101906001600160f81b031916908160001a905350610e72600a866115d6565b9450610e12565b610e838383610eac565b610e906000848484610c77565b61047f5760405162461bcd60e51b81526004016103de90611406565b6001600160a01b038216610f025760405162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f206164647265737360448201526064016103de565b6000818152600260205260409020546001600160a01b031615610f675760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e7465640000000060448201526064016103de565b6001600160a01b0382166000908152600360205260408120805460019290610f909084906113f3565b909155505060008181526002602052604080822080546001600160a01b0319166001600160a01b03861690811790915590518392907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b6001600160e01b0319811681146107ea57600080fd5b60006020828403121561101657600080fd5b8135610bc381610fee565b60005b8381101561103c578181015183820152602001611024565b50506000910152565b6000815180845261105d816020860160208601611021565b601f01601f19169290920160200192915050565b602081526000610bc36020830184611045565b60006020828403121561109657600080fd5b5035919050565b80356001600160a01b03811681146110b457600080fd5b919050565b600080604083850312156110cc57600080fd5b6110d58361109d565b946020939093013593505050565b6000806000606084860312156110f857600080fd5b6111018461109d565b925061110f6020850161109d565b9150604084013590509250925092565b60006020828403121561113157600080fd5b610bc38261109d565b6000806040838503121561114d57600080fd5b6111568361109d565b91506020830135801515811461116b57600080fd5b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b600067ffffffffffffffff808411156111a7576111a7611176565b604051601f8501601f19908116603f011681019082821181831017156111cf576111cf611176565b816040528093508581528686860111156111e857600080fd5b858560208301376000602087830101525050509392505050565b6000806000806080858703121561121857600080fd5b6112218561109d565b935061122f6020860161109d565b925060408501359150606085013567ffffffffffffffff81111561125257600080fd5b8501601f8101871361126357600080fd5b6112728782356020840161118c565b91505092959194509250565b60006020828403121561129057600080fd5b813567ffffffffffffffff8111156112a757600080fd5b8201601f810184136112b857600080fd5b61071f8482356020840161118c565b600080604083850312156112da57600080fd5b6112e38361109d565b91506112f16020840161109d565b90509250929050565b600181811c9082168061130e57607f821691505b60208210810361132e57634e487b7160e01b600052602260045260246000fd5b50919050565b6020808252602e908201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560408201526d1c881b9bdc88185c1c1c9bdd995960921b606082015260800190565b60008351611394818460208801611021565b8351908301906113a8818360208801611021565b01949350505050565b634e487b7160e01b600052601160045260246000fd5b6000600182016113d9576113d96113b1565b5060010190565b818103818111156102ab576102ab6113b1565b808201808211156102ab576102ab6113b1565b60208082526032908201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b606082015260800190565b601f82111561047f57600081815260208120601f850160051c8101602086101561147f5750805b601f850160051c820191505b8181101561149e5782815560010161148b565b505050505050565b815167ffffffffffffffff8111156114c0576114c0611176565b6114d4816114ce84546112fa565b84611458565b602080601f83116001811461150957600084156114f15750858301515b600019600386901b1c1916600185901b17855561149e565b600085815260208120601f198616915b8281101561153857888601518255948401946001909101908401611519565b50858210156115565787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6001600160a01b038581168252841660208201526040810183905260806060820181905260009061159990830184611045565b9695505050505050565b6000602082840312156115b557600080fd5b8151610bc381610fee565b634e487b7160e01b600052601260045260246000fd5b6000826115e5576115e56115c0565b500490565b6000826115f9576115f96115c0565b500690565b634e487b7160e01b600052603260045260246000fdfea2646970667358221220cbdaa87abc96517f444cae6550b089ac1a35b8b4833b9c7c3c086bd732b4f7e064736f6c63430008110033",
}

// NftApiABI is the input ABI used to generate the binding from.
// Deprecated: Use NftApiMetaData.ABI instead.
var NftApiABI = NftApiMetaData.ABI

// NftApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NftApiMetaData.Bin instead.
var NftApiBin = NftApiMetaData.Bin

// DeployNftApi deploys a new Ethereum contract, binding an instance of NftApi to it.
func DeployNftApi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NftApi, error) {
	parsed, err := NftApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NftApiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NftApi{NftApiCaller: NftApiCaller{contract: contract}, NftApiTransactor: NftApiTransactor{contract: contract}, NftApiFilterer: NftApiFilterer{contract: contract}}, nil
}

// NftApi is an auto generated Go binding around an Ethereum contract.
type NftApi struct {
	NftApiCaller     // Read-only binding to the contract
	NftApiTransactor // Write-only binding to the contract
	NftApiFilterer   // Log filterer for contract events
}

// NftApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type NftApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NftApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NftApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NftApiSession struct {
	Contract     *NftApi           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NftApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NftApiCallerSession struct {
	Contract *NftApiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NftApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NftApiTransactorSession struct {
	Contract     *NftApiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NftApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type NftApiRaw struct {
	Contract *NftApi // Generic contract binding to access the raw methods on
}

// NftApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NftApiCallerRaw struct {
	Contract *NftApiCaller // Generic read-only contract binding to access the raw methods on
}

// NftApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NftApiTransactorRaw struct {
	Contract *NftApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNftApi creates a new instance of NftApi, bound to a specific deployed contract.
func NewNftApi(address common.Address, backend bind.ContractBackend) (*NftApi, error) {
	contract, err := bindNftApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NftApi{NftApiCaller: NftApiCaller{contract: contract}, NftApiTransactor: NftApiTransactor{contract: contract}, NftApiFilterer: NftApiFilterer{contract: contract}}, nil
}

// NewNftApiCaller creates a new read-only instance of NftApi, bound to a specific deployed contract.
func NewNftApiCaller(address common.Address, caller bind.ContractCaller) (*NftApiCaller, error) {
	contract, err := bindNftApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NftApiCaller{contract: contract}, nil
}

// NewNftApiTransactor creates a new write-only instance of NftApi, bound to a specific deployed contract.
func NewNftApiTransactor(address common.Address, transactor bind.ContractTransactor) (*NftApiTransactor, error) {
	contract, err := bindNftApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NftApiTransactor{contract: contract}, nil
}

// NewNftApiFilterer creates a new log filterer instance of NftApi, bound to a specific deployed contract.
func NewNftApiFilterer(address common.Address, filterer bind.ContractFilterer) (*NftApiFilterer, error) {
	contract, err := bindNftApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NftApiFilterer{contract: contract}, nil
}

// bindNftApi binds a generic wrapper to an already deployed contract.
func bindNftApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NftApiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NftApi *NftApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NftApi.Contract.NftApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NftApi *NftApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NftApi.Contract.NftApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NftApi *NftApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NftApi.Contract.NftApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NftApi *NftApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NftApi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NftApi *NftApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NftApi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NftApi *NftApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NftApi.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_NftApi *NftApiCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NftApi.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_NftApi *NftApiSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _NftApi.Contract.BalanceOf(&_NftApi.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_NftApi *NftApiCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _NftApi.Contract.BalanceOf(&_NftApi.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_NftApi *NftApiCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NftApi.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_NftApi *NftApiSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _NftApi.Contract.GetApproved(&_NftApi.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_NftApi *NftApiCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _NftApi.Contract.GetApproved(&_NftApi.CallOpts, tokenId)
}

// GetOwner is a free data retrieval call binding the contract method 0xc41a360a.
//
// Solidity: function getOwner(uint256 _tokenId) view returns(address)
func (_NftApi *NftApiCaller) GetOwner(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NftApi.contract.Call(opts, &out, "getOwner", _tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0xc41a360a.
//
// Solidity: function getOwner(uint256 _tokenId) view returns(address)
func (_NftApi *NftApiSession) GetOwner(_tokenId *big.Int) (common.Address, error) {
	return _NftApi.Contract.GetOwner(&_NftApi.CallOpts, _tokenId)
}

// GetOwner is a free data retrieval call binding the contract method 0xc41a360a.
//
// Solidity: function getOwner(uint256 _tokenId) view returns(address)
func (_NftApi *NftApiCallerSession) GetOwner(_tokenId *big.Int) (common.Address, error) {
	return _NftApi.Contract.GetOwner(&_NftApi.CallOpts, _tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_NftApi *NftApiCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _NftApi.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_NftApi *NftApiSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _NftApi.Contract.IsApprovedForAll(&_NftApi.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_NftApi *NftApiCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _NftApi.Contract.IsApprovedForAll(&_NftApi.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NftApi *NftApiCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NftApi.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NftApi *NftApiSession) Name() (string, error) {
	return _NftApi.Contract.Name(&_NftApi.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NftApi *NftApiCallerSession) Name() (string, error) {
	return _NftApi.Contract.Name(&_NftApi.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_NftApi *NftApiCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NftApi.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_NftApi *NftApiSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _NftApi.Contract.OwnerOf(&_NftApi.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_NftApi *NftApiCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _NftApi.Contract.OwnerOf(&_NftApi.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NftApi *NftApiCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _NftApi.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NftApi *NftApiSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NftApi.Contract.SupportsInterface(&_NftApi.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NftApi *NftApiCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NftApi.Contract.SupportsInterface(&_NftApi.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NftApi *NftApiCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NftApi.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NftApi *NftApiSession) Symbol() (string, error) {
	return _NftApi.Contract.Symbol(&_NftApi.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NftApi *NftApiCallerSession) Symbol() (string, error) {
	return _NftApi.Contract.Symbol(&_NftApi.CallOpts)
}

// TokenCount is a free data retrieval call binding the contract method 0x9f181b5e.
//
// Solidity: function tokenCount() view returns(uint256)
func (_NftApi *NftApiCaller) TokenCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NftApi.contract.Call(opts, &out, "tokenCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenCount is a free data retrieval call binding the contract method 0x9f181b5e.
//
// Solidity: function tokenCount() view returns(uint256)
func (_NftApi *NftApiSession) TokenCount() (*big.Int, error) {
	return _NftApi.Contract.TokenCount(&_NftApi.CallOpts)
}

// TokenCount is a free data retrieval call binding the contract method 0x9f181b5e.
//
// Solidity: function tokenCount() view returns(uint256)
func (_NftApi *NftApiCallerSession) TokenCount() (*big.Int, error) {
	return _NftApi.Contract.TokenCount(&_NftApi.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_NftApi *NftApiCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _NftApi.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_NftApi *NftApiSession) TokenURI(tokenId *big.Int) (string, error) {
	return _NftApi.Contract.TokenURI(&_NftApi.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_NftApi *NftApiCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _NftApi.Contract.TokenURI(&_NftApi.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_NftApi *NftApiTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftApi.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_NftApi *NftApiSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftApi.Contract.Approve(&_NftApi.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_NftApi *NftApiTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftApi.Contract.Approve(&_NftApi.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0xd85d3d27.
//
// Solidity: function mint(string _tokenURI) returns(uint256)
func (_NftApi *NftApiTransactor) Mint(opts *bind.TransactOpts, _tokenURI string) (*types.Transaction, error) {
	return _NftApi.contract.Transact(opts, "mint", _tokenURI)
}

// Mint is a paid mutator transaction binding the contract method 0xd85d3d27.
//
// Solidity: function mint(string _tokenURI) returns(uint256)
func (_NftApi *NftApiSession) Mint(_tokenURI string) (*types.Transaction, error) {
	return _NftApi.Contract.Mint(&_NftApi.TransactOpts, _tokenURI)
}

// Mint is a paid mutator transaction binding the contract method 0xd85d3d27.
//
// Solidity: function mint(string _tokenURI) returns(uint256)
func (_NftApi *NftApiTransactorSession) Mint(_tokenURI string) (*types.Transaction, error) {
	return _NftApi.Contract.Mint(&_NftApi.TransactOpts, _tokenURI)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_NftApi *NftApiTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftApi.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_NftApi *NftApiSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftApi.Contract.SafeTransferFrom(&_NftApi.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_NftApi *NftApiTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftApi.Contract.SafeTransferFrom(&_NftApi.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_NftApi *NftApiTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _NftApi.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_NftApi *NftApiSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _NftApi.Contract.SafeTransferFrom0(&_NftApi.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_NftApi *NftApiTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _NftApi.Contract.SafeTransferFrom0(&_NftApi.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_NftApi *NftApiTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _NftApi.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_NftApi *NftApiSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _NftApi.Contract.SetApprovalForAll(&_NftApi.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_NftApi *NftApiTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _NftApi.Contract.SetApprovalForAll(&_NftApi.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_NftApi *NftApiTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftApi.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_NftApi *NftApiSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftApi.Contract.TransferFrom(&_NftApi.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_NftApi *NftApiTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NftApi.Contract.TransferFrom(&_NftApi.TransactOpts, from, to, tokenId)
}

// NftApiApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the NftApi contract.
type NftApiApprovalIterator struct {
	Event *NftApiApproval // Event containing the contract specifics and raw log

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
func (it *NftApiApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftApiApproval)
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
		it.Event = new(NftApiApproval)
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
func (it *NftApiApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftApiApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftApiApproval represents a Approval event raised by the NftApi contract.
type NftApiApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_NftApi *NftApiFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*NftApiApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NftApi.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NftApiApprovalIterator{contract: _NftApi.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_NftApi *NftApiFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NftApiApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NftApi.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftApiApproval)
				if err := _NftApi.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_NftApi *NftApiFilterer) ParseApproval(log types.Log) (*NftApiApproval, error) {
	event := new(NftApiApproval)
	if err := _NftApi.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftApiApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the NftApi contract.
type NftApiApprovalForAllIterator struct {
	Event *NftApiApprovalForAll // Event containing the contract specifics and raw log

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
func (it *NftApiApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftApiApprovalForAll)
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
		it.Event = new(NftApiApprovalForAll)
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
func (it *NftApiApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftApiApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftApiApprovalForAll represents a ApprovalForAll event raised by the NftApi contract.
type NftApiApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_NftApi *NftApiFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*NftApiApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _NftApi.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &NftApiApprovalForAllIterator{contract: _NftApi.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_NftApi *NftApiFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *NftApiApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _NftApi.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftApiApprovalForAll)
				if err := _NftApi.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_NftApi *NftApiFilterer) ParseApprovalForAll(log types.Log) (*NftApiApprovalForAll, error) {
	event := new(NftApiApprovalForAll)
	if err := _NftApi.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftApiTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the NftApi contract.
type NftApiTransferIterator struct {
	Event *NftApiTransfer // Event containing the contract specifics and raw log

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
func (it *NftApiTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftApiTransfer)
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
		it.Event = new(NftApiTransfer)
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
func (it *NftApiTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftApiTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftApiTransfer represents a Transfer event raised by the NftApi contract.
type NftApiTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_NftApi *NftApiFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*NftApiTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NftApi.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NftApiTransferIterator{contract: _NftApi.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_NftApi *NftApiFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NftApiTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NftApi.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftApiTransfer)
				if err := _NftApi.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_NftApi *NftApiFilterer) ParseTransfer(log types.Log) (*NftApiTransfer, error) {
	event := new(NftApiTransfer)
	if err := _NftApi.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
