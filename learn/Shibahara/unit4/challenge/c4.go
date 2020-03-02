package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 80
	height = 15
	line   = "================================================================================"
)

//Universe is a world cell
type Universe [][]bool

//NewUniverse is creating a new world
func NewUniverse() Universe {
	newuniverse := make(Universe, height)
	for i := range newuniverse {
		newuniverse[i] = make([]bool, width)
	}

	return newuniverse //u[height][width]==u[y][x]
}

//Show is printing a world
func (u Universe) Show() {
	for y := range u {
		for x := range u[y] {
			if u[y][x] == true {
				fmt.Print("*")
				continue
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

//Seed is sowing random life seed
func (u Universe) Seed() {
	for i := 0; i <= (width*height)/4; i++ {
		u[rand.Intn(height)][rand.Intn(width)] = true
	}
}

//Alive is alive or dead
func (u Universe) Alive(x, y int) bool {
	x = (x + width) % width
	y = (y + height) % height
	return u[y][x]
}

//Neighbors is check around cell
func (u Universe) Neighbors(x, y int) int {
	count := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if j == 0 && i == 0 {
				continue
			} else if u.Alive(x+j, y+i) {
				count++
			}
		}
	}

	return count
}

//Next is next generation status
func (u Universe) Next(x, y int) bool {
	if u.Neighbors(x, y) < 2 {
		return false
	} else if u.Neighbors(x, y) == 3 {
		return true
	} else if 3 < u.Neighbors(x, y) {
		return false
	}
	return u[y][x]
}

//Step is next world
func Step(a, b Universe) {
	for y := range a {
		for x := range a[y] {
			b[y][x] = a.Next(x, y)
		}
	}
}

func main() {
	u1 := NewUniverse()
	u2 := NewUniverse()

	rand.Seed(time.Now().Unix())

	u1.Seed()

	for i := 0; i < 1000; i++ {
		fmt.Println(line)
		// cmd := exec.Command("cmd", "/c", "cls")
		// cmd.Stdout = os.Stdout
		// cmd.Run()
		Step(u1, u2)
		u1, u2 = u2, u1
		u1.Show()
		fmt.Println(line)
		time.Sleep(time.Millisecond * 150)
	}
}
