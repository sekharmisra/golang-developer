package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email       string
	zipCode     int
	phoneNumber int
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:       "whatever@whatever.com",
			zipCode:     94087,
			phoneNumber: 12345678790,
		},
	}
	jim.updateName("Dimmy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newName string) {
	(*pointerToPerson).firstName = newName
}
