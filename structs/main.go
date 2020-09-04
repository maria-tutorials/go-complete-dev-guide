package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	marta := person{
		firstName: "Marta",
		lastName:  "Rodriguez",
		contact: contactInfo{
			email:   "mr@mail.com",
			zipCode: 94000,
		},
	}
	fmt.Printf("%+v", marta)
}
