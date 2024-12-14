package hangman

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

type HangManData struct {
	Word        string
	ToFind      string
	Attempts    int
	LettersUsed []string
}

func StartTheGame(wordFile, hangmanFile string) HangManData {
	words := readWordsFromFile(wordFile)
	rand.Seed(time.Now().UnixNano())
	randomWord := words[rand.Intn(len(words))]
	return HangManData{
		ToFind:   randomWord,
		Word:     HideWord(randomWord),
		Attempts: 10,
	}
}

func HideWord(randWord string) string {
	rand.Seed(time.Now().UnixNano())
	hiddenWord := []rune(randWord)
	randomletters := make(map[rune]bool)
	l := (len(randWord) / 2) - 1
	for len(randomletters) < l {
		index := rand.Intn(len(randWord))
		randomletters[rune(randWord[index])] = true
	}
	for l, letter := range randWord {
		if !randomletters[letter] {
			hiddenWord[l] = '_'
		}
	}
	return string(hiddenWord)
}

func JoseHang(hang *HangManData) {
	hangman := readWordsFromFile("hangman.txt")
	firstLine := 8
	hangLine := firstLine * (10 - hang.Attempts - 1)
	for i := 0; i < 7 && hangLine < len(hangman); i++ {
		if hangLine > len(hangman) || hang.Attempts == 10 {
			return
		}
		fmt.Println(hangman[hangLine][0:9])
		hangLine++
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

func readWordsFromFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return words
}