package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// Define any address you want to search
	address := "Sector 44 Delhi"

	// Encode the address for URL safety
	params := url.Values{}
	params.Add("q", address)

	// Build full URL
	fullURL := fmt.Sprintf("http://localhost:8080/search?%s", params.Encode())

	// Make GET request
	resp, err := http.Get(fullURL)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Print server response
	fmt.Println("Response from server:")
	fmt.Println(string(body))
}
