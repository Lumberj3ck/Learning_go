package main

import "fmt"

func defferedHelloWorld() {
	defer fmt.Println("World")

	fmt.Println("Hello")
}

func defferStack() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}

func main() {
	defferStack()
}
