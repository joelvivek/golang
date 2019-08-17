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
	// we need not declare the struct
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 634573,
		},
	}

	//jim.updateName("jimmy")
	// Even though we are trying to update the name as Jimmy it is not reflecting
	//jim.print()

	jimPointer := &jim // assigning pointer to the memory address
	jimPointer.updateName("Jimmy")
	jimPointer.print()

}

// func (p person) updateName(newFirstName string) {
// 	p.firstName = newFirstName
// }

// *person = type description, working with pointer to a person
// *pointerToPerson = we want to manipulate the value of the pointer reffering
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
