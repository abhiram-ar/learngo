package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	const (
		width  = 50
		height = 20

		ball      = '⚽'
		emptyCell = '⬛'

		framePerSec = 25
	)

	var (
		vx = 1
		vy = 1

		px = 2
		py = 5
	)

	board := make([][]bool, width)
	for w := range width {
		board[w] = make([]bool, height)
	}

	buffer := make([]rune, 0, width*height)

	screen.Clear()

	// render loop
	for {
		buffer = buffer[:0]
		board[px][py] = false

		px += vx
		py += vy

		if px == width-1 || px == 0 {
			vx *= -1
		}

		if py == height-1 || py == 0 {
			vy *= -1
		}

		board[px][py] = true

		for y := range height {
			for x := range width {
				cell := emptyCell
				if board[x][y] {
					cell = ball
				}
				// fmt.Printf("%s", cell) - lot of prints in a program is not efficient, so we use a buffer approch
				buffer = append(buffer, cell)
			}
			// fmt.Printf("\n")
			buffer = append(buffer, '\n')
		}

		screen.MoveTopLeft()
		fmt.Print(string(buffer))
		// fmt.Printf("vx=%d, vy=%d :: px=%d, py=%d\n", vx, vy, px, py)
		time.Sleep(time.Second / framePerSec)
	}
}
