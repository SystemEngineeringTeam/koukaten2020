package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	const distance = 62.1 * 1000 * 1000
	var line = " "
	var days = 0
	var types = " "
	var price = 0
	var speed = 0
	fmt.Printf("%-18v %-4v %-13v %-3v\n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Println("===========================================")
	for a := 0; a < 10; a++ {
		switch rand.Intn(3) {
		case 0:
			line = "Virgin Galactic"
		case 1:
			line = "SpaceX"
		case 2:
			line = "Space Aventures"
		}

		speed = rand.Intn(30-16) + 16 + 1
		days = distance / speed / 24 / 60 / 60

		price = speed + 20

		switch rand.Intn(2) {
		case 0:
			types = "Round-trip"
			price *= 2
		case 1:
			types = "One-trip"
		}

		fmt.Printf("%-18v %-4v %-13v $　%-3v\n", line, days, types, price)
	}
}
