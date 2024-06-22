package service

import (
	"bytes"
	"fmt"
	"os"
	"sync"

	pb "github.com/prateek69/go-messaging-grpc/proto/proto"
)

type EmployeeStore interface {
	SaveEmployee(employee *pb.Employee) error
	SearchEmployee(filter *pb.Filter, found func(*pb.Employee) error) error
}

type InMemoryStore struct {
	mutex      sync.RWMutex
	empData    map[string]*pb.Employee
	resumeData map[string]string
}

func NewInMemoryStore() *InMemoryStore {
	store := &InMemoryStore{}
	store.empData = make(map[string]*pb.Employee)
	store.resumeData = make(map[string]string)
	return store
}

func (store *InMemoryStore) SaveEmployee(employee *pb.Employee) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	id := employee.EmployeeId
	if _, ok := store.empData[id]; ok {
		return fmt.Errorf("employe id %s already exists", id)
	}

	store.empData[id] = employee
	return nil
}

func (store *InMemoryStore) SearchEmployee(filter *pb.Filter, found func(*pb.Employee) error) error {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	for _, emp := range store.empData {
		if len(emp.Projects) >= int(filter.MinProjects) {
			err := found(emp)
			if err != nil {
				return fmt.Errorf("not able to return employee with employe id %s", emp.EmployeeId)
			}
		}
	}
	return nil
}

func (store *InMemoryStore) CreateResume(employeeId string, data bytes.Buffer) error {
	imgPath := fmt.Sprintf("./%s/%s%s", os.Getenv("RESUME_OUTPUT_PATH"), employeeId, ".png")

	file, err := os.Create(imgPath)
	if err != nil {
		return fmt.Errorf("not able to store resume for employee with employe id %s", employeeId)
	}
	_, err = data.WriteTo(file)
	if err != nil {
		return fmt.Errorf("cannot write resume to file for employee %s", employeeId)
	}

	store.mutex.Lock()
	defer store.mutex.Unlock()
	store.resumeData[employeeId] = imgPath
	return nil
}
