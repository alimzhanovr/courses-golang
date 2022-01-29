package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func Program3() {
	numbersStr, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && err != io.EOF {
		return
	}
	numbersStr = strings.ReplaceAll(numbersStr, " ", "")
	numbersStr = strings.ReplaceAll(numbersStr, "\n", "")
	numbersStr = strings.ReplaceAll(numbersStr, ",", ".")
	numbersSlc := strings.Split(numbersStr, ";")
	numbers := make([]float64, 0, len(numbersSlc))
	for _, numberS := range numbersSlc {
		number, err := strconv.ParseFloat(string(numberS), 64)
		if err != nil {
			return
		}
		numbers = append(numbers, number)
	}
	fmt.Printf("%.4f\n", numbers[0]/numbers[1])
}
