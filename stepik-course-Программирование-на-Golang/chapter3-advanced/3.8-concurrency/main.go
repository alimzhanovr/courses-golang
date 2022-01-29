package main

import (
	"fmt"
	"time"
)

func main() {
	//1
	//Напишите функцию которая принимает канал и число N типа int. Необходимо вернуть значение N+1 в канал.
	//Функция должна называться task().
	//Внимание! Пакет и функция main уже объявлены, выводить и считывать ничего не нужно!
	ch := make(chan int)
	go task(ch, 41)
	fmt.Println(<-ch)

	//2
	//Напишите функцию которая принимает канал и строку. Необходимо отправить эту же строку в заданный канал 5 раз,
	//добавив к ней пробел.
	//Функция должна называться task2().
	ch2 := make(chan string)
	go task2(ch2, "hello")
	for str := range ch2 {
		fmt.Println(str)
	}

	//3
	//Напишите элемент конвейера (функцию), что запоминает предыдущее значение и отправляет значения
	//на следующий этап конвейера только если оно отличается от того, что пришло ранее.
	//Ваша функция должна принимать два канала - inputStream и outputStream, в первый вы будете получать строки,
	//во второй вы должны отправлять значения без повторов. В итоге в outputStream должны остаться значения,
	//которые не повторяются подряд. Не забудьте закрыть канал ;)
	//Функция должна называться removeDuplicates()
	//Выводить или вводить ничего не нужно!
	inputStream := make(chan string)
	outputStream := make(chan string)
	go removeDuplicates(inputStream, outputStream)
	go func() {
		defer close(inputStream)
		for _, r := range "112334456" {
			inputStream <- string(r)
		}
	}()
	for x := range outputStream {
		fmt.Print(x)
	}
	fmt.Println()
	// 123456
	////////////////////////
	time.Sleep(2 * time.Second)
}
