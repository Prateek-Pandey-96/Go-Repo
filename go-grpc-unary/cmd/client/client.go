package main

import (
	"context"
	"log"
	"time"

	pb "github.com/prateek69/go-grpc/pb/proto"
	"github.com/prateek69/go-grpc/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	serverAddr := "localhost:3001"
	conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewLaptopServiceClient(conn)
	laptop := sample.GetLaptop()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	resp, err := client.CreateLaptop(ctx, req)
	if err != nil {
		log.Fatalf("failed to create laptop: %v", err)
	}

	log.Printf("laptop created with id:%s", resp.Id)
}
