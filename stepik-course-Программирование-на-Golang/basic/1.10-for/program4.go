package main

func CountOfMaxNumber(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	max, count := numbers[0], 0
	for _, number := range numbers {
		if number > max {
			max = number
			count = 1
			continue
		}
		if number == max {
			count++
			continue
		}
	}
	return count
}
