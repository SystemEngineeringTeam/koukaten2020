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
	// todo:mars.radiusで距離を計算する
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))

	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func main() {
	columbiaMemorialStation := newLocation(coordinate{14, 34, 6.2, 'S'}, coordinate{175, 28, 21.5, 'E'})
	challengerMemorialStation := newLocation(coordinate{1, 56, 46.3, 'S'}, coordinate{354, 28, 24.2, 'E'})
	bradburyLanding := newLocation(coordinate{4, 35, 22.2, 'S'}, coordinate{137, 26, 30.1, 'E'})
	elysuimPlanitia := newLocation(coordinate{4, 30, 0.0, 'N'}, coordinate{135, 54, 0, 'E'})

	spirit := location{-14.5684, 175.472636}
	opportunity := location{-1.9462, 354.4734}
	curiosity := location{-4.5895, 137.4417}
	insight := location{4.5, 135.9}

	mars := world{3389.5}

	fmt.Println("Spirit", "Columbia Memorial Station", mars.distance(spirit, columbiaMemorialStation))

	fmt.Println("Opportunity", "Challenger Memorial Station", mars.distance(opportunity, challengerMemorialStation))

	fmt.Println("Curiosity", "Bradbury Landing", mars.distance(curiosity, bradburyLanding))
	fmt.Println("InSight", "Elysuim Plantia", mars.distance(insight, elysuimPlanitia))

}
