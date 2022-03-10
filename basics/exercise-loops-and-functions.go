package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	buildInSqrt := math.Sqrt(x)
	fmt.Println("Build in Sqrt Value:", buildInSqrt)
	z := 1.0
	for i := 1; i <= 10000; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Print(i, z)
		if buildInSqrt == z {
			fmt.Println(" Success!")
			break
		} else {
			fmt.Println()
		}
	}
	return z
}

func main() {
	x := 2.0
	fmt.Println(Sqrt(x))
	fmt.Println(math.Sqrt(x))
}
