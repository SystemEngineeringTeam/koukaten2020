package main

import "fmt"

type turtle struct {
	x, y int
}

func (t *turtle) toUpside() {
	t.y--
}

func (t *turtle) toDownside() {
	t.y++
}

func (t *turtle) toRightside() {
	t.x++
}

func (t *turtle) toLeftside() {
	t.x--
}

func (t *turtle) show() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if t.y == i && t.x == j {
				fmt.Printf("■")
			} else {
				fmt.Printf("□")
			}
		}
		fmt.Printf("\n")
	}

	fmt.Println()
}

func main() {
	turtle := &turtle{
		0, 0,
	}

	turtle.show()

	turtle.toDownside()
	turtle.toRightside()
	turtle.toLeftside()
	turtle.toRightside()
	turtle.show()

}
