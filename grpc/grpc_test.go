package grpc

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	pb "github.com/ethereum/go-ethereum/grpc/proto/protoeth"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	client pb.RpcApiClient
)

func setup(t *testing.T) {
	ctx := context.TODO()
	conn, err := grpc.DialContext(ctx, "127.0.0.1:2323", grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.Nil(t, err, "meet error")

	client = pb.NewRpcApiClient(conn)
}

func TestGrpcGetBlockNumber(t *testing.T) {
	setup(t)
	ctx := context.TODO()

	reply1, err := client.GetBlockNumber(ctx, &pb.GetBlockNumberReq{})
	assert.Nil(t, err, "GetBlockNumber error:%+v", err)
	assert.Greater(t, reply1.Number, uint64(99999))
}

func TestGrpcGetBalance(t *testing.T) {
	setup(t)
	ctx := context.TODO()

	reply1, err := client.GetBalance(
		ctx, &pb.GetBalanceReq{
			Address: "0x0000000000000000000000000000000000123456",
		},
	)
	assert.Nil(t, err, "GetBalance error:%+v", err)
	fmt.Println(reply1.Balance)
}

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
	arg := pb.NewFilterReq{}
	json.Unmarshal(bdata, &arg)
	t.Errorf("%s", string(bdata))
	v, _ := json.Marshal(arg)
	t.Errorf("%s", string(v))
}
