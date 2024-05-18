package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"time"
)

func test() {
	// Start a Selenium WebDriver server instance (using Chrome)
	opts := []selenium.ServiceOption{} // You can pass options if needed
	service, err := selenium.NewChromeDriverService("/path/to/chromedriver", 9515, opts...)
	if err != nil {
		panic(err)
	}
	defer service.Stop()

	// Connect to the WebDriver instance running locally
	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9515))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	// Navigate to Google
	if err := wd.Get("http://www.google.com"); err != nil {
		panic(err)
	}

	// More web interactions here...
	time.Sleep(5 * time.Second) // Let things render

}

func main() {
	const nihongo = "日本語"
	const sample = "фыва"
	// for i := 0; i < len(nihongo); i++ {
	//     fmt.Printf("%x ", nihongo[i])
	// }
	for idx, char := range sample {
		fmt.Printf("%# %d \n", char, idx)
	}

}
