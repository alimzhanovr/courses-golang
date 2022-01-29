package main

import (
	"bufio"
	"fmt"
	//"golang.org/x/text/unicode/bidi"
	"os"
	"strings"
)

const (
	palindrome = "Палиндром"
	no         = "Нет"
)

func reverseString(str string) string {
	if str == "" {
		return ""
	}
	startStr := []rune(str)
	count := len(startStr)
	res := make([]rune, 0, count)
	for i := count - 1; i >= 0; i-- {
		res = append(res, startStr[i])
	}
	return string(res)
}
func isPalindrome(str string) bool {
	str = strings.ToLower(str)
	//reverseStr := bidi.ReverseString(str)
	reverseStr := reverseString(str)
	if reverseStr == str {
		return true
	}
	return false
}

func Program2() {
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	str = strings.Trim(str, "\n")
	if isPalindrome(str) {
		fmt.Println(palindrome)
		return
	}
	fmt.Println(no)
}
