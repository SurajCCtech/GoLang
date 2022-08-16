package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to the working with files in the golang")
	content := "Refer the website for golang learnings- https://go.dev/doc/effective_go"

	file, err := os.Create("./Refer.txt")
	// if err != nil {
	// 	panic(err)
	// }
	checkErr(err)

	length, err := io.WriteString(file, content)
	checkErr(err)
	fmt.Println("length of file: ", length)
	defer file.Close()
	Readfile("./Refer.txt")
}

func Readfile(filename string) {
	Bytedata, err := os.ReadFile(filename)
	checkErr(err)
	fmt.Println("Data of the file:\n ", string(Bytedata)) // if i write only the Bytedata it will just print the UTF code for the characters
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
