package main

import (
	ms "exercise/mathslice"
	"fmt"
)

func main() {
	s := ms.Slice{1, 2, 3}
	sliceSum := ms.SumSlice(s)
	fmt.Println(sliceSum)
}
