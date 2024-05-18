package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var sl []int = primes[0:3]

	fmt.Println(sl, len(sl))
}
