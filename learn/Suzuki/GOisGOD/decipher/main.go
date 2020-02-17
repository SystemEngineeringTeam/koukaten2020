package main

import (
	"fmt"
)

func main() {
	var str string = "CSOITEUIWUIZNSROCNKF"
	vis(str, "GOLANG")
}
func vis(str string, key string) {
	var max int = 'Z'
	var min int = 'A'
	for i, c := range str {
		if c >= int32(min) && c <= int32(max) {
			n := int(key[i%len(key)] - 'A')
			if n < min {
				for ; n >= min; n += max {
				}
			} else if n > max {
				n = n % (max - min + 1)
			}
			c -= int32(n)
			if c < int32(min) {
				for ; c < int32(min); c += int32(max - min + 1) {
				}
			}
			for ; c > int32(max); c -= int32(max - min + 1) {
			}
		}
		fmt.Printf("%c", c)
	}
	fmt.Printf("\n")
}
