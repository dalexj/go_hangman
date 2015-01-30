package main

import (
	"fmt"
)

func main() {
	// var input string
	hangman := NewGame("hello world")
	hangman.GuessedLetters = "HEO"
	fmt.Println("Welcome to hangman, here's your phrase to guess:")
	fmt.Println(hangman.FormatPhrase())
	fmt.Println("H E L L O   W O R L D")
	// fmt.Print("> ")
	// fmt.Scanln(&input)
	// fmt.Println("H _ _ _ _   _ _ _ _ _")
	// fmt.Println("H _ _ _ _   _ _ _ _ _")
	// fmt.Println("H E _ _ _   _ _ _ _ _")
	// fmt.Println("H E L L _   _ _ _ L _")
	// fmt.Println("H E L L O   _ O _ L _")
	// fmt.Println("H E L L O   W O _ L _")
	// fmt.Println("H E L L O   W O R L _")
	fmt.Println("YOU WIN!")
}
