package main

import (
	"fmt"
	"net/http"
	"sync"
)

const (
	// Version is the version of the client.
	Version = "0.0.1"
	url     = "http://localhost:8080/hello"
	totalRequests = 10000
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < totalRequests; i++ {
		wg.Add(1)
		go sendRequest(1, &wg)
	}

	wg.Wait()

	fmt.Println("All requests completed.")
}

func sendRequest(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Request %d failed: %v\n", id, err)
		return
	}

	fmt.Printf("Request %d succeeded: %s\n", id, resp.Status)
}