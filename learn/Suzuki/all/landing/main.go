package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type location struct {
	Name string
	Lat  float64 `json:"latitude"`
	Long float64 `json:"longitude"`
}

func main() {
	locations := []location{
		{"Bradbury Landing", -4.5895, 137.4417},
		{"Columbia Memorial Station", -14.5684, 175.472636},
		{"Challenger Memorial Station", -1.9462, 354.4734},
	}
	bytes, err := json.MarshalIndent(locations, "", " ")
	exitOnErr(err)
	fmt.Println(string(bytes))
	// fmt.Println(locations)
}

func exitOnErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
