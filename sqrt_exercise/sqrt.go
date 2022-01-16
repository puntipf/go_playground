package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z, z_prev := float64(1), float64(0)

	for (z_prev - z)*(z_prev - z) > 0.00001 {
		z_prev = z
		z -= (z*z - x) / (2*z)
		fmt.Println(z)
	}
return z
}

func main() {
	fmt.Println(Sqrt(20000000))
}
