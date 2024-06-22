package service

import (
	"context"
	"testing"
	"time"

	pb "github.com/prateek69/go-grpc/pb/proto"
	"github.com/prateek69/go-grpc/sample"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateLaptopSuccess(t *testing.T) {
	t.Run("LaptopCreated-IdPresent", func(t *testing.T) {
		t.Parallel()
		laptop := sample.GetLaptop()
		createLaptopRequest := pb.CreateLaptopRequest{
			Laptop: laptop,
		}
		store := NewInMemoryLaptopStore()
		server := NewLaptopServer(store)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		res, err := server.CreateLaptop(ctx, &createLaptopRequest)
		assert.NoError(t, err)
		require.Equal(t, res.Id, laptop.Id)
	})
	t.Run("LaptopCreated-IdAbsent", func(t *testing.T) {
		t.Parallel()
		laptop := sample.GetLaptop()
		laptop.Id = ""
		createLaptopRequest := pb.CreateLaptopRequest{
			Laptop: laptop,
		}
		store := NewInMemoryLaptopStore()
		server := NewLaptopServer(store)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		res, err := server.CreateLaptop(ctx, &createLaptopRequest)
		assert.NoError(t, err)
		require.Equal(t, res.Id, laptop.Id)
	})
}

func TestCreateLaptopFailure(t *testing.T) {
	t.Run("LaptopNotCreated-IdInvalid", func(t *testing.T) {
		t.Parallel()
		laptop := sample.GetLaptop()
		laptop.Id = "invalid-uuid"
		createLaptopRequest := pb.CreateLaptopRequest{
			Laptop: laptop,
		}
		store := NewInMemoryLaptopStore()
		server := NewLaptopServer(store)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		_, err := server.CreateLaptop(ctx, &createLaptopRequest)
		assert.Error(t, err, ErrorInvalidUUID)
	})
	t.Run("LaptopNotCreated-IdAlreadyPresent", func(t *testing.T) {
		t.Parallel()
		laptop := sample.GetLaptop()
		createLaptopRequest := pb.CreateLaptopRequest{
			Laptop: laptop,
		}
		store := NewInMemoryLaptopStore()
		server := NewLaptopServer(store)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		_, err := server.CreateLaptop(ctx, &createLaptopRequest)
		assert.NoError(t, err)

		_, err = server.CreateLaptop(ctx, &createLaptopRequest)
		require.Error(t, err, ErrorLaptopAlreadyExists)
	})
}
