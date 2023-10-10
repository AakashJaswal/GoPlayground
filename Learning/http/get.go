package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type CopyToString struct{}

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	fmt.Println("Status code:", resp.Status)

	// res := make([]byte, 1000)
	// resp.Body.Read(res)
	// fmt.Println(string(res))
	cts := CopyToString{}
	io.Copy(cts, resp.Body)
}

func (CopyToString) Write(bs []byte) (n int, err error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
