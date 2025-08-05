package main

import "fmt"


func main() {
	fmt.Print("this is strings in golang >")
	fmt.Println("Hello, World!") //this is strings in golang
	fmt.Print("this is integers in golang >")
	fmt.Println(123) //this is integers in golang
	fmt.Print("this is floats in golang >")
	fmt.Println(3.14) //this is floats in golang
	fmt.Print("this is booleans in golang >")
	fmt.Println(true) //this is booleans in golang
	fmt.Print("this is arrays in golang >")
	fmt.Println([3]int{1, 2, 3})
	fmt.Print("this is slices in golang >")
	fmt.Println([]int{1, 2, 3}) //this is slices in golang
	fmt.Print("this is maps in golang >")
	fmt.Println(map[string]int{"one": 1, "two": 2}) //this is maps in golang
	fmt.Print("this is channels in golang >")
	ch := make(chan int)
	go func() {
		ch <- 42
	}()
	fmt.Println(<-ch) //this is channels in golang
	fmt.Print("this is functions in golang >")
	fmt.Println(func() string { return "Hello from a function!" }()) //this is functions in golang
	fmt.Print("this is interfaces in golang >")
	var i interface{} = "I am an interface"
	fmt.Println(i) //this is interfaces in golang
	fmt.Print("this is structs in golang >")
	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Alice", Age: 30}
	fmt.Println(p) //this is structs in golang
	fmt.Print("this is pointers in golang >")
	ptr := &p
	fmt.Println(*ptr) //this is pointers in golang
	fmt.Print("this is error handling in golang >")
	err := fmt.Errorf("this is an error")
	if err != nil {
		fmt.Println(err) //this is error handling in golang
	}
	fmt.Print("this is type assertions in golang >")
	var x interface{} = "Hello"
	s, ok := x.(string)
	if ok {
		fmt.Println(s) //this is type assertions in golang
	}
	fmt.Print("this is type switches in golang >")
	switch v := x.(type) {
	case string:
		fmt.Println("String:", v) //this is type switches in golang
	case int:
		fmt.Println("Integer:", v)
	default:
		fmt.Println("Unknown type")
	}
	fmt.Print("this is goroutines in golang >")
	go func() {
		fmt.Println("This runs in a goroutine") //this is goroutines in golang
	}()
	fmt.Print("this is defer statements in golang >")
	defer fmt.Println("This runs after the main function exits") //this is defer statements in golang
	fmt.Print("this is panic and recover in golang >")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r) //this is panic and recover in golang
		}
	}()
	panic("This is a panic example") //this will cause a panic
	fmt.Print("this is constants in golang >")
	const Pi = 3.14
	fmt.Println(Pi) //this is constants in golang
	fmt.Print("this is iota in golang >")
	const (
		First  = iota // 0
		Second        // 1
		Third         // 2
	)
	fmt.Println(First, Second, Third) //this is iota in golang
	fmt.Print("this is nil in golang >")
	var n interface{} = nil
	fmt.Println(n) //this is nil in golang
	fmt.Print("this is string concatenation in golang >")
	fmt.Println("Hello, " + "World!") //this is string concatenation in golang
	fmt.Print("this is formatted strings in golang >")
	fmt.Printf(fmt.Sprintf("Formatted: %d", 42)) //this is formatted strings in golang
}
