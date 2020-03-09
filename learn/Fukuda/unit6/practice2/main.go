package main

import "fmt"

type item string

type character struct {
	name     string
	leftHand *item
}

func (c *character) pickup(i *item) {
	c.leftHand = i
	fmt.Println(c.name + " picked " + string(*i))
}

func (c *character) give(to *character) {
	to.leftHand = c.leftHand
	c.leftHand = nil
}

func (c *character) show() string {
	if c.leftHand == nil {
		return c.name + " has no item."
	}

	return c.name + " has " + string(*c.leftHand) + "."
}

func main() {
	var arthur = character{
		name: "Arthur",
	}
	var knight = character{
		name: "Knight",
	}
	var sword item = "sword"

	arthur.pickup(&sword)
	fmt.Println(arthur.show(), knight.show())

	arthur.give(&knight)
	fmt.Println(arthur.show(), knight.show())
}
