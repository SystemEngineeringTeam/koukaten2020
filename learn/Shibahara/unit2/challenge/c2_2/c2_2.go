package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	inputText := "your message goes here"
	plainText := ""
	keyword := "GOLANG"
	text := ""

	plainText = strings.ToUpper(strings.Replace(inputText, " ", "", -1))
	fmt.Println(plainText)

	for i, j := 0, 0; i < utf8.RuneCountInString(plainText); i++ {
		text += string((plainText[i]+keyword[j])%26 + 65)
		j++
		if j >= utf8.RuneCountInString(keyword) {
			j = 0
		}
	}

	fmt.Println(text)

}
