package main

import "fmt"

func GetNumbers(skipIfLess, stopIfMore int) []int {
	numbers := make([]int, 0, 10)
	var number int
	for fmt.Scan(&number); number <= stopIfMore; fmt.Scan(&number) {
		if number < skipIfLess {
			continue
		}
		numbers = append(numbers, number)
	}
	return numbers
}
