build:
	protoc -Iproto --go_out=proto/pb --go_opt=paths=source_relative --go-grpc_out=proto/pb --go-grpc_opt=paths=source_relative proto/*.proto

server-build:
	go build -o bin/cmd/server ./cmd/server

client-build:
	go build -o bin/cmd/client ./cmd/client

prune:
	rm -f proto/pb/*.go