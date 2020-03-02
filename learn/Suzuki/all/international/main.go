package main

import (
	"fmt"
)

func main() {
	var str string = "Hola Estación Espacial Internacional"
	rot(str, 13)
}
func rot(str string, n int) {
	if n < 0 {
		for ; n > 0; n += 0x10FFFF {
		}
	} else {
		n = n % 0x10FFFF
	}
	for _, c := range str {

		c += int32(n)
		for ; c > 0x10FFFF; c -= 0x10FFFF {
		}
		fmt.Printf("%c", c)
	}
	fmt.Printf("\n")
}
