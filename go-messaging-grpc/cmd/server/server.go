package main

import (
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	pb "github.com/prateek69/go-messaging-grpc/proto/proto"
	"github.com/prateek69/go-messaging-grpc/service"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	server := service.NewEmployeeServer()
	grpcServer := grpc.NewServer()

	pb.RegisterEmployeeServiceServer(grpcServer, server)

	listener, err := net.Listen("tcp", os.Getenv("SERVER_ADDRESS"))
	if err != nil {
		log.Fatalf("error occured while starting tcp listener, %v", err)
	}

	grpcServer.Serve(listener)
}
