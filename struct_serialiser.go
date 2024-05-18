package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Person struct {
	Name        string
	Email       string
	DateOfBirth time.Time
}

func main() {
	r, o := json.Marshal(Person{"adf", "dasfd", time.Now()})
	fmt.Println(string(r), o)
}
