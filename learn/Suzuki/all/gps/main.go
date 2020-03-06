package main

import (
	"fmt"
	"math"
	"os"
)

type coordinate struct {
	d, m, s float64
	h       rune
}
type world struct {
	radius float64
}

type location struct {
	name string
	Lat  float64 `json:"latitude"`
	Long float64 `json:"longitude"`
}

type gps struct {
	current     location
	destination location
	world       world
}

type rover struct {
	gps
}

func main() {
	g := gps{location{Lat: -4.5895, Long: 137.4417}, location{Lat: 4.5, Long: 135.9}, world{3389.5}}
	curiosity := rover{g}
	//output
	fmt.Println(curiosity.massage())
}

//functions
func exitOnErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func newLocation(lat, long coordinate) location {
	return location{Lat: lat.decimal(), Long: long.decimal()}
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

//methods
func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.Lat))
	s2, c2 := math.Sincos(rad(p2.Lat))
	cl := math.Cos(rad(p1.Long - p2.Long))
	return w.radius * math.Acos(s1*s2+c1*c2*cl)
}

func (g gps) distance() float64 {
	return g.world.distance(g.current, g.destination)
}

func (l location) description() string {
	return fmt.Sprintf("%s (%.1f° %.1f°)", l.name, l.Lat, l.Long)
}

func (g gps) massage() string {
	return fmt.Sprintf("%.1fkm", g.distance())
}
