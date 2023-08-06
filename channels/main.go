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

	c:=make(chan string) // creating channel.

	for _, link := range links {
		go checkLink(link,c) // go keyword for implementing go routines
	}
	
	for i:=0; i<len(links); i++{
		fmt.Println(<-c)
	}
	
}

func checkLink(link string , c chan string) {
	_, err := http.Get(link) 
	if err != nil {
		fmt.Println(link + " might be down!")
		c <- "might be down ı think" // sending data with channels
		return
	}
	fmt.Println(link, "is up")
	c <- "yep,its up"
}





