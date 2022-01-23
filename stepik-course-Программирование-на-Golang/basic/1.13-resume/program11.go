package main

import "fmt"

const korov = "%d korov"
const korovy = "%d korovy"
const korova = "%d korova"

func lastNum(number, lastNum int) bool {
	if number%10 == lastNum {
		return true
	}
	return false
}
func equal(a, b int) bool {
	if a == b {
		return true
	}
	return false
}
func numberInNumbers(number int, numbers []int) bool {
	for _, num := range numbers {
		if equal(num, number) {
			return true
		}
	}
	return false
}
func numberKorov(number int) string {
	if lastNum(number, 1) && !equal(number, 11) {
		return fmt.Sprintf(korova, number)
	}
	case2 := lastNum(number, 2) || lastNum(number, 3) || lastNum(number, 4)
	if case2 && !numberInNumbers(number, []int{12, 13, 14}) {
		return fmt.Sprintf(korovy, number)
	}
	return fmt.Sprintf(korov, number)
}

func Program11() {
	fmt.Println(numberKorov(1))
	fmt.Println(numberKorov(11))
	fmt.Println(numberKorov(12))
	fmt.Println(numberKorov(2))
	fmt.Println(numberKorov(10))
}
