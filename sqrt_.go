package main

import "fmt"
import "math"

func module(n float64) float64 {
	if n < 0 {
		return -n
	}
	return n
}

func Sqrt(n float64) float64 {
	guess := float64(1)
	found := false
	for !found {
		prev_guess := guess
		guess -= (guess*guess - n) / (2 * guess)
		change := guess - prev_guess
		if module(change) < math.Pow(10, -6) {
			found = true
		}
	}
	return guess
}

func main() {
	number := float64(2)
	s := Sqrt(number)
	s1 := math.Sqrt(number)
	fmt.Println("My sqrt:", s)
	fmt.Println("Math library sqrt:", s1)
	fmt.Println("Difference ", module(s-s1))
}
