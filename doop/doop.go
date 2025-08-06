package main

import "fmt"
func Doop ( a int, op string, b int ) int {
	if b <= 0 && (op == "/" || op == "%") {
		return 0
	} else{
		switch op {
		case "+" :
			return (a + b)
		case "-" :
			return (a - b)
		case "*" :
			return (a * b)
		case "/" :
			return (a / b)
		case "%" :
			return (a % b)
		default :
			return 0
		}
	}
}

func main(){
	fmt.Printf("%d\n",(Doop(2, "+", 6 )))
}