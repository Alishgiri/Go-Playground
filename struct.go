package main

import "fmt"

type contactInfo struct {
	zipCode int
	email   string
}

type person struct {
	firstName string
	lastName  string

	contactInfo
	// equivalent to
	// contactInfo contactInfo
}

func structUsage() {
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			zipCode: 67678,
			email:   "alex@email.com",
		},
	}

	fmt.Println(alex)

	var betty person

	betty.firstName = "Betty"
	betty.lastName = "Periera"
	betty.contactInfo.zipCode = 99887
	betty.contactInfo.email = "bettyperiera@email.com"

	// &alex means give the RAM memory address of 'alex'
	// Turn value into address using &
	alexPointer := &alex
	alexPointer.updateFirstName("Alexa")
	alex.print()
}

// *person (person is type here) means we are working with pointers.
// Turn address into value using *
func (pointerToPerson *person) updateFirstName(newFirstName string) {
	// *pointerToPerson (pointerToPerson is a variable here) means give me the value this memory address is pointing at
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	// Print key, value use Printf and '%+v'
	fmt.Printf("%+v", p)
}
