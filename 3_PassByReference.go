package main

import "fmt"

// struct inside a struct
type contactInfo struct {
	number int
	email  string
}

type Person struct {
	name        string
	age         int
	contactinfo contactInfo
}

func main() {
	kalios := Person{
		name: "kalios",
		age:  21,
		contactinfo: contactInfo{
			number: 790302,
			email:  "kalios@example.com",
		},
	}
	fmt.Println(kalios.name)

	// creating a pointer for kalios
	kaliosAdd := &kalios

	// trying to update name of kalios
	kaliosAdd.updateName("kalios cheese")
	fmt.Println(kalios.name)
}

// creating a method for type *Person
func (P *Person) updateName(name string) {
	// de-referencing and changing the value in it
	(*P).name = name
}
