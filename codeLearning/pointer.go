package main

import "fmt"

func updateA(pointerOfNumber *int, value int) {
	*pointerOfNumber = value
}

func main() {
	// A: string, int, bool, float, array
	number := 10
	fmt.Println(number)// 10
	fmt.Println("Adresse mémoire de number:", &number)
	
	myPointer := &number // pointer
	fmt.Println(myPointer)
	fmt.Printf("La valeur de l'adresse mémoire %v est %d.\n", myPointer, *myPointer)
	
	updateA(myPointer, 20)
	fmt.Printf("La valeur de l'adresse mémoire %v a changé -> %d", myPointer, number)
}