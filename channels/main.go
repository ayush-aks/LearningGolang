// concurrency -> We can schedule work and change between them on the fly
// parallelism -> We can do multiple things at the same time.
// We only use go keyword(go routine) infront of function call.
// Channels are used to communicate in between different running go routines(Main Routine and Child Routines)
package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// channel should communicate with same data type
	c := make(chan string)

	// In order to communicate with main and child routine, argument inside function(channel) is used

	for _, lnk := range links {
		go checkLink(lnk, c)
	}

	//fmt.Println(<-c) is also a blocking code :>

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(lnk string, c chan string) {
	_, err := http.Get(lnk) // http.Get -> Blocking code
	if err != nil {
		fmt.Println(lnk, " might be down!")
		c <- "Link might be down"
		return
	}

	fmt.Println(lnk, " is up.")
	c <- "Link is up"
}
