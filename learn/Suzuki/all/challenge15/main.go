package main

import "fmt"

type celsius float64
type kelvin float64
type fahrenheit float64

func main() {
	fmt.Println("=======================")
	fmt.Println("|   °C       |   °F   |")
	fmt.Println("=======================")
	var c celsius = -40
	for ; c <= 100; c += 5 {

	}
	fmt.Println("=======================")
}

// 関数
func drawTable(a func() celsius) {
	fmt.Printf("|%-12.1f|%-8.1f|\n", a(), a().fahrenheit())
}

func m40To100celsius() {

}

// メソッド
func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}
func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}
func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit(c*9.0/5.0 + 32.0)
}
func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}
func (k kelvin) fahrenheit() fahrenheit {
	return fahrenheit(k*9.0/5.0 - 459.67)
}
func (f fahrenheit) kelvin() kelvin {
	return kelvin((f + 459.67) * 5.0 / 9.0)
}
