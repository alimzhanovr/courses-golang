package main

const (
	yearEarth = 365
	yearMars  = 687
)

func marsAge(year int) int {
	return year * yearEarth / yearMars
}
