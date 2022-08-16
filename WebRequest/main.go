package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Lco print the web request")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The type of the response is: %T\n", response)

	defer response.Body.Close()
	bytedata, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(bytedata)

	fmt.Println("The body of the response: ", content)

}
