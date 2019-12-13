package t13

import (
	"advent/common"
	"fmt"
	"github.com/gdamore/tcell"
	"time"
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
	s.buffer = make([]int, 0, 3)
	s.screen = screen

	return &s
}

type Screen struct {
	ball   Point
	handle Point

	screen tcell.Screen

	data   []int
	buffer []int

	X     int
	Y     int
	Score int
}

func (s *Screen) Put(x int64) {
	s.buffer = append(s.buffer, int(x))
	if len(s.buffer) == 3 {
		x := s.buffer[0]
		y := s.buffer[1]
		v := s.buffer[2]

		if x == -1 {
			s.Score = v
		} else {
			s.data[x+s.X*y] = v

			if s.buffer[2] == 4 {
				s.ball = Point{x, y}
			} else if s.buffer[2] == 3 {
				s.handle = Point{x, y}
			}

			if s.screen != nil {
				s.screen.SetContent(x, y, tiles[v], nil, styles[v])
			}
		}
		s.buffer = s.buffer[:0]
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
	vm := NewVM(data, 0, nil, func(o int64) {
		screen.Put(o)
	})

	vm.Run(true)

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

	vm := NewVM(data, 0, func(i int64) int64 {
		if term != nil {
			term.Show()
			time.Sleep(time.Millisecond * 1)
		}

		if screen.handle.X < screen.ball.X {
			return 1
		}

		if screen.handle.X > screen.ball.X {
			return -1
		}

		return 0
	}, func(o int64) {
		screen.Put(o)
	})

	vm.SetData(0, 2)
	vm.Run(true)

	if term != nil {
		term.Fini()
	}
	fmt.Printf("Star 2: %12.0d\n", screen.Score)
}
