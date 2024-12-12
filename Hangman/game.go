package hangman

import (
	"fmt"
	"strings"
)

func CompareChar(hang *HangManData, input string) {
	if len(input) > 1 && (input >= "a" && input <= "z") {
		clearScreen()
		if input == hang.ToFind {
			hang.Word = hang.ToFind
			JoseHang(hang)
			fmt.Println("Félicitations, tu as trouvé le mot complet : ", hang.ToFind)
			fmt.Println("Tu as gagné avec", hang.Attempts, "vies restantes. Bravo !")
		} else if input != hang.ToFind && hang.Attempts > 2 {
			clearScreen()
			hang.Attempts -= 2
			fmt.Println("Bonne tentative, mais réessaie !")
			fmt.Println("Tentatives restantes : ", hang.Attempts)
			JoseHang(hang)
		} else {
			fmt.Println("Désolé, tu as perdu. Le mot à trouver était : ", hang.ToFind)
			fmt.Println("Tentatives restantes : 0")
			JoseHang(hang)
			QuitGame()
		}
	} else {
		FoundInWord := false
		if len(input) == 0 {
			return
		}
		input1 := rune(input[0])
		for z, i := range hang.ToFind {
			if i == input1 {
				hang.Word = hang.Word[:z] + string(i) + hang.Word[z+1:]
				FoundInWord = true
			}
		}

		if !FoundInWord {
			clearScreen()
			for _, i := range hang.LettersUsed {
				if i == input {
					fmt.Println("Tu as déjà utilisé cette lettre, essaie une autre !")
					JoseHang(hang)
					fmt.Println("Lettres déjà utilisées :", hang.LettersUsed)
					return
				}
			}
			hang.Attempts--
			hang.LettersUsed = append(hang.LettersUsed, input)
			fmt.Println("Désolé, cette lettre n'est pas dans le mot.")
			fmt.Println("Tentatives restantes : ", hang.Attempts)
			JoseHang(hang)
			if hang.Attempts == 0 {
				fmt.Println("Désolé, tu as perdu. Le mot à trouver était :", hang.ToFind)
			}
		} else {
			clearScreen()
			for _, i := range hang.LettersUsed {
				if i == input {
					fmt.Println()
					fmt.Println("Tu as déjà utilisé cette lettre, essaie une autre !")
					JoseHang(hang)
					fmt.Println("Lettres déjà utilisées :", hang.LettersUsed)
					return
				}
			}
			hang.LettersUsed = append(hang.LettersUsed, input)
			fmt.Println("Bien joué, cette lettre est dans le mot !")
			fmt.Println("Tentatives restantes : ", hang.Attempts)
			JoseHang(hang)
			if hang.Word == hang.ToFind {
				fmt.Println("Félicitations, tu as trouvé le mot : ", hang.ToFind)
			}
		}
	}
}

func Input(hang *HangManData) string {
	var input string
	var result string
	var reponse int

	fmt.Print("Choisis une lettre ou un mot : ")
	fmt.Scanf("%s", &input)
	input = strings.ToLower(input)
	if input >= "a" && input <= "z" && len(input) == 1 {
		result = input
		fmt.Println()
	} else if (input < "a" || input > "z") && (input < "A" || input > "Z") {
		clearScreen()
		fmt.Print("Erreur : tu dois taper une lettre.")
		fmt.Println()
		JoseHang(hang)
		PlayGame(hang)
		result = ""
	} else if input >= "a" && input <= "z" && len(input) > 1 {
		result = input
	}
	if input == "stop" {
		input = strings.ToUpper(input)
		fmt.Println("Veux-tu sauvegarder ta partie ? Si oui, appuie sur '1', sinon appuie sur '2'. Par défaut, ta partie sera sauvegardée.")
		fmt.Scan(&reponse)
		switch reponse {
		case 1:
			save(hang)
			QuitGame()
		case 2:
			QuitGame()
		default:
			save(hang)
			QuitGame()
		}
	}
	return result
}

func PlayGame(hang *HangManData) {
	for hang.Attempts > 0 && hang.Word != hang.ToFind {
		fmt.Println("Mot actuel : ", hang.Word)
		input := Input(hang)
		CompareChar(hang, string(input))
	}
}