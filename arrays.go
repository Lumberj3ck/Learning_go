package main

import "fmt"

func main() {
	var a [10]int
	a[0] = 1
	for i := 0; i < 10; i++ {
		a[i] = i
	}
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 9, 11}
	fmt.Println(primes)
}
