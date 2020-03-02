package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const limit = 20
	HaveM := 0.0
	rand.Seed(time.Now().UnixNano())
	for limit >= HaveM {
		chance := rand.Intn(3)
		switch chance {
		case 0:
			HaveM += 0.05
		case 1:
			HaveM += 0.10
		case 2:
			HaveM += 0.25
		}
		fmt.Printf("%.2f\n", HaveM)
	}

}
