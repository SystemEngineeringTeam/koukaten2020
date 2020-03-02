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

func main() {
	worlds := map[string]world{}
	points := map[string]location{}
	// input worlds
	worlds["Mercury"] = world{2439.7}
	worlds["Venus"] = world{6051.8}
	worlds["Earth"] = world{6371.0}
	worlds["Mars"] = world{3389.5}
	worlds["Jupiter"] = world{69911}
	worlds["Saturnus"] = world{58232}
	worlds["Uranus"] = world{25362}
	worlds["Neptune"] = world{24622}
	// input points
	points["London"] = changeLocName(newLocation(coordinate{d: 51, m: 30, h: 'N'}, coordinate{d: 0, m: 8, h: 'W'}), "London")
	points["Paris"] = changeLocName(newLocation(coordinate{d: 48, m: 51, h: 'N'}, coordinate{d: 2, m: 21, h: 'E'}), "Paris")
	points["Now"] = location{"Now", 35.185223, 137.115915}
	points["Tokyo"] = changeLocName(newLocation(coordinate{35, 41, 22.4, 'N'}, coordinate{139, 41, 30.2, 'E'}), "Tokyo")
	points["Sharp"] = changeLocName(newLocation(coordinate{5, 4, 48, 'S'}, coordinate{d: 137, m: 51, h: 'E'}), "Sharp")
	points["Olympus"] = changeLocName(newLocation(coordinate{d: 18, m: 39, h: 'N'}, coordinate{d: 226, m: 12, h: 'E'}), "Olympus")
	//output
	fmt.Printf("distance from London to Paris :%.2fkm\n", worlds["Earth"].distance(points["London"], points["Paris"]))
	fmt.Printf("distance from here to Tokyo :%.2fkm\n", worlds["Earth"].distance(points["Now"], points["Tokyo"]))
	fmt.Printf("distance from Sharp to Olympus :%.2fkm\n", worlds["Mars"].distance(points["Sharp"], points["Olympus"]))
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

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.Lat))
	s2, c2 := math.Sincos(rad(p2.Lat))
	cl := math.Cos(rad(p1.Long - p2.Long))
	return w.radius * math.Acos(s1*s2+c1*c2*cl)
}
