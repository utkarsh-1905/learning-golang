package randGen

import (
	"math/rand"
	"time"
)

// RandomGenerator To generate random number
func RandomGenerator(n int) int {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(n)
	return randNum
}
