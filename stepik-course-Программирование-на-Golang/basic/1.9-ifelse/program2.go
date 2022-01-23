package main

func checkTheSameDigit(number int) string {
	set := make(map[int]struct{})
	count := 0
	for {
		lastNumber := number % 10
		set[lastNumber] = struct{}{}
		count++
		if number < 10 {
			break
		}
		number /= 10
	}
	if count != len(set) {
		return "NO"
	}
	return "YES"
}
