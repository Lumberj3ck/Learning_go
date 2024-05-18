package main

import (
	"fmt"
	"time"
)

func processMessages(messages []string) []string {
	m := make(chan struct{}, len(messages))

	for idx, message := range messages {
		go func() {
			messages[idx] = message + "-processed"
			m <- struct{}{}
		}()
	}
	for {
		if len(m) == len(messages) {
			return messages
		} else {
			time.Sleep(time.Millisecond * 50)
		}
	}
}

func main() {
	res := processMessages([]string{"asdf", "asdf", "something"})
	fmt.Println(res)
}
