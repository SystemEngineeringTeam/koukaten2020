package main

import "fmt"

type coordinate struct {
	d, m, s float64
	h       rune
}

func (c coordinate) String() string {
	return fmt.Sprintf("%.2f°%.2f’%.2f %c", c.d, c.m, c.s, c.h)
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

type location struct {
	lat, long coordinate
}

func (l location) String() string {
	return fmt.Sprintf("%v, %v", l.lat, l.long)
}

func main() {
	elysium := location{
		lat:  coordinate{4, 30, 0.0, 'N'},
		long: coordinate{135, 54, 0.0, 'E'},
	}

	fmt.Println("Elysium Planitia is at", elysium)
}
