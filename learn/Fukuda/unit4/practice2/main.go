package main

import "fmt"

// Planets is ...
type Planets []string

func (p Planets) terraform() {
	for i := range p {
		p[i] = "New" + p[i]
	}
}

func main() {
	planets := Planets{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	Planets(planets[3:4]).terraform()
	Planets(planets[6:]).terraform()

	fmt.Println(planets)

}
