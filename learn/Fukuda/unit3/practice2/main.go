package main

import "fmt"

type celcius float64
type fahrenheit float64
type kelvin float64

func (c celcius) fahrenheit() celcius {
	c = (c * 9.0 / 5.0) + 32.0
	return c
}
func (c celcius) kelvin() celcius {
	c += 273.15
	return c
}

func (f fahrenheit) celcius() fahrenheit {
	f = (f - 32) * 5.0 / 9.0
	return f
}

func (f fahrenheit) kelvin() fahrenheit {
	f = (f-32)*5.0/9.0 + 273.15
	return f
}

func (k kelvin) celcius() kelvin {
	k -= 273.15
	return k
}

func (k kelvin) fahrenheit() kelvin {
	k -= 273.15
	k = (k * 9.0 / 5.0) + 32.0
	return k
}

func main() {
	var c celcius = 0
	var k kelvin = 0
	var f fahrenheit = 0

	fmt.Println(c.fahrenheit(), c.kelvin())
	fmt.Println(k.celcius(), k.fahrenheit())
	fmt.Println(f.celcius(), f.kelvin())
}
