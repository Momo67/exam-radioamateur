package main

import (
	"fmt"
	"math/rand"
	"time"
)

func choisirElement(tableau map[int]int) (int, int) {
	// Initialiser le générateur de nombres aléatoires
	rand.NewSource(time.Now().UnixNano())

	// Convertir les clés du dictionnaire (index) en une slice d'entiers
	indexes := make([]int, 0, len(tableau))
	for index := range tableau {
		indexes = append(indexes, index)
	}

	// Sélectionner un index de manière aléatoire
	indexChoisi := indexes[rand.Intn(len(indexes))]

	// Récupérer l'élément correspondant à l'index choisi
	elementChoisi := tableau[indexChoisi]

	return indexChoisi, elementChoisi
}

func main() {
	chapitres := map[int]int{
		1:  11,
		2:  74,
		3:  47,
		4:  107,
		5:  37,
		6:  29,
		7:  40,
		8:  22,
		9:  9,
		10: 48,
		11: 10,
		12: 4,
		13: 6,
		14: 19,
		15: 29,
	}

	// L'utilisateur entre le nombre d'itérations
	var nbIterations int
	fmt.Print("Entrez le nombre d'itérations: ")
	fmt.Scan(&nbIterations)
	
	for i := 0; i < nbIterations; i++ {
		
		// Choisir un chapitre aléatoirement
		noChapitre, nbQuestions := choisirElement(chapitres)
		// Choisir une question aléatoirement
		noQuestion := rand.Intn(nbQuestions) + 1

		fmt.Printf("Chapitre %d, question %d\n", noChapitre, noQuestion)
	}


}
