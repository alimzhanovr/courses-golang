package main

import "fmt"

func work() {
	fmt.Println("someWork")
}
func myFunc(done chan struct{}) {
	work()
	close(done)
}
