package main

import "fmt"

var x int //globale

func main() {
	x = 5
	fmt.Println(x) // 5
	f()
	showY()
}

func f() {
	x := 10 //locale
	fmt.Println(x)
}
