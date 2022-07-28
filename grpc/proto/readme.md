# command to run

cp $GOPATH/bin/protoc-gen-go-grpc $GOPATH/bin/protoc-gen-go_grpc
protoc --go_grpc_out=. --go_out=$PWD \*.proto
