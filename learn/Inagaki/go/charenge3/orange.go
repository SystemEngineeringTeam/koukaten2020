package main

import "fmt"

func Line() {
	fmt.Println("====================")
}

func FtoC(F float64) float64 {
	return (F - 32) * 5 / 9
}
func CtoF(C float64) float64 {
	return C*9/5 + 32
}

func View(A, B string, ForC func(temp float64) float64) {
	Line()
	fmt.Printf("| %3.1v    |%4.4v   |\n", A, B)
	Line()
	for i := -40; i <= 100; i++ {
		if i%5 == 0 {
			temp := float64(i)
			fmt.Printf("| %-5.1f  | %-5.1f |\n", temp, ForC(temp))
		}
	}
	Line()
}

func main() {
	View("℃", "℉", FtoC)
	fmt.Println()
	View("℉", "℃", CtoF)
}
