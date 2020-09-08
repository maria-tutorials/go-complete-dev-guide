package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://reddit.com",
		"http://mariainesserra.com",
		"http://golang.org",
		"http://talkdesk.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c { //loop using channel as range
		go checkLink(l, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down :(")
		c <- link
		return
	}

	fmt.Println(link, "is up :)")
	c <- link
}
