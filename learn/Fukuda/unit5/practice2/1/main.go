package main

import (
	"fmt"
	"math"
)

type coordinate struct {
	d, m, s float64
	h       rune
}

type location struct {
	lat, long float64
}

type world struct {
	radius float64
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}

func (w world) distance(p1, p2 location) float64 {
	// todo:w.radiusで距離を計算する
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))

	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func main() {
	Spirit := newLocation(coordinate{14, 34, 6.2, 'S'}, coordinate{175, 28, 21.5, 'E'})
	Opportunity := newLocation(coordinate{1, 56, 46.3, 'S'}, coordinate{354, 28, 24.2, 'E'})
	Curiosity := newLocation(coordinate{4, 35, 22.2, 'S'}, coordinate{137, 26, 30.1, 'E'})
	InSight := newLocation(coordinate{4, 30, 0.0, 'N'}, coordinate{135, 54, 0, 'E'})
	fmt.Println("Spirit", Spirit)
	fmt.Println("Opportunity", Opportunity)
	fmt.Println("Curiosity", Curiosity)
	fmt.Println("InSight", InSight)

}
