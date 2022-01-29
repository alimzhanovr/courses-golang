package main

import "fmt"

const errror = "ошибка"

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf(errror)
	}
	return a / b, nil
}
func main() {
	//1
	//Вы должны вызвать функцию divide которая делит два числа, но она возвращает не только результат деления,
	//но и информацию об ошибке. В случае какой-либо ошибки вы должны вывести "ошибка", если ошибки нет - результат функции.
	//Функция divide(a int, b int) (int, error) получает на вход два числа которые нужно поделить и возвращает результат
	//(int) и ошибку (error). Пакет main уже объявлен, а нужные пакеты уже импортированы!
	//Не забудьте считать два числа которые необходимо поделить (передать в функцию)!
	//Sample Input:
	//10 5
	//Sample Output:
	//2
	res, err := divide(10, 5)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
