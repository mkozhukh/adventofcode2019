package t14

import (
	"advent/common"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Liquid struct {
	Name    string
	Total   int
	Have    int
	Result  int
	Sources []LiquidSource
}

type LiquidSource struct {
	l string
	c int
}

func spl1(str string) (int, string) {
	strs := strings.Split(strings.TrimSpace(str), " ")
	count, _ := strconv.Atoi(strs[0])

	return count, strs[1]
}
func getLiquids(data []string) map[string]*Liquid {
	ls := make(map[string]*Liquid)

	for _, str := range data {
		parts := strings.Split(str, "=>")
		rc, rn := spl1(parts[1])
		l := Liquid{Name: rn, Result: rc}
		ls[rn] = &l

		sources := strings.Split(parts[0], ",")
		l.Sources = make([]LiquidSource, 0, len(sources))
		for _, sstr := range sources {
			sc, sn := spl1(sstr)
			l.Sources = append(l.Sources, LiquidSource{sn, sc})
		}
	}

	return ls
}

func getOre(ls map[string]*Liquid, name string, c int) int {
	t := ls[name]

	dn := c - t.Have
	if dn < 0 {
		dn = 0
	}
	n := int(math.Ceil(float64(dn) / float64(t.Result)))

	count := 0
	for _, v := range t.Sources {

		if v.l != "ORE" {
			count += getOre(ls, v.l, v.c*n)
		} else {
			count += n * v.c
		}

	}

	t.Have += t.Result*n - c
	t.Total += t.Result * n

	return count
}

func Star1(filename string) {
	data := common.ReadLines(filename, "\n")
	ls := getLiquids(data)

	count := getOre(ls, "FUEL", 1)

	fmt.Printf("Star 1: %12.0d\n", count)
}

func Star2(filename string) {
	data := common.ReadLines(filename, "\n")
	ls := getLiquids(data)

	a := 1000000000000
	d := a / 2
	s := d/2 + d%2

	for s != 1 {
		count := getOre(ls, "FUEL", d)
		if count == a {
			s = 0
		} else {
			if count > a {
				d -= s
			} else {
				d += s
			}
		}
		s = s/2 + s%2
	}

	fmt.Printf("Star 1: %12.0d\n", d)
}
