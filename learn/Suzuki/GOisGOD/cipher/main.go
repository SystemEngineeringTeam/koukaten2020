package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "your message goes here"
	keyword := "GOLANG"
	s := ciph(strings.ToUpper(strings.Replace(str, " ", "", -1)), keyword)
	fmt.Println(str)
	fmt.Println(s)
	fmt.Println(deciph(s, keyword))
}
func ciph(str string, key string) string {
	var max int = 'Z'
	var min int = 'A'
	var ans string
	for i, c := range str {
		if c >= int32(min) && c <= int32(max) {
			n := int(key[i%len(key)] - 'A')
			if n < min {
				for ; n >= min; n += max {
				}
			} else if n > max {
				n = n % (max - min + 1)
			}
			c += int32(n)
			if c < int32(min) {
				for ; c < int32(min); c += int32(max - min + 1) {
				}
			}
			for ; c > int32(max); c -= int32(max - min + 1) {
			}
		}
		ans = ans + string(c)
	}
	return ans
}
func deciph(str string, key string) string {
	var max int = 'Z'
	var min int = 'A'
	var ans string
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
		ans = ans + string(c)
	}
	return ans
}
