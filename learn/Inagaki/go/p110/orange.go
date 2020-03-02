package main

import (
	"fmt"
)

func KelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}
func CeliusToFahrenheit(c float64) float64 {
	return (c * 9.0 / 5.0) - 32.0
}

func KelvinToFahrenheit(k float64) float64 {
	return CeliusToFahrenheit(KelvinToCelsius(k))
}

func main() {
	fmt.Println(KelvinToFahrenheit(0.0))
}
