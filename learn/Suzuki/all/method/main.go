package main

import "fmt"

type celsius float64
type kelvin float64
type fahrenheit float64

func main() {
	var k kelvin = 0.0
	c := k.celsius()
	f := k.fahrenheit()
	fmt.Print(k, "°Kは、", c, "°C、", f, "°Fです.")
}

//関数
func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}
func celsiusTokelvin(c celsius) kelvin {
	return kelvin(c + 273.15)
}
func celsiusToFahrenheit(c celsius) fahrenheit {
	return fahrenheit(c*9.0/5.0 + 32.0)
}
func fahrenheitTocelsius(f fahrenheit) celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}
func kelvinToFahrenheit(k kelvin) fahrenheit {
	return fahrenheit(k*9.0/5.0 - 459.67)
}
func fahrenheitToKelvin(f fahrenheit) kelvin {
	return kelvin((f + 459.67) * 5.0 / 9.0)
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
