package main

import (
	"fmt"
	"strings"
)

func Program3() {
	var str, sub string
	fmt.Scan(&str)
	fmt.Scan(&sub)
	fmt.Println(strings.Index(str, sub))
}
