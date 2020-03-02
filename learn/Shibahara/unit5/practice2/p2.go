package main

import "fmt"

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

func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}

func main() {
	spirit := newLocation(coordinate{14, 34, 6.2, 'S'}, coordinate{175, 28, 21.5, 'E'})
	opportunity := newLocation(coordinate{1, 56, 46.3, 'S'}, coordinate{354, 28, 24.2, 'E'})
	curiosity := newLocation(coordinate{4, 35, 22.2, 'S'}, coordinate{354, 28, 24.2, 'E'})
	insight := newLocation(coordinate{4, 30, 0.0, 'N'}, coordinate{134, 54, 0, 'E'})

	fmt.Println("spirit:", spirit)
	fmt.Println("opportunity:", opportunity)
	fmt.Println("curiosity:", curiosity)
	fmt.Println("insight:", insight)
}
