package main

import "fmt"

func dump(slice []string) {
	fmt.Println(slice, cap(slice))
}

func main() {
	var array = []string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	slice := array[0:0:0]
	capacity := 0

	for i := 0; i < 8; i++ {
		capacity = cap(slice)
		slice = append(slice, array[i])

		if capacity != cap(slice) {
			dump(slice)
		}
	}

}
