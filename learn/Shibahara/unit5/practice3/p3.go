package main

import (
	"fmt"
	"math"
)

type world struct {
	radius float64
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

type location struct {
	lat  float64
	long float64
}

type coordinate struct {
	d, m, s float64
	h       rune //N/S/E/W
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}

func main() {
	var mars = world{radius: 3389.5}
	var earth = world{radius: 6371.0}

	spirit := newLocation(coordinate{14, 34, 6.2, 'S'}, coordinate{175, 28, 21.5, 'E'})
	opportunity := newLocation(coordinate{1, 56, 46.3, 'S'}, coordinate{354, 28, 24.2, 'E'})
	curiosity := newLocation(coordinate{4, 35, 22.2, 'S'}, coordinate{354, 28, 24.2, 'E'})
	insight := newLocation(coordinate{4, 30, 0.0, 'N'}, coordinate{134, 54, 0, 'E'})

	fmt.Printf("spirit to opportunity:%.2fkm\n", mars.distance(spirit, opportunity))
	fmt.Printf("spirit to curiosity :%.2fkm\n", mars.distance(spirit, curiosity))
	fmt.Printf("spirit to insight:%.2fkm\n", mars.distance(spirit, insight))

	fmt.Printf("opportunity to curiosity:%.2fkm\n", mars.distance(opportunity, curiosity))
	fmt.Printf("opportunity to insight:%.2fkm\n", mars.distance(opportunity, insight))

	fmt.Printf("curiosity to insight:%.2fkm\n\n", mars.distance(curiosity, insight))

	london := newLocation(coordinate{51, 30, 0, 'N'}, coordinate{0, 8, 0, 'W'})
	Paris := newLocation(coordinate{48, 51, 0, 'N'}, coordinate{2, 21, 0, 'E'})

	fmt.Printf("london to Paris:%.2fkm\n\n", earth.distance(london, Paris))

}
