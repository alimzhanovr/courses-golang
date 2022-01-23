package main

import "fmt"

const formatTime = "It is %d hours %d minutes."
const secondsInHour = 3600
const minutesInHour = 60
const secondsInMinute = 60

func Time(second int) string {
	if second <= 0 && second >= 86399 {
		return ""
	}
	hour := second / secondsInHour
	minutes := second / secondsInMinute
	minute := minutes - hour*minutesInHour
	return fmt.Sprintf(formatTime, hour, minute)
}
