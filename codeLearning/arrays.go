package main

import "fmt"

// "math/rand"
// "time"

//Comment

func main() {
	// var list [n]type
	var list [4]int

	list[0] = 50
	list[1] = 20
	list[2] = 10

	fmt.Println(list)
	fmt.Println(list[0])
	fmt.Println(list[1])
	fmt.Println(list[2])

	newList := [2]int {40, 50}

	fmt.Println(newList)
	fmt.Println(newList[0])
	fmt.Println(newList[1])

	bigList := [...]int {10, 20, 30, 40, 50, 69, 420, 777777, 50085}

	for pos, value := range bigList {
		fmt.Printf("Position %d est égal à %d.\n", pos, value)
	}


}