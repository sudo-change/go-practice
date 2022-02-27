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

	// trying to update name of kalios
	kalios.updateName("kalios cheese")
	fmt.Println(kalios.name)
}

// creating a method for Person
func (P Person) updateName(name string) {
	P.name = name
}
