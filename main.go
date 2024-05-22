package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/eiannone/keyboard"
)

func choisirElement(tableau interface{}) (noChapitre interface{}, nbQuestion interface{}) {
	// Initialiser le générateur de nombres aléatoires
	rand.NewSource(time.Now().UnixNano())

	switch t := tableau.(type) {
	case map[int]int:
		// Convertir les clés du dictionnaire (index) en une slice d'entiers
		indexes := make([]int, 0, len(t))
		for index := range t {
			indexes = append(indexes, index)
		}

		// Sélectionner un index de manière aléatoire
		indexChoisi := indexes[rand.Intn(len(indexes))]

		// Récupérer l'élément correspondant à l'index choisi
		elementChoisi := t[indexChoisi]

		return indexChoisi, elementChoisi

	case []string:
		// Sélectionner un index de manière aléatoire
		indexChoisi := rand.Intn(len(t))

		// Récupérer l'élément correspondant à l'index choisi
		elementChoisi := t[indexChoisi]

		// Retourner l'index choisi et le nombre de questions
		return elementChoisi, nil

	default:
		log.Fatal("Type de tableau non supporté")
		return nil, nil
	}	
}

func main() {
	quizOFCOM := map[int]int{
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

	quizDARC := []string{
		"TA104", "TA118", "TB103", "TB104", "TB106", "TB107", "TB203", "TB204", "TB205", "TB303",
		"TB402", "TB405", "TB504", "TB506", "TB510", "TB511", "TB512", "TB605", "TB612", "TB910",
		"TB913", "TB919", "TC101", "TC203", "TC207", "TC302", "TC314", "TC401", "TC406", "TC407",
		"TC516", "TC519", "TC520", "TC618", "TC621", "TC623", "TC709", "TC716", "TD119", "TD211",
		"TD212", "TD309", "TD312", "TD313", "TD318", "TD406", "TD419", "TD424", "TD605", "TD608",
		"TE209", "TE313", "TF204", "TF320", "TG101", "TG111", "TG206", "TH105", "TH157", "TH410",
		"TJ703", "TJ802", "TJ803", "TJ807", "TK114", "TK212", "TK215",
	}

	var quiz interface{}

	// Initialiser le lecteur de clavier
	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	for {
		// L'utilisateur choisit le questionnaire
		fmt.Print("Questionnaire désiré ? (o: OFCOM(2024), d: DARC(2007)): ")
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		// Vérifier si l'utilisateur a entré un questionnaire valide
		if char == 'o' || char == 'O' {
			quiz = quizOFCOM
			break
		}	else if char == 'd' || char == 'D' {
			quiz = quizDARC
			break
		} else {
			fmt.Println("Questionnaire invalide. Veuillez réessayer.")
		}

	}

	fmt.Println("")

	// L'utilisateur entre le nombre de questions désirées
	var nbQuestions int
	fmt.Print("Nombre de questions ? : ")
	fmt.Scan(&nbQuestions)
	
	for i := 0; i < nbQuestions; i++ {
		
		// Choisir un chapitre aléatoirement
		noChapitre, nbQuestions := choisirElement(quiz)

		if nbQuestions != nil {

			// Choisir une question aléatoirement
			noQuestion := rand.Intn(nbQuestions.(int)) + 1
			//fmt.Printf("Chapitre %v, question %d\n", noChapitre, noQuestion)
			fmt.Printf("%v.%d\n", noChapitre, noQuestion)
		} else {
			fmt.Printf("%s\n", noChapitre.(string))
		}

	}

}
