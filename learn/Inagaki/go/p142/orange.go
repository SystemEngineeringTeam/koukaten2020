package main

import "fmt"

func boardView(board [8][8]rune) {
	fmt.Println("kqrbnp")
	for Right := range board {
		for Left := range board[Right] {
			fmt.Printf("%c", board[Right][Left])
		}
		fmt.Printf("\n")
	}
	fmt.Println("KQRBNP")
}

func main() {
	var board [8][8]rune
	for Right := range board {
		for Left := range board[Right] {
			if Right == 0 {
				if Left == 0 || Left == 7 {
					board[Right][Left] = 'r'
				} else if Left == 1 || Left == 6 {
					board[Right][Left] = 'n'
				} else if Left == 2 || Left == 5 {
					board[Right][Left] = 'b'
				} else if Left == 3 {
					board[Right][Left] = 'q'
				} else if Left == 4 {
					board[Right][Left] = 'k'
				}
			} else if Right == 7 {
				if Left == 0 || Left == 7 {
					board[Right][Left] = 'R'
				} else if Left == 1 || Left == 6 {
					board[Right][Left] = 'N'
				} else if Left == 2 || Left == 5 {
					board[Right][Left] = 'B'
				} else if Left == 3 {
					board[Right][Left] = 'K'
				} else if Left == 4 {
					board[Right][Left] = 'Q'
				}
			} else if Right == 1 {
				board[Right][Left] = 'p'
			} else if Right == 6 {
				board[Right][Left] = 'P'
			} else {
				board[Right][Left] = ' '
			}
		}
	}
	boardView(board)
}
