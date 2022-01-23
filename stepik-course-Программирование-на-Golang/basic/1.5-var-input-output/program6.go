package main

import "fmt"

const time = "It is %d hours %d minutes."
const minutesInHour = 60
const minutesInDegree = 2

func findTime(degree int) string {
	if 0 > degree || degree > 360 {
		return ""
	}
	minutes := degree * minutesInDegree
	fmt.Println(minutes)
	hour := minutes / minutesInHour
	minute := minutes % minutesInHour
	return fmt.Sprintf(time, hour, minute)
}
