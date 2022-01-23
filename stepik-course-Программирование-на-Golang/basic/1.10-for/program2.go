package main

func noMoreThan100(number int) bool {
	if number > 100 {
		return false
	}
	return true
}

func SumOfNumbersBetween(from, to int) int {
	if !(noMoreThan100(from) && noMoreThan100(to)) {
		return -1
	}
	if from > to {
		return -1
	}
	if from == to {
		return from
	}
	sum, curr, iterCount := 0, from, to-from+1
	for i := 0; i < iterCount; i++ {
		sum += curr
		curr++
	}
	return sum
}
