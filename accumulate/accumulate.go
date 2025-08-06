package main

import "fmt"

func accumulate(n int) int {
	if n > 0 {
		sum := 0
		for i :=0; i <= n; i++ {
			sum += i	
		}
		return sum
	}
	return 0
}

func main () {
	fmt.Println(accumulate(5))
}