package main

import (
	"fmt"
	"strings"
)

func sourceGopher(downstream chan string) {
	for _, v := range []string{"As far as eye could reach he saw nothing but the stems of the great plants about him receding in the violet shade and far overhead the multiple transparency of huge leaves filtering the sunshine to the solemn splendour of twilight in which he walked Whenever he felt able he ran again the ground continued soft and springy covered with the same resilient weed which was the first thing his hands had touched in Malacandra Once or twice a small red creature scuttled across his path but otherwise there seemed to be no life stirring in the wood nothing to fear—except the fact of wandering unprovisioned and alone in a forest of unknown vegetation thousands or millions of miles beyond the reach or knowledge of man"} {
		downstream <- v
	}
	close(downstream)
}

func filterGopher(upsteram, downstream chan string) {
	for v := range upsteram {
		str := strings.Fields(v)
		// fmt.Println(str)

		for _, s := range str {
			downstream <- s
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
