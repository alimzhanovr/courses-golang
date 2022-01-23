package main

const exists = "Существует"
const notExists = "Не существует"

func TriangleExists(a, b, c int) string {
	if (a+b > c) && (a+c > b) && (b+c > a) {
		return exists
	}
	return notExists
}
