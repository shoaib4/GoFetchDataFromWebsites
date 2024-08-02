package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func countOccurrences(content, word string) int {
	return strings.Count(content, word)
}

func main() {
	url := "https://golang.org"
	word := "Go"

	// Fetch content from the website
	resp, err := http.Get(url)
	//arr := make([]string, 0)
	if err != nil {
		fmt.Println("Error fetching the URL:", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: status code %d\n", resp.StatusCode)
		return
	}

	// Read the body content
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return
	}

	// Convert body to string
	content := string(body)

	//fmt.Printf(content)

	// Count occurrences of the word
	count := countOccurrences(content, word)
	fmt.Printf("The word '%s' occurs %d times in the content of %s\n", word, count, url)
}
