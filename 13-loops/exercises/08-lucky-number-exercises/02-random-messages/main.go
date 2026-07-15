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
// EXERCISE: Random Messages
//
//  Display a few different won and lost messages "randomly".
//
// HINTS
//  1. You can use a switch statement to do that.
//  2. Set its condition to the random number generator.
//  3. I would use a short switch.
//
// EXAMPLES
//  The Player wins: then randomly print one of these:
//
//  go run main.go 5
//    YOU WON
//  go run main.go 5
//    YOU'RE AWESOME
//
//  The Player loses: then randomly print one of these:
//
//  go run main.go 5
//    LOSER!
//  go run main.go 5
//    YOU LOST. TRY AGAIN?
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
		randInt := rand.Intn(10)
		fmt.Println("rand", randInt)
		if randInt != userGuess {
			continue
		}

		var winMsg string
		switch rand.Intn(2) {
		case 0:
			winMsg = "Heyyy you won"
		case 1:
			winMsg = "You are a champ"
		}
		fmt.Println(winMsg)
		return
	}

	var looseMsg string
	switch msgRandomizer := rand.Intn(2); msgRandomizer {
	case 0:
		looseMsg = "Heyyy you loose"
	case 1:
		looseMsg = "You are a failure"
	}
	fmt.Println(looseMsg)
}
