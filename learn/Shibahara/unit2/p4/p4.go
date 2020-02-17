package main

import "fmt"

func main() {
	message := "L fdph,L vdz,L frqtxhuhg"

	for _, c := range message {
		if 'A' <= c && c <= 'C' || 'a' <= c && c <= 'c' {
			fmt.Printf("%c", c+24)
			continue
		} else if 'D' <= c && c <= 'Z' || 'd' <= c && c <= 'z' {
			fmt.Printf("%c", c-3)
			continue
		}
		fmt.Printf("%c", c)
	}
}
