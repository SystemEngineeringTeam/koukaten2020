package main

import "fmt"

type item string
type character struct {
	Name     string
	leftHand *item
}

func main() {
	lake := item("Excalibur")
	king := character{Name: "Arthur"}
	knight := character{Name: "Bedivere"}
	king.pickup(&lake)
	king.give(&knight)
}

func (c *character) pickup(i *item) {
	if c == nil {
		return
	}
	c.leftHand = i
	fmt.Printf("%v pickup \"%v\"\n", c.Name, *i)
}
func (c *character) give(to *character) {
	if c == nil || to == nil {
		return
	}
	to.leftHand = c.leftHand
	c.leftHand = nil
	fmt.Printf("%v give \"%v\" to %v\n", c.Name, *(to.leftHand), to.Name)
}
