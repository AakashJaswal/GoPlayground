package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status code:", resp.Status)

	res := make([]byte, 1000)
	resp.Body.Read(res)
	fmt.Println(string(res))
}
