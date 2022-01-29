package main

func MaxNum(number int) int {
	max := 0
	for {
		if number < 10 {
			if number > max {
				return number
			}
			return max
		}
		lastNum := number % 10
		if lastNum > max {
			max = lastNum
		}
		number /= 10
	}
}
