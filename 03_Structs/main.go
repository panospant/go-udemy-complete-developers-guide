package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstname string
	lastname  string
	contactInfo
}

func main() {
	// var panos = person{"Panos", "Pantou"}
	var panos = person{firstname: "Panos", lastname: "Pantou"}
	fmt.Println(panos)

	var alex person
	alex.firstname = "Alex"
	alex.lastname = "Anderson"

	// fmt.Println(alex)
	//fmt.Printf("%+v", alex)

	jim := person{
		firstname: "Jim",
		lastname:  "Morrison",
		contactInfo: contactInfo{
			email:   "well...",
			zipcode: 666,
		},
	}

	jimPointer := &jim

	// the update doesn't really update the value
	jimPointer.updateName("Jimmy")

	jim.print()

	// here we got the pointer (memory address) of our value
	jim.updateName("Joe")

	jim.print()
}

// *person means that this updateName can be called by a pointer of type person
// *jimPointer actually accesses the value of the pointer
func (jimPointer *person) updateName(newFirstName string) {
	(*jimPointer).firstname = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
