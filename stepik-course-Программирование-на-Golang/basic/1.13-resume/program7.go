package main

import "fmt"

func Program7() {
	var N, number, count int
	fmt.Scan(&N)
	if N < 0 {
		return
	}
	for i := 0; i < N; i++ {
		fmt.Scan(&number)
		if number == 0 {
			count++
		}
	}
	fmt.Println(count)
}
