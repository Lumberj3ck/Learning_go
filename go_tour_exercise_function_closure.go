package main

import "fmt"

func fibonaci() func() int {
	prev, curr := 0, 1
	return func() int {
		next := prev + curr
		prev, curr = curr, next
		return next
	}
}

func main() {
	f := fibonaci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

}
