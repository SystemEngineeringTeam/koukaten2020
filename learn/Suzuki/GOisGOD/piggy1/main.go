package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var money float64
	for money < 20 {
		switch rand.Intn(3) {
		case 0:
			money += 0.05
		case 1:
			money += 0.1
		case 2:
			money += 0.25
		}
		fmt.Printf("%.2f\n", money)
	}

}
