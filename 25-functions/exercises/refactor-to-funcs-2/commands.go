package main

import (
	"fmt"
	"strconv"
	"strings"
)

func runCmd(input string, games []game, gameById map[int]game) bool {

	fmt.Println()

	cmd := strings.Fields(input)
	if len(cmd) == 0 {
		return true
	}

	switch cmd[0] {
	case "quit":
		return cmdQuit()

	case "list":
		return cmdList(games)

	case "id":
		return cmdById(cmd, games, gameById)
	}

	return true
}

func cmdQuit() bool {
	fmt.Println("bye!")
	return false
}

func cmdList(games []game) bool {
	for _, g := range games {
		printGame(g)
	}
	return true
}

func cmdById(cmd []string, games []game, gameById map[int]game) bool {
	// cmdByID
	if len(cmd) != 2 {
		fmt.Println("wrong id")
		return true
	}

	id, err := strconv.Atoi(cmd[1])
	if err != nil {
		fmt.Println("wrong id")
		return true
	}

	g, ok := gameById[id]
	if !ok {
		fmt.Println("sorry. I don't have the game")
		return true
	}

	printGame(g)
	return true

}
