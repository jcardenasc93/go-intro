package helpers

import (
	"errors"
	"math/rand"
	"time"
)

type Car struct {
	Brand string
	Model int16
	Color string
}

func GenerateRandInt(poolSize int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(poolSize)
}

func DivideFloats(x, y float32) (float32, error) {
	var result float32
	// Check zero division error
	if y == 0 {
		return result, errors.New("Zero division is not possible")
	}
	result = x / y
	return result, nil
}
