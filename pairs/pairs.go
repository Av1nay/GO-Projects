package main

import "fmt"

func Pairs() string {
	result := "" // initialize the empty string to append the output

	for i := 0; i <= 98; i++ { //looping through the first number 0 to 98
		for j := i + 1; j <= 99; j++ { //looping through the second number from 1 to 99
			pair := fmt.Sprintf("%02d %02d", i, j) // padding 0 with 2 digit width of the decimal d

			if result != "" { //adding ", " before upcomming pair
				result += ", "
			}
			result += pair
		}
	}
	return result
}

func main() {
	fmt.Println(Pairs())
}
