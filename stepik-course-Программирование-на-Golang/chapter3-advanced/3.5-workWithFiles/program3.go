package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Program3() {
	pos := 1
	f, err := os.Open("task.data")
	if err != nil {
		return
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	for {
		number, err := reader.ReadString(';')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				break
			}
		}
		if number[:len(number)-1] == "0" {
			break
		}
		//fmt.Println(number)
		pos++
	}
	fmt.Println(pos)
}
