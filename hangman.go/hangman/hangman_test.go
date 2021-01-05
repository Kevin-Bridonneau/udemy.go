package hangman

import "testing"

func TestLetterInWord(t *testing.T) {
	word := []string{"w", "o", "o", "p", "e", "r"}
	guess := "w"
	hasLetter := letterInWord(guess, word)
	if !hasLetter {
		t.Errorf("Word %s contains letter %s. got = %v", word, guess, hasLetter)
	}
}

func TestLetterNotInWord(t *testing.T) {
	word := []string{"w", "o", "o", "p", "e", "r"}
	guess := "a"
	hasLetter := letterInWord(guess, word)
	if hasLetter {
		t.Errorf("Word %s does not contains letter %s. got = %v", word, guess, hasLetter)
	}
}

func TestInvalidWord(t *testing.T) {
	_, err := New(3, "")
	if err == nil {
		t.Errorf("Error should be returned when you using a invalid word = '' ")
	}
}

func TestGameGoodGuess(t *testing.T) {
	g, _ := New(3, "bob")
	g.MakeAGuess("b")
	validState(t, "goodGuess", g.State)
}
func TestGameBadGuess(t *testing.T) {
	g, _ := New(3, "bob")
	g.MakeAGuess("a")
	validState(t, "badGuess", g.State)
}
func TestGameWon(t *testing.T) {
	g, _ := New(3, "bob")
	g.MakeAGuess("b")
	g.MakeAGuess("o")
	validState(t, "won", g.State)
}
func TestGamelost(t *testing.T) {
	g, _ := New(3, "bob")
	g.MakeAGuess("a")
	g.MakeAGuess("v")
	g.MakeAGuess("e")
	g.MakeAGuess("b")
	g.MakeAGuess("b")
	validState(t, "lost", g.State)
}

func TestGameAlreadyGuessed(t *testing.T) {
	g, _ := New(3, "bob")
	g.MakeAGuess("a")
	g.MakeAGuess("a")
	validState(t, "alreadyGuessed", g.State)
}

func validState(t *testing.T, expectedState, actualState string) bool {
	if expectedState != actualState {
		t.Errorf("State should be '%v'. got=%v", expectedState, actualState)
		return false
	}
	return true
}
