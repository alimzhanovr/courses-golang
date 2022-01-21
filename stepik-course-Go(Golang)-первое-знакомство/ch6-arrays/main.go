package main

import "fmt"

func main() {
	slc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(sum(slc...))
	fmt.Println(slc)
}
