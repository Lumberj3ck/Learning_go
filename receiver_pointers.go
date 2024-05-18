package main

import "fmt"

type Vertex struct {
	X, Y int
}

func mul(v *Vertex) {
	fmt.Println(v.X * v.Y)
}

func (v *Vertex) mul_receiver() {
	fmt.Println(v.X * v.Y)
}

func (v Vertex) vert_sum() {
	fmt.Println(v.X + v.Y)
}

func main() {
	var vert = &Vertex{1, 2}
	// var vert = Vertex{1, 2}
	mul(vert)
	vert.mul_receiver()
}
