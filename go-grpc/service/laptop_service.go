package service

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	pb "github.com/prateek69/go-grpc/pb/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LaptopServer struct {
	Store LaptopStore
	pb.UnimplementedLaptopServiceServer
}

func NewLaptopServer(store *InMemoryLaptopStore) *LaptopServer {
	return &LaptopServer{
		Store: store,
	}
}

func (server *LaptopServer) CreateLaptop(
	ctx context.Context,
	req *pb.CreateLaptopRequest,
) (*pb.CreateLaptopResponse, error) {
	// Get laptop from request
	laptop := req.GetLaptop()
	log.Printf("Create laptop request with id: %s", laptop.Id)

	// Get id and generate if doesnt exist
	if len(laptop.Id) > 0 {
		if err := uuid.Validate(laptop.Id); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop id is not a valid uuid: %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "can't generate a new laptop id: %v", err)
		}
		laptop.Id = id.String()
	}

	// Save laptop
	if err := server.Store.Save(laptop); err != nil {
		code := codes.Internal
		if errors.Is(err, ErrorLaptopAlreadyExists) {
			code = codes.AlreadyExists
		}
		return nil, status.Errorf(code, "cannot save laptop to store: %v", err)
	}

	log.Printf("laptop saved with id: %s", laptop.Id)

	// Return laptop id as response
	res := &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}
	return res, nil
}
