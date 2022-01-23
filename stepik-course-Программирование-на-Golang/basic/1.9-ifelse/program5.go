package main

func divisible(first, second int) bool {
	if first%second == 0 {
		return true
	}
	return false
}

func LeapYear(year uint) string {
	firstCase := divisible(int(year), 4) && !divisible(int(year), 100)
	leap := divisible(int(year), 400) || firstCase
	if leap {
		return "YES"
	}
	return "NO"
}
