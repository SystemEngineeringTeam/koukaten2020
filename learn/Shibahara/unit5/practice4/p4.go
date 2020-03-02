package main

import (
	"fmt"
	"math"
)

type location struct {
	name string
	lat  float64
	long float64
}

func (l location) description() string {
	return fmt.Sprintf("%v:%.2f,%.2f", l.name, l.lat, l.long)
}

type world struct {
	radius float64
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

type gps struct {
	current     location
	destination location
	world
}

func (g gps) distance() float64 {
	return g.world.distance(g.current, g.destination)
}

func (g gps) message() string {
	return fmt.Sprintf("%.2fkm %v", g.distance(), g.destination.description())
}

type rover struct {
	gps
}

func main() {
	bradbury := location{"Bradbury Landing", -4.5895, 137.4417}
	elysium := location{"Elysium Planitia", 4.5, 135.9}
	mars := world{radius: 3389.5}

	gps := gps{
		current:     bradbury,
		destination: elysium,
		world:       mars,
	}

	curiosity := rover{
		gps: gps,
	}
	fmt.Println(curiosity.message())
}
