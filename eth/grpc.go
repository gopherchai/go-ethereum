package eth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/downloader"
	"github.com/ethereum/go-ethereum/eth/filters"
	"github.com/ethereum/go-ethereum/grpc/proto/protoeth"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/log"
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
	protoeth.UnimplementedRpcApiServer
	protoeth.UnimplementedContractFilterServer
	protoeth.UnimplementedContractTransactorServer
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

func (s *GrpcService) GetTransactionReceipt(ctx context.Context, req *protoeth.GetTransactionReceiptReq) (*protoeth.GetTransactionReceiptResp, error) {

	res, err := s.TransactionAPI.GetTransactionReceipt(ctx, common.HexToHash(req.Hash))
	if err != nil {
		return nil, err
	}
	bdata, err := json.Marshal(res)

	return &protoeth.GetTransactionReceiptResp{
		Map: string(bdata),
	}, err
}

func (s *GrpcService) GetBalance(ctx context.Context, args *protoeth.GetBalanceReq) (*protoeth.GetBalanceResp, error) {
	addr := common.HexToAddress(args.Address)

	amount, err := s.BlockChainAPI.GetBalance(ctx, addr, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber))
	if err != nil {
		return nil, err
	}

	return &protoeth.GetBalanceResp{
		Balance: amount.String(),
	}, nil
}

func (s *GrpcService) GetBlockNumber(ctx context.Context, args *protoeth.GetBlockNumberReq) (*protoeth.GetBlockNumberResp, error) {
	hight := s.BlockChainAPI.BlockNumber()
	return &protoeth.GetBlockNumberResp{
		Number: hight.String(),
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

func Serve(stack *node.Node, e *Ethereum, bkd ethapi.Backend) {
	s := stack.GrpcServer()
	svr := NewGrpcService(stack, e, bkd)

	protoeth.RegisterRpcApiServer(s, svr)
	protoeth.RegisterContractFilterServer(s, svr)

	protoeth.RegisterContractTransactorServer(s, svr.UnimplementedContractTransactorServer)
	return
}

func (s *GrpcService) NewFilter(ctx context.Context, args *protoeth.NewFilterReq) (*protoeth.NewFilterResp, error) {

	bdata, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	req := filters.FilterCriteria{}
	err = json.Unmarshal(bdata, &req)
	if err != nil {
		return nil, err
	}
	id, err := s.FilterAPI.NewFilter(req)
	return &protoeth.NewFilterResp{
		Id: string(id),
	}, err

}

func (s *GrpcService) GetFilterChanges(args *protoeth.GetFilterChangeReq, stream protoeth.RpcApi_GetFilterChangesServer) error {
	//请求加参数，说明要持续多长时间
	var id rpc.ID
	t := time.NewTicker(time.Minute * 4)
	id = rpc.ID(args.Id)

	now := time.Now()
	end := now.Add(time.Minute * time.Duration(args.Timeout))
	defer t.Stop()

	for range t.C {
		if end.Before(time.Now()) {
			break
		}
		result, err := s.FilterAPI.GetFilterChanges(id)
		if err != nil {
			return err
		}
		resp := &protoeth.GetFilterChangeResp{}
		switch v := result.(type) {
		case []*types.Log:
			var logs = make([]*protoeth.Log, 0, 0)
			bdata, _ := json.Marshal(v)
			json.Unmarshal(bdata, &logs)
			resp.Logs = logs

		case []common.Hash:
			var hashes = make([]string, 0, 0)
			bdata, _ := json.Marshal(v)
			json.Unmarshal(bdata, &hashes)
			resp.Hashes = hashes
		}
		err = stream.Send(resp)
		if err != nil {
			log.Warn("send filterChange meet error:%+v", err)
		}

	}
	return nil

}

func (s *GrpcService) StartMining(ctx context.Context, args *protoeth.StartMiningReq) (*protoeth.StartMiningResp, error) {
	num := int(args.Num)
	err := s.ethMinerAPI.Start(&num)
	if err != nil {
		return nil, err
	}
	return &protoeth.StartMiningResp{}, err
}

func (s *GrpcService) StopMining(ctx context.Context, args *protoeth.StopMiningReq) (*protoeth.StopMiningResp, error) {
	s.ethMinerAPI.Stop()
	return &protoeth.StopMiningResp{}, nil
}

func (s *GrpcService) SetEtherbase(ctx context.Context, args *protoeth.SetEtherbaseReq) (*protoeth.SetEtherBaseResp, error) {

	ok := s.ethMinerAPI.SetEtherbase(common.HexToAddress(args.Address))
	if !ok {
		return nil, errors.New("not set")
	}
	return &protoeth.SetEtherBaseResp{}, nil
}

func (s *GrpcService) UnlockAccount(ctx context.Context, args *protoeth.UnlockAccountReq) (*protoeth.UnlockAccountResp, error) {
	addr := common.HexToAddress(args.Address)
	_, err := s.PersonalAccountAPI.UnlockAccount(ctx, addr, args.Password, nil)
	if err != nil {
		return nil, err
	}
	return &protoeth.UnlockAccountResp{}, err
}

func (s *GrpcService) ImportRawKey(ctx context.Context, args *protoeth.ImportRawKeyReq) (*protoeth.ImportRawKeyResp, error) {
	addr, err := s.PersonalAccountAPI.ImportRawKey(args.Key, args.Password)
	if err != nil {
		if err == keystore.ErrAccountAlreadyExists {
			return &protoeth.ImportRawKeyResp{}, nil
		}
		return nil, err
	}

	return &protoeth.ImportRawKeyResp{
		Address: addr.Hex(),
	}, nil
}

func (s *GrpcService) SendTransaction(ctx context.Context, args *protoeth.TransactionReq) (*protoeth.TransactionResp, error) {

	bdata, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	var req ethapi.TransactionArgs
	err = json.Unmarshal(bdata, &req)
	if err != nil {
		return nil, err
	}
	log.Info(string(bdata))
	gasPrice, err := s.EthereumAPI.GasPrice(ctx)
	if err != nil {
		return nil, err
	}
	nonce, err := s.TransactionAPI.GetTransactionCount(ctx, *req.From, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber))
	if err != nil {
		return nil, err
	}
	req.Nonce = nonce
	//req.ChainID = s.BlockChainAPI.ChainId()
	req.GasPrice = gasPrice
	msg := fmt.Sprintf("unmarshal tx rags :%+v", req)
	log.Info(msg)
	//s.BlockChainAPI.ChainId()

	hash, err := s.TransactionAPI.SendTransaction(ctx, req)

	return &protoeth.TransactionResp{
		TxHash: hash.String(),
	}, err
}

func (s *GrpcService) Call(ctx context.Context, args *protoeth.TransactionReq) (*protoeth.CallResp, error) {

	var req = ethapi.TransactionArgs{}
	bdata, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bdata, &req)
	if err != nil {
		return nil, err
	}

	data, err := s.BlockChainAPI.Call(ctx, req, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	if err != nil {
		return nil, err
	}
	return &protoeth.CallResp{
		Data: string(data),
	}, err
}

func (s *GrpcService) CallContract(ctx context.Context, args *protoeth.CallContractReq) (*protoeth.CallContractResp, error) {
	call := args.Call

	nonce, err := s.TransactionAPI.GetTransactionCount(ctx, common.HexToAddress(call.From), rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(args.BlockNumber)))
	if err != nil {
		return nil, err
	}

	acls := make([]*protoeth.AccessList, 0, 0)
	err = Struct2Pb(call.AccessList, &acls)
	if err != nil {
		return nil, err
	}
	s.Call(ctx, &protoeth.TransactionReq{
		From:                 call.From,
		To:                   call.To,
		Gas:                  call.GasFeeCap,
		GasPrice:             call.GasPrice,
		MaxFeePerGas:         "",
		MaxPriorityFeePerGas: "",
		Value:                call.GetValue(),
		Nonce:                nonce.String(),
		Data:                 string(call.Data),
		Input:                string(call.Data),
		ChainId:              "",
		AccessList:           acls,
	})
	return nil, nil
}

func (s *GrpcService) CodeAt(ctx context.Context, args *protoeth.CodeAtReq) (*protoeth.CodeAtResp, error) {
	data, err := s.BlockChainAPI.GetCode(ctx, common.HexToAddress(args.Contract), rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber))
	if err != nil {
		return nil, err
	}
	return &protoeth.CodeAtResp{
		Data: string(data),
	}, nil

}

type ContractTransactor struct {
	s *GrpcService
}

func (ct *ContractTransactor) HeaderByNumber(ctx context.Context, args *protoeth.HeaderByNumbeReq) (*protoeth.HeaderResp, error) {

	res, err := ct.s.ethEthereumAPI.e.APIBackend.HeaderByNumber(ctx, rpc.BlockNumber(args.Number))
	if err != nil {
		return nil, err
	}
	reply := &protoeth.HeaderResp{}
	err = Struct2Pb(res, reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (ct *ContractTransactor) PendingCodeAt(ctx context.Context, args *protoeth.PendingCodeAtReq) (*protoeth.PendingCodeAtResp, error) {
	data, err := ct.s.BlockChainAPI.GetCode(ctx, common.HexToAddress(args.Address), rpc.BlockNumberOrHashWithNumber(-2))
	if err != nil {
		return nil, err
	}
	return &protoeth.PendingCodeAtResp{
		Data: string(data),
	}, nil
}

func (ct *ContractTransactor) PendingNonceAt(ctx context.Context, args *protoeth.PendingNonceAtReq) (*protoeth.PendingNonceAtResp, error) {
	//getTransactionCount pending
	num, err := ct.s.TransactionAPI.GetTransactionCount(ctx, common.HexToAddress(args.Account), rpc.BlockNumberOrHashWithNumber(rpc.PendingBlockNumber))
	if err != nil {
		return nil, err
	}

	return &protoeth.PendingNonceAtResp{
		Nonce: hexutil.MustDecodeUint64(num.String()),
	}, nil
}

func (ct *ContractTransactor) SuggestGasPrice(ctx context.Context, args *protoeth.SuggestGasPriceReq) (*protoeth.SuggestGasPriceResp, error) {
	num, err := ct.s.EthereumAPI.GasPrice(ctx)
	if err != nil {
		return nil, err
	}
	return &protoeth.SuggestGasPriceResp{
		Price: num.ToInt().Uint64(),
	}, nil
}

func (ct *ContractTransactor) SuggestGasTipCap(ctx context.Context, args *protoeth.SuggestGasTipCapReq) (*protoeth.SuggestGasPriceResp, error) {
	num, err := ct.s.EthereumAPI.MaxPriorityFeePerGas(ctx)
	if err != nil {
		return nil, err
	}
	return &protoeth.SuggestGasPriceResp{
		Price: num.ToInt().Uint64(),
	}, nil
}

func (ct *ContractTransactor) EstimateGas(ctx context.Context, args *protoeth.EstimateGasReq) (*protoeth.EstimateGasResp, error) {
	msg := args.CallMsg
	req := ethapi.TransactionArgs{}
	err := Struct2Pb(msg, &req)
	if err != nil {
		return nil, err
	}
	num := rpc.BlockNumberOrHashWithNumber(rpc.PendingBlockNumber)
	gas, err := ct.s.BlockChainAPI.EstimateGas(ctx, req, &num)
	if err != nil {
		return nil, err
	}
	return &protoeth.EstimateGasResp{
		Gas: uint64(gas),
	}, nil
}

func (ct *ContractTransactor) SendRawTransaction(ctx context.Context, args *protoeth.SendRawTransactionReq) (*protoeth.SendRawTransactionResp, error) {

	hash, err := ct.s.TransactionAPI.SendRawTransaction(ctx,
		hexutil.Bytes(args.Data))
	if err != nil {
		return nil, err
	}
	return &protoeth.SendRawTransactionResp{
		Hash: hash.String(),
	}, nil
}

type ContractFilter struct {
	s *GrpcService
}

func (ctf *ContractFilter) FilterLogs(ctx context.Context, args *protoeth.FilterLogsReq) (*protoeth.FilterLogsResp, error) {
	req := args.Query
	var q filters.FilterCriteria
	err := Struct2Pb(req, &q)
	if err != nil {
		return nil, err
	}
	logs, err := ctf.s.FilterAPI.GetLogs(ctx, q)
	if err != nil {
		return nil, err
	}
	res := make([]*protoeth.Log, 0, 0)
	err = Struct2Pb(logs, &res)
	if err != nil {
		return nil, err
	}
	return &protoeth.FilterLogsResp{
		Logs: res,
	}, nil
}

func (ctf *ContractFilter) SubscribeFilterLogs(ctx context.Context, args *protoeth.SubscribeFilterLogsReq) (*protoeth.SubscribeFilterLogsResp, error) {

	var q filters.FilterCriteria
	err := Struct2Pb(args, &q)
	if err != nil {
		return nil, err
	}
	subs, err := ctf.s.FilterAPI.Logs(ctx, q)
	if err != nil {
		return nil, err
	}
	res := make([]*protoeth.Subscription, 0, 0)
	err = Struct2Pb(subs, &res)
	if err != nil {
		return nil, err
	}

	return &protoeth.SubscribeFilterLogsResp{
		Subs: res,
	}, nil
}

func Struct2Pb(in interface{}, out interface{}) error {
	bdata, err := json.Marshal(in)
	if err != nil {
		return err
	}
	return json.Unmarshal(bdata, out)

}
