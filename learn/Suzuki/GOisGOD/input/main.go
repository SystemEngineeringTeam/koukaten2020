package main

import (
	"fmt"
	"strconv"
)

func main() {
	var bl bool
	var str string
	var err error = nil
	str = "f"
	bl, err = strconv.ParseBool(str)
	if err == nil {
		fmt.Println(bl)
	} else {
		fmt.Println("えらー！")
	}
}
