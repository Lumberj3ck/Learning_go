package main

import "fmt"

var pow = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func main() {

	for _, v := range pow {
		fmt.Printf("2 ** %d = %v \n", v, 1<<v)
	}
}
