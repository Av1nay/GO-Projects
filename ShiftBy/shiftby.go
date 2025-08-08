package main

import "fmt"

func shiftby(r rune, step int) rune {
	indexch := (int(r-'a') + step) % 26
	if indexch < 0 {
		indexch += 26
	}
	//fmt.Printf("%c\n", 'a'+rune(indexch))
	return 'a' + rune(indexch)
}

func main() {
	fmt.Printf("%c\n", shiftby('e', 4))
}
