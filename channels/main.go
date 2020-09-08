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

	for _, link := range links {
		go checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down :(")
		return
	}

	fmt.Println(link, "is up :)")
}
