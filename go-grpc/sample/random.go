package sample

import (
	"math/rand"
	"time"

	pb "github.com/prateek69/go-grpc/pb/proto"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_AZERTY
	default:
		return pb.Keyboard_QWERTY
	}
}
func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomCPUBrand() string {
	if rand.Intn(2) == 1 {
		return "Intel"
	}
	return "Amd"
}

func randomCPUName() string {
	if rand.Intn(2) == 1 {
		return "intel-corei7"
	}
	return "amd-r7"
}

func randomInt(min int, max int) int {
	return max + rand.Intn(max-min+1)
}

func randomFloat(min float64, max float64) float64 {
	return max + rand.Float64()*(max-min)
}

func randomScreenPanel() pb.Screen_Panel {
	switch rand.Intn(3) {
	case 1:
		return pb.Screen_IPS
	case 2:
		return pb.Screen_OLED
	default:
		return pb.Screen_IPS
	}
}
