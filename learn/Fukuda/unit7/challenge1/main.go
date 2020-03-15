package main

import (
	"fmt"
	"image"
	"log"
	"math/rand"
	"sync"
	"time"
)

// Message is
type Message struct {
	Pos       image.Point
	LifeSigns int
	Rover     string
}

const (
	dayLength         = 24 * time.Second
	receiveTimePerDay = 2 * time.Second
)

func earthReceiver(msgc chan []Message) {
	for {
		time.Sleep(dayLength - receiveTimePerDay)
		receiveMarsMessages(msgc)
	}
}

func receiveMarsMessages(msgc chan []Message) {
	finished := time.After(receiveTimePerDay)
	for {
		select {
		case <-finished:
			return
		case ms := <-msgc:
			for _, m := range ms {
				log.Printf("earth received report of life sign level %d from %s at %v", m.LifeSigns, m.Rover, m.Pos)
			}
		}
	}
}

// SensorData is
type SensorData struct {
	LifeSigns int
}

type cell struct {
	groundData SensorData
	occupier   *Occupier
}

// MarsGrid is
type MarsGrid struct {
	bounds image.Rectangle
	mu     sync.Mutex
	cells  [][]cell
}

// NewMarsGrid is
func NewMarsGrid(size image.Point) *MarsGrid {
	grid := &MarsGrid{
		bounds: image.Rectangle{
			Max: size,
		},
		cells: make([][]cell, size.Y),
	}
	for y := range grid.cells {
		grid.cells[y] = make([]cell, size.X)
		for x := range grid.cells[y] {
			cell := &grid.cells[y][x]
			cell.groundData.LifeSigns = rand.Intn(1000)
		}
	}
	return grid

}

// Occupier is
type Occupier struct {
	grid *MarsGrid
	pos  image.Point
}

// RoverDriver is
type RoverDriver struct {
	commandc chan command
	occupier *Occupier
	name     string
	radio    *Radio
}

// NewRoverDriver is
func NewRoverDriver(name string, occupier *Occupier, marsToEarth chan []Message) *RoverDriver {
	r := &RoverDriver{
		commandc: make(chan command),
		occupier: occupier,
		name:     name,
		radio:    NewRadio(marsToEarth),
	}
	go r.drive()
	return r
}

func (r *RoverDriver) drive() {
	log.Printf("%s initial position %v", r.name, r.occupier.Pos())
	direction := image.Point{X: 1, Y: 0}
	updateInterval := 250 * time.Millisecond
	nextMove := time.After(updateInterval)
	for {
		select {
		case c := <-r.commandc:
			switch c {
			case right:
				direction = image.Point{
					X: direction.Y,
					Y: -direction.X,
				}
			case left:
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
			}
			log.Printf("%s new direction %v", r.name, direction)
		case <-nextMove:
			nextMove = time.After(updateInterval)
			newPos := r.occupier.Pos().Add(direction)
			if r.occupier.MoveTo(newPos) {
				log.Printf("%s moved to %v", r.name, newPos)
				r.checkForLife()
				break
			}
			log.Printf("%s blocked trying to move from %v to %v", r.name, r.occupier.Pos(), newPos)
			dir := rand.Intn(3) + 1
			for i := 0; i < dir; i++ {
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
			}
			log.Printf("%s new random direction %v", r.name, direction)

		}
	}
}

func (r *RoverDriver) checkForLife() {
	sensorData := r.occupier.Sense()
	if sensorData.LifeSigns < 900 {
		return
	}

	r.radio.SendToEarth(Message{
		Pos:       r.occupier.Pos(),
		LifeSigns: sensorData.LifeSigns,
		Rover:     r.name,
	})
}

// Sense is
func (o *Occupier) Sense() SensorData {
	o.grid.mu.Lock()
	defer o.grid.mu.Unlock()
	return o.grid.cell(o.pos).groundData
}

// MoveTo is
func (o *Occupier) MoveTo(p image.Point) bool {
	o.grid.mu.Lock()
	defer o.grid.mu.Unlock()
	newCell := o.grid.cell(p)
	if newCell == nil || newCell.occupier != nil {
		return false
	}
	o.grid.cell(o.pos).occupier = nil
	newCell.occupier = o
	o.pos = p
	return true
}

func (g *MarsGrid) cell(p image.Point) *cell {
	if !p.In(g.bounds) {
		return nil
	}
	return &g.cells[p.Y][p.X]
}

// Pos is
func (o *Occupier) Pos() image.Point {
	return o.pos
}

type command int

const (
	right command = 0
	left  command = 1
)

// Radio is
type Radio struct {
	fromRover chan Message
}

// NewRadio is
func NewRadio(toEarth chan []Message) *Radio {
	r := &Radio{
		fromRover: make(chan Message),
	}
	go r.run(toEarth)
	return r
}

func startDriver(name string, grid *MarsGrid, marsToEarth chan []Message) *RoverDriver {
	var o *Occupier
	for o == nil {
		startPoint := image.Point{X: rand.Intn(grid.Size().X), Y: rand.Intn(grid.Size().Y)}
		o = grid.Occupy(startPoint)
	}

	return NewRoverDriver(name, o, marsToEarth)
}

// Size is
func (g *MarsGrid) Size() image.Point {
	return g.bounds.Max
}

// Occupy is
func (g *MarsGrid) Occupy(p image.Point) *Occupier {
	g.mu.Lock()
	defer g.mu.Unlock()
	cell := g.cell(p)
	if cell == nil || cell.occupier != nil {
		return nil
	}
	cell.occupier = &Occupier{
		grid: g,
		pos:  p,
	}
	return cell.occupier
}

// SendToEarth is
func (r *Radio) SendToEarth(m Message) {
	r.fromRover <- m
}

func (r *Radio) run(toEarth chan []Message) {
	var buffered []Message
	for {
		toEarth1 := toEarth
		if len(buffered) == 0 {
			toEarth1 = nil
		}
		select {
		case m := <-r.fromRover:
			buffered = append(buffered, m)
		case toEarth1 <- buffered:
			buffered = nil
		}
	}
}

// Left is
func (r *RoverDriver) Left() {
	r.commandc <- left
}

// Right is
func (r *RoverDriver) Right() {
	r.commandc <- right
}

func main() {
	marsToEarth := make(chan []Message)
	go earthReceiver(marsToEarth)

	gridSize := image.Point{X: 20, Y: 10}
	grid := NewMarsGrid(gridSize)
	rover := make([]*RoverDriver, 5)
	for i := range rover {
		rover[i] = startDriver(fmt.Sprint("rover", i), grid, marsToEarth)
	}
	time.Sleep(60 * time.Second)
}
