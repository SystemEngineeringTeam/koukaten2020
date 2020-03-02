package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const limit = 2000
	var HaveM = 0.0
	rand.Seed(time.Now().UnixNano())
	for limit >= HaveM {
		chance := rand.Intn(3)
		switch chance {
		case 0:
			HaveM += 5
		case 1:
			HaveM += 10
		case 2:
			HaveM += 25
		}
		fmt.Printf("$%.2f\n", HaveM/100)
	}

}
