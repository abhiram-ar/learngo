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
// EXERCISE: Warm-up
//
//  Create and print the following maps.
//
//  1. Phone numbers by last name
//  2. Product availability by Product ID
//  3. Multiple phone numbers by last name
//  4. Shopping basket by Customer ID
//
//     Each item in the shopping basket has a Product ID and
//     quantity. Through the map, you can tell:
//     "Mr. X has bought Y bananas"
//
// ---------------------------------------------------------

func main() {
	// Hint: Store phone numbers as text

	// #1
	// Key        : Last name
	// Element    : Phone number
	nameToPhone := map[string]int{
		"sajeev":    9946202959,
		"sukumaran": 8899472950,
	}
	fmt.Printf("%#v\n", nameToPhone)

	// #2
	// Key        : Product ID
	// Element    : Available / Unavailable
	prodcutAvailability := map[string]bool{
		"d1": false,
		"d2": true,
	}
	fmt.Printf("%#v\n", prodcutAvailability)

	// #3
	// Key        : Last name
	// Element    : Phone numbers
	phoneNumbers := map[string][]int{
		"abhiram": []int{9946202959, 123456789},
		"suku":    []int{9988776655, 2233445566},
	}
	fmt.Printf("%#v\n", phoneNumbers)

	// #4
	// Key        : Customer ID
	// Element Key:
	//   Key: Product ID Element: Quantity
	customerBasket := map[string]map[string]int{
		"c1": map[string]int{"d1": 2, "d2": 3},
		"c2": map[string]int{"d1": 5, "d2": 0},
	}
	fmt.Printf("%#v\n", customerBasket)
}
