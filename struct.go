package main

import "fmt"

type contactInfo struct {
	zipCode int
	email   string
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func structUsage() {
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact: contactInfo{
			zipCode: 67678,
			email:   "alex@email.com",
		},
	}

	fmt.Println(alex)

	var betty person

	betty.firstName = "Betty"
	betty.lastName = "Periera"
	betty.contact.zipCode = 99887
	betty.contact.email = "bettyperiera@email.com"

	// Print key, value use PrintF and '%+v'
	fmt.Printf("%+v", alex)

}
