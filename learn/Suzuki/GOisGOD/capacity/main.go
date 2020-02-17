package main

import "fmt"

func main() {
	dwarfs := []string{}
	l := cap(dwarfs)
	for i := 0; i < 100; i++ {
		dwarfs = append(dwarfs, "a")
		if l != cap(dwarfs) {
			dump("dwarfs", dwarfs)
		}
		l = cap(dwarfs)
	}
}

func dump(label string, slice []string) {
	fmt.Printf("%v:長さ %v,容量 %v %v\n", label, len(slice), cap(slice), slice)
}
