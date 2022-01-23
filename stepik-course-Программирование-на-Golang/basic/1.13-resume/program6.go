package main

func ArithmeticMean(numbers ...int) float32 {
	if len(numbers) == 0 {
		return -1
	}
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return float32(sum) / float32(len(numbers))
}
