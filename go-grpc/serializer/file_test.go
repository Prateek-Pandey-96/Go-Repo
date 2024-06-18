package serializer

import (
	"testing"

	pb "github.com/prateek69/go-grpc/pb/proto"
	"github.com/prateek69/go-grpc/sample"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestWriteProtobuffToBinaryFile(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	laptop1 := sample.GetLaptop()

	err := WriteProtobuffToBinaryFile(binaryFile, laptop1)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = ReadProtobuffToBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))
}

func TestWriteProtobuffToJsonFile(t *testing.T) {
	t.Parallel()

	jsonFile := "../tmp/laptop.json"
	laptop1 := sample.GetLaptop()

	err := WriteProtobuffToJsonFile(jsonFile, laptop1)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = ReadProtobuffToJsonFile(jsonFile, laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))
}
