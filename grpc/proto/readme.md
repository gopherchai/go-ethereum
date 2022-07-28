# command to run

cp /Users/byc/.gvm/pkgsets/go1.18.4/global/bin/protoc-gen-go-grpc /Users/byc/.gvm/pkgsets/go1.18.4/global/bin/protoc-gen-go_grpc
protoc --go_grpc_out=. --go_out=$PWD \*.proto
