package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Timer struct {
	hour int
}

type Animal struct {
	name string
}

func main() {
	rand.Seed(time.Now().Unix())
	//set foods
	foods := []string{"food_A", "food_B", "food_C", "food_D", "food_E"}
	//set animals
	animals := []Animal{{"yaju1"}, {"yaju2"}, {"yaju3"}, {"yaju4"}, {"yaju5"}, {"yaju6"}, {"yaju7"}, {"yaju8"}}
	//roop 3 days
	t := Timer{hour: 0}
	for ; t.day() <= 3; t.hour++ {
		if t.hour%24 >= 6 && t.hour%24 <= 18 {
			fmt.Printf("%v ", t)
			chA := rand.Intn(len(animals))
			if rand.Intn(2) == 1 {
				fmt.Printf("%v eat %v\n", animals[chA], animals[chA].eat(foods))
			} else {
				fmt.Printf("%v %v\n", animals[chA], animals[chA].move())
			}
		}
	}
}

//functions

//methods
func (t Timer) day() int {
	return t.hour/24 + 1
}
func (t Timer) String() string {
	return fmt.Sprintf("day%v %v:00", t.day(), t.hour%24)
}
func (a Animal) move() string {

	return "move"
}

func (a Animal) eat(foods []string) string {
	f := rand.Intn(len(foods))
	return foods[f]
}
func (a Animal) String() string {
	return a.name
}
