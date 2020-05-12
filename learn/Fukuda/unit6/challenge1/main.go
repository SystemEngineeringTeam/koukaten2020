package main

import (
	"fmt"
	"os"
)

const rows, columns = 9, 9

type list struct {
	x, y       int8
	next, back *list
}

// NewSudoku is ...
func NewSudoku(field [rows][columns]int8) *[rows][columns]int8 {
	return &field
}

func newList(l []list, field *[rows][columns]int8) []list {
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if field[i][j] == 0 {
				l = append(l, list{int8(j), int8(i), nil, nil})
			}
		}
	}

	for i := range l {
		if i != len(l)-1 {
			l[i].next = &l[i+1]
		}

		if i != 0 {
			l[i].back = &l[i-1]
		}
	}

	return l
}

func show(field *[rows][columns]int8) {
	for i := 0; i < rows; i++ {
		fmt.Println(field[i])
	}
}

func checkRows(field *[rows][columns]int8, l list) bool {
	for i := 0; i < rows; i++ {
		if field[l.y][i] == field[l.y][l.x] && l.x != int8(i) {
			return false
		}
	}

	return true
}

func checkColumns(field *[rows][columns]int8, l list) bool {
	for i := 0; i < rows; i++ {
		if field[i][l.x] == field[l.y][l.x] && l.y != int8(i) {
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

func checkArea(field *[rows][columns]int8, l list) bool {
	areaX, areaY := findArea(l.x, l.y)
	for i := 0; i < 3; i++ {
		if field[l.y][l.x] == field[areaY][areaX+int8(i)] && ((l.y == areaY && l.x != areaX+int8(i)) || (l.y != areaY && l.x == areaX+int8(i)) || (l.y != areaY && l.x != areaX+int8(i))) {
			return false
		}
	}
	for i := 0; i < 3; i++ {
		if field[l.y][l.x] == field[areaY+1][areaX+int8(i)] && ((l.y == areaY+1 && l.x != areaX+int8(i)) || (l.y != areaY+1 && l.x == areaX+int8(i)) || (l.y != areaY+1 && l.x != areaX+int8(i))) {
			return false
		}
	}
	for i := 0; i < 3; i++ {
		if field[l.y][l.x] == field[areaY+2][areaX+int8(i)] && ((l.y == areaY+2 && l.x != areaX+int8(i)) || (l.y != areaY+2 && l.x == areaX+int8(i)) || (l.y != areaY+2 && l.x != areaX+int8(i))) {
			return false
		}
	}

	return true
}

func checkField(field *[rows][columns]int8, l list) bool {
	if checkArea(field, l) && checkColumns(field, l) && checkRows(field, l) {
		return true
	}

	return false
}

func incNum(field *[rows][columns]int8, l list) {
	field[l.y][l.x]++

	if field[l.y][l.x] > 9 {
		refillZero(field, l)
		incNum(field, *l.back)
	} else if checkField(field, l) && l.next != nil {
		incNum(field, *l.next)
	} else if !checkField(field, l) {
		incNum(field, l)
	}
}

func refillZero(field *[rows][columns]int8, l list) {
	field[l.y][l.x] = 0

	if l.next != nil {
		refillZero(field, *l.next)
	}
}

func main() {
	s := NewSudoku([rows][columns]int8{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 6, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 6, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 6, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 6},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	})

	var l = make([]list, 0)
	l = newList(l, s)

	show(s)
	fmt.Println()

	incNum(s, l[0])
	show(s)
}
