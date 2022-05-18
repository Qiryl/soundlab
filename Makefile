.PNOHY: lint
lint:
	golangci-lint run ./...

.PHONY: user-proto 
user-proto:
	@echo Generating user service proto
	cd proto/user/ && protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. user_messages.proto 
	cd proto/user  && protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. user.proto 
