package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type location struct {
		Name      string
		Lat, Long float64
	}
	position := []location{
		{Name: "Brabury Landing", Lat: -4.5895, Long: 137.4417},
		{Name: "Columbia Memorial", Lat: -14.5684, Long: 175.472636},
		{Name: "Challenger Memorial Station", Lat: -1.9462, Long: 354.4734},
	}

	bytes, err := json.Marshal(position)
	exitOnError(err)

	fmt.Println(string(bytes))
	// fmt.Println(position)
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
