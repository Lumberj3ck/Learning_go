package main

import (
	"fmt"
	"time"
)

func waitForDBs(numDBs int, dbChan chan struct{}) {
	for i := 0; i < numDBs; i++ {
		<-dbChan
	}
}

func getDBsChannel(numDBs int) (chan struct{}, *int) {
	count := 0
	ch := make(chan struct{})

	go func() {
		for i := 0; i < numDBs; i++ {
			ch <- struct{}{}
			fmt.Printf("Database %v is online\n", i+1)
			count++
		}
	}()

	return ch, &count
}

func addEmailsToQueue1(emails []string) chan struct{} {
	ch := make(chan struct{})
	for _, letter := range emails {
		fmt.Println(letter)
		ch <- struct{}{}
	}
	return ch
}

func addEmailsToQueue(emails []string) chan string {
	// buffCap := len(emails)
	ch := make(chan string)
	go func() {
		for true {
			<-ch
		}
	}()
	for _, letter := range emails {
		fmt.Println(letter)
		ch <- letter
	}
	// <- ch
	return ch
}

func testDeadlock() {
	// go func(){
	//     ch := make(chan struct{})
	//     ch <- struct{}{}
	// }()
	// ch <- struct{}{}
	ch := make(chan struct{})
	go func() {
		time.Sleep(time.Second * 3)
		ch <- struct{}{}
	}()
	<-ch
	fmt.Println("Block resolved")
}

func saveBackups(snapshotTicker, saveAfter <-chan time.Time, logChan chan string) {
	for {
		select {
		case <-snapshotTicker:
			go takeSnapshot(logChan)
		case <-saveAfter:
			go saveSnapshot(logChan)
			return
		default:
			go waitForData(logChan)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func takeSnapshot(logChan chan string) {
	logChan <- "Taking a backup snapshot..."
}

func saveSnapshot(logChan chan string) {
	logChan <- "All backups saved!"
	close(logChan)
}

func waitForData(logChan chan string) {
	logChan <- "Nothing to do, waiting..."
}

func main() {
	snapshotTicker := time.Tick(time.Millisecond * 400)
	saveAfter := time.After(time.Millisecond * 2800)
	logChan := make(chan string)
	go saveBackups(snapshotTicker, saveAfter, logChan)
	for log := range logChan {
		fmt.Println(log)
	}
	// testDeadlock()
	// fmt.Println("After block")
	// addEmailsToQueue([]string{"asdf", "shit"})
	// <- ch
	// time.Sleep(time.Second * 3)
	// getDBsChannel(10)
	// waitForDBs(10, ch)
}
