package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/eiannone/keyboard"
)

func choisirElement(tableau interface{}, r *rand.Rand) (noChapitre interface{}, nbQuestion interface{}) {
	switch t := tableau.(type) {
	case map[int]int:
		// Convertir les clés du dictionnaire (index) en une slice d'entiers
		indexes := make([]int, 0, len(t))
		for index := range t {
			indexes = append(indexes, index)
		}

		// Sélectionner un index de manière aléatoire
		indexChoisi := indexes[r.Intn(len(indexes))]

		// Récupérer l'élément correspondant à l'index choisi
		elementChoisi := t[indexChoisi]

		return indexChoisi, elementChoisi

	case []int:
		// Sélectionner un index de manière aléatoire
		indexChoisi := r.Intn(len(t))

		// Récupérer l'élément correspondant à l'index choisi
		elementChoisi := t[indexChoisi]

		// Retourner l'index choisi et le nombre de questions
		return elementChoisi, nil

	case []string:
		// Sélectionner un index de manière aléatoire
		indexChoisi := r.Intn(len(t))

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

	quizPrescription := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
		41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60,
		61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80,
		81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100,
		101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116,
	}

	var quiz interface{}

	// Initialiser le lecteur de clavier
	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	for {
		// L'utilisateur choisit le questionnaire
		fmt.Print("Questionnaire désiré ? (o: OFCOM(2024), p: Prescriptions, d: DARC(2007)): ")
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		// Vérifier si l'utilisateur a entré un questionnaire valide
		if char == 'o' || char == 'O' {
			quiz = quizOFCOM
			break
		} else if char == 'p' || char == 'P' {
			quiz = quizPrescription
			break
		} else if char == 'd' || char == 'D' {
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

	// Initialiser le générateur de nombres aléatoires avec l'heure actuelle
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// Map pour stocker les combinaisons déjà générées
	combinaisonsGenerees := make(map[string]bool)

	for i := 0; i < nbQuestions; i++ {
		var noQuestion interface{}
		var combinaison string

		for {
			// Choisir un chapitre aléatoirement
			noChapitre, nbQuestions := choisirElement(quiz, r)

			if nbQuestions != nil {
				// Choisir une question aléatoirement
				noQuestion = r.Intn(nbQuestions.(int)) + 1
				combinaison = fmt.Sprintf("%v.%d", noChapitre, noQuestion)
			} else {
				switch v := noChapitre.(type) {
				case int:
					combinaison = fmt.Sprintf("%d", v)
				case string:
					combinaison = fmt.Sprintf("%s", v)
				default:
					log.Fatal("Type non supporté")
				}
			}

			// Vérifier si la combinaison a déjà été générée
			if !combinaisonsGenerees[combinaison] {
				combinaisonsGenerees[combinaison] = true
				break
			}
		}

		// Afficher la combinaison générée
		fmt.Println(combinaison)
	}
}
