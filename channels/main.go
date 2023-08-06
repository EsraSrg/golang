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

	c:=make(chan string) // creating channel.

	for _, link := range links {
		go checkLink(link,c) // go keyword for implementing go routines
	}
	
	for l:= range c {
		go func(){  //function literal

		time.Sleep(5 * time.Second)
		go checkLink(l,c)
		}()
	}
	
}

func checkLink(link string , c chan string) {
	_, err := http.Get(link) 
	if err != nil {
		fmt.Println(link + " might be down!")
		c <- link // sending data with channels
		return
	}
	fmt.Println(link, "is up")
	c <- link
}





