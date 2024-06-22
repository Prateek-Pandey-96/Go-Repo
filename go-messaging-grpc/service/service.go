package service

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/prateek69/go-messaging-grpc/proto/proto"
)

type EmployeeServer struct {
	store *InMemoryStore
	pb.UnimplementedEmployeeServiceServer
}

func NewEmployeeServer() *EmployeeServer {
	server := &EmployeeServer{}
	server.store = NewInMemoryStore()
	return server
}

func (server *EmployeeServer) CreateEmployee(
	ctx context.Context,
	req *pb.CreateEmployeeRequest) (*pb.EmployeeCreatedReponse, error) {

	if err := server.store.SaveEmployee(req.Employee); err != nil {
		return nil, fmt.Errorf("error occured while saving employee %v", err)
	}

	resp := &pb.EmployeeCreatedReponse{
		EmployeeId: req.Employee.EmployeeId,
	}
	log.Printf("employee created with id %s", resp.EmployeeId)
	return resp, nil
}

func (server *EmployeeServer) SearchEmployees(
	req *pb.SearchEmployeesRequest,
	stream pb.EmployeeService_SearchEmployeesServer) error {
	err := server.store.SearchEmployee(req.Filter, func(employee *pb.Employee) error {
		resp := &pb.SearchEmployeesReponse{
			Employee: employee,
		}

		err := stream.Send(resp)
		if err != nil {
			return err
		}
		log.Printf("employee returned with id: %s", employee.EmployeeId)
		return nil
	})
	if err != nil {
		return fmt.Errorf("error occured while searching for employees %v", err)
	}
	return nil
}

func (server *EmployeeServer) CreateResume(
	stream pb.EmployeeService_CreateResumeServer) error {

	imageData := bytes.Buffer{}
	empId := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error while receiving resume %v", err)
		}

		if empId == "" {
			empId = req.Resume.EmployeeId
		}
		chunk := req.Resume.Data

		if _, err := imageData.Write(chunk); err != nil {
			return fmt.Errorf("error while storing resume %v", err)
		}
	}

	if err := server.store.CreateResume(empId, imageData); err != nil {
		return fmt.Errorf("error occured while storing resume %v", err)
	}
	return nil
}
