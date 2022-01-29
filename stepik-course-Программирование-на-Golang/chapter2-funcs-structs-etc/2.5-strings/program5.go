package main

import "strings"

func Program5(str string) string {
	res := ""
	for _, s := range str {
		if strings.Count(str, string(s)) == 1 {
			res += string(s)
		}
	}
	return res
}
