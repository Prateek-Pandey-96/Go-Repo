package main

import (
	"flag"
	"log"
	"net"

	pb "github.com/prateek69/go-grpc/pb/proto"
	"github.com/prateek69/go-grpc/service"
	"google.golang.org/grpc"
)

func main() {
	port := flag.String("port", "3000", "enter server port")
	flag.Parse()
	log.Printf("server started at port: %s", *port)

	laptopStore := service.NewInMemoryLaptopStore()
	laptopServer := service.NewLaptopServer(laptopStore)

	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	tcp_listener, err := net.Listen("tcp", "localhost:"+*port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer.Serve(tcp_listener)
}
