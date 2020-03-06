package main

import (
	"fmt"
	"math/rand"
	"time"
)

type turtle struct {
	x, y int
}

func main() {
	rand.Seed(time.Now().Unix())
	t := turtle{0, 0}
	for i := rand.Intn(5) + 5; i > 0; i-- {
		switch rand.Intn(4) {
		case 0:
			t.Up()
		case 1:
			t.Down()
		case 2:
			t.Left()
		case 3:
			t.Right()
		}
	}
	fmt.Println(t)
}
func (t *turtle) Up() {
	t.y--
}
func (t *turtle) Down() {
	t.y++
}
func (t *turtle) Left() {
	t.x--
}
func (t *turtle) Right() {
	t.x++
}
func (t turtle) String() string {
	return fmt.Sprintf("(%d,%d)", t.x, t.y)
}
