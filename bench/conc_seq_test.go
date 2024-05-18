package main

import (
	"testing"
)

func sum_seq(s []int) int {
	sum := 0
	for _, v := range s {
		// fmt.Println(v)
		sum += v
	}
	return sum
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		// fmt.Println(v)
		sum += v
	}
	c <- sum // send sum to c
}

func BenchmarkSumConcurrent(b *testing.B) {
	s := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		s[i] = i
	}
	c := make(chan int)

	for i := 0; i < b.N; i++ {
		go sum(s[:len(s)/2], c)
		go sum(s[len(s)/2:], c)
		// x, y := <-c, <-c
		// fmt.Println(x, y)
		<-c
		<-c
	}
}

func BenchmarkSumSequential(b *testing.B) {
	s := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		s[i] = i
	}

	for i := 0; i < b.N; i++ {
		sum_seq(s[:len(s)/2]) // No channel needed for sequential version
		sum_seq(s[len(s)/2:])
	}
}

func main() {}
