package t10

import (
	"advent/common"
	"fmt"
	"math"
	"sort"
)

type Point struct {
	x int
	y int

	a float64
	d bool
}

type Direction struct {
	X float64
	Y float64
}

func Star1(filename string) {
	data := common.ReadLines(filename, "\n")

	Y := len(data)
	X := len(data[0])

	p := make([]Point, 0, 100)
	for i := 0; i < Y; i++ {
		for j := 0; j < X; j++ {
			if data[i][j:j+1] == "#" {
				p = append(p, Point{j, i, 0, false})
			}
		}
	}

	var maxPoint int
	var maxValue int
	for i := range p {
		set := make(map[float64]bool)
		for j := range p {
			dx := float64(p[i].x - p[j].x)
			dy := float64(p[i].y - p[j].y)
			if dx == 0 && dy == 0 {
				continue
			}

			an := math.Atan2(dx, dy)
			set[an] = true
		}

		v := len(set)
		if v > maxValue {
			maxValue = v
			maxPoint = i
		}
	}

	fmt.Printf("Star 1: %12.0d at %d,%d\n", maxValue, p[maxPoint].x, p[maxPoint].y)
}

func Star2(filename string) {
	data := common.ReadLines(filename, "\n")

	Y := len(data)
	X := len(data[0])

	p := make([]Point, 0, 100)
	for i := 0; i < Y; i++ {
		for j := 0; j < X; j++ {
			if data[i][j:j+1] == "#" {
				p = append(p, Point{j, i, 0, false})
			}
		}
	}

	var maxPoint int
	var maxValue int
	for i := range p {
		set := make(map[float64]bool)
		for j := range p {
			dx := float64(p[i].x - p[j].x)
			dy := float64(p[i].y - p[j].y)
			if dx == 0 && dy == 0 {
				continue
			}

			an := math.Atan2(dx, dy)
			set[an] = true
		}

		v := len(set)
		if v > maxValue {
			maxValue = v
			maxPoint = i
		}
	}

	bx := p[maxPoint].x
	by := p[maxPoint].y
	for i := range p {
		p[i].a = 2*math.Pi - math.Atan2(float64(p[i].x-bx), float64(p[i].y-by)) + 0.5*math.Pi
	}
	sort.Slice(p, func(i, j int) bool {
		if p[i].a < p[j].a {
			return true
		}
		if p[i].a > p[j].a {
			return false
		}

		di := (p[i].x-bx)*(p[i].x-bx) + (p[i].y-by)*(p[i].y-by)
		dj := (p[j].x-bx)*(p[j].x-bx) + (p[j].y-by)*(p[j].y-by)
		return di < dj
	})

	i := 0
	j := 0
	fire := -1.0
	for {
		if i == maxPoint {
			continue
		}

		if p[j].a != fire && !p[j].d {
			i++
			p[j].d = true
			fire = p[j].a
		}

		if i == 200 {
			break
		}

		j++
	}

	fmt.Printf("Star 1: %12.0d at %d,%d\n", p[j].x*100+p[j].y, p[j].x, p[j].y)
}
