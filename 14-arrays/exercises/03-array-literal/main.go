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
// EXERCISE: Refactor to Array Literals
//
//  1. Use the 02-get-set-arrays exercise
//
//  2. Refactor the array assignments to array literals
//
//    1. You would need to change the array declarations to array literals
//
//    2. Then, you would need to move the right-hand side of the assignments,
//       into the array literals.
//
// EXPECTED OUTPUT
//   The output should be the same as the 02-get-set-arrays exercise.
// ---------------------------------------------------------

func main() {

	names := [3]string{"suku", "cee", "aswim"}

	for i := 0; i < len(names); i++ {
		fmt.Printf("names[%d] = %q\n", i, names[i])
	}

	distances := [5]int{11, 22, 3: 44, 55}
	for i := 0; i < len(distances); i++ {
		fmt.Printf("distances[%d] = %d\n", i, distances[i])
	}

	data := [5]byte{'A', 'B', 'C', 'D', 'E'}
	for i := 0; i < len(data); i++ {
		fmt.Printf("data[%d] = %d\n", i, data[i])
	}

	zero := [0]byte{}
	for i := 0; i < len(zero); i++ {
		fmt.Printf("data[%d] = %d\n", i, data[i])
	}
}
