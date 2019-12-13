package t13

import (
	"advent/common"
	"fmt"
)

type Point struct {
	X int
	Y int
}

func NewScreen(x, y int) *Screen {
	s := Screen{}
	s.X = x
	s.Y = y
	s.data = make([]int, x*y)
	s.buffer = make([]int, 0, 3)

	return &s
}

type Screen struct {
	ball     *Point
	lastBall *Point
	handle   Point

	data   []int
	buffer []int
	X      int
	Y      int
	Score  int
}

func (s *Screen) Put(x int64) {
	s.buffer = append(s.buffer, int(x))
	if len(s.buffer) == 3 {
		if s.buffer[0] == -1 {
			s.Score = s.buffer[2]
		} else {
			s.data[s.buffer[0]+s.X*s.buffer[1]] = s.buffer[2]

			if s.buffer[2] == 4 {
				s.lastBall = s.ball
				s.ball = &Point{s.buffer[0], s.buffer[1]}
			}

			if s.buffer[2] == 3 {
				s.handle = Point{s.buffer[0], s.buffer[1]}
			}
		}
		s.buffer = s.buffer[:0]
	}
}

func (s *Screen) Print() {
	for i := 0; i < s.Y; i++ {
		for j := 0; j < s.X; j++ {
			switch s.data[i*s.X+j] {
			case 0:
				fmt.Print(" ")
			case 1:
				fmt.Print("#")
			case 2:
				fmt.Print("@")
			case 3:
				fmt.Print("=")
			case 4:
				fmt.Print("o")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print(s.Score, "\n")
	fmt.Print("\n")
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

	screen := NewScreen(37, 26)
	vm := NewVM(data, 0, nil, func(o int64) {
		screen.Put(o)
	})

	vm.Run(true)

	screen.Print()
	fmt.Printf("Star 1: %12.0d\n", screen.Count(2))
}

func Star2(filename string) {
	data := common.ReadInt64Lines(filename, ",")

	screen := NewScreen(37, 26)
	vm := NewVM(data, 0, func(i int64) int64 {
		if screen.lastBall != nil {
			dx := screen.ball.X - screen.lastBall.X
			tx := screen.ball.X + dx*(screen.Y-common.AbsInt(screen.ball.Y)-(screen.Y-screen.handle.Y)-1)
			if screen.handle.X < tx {
				return 1
			}

			if screen.handle.X > tx {
				return -1
			}

			return 0
		}
		return 0
	}, func(o int64) {
		screen.Put(o)
	})

	vm.SetData(0, 2)
	vm.Run(true)

	screen.Print()
	fmt.Printf("Star 2: %12.0d\n", screen.Count(2))
}
