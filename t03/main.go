package t03

import (
	"advent/common"
	"fmt"
	"strconv"
	"strings"
)

type Segment struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func NewSegments(cm []string) []Segment {
	out := make([]Segment, len(cm))
	var x1, y1, x2, y2 int

	for i, s := range cm {
		step, _ := strconv.Atoi(s[1:])
		switch s[0:1] {
		case "U":
			x2 = x1
			y2 = y1 + step
			out[i] = Segment{x1, y1, x2, y2}
		case "D":
			x2 = x1
			y2 = y1 - step
			out[i] = Segment{x2, y2, x1, y1}
		case "R":
			x2 = x1 + step
			y2 = y1
			out[i] = Segment{x1, y1, x2, y2}
		case "L":
			x2 = x1 - step
			y2 = y1
			out[i] = Segment{x2, y2, x1, y1}
		}
		x1 = x2
		y1 = y2
	}

	return out
}

func (a *Segment) Intersect(b *Segment) (bool, int, int) {
	ax := a.y1 == a.y2
	bx := b.y1 == b.y2

	if ax == bx {
		return false, 0, 0
	}

	if bx {
		return a.intersectY(b)
	} else {
		return b.intersectY(a)
	}
}

func (a *Segment) Length() int {
	return a.y2 - a.y1 + a.x2 - a.x1
}

func (a *Segment) DirectDist(x, y int) int {
	if a.x1 == a.x2 && a.x1 == x {
		return y - a.y1
	}

	if a.y1 == a.y2 && a.y1 == y {
		return x - a.x1
	}

	return 0
}

func (a *Segment) intersectY(b *Segment) (bool, int, int) {
	//  a - vertical, b - horizontal
	if a.y1 < b.y1 && a.y2 > b.y1 && b.x1 < a.x1 && b.x2 > a.x1 {
		return true, a.x1, b.y1
	}

	return false, 0, 0
}

func Star1(filename string) {
	paths := common.ReadLines(filename, "\n")

	a := NewSegments(strings.Split(paths[0], ","))
	b := NewSegments(strings.Split(paths[1], ","))

	var dist int
	for _, i := range a {
		for _, j := range b {
			res, tx, ty := i.Intersect(&j)
			if res {
				tist := common.AbsInt(tx) + common.AbsInt(ty)
				if tist == 0 {
					continue
				}
				if dist == 0 || tist < dist {
					dist = tist
				}
			}
		}
	}

	fmt.Printf("Star 1: %12.0d\n", dist)
}

func Star2(filename string) {
	paths := common.ReadLines(filename, "\n")

	a := NewSegments(strings.Split(paths[0], ","))
	b := NewSegments(strings.Split(paths[1], ","))

	var dist, li, lj int
	for _, i := range a {
		lj = 0
		for _, j := range b {
			res, tx, ty := i.Intersect(&j)
			if res {
				tist := i.DirectDist(tx, ty) + j.DirectDist(tx, ty) + li + lj
				if tist == 0 {
					continue
				}
				if dist == 0 || tist < dist {
					dist = tist
				}

				// any other intersection will have longer delay
				break
			}

			lj += j.Length()
		}
		li += i.Length()
	}

	fmt.Printf("Star 1: %12.0d\n", dist)
}
