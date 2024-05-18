package main

import (
	"bufio"
	"fmt"
	"os"
)

func inc(c *int) {
	*c++
}

func main() {
	// Получаем читателя пользовательского ввода
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Interaction counter")

	cnt := 0
	for {
		fmt.Print("-> ")
		_, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		inc(&cnt)

		fmt.Printf("User input %d lines\n", cnt)
	}
}
