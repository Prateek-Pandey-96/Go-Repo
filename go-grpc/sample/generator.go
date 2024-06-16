package sample

import (
	"github.com/google/uuid"
	pb "github.com/prateek69/go-grpc/pb/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func GenerateKeyboard() *pb.Keyboard {
	return &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}
}

func GenerateCPU() *pb.CPU {
	cpu := &pb.CPU{
		Brand:         randomCPUBrand(),
		Name:          randomCPUName(),
		NumberCores:   uint32(randomInt(2, 8)),
		NumberThreads: uint32(randomInt(2, 8)),
		MinGhz:        randomFloat(2.0, 3.5),
		MaxGhz:        randomFloat(3.5, 5.0),
	}
	return cpu
}

func GenerateGPU() *pb.GPU {
	memory := &pb.Memory{
		Value: uint64(randomInt(2, 8)),
		Unit:  pb.Memory_GIGABYTE,
	}

	gpu := &pb.GPU{
		// using same functions here as CPU
		Brand:  randomCPUBrand(),
		Name:   randomCPUName(),
		MinGhz: randomFloat(2.0, 3.5),
		MaxGhz: randomFloat(3.5, 5.0),
		Memory: memory,
	}
	return gpu
}

func GenerateRAM() *pb.Memory {
	ram := &pb.Memory{
		Value: uint64(randomInt(2, 8)),
		Unit:  pb.Memory_GIGABYTE,
	}
	return ram
}

func GenerateSSD() *pb.Storage {
	ssd := &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(2, 8)),
			Unit:  pb.Memory_GIGABYTE,
		},
	}
	return ssd
}

func GenerateHDD() *pb.Storage {
	hdd := &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(2, 8)),
			Unit:  pb.Memory_GIGABYTE,
		},
	}
	return hdd
}

func GenerateScreen() *pb.Screen {
	screen := &pb.Screen{
		SizeInt: float32(randomFloat(13, 17)),
		Resolution: &pb.Screen_Resolution{
			Height: uint32(randomInt(1000, 4200)),
			Width:  uint32(randomInt(2000, 8400)),
		},
		Panel:      randomScreenPanel(),
		Multitouch: randomBool(),
	}
	return screen
}

// Generate a laptop now
func GetLaptop() *pb.Laptop {
	laptop := &pb.Laptop{
		Id:       uuid.New().String(),
		Brand:    "Apple",
		Name:     "Mac Book",
		Cpu:      GenerateCPU(),
		Ram:      GenerateRAM(),
		Gpus:     []*pb.GPU{GenerateGPU()},
		Storages: []*pb.Storage{GenerateHDD(), GenerateSSD()},
		Screen:   GenerateScreen(),
		Weight: &pb.Laptop_WeightKg{
			WeightKg: randomFloat(1.0, 3.0),
		},
		PriceUsd:    randomFloat(1500, 3000),
		ReleaseYear: uint32(randomInt(2015, 2019)),
		UpdatedAt:   timestamppb.Now(),
	}
	return laptop
}
