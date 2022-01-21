package main

import "strconv"

func digitCounter(number int, res map[int]int) {
	numbers := strconv.Itoa(number)
	for _, num := range numbers {
		n, err := strconv.Atoi(string(num))
		if err != nil {
			return
		}
		res[n]++
	}
}
