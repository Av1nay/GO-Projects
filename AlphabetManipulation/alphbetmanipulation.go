package main

import "fmt"

func num(n int) string {
	if n > 0 && n < 27 { // ensuring the n is inside the range
		result := ""             // initializing a empty result where each letter will be appended
		for i := 0; i < n; i++ { //loops from 0 to n-1
			alphabet := 'a' + i        //rune ensure working with Unicode characters and adding i gives next letters if i is 0 then alphabet will be a, if 1 then b
			result += string(alphabet) // converts rune to string and append the output to result
		}
		return result //returning the result here
	}
	return "Please input from 1 to 26" // returning error or invalid input
}

func main() {
	fmt.Println(num(1))
}
