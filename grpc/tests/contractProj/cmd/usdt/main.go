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
	account0 = `{"address":"d961421f8351565be929cb6283c481f83cb5ec00","crypto":{"cipher":"aes-128-ctr","ciphertext":"48faabfa50e07a0309d1c8dedb3231de5865be2a74f1d20fe14a8c35e86e2805","cipherparams":{"iv":"5a0a141042166733db3bfa2746b695e5"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":4096,"p":6,"r":8,"salt":"6d77e7042e51cd1114afd827b09b33457facd2756dcd2112bdf27a665d7c3fa2"},"mac":"fb8d7e36ee32a81c0b871cbcfc0cc641a215e3d48251f39260d67332ef9fbd45"},"id":"40a5cce3-de6a-4920-b8f2-a01efe525c3b","version":3}`
	address0 = common.HexToAddress(`0xd961421f8351565be929cb6283c481f83cb5ec00`)
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
	time.Sleep(time.Second * 2)
	log.Println(address0.String())
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
