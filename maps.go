package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func maps_keys() {
	var m map[string]Vertex
	m = make(map[string]Vertex)
	m["Some shit"] = Vertex{Lat: 1.1, Long: 1.1}
	m["A"] = Vertex{1.1, 1.1}
	fmt.Println(m)
}

func map_literals() {
	var literal_map = map[string]Vertex{
		"Google": Vertex{1.1, 123.432},
		// we can omit type name
		// here we omit the Vertex type name
		"Amazon": {11.44, 123.432},
	}
	fmt.Println(literal_map)
}

func main() {
	map_literals()
}
