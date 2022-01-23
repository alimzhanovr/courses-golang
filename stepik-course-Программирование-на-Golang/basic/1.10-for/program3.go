package main

func doubleDigit(number int) bool {
	if 9 < number && number < 100 {
		return true
	}
	return false
}
func multiple8(number int) bool {
	if number%8 == 0 {
		return true
	}
	return false
}
func multiple8AndDoubleDigit(numbers []int) []int {
	if len(numbers) == 0 {
		return []int{}
	}
	res := make([]int, 0, len(numbers))
	for _, number := range numbers {
		if multiple8(number) && doubleDigit(number) {
			res = append(res, number)
		}
	}
	return res
}

func SumOfMultiple8AndDoubleDigit(numbers []int) int {
	numbers = multiple8AndDoubleDigit(numbers)
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
