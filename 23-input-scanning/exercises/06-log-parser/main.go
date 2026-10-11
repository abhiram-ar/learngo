// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Log Parser from Stratch
//
//  You've watched the lecture. Now, try to create the same
//  log parser program on your own. Do not look at the lecture,
//  and the existing source code.
//
//
// EXPECTED OUTPUT
//
//  go run main.go < log.txt
//
//   DOMAIN                             VISITS
//   ---------------------------------------------
//   blog.golang.org                        30
//   golang.org                             10
//   learngoprogramming.com                 20
//
//   TOTAL                                  60
//
//
//  go run main.go < log_err_missing.txt
//
//   wrong input: [golang.org] (line #3)
//
//
//  go run main.go < log_err_negative.txt
//
//   wrong input: "-100" (line #3)
//
//
//  go run main.go < log_err_str.txt
//
//   wrong input: "FOUR" (line #3)
//
// ---------------------------------------------------------

func main() {
	var (
		domainVisits = make(map[string]int)
		total        = 0
		line         = 0
	)
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		line++
		sections := strings.Fields(sc.Text())

		if len(sections) < 2 {
			fmt.Printf("wrong input: %v (line #%d)", sections, line)
			return
		}

		domain := sections[0]
		visits, err := strconv.Atoi(sections[1])

		if err != nil || visits < 0 {
			fmt.Printf("wrong input: %q (line %d)\n", sections[1], line)
			return
		}

		domainVisits[domain] += visits
		total += visits
	}

	if err := sc.Err(); err != nil {
		fmt.Println("Error reading stream:", err)
		return
	}

	fmt.Printf("\n%-30s%-10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 40))
	for k, v := range domainVisits {
		fmt.Printf("%-30s%-10d\n", k, v)
	}
	fmt.Printf("\n%-30s%-10d\n", "TOTAL", total)
}
