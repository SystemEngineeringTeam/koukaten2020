package main

import (
	"fmt"
	"unicode"
)

func printBoard(board [8][8]rune) {
	for i := range board {
		for j := range board[0] {
			c := board[i][j]

			if c == 0 {
				c = ' '
			}
			fmt.Printf("%c", c)
		}
		fmt.Printf("\n")
	}
}

func main() {
	var board [8][8]rune

	board[0][0] = 'r'
	board[0][1] = 'n'
	board[0][2] = 'b'
	board[0][3] = 'q'
	board[0][4] = 'k'
	board[0][5] = 'b'
	board[0][6] = 'n'
	board[0][7] = 'r'

	for column := range board[1] {
		board[1][column] = 'p'
		board[6][column] = 'p'
	}

	for column := range board[7] {
		board[7][column] = unicode.ToUpper(board[0][column])
	}

	printBoard(board)

}
