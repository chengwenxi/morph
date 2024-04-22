// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/morph-l2/bindings/solc"
)

const L2CrossDomainMessengerStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)1017_storage\"},{\"astId\":1003,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)1016_storage\"},{\"astId\":1005,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"_paused\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_bool\"},{\"astId\":1006,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_array(t_uint256)1016_storage\"},{\"astId\":1007,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"_status\",\"offset\":0,\"slot\":\"151\",\"type\":\"t_uint256\"},{\"astId\":1008,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"152\",\"type\":\"t_array(t_uint256)1016_storage\"},{\"astId\":1009,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"xDomainMessageSender\",\"offset\":0,\"slot\":\"201\",\"type\":\"t_address\"},{\"astId\":1010,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"counterpart\",\"offset\":0,\"slot\":\"202\",\"type\":\"t_address\"},{\"astId\":1011,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"feeVault\",\"offset\":0,\"slot\":\"203\",\"type\":\"t_address\"},{\"astId\":1012,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"204\",\"type\":\"t_array(t_uint256)1015_storage\"},{\"astId\":1013,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"messageSendTimestamp\",\"offset\":0,\"slot\":\"250\",\"type\":\"t_mapping(t_bytes32,t_uint256)\"},{\"astId\":1014,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"isL1MessageExecuted\",\"offset\":0,\"slot\":\"251\",\"type\":\"t_mapping(t_bytes32,t_bool)\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)1015_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[46]\",\"numberOfBytes\":\"1472\"},\"t_array(t_uint256)1016_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_array(t_uint256)1017_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_bytes32,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_bool\"},\"t_mapping(t_bytes32,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_uint256\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L2CrossDomainMessengerStorageLayout = new(solc.StorageLayout)

var L2CrossDomainMessengerDeployedBin = "0x6080604052600436106100f2575f3560e01c80638da5cb5b11610087578063c4d66de811610057578063c4d66de8146102c9578063e70fc93b146102e8578063ecc7042814610321578063f2fde38b14610335575f80fd5b80638da5cb5b1461024e5780638ef1332e14610278578063b2267a7b14610297578063bedb86fb146102aa575f80fd5b80635f7b1577116100c25780635f7b1577146101cf5780636e296e45146101e2578063715018a61461020e578063797594b014610222575f80fd5b806302345b50146101055780632a6cccb214610148578063478222c2146101675780635c975abb146101b8575f80fd5b36610101576100ff610354565b005b5f80fd5b348015610110575f80fd5b5061013361011f366004611868565b60fb6020525f908152604090205460ff1681565b60405190151581526020015b60405180910390f35b348015610153575f80fd5b506100ff6101623660046118a7565b6103dc565b348015610172575f80fd5b5060cb546101939073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161013f565b3480156101c3575f80fd5b5060655460ff16610133565b6100ff6101dd3660046118c7565b6104d7565b3480156101ed575f80fd5b5060c9546101939073ffffffffffffffffffffffffffffffffffffffff1681565b348015610219575f80fd5b506100ff610529565b34801561022d575f80fd5b5060ca546101939073ffffffffffffffffffffffffffffffffffffffff1681565b348015610259575f80fd5b5060335473ffffffffffffffffffffffffffffffffffffffff16610193565b348015610283575f80fd5b506100ff610292366004611a39565b61053a565b6100ff6102a5366004611aa6565b6106e5565b3480156102b5575f80fd5b506100ff6102c4366004611b01565b6106ff565b3480156102d4575f80fd5b506100ff6102e33660046118a7565b610720565b3480156102f3575f80fd5b50610313610302366004611868565b60fa6020525f908152604090205481565b60405190815260200161013f565b34801561032c575f80fd5b506103136108fc565b348015610340575f80fd5b506100ff61034f3660046118a7565b610983565b60335473ffffffffffffffffffffffffffffffffffffffff1633146103da576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b565b6103e4610354565b73ffffffffffffffffffffffffffffffffffffffff8116610461576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f6665655661756c742063616e6e6f74206265206164647265737328302900000060448201526064016103d1565b60cb805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f4aadc32827849f797733838c61302f7f56d2b6db28caa175eb3f7f8e5aba25f5905f90a35050565b6104df610a37565b610521868686868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250889250610aa4915050565b505050505050565b610531610354565b6103da5f610d43565b610542610a37565b60ca5473ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffeeeeffffffffffffffffffffffffffffffffeeef330173ffffffffffffffffffffffffffffffffffffffff1614610620576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f43616c6c6572206973206e6f74204c3143726f7373446f6d61696e4d6573736560448201527f6e6765720000000000000000000000000000000000000000000000000000000060648201526084016103d1565b5f61062e8686868686610db9565b80516020918201205f81815260fb90925260409091205490915060ff16156106d8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f4d6573736167652077617320616c7265616479207375636365737366756c6c7960448201527f206578656375746564000000000000000000000000000000000000000000000060648201526084016103d1565b6105218686868585610e55565b6106ed610a37565b6106f984848484610aa4565b50505050565b610707610354565b80156107185761071561110b565b50565b610715611190565b5f54610100900460ff161580801561073e57505f54600160ff909116105b806107575750303b15801561075757505f5460ff166001145b6107e3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016103d1565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561083f575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b73ffffffffffffffffffffffffffffffffffffffff821661088c576040517fecc6fdf000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610896825f6111e7565b80156108f8575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b5f73530000000000000000000000000000000000000173ffffffffffffffffffffffffffffffffffffffff1663b58343bb6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561095a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061097e9190611b20565b905090565b61098b610354565b73ffffffffffffffffffffffffffffffffffffffff8116610a2e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016103d1565b61071581610d43565b60655460ff16156103da576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064016103d1565b610aac611334565b823414610b15576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f6d73672e76616c7565206d69736d61746368000000000000000000000000000060448201526064016103d1565b5f73530000000000000000000000000000000000000190505f8173ffffffffffffffffffffffffffffffffffffffff1663b58343bb6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b77573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b9b9190611b20565b90505f610bab3388888589610db9565b80516020918201205f81815260fa90925260409091205490915015610c2c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f4475706c696361746564206d657373616765000000000000000000000000000060448201526064016103d1565b5f81815260fa602052604090819020429055517f600a2e770000000000000000000000000000000000000000000000000000000081526004810182905273ffffffffffffffffffffffffffffffffffffffff84169063600a2e77906024016020604051808303815f875af1158015610ca6573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610cca9190611b20565b5073ffffffffffffffffffffffffffffffffffffffff87163373ffffffffffffffffffffffffffffffffffffffff167f104371f3b442861a2a7b82a070afbbaab748bb13757bf47769e170e37809ec1e8885888a604051610d2e9493929190611ba2565b60405180910390a35050506106f96001609755565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b60608585858585604051602401610dd4959493929190611bd0565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f8ef1332e00000000000000000000000000000000000000000000000000000000179052905095945050505050565b7fffffffffffffffffffffffffacffffffffffffffffffffffffffffffffffffff73ffffffffffffffffffffffffffffffffffffffff851601610f1a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f466f7262696420746f2063616c6c206c3220746f206c31206d6573736167652060448201527f706173736572000000000000000000000000000000000000000000000000000060648201526084016103d1565b610f23846113ae565b60c95473ffffffffffffffffffffffffffffffffffffffff90811690861603610fa8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f496e76616c6964206d6573736167652073656e6465720000000000000000000060448201526064016103d1565b60c980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff878116919091179091556040515f918616908590611002908690611c1f565b5f6040518083038185875af1925050503d805f811461103c576040519150601f19603f3d011682016040523d82523d5f602084013e611041565b606091505b505060c980547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead179055905080156110d9575f82815260fb602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c91a2610521565b60405182907f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f905f90a2505050505050565b611113610a37565b606580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586111663390565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b61119861142d565b606580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa33611166565b5f54610100900460ff1661127d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016103d1565b611285611499565b61128d611537565b6112956115d5565b60c980547fffffffffffffffffffffffff000000000000000000000000000000000000000090811661dead1790915560ca805473ffffffffffffffffffffffffffffffffffffffff858116919093161790558116156108f85760cb805473ffffffffffffffffffffffffffffffffffffffff83167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790555050565b6002609754036113a0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016103d1565b6002609755565b6001609755565b3073ffffffffffffffffffffffffffffffffffffffff821603610715576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f4d657373656e6765723a20466f7262696420746f2063616c6c2073656c66000060448201526064016103d1565b60655460ff166103da576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f742070617573656400000000000000000000000060448201526064016103d1565b5f54610100900460ff1661152f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016103d1565b6103da611673565b5f54610100900460ff166115cd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016103d1565b6103da611712565b5f54610100900460ff1661166b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016103d1565b6103da6117d2565b5f54610100900460ff16611709576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016103d1565b6103da33610d43565b5f54610100900460ff166117a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016103d1565b606580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b5f54610100900460ff166113a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016103d1565b5f60208284031215611878575f80fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff811681146118a2575f80fd5b919050565b5f602082840312156118b7575f80fd5b6118c08261187f565b9392505050565b5f805f805f8060a087890312156118dc575f80fd5b6118e58761187f565b955060208701359450604087013567ffffffffffffffff80821115611908575f80fd5b818901915089601f83011261191b575f80fd5b813581811115611929575f80fd5b8a602082850101111561193a575f80fd5b602083019650809550505050606087013591506119596080880161187f565b90509295509295509295565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f82601f8301126119a1575f80fd5b813567ffffffffffffffff808211156119bc576119bc611965565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715611a0257611a02611965565b81604052838152866020858801011115611a1a575f80fd5b836020870160208301375f602085830101528094505050505092915050565b5f805f805f60a08688031215611a4d575f80fd5b611a568661187f565b9450611a646020870161187f565b93506040860135925060608601359150608086013567ffffffffffffffff811115611a8d575f80fd5b611a9988828901611992565b9150509295509295909350565b5f805f8060808587031215611ab9575f80fd5b611ac28561187f565b935060208501359250604085013567ffffffffffffffff811115611ae4575f80fd5b611af087828801611992565b949793965093946060013593505050565b5f60208284031215611b11575f80fd5b813580151581146118c0575f80fd5b5f60208284031215611b30575f80fd5b5051919050565b5f5b83811015611b51578181015183820152602001611b39565b50505f910152565b5f8151808452611b70816020860160208601611b37565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b848152836020820152826040820152608060608201525f611bc66080830184611b59565b9695505050505050565b5f73ffffffffffffffffffffffffffffffffffffffff808816835280871660208401525084604083015283606083015260a06080830152611c1460a0830184611b59565b979650505050505050565b5f8251611c30818460208701611b37565b919091019291505056fea164736f6c6343000818000a"

func init() {
	if err := json.Unmarshal([]byte(L2CrossDomainMessengerStorageLayoutJSON), L2CrossDomainMessengerStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2CrossDomainMessenger"] = L2CrossDomainMessengerStorageLayout
	deployedBytecodes["L2CrossDomainMessenger"] = L2CrossDomainMessengerDeployedBin
}
