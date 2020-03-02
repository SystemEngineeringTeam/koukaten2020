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
	world
}

type rover struct {
	gps
}

func (l location) descrption() string {
	return fmt.Sprintf("%v (%.1f°, %.1f°)", l.name, l.lat, l.long)
}

func (w world) distance(p1, p2 location) float64 {
	// todo:mars.radiusで距離を計算する
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))

	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func (g gps) message() string {
	return fmt.Sprintf("%v", g.distance(g.current, g.destination))
}

func main() {
	curiosity := rover{gps{location{name: "Bradbury Landing", lat: -4.5895, long: 137.4417}, location{name: "Elysium Planitia", lat: -4.5, long: 135.9}, world{radius: 3389.5}}}

	fmt.Println(curiosity.message())

}
