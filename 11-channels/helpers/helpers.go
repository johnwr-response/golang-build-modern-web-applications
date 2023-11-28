package helpers

import (
	"math/rand"
)

func RandomNumber(n int) int {
	// This is no longer needed as of Go 1.20
	// rand.Seed(time.Now().UnixNano())
	value := rand.Intn(n)
	return value
}
