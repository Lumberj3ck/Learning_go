package main

import "fmt"

func main() {
	a := make([]int, 100)
	for i := 0; i < 100; i++ {
		a[i] = i + 1
	}

	b := append(a[:10], a[90:]...)

	fmt.Println(b)

	dim := len(a)
	for i := range a[:dim/2] {
		a[i], a[dim-i-1] = a[dim-i-1], a[i]
	}
	fmt.Println(a)
}
