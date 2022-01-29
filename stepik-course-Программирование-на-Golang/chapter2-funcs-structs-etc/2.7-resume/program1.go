package main

import "math"

func FindHypotenuse(cathetA, cathetB int) int {
	return int(math.Sqrt(float64(cathetA*cathetA + cathetB*cathetB)))
}
