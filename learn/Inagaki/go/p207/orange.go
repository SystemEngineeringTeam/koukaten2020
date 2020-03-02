package main

import (
	"fmt"
	"math"
)

type location struct {
	name      string
	lat, long float64
}

type world struct {
	radius float64
}

type gps struct {
	current, destination location
	world                world
}

func main() {
	mars := world{3389.5}
	current := location{"Bradbury", -4.5895, 137.4417}
	destination := location{"Elysium", 4.5, 135.9}
	gps := gps{current, destination, mars}
	fmt.Println(gps.message())
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func (g gps) message() string {
	return fmt.Sprintf("%fkm", g.world.distance(g.current, g.destination))
}
