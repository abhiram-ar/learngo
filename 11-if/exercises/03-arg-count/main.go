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
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Arg Count
//
//  1. Get arguments from command-line.
//  2. Print the expected outputs below depending on the number
//     of arguments.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me args
//
//  go run main.go hello
//    There is one: "hello"
//
//  go run main.go hi there
//    There are two: "hi there"
//
//  go run main.go I wanna be a gopher
//    There are 5 arguments
// ---------------------------------------------------------

func main() {
	argsLen := len(os.Args)
	if argsLen == 0 {
		fmt.Println("Give me args")
	} else if argsLen == 1 {
		fmt.Println("There is one:", os.Args[1])
	} else if argsLen == 2 {
		fmt.Printf("There are two: %s %s \n", os.Args[1], os.Args[2])
	} else {
		fmt.Printf("There are %d arguuments", argsLen)
	}
}
