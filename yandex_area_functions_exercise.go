package main

import "fmt"
import "math"

type figures int

const (
	square figures = iota + 1
	circle
	triangle
)

func area(f figures) (func(float64) float64, bool) {
	known := f < 4 && f > 0

	return func(s float64) float64 {
		if f == square {
			return s * s
		} else if f == circle {
			return 3.14 * s * s
		} else if f == triangle {
			return math.Sqrt(s) * s * s / 4
		}
		return 0
	}, known
}

func main() {
	ar, ok := area(1)
	if !ok {
		fmt.Println("asdf")
	}
	fmt.Println(ar(4))
}
