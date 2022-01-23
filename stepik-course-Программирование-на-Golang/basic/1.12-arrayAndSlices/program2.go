package main

func Element4ofSlice(numbers []int) int {
	if len(numbers) <= 4 {
		return 0
	}
	//4 element it is idx=3
	return numbers[3]
}
