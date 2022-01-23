package main

import "strconv"

func checkTicket(number int) bool {
	if number >= 100000 && number <= 999999 {
		return true
	}
	return false
}
func LuckyTicket(number int) string {
	if !checkTicket(number) {
		return ""
	}
	numberS := strconv.Itoa(number)
	firstHalfSum := findSum(numberS[:3])
	secondHalfSum := findSum(numberS[3:])

	if firstHalfSum != secondHalfSum {
		return "NO"
	}
	return "YES"
}

func findSum(number string) int {
	res := 0
	for _, numb := range number {
		num, err := strconv.Atoi(string(numb))
		if err != nil {
			return -1
		}
		res += num
	}
	return res
}
