package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{"Alex", "Anderson"}
	fmt.Println(alex)

	var neel person
	neel.firstName = "Neel"
	neel.lastName = "Anderson"

	fmt.Println(neel)
	fmt.Printf("%+v", neel)

}
