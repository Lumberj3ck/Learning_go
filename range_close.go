package main

import (
	"fmt"
	// "time"
)

func fibonaci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func testSelect(ch <-chan int, ch1 <-chan int) {
	// time.Sleep(time.Millisecond * 100)
	select {
	case v := <-ch:
		fmt.Println("Got V:", v)
	case v := <-ch1:
		fmt.Println("Got V1:", v)
	}
}

func createSelectChan() {
	ch := make(chan int)
	ch1 := make(chan int)
	go func() {
		ch <- 2
	}()
	go func() {
		ch <- 1
	}()
	testSelect(ch, ch1)
}

func runFib() {
	c := make(chan int, 10)
	go fibonaci(10, c)

	for i := range c {
		fmt.Println(i)
	}
}

func main() {
	createSelectChan()
}
