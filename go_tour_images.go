package main

import (
	"fmt"
	"image"
	// "image/color"
)

func main() {
	// m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	// fmt.Println(m.Bounds())
	// fmt.Println(m.At(8, 10).RGBA())
	fmt.Println(image.Rect(0, 0, 150, 150).Max.Y)

}
