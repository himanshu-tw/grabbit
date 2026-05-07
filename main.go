package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Welcome to Video Downloader.")

	var url string
	fmt.Println("Enter the URL")
	fmt.Scan(&url)

	fmt.Println("The URL you entered is :-", url)
}
