package main

import (
	"fmt"
)

type coordinate struct {
	d, m, s float64
	h       rune
}

type location struct {
	name string
	Lat  float64 `json:"latitude"`
	Long float64 `json:"longitude"`
}

func main() {
	locmap := map[string]location{}
	locmap["Spirit"] = changeLocName(newLocation(coordinate{14, 34, 6.2, 's'}, coordinate{175, 28, 21.5, 'E'}), "Columbia Memorial Station")
	locmap["Opportunity"] = changeLocName(newLocation(coordinate{1, 56, 46.3, 's'}, coordinate{354, 28, 24.2, 'E'}), "Challenger Memorial Station")
	locmap["Curiosity"] = changeLocName(newLocation(coordinate{4, 35, 22.2, 's'}, coordinate{137, 26, 30.1, 'E'}), "Bradbury Landing")
	locmap["InSight"] = changeLocName(newLocation(coordinate{4, 30, 0.0, 'n'}, coordinate{135, 54, 0, 'E'}), "Elysium Planitia")
	for str := range locmap {
		loc := locmap[str]
		fmt.Printf("%-12s%-30s%-8.2f%-8.2f\n", str, loc.name, loc.Lat, loc.Long)
	}
}

//functions

func newLocation(lat, long coordinate) location {
	return location{Lat: lat.decimal(), Long: long.decimal()}
}

func changeLocName(l location, str string) location {
	l.name = str
	return l
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
