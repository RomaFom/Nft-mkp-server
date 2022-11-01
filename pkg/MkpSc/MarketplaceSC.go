// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mkp_api

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

// MkpApiMetaData contains all meta data concerning the MkpApi contract.
var MkpApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_feePercent\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"Bought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"Offered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_itemId\",\"type\":\"uint256\"}],\"name\":\"buyItem\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeAccount\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feePercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_itemId\",\"type\":\"uint256\"}],\"name\":\"getFinalPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"itemCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"items\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC721\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingPrice\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isSold\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_itemId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"listItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC721\",\"name\":\"_nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"makeItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051610cc4380380610cc483398101604081905261002f9161004e565b6001600081905580546001600160a01b03191633179055608052610067565b60006020828403121561006057600080fd5b5051919050565b608051610c3b6100896000396000818161011301526102400152610c3b6000f3fe60806040526004361061007b5760003560e01c8063883efa671161004e578063883efa6714610135578063bfb231d214610157578063e7fb74c714610204578063fa00afc71461021757600080fd5b80634b5549541461008057806365e17c9d146100b35780636bfb0d01146100eb5780637fd6f15c14610101575b600080fd5b34801561008c57600080fd5b506100a061009b366004610a94565b610237565b6040519081526020015b60405180910390f35b3480156100bf57600080fd5b506001546100d3906001600160a01b031681565b6040516001600160a01b0390911681526020016100aa565b3480156100f757600080fd5b506100a060025481565b34801561010d57600080fd5b506100a07f000000000000000000000000000000000000000000000000000000000000000081565b34801561014157600080fd5b50610155610150366004610aad565b610291565b005b34801561016357600080fd5b506101c3610172366004610a94565b600360208190526000918252604090912080546001820154600283015493830154600484015460059094015492946001600160a01b039283169490939192909190811690600160a01b900460ff1687565b604080519788526001600160a01b0396871660208901528701949094526060860192909252608085015290911660a0830152151560c082015260e0016100aa565b610155610212366004610a94565b610563565b34801561022357600080fd5b50610155610232366004610ae7565b61089d565b600060646102657f000000000000000000000000000000000000000000000000000000000000000082610b32565b6000848152600360205260409020600401546102819190610b45565b61028b9190610b5c565b92915050565b6002600054036102bc5760405162461bcd60e51b81526004016102b390610b7e565b60405180910390fd5b600260008181558381526003602052604090819020600181015492015490516331a9108f60e11b8152600481019190915233916001600160a01b031690636352211e90602401602060405180830381865afa15801561031f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103439190610bb5565b6001600160a01b0316146103a35760405162461bcd60e51b815260206004820152602160248201527f596f75206d757374206f776e2074686520746f6b656e20746f206c69737420696044820152601d60fa1b60648201526084016102b3565b600082815260036020526040902060050154600160a01b900460ff1615156001146104075760405162461bcd60e51b8152602060048201526014602482015273125d195b481a5cc8185b1c9958591e481cdbdb1960621b60448201526064016102b3565b600081116104575760405162461bcd60e51b815260206004820152601c60248201527f5072696365206d7573742062652067726561746572207468616e20300000000060448201526064016102b3565b6000828152600360205260409081902060048082018490556001820154600283015493516323b872dd60e01b81523392810192909252306024830152604482019390935290916001600160a01b0316906323b872dd90606401600060405180830381600087803b1580156104ca57600080fd5b505af11580156104de573d6000803e3d6000fd5b505050506005810180546001600160a81b0319163360ff60a01b198116919091179091556001820154600283015460408051878152602081019290925281018590526001600160a01b03909116907f655a0cf9c8db81512be9a76dc1c5ae5380b8816ce6ad659cd61b715e2999d59a9060600160405180910390a35050600160005550565b6002600054036105855760405162461bcd60e51b81526004016102b390610b7e565b6002600090815581815260036020526040902060050154600160a01b900460ff16156105ea5760405162461bcd60e51b8152602060048201526014602482015273125d195b481a5cc8185b1c9958591e481cdbdb1960621b60448201526064016102b3565b6000811180156105fc57506002548111155b61063e5760405162461bcd60e51b8152602060048201526013602482015272125d195b48191bd95cc81b9bdd08195e1a5cdd606a1b60448201526064016102b3565b600081815260036020526040902060050154336001600160a01b03909116036106a95760405162461bcd60e51b815260206004820152601c60248201527f596f752063616e6e6f742062757920796f7572206f776e206974656d0000000060448201526064016102b3565b60006106b482610237565b6000838152600360205260409020909150348211156107155760405162461bcd60e51b815260206004820152601c60248201527f4e6f7420656e6f756768206d6f6e657920746f20627579206974656d0000000060448201526064016102b3565b600581015460038201546040516001600160a01b039092169181156108fc0291906000818181858888f19350505050158015610755573d6000803e3d6000fd5b5060015460038201546001600160a01b03909116906108fc906107789085610bd9565b6040518115909202916000818181858888f193505050501580156107a0573d6000803e3d6000fd5b50600181015460028201546040516323b872dd60e01b815230600482015233602482015260448101919091526001600160a01b03909116906323b872dd90606401600060405180830381600087803b1580156107fb57600080fd5b505af115801561080f573d6000803e3d6000fd5b50505050600581018054600160a01b60ff60a01b198216179091556004820154600383018190556001830154600284015460408051888152602081019290925281019290925233926001600160a01b03908116929116907f8b4c9c8a607d67b321582dd8461041b1dc2ceeca70c8b7f37f8e02095cf2e76d9060600160405180910390a45050600160005550565b6002600054036108bf5760405162461bcd60e51b81526004016102b390610b7e565b6002600055806109115760405162461bcd60e51b815260206004820152601c60248201527f5072696365206d7573742062652067726561746572207468616e20300000000060448201526064016102b3565b6002805490600061092183610bec565b90915550506040516323b872dd60e01b8152336004820152306024820152604481018390526001600160a01b038416906323b872dd90606401600060405180830381600087803b15801561097457600080fd5b505af1158015610988573d6000803e3d6000fd5b50506040805160e081018252600280548083526001600160a01b0389811660208086018281528688018c815260608089018d815260808a018e81523360a08c01818152600060c08e018181529b81526003808a52908f90209d518e55965160018e018054918c166001600160a01b031990921691909117905594518c8c01559151948b0194909455925160048a015590516005909801805496511515600160a01b026001600160a81b0319909716989095169790971794909417909255925485519081529081018990529384018790528695509350917f655a0cf9c8db81512be9a76dc1c5ae5380b8816ce6ad659cd61b715e2999d59a910160405180910390a3505060016000555050565b600060208284031215610aa657600080fd5b5035919050565b60008060408385031215610ac057600080fd5b50508035926020909101359150565b6001600160a01b0381168114610ae457600080fd5b50565b600080600060608486031215610afc57600080fd5b8335610b0781610acf565b95602085013595506040909401359392505050565b634e487b7160e01b600052601160045260246000fd5b8082018082111561028b5761028b610b1c565b808202811582820484141761028b5761028b610b1c565b600082610b7957634e487b7160e01b600052601260045260246000fd5b500490565b6020808252601f908201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604082015260600190565b600060208284031215610bc757600080fd5b8151610bd281610acf565b9392505050565b8181038181111561028b5761028b610b1c565b600060018201610bfe57610bfe610b1c565b506001019056fea264697066735822122027d28c75fb8018598a1f90c7237babaced757e7fd083582b749a048b265ae38864736f6c63430008110033",
}

// MkpApiABI is the input ABI used to generate the binding from.
// Deprecated: Use MkpApiMetaData.ABI instead.
var MkpApiABI = MkpApiMetaData.ABI

// MkpApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MkpApiMetaData.Bin instead.
var MkpApiBin = MkpApiMetaData.Bin

// DeployMkpApi deploys a new Ethereum contract, binding an instance of MkpApi to it.
func DeployMkpApi(auth *bind.TransactOpts, backend bind.ContractBackend, _feePercent *big.Int) (common.Address, *types.Transaction, *MkpApi, error) {
	parsed, err := MkpApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MkpApiBin), backend, _feePercent)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MkpApi{MkpApiCaller: MkpApiCaller{contract: contract}, MkpApiTransactor: MkpApiTransactor{contract: contract}, MkpApiFilterer: MkpApiFilterer{contract: contract}}, nil
}

// MkpApi is an auto generated Go binding around an Ethereum contract.
type MkpApi struct {
	MkpApiCaller     // Read-only binding to the contract
	MkpApiTransactor // Write-only binding to the contract
	MkpApiFilterer   // Log filterer for contract events
}

// MkpApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type MkpApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MkpApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MkpApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MkpApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MkpApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MkpApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MkpApiSession struct {
	Contract     *MkpApi           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MkpApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MkpApiCallerSession struct {
	Contract *MkpApiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MkpApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MkpApiTransactorSession struct {
	Contract     *MkpApiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MkpApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type MkpApiRaw struct {
	Contract *MkpApi // Generic contract binding to access the raw methods on
}

// MkpApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MkpApiCallerRaw struct {
	Contract *MkpApiCaller // Generic read-only contract binding to access the raw methods on
}

// MkpApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MkpApiTransactorRaw struct {
	Contract *MkpApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMkpApi creates a new instance of MkpApi, bound to a specific deployed contract.
func NewMkpApi(address common.Address, backend bind.ContractBackend) (*MkpApi, error) {
	contract, err := bindMkpApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MkpApi{MkpApiCaller: MkpApiCaller{contract: contract}, MkpApiTransactor: MkpApiTransactor{contract: contract}, MkpApiFilterer: MkpApiFilterer{contract: contract}}, nil
}

// NewMkpApiCaller creates a new read-only instance of MkpApi, bound to a specific deployed contract.
func NewMkpApiCaller(address common.Address, caller bind.ContractCaller) (*MkpApiCaller, error) {
	contract, err := bindMkpApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MkpApiCaller{contract: contract}, nil
}

// NewMkpApiTransactor creates a new write-only instance of MkpApi, bound to a specific deployed contract.
func NewMkpApiTransactor(address common.Address, transactor bind.ContractTransactor) (*MkpApiTransactor, error) {
	contract, err := bindMkpApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MkpApiTransactor{contract: contract}, nil
}

// NewMkpApiFilterer creates a new log filterer instance of MkpApi, bound to a specific deployed contract.
func NewMkpApiFilterer(address common.Address, filterer bind.ContractFilterer) (*MkpApiFilterer, error) {
	contract, err := bindMkpApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MkpApiFilterer{contract: contract}, nil
}

// bindMkpApi binds a generic wrapper to an already deployed contract.
func bindMkpApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MkpApiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MkpApi *MkpApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MkpApi.Contract.MkpApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MkpApi *MkpApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MkpApi.Contract.MkpApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MkpApi *MkpApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MkpApi.Contract.MkpApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MkpApi *MkpApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MkpApi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MkpApi *MkpApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MkpApi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MkpApi *MkpApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MkpApi.Contract.contract.Transact(opts, method, params...)
}

// FeeAccount is a free data retrieval call binding the contract method 0x65e17c9d.
//
// Solidity: function feeAccount() view returns(address)
func (_MkpApi *MkpApiCaller) FeeAccount(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MkpApi.contract.Call(opts, &out, "feeAccount")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeAccount is a free data retrieval call binding the contract method 0x65e17c9d.
//
// Solidity: function feeAccount() view returns(address)
func (_MkpApi *MkpApiSession) FeeAccount() (common.Address, error) {
	return _MkpApi.Contract.FeeAccount(&_MkpApi.CallOpts)
}

// FeeAccount is a free data retrieval call binding the contract method 0x65e17c9d.
//
// Solidity: function feeAccount() view returns(address)
func (_MkpApi *MkpApiCallerSession) FeeAccount() (common.Address, error) {
	return _MkpApi.Contract.FeeAccount(&_MkpApi.CallOpts)
}

// FeePercent is a free data retrieval call binding the contract method 0x7fd6f15c.
//
// Solidity: function feePercent() view returns(uint256)
func (_MkpApi *MkpApiCaller) FeePercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MkpApi.contract.Call(opts, &out, "feePercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeePercent is a free data retrieval call binding the contract method 0x7fd6f15c.
//
// Solidity: function feePercent() view returns(uint256)
func (_MkpApi *MkpApiSession) FeePercent() (*big.Int, error) {
	return _MkpApi.Contract.FeePercent(&_MkpApi.CallOpts)
}

// FeePercent is a free data retrieval call binding the contract method 0x7fd6f15c.
//
// Solidity: function feePercent() view returns(uint256)
func (_MkpApi *MkpApiCallerSession) FeePercent() (*big.Int, error) {
	return _MkpApi.Contract.FeePercent(&_MkpApi.CallOpts)
}

// GetFinalPrice is a free data retrieval call binding the contract method 0x4b554954.
//
// Solidity: function getFinalPrice(uint256 _itemId) view returns(uint256)
func (_MkpApi *MkpApiCaller) GetFinalPrice(opts *bind.CallOpts, _itemId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MkpApi.contract.Call(opts, &out, "getFinalPrice", _itemId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFinalPrice is a free data retrieval call binding the contract method 0x4b554954.
//
// Solidity: function getFinalPrice(uint256 _itemId) view returns(uint256)
func (_MkpApi *MkpApiSession) GetFinalPrice(_itemId *big.Int) (*big.Int, error) {
	return _MkpApi.Contract.GetFinalPrice(&_MkpApi.CallOpts, _itemId)
}

// GetFinalPrice is a free data retrieval call binding the contract method 0x4b554954.
//
// Solidity: function getFinalPrice(uint256 _itemId) view returns(uint256)
func (_MkpApi *MkpApiCallerSession) GetFinalPrice(_itemId *big.Int) (*big.Int, error) {
	return _MkpApi.Contract.GetFinalPrice(&_MkpApi.CallOpts, _itemId)
}

// ItemCount is a free data retrieval call binding the contract method 0x6bfb0d01.
//
// Solidity: function itemCount() view returns(uint256)
func (_MkpApi *MkpApiCaller) ItemCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MkpApi.contract.Call(opts, &out, "itemCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ItemCount is a free data retrieval call binding the contract method 0x6bfb0d01.
//
// Solidity: function itemCount() view returns(uint256)
func (_MkpApi *MkpApiSession) ItemCount() (*big.Int, error) {
	return _MkpApi.Contract.ItemCount(&_MkpApi.CallOpts)
}

// ItemCount is a free data retrieval call binding the contract method 0x6bfb0d01.
//
// Solidity: function itemCount() view returns(uint256)
func (_MkpApi *MkpApiCallerSession) ItemCount() (*big.Int, error) {
	return _MkpApi.Contract.ItemCount(&_MkpApi.CallOpts)
}

// Items is a free data retrieval call binding the contract method 0xbfb231d2.
//
// Solidity: function items(uint256 ) view returns(uint256 itemId, address nft, uint256 tokenId, uint256 price, uint256 listingPrice, address seller, bool isSold)
func (_MkpApi *MkpApiCaller) Items(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ItemId       *big.Int
	Nft          common.Address
	TokenId      *big.Int
	Price        *big.Int
	ListingPrice *big.Int
	Seller       common.Address
	IsSold       bool
}, error) {
	var out []interface{}
	err := _MkpApi.contract.Call(opts, &out, "items", arg0)

	outstruct := new(struct {
		ItemId       *big.Int
		Nft          common.Address
		TokenId      *big.Int
		Price        *big.Int
		ListingPrice *big.Int
		Seller       common.Address
		IsSold       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ItemId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Nft = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Price = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ListingPrice = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Seller = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.IsSold = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// Items is a free data retrieval call binding the contract method 0xbfb231d2.
//
// Solidity: function items(uint256 ) view returns(uint256 itemId, address nft, uint256 tokenId, uint256 price, uint256 listingPrice, address seller, bool isSold)
func (_MkpApi *MkpApiSession) Items(arg0 *big.Int) (struct {
	ItemId       *big.Int
	Nft          common.Address
	TokenId      *big.Int
	Price        *big.Int
	ListingPrice *big.Int
	Seller       common.Address
	IsSold       bool
}, error) {
	return _MkpApi.Contract.Items(&_MkpApi.CallOpts, arg0)
}

// Items is a free data retrieval call binding the contract method 0xbfb231d2.
//
// Solidity: function items(uint256 ) view returns(uint256 itemId, address nft, uint256 tokenId, uint256 price, uint256 listingPrice, address seller, bool isSold)
func (_MkpApi *MkpApiCallerSession) Items(arg0 *big.Int) (struct {
	ItemId       *big.Int
	Nft          common.Address
	TokenId      *big.Int
	Price        *big.Int
	ListingPrice *big.Int
	Seller       common.Address
	IsSold       bool
}, error) {
	return _MkpApi.Contract.Items(&_MkpApi.CallOpts, arg0)
}

// BuyItem is a paid mutator transaction binding the contract method 0xe7fb74c7.
//
// Solidity: function buyItem(uint256 _itemId) payable returns()
func (_MkpApi *MkpApiTransactor) BuyItem(opts *bind.TransactOpts, _itemId *big.Int) (*types.Transaction, error) {
	return _MkpApi.contract.Transact(opts, "buyItem", _itemId)
}

// BuyItem is a paid mutator transaction binding the contract method 0xe7fb74c7.
//
// Solidity: function buyItem(uint256 _itemId) payable returns()
func (_MkpApi *MkpApiSession) BuyItem(_itemId *big.Int) (*types.Transaction, error) {
	return _MkpApi.Contract.BuyItem(&_MkpApi.TransactOpts, _itemId)
}

// BuyItem is a paid mutator transaction binding the contract method 0xe7fb74c7.
//
// Solidity: function buyItem(uint256 _itemId) payable returns()
func (_MkpApi *MkpApiTransactorSession) BuyItem(_itemId *big.Int) (*types.Transaction, error) {
	return _MkpApi.Contract.BuyItem(&_MkpApi.TransactOpts, _itemId)
}

// ListItem is a paid mutator transaction binding the contract method 0x883efa67.
//
// Solidity: function listItem(uint256 _itemId, uint256 _price) returns()
func (_MkpApi *MkpApiTransactor) ListItem(opts *bind.TransactOpts, _itemId *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _MkpApi.contract.Transact(opts, "listItem", _itemId, _price)
}

// ListItem is a paid mutator transaction binding the contract method 0x883efa67.
//
// Solidity: function listItem(uint256 _itemId, uint256 _price) returns()
func (_MkpApi *MkpApiSession) ListItem(_itemId *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _MkpApi.Contract.ListItem(&_MkpApi.TransactOpts, _itemId, _price)
}

// ListItem is a paid mutator transaction binding the contract method 0x883efa67.
//
// Solidity: function listItem(uint256 _itemId, uint256 _price) returns()
func (_MkpApi *MkpApiTransactorSession) ListItem(_itemId *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _MkpApi.Contract.ListItem(&_MkpApi.TransactOpts, _itemId, _price)
}

// MakeItem is a paid mutator transaction binding the contract method 0xfa00afc7.
//
// Solidity: function makeItem(address _nft, uint256 _tokenId, uint256 _price) returns()
func (_MkpApi *MkpApiTransactor) MakeItem(opts *bind.TransactOpts, _nft common.Address, _tokenId *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _MkpApi.contract.Transact(opts, "makeItem", _nft, _tokenId, _price)
}

// MakeItem is a paid mutator transaction binding the contract method 0xfa00afc7.
//
// Solidity: function makeItem(address _nft, uint256 _tokenId, uint256 _price) returns()
func (_MkpApi *MkpApiSession) MakeItem(_nft common.Address, _tokenId *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _MkpApi.Contract.MakeItem(&_MkpApi.TransactOpts, _nft, _tokenId, _price)
}

// MakeItem is a paid mutator transaction binding the contract method 0xfa00afc7.
//
// Solidity: function makeItem(address _nft, uint256 _tokenId, uint256 _price) returns()
func (_MkpApi *MkpApiTransactorSession) MakeItem(_nft common.Address, _tokenId *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _MkpApi.Contract.MakeItem(&_MkpApi.TransactOpts, _nft, _tokenId, _price)
}

// MkpApiBoughtIterator is returned from FilterBought and is used to iterate over the raw logs and unpacked data for Bought events raised by the MkpApi contract.
type MkpApiBoughtIterator struct {
	Event *MkpApiBought // Event containing the contract specifics and raw log

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
func (it *MkpApiBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MkpApiBought)
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
		it.Event = new(MkpApiBought)
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
func (it *MkpApiBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MkpApiBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MkpApiBought represents a Bought event raised by the MkpApi contract.
type MkpApiBought struct {
	ItemId  *big.Int
	Nft     common.Address
	TokenId *big.Int
	Price   *big.Int
	Seller  common.Address
	Buyer   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBought is a free log retrieval operation binding the contract event 0x8b4c9c8a607d67b321582dd8461041b1dc2ceeca70c8b7f37f8e02095cf2e76d.
//
// Solidity: event Bought(uint256 itemId, address indexed nft, uint256 tokenId, uint256 price, address indexed seller, address indexed buyer)
func (_MkpApi *MkpApiFilterer) FilterBought(opts *bind.FilterOpts, nft []common.Address, seller []common.Address, buyer []common.Address) (*MkpApiBoughtIterator, error) {

	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _MkpApi.contract.FilterLogs(opts, "Bought", nftRule, sellerRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return &MkpApiBoughtIterator{contract: _MkpApi.contract, event: "Bought", logs: logs, sub: sub}, nil
}

// WatchBought is a free log subscription operation binding the contract event 0x8b4c9c8a607d67b321582dd8461041b1dc2ceeca70c8b7f37f8e02095cf2e76d.
//
// Solidity: event Bought(uint256 itemId, address indexed nft, uint256 tokenId, uint256 price, address indexed seller, address indexed buyer)
func (_MkpApi *MkpApiFilterer) WatchBought(opts *bind.WatchOpts, sink chan<- *MkpApiBought, nft []common.Address, seller []common.Address, buyer []common.Address) (event.Subscription, error) {

	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _MkpApi.contract.WatchLogs(opts, "Bought", nftRule, sellerRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MkpApiBought)
				if err := _MkpApi.contract.UnpackLog(event, "Bought", log); err != nil {
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

// ParseBought is a log parse operation binding the contract event 0x8b4c9c8a607d67b321582dd8461041b1dc2ceeca70c8b7f37f8e02095cf2e76d.
//
// Solidity: event Bought(uint256 itemId, address indexed nft, uint256 tokenId, uint256 price, address indexed seller, address indexed buyer)
func (_MkpApi *MkpApiFilterer) ParseBought(log types.Log) (*MkpApiBought, error) {
	event := new(MkpApiBought)
	if err := _MkpApi.contract.UnpackLog(event, "Bought", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MkpApiOfferedIterator is returned from FilterOffered and is used to iterate over the raw logs and unpacked data for Offered events raised by the MkpApi contract.
type MkpApiOfferedIterator struct {
	Event *MkpApiOffered // Event containing the contract specifics and raw log

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
func (it *MkpApiOfferedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MkpApiOffered)
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
		it.Event = new(MkpApiOffered)
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
func (it *MkpApiOfferedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MkpApiOfferedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MkpApiOffered represents a Offered event raised by the MkpApi contract.
type MkpApiOffered struct {
	ItemId  *big.Int
	Nft     common.Address
	TokenId *big.Int
	Price   *big.Int
	Seller  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOffered is a free log retrieval operation binding the contract event 0x655a0cf9c8db81512be9a76dc1c5ae5380b8816ce6ad659cd61b715e2999d59a.
//
// Solidity: event Offered(uint256 itemId, address indexed nft, uint256 tokenId, uint256 price, address indexed seller)
func (_MkpApi *MkpApiFilterer) FilterOffered(opts *bind.FilterOpts, nft []common.Address, seller []common.Address) (*MkpApiOfferedIterator, error) {

	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _MkpApi.contract.FilterLogs(opts, "Offered", nftRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &MkpApiOfferedIterator{contract: _MkpApi.contract, event: "Offered", logs: logs, sub: sub}, nil
}

// WatchOffered is a free log subscription operation binding the contract event 0x655a0cf9c8db81512be9a76dc1c5ae5380b8816ce6ad659cd61b715e2999d59a.
//
// Solidity: event Offered(uint256 itemId, address indexed nft, uint256 tokenId, uint256 price, address indexed seller)
func (_MkpApi *MkpApiFilterer) WatchOffered(opts *bind.WatchOpts, sink chan<- *MkpApiOffered, nft []common.Address, seller []common.Address) (event.Subscription, error) {

	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _MkpApi.contract.WatchLogs(opts, "Offered", nftRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MkpApiOffered)
				if err := _MkpApi.contract.UnpackLog(event, "Offered", log); err != nil {
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

// ParseOffered is a log parse operation binding the contract event 0x655a0cf9c8db81512be9a76dc1c5ae5380b8816ce6ad659cd61b715e2999d59a.
//
// Solidity: event Offered(uint256 itemId, address indexed nft, uint256 tokenId, uint256 price, address indexed seller)
func (_MkpApi *MkpApiFilterer) ParseOffered(log types.Log) (*MkpApiOffered, error) {
	event := new(MkpApiOffered)
	if err := _MkpApi.contract.UnpackLog(event, "Offered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
