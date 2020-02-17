package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var num, comNum int
	rand.Seed(time.Now().Unix())

	fmt.Println("input number(1-10):")
	fmt.Scan(&num)

	for {
		comNum = rand.Intn(11) + 1

		if comNum > num {
			fmt.Printf("%d is too big\n", comNum)
		} else if comNum < num {
			fmt.Printf("%d is too small\n", comNum)
		} else {
			fmt.Println("hit!")
			break
		}
	}

}
