package main

import "fmt"

func BetweenLimits (from, to rune) string {
	var alphabet string
	if from > to {
		from, to = to, from
	}
	for i := from + 1; i < to; i++{
			alphabet += string(i)
	}
	return alphabet
}

func main (){
	fmt.Println(BetweenLimits('e', 'a'))
}