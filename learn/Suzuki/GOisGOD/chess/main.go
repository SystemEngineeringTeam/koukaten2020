package main

import "fmt"

func main() {
	var board [8][8]rune
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			board[i][j] = ' '
		}
	}
	tp := "rbnkq"
	for i := 0; i < 5; i++ {
		board[0][i] = rune(tp[i])
		if i < 3 {
			board[0][7-i] = rune(tp[i])
			board[7][i] = rune(tp[i]) - ('a' - 'A')
		}
		board[7][7-i] = rune(tp[i]) - ('a' - 'A')
	}
	for column := range board[1] {
		board[1][column] = 'p'
		board[6][column] = 'P'
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Printf("%c", board[i][j])
		}
		fmt.Printf("\n")
	}
}
