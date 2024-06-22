package serializer

import (
	"fmt"
	"os"

	pb "github.com/prateek69/go-messaging-grpc/proto/proto"
	"google.golang.org/protobuf/proto"
)

func ConvertProtoToBinaryFile(fileName string, employee *pb.Employee) error {
	data, err := proto.Marshal(employee)
	if err != nil {
		return fmt.Errorf("error while marshaling the message, %v", err)
	}

	if err := os.WriteFile(fileName, data, 0644); err != nil {
		return fmt.Errorf("error while saving the binary data to file, %v", err)
	}

	return nil
}

func ConvertBinaryFileToProto(fileName string) (*pb.Employee, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("error while reading the binary file, %v", err)
	}

	employee := &pb.Employee{}
	if err := proto.Unmarshal(data, employee); err != nil {
		return nil, fmt.Errorf("error while unmarshaling the message, %v", err)
	}

	return employee, nil
}
