package main

import "fmt"

type Example struct {
	a string
	b int
	c bool
}

func main() {
	myExample := Example{
		a: "Alex",
		b: 30,
		c: true,
	}

	// foolExample := Example{"Alex"}

	fmt.Println(myExample)
	// fmt.Println(foolExample)

	myContact := newContact("Alex", 30)

	fmt.Println(myContact)

}
