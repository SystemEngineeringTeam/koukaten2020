package main

import (
	"fmt"
	"os"
)

const rows, columns = 9, 9

// NewSudoku is ...
func NewSudoku(field [rows][columns]int8) *[rows][columns]int8 {
	return &field
}

func show(field *[rows][columns]int8) {
	for i := 0; i < rows; i++ {
		fmt.Println(field[i])
	}
}

func checkRows(field *[rows][columns]int8, x, y int8) bool {
	for i := 0; i < rows; i++ {
		if field[y][i] == field[y][x] && x != int8(i) {
			return false
		}
	}

	return true
}

func checkColumns(field *[rows][columns]int8, x, y int8) bool {
	for i := 0; i < rows; i++ {
		if field[i][x] == field[y][x] && x != int8(i) {
			return false
		}
	}

	return true
}

func findArea(x, y int8) (int8, int8) {
	if x < 3 {
		if y < 3 {
			return 0, 0
		}
		if y < 6 {
			return 0, 3
		}

		if y < 9 {
			return 0, 6
		}
	}

	if x < 6 {
		if y < 3 {
			return 3, 0
		}
		if y < 6 {
			return 3, 3
		}

		if y < 9 {
			return 3, 6
		}
	}

	if x < 9 {
		if y < 3 {
			return 6, 0
		}
		if y < 6 {
			return 6, 3
		}

		if y < 9 {
			return 6, 6
		}
	}

	os.Exit(1)
	return 0, 0
}

func checkArea(field *[rows][columns]int8, x, y int8) bool {
	areaX, areaY := findArea(x, y)
	for i := 0; i < 3; i++ {
		if field[y][x] == field[areaY][areaX+int8(i)] && y != areaY && x != areaX+int8(i) {
			return false
		}
	}
	for i := 0; i < 3; i++ {
		if field[y][x] == field[areaY+1][areaX+int8(i)] && y != areaY+1 && x != areaX+int8(i) {
			return false
		}
	}
	for i := 0; i < 3; i++ {
		if field[y][x] == field[areaY+2][areaX+int8(i)] && y != areaY+2 && x != areaX+int8(i) {
			return false
		}
	}

	return true
}

func checkField(field *[rows][columns]int8, x, y int8) bool {
	if checkArea(field, x, y) && checkColumns(field, x, y) && checkRows(field, x, y) {
		return true
	}

	return false
}

func putNum(field *[rows][columns]int8, x, y int8) {
	for i := 1; i <= 9; i++ {
		field[y][x] = int8(i)
		if checkField(field, x, y) {
			continue
		}
	}

	if x == 0 && y != 0 {
		x = 8
		y--

	} else {
		x--

	}
	// fmt.Println(x, y)

	field[y][x] = 0
	// putNum(field, x, y)
}

func fillField(field *[rows][columns]int8) {
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if field[i][j] == 0 {
				fmt.Println(j, i)

				putNum(field, int8(j), int8(i))
			}
		}
	}
}

func main() {
	s := NewSudoku([rows][columns]int8{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	})

	show(s)

	fmt.Println()
	fillField(s)
	show(s)
}
