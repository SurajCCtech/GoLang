package main

import "fmt"
import "MoreOnMethods/model"

//Test function
func Test() {
	//STRUCTURE IDENTIFIER
	p := model.Person{
		Name: "test",
		age:  21,
	}
	// fmt.Println(p)
	// c := model.company{}
	// fmt.Println(c)

	//STRUCTURE'S FIELDS
	fmt.Println(p.Name)
	fmt.Println(p.age)

	//STRUCTURE'S METHOD
	fmt.Println(p.GetAge())
	fmt.Println(p.getName())

}
