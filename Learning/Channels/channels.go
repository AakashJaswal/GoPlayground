package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	links := []string{
		"http://aws.com",
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
	}

	c := make(chan string)
	for _, link := range links {
		go getHttpStatus(link, c)

	}
	for i := range c {
		go func(i string) {
			time.Sleep(2 * time.Second)
			go getHttpStatus(i, c)
		}(i)
	}
}

func getHttpStatus(l string, c chan string) {
	resp, err := http.Get(l)

	if err != nil {
		fmt.Println("Error:", err)
		c <- l + ": " + err.Error()
		os.Exit(1)
	}

	fmt.Println(l + ": " + resp.Status)
	// c <- l + ": " + resp.Status
	c <- l
}
