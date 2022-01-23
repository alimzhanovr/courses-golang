package main

import "fmt"

func main() {
	fmt.Println(RepeatString(ILikeGo(), 3))
}

func ILikeGo() string {
	return "I like Go!"
}

func RepeatString(str string, num int) string {
	res := ""
	if num < 0 {
		return res
	}
	if num == 0 || num == 1 {
		return str
	}
	for i := 0; i < num; i++ {
		res += str + "\n"
	}
	return res[:len(res)-1]
}
