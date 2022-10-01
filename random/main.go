package main

import (
	"fmt"
	"time"
)

func main() {
	Seed()

	for {
		i := RandomIntRange(-100, 100)
		ii := RandomInt(150)
		f := RandomFloatRange(-100, 100)
		ff := RandomFloat(150)
		str := RandomString(10)
		fmt.Printf("i   = %d\nii  = %d\nf   = %.6f\nff  = %.6f\ns   = %s\n\n", i, ii, f, ff, str)

		time.Sleep(1 * time.Second)
	}
}
