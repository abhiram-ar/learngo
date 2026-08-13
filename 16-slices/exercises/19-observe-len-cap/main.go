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
)

// ---------------------------------------------------------
// EXERCISE: Observe the length and capacity
//
//  Follow the instructions inside the code below to
//  gain more intuition about the length and capacity of a slice.
//
// ---------------------------------------------------------

func main() {
	// --- #1 ---
	// 1. create a new slice named: games
	// games := []int{1, 2, 3, 4}

	// 2. print the length and capacity of the games slice
	// fmt.Printf("games len=%d cap=%d\n", len(games), cap(games))

	// 3. comment out the games slice
	//    then declare it as an empty slice
	// games := []string{}

	// 4. print the length and capacity of the games slice
	// fmt.Printf("games len=%d cap=%d\n", len(games), cap(games))

	// 5. append the elements: "pacman", "mario", "tetris", "doom"
	// games = append(games, "pacman", "mario", "tetris", "dooom")

	// 6. print the length and capacity of the games slice
	// fmt.Printf("games len=%d cap=%d\n", len(games), cap(games))

	// 7. comment out everything
	//
	// 8. declare the games slice again using a slice literal
	//    (use the same elements from step 5)
	games := []string{"pacman", "mario", "tetris", "dooom"}

	// --- #2 ---
	// 1. use a loop from 0 to 4 to slice the games slice, element by element.
	// 2. print its length and capacity along the way (in the loop).
	fmt.Println()
	for i := range len(games) + 1 {
		slice := games[:i]
		fmt.Printf("games[:%d]'s len: %d cap: %d\n", i, len(slice), cap(slice))
	}
	fmt.Println()

	// --- #3 ---
	// 1. slice the games slice up to zero element
	//    (save the result to a new slice named: "zero")
	var zero = games[:0]

	// 2. print the games and the new slice's len and cap
	fmt.Printf("games len=%d cap=%d\n", len(games), cap(games))
	fmt.Printf("zero  len=%d cap=%d\n", len(zero), cap(zero))
	fmt.Println()

	// 3. append a new element to the new slice
	zero = append(zero, "newElem")

	// 4. print the new slice's lens and caps
	fmt.Printf("zero  len=%d cap=%d\n", len(zero), cap(zero))

	// 5. repeat the last two steps 5 times (use a loop)
	data := []string{"ultima", "dagger", "pong", "coldspot", "zetra"}
	for i := range 5 {
		zero = append(zero, data[i])
		fmt.Printf("zero  len=%d cap=%d\n", len(zero), cap(zero))
	}

	// 6. notice the growth of the capacity after the 5th append
	fmt.Println()

	// --- #4 ---
	// using a range loop, slice the zero slice element by element,
	// and print its length and capacity along the way.
	for i := range len(zero) + 1 {
		slice := zero[:i]
		fmt.Printf("zero[:%d]'s len:%d cap:%d\n", i, len(slice), cap(slice))
	}

	// observe that, the range loop only loops for the length, not the cap.
	fmt.Println()

	// --- #5 ---
	// 1. do the 3rd step above again but this time, start by slicing
	//    the zero slice up to its capacity (use the cap function).
	// 2. print the elements of the zero slice in the loop.
	fmt.Println()

	zero = zero[:cap(zero)]
	for i := range len(zero) + 1 {
		slice := zero[:i]
		fmt.Printf("zero[:%d]'s  len: %d cap: %d - %q\n", i, len(slice), cap(slice), slice)
	}

	// --- #6 ---
	// 1. change the one of the elements of the zero slice
	zero[2] = "changed-elem"

	// 2. change the same element of the games slice
	games[2] = "changed-elem"

	// 3. print the games and the zero slices
	fmt.Println()
	fmt.Printf("games len:%d cap:%d ::%q\n", len(games), cap(games), games)
	fmt.Printf("zero  len:%d cap:%d ::%q\n", len(zero), cap(zero), zero)

	// 4. observe that they don't have the same backing array
	fmt.Println()

	// --- #7 ---
	// try to slice the games slice beyond its capacity
	games = games[:3]
	fmt.Printf("games ::%q\n", games)

}
