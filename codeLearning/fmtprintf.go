package main

import "fmt"

func main() {
	// //Amener à disparaître dans le futur
	// println("Salut à tous!")
	// print("Salut à tous!")

	// //Utiliser fmt
	// fmt.Println("Salut à tous!")
	// fmt.Print("Salut à tous!")

	//Printf

	x, y := 50, 35

	fmt.Printf("Salut à tous!\n")
	fmt.Printf("Mon nombre (default) est: %v\n", x)
	fmt.Printf("Mon nombre (base 2) est: %b\n", x)
	fmt.Printf("Mon nombre (base 8) est: %O\n ", x)
	fmt.Printf("Mon nombre (base 10) est: %v\n", x)
	fmt.Printf("Mon nombre est (base 16): %X\n", x)

	fmt.Printf("La valeur de X est égal à la valeur Y -> %t\n", x > y)

	fmt.Printf("L'écriture scientifique de 1258.16559 > %E", 1258.16559)
}
