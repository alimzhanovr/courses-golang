package main

import "fmt"

func main() {
	fmt.Println(FormatOut(-000012.2123))
	fmt.Println(FormatOut(1000001))
}

func FormatOut(number float64) string {
	if number <= 0 {
		return fmt.Sprintf("число %4.2f не подходит", number)
	}
	if number > 10000 {
		return fmt.Sprintf("%e", number)
	}
	return fmt.Sprintf("%.4f", number*number)
}
