package main

import "fmt"

func repeatStrings(str string, count int) (string, error) {
	if count < 0 {
		return "", fmt.Errorf("the count must be greater than zero")
	}
	if count == 0 {
		return "", nil
	}
	res := ""
	for i := 0; i < count; i++ {
		res += str
	}
	return res, nil
}
