package t13

import (
	"advent/common"
	"fmt"
	"github.com/gdamore/tcell"
)

type Point struct {
	X int
	Y int
}

func NewScreen(x, y int, screen tcell.Screen) *Screen {
	s := Screen{}
	s.X = x
	s.Y = y
	s.data = make([]int, x*y)
	s.screen = screen

	return &s
}

type Screen struct {
	ball   Point
	handle Point

	screen tcell.Screen

	data []int

	X     int
	Y     int
	Score int
}

func (s *Screen) Put(x, y, v int) {
	if x == -1 {
		s.Score = v
	} else {
		s.data[x+s.X*y] = v

		if v == 4 {
			s.ball = Point{x, y}
		} else if v == 3 {
			s.handle = Point{x, y}
		}

		if s.screen != nil {
			s.screen.SetContent(x, y, tiles[v], nil, styles[v])
		}
	}
}

var tiles = map[int]rune{
	0: ' ',
	1: '@',
	2: 'X',
	3: '-',
	4: 'o',
}

var backgroundStyle = tcell.StyleDefault.Background(tcell.ColorBlack)
var styles = map[int]tcell.Style{
	0: backgroundStyle,
	1: backgroundStyle.Foreground(tcell.ColorNames["darkgoldenrod"]),
	2: backgroundStyle.Foreground(tcell.ColorNames["silver"]),
	3: backgroundStyle.Foreground(tcell.ColorNames["red"]),
	4: backgroundStyle.Foreground(tcell.ColorNames["yellow"]),
}

func (s *Screen) Count(c int) int {
	count := 0
	for _, v := range s.data {
		if v == c {
			count += 1
		}
	}

	return count
}

func Star1(filename string) {
	data := common.ReadInt64Lines(filename, ",")

	screen := NewScreen(37, 26, nil)
	vm := NewVM(data, 0)
	go vm.Run()

	done := false
	for {
		select {
		case <-vm.exit:
			done = true
		case x := <-vm.output:
			y := <-vm.output
			v := <-vm.output
			screen.Put(int(x), int(y), int(v))
		}

		if done {
			break
		}
	}

	fmt.Printf("Star 1: %12.0d\n", screen.Count(2))
}

func Star2(filename string) {
	data := common.ReadInt64Lines(filename, ",")

	var term tcell.Screen
	term, _ = tcell.NewScreen()
	term.Init()
	term.SetStyle(tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite))
	term.Clear()

	screen := NewScreen(37, 26, term)

	vm := NewVM(data, 0)
	vm.SetData(0, 2)
	go vm.Run()

	done := false
	for {
		select {
		case <-vm.exit:
			done = true
		case x := <-vm.output:
			y := <-vm.output
			v := <-vm.output
			screen.Put(int(x), int(y), int(v))
		case <-vm.inputRequest:
			if term != nil {
				term.Show()
			}
			if screen.handle.X < screen.ball.X {
				vm.input <- 1
			} else if screen.handle.X > screen.ball.X {
				vm.input <- -1
			} else {
				vm.input <- 0
			}
		}

		if done {
			break
		}
	}

	if term != nil {
		term.Fini()
	}
	fmt.Printf("Star 2: %12.0d\n", screen.Score)
}
