package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	text := ""
	keyIndex := 0

	for i := 0; i < utf8.RuneCountInString(cipherText); i++ {
		c := cipherText[i] - 'A'
		k := keyword[keyIndex] - 'A'

		c = (c-k+26)%26 + 'A'
		text += string(c)

		keyIndex++
		keyIndex %= utf8.RuneCountInString(keyword)
	}

	fmt.Println(text)
}
