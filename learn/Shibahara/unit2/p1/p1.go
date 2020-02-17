package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	var piggy float32

	for piggy < 20 {

		switch rand.Intn(3) {
		case 0:
			piggy += 0.05
		case 1:
			piggy += 0.10
		case 2:
			piggy += 0.25
		}
		fmt.Printf("$%.2f\n", piggy)
	}
	println("saving stop")
}
