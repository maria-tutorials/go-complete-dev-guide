package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	//marie := person{"Marie", "Currie"}
	//marie := person{firstName: "Marie", lastName: "Currie"}
	var marie person
	marie.firstName = "Marie"
	marie.lastName = "Currie"

	fmt.Printf("%+v", marie)
}
