package main

func WorkAndArray(numbers [10]uint8, change []int) [10]uint8 {
	if len(change)%2 != 0 {
		return [10]uint8{}
	}

	for i := 0; i < len(change)-1; i += 2 {
		numbers[change[i]], numbers[change[i+1]] = numbers[change[i+1]], numbers[change[i]]
	}
	return numbers
}
