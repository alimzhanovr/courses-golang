package main

import (
	"strconv"
)

func Program4(num int) int {
	if num <= 0 {
		return 0
	}
	numberS := strconv.Itoa(num)
	res := 0
	for idx := range numberS {
		number, _ := strconv.Atoi(string(numberS[idx]))
		square := number * number
		if square > 9 {
			res = res*100 + square
			continue
		}
		res = res*10 + square
	}
	return res
}
