cd  ./server/proto
protoc --proto_path=. --go_out=. *.proto
cd ../../rpcapi
protoc --proto_path=. --go_out=. *.proto
protoc --go-grpc_out=. *.proto
