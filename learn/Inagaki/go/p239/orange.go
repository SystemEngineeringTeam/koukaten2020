package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 亀はturtle
type tartle struct {
	x, y point
}
type point int

func main() {
	rand.Seed(time.Now().UnixNano())
	kame := tartle{
		x: 4, y: 4,
	}
	for i := 0; i < 10; i++ {
		kame.show()
		kame.run()
	}
}

func (t *tartle) run() {
	t.x += arrow()
	t.x.over()
	t.y += arrow()
	t.x.over()
}

func (i *point) over() {
	if *i < 0 {
		*i = 6
	} else if *i > 6 {
		*i = 0
	}
}

func arrow() point {
	if rand.Intn(2) == 0 {
		return -1
	}
	return 1
}

func (t tartle) show() {
	for i := 0; i < 7; i++ {
		for m := 0; m < 7; m++ {
			if int(t.x) == i && int(t.y) == m {
				fmt.Printf("0")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Println("================================")
}
