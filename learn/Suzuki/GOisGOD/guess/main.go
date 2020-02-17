package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var r, g int
	min := 1
	max := 100
	var i int
	r = rand.Intn(100) + 1
	for i = 1; r != g; i++ {
		g = rand.Intn(max-min+1) + min
		if r == g {
			fmt.Println("あぁぁ入る入る入る…")
		} else if r < g {
			fmt.Println("大きすぎるッピ！", g)
			max = g - 1
		} else {
			min = g + 1
			fmt.Println("小さすぎるッピ！", g)
		}
	}
	fmt.Println(r)
}
