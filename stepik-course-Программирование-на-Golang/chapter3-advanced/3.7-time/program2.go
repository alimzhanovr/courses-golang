package main

import (
	"fmt"
	"time"
)

func Program2() {
	const format = "2006-01-02 15:04:05"
	date, h := "", ""
	fmt.Scan(&date, &h)
	ts := date + " " + h
	td, err := time.Parse(format, ts)
	if err != nil {
		return
	}
	if td.Hour() < 13 {
		fmt.Println(ts)
		return
	}
	td = td.AddDate(0, 0, 1)
	fmt.Println(td.Format(format))
}
