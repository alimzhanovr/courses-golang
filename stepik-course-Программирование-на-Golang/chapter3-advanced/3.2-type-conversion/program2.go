package main

import (
	"strconv"
	"unicode"
)

/*
func adding(badString string) int {
	sum := 0
	for _, str := range strings.Fields(badString) {
		normal := ""
		for _, s := range str {
			if unicode.IsDigit(s) {
				normal += string(s)
			}
		}
		num, _ := strconv.Atoi(normal)
		sum += num
	}
	return sum
}*/

func adding(str1, str2 string) string {
	strs := []string{str1, str2}
	for idx, str := range strs {
		normal := ""
		for _, s := range str {
			if unicode.IsDigit(s) {
				normal += string(s)
			}
		}
		strs[idx] = normal
	}

	sum := func(strs []string) int {
		sum := 0
		for _, str := range strs {
			num, _ := strconv.Atoi(str)
			sum += num
		}
		return sum
	}
	return strconv.Itoa(sum(strs))
}
