// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Enough Picks
//
//  If the player's guess number is below 10;
//  then make sure that the computer generates a random
//  number between 0 and 10.
//
//  However, if the guess number is above 10; then the
//  computer should generate the numbers
//  between 0 and the guess number.
//
// WHY?
//  This way the game will be harder.
//
//  Because, in the current version of the game, if
//  the player's guess number is for example 3; then the
//  computer picks a random number between 0 and 3.
//
//  So, currently a player can easily win the game.
//
// EXAMPLE
//  Suppose that the player runs the game like this:
//    go run main.go 9
//
//  Or like this:
//    go run main.go 2
//
//    Then the computer should pick a random number
//    between 0-10.
//
//  Or, if the player runs it like this:
//    go run main.go 15
//
//    Then the computer should pick a random number
//    between 0-15.
// ---------------------------------------------------------

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("No input, give me a number between 0 and 9, -v for verbose mode")
		return
	}

	verbose := false
	if args[0] == "-v" {
		verbose = true
	}

	userGuess, err := strconv.Atoi(args[len(args)-1])
	if err != nil {
		fmt.Printf("Invalid number %q", args[len(args)-1])
		return
	}

	if userGuess < 0 {
		fmt.Println("Guess should be zero or more")
		return
	}

	randUpperlimit := 10
	if userGuess > 10 {
		randUpperlimit = userGuess
	}

	for turn := 0; turn < 5; turn++ {

		rand := rand.Intn(randUpperlimit + 1)

		if verbose {
			fmt.Printf("turn %d :: randomGen %d \n", turn+1, rand)
		}

		if rand != userGuess {
			continue
		}

		if turn == 0 {
			fmt.Println("You won on firrt turn")
		} else {
			fmt.Println("You won")
		}
		return
	}

	fmt.Println("you loose")
}
