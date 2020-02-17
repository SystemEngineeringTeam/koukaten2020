package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const (
	height = 15
	width  = 80
)

//Universe is a world cell
type Universe [][]bool

//NewUniverse is creating a new world
func NewUniverse() Universe {
	newuniverse := make(Universe, height)
	for i := range newuniverse {
		newuniverse[i] = make([]bool, width)
	}

	return newuniverse
}

//Show is printing a world
func (u Universe) Show() {
	for i := range u {
		for j := range u[i] {
			if u[i][j] == true {
				fmt.Print("*")
				continue
			}
			fmt.Print(".")
		}
		fmt.Println()
	}
}

//Seed is sowing random life seed
func (u Universe) Seed() {
	rand.Seed(time.Now().Unix())
	if rand.Intn(3) == 0 {
		u[rand.Intn(height)][rand.Intn(width)] = true
	}
}

func main() {
	u := NewUniverse()

	u.Show()
	for {
		u.Show()
		time.Sleep(time.Millisecond * 300)
		cmd := exec.Command("powershell", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		u.Seed()
	}

}
