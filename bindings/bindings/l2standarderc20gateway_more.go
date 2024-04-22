// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/morph-l2/bindings/solc"
)

const L2StandardERC20GatewayStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L2/gateways/L2StandardERC20Gateway.sol:L2StandardERC20Gateway\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"contracts/L2/gateways/L2StandardERC20Gateway.sol:L2StandardERC20Gateway\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/L2/gateways/L2StandardERC20Gateway.sol:L2StandardERC20Gateway\",\"label\":\"_status\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_uint256\"},{\"astId\":1003,\"contract\":\"contracts/L2/gateways/L2StandardERC20Gateway.sol:L2StandardERC20Gateway\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_array(t_uint256)1015_storage\"},{\"astId\":1004,\"contract\":\"contracts/L2/gateways/L2StandardERC20Gateway.sol:L2StandardERC20Gateway\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_array(t_uint256)1016_storage\"},{\"astId\":1005,\"contract\":\"contracts/L2/gateways/L2StandardERC20Gateway.sol:L2StandardERC20Gateway\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_address\"},{\"astId\":1006,\"contract\":\"contracts/L2/gateways/L2StandardERC20Gateway.sol:L2StandardERC20Gateway\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_array(t_uint256)1015_storage\"},{\"astId\":1007,\"contract\":\"contracts/L2/gateways/L2StandardERC20Gateway.sol:L2StandardERC20Gateway\",\"label\":\"counterpart\",\"offset\":0,\"slot\":\"151\",\"type\":\"t_address\"},{\"astId\":1008,\"contract\":\"contracts/L2/gateways/L2StandardERC20Gateway.sol:L2StandardERC20Gateway\",\"label\":\"router\",\"offset\":0,\"slot\":\"152\",\"type\":\"t_address\"},{\"astId\":1009,\"contract\":\"contracts/L2/gateways/L2StandardERC20Gateway.sol:L2StandardERC20Gateway\",\"label\":\"messenger\",\"offset\":0,\"slot\":\"153\",\"type\":\"t_address\"},{\"astId\":1010,\"contract\":\"contracts/L2/gateways/L2StandardERC20Gateway.sol:L2StandardERC20Gateway\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"154\",\"type\":\"t_array(t_uint256)1014_storage\"},{\"astId\":1011,\"contract\":\"contracts/L2/gateways/L2StandardERC20Gateway.sol:L2StandardERC20Gateway\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"200\",\"type\":\"t_array(t_uint256)1016_storage\"},{\"astId\":1012,\"contract\":\"contracts/L2/gateways/L2StandardERC20Gateway.sol:L2StandardERC20Gateway\",\"label\":\"tokenMapping\",\"offset\":0,\"slot\":\"250\",\"type\":\"t_mapping(t_address,t_address)\"},{\"astId\":1013,\"contract\":\"contracts/L2/gateways/L2StandardERC20Gateway.sol:L2StandardERC20Gateway\",\"label\":\"tokenFactory\",\"offset\":0,\"slot\":\"251\",\"type\":\"t_address\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)1014_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[46]\",\"numberOfBytes\":\"1472\"},\"t_array(t_uint256)1015_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_array(t_uint256)1016_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_mapping(t_address,t_address)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e address)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_address\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L2StandardERC20GatewayStorageLayout = new(solc.StorageLayout)

var L2StandardERC20GatewayDeployedBin = "0x6080604052600436106100d9575f3560e01c80638da5cb5b1161007c578063e77772fe11610057578063e77772fe1461024d578063f2fde38b14610279578063f887ea4014610298578063f8c8765e146102c4575f80fd5b80638da5cb5b146101f1578063a93a4af91461021b578063c676ad291461022e575f80fd5b80636c07ea43116100b75780636c07ea431461018b578063715018a61461019e578063797594b0146101b25780638431f5c1146101de575f80fd5b80633cb747bf146100dd57806354bbd59c14610132578063575361b614610176575b5f80fd5b3480156100e8575f80fd5b506099546101099073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b34801561013d575f80fd5b5061010961014c36600461193e565b73ffffffffffffffffffffffffffffffffffffffff9081165f90815260fa60205260409020541690565b610189610184366004611960565b6102e3565b005b610189610199366004611a02565b61032e565b3480156101a9575f80fd5b5061018961036c565b3480156101bd575f80fd5b506097546101099073ffffffffffffffffffffffffffffffffffffffff1681565b6101896101ec366004611af5565b61037f565b3480156101fc575f80fd5b5060655473ffffffffffffffffffffffffffffffffffffffff16610109565b610189610229366004611bc0565b610982565b348015610239575f80fd5b5061010961024836600461193e565b610994565b348015610258575f80fd5b5060fb546101099073ffffffffffffffffffffffffffffffffffffffff1681565b348015610284575f80fd5b5061018961029336600461193e565b610a32565b3480156102a3575f80fd5b506098546101099073ffffffffffffffffffffffffffffffffffffffff1681565b3480156102cf575f80fd5b506101896102de366004611c03565b610ae9565b61032686868686868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250889250610db6915050565b505050505050565b6103678333845f5b6040519080825280601f01601f191660200182016040528015610360576020820181803683370190505b5085610db6565b505050565b6103746111b9565b61037d5f61123a565b565b60995473ffffffffffffffffffffffffffffffffffffffff16338114610406576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f6f6e6c79206d657373656e6765722063616e2063616c6c00000000000000000060448201526064015b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa15801561044f573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104739190611c5c565b60975473ffffffffffffffffffffffffffffffffffffffff9081169116146104f7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6f6e6c792063616c6c20627920636f756e74657270617274000000000000000060448201526064016103fd565b6104ff6112b0565b3415610567576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f6e6f6e7a65726f206d73672e76616c756500000000000000000000000000000060448201526064016103fd565b73ffffffffffffffffffffffffffffffffffffffff87166105e4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f746f6b656e20616464726573732063616e6e6f7420626520300000000000000060448201526064016103fd565b60fb546040517f61e98ca100000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff89811660248301525f9216906361e98ca190604401602060405180830381865afa158015610658573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061067c9190611c5c565b90508073ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff1614610713576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f6c3220746f6b656e206d69736d6174636800000000000000000000000000000060448201526064016103fd565b505f828060200190518101906107299190611ce3565b935090506060808215610753578480602001905181019061074a9190611d35565b925090506107e8565b73ffffffffffffffffffffffffffffffffffffffff8981165f90815260fa60205260409020548116908b16146107e5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f746f6b656e206d617070696e67206d69736d617463680000000000000000000060448201526064016103fd565b50835b73ffffffffffffffffffffffffffffffffffffffff89163b6108605773ffffffffffffffffffffffffffffffffffffffff8981165f90815260fa6020526040902080547fffffffffffffffffffffffff000000000000000000000000000000000000000016918c16919091179055610860828b611323565b6040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8881166004830152602482018890528a16906340c10f19906044015f604051808303815f87803b1580156108cd575f80fd5b505af11580156108df573d5f803e3d5ffd5b505050506108ed8782611451565b8773ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff168b73ffffffffffffffffffffffffffffffffffffffff167f165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba348a8a8660405161096593929190611dd4565b60405180910390a450505061097960018055565b50505050505050565b61098e8484845f610336565b50505050565b60fb546040517f61e98ca100000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff83811660248301525f9216906361e98ca190604401602060405180830381865afa158015610a08573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a2c9190611c5c565b92915050565b610a3a6111b9565b73ffffffffffffffffffffffffffffffffffffffff8116610add576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016103fd565b610ae68161123a565b50565b5f54610100900460ff1615808015610b0757505f54600160ff909116105b80610b205750303b158015610b2057505f5460ff166001145b610bac576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016103fd565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610c08575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b73ffffffffffffffffffffffffffffffffffffffff8416610c85576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f7a65726f20726f7574657220616464726573730000000000000000000000000060448201526064016103fd565b610c90858585611501565b73ffffffffffffffffffffffffffffffffffffffff8216610d0d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f7a65726f20746f6b656e20666163746f7279000000000000000000000000000060448201526064016103fd565b60fb80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84161790558015610daf575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b610dbe6112b0565b5f8311610e27576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f7769746864726177207a65726f20616d6f756e7400000000000000000000000060448201526064016103fd565b609854339073ffffffffffffffffffffffffffffffffffffffff16819003610e625782806020019051810190610e5d9190611e11565b935090505b73ffffffffffffffffffffffffffffffffffffffff8087165f90815260fa60205260409020541680610ef0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f6e6f20636f72726573706f6e64696e67206c3120746f6b656e0000000000000060448201526064016103fd565b6040517f9dc29fac00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff838116600483015260248201879052881690639dc29fac906044015f604051808303815f87803b158015610f5d575f80fd5b505af1158015610f6f573d5f803e3d5ffd5b505050505f818884898989604051602401610f8f96959493929190611e2d565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f84bd13b00000000000000000000000000000000000000000000000000000000017905260995482517fecc7042800000000000000000000000000000000000000000000000000000000815292519394505f9373ffffffffffffffffffffffffffffffffffffffff9091169263ecc704289260048083019391928290030181865afa158015611073573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110979190611e87565b6099546097546040517fb2267a7b00000000000000000000000000000000000000000000000000000000815292935073ffffffffffffffffffffffffffffffffffffffff9182169263b2267a7b9234926110fc929116905f9088908c90600401611e9e565b5f604051808303818588803b158015611113575f80fd5b505af1158015611125573d5f803e3d5ffd5b50505050508373ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fa9967b740f3fc69dfbf4744b4b1c7dfdb0b1b63f1fa4cf573bcdcb9f3ac687c48b8b8b876040516111a49493929190611e9e565b60405180910390a450505050610daf60018055565b60655473ffffffffffffffffffffffffffffffffffffffff16331461037d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103fd565b6065805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b60026001540361131c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016103fd565b6002600155565b60fb546040517f7bdbcbbf00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff83811660248301525f921690637bdbcbbf906044016020604051808303815f875af1158015611398573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906113bc9190611c5c565b90505f805f858060200190518101906113d59190611ee3565b9250925092508373ffffffffffffffffffffffffffffffffffffffff1663c820f146838584308a6040518663ffffffff1660e01b815260040161141c959493929190611f5b565b5f604051808303815f87803b158015611433575f80fd5b505af1158015611445573d5f803e3d5ffd5b50505050505050505050565b5f815111801561147757505f8273ffffffffffffffffffffffffffffffffffffffff163b115b156114f7576040517f444b281f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83169063444b281f906114ce908490600401611fb7565b5f604051808303815f87803b1580156114e5575f80fd5b505af1158015610326573d5f803e3d5ffd5b5050565b60018055565b73ffffffffffffffffffffffffffffffffffffffff831661157e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f7a65726f20636f756e746572706172742061646472657373000000000000000060448201526064016103fd565b73ffffffffffffffffffffffffffffffffffffffff81166115fb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f7a65726f206d657373656e67657220616464726573730000000000000000000060448201526064016103fd565b6116036116ac565b61160b61174a565b6097805473ffffffffffffffffffffffffffffffffffffffff8086167fffffffffffffffffffffffff000000000000000000000000000000000000000092831617909255609980548484169216919091179055821615610367576098805473ffffffffffffffffffffffffffffffffffffffff84167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116179055505050565b5f54610100900460ff16611742576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016103fd565b61037d6117e8565b5f54610100900460ff166117e0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016103fd565b61037d61187e565b5f54610100900460ff166114fb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016103fd565b5f54610100900460ff16611914576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016103fd565b61037d3361123a565b73ffffffffffffffffffffffffffffffffffffffff81168114610ae6575f80fd5b5f6020828403121561194e575f80fd5b81356119598161191d565b9392505050565b5f805f805f8060a08789031215611975575f80fd5b86356119808161191d565b955060208701356119908161191d565b945060408701359350606087013567ffffffffffffffff808211156119b3575f80fd5b818901915089601f8301126119c6575f80fd5b8135818111156119d4575f80fd5b8a60208285010111156119e5575f80fd5b602083019550809450505050608087013590509295509295509295565b5f805f60608486031215611a14575f80fd5b8335611a1f8161191d565b95602085013595506040909401359392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611aa857611aa8611a34565b604052919050565b5f67ffffffffffffffff821115611ac957611ac9611a34565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b5f805f805f8060c08789031215611b0a575f80fd5b8635611b158161191d565b95506020870135611b258161191d565b94506040870135611b358161191d565b93506060870135611b458161191d565b92506080870135915060a087013567ffffffffffffffff811115611b67575f80fd5b8701601f81018913611b77575f80fd5b8035611b8a611b8582611ab0565b611a61565b8181528a6020838501011115611b9e575f80fd5b816020840160208301375f602083830101528093505050509295509295509295565b5f805f8060808587031215611bd3575f80fd5b8435611bde8161191d565b93506020850135611bee8161191d565b93969395505050506040820135916060013590565b5f805f8060808587031215611c16575f80fd5b8435611c218161191d565b93506020850135611c318161191d565b92506040850135611c418161191d565b91506060850135611c518161191d565b939692955090935050565b5f60208284031215611c6c575f80fd5b81516119598161191d565b5f5b83811015611c91578181015183820152602001611c79565b50505f910152565b5f82601f830112611ca8575f80fd5b8151611cb6611b8582611ab0565b818152846020838601011115611cca575f80fd5b611cdb826020830160208701611c77565b949350505050565b5f8060408385031215611cf4575f80fd5b82518015158114611d03575f80fd5b602084015190925067ffffffffffffffff811115611d1f575f80fd5b611d2b85828601611c99565b9150509250929050565b5f8060408385031215611d46575f80fd5b825167ffffffffffffffff80821115611d5d575f80fd5b611d6986838701611c99565b93506020850151915080821115611d7e575f80fd5b50611d2b85828601611c99565b5f8151808452611da2816020860160208601611c77565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b73ffffffffffffffffffffffffffffffffffffffff84168152826020820152606060408201525f611e086060830184611d8b565b95945050505050565b5f8060408385031215611e22575f80fd5b8251611d038161191d565b5f73ffffffffffffffffffffffffffffffffffffffff80891683528088166020840152808716604084015280861660608401525083608083015260c060a0830152611e7b60c0830184611d8b565b98975050505050505050565b5f60208284031215611e97575f80fd5b5051919050565b73ffffffffffffffffffffffffffffffffffffffff85168152836020820152608060408201525f611ed26080830185611d8b565b905082606083015295945050505050565b5f805f60608486031215611ef5575f80fd5b835167ffffffffffffffff80821115611f0c575f80fd5b611f1887838801611c99565b94506020860151915080821115611f2d575f80fd5b50611f3a86828701611c99565b925050604084015160ff81168114611f50575f80fd5b809150509250925092565b60a081525f611f6d60a0830188611d8b565b8281036020840152611f7f8188611d8b565b60ff969096166040840152505073ffffffffffffffffffffffffffffffffffffffff9283166060820152911660809091015292915050565b602081525f6119596020830184611d8b56fea164736f6c6343000818000a"

func init() {
	if err := json.Unmarshal([]byte(L2StandardERC20GatewayStorageLayoutJSON), L2StandardERC20GatewayStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2StandardERC20Gateway"] = L2StandardERC20GatewayStorageLayout
	deployedBytecodes["L2StandardERC20Gateway"] = L2StandardERC20GatewayDeployedBin
}
