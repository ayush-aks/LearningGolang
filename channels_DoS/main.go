package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, lnk := range links {
		go checkLink(lnk, c)
	}

	// Function literal in Go -> Anonymous function/ Unnamed function

	for l := range c {
		// go checkLink(l, c) -> Run for continuous call instead of sleep.
		go func(lnk string) {
			time.Sleep(2 * time.Second)
			checkLink(lnk, c)
		}(l)
	}
}

func checkLink(lnk string, c chan string) {
	_, err := http.Get(lnk) // http.Get -> Blocking code
	if err != nil {
		fmt.Println(lnk, " might be down!")
		c <- lnk
		return
	}

	fmt.Println(lnk, " is up.")
	c <- lnk
}
