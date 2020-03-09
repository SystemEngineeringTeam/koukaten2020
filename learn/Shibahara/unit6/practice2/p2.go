package main

import "fmt"

type item struct {
	name string
}

type character struct {
	name     string
	leftHand *item
}

func (c *character) pickup(i *item) {
	if c == nil || i == nil {
		return
	}
	fmt.Printf("%s picks up %s\n", c.name, i.name)
	c.leftHand = i
}

func (c *character) give(givenChara *character) {
	if c.leftHand == nil {
		fmt.Printf("%s has nothing\n", c.name)
		return
	}
	if givenChara.leftHand != nil {
		fmt.Printf("%s's lefthand is full\n", givenChara.name)
	}
	givenChara.leftHand = c.leftHand
	c.leftHand = nil
	fmt.Printf("%s gives %s %s\n", c.name, givenChara.leftHand.name, givenChara.name)
}

func main() {
	arthur := character{name: "Arthur"}
	sword := item{name: "sword"}
	arthur.pickup(&sword)

	knight := character{name: "knight"}
	arthur.give(&knight)
}
