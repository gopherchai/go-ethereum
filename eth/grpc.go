package eth

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/eth/downloader"
	"github.com/ethereum/go-ethereum/eth/filters"
	"github.com/ethereum/go-ethereum/grpc/proto/protoeth"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/rpc"
)

type GrpcService struct {
	//*Ethereum
	ethEthereumAPI *EthereumAPI
	ethMinerAPI    *MinerAPI
	*downloader.DownloaderAPI
	*filters.FilterAPI
	ethAdminAPI *AdminAPI
	ethDebugAPI *DebugAPI
	*ethapi.NetAPI

	*ethapi.EthereumAPI
	*ethapi.BlockChainAPI
	*ethapi.TransactionAPI
	*ethapi.TxPoolAPI
	*ethapi.DebugAPI
	*ethapi.EthereumAccountAPI
	*ethapi.PersonalAccountAPI
	protoeth.UnimplementedBalanceServer
}

func NewGrpcService(node *node.Node, e *Ethereum, bkd ethapi.Backend) *GrpcService {

	//
	apis := e.APIs()
	var s = &GrpcService{}
	for _, api := range apis {
		switch v := api.Service.(type) {
		case *ethapi.EthereumAPI:
			s.EthereumAPI = v
		case *ethapi.BlockChainAPI:
			s.BlockChainAPI = v
		case *ethapi.TransactionAPI:
			s.TransactionAPI = v
		case *ethapi.TxPoolAPI:
			s.TxPoolAPI = v
		case *ethapi.DebugAPI:
			s.DebugAPI = v
		case *ethapi.EthereumAccountAPI:
			s.EthereumAccountAPI = v
		case *ethapi.PersonalAccountAPI:
			s.PersonalAccountAPI = v
		case *EthereumAPI:
			s.ethEthereumAPI = v
		case *MinerAPI:
			s.ethMinerAPI = v
		case *downloader.DownloaderAPI:
			s.DownloaderAPI = v
		case *filters.FilterAPI:
			s.FilterAPI = v
		case *AdminAPI:
			s.ethAdminAPI = v
		case *DebugAPI:
			s.ethDebugAPI = v
		case *ethapi.NetAPI:
			s.NetAPI = v
		}
	}

	return s

}

func (s *GrpcService) GetBlockNumber(ctx context.Context, args *protoeth.GetBlockNumberReq) (*protoeth.GetBlockNumberResp, error) {
	hight := s.BlockChainAPI.BlockNumber()
	return &protoeth.GetBlockNumberResp{
		Number: uint64(hight),
	}, nil
}

//howto method is an template for add new method for GrpcService
//Before rewrite this method, please define the args , reply
//and the rpc method in the file grpc/proto/eth.pro. After that ,
//run `protoc --go_grpc_out=. --go_out=$PWD \*.proto` in the directory grpc/proto.For `protoc`  please read `https://grpc.io/docs/languages/go/quickstart/`
//to run geth with grpc please add option --grpc ,then the grpc service will be supply at 127.0.0.1:2323
// func (s *GrpcService) howTo(ctx context.Context, args interface{}) (reply interface{}, err error) {
// 	//we need change the args of `GetTransactionByHash` with args
// 	trx, err := s.TransactionAPI.GetTransactionByHash(ctx, common.Hash{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	//translate the result to the format of reply
// 	bdata, err := json.Marshal(trx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = json.Unmarshal(bdata, reply)
// 	return
// }

func (s *GrpcService) GetBalance(ctx context.Context, args *protoeth.GetBalanceReq) (*protoeth.GetBalanceResp, error) {
	addr := common.BytesToAddress([]byte(args.Address))

	amount, err := s.BlockChainAPI.GetBalance(ctx, addr, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber))
	if err != nil {
		return nil, err
	}

	return &protoeth.GetBalanceResp{
		Balance: amount.ToInt().String(),
	}, nil
}

func Serve(stack *node.Node, e *Ethereum, bkd ethapi.Backend) {
	s := stack.GrpcServer()
	protoeth.RegisterBalanceServer(s, NewGrpcService(stack, e, bkd))
	return
}
