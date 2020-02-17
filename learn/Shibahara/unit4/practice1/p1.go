package main

import "fmt"

func printBoard(board [8][8]rune) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Printf("%c", board[i][j])
		}
		fmt.Println()
	}
}

func main() {
	board := [8][8]rune{
		{'r', 'n', 'b', 'q', 'k', 'b', 'n', 'r'},
		{},
		{},
		{},
		{},
		{},
		{},
		{'R', 'N', 'B', 'Q', 'K', 'B', 'N', 'R'},
	}
	for i := 1; i < 7; i++ {
		for j := 0; j < 8; j++ {
			if i == 1 {
				board[i][j] = 'p'
				continue
			} else if i == 6 {
				board[i][j] = 'P'
				continue
			}
			board[i][j] = '.'
		}
	}

	printBoard(board)

}
