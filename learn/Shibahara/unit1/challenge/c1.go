package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const distance = 62100000
	var spaceline, triptype string
	var speed, days, price int

	rand.Seed(time.Now().Unix())

	fmt.Printf("%-18s%-6s%-13s%-5s\n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Println("==========================================")
	// fmt.Println("Virgin Galactic   23    Round-trip   $ 96")

	for i := 0; i < 10; i++ {
		//航宙会社を選ぶ
		switch rand.Intn(3) {
		case 0:
			spaceline = "Space Adventures"
		case 1:
			spaceline = "SpaceX"
		case 2:
			spaceline = "Virgin Galactic"
		}

		speed = rand.Intn(15) + 16
		price = speed + 20
		days = distance / (speed * 60 * 60 * 24)

		switch rand.Intn(2) {
		case 0:
			triptype = "Round-trip"
			price *= 2
		case 1:
			triptype = "One-way"
		}

		fmt.Printf("%-18s%-6d%-13s$ %-5d\n", spaceline, days, triptype, price)

	}
}
