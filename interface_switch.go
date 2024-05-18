package main

import "fmt"

func printSomething(smth interface{}) {
	switch t := smth.(type) {
	case int:
		fmt.Printf("%T -> %v\n", t, t)
	case string:
		fmt.Printf("%T -> %v\n", t, t)
	default:
		fmt.Printf("%T -> %v\n", t, t)
	}
}

func main() {
	printSomething(uint(123)) // in default case
	printSomething(123)       // in int case
}
