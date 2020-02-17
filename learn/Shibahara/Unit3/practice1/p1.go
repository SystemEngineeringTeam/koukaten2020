package main

import "fmt"

func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}
func celsiusToFahrenhait(c float64) float64 {
	c = (c * 9.0 / 5.0) + 32.0
	return c
}
func kelvinToFahrenhait(k float64) float64 {
	k = ((k - 273.15) * 9.0 / 5.0) + 32.0
	return k
}

func main() {
	kelvin := 233.0
	var celsius float64
	fmt.Print("input celsius:")
	fmt.Scan(&celsius)

	fmt.Println("kelvimToCelsius: ", kelvinToCelsius(kelvin), "℃")
	fmt.Println("celsiusToFahrenhait: ", celsiusToFahrenhait(celsius), "K")
	fmt.Print("input kelvin:")
	fmt.Scan(&kelvin)
	fmt.Println("kelvimToFahrenhait: ", kelvinToFahrenhait(kelvin), "F")

}
