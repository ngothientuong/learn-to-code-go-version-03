package main

import (
	"bufio"
	"fmt"
	"net/http"
)

// This function send a request to an url and get response status and response body back
func main() {

	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	// for i := 0; scanner.Scan() && i < 5; i++ {
	for i := 0; scanner.Scan(); i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
