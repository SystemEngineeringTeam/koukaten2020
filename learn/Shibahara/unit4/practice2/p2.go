package main

import (
	"fmt"
)

//Planets is string's slice
type Planets []string

func (ps Planets) terraform() Planets {
	for i, p := range ps {
		if p == "Mars" || p == "Uranus" || p == "Neptune" {
			ps[i] = "New_" + p
		}
	}
	return ps
}

func main() {
	planets := Planets{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}
	planets.terraform()
	fmt.Println(planets)
}
