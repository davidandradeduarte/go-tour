package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func Sqrt(x float64) (float64, error) {

	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := 1.0
	for i := z; math.Round(z*100)/100 >= math.Round(i*100)/100; {
		i = z
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
		fmt.Println(i)
	}
	return z, nil
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
