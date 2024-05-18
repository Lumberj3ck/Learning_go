package main

import "fmt"

func Pic(dy, dx int) [][]int {
	var dy_len = make([][]int, dy)

	for i := range dy {
		dy_len[i] = make([]int, dx)
		for j := range dx {
			dy_len[i][j] = j
		}
	}
	fmt.Println(dy_len)
	return dy_len
}

func main() {
	Pic(2, 3)
}
