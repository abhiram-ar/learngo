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
// EXERCISE: Double Guesses
//
//  Let the player guess two numbers instead of one.
//
// HINT:
//  Generate random numbers using the greatest number
//  among the guessed numbers.
//
// EXAMPLES
//  go run main.go 5 6
//  Player wins if the random number is either 5 or 6.
// ---------------------------------------------------------

func main() {
	if len(os.Args) != 3 {
		fmt.Println("No input, give me 2 numbers between 0 and 9")
		return
	}

	userGuess1, err1 := strconv.Atoi(os.Args[1])
	userGuess2, err2 := strconv.Atoi(os.Args[2])

	if err1 != nil || err2 != nil ||
		(userGuess1 < 0 || userGuess1 > 9) ||
		(userGuess2 < 0 || userGuess2 > 9) {
		fmt.Println("Guess should be between 0 and 9")
		return
	}

	var min = userGuess1
	if userGuess2 < min {
		min = userGuess2
	}

	for turn := 0; turn < 5; turn++ {
		rand := rand.Intn(min + 1)
		fmt.Println("rand", rand)
		if !(rand == userGuess1 || rand == userGuess2) {
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
