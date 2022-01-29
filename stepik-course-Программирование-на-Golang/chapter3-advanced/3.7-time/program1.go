package main

import (
	"fmt"
	"time"
)

func Program1() {
	td, err := time.Parse(time.RFC3339, "1986-04-16T05:20:00+06:00")
	if err != nil {
		return
	}
	unixD := td.Format(time.UnixDate)
	fmt.Println(unixD)
}
