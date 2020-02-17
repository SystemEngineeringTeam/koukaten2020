package main

import (
	"fmt"
	"math/rand"
	"time"
)

var era = "AD"

func main() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		year := rand.Intn(114514) + 1
		month := rand.Intn(12) + 1
		daysInMonth := 31
		switch month {
		case 2:
			if uru(year) == 1 {
				daysInMonth = 28
			} else {
				daysInMonth = 29
			}
		case 4, 6, 9, 11:
			daysInMonth = 30
		}
		day := rand.Intn(daysInMonth) + 1
		fmt.Println(era, year, month, day)
	}
}

func uru(a int) int {
	if a%100 == 0 {
		if (a/100)%4 == 0 {
			return 1
		} else {
			return 0
		}
	} else if a%4 == 0 {
		return 1
	} else {
		return 0
	}
}
