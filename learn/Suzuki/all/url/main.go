package main

import (
	"fmt"
	"net/url"
)

func main() {
	Url := "https://a b.com/"
	_, err := url.Parse(Url)
	fmt.Printf("%#v\n", err)
	if Errs, ok := err.(*url.Error); ok {
		fmt.Println(Errs)
	}
}
