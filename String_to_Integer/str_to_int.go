package main

import "fmt"

func StrToInt(s string) int {
	if len(s) == 0 { //if inout is empty
		return 0
	}

	//initializing the sign and the strat character
	sign := 1  //assuming the number is positive
	start := 0 //start from the first character

	switch s[0] {
	case '-':
		sign = -1
		start = 1
	case '+':
		start = 1
	}
	// checking for multiple signs like "--12" or "++12"
	if start == 1 && (len(s) > 1 && (s[1] == '-' || s[1] == '+')) {
		return 0
	}

	result := 0

	for i := start; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		result = result*10 + int(s[i]-'0')
	}
	return sign * result
}

func main() {
	fmt.Println(StrToInt("--123"))
	fmt.Println(StrToInt("++123"))
	fmt.Println(StrToInt("123"))
	fmt.Println(StrToInt("0"))
	fmt.Println(StrToInt("abc"))
	fmt.Println(StrToInt("87 123"))
}
