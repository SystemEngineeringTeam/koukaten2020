package main

import (
	"fmt"
	"strings"
)

func main() {
	var plainText string = "OrangeSter"
	var key string = "Otter"
	var messege string
	var count int = 0

	key = strings.ToLower(key)

	for i := range plainText {
		keypoint := key[count] - 'a'
		if 'A' <= plainText[i] && plainText[i] <= 'Z' {
			if plainText[i]+keypoint > 'Z' {
				messege += string(plainText[i] + keypoint + 'A' - 'Z')
			} else {
				messege += string(plainText[i] + keypoint)
			}
		} else if 'a' <= plainText[i] && plainText[i] <= 'z' {
			if plainText[i]+keypoint > 'z' {
				messege += string(plainText[i] + keypoint + 'a' - 'z')
			} else {
				messege += string(plainText[i] + keypoint)
			}
		} else {
			messege += string(plainText)
		}
		count++
		count = count % len(key)
	}
	fmt.Println(messege)
}
