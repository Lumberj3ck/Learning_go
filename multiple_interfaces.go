package main

import "fmt"

type mathAction interface {
	action() int
}

type tripplePoint struct {
	x, y, c int
}

// tripplePoint doesn't implement mathAction interface as you would expect
// but *tripplePoint implement
func (p *tripplePoint) action() int {
	p.c = int(p.x) * int(p.y)
	return p.c
}

type doublePoint struct {
	x, y int
}

func (p *doublePoint) action() int {
	return int(p.x) * int(p.y)
}

func mathExecute(action mathAction) int {
	// here we don't accept a pointer but because go gives us
	// an ability to not write like this &action.action()
	re := action.action()
	d, ok := action.(*doublePoint)
	t, ok := action.(*tripplePoint)
	fmt.Println(d)
	if ok {
		fmt.Println("Underlying type of interface is tripplePoint")
		fmt.Println(t.x, t.y, t.c)
	}
	return re
}

func main() {
	test_number := tripplePoint{3, 4, 0}
	// interface gives us ability to hide implementation of methods
	// and do not care about their implementation
	// interface gives us ability just call some methods that we ask inside of interface definition
	// test_number := doublePoint{3, 4}
	fmt.Println(mathExecute(&test_number))
	// fmt.Println(test_number.x)
	// fmt.Println(test_number.c)
}
