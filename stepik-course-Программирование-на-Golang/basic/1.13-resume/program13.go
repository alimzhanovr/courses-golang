package main

import "fmt"

func fibonacci(before int) (idx int) {
	if before < 0 {
		return -1
	}
	if before == 0 {
		return 0
	}
	prev, curr := 0, 1
	counter := 1
	for {
		if curr > before {
			return -1
		}
		if curr == before {
			return counter
		}
		prev, curr = curr, prev+curr
		counter++
	}
}

func Program13() {
	var num int
	fmt.Scan(&num)
	fmt.Println(fibonacci(num))
}
