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

	alex.updateFirstName("Randy")
	alex.print()
}

func (p *person) updateFirstName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	// Print key, value use Printf and '%+v'
	fmt.Printf("%+v", p)
}
