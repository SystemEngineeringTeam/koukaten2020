package main

import (
	"fmt"
	"math/rand"
	"time"
)

type ticket struct {
	Spaceline string
	Days      int
	triptype  int
	Price     int
}

func main() {
	rand.Seed(time.Now().Unix())
	l := 62100000
	//list
	fmt.Printf("%-17s %-5s %-12s %-5s\n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Println("==========================================")
	// randomtickets
	var tkt ticket
	var spd int
	var str string
	for i := 0; i < 10; i++ {
		tkt.triptype = rand.Intn(2)
		switch tkt.triptype {
		case 0:
			str = "Round-trip"
		case 1:
			str = "One-way"
		}
		switch rand.Intn(3) {
		case 0:
			tkt.Spaceline = "Space Adventures"
		case 1:
			tkt.Spaceline = "SpaceX"
		case 2:
			tkt.Spaceline = "Virgin Galactic"
		}
		spd = rand.Intn(15) + 16
		tkt.Days = (l / 86400) / spd
		tkt.Price = spd + 20
		if tkt.triptype == 1 {
			tkt.Price *= 2
		}
		fmt.Printf("%-17s %-5d %-12s %-5d\n", tkt.Spaceline, tkt.Days, str, tkt.Price)
	}
}
