package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs"

func main() {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myurl)

	// Parsing the URL
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	queryParameters := result.Query()
	fmt.Printf("The type of the query parameter is: %T\n", queryParameters)

	fmt.Println(queryParameters["coursename"])

	// Creating the URL
	partsURL := &url.URL{
		Scheme:   "https",
		Host:     "suraj.com",
		Path:     "/tutcss",
		RawQuery: "user=Suraj",
	}
	another := partsURL.String()
	fmt.Println(another)

}
