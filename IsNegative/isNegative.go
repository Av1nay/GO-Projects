package main

import "fmt"

func isNegative(n int) bool {
	return n < 0
}
func main() {
	fmt.Println(isNegative(-1))
	fmt.Println(isNegative(0))
	fmt.Println(isNegative(1))
}