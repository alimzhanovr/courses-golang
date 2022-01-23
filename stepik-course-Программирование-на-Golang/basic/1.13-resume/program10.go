package main

func BigMultiple(from, to, multiple int) int {
	if from > to {
		return -1
	}
	for i := to; i >= from; i-- {
		if i%multiple == 0 {
			return i
		}
	}
	return -1
}
