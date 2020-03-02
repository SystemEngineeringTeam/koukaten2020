package main

import (
	"fmt"
)

func main() {
	keyCode := "L fdph,L vdz, L frqtxhuhg"
	for i := 0; i < (len(keyCode)); i++ {
		c := keyCode[i]
		if 'A' <= keyCode[i] && keyCode[i] <= 'Z' {
			c -= 3
			if c < 'A' {
				c -= 26
			}
		} else if 'a' <= keyCode[i] && keyCode[i] <= 'z' {
			c -= 3
			if c < 'a' {
				c -= 26
			}
		}
		fmt.Printf("%c", c)
	}
}
