package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var l float64
	var lightyear float64
	l = 236000000000000000
	lightyear = l / 9460730472580800
	fmt.Printf("%.2fly\n", lightyear)
}
