gen:
	protoc --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto

clean:	
	rm pb/proto/*.go

server:
	go run cmd/server/server.go --port 3001

client:
	go run cmd/client/client.go

test:
	go test -cover ./...