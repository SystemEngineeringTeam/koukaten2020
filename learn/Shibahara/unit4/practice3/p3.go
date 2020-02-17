package main

import "fmt"

func main() {
	slice := []int{}
	beforeCap := cap(slice)

	for i := 0; i < 100; i++ {
		beforeCap = cap(slice)

		slice = append(slice, i)
		if cap(slice) > beforeCap {
			fmt.Println("cap:", cap(slice))
		}
	}
}
