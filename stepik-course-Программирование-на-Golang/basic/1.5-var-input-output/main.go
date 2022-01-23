package main

import "fmt"

func main() {
	//first task
	//Напишите программу, которая последовательно делает следующие операции с введённым числом:
	//Число умножается на 2;
	//Затем к числу прибавляется 100.
	someNumber := 100
	res1 := multiplyAndAdd(someNumber, 2, 100)
	fmt.Println(res1)

	//second task
	//Петя торопился в школу и неправильно написал программу,
	//которая сначала находит квадраты двух чисел, а затем их суммирует. Исправьте его программу.
	fmt.Println(SquareAndSum(2, 3)) // 2*2 + 3*3 == 13

	//3 task
	//По данному целому числу, найдите его квадрат.
	fmt.Println(FindSquare(2)) //2*2 = 4

	//4 task
	//Дано натуральное число, выведите его последнюю цифру.
	//На вход дается натуральное число N, не превосходящее 10000.
	fmt.Println(FindLastNumber(9999))

	//5 task
	//Дано неотрицательное целое число.
	//Найдите число десятков (то есть вторую цифру справа).
	fmt.Println(findNumberOfTens(1234))

	//6 task
	//Часовая стрелка повернулась с начала суток на d градусов. Определите, сколько сейчас целых часов h и целых минут m.
	//Входные данные
	//На вход программе подается целое число d (0 < d < 360).
	//Выходные данные
	//Выведите на экран фразу:
	//It is ... hours ... minutes.
	//Sample Input:
	//90
	//Sample Output:
	//It is 3 hours 0 minutes.
	fmt.Println(findTime(90))
}
