gen:
	protoc --go_out=proto/ --go_opt=paths=source_relative \
    --go-grpc_out=proto/ --go-grpc_opt=paths=source_relative \
    proto/*.proto

clean:	
	rm proto/proto/*.go

server:
	go run cmd/server/server.go

client:
	go run cmd/client/client.go

test:
	go test -cover ./...