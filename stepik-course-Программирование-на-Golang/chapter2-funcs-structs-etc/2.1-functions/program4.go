package main

func FibNumWithIndx(indx int) int {
	if indx <= 0 {
		return 0
	}
	if indx == 1 || indx == 2 {
		return 1
	}
	prev, curr := 1, 1
	for i := 3; i <= indx; i++ {
		prev, curr = curr, prev+curr
	}
	return curr
}
