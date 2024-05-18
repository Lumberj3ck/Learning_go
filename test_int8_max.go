package main

import "fmt"

func test() uint16 {
	var test_int uint16
	test_int = 32768
	return test_int
}

func main() {
	res := test()
	fmt.Println(res)
}
