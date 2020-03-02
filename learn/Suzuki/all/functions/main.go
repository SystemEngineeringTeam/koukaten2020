package main

import "fmt"

func main() {
	kelvin := 0.0
	celsius := kelvinToCelsius(kelvin)
	f := celsiusToFahrenheit(celsius)
	fmt.Print(kelvin, "°Kは、", celsius, "°C、", f, "°Fです.")
}
func kelvinToCelsius(k float64) float64 {
	return (k - 273.15)
}
func celsiusToFahrenheit(c float64) float64 {
	return (c*9.0/5.0 + 32.0)
}
