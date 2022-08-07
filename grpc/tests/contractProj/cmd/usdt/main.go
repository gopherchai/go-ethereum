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
	account0 = `{"address":"8ab3a8a9205fb6c279c973f08b36f989afbd68ad","crypto":{"cipher":"aes-128-ctr","ciphertext":"9470422e84cf766cb0d4041f0b8672e74a64fba936eb2d4519c531db8a003a1e","cipherparams":{"iv":"e26071fab78f650ddee207e413f2f25b"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":4096,"p":6,"r":8,"salt":"998501be68e21cb12e6a7c861723b0b99c95d614962cf7c27cdb986386543b39"},"mac":"35f473c14a119264433722f6a44b4c1b8d044d704ba5245a212ea6cff3b80d1c"},"id":"f8eafc21-755a-4c8a-8c3c-0538e2d3be6d","version":3}`
	address0 = common.HexToAddress(`8ab3a8a9205fb6c279c973f08b36f989afbd68ad`)
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
	log.Println(address1.String())

	tx, err = instance.Transfer(&bind.TransactOpts{
		From: address0,
		//GasTipCap: big.NewInt(52),
		GasFeeCap: big.NewInt(699086302),
		//GasPrice: big.NewInt(100),
	}, address1, big.NewInt(200))

	if err != nil {
		log.Fatalf("%+v", err)
	}
	log.Println(tx.Hash())

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
