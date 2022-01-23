package main

func DegreesOfTwo(number int) []int {
	deg := 1
	res := make([]int, 0)
	for deg <= number {
		res = append(res, deg)
		deg *= 2
	}
	return res
}
