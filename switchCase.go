package main

import "fmt"
import "time"

func checkWhoYouAre(year int) {
	switch {
	case year >= 1946 && year <= 1964:
		fmt.Println("Hello, boomer")
	case year >= 1965 && year <= 1980:
		fmt.Println("Hey, X")
	case year >= 1981 && year <= 1996:
		fmt.Println("Hey, milineal")
	case year >= 1997 && year <= 2012:
		fmt.Println("Hey, zoomer!")
	case year >= 2013:
		fmt.Println("Hey, alpha!")
	default:
		fmt.Println("Gutten Tag!")
	}
}

func timeToday() {
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today is Saturday")
	case today + 1:
		fmt.Println("Tommorow is Saturday")
	case today + 2:
		fmt.Println("In two days is Saturday")
	default:
		fmt.Println("Too far away.")
	}
}

func main() {
	timeToday()
	// var year = 1048
	// checkWhoYouAre(year)
}
