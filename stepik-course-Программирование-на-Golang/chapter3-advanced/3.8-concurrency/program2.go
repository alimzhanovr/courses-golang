package main

func task2(ch chan string, s string) {
	s += " "
	for i := 0; i < 5; i++ {
		ch <- s
	}
	close(ch)
}
