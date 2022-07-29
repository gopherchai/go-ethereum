package grpc

import (
	"context"
	"testing"

	pb "github.com/ethereum/go-ethereum/grpc/proto/protoeth"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	client pb.BalanceClient
)

func setup(t *testing.T) {
	ctx := context.TODO()
	conn, err := grpc.DialContext(ctx, "127.0.0.1:2323", grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.Nil(t, err, "meet error")

	client = pb.NewBalanceClient(conn)
}
func TestGrpc(t *testing.T) {

	setup(t)
	ctx := context.TODO()

	reply1, err := client.GetBlockNumber(ctx, &pb.GetBlockNumberReq{})
	assert.Nil(t, err, "GetBlockNumber error:%+v", err)
	assert.Greater(t, reply1.Number, uint64(99999))

}
