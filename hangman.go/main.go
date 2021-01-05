package main

import (
	"fmt"
	"os"

	"udemy.go/hangman.go/dictionary"
	"udemy.go/hangman.go/hangman"
)

func main() {

	err := dictionary.Load("words.txt")

	if err != nil {
		fmt.Printf("Could not load dictionary: %v", err)
		os.Exit(1)
	}

	g := hangman.New(8, dictionary.PickWord())

	hangman.DrawWelcome()

	guess := ""

	for {
		hangman.Draw(g, guess)

		switch g.State {
		case "won", "lost":
			os.Exit(1)

		}
		l, err := hangman.ReadGuess()
		if err != nil {
			fmt.Printf("Could not read from terminal: %v\n", err)
			os.Exit(1)
		}
		guess = l

		g.MakeAGuess(guess)
	}

}
