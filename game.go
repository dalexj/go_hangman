package main

import (
	"strings"
)

type Game struct {
	Phrase         string
	GuessedLetters string
	Guesses        int
}

// returns the Phrase of the game in a pretty format
// where guessed letters are shown and unguessed letters
// are left as underscores
func (g Game) FormatPhrase() string {
	formatted := ""
	for i, _ := range g.Phrase {
		if string(g.Phrase[i]) == " " {
			formatted += " "
		} else if strings.ContainsAny(string(g.Phrase[i]), g.GuessedLetters) {
			formatted += string(g.Phrase[i])
		} else {
			formatted += "_"
		}
		formatted += " "
	}
	return formatted
}

// creates a new Game struct, uppercases the phrase given
func NewGame(phrase string) Game {
	return Game{
		GuessedLetters: "",
		Guesses:        0,
		Phrase:         strings.ToUpper(phrase),
	}
}
