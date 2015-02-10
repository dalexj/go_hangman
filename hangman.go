package main

import (
	"fmt"
	"strings"
)

func isLetter(letter string) bool {
	return letter[0] >= 65 && letter[0] <= 90
}

func getInput(input *string) {
	fmt.Print("> ")
	fmt.Scanln(input)
	*input = strings.ToUpper(strings.TrimSpace(*input))
	if len(*input) != 1 || !isLetter(*input) {
		fmt.Println("Invalid input, please enter a single letter")
		getInput(input)
	}
}

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
	// fmt.Println("H _ _ _ _   _ _ _ _ _")
	// fmt.Println("H _ _ _ _   _ _ _ _ _")
	// fmt.Println("H E _ _ _   _ _ _ _ _")
	// fmt.Println("H E L L _   _ _ _ L _")
	// fmt.Println("H E L L O   _ O _ L _")
	// fmt.Println("H E L L O   W O _ L _")
	// fmt.Println("H E L L O   W O R L _")
	fmt.Println("YOU WIN!")
}
