package main

import "fmt"

func Countdown(n int) string {
	result := ""
	if n == 0 {
		return "0!"
	}
	for i := n; i > 0; i -= 2 {
		result += string('0'+i) + ","

	}
	result += "0!"
	return result
}

func main() {
	fmt.Println(Countdown(7))
}
