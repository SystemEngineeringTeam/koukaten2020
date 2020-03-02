package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Scanln(&a)
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	fmt.Println(b)
}
