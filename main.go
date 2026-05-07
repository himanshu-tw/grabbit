package main

import (
	"fmt"
	"net/url"
	"log"
)

func main() {
	fmt.Println("Welcome to Video Downloader.")

	var urlInp string
	fmt.Println("Enter the URL")
	fmt.Scan(&urlInp)

	fmt.Println("The URL you entered is :-", urlInp)

	// checking if the url is correct
	u, err := url.Parse(urlInp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)
}
