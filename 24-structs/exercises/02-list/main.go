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
)

// ---------------------------------------------------------
// EXERCISE: List
//
//  Now, it's time to add an interface to your program using
//  the bufio.Scanner. So the users can list the games, or
//  search for the games by id.
//
//  1. Scan for the input in a loop (use bufio.Scanner)
//
//  2. Print the available commands.
//
//  3. Implement the quit command: Quits from the loop.
//
//  4. Implement the list command: Lists all the games.
//
//
// EXPECTED OUTPUT
//  Please run the solution and try the program with list and
//  quit commands.
// ---------------------------------------------------------

type command struct {
	key string
	msg string
}

var quit = command{key: "q", msg: "quit program"}
var list = command{key: "l", msg: "list all games"}
var commands = []command{quit, list}

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

func main() {
	// use your solution from the previous exercise
	sc := bufio.NewScanner(os.Stdin)

	games := []game{
		{item: item{id: 1, name: "god of war", price: 50}, genre: "action adventure"},
		{item: item{id: 2, name: "x-com 2", price: 30}, genre: "strategy"},
		{item: item{id: 3, name: "minecreaft", price: 20}, genre: "sandbox"},
	}

	fmt.Println("Enter:")
	for _, command := range commands {
		fmt.Printf("[%s] for %s\n", command.key, command.msg)
	}

mainLoop:
	for {
		hasValue := sc.Scan()
		if !hasValue {
			continue
		}

		switch sc.Text() {
		case quit.key:
			break mainLoop
		case list.key:

			fmt.Printf("%-10s %-20s %-10s %-20s\n\n", "id", "name", "price", "minecraft")
			for _, game := range games {
				fmt.Printf("%-10d %-20s %-10d %-20s\n", game.id, game.name, game.price, game.genre)
			}
		}

	}

	if err := sc.Err(); err != nil {
		fmt.Println(err)
	}

}
