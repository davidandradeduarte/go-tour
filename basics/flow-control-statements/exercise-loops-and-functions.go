package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := z; math.Round(z*100)/100 >= math.Round(i*100)/100; {
		i = z
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
		fmt.Println(i)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
