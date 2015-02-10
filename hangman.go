package main

import (
	"fmt"
	"strings"
)

// returns true if letter is uppercase A-Z
func isLetter(letter string) bool {
	return letter[0] >= 65 && letter[0] <= 90
}

// prompt user for a single letter(a-z/A-Z) until a valid input is given
// takes in a string pointer that the input is set to
func getInput(input *string) {
	fmt.Print("> ")
	fmt.Scanln(input)
	*input = strings.ToUpper(strings.TrimSpace(*input))
	if len(*input) != 1 || !isLetter(*input) {
		fmt.Println("Invalid input, please enter a single letter")
		getInput(input)
	}
}

// welcomes the player to hangman
func welcomeMessage() {
	fmt.Println("Welcome to hangman, here's your phrase to guess:")
}

func main() {
	var input string
	hangman := NewGame("hello world")
	hangman.GuessedLetters = "HEO"
	welcomeMessage()
	fmt.Println(hangman.FormatPhrase())
	getInput(&input)
	fmt.Println("YOU WIN!")
}
