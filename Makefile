test: 
	go test -v -race -cover  ./...
generate-proto:
	protoc -I=./proto --go_out=ola_proto --proto_path=ola/common ola/common/protocol/Ola.proto
	protoc -I=./proto --go_out=ola_proto --proto_path=ola/common ola/common/rpc/Rpc.proto
