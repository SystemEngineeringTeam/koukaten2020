package main

import (
	"fmt"
	"net/url"
)

func main() {
	_, url := url.Parse("https://a b.com/")
	fmt.Printf("%#v", url.Error())
}
