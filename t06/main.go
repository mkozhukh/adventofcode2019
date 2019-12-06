package t06

import (
	"advent/common"
	"fmt"
	"strings"
)

type Body struct {
	ID     string
	Parent *Body
	Childs []*Body
}

type Map map[string]*Body

func addOrbit(m Map, a, b string) {
	bodyA := m[a]
	bodyB := m[b]

	if bodyA == nil {
		bodyA = &Body{ID: a, Childs: make([]*Body, 0)}
		m[a] = bodyA
	}
	if bodyB == nil {
		bodyB = &Body{ID: a, Childs: make([]*Body, 0)}
		m[b] = bodyB
	}

	bodyA.Childs = append(bodyA.Childs, bodyB)
	bodyB.Parent = bodyA
}

func countOrbits(b *Body, level int) int {
	count := len(b.Childs) * level

	for _, c := range b.Childs {
		count += countOrbits(c, level+1)
	}
	return count
}

func getOrbits(b *Body) []*Body {
	out := make([]*Body, 0)
	for b.Parent != nil {
		out = append(out, b.Parent)
		b = b.Parent
	}
	return out
}

func Star1(filename string) {
	orbits := common.ReadLines(filename, "\n")
	bodies := make(Map)

	for _, orbit := range orbits {
		parts := strings.Split(orbit, ")")
		addOrbit(bodies, parts[0], parts[1])
	}

	count := countOrbits(bodies["COM"], 1)

	fmt.Printf("Star 1: %12.0d\n", count)
}

func Star2(filename string) {
	orbits := common.ReadLines(filename, "\n")
	bodies := make(Map)

	for _, orbit := range orbits {
		parts := strings.Split(orbit, ")")
		addOrbit(bodies, parts[0], parts[1])
	}

	path1 := getOrbits(bodies["YOU"])
	path2 := getOrbits(bodies["SAN"])

	i1 := len(path1) - 1
	i2 := len(path2) - 1

	for i1 >= 0 && i2 >= 0 {
		if path1[i1] == path2[i2] {
			i1 -= 1
			i2 -= 1
		} else {
			break
		}
	}

	fmt.Printf("Star 1: %12.0d\n", i1+2+i2)
}
