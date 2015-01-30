package main

import (
	"strings"
)

type Game struct {
	Phrase         string
	GuessedLetters string
	Guesses        int
}

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

func NewGame(phrase string) Game {
	return Game{
		GuessedLetters: "",
		Guesses:        0,
		Phrase:         strings.ToUpper(phrase),
	}
}
