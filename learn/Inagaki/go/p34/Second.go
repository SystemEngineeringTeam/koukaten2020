package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	const right = 88
	for answer := 0; answer != right; answer = rand.Intn(100) + 1 {
		fmt.Println(answer)
		if answer > right {
			fmt.Println("bigger")
		} else if answer < right {
			fmt.Println("smaller")
		}
	}
	fmt.Println("Ok")
}
