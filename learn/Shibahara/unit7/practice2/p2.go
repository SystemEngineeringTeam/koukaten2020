package main

import (
	"fmt"
	"image"
	"time"
)

func main() {
	go worker()
	time.Sleep(time.Second * 10)
}

func worker() {
	pos := image.Point{X: 10, Y: 10}

	direction := image.Point{X: 1, Y: 0}

	delay := time.Second / 2
	next := time.After(delay)
	for {
		select {
		case <-next:
			pos = pos.Add(direction)

			fmt.Println("current position is", pos)

			next = time.After(delay)
			delay += time.Second / 2

		}
	}
}