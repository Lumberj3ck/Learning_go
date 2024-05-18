package main

import (
	"fmt"
	"time"
)

func sendEmail(message string) {
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
	}()
	fmt.Printf("Email sent: '%s'\n", message)
}

// Don't touch below this line

func test(message string) {
	sendEmail(message)
	time.Sleep(time.Millisecond * 555)
	fmt.Println("========================")
}

func testParalelism() {
	go func() {
		for i := 1; i < 99; i++ {
			fmt.Println(i)
		}
	}()
	func() {
		for i := 999; i < 1050; i++ {
			fmt.Println(i)
		}
	}()
	// time.Sleep(time.Millisecond * 500)
}

func say(s string) {
	for i := 0; i < 5; i++ {
		// time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// test("Hello there Kaladin!")
	// test("Hi there Shallan!")
	// test("Hey there Dalinar!")
	testParalelism()
	// go say("world")
	// say("hello")
}
