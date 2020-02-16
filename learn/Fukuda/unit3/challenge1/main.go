package main

import "fmt"

type celsius float64
type fahrenheit float64

func (c celsius) fahrenheit() celsius {
	c = (c * 9.0 / 5.0) + 32.0
	return c
}

func drawTable(c celsius) {
	fmt.Println("===============")
	fmt.Println("| °C   | °F   |")
	fmt.Println("===============")

	for ; c <= 100; c += 5 {
		fmt.Printf("|%-5.1f |%-5.1f |\n", c, c.fahrenheit())
	}

	fmt.Println("===============")

}

func main() {
	var c celsius = -40

	drawTable(c)

}
