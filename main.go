package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://google.com",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for {
		go checkLink(<-c, c)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be dead.")
		c <- link
		return
	}

	fmt.Println(link, "is up an runnin'!")
	c <- link
}
