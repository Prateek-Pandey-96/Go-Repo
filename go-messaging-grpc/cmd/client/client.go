package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	pb "github.com/prateek69/go-messaging-grpc/proto/proto"
	"github.com/prateek69/go-messaging-grpc/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	conn, err := grpc.NewClient(os.Getenv("SERVER_ADDRESS"), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	client := pb.NewEmployeeServiceClient(conn)

	// Type-1 unary rpc
	createEmployees(client, &ctx)
	// Type-2 server streaming
	searchEmployee(client, &ctx)
	// Type-3 client streaming
	uploadResume(client, &ctx)
}

func createEmployees(client pb.EmployeeServiceClient, ctx *context.Context) {
	for i := 0; i < 10; i++ {
		employee := sample.GetEmployee()
		req := &pb.CreateEmployeeRequest{
			Employee: employee,
		}
		res, err := client.CreateEmployee(*ctx, req)
		if err != nil {
			log.Fatalf("failed to create employee: %v", err)
		}
		log.Printf("employee created with id:%s", res.EmployeeId)
	}
}

func searchEmployee(client pb.EmployeeServiceClient, ctx *context.Context) {
	filter := &pb.Filter{
		MinProjects: 3,
	}
	req := &pb.SearchEmployeesRequest{
		Filter: filter,
	}

	stream, err := client.SearchEmployees(*ctx, req)
	if err != nil {
		log.Printf("error while receiving response %v", err)
	}
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("error while receiving response %v", err)
		}
		log.Printf("employee found with name: %s with %d projects under his belt", resp.Employee.Name, len(resp.Employee.Projects))
	}
}

func uploadResume(client pb.EmployeeServiceClient, ctx *context.Context) {
	resumePath := os.Getenv("RESUME_INPUT_PATH") + "resume.png"
	file, err := os.Open(resumePath)
	if err != nil {
		log.Printf("unable to open file %v", err)
	}
	defer file.Close()

	stream, err := client.CreateResume(*ctx)
	if err != nil {
		log.Printf("cannot upload resume %v", err)
	}

	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("cannot read file")
		}

		req := &pb.CreateResumeRequest{
			Resume: &pb.Resume{
				EmployeeId: uuid.NewString(),
				Data:       buffer[:n],
			},
		}
		if err := stream.Send(req); err != nil {
			log.Printf("unable to send file to server %v", err)
		}
	}

	_, err = stream.CloseAndRecv()
	if err == io.EOF {
		log.Printf("resume saved successfully!")
		os.Exit(0)
	}
	if err != nil {
		log.Printf("unable to send file to server %v", err)
		os.Exit(1)
	}
}
