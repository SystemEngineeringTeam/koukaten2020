package main

import (
	"fmt"
	"strings"
)

func main() {
	message := "your message goes here"
	keyword := "golang"
	keyIndex := 0
	cipherText := ""

	message = strings.ToUpper(strings.Replace(message, " ", "", -1))
	keyword = strings.ToUpper(strings.Replace(keyword, " ", "", -1))

	for i := 0; i < len(message); i++ {
		c := message[i]
		if c >= 'A' && c <= 'Z' {
			c -= 'A'
			k := keyword[keyIndex] - 'A'

			c = (c+k)%26 + 'A'

			keyIndex++
			keyIndex %= len(keyword)
		}
		cipherText += string(c)
	}
	fmt.Println(cipherText)
}
