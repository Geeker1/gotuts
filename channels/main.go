package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string {
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _,link := range links{
		go checkLink(link,c)
	}

	for l := range c {
		time.Sleep(5 * time.Second)
		go func(l string){
			time.Sleep(5 * time.Second)
			checkLink(l,c)
		}(l)
	}
}



func checkLink(link string,c chan string) {
	_, err := http.Get(link)
	if err != nil{
		fmt.Println(link, "mnight be down")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}