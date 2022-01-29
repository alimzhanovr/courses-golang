package main

import "sync"

func myFunc2(wg *sync.WaitGroup) {
	defer wg.Done()
	work()
}
