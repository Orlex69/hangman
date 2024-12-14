/*package hangman

import (
	"fmt"
)

func clearScreen() {

}

func ChooseLevel() string {
	var level string

	fmt.Println("Choose the difficulty : 1, 2 ou 3")
	fmt.Println("Level 1 : Easy")
	fmt.Println("Level 2 : Medium")
	fmt.Println("Level 3 : Hard")
	fmt.Scan(&level)

	var wordFile string
	switch level {
	case "1":
		wordFile = "data/word.txt1"
	case "2":
		wordFile = "data/word.txt2"
	case "3":
		wordFile = "data/word.txt3"
	default:
		wordFile = "data/word.txt1"
	}

	clearScreen()
	fmt.Println("Good luck ! If you want to save your game and quit, press 'stop' or 'STOP'")
	return wordFile
}