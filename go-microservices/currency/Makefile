.PHONY: protos

protos:
	protoc -I protos/ --go_out=. --go-grpc_out=. --go-grpc_opt=require_unimplemented_servers=false protos/currency.proto