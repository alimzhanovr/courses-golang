package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func Program3() {
	const format = "02.01.2006 15:04:05"
	time.Now()
	str, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && err != io.EOF {
		panic(any(err))
	}
	dates := strings.Split(str[:len(str)-1], ",")
	t1, err := time.Parse(format, dates[0])
	if err != nil {
		return
	}
	t2, err := time.Parse(format, dates[1])
	if err != nil {
		return
	}
	since := time.Since(t1).Seconds() - time.Since(t2).Seconds()
	fmt.Println(time.Duration(since).Seconds())
}
