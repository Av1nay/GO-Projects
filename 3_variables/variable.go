package main

import "fmt"

func main() {
	// This is a simple Go program that demonstrates the use of variables.
	fmt.Println("This is a variable in Go:")
	
	// Declare a variable of type int
	var age int = 30
	fmt.Println("Age:", age)

	age2 := 25 
	fmt.Println("This Short variable declaration Using age2 := ", age2) // Using short variable declaration

	// Declare a variable of type float64
	var height float64 = 5.9
	fmt.Println("Height:", height)

	//short variable declaration
	height2 := 6.1
	fmt.Println("This Short variable declaration Using height2 :=", height2) // Using short variable declaration

	// Declare a variable of type complex128
	var complexNum complex128 = 1 + 2i
	fmt.Println("Complex Number:", complexNum)

	// Declare a variable of type rune
	var char rune = 'A'
	fmt.Println("Convert rune to string 'A':", rune(char)) // Convert rune to string for printing
	

	// Declare a variable of type string
	var name string = "John Doe"
	fmt.Println("Name:", name)

	// Declare a variable of type bool
	var isEmployed bool = true
	fmt.Println("Is Employed:", isEmployed)

	// Declare a variable of type array
	var scores [3]int = [3]int{90, 85, 88}
	fmt.Println("Scores:", scores)

	// Declare a variable of type slice
	var grades []int = []int{90, 85, 88}
	fmt.Println("Grades:", grades)
}