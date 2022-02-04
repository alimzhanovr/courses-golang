package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	sl := make([][]uint8, 0, dy)
	for y := 0; y < dy; y++ {
		inner_sl := make([]uint8, 0, dx)
		for x := 0; x < dx; x++ {
			inner_sl = append(inner_sl, uint8((x+y)/2+x*y*100+x^y))
		}
		sl = append(sl, inner_sl)
	}
	return sl
}

func main() {
	pic.Show(Pic)
}
