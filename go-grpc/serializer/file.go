package serializer

import (
	"fmt"
	"os"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func WriteProtobuffToBinaryFile(fileName string, message proto.Message) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("error while coverting the proto message to binary, %w", err)
	}

	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		return fmt.Errorf("error while writing data to file, %w", err)
	}

	return nil
}

func ReadProtobuffToBinaryFile(fileName string, message proto.Message) error {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("error while reading the binary file, %w", err)
	}

	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("error while unmarshling the binary data, %w", err)
	}

	return nil
}

func WriteProtobuffToJsonFile(fileName string, message proto.Message) error {
	marshaler := protojson.MarshalOptions{
		Indent:            "  ",
		UseEnumNumbers:    false,
		EmitDefaultValues: true,
	}
	string_data, err := marshaler.Marshal(message)
	if err != nil {
		return fmt.Errorf("error while marsheling the data to json, %w", err)
	}

	err = os.WriteFile(fileName, []byte(string_data), 0644)
	if err != nil {
		return fmt.Errorf("error while writing json data to file, %w", err)
	}

	return nil
}

func ReadProtobuffToJsonFile(fileName string, message proto.Message) error {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("error while reading the json file, %w", err)
	}

	err = protojson.Unmarshal([]byte(data), message)
	if err != nil {
		return fmt.Errorf("error while unmarshling json data, %w", err)
	}

	return nil
}
