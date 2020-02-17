package main

import (
	"fmt"
	"math/big"
)

func main() {
	lightspeed := big.NewInt(299792) //km/s

	distance := new(big.Int) //km
	distance.SetString("236000000000000000", 10)

	secondPerYear := big.NewInt(60 * 60 * 24 * 365) //year

	seconds := new(big.Int)
	seconds.Div(distance, lightspeed)

	years := new(big.Int)
	years.Div(seconds, secondPerYear)

	fmt.Println(years, "light year")
}
