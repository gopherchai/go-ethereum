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

const account1 = `
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

func main() {
	log.SetFlags(log.Lshortfile)
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Panicln(err)
		return
	}
	key, err := keystore.DecryptKey([]byte(account1), "")
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
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(1000000) // in units
	auth.GasPrice = big.NewInt(100000)
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
	tx, err = instance.Transfer(&bind.TransactOpts{
		GasPrice: big.NewInt(10000),
	}, common.HexToAddress("0x3e09c78573e56fda7168d86bbd0e287b11ea1f00"), big.NewInt(200))
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
