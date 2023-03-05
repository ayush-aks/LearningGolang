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
	for _, lnk := range links {
		checkLink(lnk)
	}
}

func checkLink(lnk string) {
	_, err := http.Get(lnk) // htt.Get -> Blocking code
	if err != nil {
		fmt.Println(lnk, " might be down!")
		return
	}

	fmt.Println(lnk, " is up.")
}

// The problem with this approach is that it is time consuming.
// This is because http.Get is waiting for response from each link - sequentially
// A faster and better approach would be to use channels and Go Routines (Making multiple requests)
// Refer channels folder for better approach.
// concurrency -> We can schedule work and change between them on the fly
// parallelism -> We can do multiple things at the same time.
// We only use go keyword(go routine) infront of function call.
// Channels are used to communicate in between different running go routines(Main Routine and Child Routines)
