package main

import (
	"time"
)

type email struct {
	body string
	date time.Time
}

func checkEmailAge(emails [3]email) {
	// go func(){
	//     isOld := [3]bool{}
	//     isOld[0] = <-isOldChan
	//     isOld[1] = <-isOldChan
	//     isOld[2] = <-isOldChan
	// }()
	sendIsOld(emails)

	// isOld := [3]bool{}
	// for i := range isOldChan{
	//     isOld[d] = i
	// }

	// return isOld
}

// don't touch below this line

func sendIsOld(emails [3]email) {
	isOldChan := make(chan bool, 10)
	for _, e := range emails {
		if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
			isOldChan <- true
			continue
		}
		isOldChan <- false
	}
}

func main() {
	checkEmailAge([3]email{{"asdf", time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC)}, {"asdf", time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC)}, {"asdf", time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC)}})
}
