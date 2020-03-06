package main

import (
	"fmt"
)

type character struct {
	item have
}

type have *string

func main() {
	arthur := character{}
	knight := character{}
	var have have
	*have = "game"
	arthur.pickup(&have)
	arthur.give(&knight)
	fmt.Println(arthur, knight)
}

func (me character) pickup(i *have) {
	me.item = *i
}

func (me character) give(to *character) {
	if me.item == nil {
		fmt.Println("I don't have item")
		return
	}
	*to.item = *me.item
	me.item = nil
}
