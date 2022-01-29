package main

func Vote(x int, y int, z int) int {
	if x == y {
		return x
	}
	if x == z {
		return x
	}
	return y
}