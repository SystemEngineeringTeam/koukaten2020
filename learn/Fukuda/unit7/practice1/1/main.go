package main

import (
	"fmt"
)

func sourceGopher(downstream chan string) {
	for _, v := range []string{"hello", "world", "world", "new", "new", "new", "hello new world"} {
		downstream <- v
	}
	close(downstream)
}

func filterGopher(upsteram, downstream chan string) {
	tmp := ""
	for v := range upsteram {
		if v != tmp {
			downstream <- v
			tmp = v
		}
	}

	close(downstream)
}

func printGopher(upsteram chan string) {
	for v := range upsteram {
		fmt.Println(v)
	}
}

func main() {
	c0 := make(chan string)
	c1 := make(chan string)

	go sourceGopher(c0)

	go filterGopher(c0, c1)
	printGopher(c1)
}
