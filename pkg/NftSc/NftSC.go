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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_tokenURI\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040518060400160405280600881526020016711105c1c0813919560c21b815250604051806040016040528060048152602001630444150560e41b81525081600090816200006191906200018e565b5060016200007082826200018e565b5050506200008d620000876200009360201b60201c565b62000097565b6200025a565b3390565b600780546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b634e487b7160e01b600052604160045260246000fd5b600181811c908216806200011457607f821691505b6020821081036200013557634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200018957600081815260208120601f850160051c81016020861015620001645750805b601f850160051c820191505b81811015620001855782815560010162000170565b5050505b505050565b81516001600160401b03811115620001aa57620001aa620000e9565b620001c281620001bb8454620000ff565b846200013b565b602080601f831160018114620001fa5760008415620001e15750858301515b600019600386901b1c1916600185901b17855562000185565b600085815260208120601f198616915b828110156200022b578886015182559484019460019091019084016200020a565b50858210156200024a5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6117be806200026a6000396000f3fe608060405234801561001057600080fd5b50600436106101215760003560e01c80638da5cb5b116100ad578063c41a360a11610071578063c41a360a1461024e578063c87b56dd14610261578063d0def52114610274578063e985e9c514610287578063f2fde38b146102c357600080fd5b80638da5cb5b1461020557806395d89b41146102165780639f181b5e1461021e578063a22cb46514610228578063b88d4fde1461023b57600080fd5b806323b872dd116100f457806323b872dd146101a357806342842e0e146101b65780636352211e146101c957806370a08231146101dc578063715018a6146101fd57600080fd5b806301ffc9a71461012657806306fdde031461014e578063081812fc14610163578063095ea7b31461018e575b600080fd5b61013961013436600461115f565b6102d6565b60405190151581526020015b60405180910390f35b610156610328565b60405161014591906111cc565b6101766101713660046111df565b6103ba565b6040516001600160a01b039091168152602001610145565b6101a161019c366004611214565b6103e1565b005b6101a16101b136600461123e565b6104fb565b6101a16101c436600461123e565b61052c565b6101766101d73660046111df565b610547565b6101ef6101ea36600461127a565b6105a7565b604051908152602001610145565b6101a161062d565b6007546001600160a01b0316610176565b610156610641565b6008546101ef9081565b6101a1610236366004611295565b610650565b6101a161024936600461135d565b61065f565b61017661025c3660046111df565b610697565b61015661026f3660046111df565b6106a2565b6101ef6102823660046113d9565b6107b2565b61013961029536600461143b565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205460ff1690565b6101a16102d136600461127a565b6107f2565b60006001600160e01b031982166380ac58cd60e01b148061030757506001600160e01b03198216635b5e139f60e01b145b8061032257506301ffc9a760e01b6001600160e01b03198316145b92915050565b6060600080546103379061146e565b80601f01602080910402602001604051908101604052809291908181526020018280546103639061146e565b80156103b05780601f10610385576101008083540402835291602001916103b0565b820191906000526020600020905b81548152906001019060200180831161039357829003601f168201915b5050505050905090565b60006103c58261086b565b506000908152600460205260409020546001600160a01b031690565b60006103ec82610547565b9050806001600160a01b0316836001600160a01b03160361045e5760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656044820152603960f91b60648201526084015b60405180910390fd5b336001600160a01b038216148061047a575061047a8133610295565b6104ec5760405162461bcd60e51b815260206004820152603e60248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60448201527f6b656e206f776e6572206e6f7220617070726f76656420666f7220616c6c00006064820152608401610455565b6104f683836108ca565b505050565b6105053382610938565b6105215760405162461bcd60e51b8152600401610455906114a8565b6104f68383836109b6565b6104f68383836040518060200160405280600081525061065f565b6000818152600260205260408120546001600160a01b0316806103225760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b6044820152606401610455565b60006001600160a01b0382166106115760405162461bcd60e51b815260206004820152602960248201527f4552433732313a2061646472657373207a65726f206973206e6f7420612076616044820152683634b21037bbb732b960b91b6064820152608401610455565b506001600160a01b031660009081526003602052604090205490565b610635610b52565b61063f6000610bac565b565b6060600180546103379061146e565b61065b338383610bfe565b5050565b6106693383610938565b6106855760405162461bcd60e51b8152600401610455906114a8565b61069184848484610ccc565b50505050565b600061032282610547565b60606106ad8261086b565b600082815260066020526040812080546106c69061146e565b80601f01602080910402602001604051908101604052809291908181526020018280546106f29061146e565b801561073f5780601f106107145761010080835404028352916020019161073f565b820191906000526020600020905b81548152906001019060200180831161072257829003601f168201915b50505050509050600061075d60408051602081019091526000815290565b9050805160000361076f575092915050565b8151156107a15780826040516020016107899291906114f6565b60405160208183030381529060405292505050919050565b6107aa84610cff565b949350505050565b60006107bc610b52565b6107ca600880546001019055565b60006107d560085490565b90506107e18482610d72565b6107eb8184610eb4565b9392505050565b6107fa610b52565b6001600160a01b03811661085f5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610455565b61086881610bac565b50565b6000818152600260205260409020546001600160a01b03166108685760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b6044820152606401610455565b600081815260046020526040902080546001600160a01b0319166001600160a01b03841690811790915581906108ff82610547565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b60008061094483610547565b9050806001600160a01b0316846001600160a01b0316148061098b57506001600160a01b0380821660009081526005602090815260408083209388168352929052205460ff165b806107aa5750836001600160a01b03166109a4846103ba565b6001600160a01b031614949350505050565b826001600160a01b03166109c982610547565b6001600160a01b031614610a2d5760405162461bcd60e51b815260206004820152602560248201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060448201526437bbb732b960d91b6064820152608401610455565b6001600160a01b038216610a8f5760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b6064820152608401610455565b610a9a6000826108ca565b6001600160a01b0383166000908152600360205260408120805460019290610ac390849061153b565b90915550506001600160a01b0382166000908152600360205260408120805460019290610af190849061154e565b909155505060008181526002602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b6007546001600160a01b0316331461063f5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610455565b600780546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b816001600160a01b0316836001600160a01b031603610c5f5760405162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c6572000000000000006044820152606401610455565b6001600160a01b03838116600081815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b610cd78484846109b6565b610ce384848484610f47565b6106915760405162461bcd60e51b815260040161045590611561565b6060610d0a8261086b565b6000610d2160408051602081019091526000815290565b90506000815111610d4157604051806020016040528060008152506107eb565b80610d4b84611048565b604051602001610d5c9291906114f6565b6040516020818303038152906040529392505050565b6001600160a01b038216610dc85760405162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f20616464726573736044820152606401610455565b6000818152600260205260409020546001600160a01b031615610e2d5760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e746564000000006044820152606401610455565b6001600160a01b0382166000908152600360205260408120805460019290610e5690849061154e565b909155505060008181526002602052604080822080546001600160a01b0319166001600160a01b03861690811790915590518392907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b6000828152600260205260409020546001600160a01b0316610f2f5760405162461bcd60e51b815260206004820152602e60248201527f45524337323155524953746f726167653a2055524920736574206f66206e6f6e60448201526d32bc34b9ba32b73a103a37b5b2b760911b6064820152608401610455565b60008281526006602052604090206104f68282611601565b60006001600160a01b0384163b1561103d57604051630a85bd0160e11b81526001600160a01b0385169063150b7a0290610f8b9033908990889088906004016116c1565b6020604051808303816000875af1925050508015610fc6575060408051601f3d908101601f19168201909252610fc3918101906116fe565b60015b611023573d808015610ff4576040519150601f19603f3d011682016040523d82523d6000602084013e610ff9565b606091505b50805160000361101b5760405162461bcd60e51b815260040161045590611561565b805181602001fd5b6001600160e01b031916630a85bd0160e11b1490506107aa565b506001949350505050565b60608160000361106f5750506040805180820190915260018152600360fc1b602082015290565b8160005b811561109957806110838161171b565b91506110929050600a8361174a565b9150611073565b60008167ffffffffffffffff8111156110b4576110b46112d1565b6040519080825280601f01601f1916602001820160405280156110de576020820181803683370190505b5090505b84156107aa576110f360018361153b565b9150611100600a8661175e565b61110b90603061154e565b60f81b81838151811061112057611120611772565b60200101906001600160f81b031916908160001a905350611142600a8661174a565b94506110e2565b6001600160e01b03198116811461086857600080fd5b60006020828403121561117157600080fd5b81356107eb81611149565b60005b8381101561119757818101518382015260200161117f565b50506000910152565b600081518084526111b881602086016020860161117c565b601f01601f19169290920160200192915050565b6020815260006107eb60208301846111a0565b6000602082840312156111f157600080fd5b5035919050565b80356001600160a01b038116811461120f57600080fd5b919050565b6000806040838503121561122757600080fd5b611230836111f8565b946020939093013593505050565b60008060006060848603121561125357600080fd5b61125c846111f8565b925061126a602085016111f8565b9150604084013590509250925092565b60006020828403121561128c57600080fd5b6107eb826111f8565b600080604083850312156112a857600080fd5b6112b1836111f8565b9150602083013580151581146112c657600080fd5b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b600067ffffffffffffffff80841115611302576113026112d1565b604051601f8501601f19908116603f0116810190828211818310171561132a5761132a6112d1565b8160405280935085815286868601111561134357600080fd5b858560208301376000602087830101525050509392505050565b6000806000806080858703121561137357600080fd5b61137c856111f8565b935061138a602086016111f8565b925060408501359150606085013567ffffffffffffffff8111156113ad57600080fd5b8501601f810187136113be57600080fd5b6113cd878235602084016112e7565b91505092959194509250565b600080604083850312156113ec57600080fd5b6113f5836111f8565b9150602083013567ffffffffffffffff81111561141157600080fd5b8301601f8101851361142257600080fd5b611431858235602084016112e7565b9150509250929050565b6000806040838503121561144e57600080fd5b611457836111f8565b9150611465602084016111f8565b90509250929050565b600181811c9082168061148257607f821691505b6020821081036114a257634e487b7160e01b600052602260045260246000fd5b50919050565b6020808252602e908201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560408201526d1c881b9bdc88185c1c1c9bdd995960921b606082015260800190565b6000835161150881846020880161117c565b83519083019061151c81836020880161117c565b01949350505050565b634e487b7160e01b600052601160045260246000fd5b8181038181111561032257610322611525565b8082018082111561032257610322611525565b60208082526032908201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b606082015260800190565b601f8211156104f657600081815260208120601f850160051c810160208610156115da5750805b601f850160051c820191505b818110156115f9578281556001016115e6565b505050505050565b815167ffffffffffffffff81111561161b5761161b6112d1565b61162f81611629845461146e565b846115b3565b602080601f831160018114611664576000841561164c5750858301515b600019600386901b1c1916600185901b1785556115f9565b600085815260208120601f198616915b8281101561169357888601518255948401946001909101908401611674565b50858210156116b15787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6001600160a01b03858116825284166020820152604081018390526080606082018190526000906116f4908301846111a0565b9695505050505050565b60006020828403121561171057600080fd5b81516107eb81611149565b60006001820161172d5761172d611525565b5060010190565b634e487b7160e01b600052601260045260246000fd5b60008261175957611759611734565b500490565b60008261176d5761176d611734565b500690565b634e487b7160e01b600052603260045260246000fdfea26469706673582212200e2b1f23b23b944a9011132ecf21f942dc695b6299d469860617c56a7943958264736f6c63430008110033",
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

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NftApi *NftApiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NftApi.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NftApi *NftApiSession) Owner() (common.Address, error) {
	return _NftApi.Contract.Owner(&_NftApi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NftApi *NftApiCallerSession) Owner() (common.Address, error) {
	return _NftApi.Contract.Owner(&_NftApi.CallOpts)
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
// Solidity: function tokenCount() view returns(uint256 _value)
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
// Solidity: function tokenCount() view returns(uint256 _value)
func (_NftApi *NftApiSession) TokenCount() (*big.Int, error) {
	return _NftApi.Contract.TokenCount(&_NftApi.CallOpts)
}

// TokenCount is a free data retrieval call binding the contract method 0x9f181b5e.
//
// Solidity: function tokenCount() view returns(uint256 _value)
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

// Mint is a paid mutator transaction binding the contract method 0xd0def521.
//
// Solidity: function mint(address recipient, string _tokenURI) returns(uint256)
func (_NftApi *NftApiTransactor) Mint(opts *bind.TransactOpts, recipient common.Address, _tokenURI string) (*types.Transaction, error) {
	return _NftApi.contract.Transact(opts, "mint", recipient, _tokenURI)
}

// Mint is a paid mutator transaction binding the contract method 0xd0def521.
//
// Solidity: function mint(address recipient, string _tokenURI) returns(uint256)
func (_NftApi *NftApiSession) Mint(recipient common.Address, _tokenURI string) (*types.Transaction, error) {
	return _NftApi.Contract.Mint(&_NftApi.TransactOpts, recipient, _tokenURI)
}

// Mint is a paid mutator transaction binding the contract method 0xd0def521.
//
// Solidity: function mint(address recipient, string _tokenURI) returns(uint256)
func (_NftApi *NftApiTransactorSession) Mint(recipient common.Address, _tokenURI string) (*types.Transaction, error) {
	return _NftApi.Contract.Mint(&_NftApi.TransactOpts, recipient, _tokenURI)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NftApi *NftApiTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NftApi.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NftApi *NftApiSession) RenounceOwnership() (*types.Transaction, error) {
	return _NftApi.Contract.RenounceOwnership(&_NftApi.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NftApi *NftApiTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NftApi.Contract.RenounceOwnership(&_NftApi.TransactOpts)
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

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NftApi *NftApiTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NftApi.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NftApi *NftApiSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NftApi.Contract.TransferOwnership(&_NftApi.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NftApi *NftApiTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NftApi.Contract.TransferOwnership(&_NftApi.TransactOpts, newOwner)
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

// NftApiOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NftApi contract.
type NftApiOwnershipTransferredIterator struct {
	Event *NftApiOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NftApiOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftApiOwnershipTransferred)
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
		it.Event = new(NftApiOwnershipTransferred)
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
func (it *NftApiOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftApiOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftApiOwnershipTransferred represents a OwnershipTransferred event raised by the NftApi contract.
type NftApiOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NftApi *NftApiFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NftApiOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NftApi.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NftApiOwnershipTransferredIterator{contract: _NftApi.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NftApi *NftApiFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NftApiOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NftApi.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftApiOwnershipTransferred)
				if err := _NftApi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NftApi *NftApiFilterer) ParseOwnershipTransferred(log types.Log) (*NftApiOwnershipTransferred, error) {
	event := new(NftApiOwnershipTransferred)
	if err := _NftApi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
