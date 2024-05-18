package main

import "strings"
import "fmt"

func solution(str string) map[string]int {
	words_amount := make(map[string]int)
	res := strings.Fields(str)
	for _, word := range res {
		if _, ok := words_amount[word]; ok {
			words_amount[word] = words_amount[word] + 1
		} else {
			words_amount[word] = 1
		}
	}
	return words_amount
}

func main() {
	resp := solution("asdf asdf")
	fmt.Println(resp)
}
