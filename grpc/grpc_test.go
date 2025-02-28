package grpc

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth/filters"
	"github.com/ethereum/go-ethereum/grpc/proto/protoeth"
	pb "github.com/ethereum/go-ethereum/grpc/proto/protoeth"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	usdt "github.com/ethereum/go-ethereum/grpc/tests/contractProj/build/contracts/USDT"
	vote "github.com/ethereum/go-ethereum/grpc/tests/contractProj/build/contracts/vote"
)

var (
	client      pb.RpcApiClient
	account1Key string
	account2Key string
	address1    common.Address
	address2    common.Address
	key1        *keystore.Key
	key2        *keystore.Key
)

func setup(t *testing.T) {
	ctx := context.TODO()
	conn, err := grpc.DialContext(ctx, "127.0.0.1:2323", grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.Nil(t, err, "meet error")
	key1, _ = keystore.DecryptKey([]byte(account1), "")
	address1 = key1.Address
	account1Key = hex.EncodeToString(crypto.FromECDSA(key1.PrivateKey))
	key2, _ = keystore.DecryptKey([]byte(account2), "")
	address2 = key2.Address
	account2Key = hex.EncodeToString(crypto.FromECDSA(key2.PrivateKey))

	client = pb.NewRpcApiClient(conn)
}

func testKeyImport(t *testing.T) {
	ctx := context.TODO()
	_, err := client.ImportRawKey(ctx, &pb.ImportRawKeyReq{Key: account1Key})
	require.Nil(t, err)
	_, err = client.ImportRawKey(ctx, &pb.ImportRawKeyReq{Key: account2Key})

	require.Nil(t, err)
	t.Logf("get addr1:%+v", address1)
}

var err error

func testMiningStop(t *testing.T) {
	_, err = client.StopMining(ctx, &pb.StopMiningReq{})
	require.Nil(t, err)
	time.Sleep(time.Second)
	numReply, err := client.GetBlockNumber(ctx, &pb.GetBlockNumberReq{})

	require.Nil(t, err)
	require.NotNil(t, numReply)
	numReply2, err := client.GetBlockNumber(ctx, &pb.GetBlockNumberReq{})
	require.Nil(t, err)
	require.NotNil(t, numReply2)
	require.Equal(t, numReply.Number, numReply2.Number)
}

func testStartMiningAndSetEtherbase(t *testing.T) {
	_, err = client.SetEtherbase(ctx, &pb.SetEtherbaseReq{
		Address: address1.String(),
	})
	require.Nil(t, err)
	balance, err := client.GetBalance(ctx, &pb.GetBalanceReq{
		Address: address1.String(),
	})
	require.Nil(t, err)
	num, err := client.GetBlockNumber(ctx, &pb.GetBlockNumberReq{})
	require.Nil(t, err)

	_, err = client.StartMining(ctx, &pb.StartMiningReq{
		Num: 1,
	})
	require.Nil(t, err)
	time.Sleep(time.Second * 3)
	num2, err := client.GetBlockNumber(ctx, &pb.GetBlockNumberReq{})
	require.Nil(t, err)
	require.Greater(t, num2.Number, num.Number)
	time.Sleep(time.Second)
	balance2, err := client.GetBalance(ctx, &pb.GetBalanceReq{
		Address: address1.String(),
	})
	require.Nil(t, err)
	b1, _ := hexutil.DecodeBig(balance.Balance)
	b2, _ := hexutil.DecodeBig(balance2.Balance)
	require.Equal(t, 1, b2.Cmp(b1))

}

func TestGrpcMiner(t *testing.T) {
	//startMin,stopMin,getBalancer,getBlockNumbert
	//startMining again

	setup(t)
	testKeyImport(t)
	testMiningStop(t)
	testStartMiningAndSetEtherbase(t)

	//before setEtherbase, we should make sure the mining stoped;

	//TODO unlockAccount,转账
	//检查转账结果
	//部署合约
	//调用合约
	//检查合约调用结果

}

var (
	ctx = context.TODO()
)

func TestTransaction(t *testing.T) {
	setup(t)

	client.UnlockAccount(ctx, &pb.UnlockAccountReq{
		Address: address1.String(),
	})

	gas := hexutil.Uint64(uint64(100000))
	//mfpg := hexutil.Big(*big.NewInt(int64(100)))

	var v hexutil.Big
	err = v.UnmarshalText([]byte("0x1239"))
	require.Nil(t, err)
	//v := hexutil.Big(*big.NewInt(int64(3941913128610986)))

	req := ethapi.TransactionArgs{
		From:                 &address1,
		To:                   &address2,
		Gas:                  &gas,
		GasPrice:             nil,
		MaxFeePerGas:         nil,
		MaxPriorityFeePerGas: nil,
		Value:                &v,
		Nonce:                nil,
		Data:                 nil,
		Input:                nil,
		AccessList:           nil,
		ChainID:              nil,
	}
	bdata, _ := json.Marshal(req)
	arg := pb.TransactionReq{}
	json.Unmarshal(bdata, &arg)
	balance, err := client.GetBalance(ctx, &pb.GetBalanceReq{
		Address: address1.String(),
	})
	balance2, err := client.GetBalance(ctx, &pb.GetBalanceReq{
		Address: address2.String(),
	})
	require.Nil(t, err, grpc.ErrorDesc(err))
	resp, err := client.SendTransaction(ctx, &arg)

	require.Nil(t, err, grpc.ErrorDesc(err)+address1.String()+balance.Balance+balance2.Balance)
	require.NotNil(t, resp)
	time.Sleep(time.Second * 4)
	balance2, err = client.GetBalance(ctx, &pb.GetBalanceReq{
		Address: address2.String(),
	})

	require.Nil(t, err, grpc.ErrorDesc(err))
	b1, _ := hexutil.DecodeBig(balance.Balance)
	b2, _ := hexutil.DecodeBig(balance2.Balance)
	t.Logf("b:%+v,%+v", balance, balance2)
	require.Equal(t, b2.Cmp(b1), 1, "b:%+v,%+v", balance, balance2)

}

func TestDeployContract(t *testing.T) {
	setup(t)
	_, err := client.UnlockAccount(ctx, &pb.UnlockAccountReq{
		Address: address1.String(),
	})
	require.Nil(t, err)
	bin := vote.VoteBin
	voteabi, err := vote.VoteMetaData.GetAbi()
	require.Nil(t, err)
	params := make([][32]byte, 0, 0)
	for i := 0; i < 5; i++ {
		var tmp [32]byte
		str := tmp[:]
		b := []byte(fmt.Sprintf("0x4592d8f8d7b001evote proposal-%d", i))
		copy(str, b)
		params = append(params, tmp)
	}
	input, err := voteabi.Pack("", params)
	require.Nil(t, err)
	input = append(common.FromHex(bin), input...)
	gas := hexutil.Uint64(uint64(3941918))
	//mfpg := hexutil.Big(*big.NewInt(int64(100)))
	v := hexutil.Big(*big.NewInt(int64(394190986)))

	binput := hexutil.Bytes(input)
	req := ethapi.TransactionArgs{
		From:                 &address1,
		To:                   nil,
		Gas:                  &gas,
		GasPrice:             nil,
		MaxFeePerGas:         nil,
		MaxPriorityFeePerGas: nil,
		Value:                &v,
		Nonce:                nil,
		Data:                 nil,
		Input:                &binput,
		AccessList:           nil,
		ChainID:              nil,
	}

	bdata, err := json.Marshal(req)
	assert.Nil(t, err)
	arg := pb.TransactionReq{}
	err = json.Unmarshal(bdata, &arg)
	assert.Nil(t, err)
	resp, err := client.SendTransaction(ctx, &arg)

	assert.Nil(t, err, grpc.ErrorDesc(err))

	assert.NotNil(t, resp)
	t.Errorf("get resp:%s", resp.TxHash)
	time.Sleep(time.Second * 4)

}

//TODO filter,filter logs

func TestNewFilter(t *testing.T) {
	setup(t)

	//transfer(address _to, uint256 _value)
	sig := fmt.Sprintf("%v(%v)", "transfer", strings.Join([]string{"uint256", "address"}, ","))

	topic := common.BytesToHash(crypto.Keccak256([]byte(sig)))
	req := filters.FilterCriteria{
		BlockHash: nil,
		FromBlock: big.NewInt(100),
		ToBlock:   big.NewInt(1000),
		Addresses: []common.Address{address1, address2},
		Topics:    [][]common.Hash{[]common.Hash{topic}},
	}
	bdata, err := json.Marshal(req)
	require.Nil(t, err)
	var args = pb.NewFilterReq{}
	err = json.Unmarshal(bdata, &args)
	require.Nil(t, err)
	resp, err := client.NewFilter(ctx, &args)
	require.Nil(t, err)

	resp2, err := client.GetFilterChanges(ctx, &pb.GetFilterChangeReq{
		Id: resp.Id,
	})

	require.Nil(t, err)
	require.NotNil(t, resp2)
	for {
		msg, err := resp2.Recv()
		require.Nil(t, err)
		require.Nil(t, msg)
	}
}

func Sleep80() {
	time.Sleep(time.Second * 1)
}
func TestUSDTDeployAndCallContract(t *testing.T) {
	setup(t)
	address0Str := `0x8ab3a8a9205fb6c279c973f08b36f989afbd68ad`
	_, err := client.UnlockAccount(ctx, &pb.UnlockAccountReq{
		Address: address0Str,
	})
	address0 := common.HexToAddress(address0Str)
	require.Nil(t, err, grpc.ErrorDesc(err))
	bin := usdt.UsdtBin
	usdtabi, err := usdt.UsdtMetaData.GetAbi()
	require.Nil(t, err)
	params := make([]interface{}, 0, 0)

	params = append(params, big.NewInt(1000000000000000010), "USDT", "usdtsymbol", big.NewInt(4))

	input, err := usdtabi.Pack("", params...)
	require.Nil(t, err)
	input = append(common.FromHex(bin), input...)

	binput := hexutil.Bytes(input)
	t.Log("get input", binput.String())
	gas := hexutil.Uint64(uint64(2604788 * 2))
	//mfpg := hexutil.Big(*big.NewInt(int64(100)))
	//v := hexutil.Big(*big.NewInt(int64(39419099986)))

	req := ethapi.TransactionArgs{
		From:                 &address0,
		To:                   nil,
		Gas:                  &gas,
		GasPrice:             nil,
		MaxFeePerGas:         nil,
		MaxPriorityFeePerGas: nil,
		Value:                nil,
		Nonce:                nil,
		Data:                 &binput,
		//	Input:      &input1,
		AccessList: nil,
		ChainID:    nil,
	}
	arg := pb.TransactionReq{}
	structToPb(t, req, &arg)
	require.Equal(t, req.Data.String(), arg.Data)
	bdata, err := json.Marshal(req)
	assert.Nil(t, err)

	err = json.Unmarshal(bdata, &arg)
	assert.Nil(t, err)
	Sleep80()
	//需要追踪rpc的sendTransaction与grpc的实现的差异
	resp, err := client.SendTransaction(ctx, &arg)

	assert.Nil(t, err, grpc.ErrorDesc(err))

	assert.NotNil(t, resp)
	t.Errorf("get11 ressp:%s", resp.TxHash)
	Sleep80()

	res, err := client.GetTransactionReceipt(ctx, &pb.GetTransactionReceiptReq{
		Hash: resp.TxHash,
	})
	require.Nil(t, err)
	m := new(types.Receipt)

	t.Logf("%s", res.Map)
	err = json.Unmarshal([]byte(res.Map), m)
	require.Nil(t, err)

	contractAddr := m.ContractAddress

	binput, err = usdtabi.Pack("transfer", address1, big.NewInt(1000))
	require.Nil(t, err)

	req = ethapi.TransactionArgs{
		From: &address0,
		To:   &contractAddr,
		Gas:  &gas,
		//Value: (*hexutil.Big)(big.NewInt(3)),
		Input: &binput,
		Data:  &binput,
	}

	arg.Reset()
	structToPb(t, req, &arg)
	Sleep80()
	resp, err = client.SendTransaction(ctx, &arg)
	require.Nil(t, err)
	t.Logf("%s", resp.TxHash)

	var callRes *protoeth.CallResp
	binput, err = usdtabi.Pack("balanceOf", address1)
	require.Nil(t, err)
	req = ethapi.TransactionArgs{
		From:  &address0,
		To:    &contractAddr,
		Gas:   &gas,
		Input: &binput,
	}
	arg.Reset()
	structToPb(t, req, &arg)
	Sleep80()
	t.Logf("%+v", arg)
	callRes, err = client.Call(ctx, &arg)
	assert.Nil(t, callRes)
	require.Nil(t, err, grpc.ErrorDesc(err))

	t.Logf("%s", callRes.Data)
	num := big.NewInt(0)
	err = num.UnmarshalText([]byte(callRes.Data))
	require.Nil(t, err)
	require.Equal(t, 1, num.Cmp(big.NewInt(0)))

	outs, err := usdtabi.Unpack("balanceOf", []byte(callRes.Data))
	require.Nil(t, err)
	out0 := *abi.ConvertType(outs[0], new(*big.Int)).(**big.Int)
	require.Equal(t, 0, out0.Cmp(big.NewInt(1000)))
	//todo 需要再调用balanceof合约
	//time.Sleep(time.Second * 4)
}

func structToPb(t *testing.T, from interface{}, to interface{}) {
	bdata, err := json.Marshal(from)
	assert.Nil(t, err)
	err = json.Unmarshal(bdata, to)
	assert.Nil(t, err)
}

func TestStop(t *testing.T) {
	setup(t)
	resp, err := client.StopMining(ctx, &pb.StopMiningReq{})
	require.Nil(t, err)
	require.NotNil(t, resp)
	time.Sleep(time.Second * 5)
	r, err := client.StartMining(ctx, &pb.StartMiningReq{Num: 1})
	require.Nil(t, err)
	require.NotNil(t, r)
}

func TestGrpcGetBlockNumber(t *testing.T) {
	setup(t)
	ctx := context.TODO()

	reply1, err := client.GetBlockNumber(ctx, &pb.GetBlockNumberReq{})
	assert.Nil(t, err, "GetBlockNumber error:%+v", err)

	assert.Greater(t, reply1.Number, "0x7d")
}

func TestCmp(t *testing.T) {
	a := hexutil.MustDecodeBig("0x11")
	b := hexutil.MustDecodeBig("0x11")
	assert.Equal(t, a, b)
}

const (
	account0 = `{"address":"8ab3a8a9205fb6c279c973f08b36f989afbd68ad","crypto":{"cipher":"aes-128-ctr","ciphertext":"9470422e84cf766cb0d4041f0b8672e74a64fba936eb2d4519c531db8a003a1e","cipherparams":{"iv":"e26071fab78f650ddee207e413f2f25b"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":4096,"p":6,"r":8,"salt":"998501be68e21cb12e6a7c861723b0b99c95d614962cf7c27cdb986386543b39"},"mac":"35f473c14a119264433722f6a44b4c1b8d044d704ba5245a212ea6cff3b80d1c"},"id":"f8eafc21-755a-4c8a-8c3c-0538e2d3be6d","version":3}`
	account1 = `
	{
		"address": "f28bba82b11d654428340e910dd602193354a2b0",
		"crypto": {
			"cipher": "aes-128-ctr",
			"ciphertext": "6715c81619366a65a278df8dc645ee1e37018f7f69587b2d052ff722d4ba4d1e",
			"cipherparams": {
				"iv": "c15f273af10280556d933c80cea6a8b9"
			},
			"kdf": "scrypt",
			"kdfparams": {
				"dklen": 32,
				"n": 4096,
				"p": 6,
				"r": 8,
				"salt": "f836fce9e561a38e83a3cee54331ee9abfac2e63babec17530e7a2f8ef3aac23"
			},
			"mac": "3e83e4964bf2f7622efc7fe42a4803ae44c5da7e2732b976d00276f8c5007582"
		},
		"id": "e18c3d2a-48c3-4bf8-849a-d61223ef0d28",
		"version": 3
	}`
	account2 = `{
		"address": "3e09c78573e56fda7168d86bbd0e287b11ea1f00",
		"crypto": {
			"cipher": "aes-128-ctr",
			"ciphertext": "de3d7fc4b00530408d9fe73e9ed9ca5ce2832f7018405119e79cc39a07d5f518",
			"cipherparams": {
				"iv": "45afac047c02a64b04ab8827dbf225a9"
			},
			"kdf": "scrypt",
			"kdfparams": {
				"dklen": 32,
				"n": 4096,
				"p": 6,
				"r": 8,
				"salt": "7bc295f6914e880261b5e5ae33672b8711622ec6bc6a863a5c106dc341a38d20"
			},
			"mac": "3966468d01fb014ecf23a2cb0283d86fb74307da3b934add4ffbcd3335db2e96"
		},
		"id": "fec29597-2348-4c72-ae32-f8bba6cd4b88",
		"version": 3
	}`
)

func TestMashal(t *testing.T) {
	from := big.NewInt(int64(1000232343223132))
	to := big.NewInt(int64(10002323432231))

	bh := common.BytesToHash([]byte(`0xb7e4861d3e312e77d202e963b96ee3a9d794eb0f33046fb53e35d9ab01b0d70c`))
	addr := common.BytesToAddress([]byte(`0x00192fb10df37c9fb26829eb2cc623cd1bf599e8`))
	q := ethereum.FilterQuery{
		BlockHash: &bh,
		FromBlock: from,
		ToBlock:   to,
		Addresses: []common.Address{addr},
		Topics:    [][]common.Hash{[]common.Hash{bh}, []common.Hash{bh}},
	}
	bdata, _ := json.Marshal(q)
	// arg := pb.NewFilterReq{}
	// json.Unmarshal(bdata, &arg)

	// t.Errorf("%s", string(bdata))
	// v, _ := json.Marshal(arg)
	t.Errorf("%s", string(bdata))
	data := []byte(`{"BlockHash":"0x6437393465623066333330343666623533653335643961623031623064373063","FromBlock":"1000232343223132","ToBlock":"10002323432231","Addresses":["0x3239656232636336323363643162663539396538"],"Topics":[["0x6437393465623066333330343666623533653335643961623031623064373063"],["0x6437393465623066333330343666623533653335643961623031623064373063"]]}`)
	var q2 ethereum.FilterQuery
	err := json.Unmarshal(data, &q2)
	t.Errorf("%+v", q2)
	t.Errorf("%+v", err)
}
