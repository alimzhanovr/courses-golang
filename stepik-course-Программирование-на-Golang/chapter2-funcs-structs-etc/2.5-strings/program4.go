package main

func OddStr(str string) string {
	res := ""
	for idx := range str {
		if idx%2 == 1 {
			res += string(str[idx])
		}
	}
	return res
}
