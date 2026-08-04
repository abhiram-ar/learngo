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
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Housing Prices
//
//  We have received housing prices. Your task is to load the data into
//  appropriately typed slices then print them.
//
//  1. Check out the expected output
//
//
//  2. Check out the code below
//
//     1. header   : stores the column headers
//     2. data     : stores the real data related to each column
//     3. separator: you will use it to separate the data
//
//
//  3. Parse the data
//
//     1. Separate it into rows by using the newline character ("\n")
//
//     2. For each row, separate it by using the separator (",")
//
//
//  4. Load the data into distinct slices
//
//     1. Load the locations into a []string
//     2. Load the sizes into []int
//     3. Load the beds into []int
//     4. Load the baths into []int
//     5. Load the prices into []int
//
//
//  5. Print the header
//
//     1. Separate it by using the separator
//
//     2. Print each column 15 character wide ("%-15s")
//
//
//  6. Print the rows from the slices that you've created, line by line
//
//
// EXPECTED OUTPUT
//
//  Location       Size           Beds           Baths          Price
//  ===========================================================================
//  New York       100            2              1              100000
//  New York       150            3              2              200000
//  Paris          200            4              3              400000
//  Istanbul       500            10             5              1000000
//
//
// HINTS
//
//  + strings.Split function can separate a string into a []string
//
// ---------------------------------------------------------

func main() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	for _, header := range strings.Split(header, ",") {
		fmt.Printf("%-15s", header)
	}
	fmt.Printf("\n")
	fmt.Println(strings.Repeat("=", 75))

	var (
		locations []string
		sizes     []int
		beds      []int
		baths     []int
		prices    []int
	)

	noOfRows := 0
	for row := range strings.SplitSeq(data, "\n") {
		rowData := strings.Split(row, separator)

		locations = append(locations, rowData[0])

		size, _ := strconv.Atoi(rowData[1])
		sizes = append(sizes, size)

		bed, _ := strconv.Atoi(rowData[2])
		beds = append(beds, bed)

		bath, _ := strconv.Atoi(rowData[3])
		baths = append(baths, bath)

		price, _ := strconv.Atoi(rowData[4])
		prices = append(prices, price)

		noOfRows++
	}

	for idx := range noOfRows {
		fmt.Printf("%-15s%-15d%-15d%-15d%-15d\n", locations[idx], sizes[idx], beds[idx], baths[idx], prices[idx])
	}
}
