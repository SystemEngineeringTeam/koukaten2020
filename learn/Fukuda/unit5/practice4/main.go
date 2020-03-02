package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type coordinate struct {
	d float64 // `json:"degrees"`
	m float64 // `json:"minutes"`
	s float64 // `json:"seconds"`
	h rune    // `json:"hemisphere"`
}

// func (c coordinate) String() string {
// 	return fmt.Sprintf("%v°%v'%.1f %c", c.d, c.m, c.s, c.h)
// }

// func (l location) String() string {
// 	return fmt.Sprintf("%v,%v", l.lat, l.long)
// }

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

func (c coordinate) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("{\"decimal\": %.1f,\"dms\": \"%v°%v'%.1v\" %c\",\"degrees\": %v,\"minutes\": %v,\"seconds\": %v,\"hemisphere\": \"%c\"}", c.decimal(), c.d, c.m, c.s, c.h, c.d, c.m, c.s, c.h)), nil
}

func (l location) MarshalJSON() ([]byte, error) {
	var lat, _ = json.Marshal(l.lat)
	var long, _ = json.Marshal(l.long)
	return []byte(fmt.Sprintf("{\"lat\":%v,\"long\":%v}", lat, long)), nil
}

func main() {
	elysuim := location{
		lat:  coordinate{4, 30, 0.0, 'N'},
		long: coordinate{135, 54, 0.0, 'E'},
	}

	bytes, err := json.Marshal(elysuim)
	exitOnError(err)

	fmt.Println(string(bytes))
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
