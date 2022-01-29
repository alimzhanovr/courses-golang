package main

func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
	res := make(chan int)
	go func() {
		defer close(res)
		sum := 0
		for {
			select {
			case num := <-arguments:
				sum += num
			case <-done:
				res <- sum
				return
			}
		}
	}()
	return res
}
