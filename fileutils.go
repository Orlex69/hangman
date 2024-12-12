package hangman

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func save(hang *HangManData) {
	partie, err1 := json.Marshal(hang)
	if err1 != nil {
		log.Fatal(err1)
	}
	err2 := os.WriteFile("save1.txt", partie, 0644)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("Your game has been saved in save1.txt")
}

func QuitGame() {
	fmt.Println("Thanks for playing this game, see you next time ! :)")
	os.Exit(0)
}