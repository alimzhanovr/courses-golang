package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(Abbreviation("Today I learned "))
}
func Abbreviation(text string) string {
	res := ""
	text = strings.TrimSpace(text)
	strs := strings.Fields(text)
	for _, str := range strs {
		firstLetter := []rune(str)[0]
		if !unicode.IsLetter(firstLetter) {
			continue
		}
		res += string(unicode.ToUpper(rune(firstLetter)))
	}
	return res
}
