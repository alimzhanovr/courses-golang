package main

import "strconv"

func FlipNumber(number int) int {
	numberS := strconv.Itoa(number)
	mult := 1
	res := 0
	for _, numS := range numberS {
		num, _ := strconv.Atoi(string(numS))
		res += num * mult
		mult *= 10
	}
	return res
}
