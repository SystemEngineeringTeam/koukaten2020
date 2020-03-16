package main

import (
	"fmt"
	"net/url"
)

func main() {
	_, err := url.Parse("https://a b.com")

	if err != nil {
		fmt.Printf("%#v\n", err)

		fmt.Printf("%#v\n", err.Error())

	}

	if e, ok := err.(*url.Error); ok {
		fmt.Printf("%#v\n", e.Op)
		fmt.Printf("%#v\n", e.Err)
		fmt.Printf("%#v\n", e.URL)
	}

}
