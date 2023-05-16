package main

import "fmt"

func updateA(number int) int {
	number = 5
	return number
}

func updateB(item map[string]int) {
	item["Bonbon"] =  4
	item["tournevis"] = 7
}

func main() {
	// A: string, int, bool, float, array
	number := 10
	number = updateA(number)
	fmt.Println(number) //5

	// B: maps, functions
	groceryList := map[string]int {
		"prince": 6,
		"viande": 10,
	}

	updateB(groceryList)
	fmt.Println(groceryList)
}
