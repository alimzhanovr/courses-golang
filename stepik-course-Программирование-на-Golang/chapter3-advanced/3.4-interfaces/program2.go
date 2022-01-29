package main

import (
	"fmt"
	"strings"
)

type battery string

func (b battery) String() string {
	str := string(b)
	count := 0
	for _, s := range str {
		if s == '1' {
			count++
		}
	}
	return fmt.Sprintf("[%s%s]", strings.Repeat(" ", len(b)-count), strings.Repeat("X", count))
}
