package main

func task(ch chan int, N int) {
	ch <- N + 1
}
