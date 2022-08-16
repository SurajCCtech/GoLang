package model

//Person struct
type Person struct {
	Name string
	age  int
}

//GetAge of person
func (p Person) GetAge() int {
	return p.age
}

func (p Person) getName() string {
	return p.Name
}

type company struct {
}
