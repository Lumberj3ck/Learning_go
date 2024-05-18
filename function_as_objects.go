package main

import "fmt"

func div(a, b float64) float64 {
	return a / b
}

func mul(a, b float64) float64 {
	return a * b
}

func calculation(fn func(float64, float64) float64, a float64, b float64) {
	fmt.Println(fn(3, 4))
}

func main() {
	operations := []func(float64, float64) float64{div, mul}
	fmt.Println(len(operations[1:]), cap(operations))
	for _, op := range operations {
		calculation(op, 9, 4)
	}

	var s = mul
	fmt.Println(s(18, 99))
}
