package main

import (
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"time"

	"github.com/eiannone/keyboard"
)

func choisirElement(tableau interface{}, r *rand.Rand) (noChapitre interface{}, nbQuestion interface{}) {
	switch t := tableau.(type) {
	case map[int]int:
		indexes := make([]int, 0, len(t))
		for index := range t {
			indexes = append(indexes, index)
		}
		indexChoisi := indexes[r.Intn(len(indexes))]
		elementChoisi := t[indexChoisi]
		return indexChoisi, elementChoisi

	case []int:
		indexChoisi := r.Intn(len(t))
		elementChoisi := t[indexChoisi]
		return elementChoisi, nil

	case []string:
		indexChoisi := r.Intn(len(t))
		elementChoisi := t[indexChoisi]
		return elementChoisi, nil

	default:
		log.Fatal("Type de tableau non supporté")
		return nil, nil
	}
}

func choisirQuestionsChapitre(quiz map[int]int, chapitre int, nbQuestions int, r *rand.Rand) []string {
	nbTotalQuestions, exists := quiz[chapitre]
	if !exists {
		log.Fatalf("Chapitre %d n'existe pas dans le quiz", chapitre)
	}

	combinaisonsGenerees := make(map[string]bool)
	var questions []string

	for i := 0; i < nbQuestions; i++ {
		var combinaison string

		for {
			noQuestion := r.Intn(nbTotalQuestions) + 1
			combinaison = fmt.Sprintf("%d.%d", chapitre, noQuestion)

			if !combinaisonsGenerees[combinaison] {
				combinaisonsGenerees[combinaison] = true
				break
			}
		}

		questions = append(questions, combinaison)
	}

	return questions
}

func choisirQuestionsTousChapitres(quiz map[int]int, nbQuestions int, r *rand.Rand) []string {
	combinaisonsGenerees := make(map[string]bool)
	var questions []string

	for i := 0; i < nbQuestions; i++ {
		var combinaison string

		for {
			chapitre, nbTotalQuestions := choisirElement(quiz, r)
			noQuestion := r.Intn(nbTotalQuestions.(int)) + 1
			combinaison = fmt.Sprintf("%d.%d", chapitre, noQuestion)

			if !combinaisonsGenerees[combinaison] {
				combinaisonsGenerees[combinaison] = true
				break
			}
		}

		questions = append(questions, combinaison)
	}

	return questions
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
	for {
		fmt.Print("Questionnaire désiré (o: OFCOM(2024), p: Prescriptions, d: DARC(2007)) ? ")
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

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

	var nbQuestions int
	fmt.Print("Nombre de questions ? ")
	fmt.Scan(&nbQuestions)

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	if reflect.DeepEqual(quiz, quizOFCOM) {
		fmt.Print("Souhaitez-vous sélectionner un chapitre spécifique ou tous les chapitres (s: spécifique, t: tous) ? ")
		choixChapitre, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("")

		if choixChapitre == 's' || choixChapitre == 'S' {
			var chapitre int
			fmt.Print("Chapitre désiré pour le quiz OFCOM ? ")
			fmt.Scan(&chapitre)

			nbTotalQuestions, exists := quizOFCOM[chapitre]
			if !exists {
				log.Fatalf("Chapitre %d n'existe pas dans le quiz OFCOM", chapitre)
			}

			if nbQuestions > nbTotalQuestions {
				log.Fatalf("Le nombre de questions demandé dépasse le nombre de questions disponibles dans le chapitre %d (disponibles: %d)", chapitre, nbTotalQuestions)
			}

			questions := choisirQuestionsChapitre(quizOFCOM, chapitre, nbQuestions, r)
			for _, question := range questions {
				fmt.Println(question)
			}
		} else if choixChapitre == 't' || choixChapitre == 'T' {
			var nbTotalQuestions int
			for _, nb := range quizOFCOM {
				nbTotalQuestions += nb
			}

			if nbQuestions > nbTotalQuestions {
				log.Fatalf("Le nombre de questions demandé dépasse le nombre total de questions disponibles (%d)", nbTotalQuestions)
			}

			questions := choisirQuestionsTousChapitres(quizOFCOM, nbQuestions, r)
			for _, question := range questions {
				fmt.Println(question)
			}
		} else {
			log.Fatal("Choix invalide")
		}
	} else {
		combinaisonsGenerees := make(map[string]bool)

		for i := 0; i < nbQuestions; i++ {
			var noQuestion interface{}
			var combinaison string

			for {
				noChapitre, nbQuestions := choisirElement(quiz, r)

				if nbQuestions != nil {
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

				if !combinaisonsGenerees[combinaison] {
					combinaisonsGenerees[combinaison] = true
					break
				}
			}

			fmt.Println(combinaison)
		}
	}
}
