package t11

import (
	"advent/common"
	"fmt"
)

type Point struct {
	X int
	Y int
}

func Star1(filename string) {
	data := common.ReadInt64Lines(filename, ",")

	visited := make(map[Point]int64)
	now := Point{0, 0}
	dir := Point{0, 1}

	mode := true
	vm := NewVM(data, 0, func(i int64) int64 {
		color, ok := visited[now]
		if !ok {
			return 0
		}
		return color
	}, func(o int64) {
		if mode {
			visited[now] = o
		} else {
			if o == 1 {
				dir = Point{dir.Y, -dir.X}
			} else {
				dir = Point{-dir.Y, dir.X}
			}

			now.X += dir.X
			now.Y += dir.Y
		}

		mode = !mode
	})

	for {
		if vm.Run() == -2 {
			break
		}
	}

	fmt.Printf("Star 1: %12.0d\n", len(visited))
}

func Star2(filename string) {
	data := common.ReadInt64Lines(filename, ",")

	visited := make(map[Point]int64)
	now := Point{0, 0}
	dir := Point{0, 1}
	visited[Point{0, 0}] = 1

	mode := true
	vm := NewVM(data, 0, func(i int64) int64 {
		color, ok := visited[now]
		if !ok {
			return 0
		}
		return color
	}, func(o int64) {
		if mode {
			visited[now] = o
		} else {
			if o == 1 {
				dir = Point{dir.Y, -dir.X}
			} else {
				dir = Point{-dir.Y, dir.X}
			}

			now.X += dir.X
			now.Y += dir.Y
		}

		mode = !mode
	})

	for {
		if vm.Run() == -2 {
			break
		}
	}

	var minX, minY, maxX, maxY int

	for p := range visited {
		if p.X < minX {
			minX = p.X
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	for i := maxY; i >= minY; i-- {
		for j := minX; j <= maxX; j++ {
			color, ok := visited[Point{j, i}]
			if !ok || color == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Print("\n")
	}

	fmt.Printf("Star 1: %12.0d\n", len(visited))
}
