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
// EXERCISE: Query By Id
//
//	Add a new command: "id". So the users can query the games
//	by id.
//
//	1. Before the loop, index the games by id (use a map).
//
//	2. Add the "id" command.
//	   When a user types: id 2
//	   It should print only the game with id: 2.
//
//	3. Handle the errors:
//
//	   id
//	   wrong id
//
//	   id HEY
//	   wrong id
//
//	   id 10
//	   sorry. I don't have the game
//
//	   id 1
//	   #1: "god of war" (action adventure) $50
//
//	   id 2
//	   #2: "x-com 2" (strategy) $40
//
// EXPECTED OUTPUT
//
//	Please also run the solution and try the program with
//	list, quit, and id commands to see it in action.
//
// ---------------------------------------------------------

type command struct {
	cmd string
	msg string
}

var quit = command{cmd: "quit", msg: "quit program"}
var list = command{cmd: "list", msg: "list all games"}
var id = command{cmd: "id", msg: "get game by id"}
var commands = []command{quit, list, id}

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
	sc := bufio.NewScanner(os.Stdin)

	games := []game{
		{item: item{id: 1, name: "god of war", price: 50}, genre: "action adventure"},
		{item: item{id: 2, name: "x-com 2", price: 30}, genre: "strategy"},
		{item: item{id: 3, name: "minecreaft", price: 20}, genre: "sandbox"},
	}

	gamesById := make(map[int]game, len(games))
	for idx := range games {
		gamesById[games[idx].id] = games[idx]
	}

	fmt.Println("Enter:")
	for _, command := range commands {
		fmt.Printf("[%s] :: %s\n", command.cmd, command.msg)
	}

mainLoop:
	for {
		fmt.Println()
		hasValue := sc.Scan()
		if !hasValue {
			continue
		}

		cmds := strings.Fields(sc.Text())
		if len(cmds) == 0 {
			continue
		}

		switch cmds[0] {

		case quit.cmd:
			break mainLoop

		case list.cmd:
			fmt.Printf("%-10s %-20s %-10s %-20s\n\n", "id", "name", "price", "minecraft")
			for _, game := range games {
				fmt.Printf("%-10d %-20s %-10d %-20s\n", game.id, game.name, game.price, game.genre)
			}

		case id.cmd:
			fields := strings.Fields(sc.Text())
			if len(fields) < 2 || fields[1] == "" {
				fmt.Println("wrong id")
				continue
			}

			id, err := strconv.Atoi(fields[1])
			if err != nil {
				fmt.Println("wrong id")
				continue
			}

			game, ok := gamesById[id]
			if !ok {
				fmt.Println("sorry. I don't have the game")
				continue
			}

			fmt.Printf("#%d: %q (%s) $%d\n", game.id, game.name, game.genre, game.price)

		}

	}

	if err := sc.Err(); err != nil {
		fmt.Println(err)
	}

}
