package main

import "fmt"

//Planets is slice of strings
type Planets []string

func main() {
	planets := Planets{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	fmt.Println(planets)
	planets.terraform()
	fmt.Println(planets)

}

//methods
func (pla Planets) terraform() {
	tf := Planets{"Mars", "Uranus", "Neptune"}
	for i := range pla {
		for j := range tf {
			if pla[i] == tf[j] {
				pla[i] = "New " + pla[i]
				break
			}
		}
	}
}
