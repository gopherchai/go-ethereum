package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	usdt "github.com/ethereum/go-ethereum/grpc/tests/contractProj/build/contracts/usdt"
)

var (
	account0 = `{"address":"17a8bdcc966016e68f8df27b85df29fafa0f6b43","crypto":{"cipher":"aes-128-ctr","ciphertext":"667e3a71c59dfe6e30da90efbf6e347ca5c854ed2685733f967e49e251acda99","cipherparams":{"iv":"7810570381616b4fb95ad7539928e820"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":4096,"p":6,"r":8,"salt":"5416614f984c774c08e6fd5b7ed7b2c33ee0d9ed09fe438f04effe305ab7cc94"},"mac":"886f1f3bc19608986e56f3e5293c0d301c3c2f21577e27cf65dbcc5c85c83102"},"id":"80b7af94-0427-44a6-bfb2-825194c7577b","version":3}`
	address0 = common.HexToAddress(`17a8bdcc966016e68f8df27b85df29fafa0f6b43`)
	address2 = common.HexToAddress(`f28bba82b11d654428340e910dd602193354a2b0`)
	address1 = common.HexToAddress(`3e09c78573e56fda7168d86bbd0e287b11ea1f00`)
)

func main() {
	log.SetFlags(log.Lshortfile)
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Panicln(err)
		return
	}
	key, err := keystore.DecryptKey([]byte(account0), "")
	if err != nil {
		log.Panicln(err)
		return
	}
	account1Key := hex.EncodeToString(crypto.FromECDSA(key.PrivateKey))

	privateKey, err := crypto.HexToECDSA(account1Key)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// gasPrice, err := client.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	///auth.Value = big.NewInt(0) // in wei
	//auth.GasLimit = uint64(1000000) // in units
	auth.GasPrice = big.NewInt(875000000)

	log.Printf("%d\n", nonce)

	addr, tx, instance, err := usdt.DeployUsdt(auth, client, big.NewInt(1000000000), "usdt", "usdts", big.NewInt(3))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s,%s\n", addr.Hex(), tx.Hash())
	time.Sleep(time.Second * 5)
	if instance == nil {
		panic("nil")
	}
	time.Sleep(time.Second * 12)
	total, err := instance.BalanceOf(&bind.CallOpts{
		Pending:     false,
		From:        address0,
		BlockNumber: big.NewInt(-1),
	}, address0)
	if err != nil {
		//出现execution reverted报错，需要进一步研究
		//追踪到core/vm/interpreter.go:141行文件中
		//可能是智能合约有问题，也可能是别的问题
		log.Fatalf("%+v", err)
	}
	log.Println("total:", total.Int64())

	tx, err = instance.Transfer(&bind.TransactOpts{
		From: address0,
		//GasTipCap: big.NewInt(52),
		GasFeeCap: big.NewInt(659086302),
		//GasPrice: big.NewInt(100),
	}, address1, big.NewInt(200))
	if err != nil {
		log.Fatalf("%+v", err)
	}
	log.Println(tx.Hash())
	num, _ := client.BlockNumber(context.TODO())
	supply, err := instance.TotalSupply(&bind.CallOpts{
		From:        fromAddress,
		Pending:     false,
		BlockNumber: big.NewInt(int64(num)),
		Context:     nil,
	})
	if err != nil {

		log.Panicln(err, err.Error())
	}

	log.Println(supply)
}
