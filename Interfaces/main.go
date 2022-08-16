package main

import "fmt"

type animal interface {
	breath()
	walk()
}

type lion struct {
	age int
}

func (l *lion) breath() {
	fmt.Println("Lion breathes")
}

func (l *lion) walk() {
	fmt.Println("Lion walks")
}

func main() {
	var a animal
	a = &lion{age: 10}
	a.breath()
	a.walk()
}
