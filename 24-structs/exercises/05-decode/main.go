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
	"encoding/json/jsontext"
	"encoding/json/v2"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Decode
//
//  At the beginning of the file:
//
//  1. Load the initial data to the game store from json.
//     (see the data constant below)
//
//  2. Load the decoded values into the usual `game` values (to the games slice as well).
//
//     So the rest of the program can work intact.
//
//
// HINT
//
//  Move the jsonGame type to the top and reuse it both when
//  loading the initial data, and in the "save" command.
//
//
// EXPECTED OUTPUT
//  Please run the solution to see the output.
// ---------------------------------------------------------

const data = `
[
        {
                "id": 1,
                "name": "god of war",
                "genre": "action adventure",
                "price": 50
        },
        {
                "id": 2,
                "name": "x-com 2",
                "genre": "strategy",
                "price": 40
        },
        {
                "id": 3,
                "name": "minecraft",
                "genre": "sandbox",
                "price": 20
        }
]`

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

	decodedGames := []exportedGame{}
	err := json.Unmarshal([]byte(data), &decodedGames)
	if err != nil {
		fmt.Println("error while decoding")
		return
	}

	games := []game{}
	for _, decodedGame := range decodedGames {
		games = append(games,
			game{
				item:  item{id: decodedGame.Id, name: decodedGame.Name, price: decodedGame.Price},
				genre: decodedGame.Genre,
			},
		)
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

			exportedGames, err := json.Marshal(gamesToExport, jsontext.Multiline(true))
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
