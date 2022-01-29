package main

func merge2Channels(f func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	idx := make(chan int) // Для передачи индекса
	val := make(chan int) // Для передачи значения

	go func() {
		for i := 0; i < n; i++ {
			x1 := <-in1
			x2 := <-in2

			done := make(chan int)
			go func() {
				done <- f(x1)
			}()
			go func() {
				done <- f(x2)
			}()

			go func(i int) {
				idx <- i
				val <- (<-done) + (<-done)
			}(i)
		}
	}()

	go func() {
		m := make(map[int]int)
		for v := range val {
			m[<-idx] = v
		}

		// Для возврата значений по порядку
		for i := 0; i < n; i++ {
			out <- m[i]
		}
	}()
}
