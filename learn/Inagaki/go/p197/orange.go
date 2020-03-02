package main

import "fmt"

func main() {
	posotion := []location{
		newLocation(cordinate{14, 34, 6.2, 'S'}, cordinate{175, 28, 21.5, 'E'}, "Clolumbia"),
		newLocation(cordinate{1, 56, 46.3, 'S'}, cordinate{354, 28, 242, 'E'}, "Challenger"),
	}
	fmt.Println(posotion)
}

type cordinate struct {
	d, m, s float64
	h       rune
}

func (c cordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

type location struct {
	lat, long float64
	name      string
}

func newLocation(lat, long cordinate, name string) location {
	return location{lat.decimal(), long.decimal(), name}
}
