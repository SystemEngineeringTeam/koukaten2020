package main

import (
	"fmt"
	"image"
	"log"
	"time"
)

type command int

const (
	right = command(0)
	left  = command(1)
	start = command(2)
	stop  = command(3)
)

// RoverDriver is
type RoverDriver struct {
	commandc chan command
}

// Left is
func (r *RoverDriver) Left() {
	r.commandc <- left
}

// Right is
func (r *RoverDriver) Right() {
	r.commandc <- right
}

// Start is
func (r *RoverDriver) Start() {
	r.commandc <- start
}

// Stop is
func (r *RoverDriver) Stop() {
	r.commandc <- stop
}

func (r *RoverDriver) drive() {
	pos := image.Point{X: 0, Y: 0}
	direction := image.Point{X: 1, Y: 0}

	isMove := true

	updateInterval := 250 * time.Millisecond
	nextMove := time.After(updateInterval)
	for {
		select {
		case c := <-r.commandc:
			switch c {
			case right:
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
			case left:
				direction = image.Point{
					X: direction.Y,
					Y: -direction.X,
				}

			case stop:
				isMove = false
			case start:
				isMove = true
			}
			log.Printf("new direction %v", direction)

		case <-nextMove:

			if isMove {
				pos = pos.Add(direction)

			}

			log.Printf("moved to %v", pos)
			nextMove = time.After(updateInterval)
		}
	}
}

// NewRoverDriver is
func NewRoverDriver() *RoverDriver {
	r := &RoverDriver{
		commandc: make(chan command),
	}
	go r.drive()
	return r

}

func worker() {
	pos := image.Point{X: 10, Y: 10}
	direction := image.Point{X: 1, Y: 0}

	next := time.After(time.Second)
	delay := time.Second

	for {
		select {
		case <-next:
			pos = pos.Add(direction)
			fmt.Println("current position is ", pos)
			delay += time.Second / 2
			next = time.After(delay)

		}
	}
}

func main() {
	r := NewRoverDriver()
	time.Sleep(3 * time.Second)
	r.Left()
	time.Sleep(3 * time.Second)
	r.Right()
	time.Sleep(3 * time.Second)
	r.Stop()
	time.Sleep(3 * time.Second)
	r.Start()
	time.Sleep(3 * time.Second)

}
