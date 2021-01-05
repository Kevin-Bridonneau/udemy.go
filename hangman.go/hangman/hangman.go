package hangman

import (
	"fmt"
	"strings"
)

type Game struct {
	State        string   // Game state
	Letters      []string // letters in a word to find
	FoundLetters []string // Good guesses
	UsedLetters  []string // Used letters
	TurnsLeft    int      // remaining attemps
}

func New(turns int, word string) (*Game, error) {
	if len(word) < 3 {
		return nil, fmt.Errorf("Word '%s' must be at least 3 characters. got = %v", word, len(word))
	}
	letters := strings.Split(strings.ToUpper(word), "")
	found := make([]string, len(letters))
	for i := 0; i < len(letters); i++ {
		found[i] = "_"
	}
	g := &Game{
		State:        "",
		Letters:      letters,
		FoundLetters: found,
		UsedLetters:  []string{},
		TurnsLeft:    turns,
	}
	return g, nil
}

func (g *Game) MakeAGuess(guess string) {
	guess = strings.ToUpper(guess)

	switch g.State {
	case "won", "lost":
		return
	}

	if letterInWord(guess, g.UsedLetters) {
		g.State = "alreadyGuessed"
	} else if letterInWord(guess, g.Letters) {
		g.State = "goodGuess"
		g.RevealLetter(guess)
		if hasWon(g.Letters, g.FoundLetters) {
			g.State = "won"
		}
	} else {
		g.UsedLetters = append(g.UsedLetters, guess)
		g.State = "badGuess"
		g.TurnsLeft--

		if g.TurnsLeft <= 0 {
			g.State = "lost"
		}
	}

}

func hasWon(letters []string, foundletters []string) bool {
	for i := range letters {
		if letters[i] != foundletters[i] {
			return false
		}
	}
	return true
}

func (g *Game) RevealLetter(guess string) {
	g.UsedLetters = append(g.UsedLetters, guess)
	for i, l := range g.Letters {
		if guess == l {
			g.FoundLetters[i] = guess
		}
	}
}

func letterInWord(guess string, letters []string) bool {
	for _, l := range letters {
		if l == guess {
			return true
		}
	}
	return false
}
