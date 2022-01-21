package main

import "fmt"

func max(a, b int) (int, error) {
	if a > b {
		return a, nil
	} else if a < b {
		return b, nil
	}
	return 0, fmt.Errorf("a equal b, %d==%d", a, b)
}
