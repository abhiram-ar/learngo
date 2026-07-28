// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

var (
	digitchar     = "█"
	separatorChar = "░"
)

type Digit [5]string

var zero = Digit{
	"███",
	"█ █",
	"█ █",
	"█ █",
	"███",
}

var one = Digit{
	"██ ",
	" █ ",
	" █ ",
	" █ ",
	"███",
}

var two = Digit{
	"███",
	"  █",
	"███",
	"█  ",
	"███",
}

var three = Digit{
	"███",
	"  █",
	"███",
	"  █",
	"███",
}

var four = Digit{
	"█ █",
	"█ █",
	"███",
	"  █",
	"  █",
}

var five = Digit{
	"███",
	"█  ",
	"███",
	"  █",
	"███",
}

var six = Digit{
	"███",
	"█  ",
	"███",
	"█ █",
	"███",
}

var seven = Digit{
	"███",
	"  █",
	"  █",
	"  █",
	"  █",
}

var eight = Digit{
	"███",
	"█ █",
	"███",
	"█ █",
	"███",
}

var nine = Digit{
	"███",
	"█ █",
	"███",
	"  █",
	"███",
}

var digits = [...]Digit{zero, one, two, three, four, five, six, seven, eight, nine}

func main() {
	for line := range digits[0] {
		for j := range digits {
			fmt.Printf("%s ", digits[j][line])
		}
		fmt.Println("")
	}
}
