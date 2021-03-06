package main

import "fmt"

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
	marta := person{
		firstName: "Marta",
		lastName:  "Rodriguez",
		contactInfo: contactInfo{
			email:   "mr@mail.com",
			zipCode: 94000,
		},
	}

	marta.updateName("joe")
	marta.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
