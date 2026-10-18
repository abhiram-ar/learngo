package main

import "fmt"

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

func newGame(id int, price int, name string, genre string) game {
	return game{item: item{id: id, name: name, price: price}, genre: genre}
}

func addGame(games []game, game game) []game {
	games = append(games, game)

	return games
}

func load() []game {
	games := []game{}

	godOfWarGame := newGame(1, 50, "god of war", "action adventure")
	games = append(games, godOfWarGame)

	xcom := newGame(2, 40, "x-com 2", "strategy")
	games = append(games, xcom)

	minecraft := newGame(3, 20, "minecraft", "sandbox")
	games = append(games, minecraft)

	return games
}

func indexById(games []game) map[int]game {
	byID := make(map[int]game)
	for _, g := range games {
		byID[g.id] = g
	}

	return byID

}

func printGame(g game) {
	fmt.Printf("#%d: %-15q %-20s $%d\n", g.id, g.name, "("+g.genre+")", g.price)
}
