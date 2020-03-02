package main

import (
	"fmt"
)

func dump(slice []string) {
	fmt.Printf("%v len(%v) cap(%v)\n",
		slice, len(slice), cap(slice))
}

func main() {
	dwarfs := []string{"Ceres", "Pluto", "Hanuma", "Makeke", "Eris"}
	dwarfs1 := dwarfs[0:2:2]
	dump(dwarfs1)
}
