package main

import "fmt"
import "time"

func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}

func Fib(n int) int {
	if n <= 1 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func FibClosure() func() int {
	x1, x2 := 0, 1

	return func() int {
		x1, x2 = x2, x2+x1
		return x1
	}
}

func MetricCall(f func(string)) func(string) {
	return func(s string) {
		start := time.Now()
		f(s)
		fmt.Println(time.Now().Sub(start))
	}
}

func say(s string) {
	fmt.Println(s)
}

func main() {
	say := MetricCall(say)
	say("Hello")
}
