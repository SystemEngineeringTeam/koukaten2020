package main

import (
	"fmt"
	"math/rand"
	"time"
)

type sol struct {
	day, time int
}

type animaler interface {
	move() string
	eat() string
}

type animal struct {
	name string
}

func (a animal) move() string {

	var str string

	switch rand.Intn(4) {
	case 0:
		str = "walking"
	case 1:
		str = "balking"
	case 2:
		str = "climbing"
	case 3:
		str = "writing code!!!"
	}

	return str

}

func (a animal) eat() string {
	var str string

	switch rand.Intn(5) {
	case 0:
		str = "banana"
	case 1:
		str = "strawberry"
	case 2:
		str = "apple"
	case 3:
		str = "orange"
	case 4:
		str = "grass"
	}

	return str
}

func whitch(a animaler) {
	sol := sol{1, 0}

	for ; sol.day <= 3; sol.day++ {
		sol.time = 0
		for ; sol.time < 24; sol.time++ {

			if sol.time < 7 || sol.time > 19 {
				fmt.Println("Day", sol.day, "Time", sol.time, "sleeping")
			} else {
				switch rand.Intn(2) {
				case 0:
					fmt.Println("Day", sol.day, "Time", sol.time, a.move())
				case 1:
					fmt.Println("Day", sol.day, "Time", sol.time, a.eat())
				}
			}

		}
	}
}

func main() {
	rand.Seed(time.Now().Unix())

	sheep := animal{"Sheep"}
	monkey := animal{"Monkey"}

	fmt.Println(sheep.name)
	whitch(sheep)
	fmt.Println(monkey.name)
	whitch(monkey)
}
