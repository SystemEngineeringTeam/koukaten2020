package main

import (
	"fmt"
	"net/url"
)

func main() {
	_, err := url.Parse("https://a b.com/")
	fmt.Printf("%#v\n", err)

	if e, ok := err.(*url.Error); ok {
		fmt.Println(e.Err)
		fmt.Println(e.Op)
		fmt.Println(e.URL)
	}
}
