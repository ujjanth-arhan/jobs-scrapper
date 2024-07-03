package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Worker started...")
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	defer resp.Body.Close()

	page, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(string(page))

	fmt.Println("\nWorker finished...")
}
