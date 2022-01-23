package main

func SquareOfNumbers(from, to int) []int {
	//[from; to]
	iterCount := to - from + 1
	res := make([]int, 0, iterCount)
	if from > to {
		return []int{}
	}
	if from == to {
		return []int{from}
	}
	curr := from
	for i := 0; i < iterCount; i++ {
		res = append(res, curr*curr)
		curr++
	}
	return res
}
