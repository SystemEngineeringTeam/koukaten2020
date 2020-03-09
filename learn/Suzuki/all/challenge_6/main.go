package main

import "fmt"

/*型・数値ーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーー*/

const root int = 3
const rows int = root * root
const columns int = root * root
const numbers int = root * root

//Sudoku 数独問題データの型
type Sudoku struct {
	dt  [rows][columns]int8          //数独本体の配列データ
	can [rows][columns][numbers]bool //１〜９まで値挿入可否(true:不可 false:可)
}

/* ほんへ ーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーー*/

func main() {
	//問題
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
	solve(&s)
	//出力
	fmt.Println(s)
}

/* 関数 ーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーー*/

//NewSudoku Sudoku型の生成
func NewSudoku(s [rows][columns]int8) Sudoku {
	return Sudoku{dt: s}
}

//solve Sudoku型の数独を解く(解法は十字と3*3マスのみ)
func solve(s *Sudoku) {
	for r := range s.dt {
		for c := range s.dt[r] {
			s.checkLoops(r, c)
		}
	}
}

/* メソッド ーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーー*/

//Sudoku型のStringer
func (s Sudoku) String() string {
	var str string
	for i, t := range s.dt {
		if i != 0 {
			str += "\n"
		}
		str += fmt.Sprint(t)
	}
	/* デバッグ用フラグ可視化
	str += "\n"
	for i := range s.can {
		for _, t := range s.can[i] {
			for _, c := range t {
				if c {
					str += "x"
				} else {
					str += "o"
				}
			}
			str += " "
		}
		str += "\n"
	}
	*/
	return str
}

//Sudoku型の(r行,c列)に対して候補が一つに絞られた場合、値の決定をし、trueを返す
//値が既に入っている場合その候補データを更新（←今回はしなくても解ける問題なので必要はない）
func (s *Sudoku) checkOne(r, c int) bool {
	a := int(s.dt[r][c])
	if a == 0 {
		n, m := 0, 0
		for i, b := range s.can[r][c] {
			if b == false {
				m = i
				n++
			}
		}
		if n == 1 {
			s.dt[r][c] = int8(m + 1)
			return true
		}
	} else {
		s.can[r][c] = [numbers]bool{true, true, true, true, true, true, true, true, true}
		s.can[r][c][a-1] = false
	}
	return false
}

//Sudoku型の(r行,c列)を基準に十字方向とroot*rootの他のマスを判定・値の代入をする
//値が代入されると、そのマスを基準に同様の判定を行いtrueを返す
func (s *Sudoku) checkLoops(r, c int) bool {
	ret := false
	d := int(s.dt[r][c])
	if d != 0 {
		//横列
		for i, a := range s.dt[r] {
			if i != c {
				s.can[r][i][d-1] = true
				if a == 0 {
					if s.checkOne(r, i) {
						ret = true
						s.checkLoops(r, i)
					}
				}
			}
		}
		//縦列
		for i, a := range s.dt {
			if i != r {
				s.can[i][c][d-1] = true

				if a[c] == 0 {
					if s.checkOne(i, c) {
						ret = true
						s.checkLoops(i, c)
					}
				}
			}
		}
		//箱内
		boxr, boxc := (r/root)*root, (c/root)*root
		for i := boxr; i < boxr+root; i++ {
			if i != r {
				for j := boxc; j < boxr+root; j++ {
					if j != c {
						s.can[i][j][d-1] = true
						if s.checkOne(i, j) {
							ret = true
							s.checkLoops(i, j)
						}
					}
				}
			}
		}
	} else {
		if s.checkOne(r, c) {
			s.checkLoops(r, c)
		}
	}
	return ret
}
