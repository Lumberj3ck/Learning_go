package main

import (
	"fmt"
	"time"
)

func testDefer() int {
	defer fmt.Println("World")

	defer fmt.Println("My")
	fmt.Println("Hello")
	return 3 + 3
}
func intuitive() string {
	value := "Казалось бы"
	defer func() { value = "На самом деле" }()
	return value
}

func unintuitive() (value string) {
	defer func() { value = "На самом деле" }() // круглые скобки в конце означают, что функция вызывается
	return "Казалось бы"
}

func evaluatingArgs() {
	a := "some text"

	// function will evaluate
	// args of the defered func when
	// meets defer operator
	defer func(s string) {
		fmt.Println(s)
	}(a)
	a = "new text"
}

func MetricTime(start time.Time) {
	fmt.Println(time.Now().Sub(start))
}

func main() {
	// fmt.Println(testDefer())
	// fmt.Println(unintuitive())
	defer MetricTime(time.Now())
	evaluatingArgs()
}
