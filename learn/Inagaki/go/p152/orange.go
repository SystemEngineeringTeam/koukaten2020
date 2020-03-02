package main

import "fmt"

type PLANET []string

func (pranets PLANET) terraform() {
	for i := range pranets {
		pranets[i] = "New " + pranets[i]
	}
}

func main() {
	planet := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}
	PLANET(planet[3:4]).terraform()
	PLANET(planet[6:]).terraform()
	fmt.Println(planet)
}
