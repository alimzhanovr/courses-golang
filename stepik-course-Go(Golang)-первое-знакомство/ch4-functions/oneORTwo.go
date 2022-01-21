package main

const (
	one = "one"
	two = "two"
)

func oneORTwo(a, b int, s string) (int, string) {
	if s == one {
		return a, s
	}
	if s == two {
		return b, s
	}
	return 0, s
}
