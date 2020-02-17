package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	text := ""

	for i, j := 0, 0; i < utf8.RuneCountInString(cipherText); i++ {
		text += string((cipherText[i]-keyword[j]+26)%26 + 65)
		j++
		if j >= utf8.RuneCountInString(keyword) {
			j = 0
		}
	}

	fmt.Println(text)

}
