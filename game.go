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

func InputHandler(w http.ResponseWriter, r *http.Request, game *HangManData, renderTemplate func(http.ResponseWriter, string, interface{})) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Erreur lors du traitement du formulaire.", http.StatusBadRequest)
		return
	}

	input := r.Form.Get("input")
	if input == "" {
		http.Error(w, "Aucune entrée reçue.", http.StatusBadRequest)
		return
	}

	// Convertir l'entrée en minuscules
	input = strings.ToLower(input)

	// Gestion du jeu
	CompareChar(game, input)

	// Préparer les données pour le template
	data := struct {
		Word      string
		Attempts  int
		Letters   []string
		HangState string
	}{
		Word:      game.Word,
		Attempts:  game.Attempts,
		Letters:   game.LettersUsed,
		HangState: "Continuez à deviner !",
	}

	if game.Attempts == 0 {
		data.HangState = "Désolé, vous avez perdu. Le mot était : " + game.ToFind
	} else if game.Word == game.ToFind {
		data.HangState = "Félicitations, vous avez trouvé le mot : " + game.ToFind
	}

	// Renvoyer les données au template HTML
	renderTemplate(w, "Game.html", data)
}


func PlayGame(hang *HangManData) {
	for hang.Attempts > 0 && hang.Word != hang.ToFind {
		fmt.Println("Mot actuel : ", hang.Word)
		input := Input(hang)
		CompareChar(hang, string(input))
	}
}