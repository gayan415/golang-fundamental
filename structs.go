package main

import (
	"fmt"
)

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// alex := person{"alex", "Anderson"}
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email: "gayan@bluedrop",
			zip:   1234,
		},
	}

	jim.updateName("Jimmy")
	jim.print()

	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// alex.contact.email = "gayan@bluedrop.com"
	// alex.contact.zip = 1234
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
