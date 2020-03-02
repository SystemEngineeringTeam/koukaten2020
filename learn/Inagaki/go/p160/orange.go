package main

import "fmt"

func main() {
	mama := make([]string, 0, 5)
	var mamaNum int
	for i := 0; i < 10000; i++ {
		mamaNum = cap(mama)
		mama = append(mama, "MAMA", "mama")
		if cap(mama) != mamaNum {
			fmt.Println(cap(mama))
		}
	}
}
