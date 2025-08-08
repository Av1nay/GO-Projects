package main

import "fmt"

func ReverseAlphabetValue(ch rune) rune {
	return 'z' - (ch - 'a')
}

func main() {
	fmt.Println(ReverseAlphabetValue('h'))
}
