package main

import "fmt"

type Vertex struct {
	x int
	y int
}

func main() {
	i, j := 42, 2701

	p := &i

	fmt.Println(*p)

	i = 389
	*p = 23

	vert := Vertex{1, 5}

	fmt.Println(vert)
	fmt.Println("Vert x", vert.x)

	vert_pointer := &vert
	vert_pointer.x = 99
	fmt.Println("Vert pointer", vert_pointer.x)

	fmt.Println("J variable", j)
}
