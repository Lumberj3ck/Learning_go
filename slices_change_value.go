package main

import "fmt"

func main() {
	var names [4]string

	a := names[1:3]
	b := names[2:3]
	a[1] = "asdaf"

	fmt.Println("Slice B", b)
	fmt.Println("Slice a", b)
}
