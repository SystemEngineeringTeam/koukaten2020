package main

import (
	"fmt"
	"math/rand"
	"time"
)

type fennec struct {
	name string
}

func (f fennec) String() string {
	return f.name
}
func (f fennec) move() string {
	return "こゃーん"
}
func (f fennec) eat() string {
	return "くだもの"
}

type serval struct {
	name string
}

func (s serval) String() string {
	return s.name
}
func (s serval) move() string {
	return "にゃー"
}
func (s serval) eat() string {
	return "にく"
}

type racoon struct {
	name string
}

func (r racoon) String() string {
	return r.name
}
func (r racoon) move() string {
	return "なのだ"
}
func (r racoon) eat() string {
	return "なんでも"
}

type animal interface {
	move() string
	eat() string
	String() string
}

func next(a animal) string {
	switch rand.Intn(2) {
	case 0:
		//move method
		return fmt.Sprintf("%s が %s してるよ！かわいいね！", a.String(), a.move())
	default:
		//eat method
		return fmt.Sprintf("%s が %sをたべてるよ！かわいいね！", a.String(), a.eat())
	}
}

func main() {
	rand.Seed(time.Now().Unix())

	animals := []animal{
		fennec{name: "fennec"},
		racoon{name: "araisan"},
		serval{name: "servel-chan"},
	}

	for day := 1; day <= 3; day++ {
		for hour := 1; hour <= 24; hour++ {
			fmt.Printf("%d日目 %d時\n", day, hour)
			if 7 <= hour && hour <= 18 {
				for _, ani := range animals {
					fmt.Println(next(ani))
				}
			} else {
				fmt.Println("すやっすや")
			}
		}
	}
}
