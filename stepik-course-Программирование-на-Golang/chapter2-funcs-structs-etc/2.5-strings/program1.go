package main

import (
	"unicode"
)

const (
	right = "Right"
	wrong = "Wrong"
)

func startWithUpperLetter(text string) bool {
	if unicode.IsUpper([]rune(text)[0]) {
		return true
	}
	return false
}
func CorrectString(text string) string {
	if startWithUpperLetter(text) && text[len(text)-1] == '.' {
		return right
	}
	return wrong
}
