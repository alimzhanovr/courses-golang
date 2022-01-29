package main

func calculator1(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	res := make(chan int)
	go func() {
		defer close(res)
		select {
		case num := <-firstChan:
			res <- num * num
		case num := <-secondChan:
			res <- num * 3
		case <-stopChan:
			return
		}
	}()
	return res
}
