package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64

type sensor func() kelvin

func realSensor() kelvin {
	return 0
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func main() {
	var val kelvin = 3
	sensor := calibrate(fakeSensor, val)
	fmt.Println(sensor())
	val += 9
	fmt.Println(sensor())
}
