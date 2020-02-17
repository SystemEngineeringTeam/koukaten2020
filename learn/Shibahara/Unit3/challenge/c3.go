package main

import "fmt"

const (
	tempFormat = "%.1f"
)

type fahrenheit float64

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32) * 5 / 9)
}

type celsius float64

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit(c*9/5 + 32)
}

func ftoc(raw int) (string, string) {
	f := fahrenheit(raw)
	c := f.celsius()

	cell1 := fmt.Sprintf(tempFormat, f)
	cell2 := fmt.Sprintf(tempFormat, c)

	return cell1, cell2
}
func ctof(raw int) (string, string) {
	c := celsius(raw)
	f := c.fahrenheit()

	cell1 := fmt.Sprintf(tempFormat, c)
	cell2 := fmt.Sprintf(tempFormat, f)
	return cell1, cell2
}

type getTemperature func(raw int) (string, string)

func drawTable(unit1, unit2 string, getT getTemperature) {
	const (
		line       = "================"
		dataFormat = "|%5s|%5s|\n"
	)
	fmt.Println(line)
	fmt.Printf(dataFormat, unit1, unit2)
	fmt.Println(line)

	for i := -40; i <= 100; i += 5 {
		cell1, cell2 := getT(i)
		fmt.Println()
	}
}

func main() {

}
