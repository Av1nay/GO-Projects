package main

import "fmt"

func IntVsFloat (i int, f float32) string{
	if f == float32(i) {
		return "Equal"
	} else if f > float32(i) {
		return "float"
	} else {	
		return "Integer"
	}
}

func main () {
	fmt.Println(IntVsFloat(1, 2.0))
}