package main

import "fmt"

func SimpleStrToInt(s string)int{
	result:=0 //initializing the result to 0 

	for _, ch := range s { // _ ignore the index position of the string and only cacrhes the character ch from the string range s
		if ch < '0' || ch >'9'{
			return 0 // if the input is invalid or input is other than 0-9
		}
		result = result*10 + int(ch-'0') // convert the character to its numeric value and add to the result
	}
	return result
}

func main(){
	fmt.Println(SimpleStrToInt("00000012345"))
}