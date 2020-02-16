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

// Universe is a field.
type Universe [][]bool

// NewUniverse is a function.
func NewUniverse() Universe {
	u := make(Universe, height)

	for i := range u {
		for j := 0; j < width; j++ {
			u[i] = append(u[i], false)
		}
	}
	return u
}

// Show is...
func (u Universe) Show() {

	for i := range u {
		for j := range u[i] {
			if u[i][j] == true {
				fmt.Printf("*")
			} else if u[i][j] == false {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}

// Seed is ...
func (u Universe) Seed() {

	for i := range u {
		for j := range u[i] {
			r := rand.Intn(4)
			if r == 0 {
				u[i][j] = true
			}
		}
	}
}

// Alive is ...
func (u Universe) Alive(x, y int) bool {
	y = (y + height) % height
	x = (x + width) % width

	if u[y][x] == true {
		return true
	}
	return false
}

// Neighbors is ...
func (u Universe) Neighbors(x, y int) int {

	count := 0

	xDir := []int{-1, 0, 1}
	yDir := []int{-1, 0, 1}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if u.Alive(x+xDir[j], y+yDir[i]) && (i != 0 && j != 0) {
				count++
			}
		}
	}

	return count

}

// Next is ...
func (u Universe) Next(x, y int) bool {
	if u.Neighbors(x, y) == 2 || u.Neighbors(x, y) == 3 {
		return true
	}

	return false
}

// Step is ...
func Step(a, b Universe) {
	count := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			a[i][j] = a.Next(j, i)
		}
	}

	a, b = b, a

	fmt.Printf("\033[2J")
	a.Show()
	time.Sleep(time.Second * 5 / 10)

	if count < 1000 {
		Step(a, b)
	}
	count++

}

func main() {
	rand.Seed(time.Now().Unix())
	u := NewUniverse()
	n := NewUniverse()
	u.Seed()
	u.Show()

	Step(u, n)

}
