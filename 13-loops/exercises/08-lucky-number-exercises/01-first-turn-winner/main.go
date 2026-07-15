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
// EXERCISE: First Turn Winner
//
//  If the player wins on the first turn, then display
//  a special bonus message.
//
// RESTRICTION
//  The first picked random number by the computer should
//  match the player's guess.
//
// EXAMPLE SCENARIO
//  1. The player guesses 6
//  2. The computer picks a random number and it happens
//     to be 6
//  3. Terminate the game and print the bonus message
// ---------------------------------------------------------

func main() {
	if len(os.Args) != 2 {
		fmt.Println("No input, give me a number between 0 and 9")
		return
	}

	userGuess, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Invalid number %q", os.Args[1])
		return
	}

	if userGuess < 0 || userGuess > 9 {
		fmt.Println("Guess should be between 0 and 9")
		return
	}

	for turn := 0; turn < 5; turn++ {
		rand := rand.Intn(10)
		fmt.Println("rand", rand)
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
