package main

import (
	"image"
	"log"
	"sync"
	"time"
)

func main() {
	r := NewRoverDriver()
	time.Sleep(3 * time.Second)
	r.Left()
	time.Sleep(3 * time.Second)
	r.Right()
	time.Sleep(3 * time.Second)
}

type command int

const (
	right = command(0)
	left  = command(1)
	stop  = command(2)
)

//RoverDriver is
type RoverDriver struct {
	commandc chan command
}

//NewRoverDriver is
func NewRoverDriver() *RoverDriver {
	r := &RoverDriver{
		commandc: make(chan command),
	}
	go r.drive()
	return r
}

func (r *RoverDriver) drive() {

	pos := image.Point{X: 0, Y: 0}

	direction := image.Point{X: 1, Y: 0}

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
			}
			log.Printf("new direction %v", direction)

		case <-nextMove:
			pos = pos.Add(direction)
			log.Printf("moved to %v", pos)
			nextMove = time.After(updateInterval)
		}
	}
}

//Left is
func (r *RoverDriver) Left() {
	r.commandc <- left
}

//Right is
func (r *RoverDriver) Right() {
	r.commandc <- right
}

//MarsGrid is
type MarsGrid struct {
	mu   sync.Mutex
	grid Occupier
}

//Occupy is
func (g *MarsGrid) Occupy(p image.Point) *Occupier {

}

//Occupier is
type Occupier struct {
	cell image.Point
}

//MoveTo is
func (g *Occupier) MoveTo(p image.Point) bool {

}
