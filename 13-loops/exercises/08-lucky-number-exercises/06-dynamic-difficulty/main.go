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
// EXERCISE: Dynamic Difficulty
//
//  Current game picks only 5 numbers (5 turns).
//
//  Make sure that the game adjust its own difficulty
//  depending on the guess number.
//
// RESTRICTION
//  Do not make the game too easy. Only adjust the
//  difficulty if the guess is above 10.
//
// EXPECTED OUTPUT
//  Suppose that the player runs the game like this:
//    go run main.go 5
//
//    Then the computer should pick 5 random numbers.
//
//  Or, if the player runs it like this:
//    go run main.go 25
//
//    Then the computer may pick 11 random numbers
//    instead.
//
//  Or, if the player runs it like this:
//    go run main.go 100
//
//    Then the computer may pick 30 random numbers
//    instead.
//
//  As you can see, greater guess number causes the
//  game to increase the game turns, which in turn
//  adjust the game's difficulty dynamically.
//
//  Because, greater guess number makes it harder to win.
//  But, this way, game's difficulty will be dynamic.
//  It will adjust its own difficulty depending on the
//  guess number.
// ---------------------------------------------------------

func main() {
	args := os.Args[1:]
	maxTurns := 5

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

	if userGuess > 15 {
		maxTurns = userGuess / 3
	}

	for turn := 0; turn < maxTurns; turn++ {

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
