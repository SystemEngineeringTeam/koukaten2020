package main

import "fmt"

type celsius float64
type kelvin float64
type fahrenheit float64

type tempfunction func(a float64) (string, string)

func main() {
	drawTable("C°", "F°", c2f)
	drawTable("F°", "C°", f2c)
}

// 関数
func drawTable(s1, s2 string, t tempfunction) {
	fmt.Println("=======================")
	fmt.Printf("|%-12s|%-8s|\n", s1, s2)
	fmt.Println("=======================")
	var c float64 = -40
	for ; c <= 100; c += 5 {
		s1, s2 := t(c)
		fmt.Printf("|%-12s|%-8s|\n", s1, s2)
	}
	fmt.Println("=======================")
}

func m40To100celsius() {

}

func c2f(a float64) (string, string) {
	c := celsius(a)
	return fmt.Sprintf("%-.1f", c), fmt.Sprintf("%-.1f", c.fahrenheit())
}
func f2c(a float64) (string, string) {
	f := fahrenheit(a)
	return fmt.Sprintf("%-.1f", f), fmt.Sprintf("%-.1f", f.celsius())
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
