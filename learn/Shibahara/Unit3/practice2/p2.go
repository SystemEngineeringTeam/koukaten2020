package main

type celsius float64

func (c celsius) fahrenheit() fahrenheit {
	c = (c * 9.0 / 5.0) + 32.0
	return fahrenheit(c)
}
func (c celsius) kelvin() kelvin {
	c += 273.15
	return kelvin(c)
}

type kelvin float64

func (k kelvin) celsius() celsius {
	k -= 273.15
	return celsius(k)
}
func (k kelvin) fahrenheit() fahrenheit {
	k = (k-273.15)*9/5 + 32
	return fahrenheit(k)
}

type fahrenheit float64

func (f fahrenheit) celsius() celsius {
	f = (f-32)*5/9 + 273.15
	return celsius(f)
}
func (f fahrenheit) kelvin() kelvin {
	f = (f - 32) * 5 / 9
	return kelvin(f)
}

func main() {

}
