package main

import "fmt"

type turtle struct {
	posX, posY int
}

func (t *turtle) up() {
	t.posY--
}
func (t *turtle) down() {
	t.posY++
}
func (t *turtle) left() {
	t.posX--
}
func (t *turtle) right() {
	t.posX++
}

func main() {
	t := turtle{0, 0}
	fmt.Println(t)

	t.up()
	t.left()
	fmt.Println(t)

	t.down()
	t.right()
	fmt.Println(t)

}
