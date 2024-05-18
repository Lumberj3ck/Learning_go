package main

import "fmt"

var Global = 5

func changePrint() {
	defer func(s int) {
		Global = s
	}(Global)
	Global = 99999
	fmt.Println(Global)
}

func main() {
	changePrint()
	fmt.Println(Global)
}
