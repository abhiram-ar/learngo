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
// EXERCISE: Housing Prices and Averages
//
//  Use the previous exercise to solve this exercise (Housing Prices).
//
//  Your task is to print the averages of the sizes, beds, baths, and prices.
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
//  ===========================================================================
//                 237.50         4.75           2.75           425000.00
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

	// Solve this exercise by using your previous solution for
	// the "Housing Prices" exercise.

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
	fmt.Println()
	fmt.Println(strings.Repeat("=", 75))
	fmt.Printf("%-15s%-15.2f%-15.2f%-15.2f%-15.2f\n", "", findAvg(sizes), findAvg(beds), findAvg(baths), findAvg(prices))

}

func findAvg(nums []int) float64 {
	if len(nums) == 0 {
		return 0
	}

	sum := 0.0
	for idx := range len(nums) {
		sum += float64(nums[idx])
	}
	return sum / float64(len(nums))
}
