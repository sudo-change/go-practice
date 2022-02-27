package main

import "fmt"

// when functions went into classes we call them as methods
// in golang we dont have classes
type Person struct {
	name  string
	age   int
	email string
}

func main() {
	kalios := Person{
		name:  "kalios",
		age:   21,
		email: "kalios@example.com",
	}
	kalios.WStart("Mastery")

}

func (P Person) WStart(text string) {
	fmt.Printf("to feel %v in one thing in the next 2 years i.e within %v \n", text, (P.age)+2)
}
