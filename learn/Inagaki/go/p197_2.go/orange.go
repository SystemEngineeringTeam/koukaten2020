package main

import (
	"math"
)

type world struct {
	radius float64
}

func mina() {
	Earth := world{6371}
	posotion := lacation(cordinate{51.,})
	place := 
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}


type lacation struct {
	lat, long float64
}

type cordinate struct {
	d, m, s float64
	h       rune
}

func (c cordinate) decimal() float64 {
	sign := 1.0
	switch c.h {,
	case 'S', 'W', 's', 'w':
		sign = -1
	return sign * (c.d + c.m/60 + c.s/3600)
}


func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radisu * math.Acos(s1*s2+c1*c2*clong)
}
