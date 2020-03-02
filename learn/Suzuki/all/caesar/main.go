package main

import "fmt"

func main() {
	var str string = "L fdph,L vdz,L frqtxhuhg."
	for i := 0; i < len(str); i++ {
		if (str[i] >= 'a' && str[i] <= 'z') || (str[i] >= 'A' && str[i] <= 'Z') {
			a := (str[i] > 'Z')
			c := str[i]
			if a {
				c -= 'a'
			} else {
				c -= 'A'
			}

			c -= 3

			for c > 'z'-'a'+1 {
				c += 'z' - 'a' + 1
			}
			if a {
				c += 'a'
			} else {
				c += 'A'
			}
			fmt.Printf("%c", c)
		} else {
			fmt.Printf("%c", str[i])
		}
	}
}
