package main

import "fmt"

func Program8() {
	var N, min, count int
	numbers := make([]int, N)
	fmt.Scan(&N)
	if N < 0 {
		return
	}
	for i := 0; i < N; i++ {
		fmt.Scan(&numbers[i])
	}
	min = numbers[0]
	for _, number := range numbers {
		if number < min {
			min = number
			count = 1
			continue
		}
		if number == min {
			count++
			continue
		}
	}
	fmt.Println(count)
}
