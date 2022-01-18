package main

func doubleM(a, b int) int {
	var sum int
	if b == a {
		return a
	}

	if b > a {
		swap(&a, &b)
	}

	for a <= b {
		sum += (a * a)
		a++
	}
	return sum
}
