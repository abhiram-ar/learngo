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
// EXERCISE: Verbose Mode
//
//  When the player runs the game like this:
//
//    go run main.go -v 4
//
//  Display each generated random number:

//    1 3 4 🎉  YOU WIN!
//
//  In this example, computer picks 1, 3, and 4. And the
//  player wins.
//
// HINT
//  You need to get and interpret the command-line arguments.
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

	if userGuess < 0 || userGuess > 9 {
		fmt.Println("Guess should be between 0 and 9")
		return
	}

	for turn := 0; turn < 5; turn++ {
		rand := rand.Intn(10)

		if verbose {
			fmt.Printf("turn %d\n", turn+1)
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
