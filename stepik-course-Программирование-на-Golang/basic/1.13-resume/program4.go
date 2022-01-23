package main

const rTriangle = "Прямоугольный"
const noRTriangle = "Непрямоугольный"

func Triangle(a, b, c int) string {
	if !(a < b && b < c) {
		return ""
	}
	if c*c == a*a+b*b {
		return rTriangle
	}
	return noRTriangle
}
