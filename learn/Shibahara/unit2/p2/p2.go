package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	var piggy int

	for piggy < 2000 {

		switch rand.Intn(3) {
		case 0:
			piggy += 5
		case 1:
			piggy += 10
		case 2:
			piggy += 25
		}

		fmt.Printf("$%d.%02d\n", piggy/100, piggy%100)
	}
	println("saving stop")
}
