package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	var ayush person
	ayush.firstName = "Ayush"
	ayush.lastName = "Shrivastav"
	ayush.contactInfo.email = "ayush.kumar-shrivastav@thalesgroup.com"
	ayush.contactInfo.zipCode = 202020
	// aksPointer := &ayush
	// aksPointer.updateName("abcd")
	// Shortcut:- No need to create above pointer, Go can directly pass the reference when calling updateName with pointer
	ayush.updateName("abcd")
	ayush.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
