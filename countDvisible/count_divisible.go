package main

import (
	"fmt"
)

func CountDivisible( from, to, step, divisor int) int{
	if step <= 0 || divisor <=0 || from >= to{
		return 0
	} else{
		count:=0
		for i :=from; i < to; i+=step {
			if i%divisor == 0{
				fmt.Printf("%d, ", i)
				count++
			}	
		}
		fmt.Println()
		return count
	}	
}

func main(){
	result := CountDivisible(1, 20,2,3)
	fmt.Println(result)
}