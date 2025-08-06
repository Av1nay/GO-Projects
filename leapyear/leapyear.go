package main

import "fmt"

func IsLeapYear(year int) bool {
	if year % 4 == 0 {
		if (year % 100 == 0 ){
			if (year % 400 == 0){
				return true
			}
			return false
		}
		return true
	} else {
		return  false
	}
}

func main(){
	year:= 1900
	fmt.Printf("Is %d a leap year: %t\n", year, IsLeapYear(year))
}