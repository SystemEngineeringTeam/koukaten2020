package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)

	var torf bool

	switch str {
	case "true", "yes", "1":
		torf = true
	case "false", "no", "0":
		torf = false
	default:
		fmt.Println("error")
	}

	fmt.Println(torf)

}
