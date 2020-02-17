package main

import "fmt"

func main() {
	const hoursperday = 24
	distance, days := 56000000, 28

	fmt.Print(distance/(days*hoursperday), "km/h")
}
