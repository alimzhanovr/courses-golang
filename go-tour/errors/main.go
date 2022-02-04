package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("x=%f must be more than 0", e)
}
func sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < int(x); i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return -1, ErrNegativeSqrt(x)
	}
	return sqrt(x), nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
