package sample

import (
	"math/rand"
	"strconv"

	"github.com/google/uuid"
	pb "github.com/prateek69/go-messaging-grpc/proto/proto"
)

func GenerateProject() *pb.Project {
	project_id := uuid.New().String()
	return &pb.Project{
		ProjectId:    project_id,
		ProjectTitle: "project_title_" + project_id,
	}
}

func GetEmployee() *pb.Employee {
	return &pb.Employee{
		Name:       getRandomName(),
		Age:        25 + getRandomInt(),
		EmployeeId: uuid.New().String(),
		Projects:   getProjects(int(getRandomInt())),
	}
}

func getProjects(n int) []*pb.Project {
	projects := []*pb.Project{}
	for i := 0; i < n; i++ {
		projects = append(projects, GenerateProject())
	}
	return projects
}

func getRandomInt() int32 {
	x := rand.Intn(5)
	return int32(x)
}

func getRandomName() string {
	return "Employee" + strconv.Itoa(rand.Intn(10))
}
