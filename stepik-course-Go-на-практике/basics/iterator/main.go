package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Определите недостающие типы:

// element - интерфейс элемента последовательности
// (пустой, потому что элемент может быть любым)
type element interface{}

// weightFunc - функция, которая возвращает вес элемента
type weightFunc func(element) int

// iterable - интерфейс последовательности, которую можно поэлементно перебирать
type iterable interface {
	// методы последовательности
	next() bool
	val() int
}

// Тип ints - последовательность из целых чисел
// (реализует интерфейс iterable)
type ints struct {
	// поля структуры
	numbers []int
	point   int
	len     int
}

func (i *ints) next() bool {
	if i.len > i.point {
		return true
	}
	return false
}
func (i *ints) val() int {
	temp := i.point
	i.point++
	return i.numbers[temp]
}

// методы ints, которые реализуют интерфейс iterable

// конструктор ints
func newInts(nums []int) *ints {
	return &ints{nums, 0, len(nums)}
}

// ┌─────────────────────────────────────────────┐
// │ не меняйте код ниже этой строки             │
// └─────────────────────────────────────────────┘

// main находит максимальное число из переданных на вход программы.
func main() {
	nums := readInput()
	it := newInts(nums)
	weight := func(el element) int {
		return el.(int)
	}
	m := max(it, weight)
	fmt.Println(m)
}

// max возвращает максимальный элемент в последовательности it.
// Для сравнения элементов используется вес, который возвращает
// функция weight.
func max(it iterable, weight weightFunc) element {
	var maxEl element
	for it.next() {
		curr := it.val()
		if maxEl == nil || weight(curr) > weight(maxEl) {
			maxEl = curr
		}
	}
	return maxEl
}

// readInput считывает последовательность целых чисел из os.Stdin.
func readInput() []int {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}
	return nums
}
