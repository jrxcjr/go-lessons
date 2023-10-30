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

	//This is the proper syntax to define and fulfill a struct.
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 12345,
		},
	}

	//This is the Third way of assigning values to a struct. Also proper. Similar to a PHP interface
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	alex.contactInfo.email = "trololo@mail.com"
	alex.contactInfo.zipCode = 12345

	//this is the pointer to the original address in value, of the original jim struct
	jimPointer := &jim

	//this function will pass as a receiver, the memory address of the original jim struct.
	jimPointer.updateName("Jimmy")
	jim.print()
}

// This receiver is a type description, it means that we're working with a pointer to a person.
// we pass the memory address to the receiver to use the original address inside the function.
func (pointerToPerson *person) updateName(newFirstName string) {

	//This is an operator, it means we want to manipulate the value that the pointer is referencing, that was received in the receiver.
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
