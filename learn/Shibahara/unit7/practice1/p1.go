package main

import (
	"fmt"
)

func main() {
	ch0 := make(chan string)
	ch1 := make(chan string)

	go sourceGopher(ch0)
	go filterGopher(ch0, ch1)
	printGopher(ch1)
}

func sourceGopher(downstream chan string) {
	for _, v := range []string{"Hello", "Hello", "Hello", "Hello"} {
		downstream <- v
	}
	close(downstream)
}

func filterGopher(upstream, downstream chan string) {
	var before string
	for {
		item, ok := <-upstream
		//上流が閉じたらここも閉じる
		if !ok { // okじゃないなら。こんな書き方をする
			close(downstream)
			return
		}
		if item != before {
			downstream <- item
			before = item
		}
	}
}

func printGopher(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}
