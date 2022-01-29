package main

import (
	"fmt"
	"unicode"
)

const (
	ok            = "Ok"
	wrongPassword = "Wrong password"
)

func validPassword(password string) bool {
	if len([]rune(password)) < 5 {
		return false
	}
	for _, letter := range password {
		if !(unicode.IsDigit(letter) || unicode.Is(unicode.Latin, letter)) {
			return false
		}
	}
	return true
}

func Program6(password string) {
	if !validPassword(password) {
		fmt.Println(wrongPassword)
		return
	}
	fmt.Println(ok)
}
