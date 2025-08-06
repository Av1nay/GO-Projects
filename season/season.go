package main

import "fmt"


func Season(month string) string{
	switch month {
	case "dec", "jan", "feb":
		return "winter"
	case "mar", "apr", "may":	
		return "spring"
	case "jun", "jul", "aug":
		return "summer"	
	case "sep", "oct", "nov":	
		return "autumn"		
	default:
		return "invalid month" + month
	}
}

func main() {
	fmt.Println(Season("dec"))
	fmt.Println(Season("mar"))
	fmt.Println(Season("jul"))
	fmt.Println(Season("nov"))
}
