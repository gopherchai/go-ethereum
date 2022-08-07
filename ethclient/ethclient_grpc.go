package ethclient

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/grpc/proto/protoeth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClient struct {
	api protoeth.RpcApiClient
	cfc protoeth.ContractFilterClient
	ctc protoeth.ContractTransactorClient
	//TODO 添加秘钥，用于对交易签名
}

func NewGrpc(address string) (*GrpcClient, error) {
	cli, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &GrpcClient{
		api: protoeth.NewRpcApiClient(cli),
		cfc: protoeth.NewContractFilterClient(cli),
		ctc: protoeth.NewContractTransactorClient(cli),
	}, nil
}

// CodeAt returns the code of the given account. This is needed to differentiate
// between contract internal errors and the local chain being out of sync.
func (gc *GrpcClient) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	resp, err := gc.api.CodeAt(ctx, &protoeth.CodeAtReq{
		BlockNumber: toBlockNumArg(blockNumber),
	})
	if err != nil {
		return nil, err
	}
	return []byte(resp.Data), nil
}

// CallContract executes an Ethereum contract call with the specified data as the
// input.
func (gc *GrpcClient) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {

	resp, err := gc.api.CallContract(ctx, &protoeth.CallContractReq{
		Call:        &protoeth.CallMsg{},
		BlockNumber: blockNumber.Uint64(),
	})
	if err != nil {
		return nil, err
	}
	return []byte(resp.Data), nil
}

// PendingCodeAt returns the code of the given account in the pending state.
func (gc *GrpcClient) PendingCodeAt(ctx context.Context, contract common.Address) ([]byte, error) {
	panic("not implemented") // TODO: Implement
}

// PendingCallContract executes an Ethereum contract call against the pending state.
func (gc *GrpcClient) PendingCallContract(ctx context.Context, call ethereum.CallMsg) ([]byte, error) {
	panic("not implemented") // TODO: Implement
}

// HeaderByNumber returns a block header from the current canonical chain. If
// number is nil, the latest known header is returned.
func (gc *GrpcClient) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	panic("not implemented") // TODO: Implement
}

// PendingNonceAt retrieves the current pending nonce associated with an account.
func (gc *GrpcClient) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	panic("not implemented") // TODO: Implement
}

// SuggestGasPrice retrieves the currently suggested gas price to allow a timely
// execution of a transaction.
func (gc *GrpcClient) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	panic("not implemented") // TODO: Implement
}

// SuggestGasTipCap retrieves the currently suggested 1559 priority fee to allow
// a timely execution of a transaction.
func (gc *GrpcClient) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	panic("not implemented") // TODO: Implement
}

// EstimateGas tries to estimate the gas needed to execute a specific
// transaction based on the current pending state of the backend blockchain.
// There is no guarantee that this is the true gas limit requirement as other
// transactions may be added or removed by miners, but it should provide a basis
// for setting a reasonable default.
func (gc *GrpcClient) EstimateGas(ctx context.Context, call ethereum.CallMsg) (gas uint64, err error) {
	panic("not implemented") // TODO: Implement
}

// SendTransaction injects the transaction into the pending pool for execution.
func (gc *GrpcClient) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	gc.ctc.SendRawTransaction(ctx, &protoeth.SendRawTransactionReq{})
	panic("not implemented") // TODO: Implement
}

// FilterLogs executes a log filter operation, blocking during execution and
// returning all the results in one batch.
//
// TODO(karalabe): Deprecate when the subscription one can return past data too.
func (gc *GrpcClient) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	panic("not implemented") // TODO: Implement
}

// SubscribeFilterLogs creates a background log filtering operation, returning
// a subscription immediately, which can be used to stream the found events.
func (gc *GrpcClient) SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	panic("not implemented") // TODO: Implement
}

type Sub struct{}

// Unsubscribe cancels the sending of events to the data channel
// and closes the error channel.
func (sub *Sub) Unsubscribe() {
	panic("not implemented") // TODO: Implement
}

// Err returns the subscription error channel. The error channel receives
// a value if there is an issue with the subscription (e.g. the network connection
// delivering the events has been closed). Only one value will ever be sent.
// The error channel is closed by Unsubscribe.
func (sub *Sub) Err() <-chan error {
	panic("not implemented") // TODO: Implement
}
