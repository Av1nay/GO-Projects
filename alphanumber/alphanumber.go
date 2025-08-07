package main

import "fmt"

func AlphaNumber(n int) string{
	isNegative:= n<0
	if isNegative{
		n=-n
	}
	digits:=[]int{} // initialize empty slice to hold digit of n in reverse order
	for n > 0{
		digits = append(digits, n%10) // add extracted digit in slice in reverse order
		n/=10 //removes last digit e.g. 123/10=12 as it is int
	}

	//building the result in reverse order of the slice
	result:=""
	for i :=len(digits)-1;i>=0; i--{ //loops backward through the digit slice
		result+=string('a'+digits[i])//converting digit to alphabet
	}
	if isNegative{
		result ="-"+result
	}
	return result
}

func main(){
	fmt.Println(AlphaNumber(-1234))
}