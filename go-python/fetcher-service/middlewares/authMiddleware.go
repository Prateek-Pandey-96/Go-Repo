package middlewares

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prateek69/fetcher-service/models"
	pb "github.com/prateek69/fetcher-service/pb/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func AuthMiddleware(c *gin.Context) {
	serverAddr := "[::]:50051"
	conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewAuthServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var userReq models.UserReq
	if err := c.BindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "improper user object"})
		return
	}

	req := &pb.VerifyRequest{
		Token: userReq.Token,
	}

	resp, err := client.VerifyToken(ctx, req)
	if err != nil {
		log.Printf("failed to verify token: %v", err)
	}

	if !resp.LoggedIn {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "user not logged in"})
		return
	}

	c.Next()
}
