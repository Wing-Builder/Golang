package main

import "fmt"

func main() {
	// supermarketPrice := map[string]int{
	// 	"prince": 8,
	// 	"eau": 2,
	// 	"viande": 6,
	// }

	// fmt.Println(supermarketPrice["prince"])

	// for key, value := range supermarketPrice{
	// 	fmt.Println(key, value)
	// }

	/**********exo***********/

	/*Mettre les mois dans une map avec le nombre de jour associé
	Indiquer le nombre de jours dans une année en passant par la map*/

	// fmt.Println(Months["Janvier"+"Février"+"Mars"+"Avril"+
	// "Mai"+"Juin"+"Juillet"+"Août"+"Septembre"+"Octobre"+"Novembre"+
	// "Décembre"])

	// fmt.Println(Months["Janvier"])

	daysInAYear := 0

	daysInAMonth := map[string]int{
		"Janvier":   31,
		"Février":   28,
		"Mars":      31,
		"Avril":     30,
		"Mai":       31,
		"Juin":      30,
		"Juillet":   31,
		"Août":      31,
		"Septembre": 30,
		"Octobre":   31,
		"Novembre":  30,
		"Décembre":  31,
	}

	for _, value := range daysInAMonth {
		daysInAYear = daysInAYear + value
	}

	fmt.Printf("Nombre de jours dans une année: %d jours!\n", daysInAYear)

	for key, value := range daysInAMonth {
		fmt.Printf("Le mois de %v possède %d jours!\n", key, value)
	}
}
