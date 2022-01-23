package main

import (
	"strconv"
	"strings"
)

func DeleteNumber(number, delNum int) int {
	if delNum < 0 || delNum > 9 {
		return number
	}
	numberS, delNumS := strconv.Itoa(number), strconv.Itoa(delNum)
	if !strings.Contains(numberS, delNumS) {
		return number
	}
	deletedNumber := ""
	for _, numS := range numberS {
		if string(numS) == delNumS {
			continue
		}
		deletedNumber += string(numS)
	}
	res, err := strconv.Atoi(deletedNumber)
	if err != nil {
		return number
	}
	return res
}
