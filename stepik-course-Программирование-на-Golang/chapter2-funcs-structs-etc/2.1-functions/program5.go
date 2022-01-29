package main

func SumNumbers(numbers ...int) (sum int, numsLen int) {
	numsLen = len(numbers)
	for _, number := range numbers {
		sum += number
	}
	return
}
