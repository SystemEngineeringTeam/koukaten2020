package main

import (
	"fmt"
)

type celsius float64
type kelvin float64
type fahrenheit float64

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit(c*9.0/5.0 + 32.0)
}
func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}
func (k kelvin) fahrenheit() fahrenheit {
	return fahrenheit((k-273.15)*9.0/5.0 + 32.0)
}

func (f fahrenheit) kelvin() kelvin {
	return kelvin((f-32.0)*5.0/9.0 + 273.15)
}
func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

func main() {
	var k kelvin = 0
	var c celsius = 0
	var f fahrenheit = 0
	fmt.Println("0K= ", k.celsius(), "C", "0K=", k.fahrenheit(), "F")
	fmt.Println("0C= ", c.kelvin(), "K", "0C=", c.fahrenheit(), "F")
	fmt.Println("0F= ", f.kelvin(), "K", "0F=", f.celsius(), "C")
}
