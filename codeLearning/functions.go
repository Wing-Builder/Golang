package main

import (
	"errors"
	"fmt"
)

// "math/rand"
// "time"

//Comment
func sayhello(name string) {
	fmt.Printf("Bonjour à tous, je m'appelle %v.\n", name)
}

func calculatePerimRect(lo, la float64) float64 {
	return 2 * (lo + la)
}

func sayBye(name string) (string, error) {
	if name == "" {
		return "", errors.New("\bErreur: Vous avez oublié de spécifier un nom!")
	}
	byeMessage := fmt.Sprintf("%v s'en va! Bonne soirée", name)
	return byeMessage, nil /* si erreur on renvoie une string vide et on déclenche
	l'erreur*/
}

func main() {
	sayhello("Alex")
	rl := calculatePerimRect(5.5, 2.4)
	fmt.Printf("le périmètre de mon rectangle est: %v.\n", rl)
	fmt.Println(sayBye(""))
}