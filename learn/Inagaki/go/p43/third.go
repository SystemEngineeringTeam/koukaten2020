package main

import (
	"fmt"
	"math/rand"
	"time"
)

var era = "AD"

func main() {
	rand.Seed(time.Now().UnixNano())
	i := 0
	for i < 10 {
		yaer := rand.Intn(9999) + 1
		month := rand.Intn(12) + 1
		MonthInDay := 0

		switch month {
		case 2:
			MonthInDay = 28
			if yaer%400 == 0 {
				MonthInDay = 29
			}
			if yaer%4 == 0 && yaer%100 != 0 {
				MonthInDay = 29
			}
		case 4, 6, 9, 11:
			MonthInDay = 30
		default:
			MonthInDay = 31
		}

		day := rand.Intn(MonthInDay) + 1
		fmt.Println(era, yaer, month, day)
		i++
	}
}
