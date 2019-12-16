package t15

import (
	"advent/common"
	"fmt"
	"github.com/gdamore/tcell"
)

type Point struct {
	X int
	Y int
}

func (a *Point) Add(b Point) {
	a.X = a.X + b.X
	a.Y = a.Y + b.Y
}

var TOP = Point{X: 0, Y: -1}
var BOTTOM = Point{X: 0, Y: 1}
var RIGHT = Point{X: 1, Y: 0}
var LEFT = Point{X: -1, Y: 0}

var BACK = map[int]int{
	1: 2,
	3: 4,
	4: 3,
	2: 1,
}

var DIR = map[int]Point{
	1: TOP,
	2: BOTTOM,
	3: LEFT,
	4: RIGHT,
}

var tiles = map[int]string{
	0: ".",
	3: "#",
	4: " ",
	5: "X",
}

var backgroundStyle = tcell.StyleDefault.Background(tcell.ColorBlack)
var styles = map[int]tcell.Style{
	0: backgroundStyle,
	3: backgroundStyle.Foreground(tcell.ColorNames["darkgoldenrod"]),
	4: backgroundStyle.Foreground(tcell.ColorNames["silver"]),
	5: backgroundStyle.Foreground(tcell.ColorNames["red"]),
}

func printMap(mp []int, start, end int) {

	fmt.Println()
	for i := start; i < end; i++ {
		for j := start; j < end; j++ {
			fmt.Print(tiles[mp[i*500+j]])
		}
		fmt.Println()
	}
}

func goTo(vm *VM, path []int, p Point) Point {
	for _, d := range path {
		p.Add(DIR[d])
		<-vm.inputRequest
		vm.input <- int64(d)
		<-vm.output
	}

	return p
}

func goBack(vm *VM, path []int) {
	for i := len(path) - 1; i >= 0; i-- {
		<-vm.inputRequest
		vm.input <- int64(BACK[path[i]])
		<-vm.output
	}
}

func testDir(vm *VM, mp []int, p Point, i int) (int, *Point) {
	p.Add(DIR[i])
	n := 500*p.Y + p.X
	if mp[n] > 0 {
		return -2, &p
	}

	<-vm.inputRequest
	vm.input <- int64(i)
	t := <-vm.output

	mp[n] = int(t + 3)

	if t == 1 || t == 2 {
		<-vm.inputRequest
		vm.input <- int64(BACK[i])
		<-vm.output
		return int(t - 1), &p
	}

	return -1, &p
}
func testAround(vm *VM, mp []int, p Point, path []int) (*Point, [][]int) {
	next := make([][]int, 0, 4)
	p = goTo(vm, path, p)

	var found *Point
	for i := 1; i <= 4; i++ {
		a, pt := testDir(vm, mp, p, i)
		if a >= 0 {
			t := make([]int, len(path)+1)
			copy(t, path)
			t[len(path)] = i
			next = append(next, t)

			if a == 1 {
				found = pt
			}
		}
	}

	goBack(vm, path)
	return found, next
}

func Star1(filename string) {
	data := common.ReadInt64Lines(filename, ",")

	vm := NewVM(data, 0)
	go vm.Run()

	mp := make([]int, 500*500)
	next := make([][]int, 0, 10000)
	now := make([][]int, 1, 10000)

	c := 1
	done := false
	start := Point{250, 250}
	mp[500*250+250] = 4

	now[0] = make([]int, 0)
	for {
		next = next[:0]
		for _, v := range now {
			r, sub := testAround(vm, mp, start, v)
			if r != nil {
				done = true
			}
			next = append(next, sub...)
		}
		if done {
			break
		}

		c += 1

		now, next = next, now
	}

	printMap(mp, 225, 280)
	fmt.Printf("Star 1: %12.0d\n", c)
}

func Star2(filename string) {
	data := common.ReadInt64Lines(filename, ",")

	vm := NewVM(data, 0)
	go vm.Run()

	mp := make([]int, 500*500)
	next := make([][]int, 0, 10000)
	now := make([][]int, 1, 10000)

	start := Point{250, 250}
	mp[500*250+250] = 4

	ox := make([]*Point, 0, 20)
	oxNext := make([]*Point, 0, 20)

	now[0] = make([]int, 0)
	for {
		next = next[:0]
		for _, v := range now {
			r, sub := testAround(vm, mp, start, v)
			if r != nil {
				ox = append(ox, r)
			}
			next = append(next, sub...)
		}

		now, next = next, now
		if len(next) == 0 {
			break
		}
	}

	//fill
	oxCount := 0
	for len(ox) > 0 {
		oxNext = oxNext[:0]
		for _, v := range ox {
			for i := 1; i <= 4; i++ {
				t := Point{v.X, v.Y}
				t.Add(DIR[i])
				if mp[500*t.Y+t.X] == 4 {
					mp[500*t.Y+t.X] = 5
					oxNext = append(oxNext, &t)
				}
			}
		}
		oxCount += 1
		ox, oxNext = oxNext, ox
	}

	printMap(mp, 225, 280)
	fmt.Printf("Star 1: %12.0d\n", oxCount)
}
