// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Refactor to Ellipsis
//
//  1. Use the 03-array-literal exercise
//
//  2. Refactor the length of the array literals to ellipsis
//
//    This means: Use the ellipsis instead of defining the array's length
//                manually.
//
// EXPECTED OUTPUT
//   The output should be the same as the 03-array-literal exercise.
// ---------------------------------------------------------

func main() {
	names := [...]string{"suku", "cee", "aswim"}

	for i := 0; i < len(names); i++ {
		fmt.Printf("names[%d] = %q\n", i, names[i])
	}

	distances := [...]int{11, 22, 3: 44, 55}
	for i := 0; i < len(distances); i++ {
		fmt.Printf("distances[%d] = %d\n", i, distances[i])
	}

	data := [...]byte{'A', 'B', 'C', 'D', 'E'}
	for i := 0; i < len(data); i++ {
		fmt.Printf("data[%d] = %d\n", i, data[i])
	}

	zero := [...]byte{}
	for i := 0; i < len(zero); i++ {
		fmt.Printf("data[%d] = %d\n", i, data[i])
	}
}
