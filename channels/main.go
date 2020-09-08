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

	fmt.Println(<-c)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down :(")
		c <- "Might be down I think"
		return
	}

	fmt.Println(link, "is up :)")
	c <- "Yup it's up"
}
