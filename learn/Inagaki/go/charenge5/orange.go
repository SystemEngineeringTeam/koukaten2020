package main

import (
	"fmt"
	"os"
)

const rows int = 9
const columns int = 9

type NewSudoku [rows][columns]int8

func main() {
	s := NewSudoku([rows][columns]int8{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	})
	k := s
	for {
		for i := range s {
			for v := range s[i] {
				if s[i][v] == 0 {
					if k[i][v] == 9 {
						k[i][v] = 1
					} else {
						k[i][v]++
						fmt.Println(k)
						if k.check() == 1 {
							os.Exit(1)
						}
					}
				}
			}
		}
	}
}

//一気に確認
func (N NewSudoku) check() int64 {
	if N.nineNineCheck()+N.juuziCheck()+N.zeroCheck() == 3 {
		return 1
	}
	return 0
}

//0が残ってるか確認
func (N NewSudoku) zeroCheck() int64 {
	for i := range N {
		for v := range N[i] {
			if N[i][v] == 0 {
				return 0
			}
		}
	}
	return 1
}

//十字の確認
func (N NewSudoku) juuziCheck() int64 {
	var mix int
	for i := range N {
		mix = 0
		for v := range N[i] {
			mix += int(N[i][v])
		}
		if mix != 45 {
			return 0
		}
	}
	for i := range N {
		mix = 0
		for v := range N[i] {
			mix += int(N[v][i])
		}
		if mix != 45 {
			return 0
		}
	}
	return 1
}

//３×３マスの確認
func (N NewSudoku) nineNineCheck() int64 {
	var mix int
	three := [3]int{0, 3, 6}
	for i := range three {
		for v := range three {
			mix = 0
			for x := i; x < i+3; x++ {
				for y := v; y < v+3; v++ {
					mix = +int(N[x][y])
				}
			}
			if mix != 45 {
				return 0
			}
		}
	}
	return 1
}
