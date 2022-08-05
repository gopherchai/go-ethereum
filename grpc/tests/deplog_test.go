package tests

import (
	"encoding/hex"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	pb "github.com/ethereum/go-ethereum/grpc/proto/protoeth"
	"google.golang.org/grpc"
)

var (
	client      pb.RpcApiClient
	account1Key string
	account2Key string
	address1    common.Address
	address2    common.Address
	key1        *keystore.Key
	key2        *keystore.Key
)

func setup(t *testing.T) {
	cc, _ := grpc.Dial("127.0.0.1:2323")
	client = pb.NewRpcApiClient(cc)
	//TODO sendRawTransaction client and server
	//makeRawTransaction
	key1, _ = keystore.DecryptKey([]byte(account1), "")
	address1 = key1.Address
	account1Key = hex.EncodeToString(crypto.FromECDSA(key1.PrivateKey))
	key2, _ = keystore.DecryptKey([]byte(account2), "")
	address2 = key2.Address
	account2Key = hex.EncodeToString(crypto.FromECDSA(key2.PrivateKey))
}
func TestContractDeploy(t *testing.T) {
	setup(t)
	// pubkey1 := key1.PrivateKey.PublicKey
	// rpc.Client
	// bind.new
	// bind.NewTransactor()
	// bind.DeployContract()

}

var (
	account1 = `
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
	account2 = `{
		"address": "3e09c78573e56fda7168d86bbd0e287b11ea1f00",
		"crypto": {
			"cipher": "aes-128-ctr",
			"ciphertext": "de3d7fc4b00530408d9fe73e9ed9ca5ce2832f7018405119e79cc39a07d5f518",
			"cipherparams": {
				"iv": "45afac047c02a64b04ab8827dbf225a9"
			},
			"kdf": "scrypt",
			"kdfparams": {
				"dklen": 32,
				"n": 4096,
				"p": 6,
				"r": 8,
				"salt": "7bc295f6914e880261b5e5ae33672b8711622ec6bc6a863a5c106dc341a38d20"
			},
			"mac": "3966468d01fb014ecf23a2cb0283d86fb74307da3b934add4ffbcd3335db2e96"
		},
		"id": "fec29597-2348-4c72-ae32-f8bba6cd4b88",
		"version": 3
	}`
)
