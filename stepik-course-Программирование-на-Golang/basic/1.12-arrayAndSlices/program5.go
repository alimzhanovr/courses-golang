package main

import "fmt"

func Program5() {
	var n int
	fmt.Scan(&n)
	if 1 > n && n > 100 {
		return
	}
	numbers := make([]int, n)
	for idx := range numbers {
		fmt.Scan(&numbers[idx])
	}
	fmt.Println(positiveNumbers(numbers))
}
func isPositive(number int) bool {
	return number > 0
}
func positiveNumbers(numbers []int) []int {
	res := make([]int, 0, len(numbers))
	for idx := range numbers {
		if isPositive(numbers[idx]) {
			res = append(res, numbers[idx])
		}
	}
	return res
}
