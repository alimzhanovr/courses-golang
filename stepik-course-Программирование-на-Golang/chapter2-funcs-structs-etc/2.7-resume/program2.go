package main

import (
	"strings"
)

func AddSep(str, sep string) string {
	slc := strings.Split(str, "")
	res := strings.Join(slc, sep)
	return res
}
