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
	"unicode/utf8"
)

// ---------------------------------------------------------
// EXERCISE: Rune Manipulator
//
//  Please read the comments inside the following code.
//
// EXPECTED OUTPUT
//  Please run the solution.
// ---------------------------------------------------------

func main() {
	words := []string{
		"cool",
		"güzel",
		"jīntiān",
		"今天",
		"read 🤓",
	}

	for _, w := range words {

		// Print the byte and rune length of the strings
		// Hint: Use len and utf8.RuneCountInString
		fmt.Println("--len and rune count--")
		fmt.Printf("len=%d runes=%d :: %s\n", len(w), utf8.RuneCountInString(w), w)

		// Print the bytes of the strings in hexadecimal
		// Hint: Use % x verb
		fmt.Println("--Bytes of string in hexadecimal--")
		fmt.Printf("% -30x :: %s\n", w, w)

		// Print the runes of the strings in hexadecimal
		// Hint: Use % x verb
		fmt.Println("--Runes in hexadecimal--")
		for _, r := range w {
			fmt.Printf("% 2x\n", r)
		}

		// Print the runes of the strings as rune literals
		// Hint: Use for range
		fmt.Println("--Runes as literals--")
		for _, s := range w {
			fmt.Printf("%q \n", s)
		}

		// Print the first rune and its byte size of the strings
		// Hint: Use utf8.DecodeRuneInString
		fmt.Println("--First rune and bytesize--")
		firstRune, size := utf8.DecodeRuneInString(w)
		fmt.Printf("first=%c size=%d :: %s\n", firstRune, size, w)

		// Print the last rune of the strings
		// Hint: Use utf8.DecodeLastRuneInString
		fmt.Println("\n--Last rune of string--")
		lastRune, _ := utf8.DecodeLastRuneInString(w)
		fmt.Printf("last=%c  :: %s\n", lastRune, w)

		// Slice and print the first two runes of the strings
		fmt.Println("\n--Slice first two runes--")
		_, firstRuneSize := utf8.DecodeRuneInString(w)
		_, secondRuneSize := utf8.DecodeRuneInString(w[firstRuneSize:])
		fmt.Printf("first two runes of %q => %q", w, w[:firstRuneSize+secondRuneSize])

		// Slice and print the last two runes of the strings
		fmt.Println("\n--Slice last two runes--")
		_, LastRuneSize := utf8.DecodeLastRuneInString(w)
		_, SecondLastRuneSize := utf8.DecodeLastRuneInString(w[:len(w)-LastRuneSize])
		fmt.Printf("last two runes of %q => %q", w, w[len(w)-LastRuneSize-SecondLastRuneSize:])

		// Convert the string to []rune
		// Print the first and last two runes
		fmt.Println("\n--fist and last two runes of []rune--")
		strRune := []rune(w)
		fmt.Printf("first 2 runes = %q", strRune[:2])
		fmt.Printf("last 2 runes = %q", strRune[len(strRune)-2:])
	}
}
