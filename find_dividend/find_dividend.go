package main

import "fmt"


func FindDividend(from, to, divisor int) int{
	if divisor == 0{
		return 0
	}
	for i:= from; i < to; i++{
		if i%divisor == 0{
			return i
		}
	}
	return -1
}
func main(){
	fmt.Println("the fisrt dividend is : ", FindDividend(50, 511, -26) )
}