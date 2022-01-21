package main

func filter(predicate func(int) bool, numbers []int) []int {
	res := make([]int, 0, len(numbers))
	for _, val := range numbers {
		if predicate(val) {
			res = append(res, val)
		}
	}
	return res
}
