package main

import "fmt"

func findMax(numbers [5]int) int {
	max := numbers[0]
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}
func Program3() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	// здесь ваш код
	// ...
	fmt.Println(findMax(array))
}
