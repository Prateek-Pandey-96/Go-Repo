package serializer

import (
	"testing"

	"github.com/prateek69/go-messaging-grpc/sample"
	"github.com/stretchr/testify/require"
)

func TestConvertProtoToBinaryFile(t *testing.T) {
	t.Parallel()
	employee := sample.GetEmployee()
	filename := "../output/employee"
	err := ConvertProtoToBinaryFile(filename, employee)
	require.NoError(t, err)

	readEmployee, err := ConvertBinaryFileToProto(filename)
	require.NoError(t, err)
	require.Equal(t, employee.EmployeeId, readEmployee.EmployeeId)
}
