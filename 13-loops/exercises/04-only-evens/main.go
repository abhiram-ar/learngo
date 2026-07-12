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
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Only Evens
//
//  1. Extend the "Sum up to N" exercise
//  2. Sum only the even numbers
//
// RESTRICTIONS
//  Skip odd numbers using the `continue` statement
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Give 2 numbers")
		return
	}

	min, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Invalid number %q", os.Args[1])
		return
	}

	max, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Invalid number %q", os.Args[2])
		return
	}

	if min > max {
		fmt.Println("first argument is less than second")
		return
	}

	var sum int
	for i := min; i <= max; i++ {
		if i%2 == 1 {
			continue
		}

		sum += i

		fmt.Print(i)
		if i < max {
			fmt.Print(" + ")
		}
	}
	fmt.Printf(" = %d\n", sum)
}
