package main

import (
	"fmt"
	"math"
)

func abs(n float64) float64 {
	if n < 0 {
		return -n
	}
	return n
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	guess := float64(1)
	found := false
	if x <= 0 {
		return 0, ErrNegativeSqrt(x)
	}
	for !found {
		prev_guess := guess
		guess -= (guess*guess - x) / (2 * guess)
		change := guess - prev_guess
		if abs(change) < math.Pow(10, -6) {
			found = true
		}
	}
	return guess, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
