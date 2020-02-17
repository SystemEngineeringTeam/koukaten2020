package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var money int
	var ans float64
	for money < 2000 {
		switch rand.Intn(3) {
		case 0:
			money += 5
		case 1:
			money += 10
		case 2:
			money += 25
		}
		ans = float64(money) / 100
		fmt.Printf("$%.2f\n", ans)
	}

}
