package main

import (
	"fmt"
	"github.com/cybersamx/go-recipes/random/rand"
	"time"
)

func main() {
	rand.Seed()

	for {
		i := rand.RandomIntRange(-100, 100)
		ii := rand.RandomInt(150)
		f := rand.RandomFloatRange(-100, 100)
		ff := rand.RandomFloat(150)
		str := rand.RandomString(10)
		fmt.Printf("i   = %d\nii  = %d\nf   = %.6f\nff  = %.6f\ns   = %s\n\n", i, ii, f, ff, str)

		time.Sleep(1 * time.Second)
	}
}
