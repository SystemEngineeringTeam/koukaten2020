package main

import "fmt"

type kelvin float64

type sensor func() kelvin

func realSensor() kelvin {
	return 0
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func main() {
	var num kelvin = 3
	sensor := calibrate(realSensor, num)
	fmt.Println(sensor())
}
