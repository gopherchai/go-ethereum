package node

import (
	"context"
	"net"

	"github.com/ethereum/go-ethereum/common"

	//"github.com/ethereum/go-ethereum/grpc"
	"github.com/ethereum/go-ethereum/grpc/proto/protoeth"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/rpc"
	"google.golang.org/grpc"
)

type BalanceService struct {
	//*eth.Ethereum
	//*eth.MinerAPI
	//*eth.AdminAPI
	*ethapi.EthereumAPI
	*ethapi.BlockChainAPI
	*ethapi.TransactionAPI
	*ethapi.TxPoolAPI
	*ethapi.DebugAPI
	protoeth.UnimplementedBalanceServer
}

func NewService(bkd ethapi.Backend) *BalanceService {

	//cfg := n.Config()
	//bkd := &eth.EthAPIBackend{}
	//cfg.ExtRPCEnabled(), cfg.AllowUnprotectedTxs, e
	lock := new(ethapi.AddrLocker)
	return &BalanceService{
		EthereumAPI:    ethapi.NewEthereumAPI(bkd),
		BlockChainAPI:  ethapi.NewBlockChainAPI(bkd),
		TransactionAPI: ethapi.NewTransactionAPI(bkd, lock),
		TxPoolAPI:      ethapi.NewTxPoolAPI(bkd),
		DebugAPI:       ethapi.NewDebugAPI(bkd),
	}
}

func (blc *BalanceService) GetBlockNumber(ctx context.Context, args *protoeth.GetBlockNumberReq) (*protoeth.GetBlockNumberResp, error) {
	hight := blc.BlockChainAPI.BlockNumber()
	return &protoeth.GetBlockNumberResp{
		Number: uint64(hight),
	}, nil
}

func (blc *BalanceService) GetBalance(ctx context.Context, args *protoeth.GetBalanceReq) (*protoeth.GetBalanceResp, error) {
	addr := common.BytesToAddress([]byte(args.Address))

	amount, err := blc.BlockChainAPI.GetBalance(ctx, addr, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber))
	if err != nil {
		return nil, err
	}

	return &protoeth.GetBalanceResp{
		Balance: amount.ToInt().String(),
	}, nil

}

func Serve(s *grpc.Server, endpoint string, bkd ethapi.Backend) error {
	var err error
	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		return err
	}
	protoeth.RegisterBalanceServer(s, NewService(bkd))

	go s.Serve(lis)

	return err

}
