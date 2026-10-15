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
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Encode
//
//  Add a new command: "save". Encode the games to json, and
//  print it, then terminate the loop.
//
//  1. Create a new struct type with exported fields: ID, Name, Genre and Price.
//
//  2. Create a new slice using the new struct type.
//
//  3. Save the games into the new slice.
//
//  4. Encode the new slice.
//
//
// RESTRICTION
//  Do not export the fields of the game struct.
//
//
// EXPECTED OUTPUT
//  Inanc's game store has 3 games.
//
//    > list   : lists all the games
//    > id N   : queries a game by id
//    > save   : exports the data to json and quits
//    > quit   : quits
//
//  save
//
//  [
//          {
//                  "id": 1,
//                  "name": "god of war",
//                  "genre": "action adventure",
//                  "price": 50
//          },
//          {
//                  "id": 2,
//                  "name": "x-com 2",
//                  "genre": "strategy",
//                  "price": 40
//          },
//          {
//                  "id": 3,
//                  "name": "minecraft",
//                  "genre": "sandbox",
//                  "price": 20
//          }
//  ]
//
// ---------------------------------------------------------

type command struct {
	cmd string
	msg string
}

var quit = command{cmd: "quit", msg: "quit program"}
var list = command{cmd: "list", msg: "list all games"}
var id = command{cmd: "id", msg: "get game by id"}
var save = command{cmd: "save", msg: "exports the data to json and quits"}
var commands = []command{quit, list, save, id}

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

type exportedGame struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
	Price int    `json:"price"`
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

		case save.cmd:
			gamesToExport := []exportedGame{}
			for _, game := range games {
				gamesToExport = append(gamesToExport, exportedGame{Id: game.id, Name: game.name, Genre: game.genre, Price: game.price})
			}

			exportedGames, err := json.MarshalIndent(gamesToExport, "", "\t")
			if err != nil {
				fmt.Println("error while converting to json")
				continue
			}

			fmt.Printf("%s", exportedGames)
			break mainLoop
		}

	}

	if err := sc.Err(); err != nil {
		fmt.Println(err)
	}

}
