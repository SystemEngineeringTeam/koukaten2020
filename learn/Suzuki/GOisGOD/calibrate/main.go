package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64
type sensor func() kelvin

func main() {
	sensor := calibrate(realSensor, 5)
	for i := 0; i < 5; i++ {
		fmt.Println(sensor())
	}
	sensor = calibrate(fakesensor, 5)
	for i := 0; i < 5; i++ {
		fmt.Println(sensor())
	}
}
func realSensor() kelvin {
	return 0
}
func fakesensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}
func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}
