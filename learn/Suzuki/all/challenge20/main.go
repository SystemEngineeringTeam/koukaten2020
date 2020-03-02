package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 80
	height = 15
)

//field
type Universe [][]bool

func main() {
	a := NewUniverse()
	a.Seed()
	Step(a, NewUniverse())
}

func NewUniverse() Universe {
	a := make([][]bool, width)
	for i := range a {
		a[i] = make([]bool, height)
	}
	return a
}

func (u Universe) Show() {
	// fmt.Printf("========================================================================\n")
	for i := range u[0] {
		for j := range u {
			if u[j][i] {
				fmt.Printf("o")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}
func (u Universe) Seed() {
	rand.Seed(time.Now().Unix())
	for i := range u {
		for j := range u[i] {
			if rand.Intn(4) == 0 {
				u[i][j] = true
			}
		}
	}
}
func (u Universe) Alive(x, y int) bool {
	return u[x][y]
}
func (u Universe) Neighbors(x, y int) int {
	a := 0
	for i := -1; i <= 1; i++ {
		if x+i > 0 && x+i < len(u)-1 {
			for j := -1; j <= 1; j++ {
				if y+j > 0 && y+j < len(u[x])-1 {
					if u[x+i][y+j] {
						a++
					}
				}
			}
		}
	}
	if u[x][y] {
		a--
	}
	return a
}
func (u Universe) Next(x, y int) bool {
	ans := false
	if u[x][y] {
		a := u.Neighbors(x, y)
		if a == 2 || a == 3 {
			ans = true
		}
	} else {
		if u.Neighbors(x, y) == 3 {
			ans = true
		}
	}
	return ans
}
func Step(a, b Universe) {
	for n := 0; ; n++ {
		for i := range a {
			for j := range a[i] {
				b[i][j] = a.Next(i, j)
			}
		}
		a.Show()
		time.Sleep(time.Second / 15)
		a, b = b, a
	}
}
