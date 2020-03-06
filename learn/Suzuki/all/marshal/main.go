package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
)

type coordinate struct {
	D float64 `json:"degrees"`
	M float64 `json:"minutes"`
	S float64 `json:"seconds"`
	H rune    `json:"hemisphere,string"`
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
	dt := coordinate{135, 54, 0, 'E'}
	bytes, err := json.MarshalIndent(dt, "", " ")
	exitOnErr(err)
	fmt.Println(string(bytes))
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
	switch c.H {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.D + c.M/60 + c.S/3600)
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.Lat))
	s2, c2 := math.Sincos(rad(p2.Lat))
	cl := math.Cos(rad(p1.Long - p2.Long))
	return w.radius * math.Acos(s1*s2+c1*c2*cl)
}

func (l location) description() string {
	return fmt.Sprintf("%s (%.1f° %.1f°)", l.name, l.Lat, l.Long)
}

func (c coordinate) String() string {
	return fmt.Sprintf("%.0f°%.0f'%.1f\" %c", c.D, c.M, c.S, c.H)
}

//json.marshaler
func (c coordinate) MarshalJSON() ([]byte, error) {
	type Alias coordinate
	return json.Marshal(&struct {
		Dec float64 `json:"decimal"`
		Str string  `json:"dms"`
		Alias
		Hem string `json:"hemisphere"`
	}{
		Dec:   c.decimal(),
		Str:   c.String(),
		Alias: (Alias)(c),
		Hem:   fmt.Sprintf("%c", c.H),
	})
}
