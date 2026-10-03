package main

import (
	"fmt"
	"os"
)

const prefix = "http://"
const maskChar = '*'

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Give me a string")
		return
	}

	str := args[0]
	buf := make([]byte, len(str))

	var c byte
	var matchLen = 0

	for i := range len(str) {
		c = str[i]

		if matchLen != len(prefix) {
			if c == prefix[matchLen] {
				matchLen++
			} else if c != prefix[matchLen] {
				matchLen = 0
			}
		} else if c == ' ' || c == '\t' || c == '\n' {
			matchLen = 0
		}

		if matchLen == len(prefix) {
			c = maskChar
		}

		buf[i] = c
	}

	fmt.Println("output::", string(buf))
}
