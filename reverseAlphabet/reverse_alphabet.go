package main

import "fmt"

func ReverseAlphabet(step int) string{
	// ensuring function behaves correctly after invalid step input and makes to default 1 even if step is 0 or negative
	if step <= 0{
		step = 1
	}
	result := "" //declairing empty string to append final output

	for i := 'z'; i >='a'; i -=rune(step){ // initializing i as rune 'z' and stops at 'a' and i decreases by "step" each time in loop begins
		result += string(i) // converting rune to string and append the output to result for final output
	}
	return result //returning the reuslt after the for loop

}

func main(){
	fmt.Println(ReverseAlphabet(6))
}
