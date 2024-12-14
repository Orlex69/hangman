package hangman

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func QuitGame() {
	fmt.Println("Thanks for playing this game, see you next time ! :)")
	os.Exit(0)
}