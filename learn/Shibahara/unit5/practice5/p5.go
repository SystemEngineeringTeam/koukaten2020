package main

import "fmt"

var t interface {
	talk() string
}

func main() {
	t = martian{}
	fmt.Println
}
