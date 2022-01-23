package main

import "fmt"

func Program4() {
	var n int
	fmt.Scan(&n)
	if 1 > n && n > 100 {
		return
	}
	numbers := make([]int, n)
	for idx := range numbers {
		fmt.Scan(&numbers[idx])
	}
	fmt.Println(evenNumbers(numbers))
}
func isEven(number int) bool {
	return number%2 == 0
}
func evenNumbers(numbers []int) []int {
	res := make([]int, 0, len(numbers))
	for idx := range numbers {
		if isEven(idx) {
			res = append(res, numbers[idx])
		}
	}
	return res
}
