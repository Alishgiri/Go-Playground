package main

import (
	"fmt"
	"net/http"
	"time"
)

func routinesAndChannelsExample() {
	links := []string{
		"http://amazon.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://google.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for i := 0; i < len(links); i++ {

	// This is equivalent to the for loop BELOW
	// for {
	// 	go checkLink(<-c, c)
	// }

	// This is equivalent to the for loop ABOVE
	for l := range c {
		go func() {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}()
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down!")
		c <- link
		return
	}
	fmt.Println(link, " is UP!")
	c <- link
}
