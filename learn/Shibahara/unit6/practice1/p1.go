package main

import "fmt"

type turtle struct {
	posx, posy int
}

func (t *turtle) up() {
	t.posy++
}
func (t *turtle) down() {
	t.posy--
}
func (t *turtle) left() {
	t.posx--
}
func (t *turtle) right() {
	t.posx++
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
