package rand

import (
	"math/rand"
	"time"
)

const alphanumeric = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

// Seed initialize the rand package to a deterministic state of the current system clock in nanoseconds.
func Seed() {
	rand.Seed(time.Now().UnixNano())
}

// RandomIntRange returns a random integer n where n >= lower && n < upper.
func RandomIntRange(lower, upper int) int {
	return lower + rand.Intn(upper-lower)
}

// RandomInt returns a random integer n where n >= 0 && n < upper.
func RandomInt(upper int) int {
	return RandomIntRange(0, upper)
}

// RandomFloatRange returns a random float64 n where n >= lower && n < upper.
func RandomFloatRange(lower, upper int) float64 {
	return float64(lower) + rand.Float64()*float64(upper-lower)
}

// RandomFloat returns a random float n where n >= 0 && n < upper.
func RandomFloat(upper int) float64 {
	return RandomFloatRange(0, upper)
}

// RandomString returns a random ASCII String of length n.
func RandomString(n int) string {
	length := len(alphanumeric)

	bytes := make([]byte, n)

	for i := 0; i < n; i++ {
		bytes[i] = alphanumeric[RandomInt(length)]
	}

	return string(bytes)
}
