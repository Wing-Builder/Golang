package main

import "fmt"

type Example struct {
	a string
	b int
	c bool
}

func main() {

	myContact := newContact("Alex", 30, map[string]string{"Fixe":"10.14.25.36.25", 
	"Portable":"17.54.14.25.62"})

	fmt.Println(myContact)
	// fmt.Println(displayContact())
	fmt.Println(myContact.displayContact())

	myNewContact := newContact("Justine", 27, map[string]string{"Fixe":"20.14.25.05.89", 
	"Portable":"34.54.96.58.74"})

	fmt.Println(myContact)
	// fmt.Println(displayContact())
	fmt.Println(myContact.displayContact())

	fmt.Println(myNewContact)
	fmt.Println(myNewContact.displayContact())

}
