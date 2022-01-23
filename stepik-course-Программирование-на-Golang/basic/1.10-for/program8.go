package main

import (
	"strconv"
	"strings"
)

func validNumber(number int) bool {
	if 0 <= number && number <= 10000 {
		return true
	}
	return false
}
func DetermineNumbersAreInBothNumbers(number1, number2 int) []int {
	if !(validNumber(number1) && validNumber(number2)) {
		return []int{}
	}
	numbersKeeperSet := make(map[int]struct{})
	number1S, number2S := strconv.Itoa(number1), strconv.Itoa(number2)
	for _, number := range number1S {
		if strings.ContainsRune(number2S, number) {
			num, _ := strconv.Atoi(string(number))
			numbersKeeperSet[num] = struct{}{}
		}
	}

	res := make([]int, 0, len(numbersKeeperSet))
	for number, _ := range numbersKeeperSet {
		res = append(res, number)
	}
	return res
}
