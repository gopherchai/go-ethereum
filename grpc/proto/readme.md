# command to run

## generate proto grpc

```go
cp $GOPATH/bin/protoc-gen-go-grpc $GOPATH/bin/protoc-gen-go_grpc

protoc --go-grpc_out=. --go_out=. ./*.proto

```

rm -rf /Users/byc/code/gopherchai/data/go-ethereum/cmd/geth/data/geth
