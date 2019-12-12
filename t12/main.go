package t12

import (
	"advent/common"
	"fmt"
	"strconv"
	"strings"
)

func NewMoonFromString(in string) *Moon {
	// <x=15, y=-2, z=-6>
	parts := strings.Split(in, ",")

	x, _ := strconv.Atoi(strings.Trim(parts[0], "<=> xyz"))
	y, _ := strconv.Atoi(strings.Trim(parts[1], "<=> xyz"))
	z, _ := strconv.Atoi(strings.Trim(parts[2], "<=> xyz"))
	return NewMoon(x, y, z)
}
func NewMoon(x, y, z int) *Moon {
	m := Moon{}

	m.speed = make([]int, 3)
	m.pos = []int{x, y, z}
	return &m
}

type Moon struct {
	pos   []int
	speed []int
}

func (m *Moon) Move() {
	for i := range m.pos {
		m.pos[i] += m.speed[i]
	}
}

func (m *Moon) Energy() int {
	var p, k int
	for i := range m.pos {
		p += common.AbsInt(m.pos[i])
		k += common.AbsInt(m.speed[i])
	}
	return p * k
}

func (m *Moon) Gravity(o *Moon) {
	for i := range m.pos {
		if o.pos[i] > m.pos[i] {
			m.speed[i] += 1
		} else if o.pos[i] < m.pos[i] {
			m.speed[i] -= 1
		}
	}
}

func (m *Moon) Equal(o *Moon) bool {
	for i := range m.pos {
		if o.pos[i] != m.pos[i] {
			return false
		}
	}

	return true
}

func Star1(filename string) {
	data := common.ReadLines(filename, "\n")

	moons := make([]*Moon, 4)
	for i := range data {
		moons[i] = NewMoonFromString(data[i])
	}

	for i := 0; i < 10; i++ {
		for _, m1 := range moons {
			for _, m2 := range moons {
				if m1 != m2 {
					m1.Gravity(m2)
				}
			}
		}

		for _, m := range moons {
			m.Move()
		}
	}

	sum := 0
	for _, m := range moons {
		sum += m.Energy()
	}

	fmt.Printf("Star 1: %12.0d\n", sum)
}

func Star2(filename string) {
	data := common.ReadLines(filename, "\n")

	moons := make([]*Moon, 4)
	base := make([]*Moon, 4)
	for i := range moons {
		moons[i] = NewMoonFromString(data[i])
		base[i] = NewMoonFromString(data[i])
	}

	keys := []int{0, 0, 0}

	c := 0
	for {
		for _, m1 := range moons {
			for _, m2 := range moons {
				if m1 != m2 {
					m1.Gravity(m2)
				}
			}
		}

		for _, m := range moons {
			m.Move()
		}
		c += 1

		for k := 0; k < 3; k++ {
			equal := true
			if keys[k] != 0 {
				continue
			}
			for i := range moons {
				if moons[i].pos[k] != base[i].pos[k] || moons[i].speed[k] != base[i].speed[k] {
					equal = false
					break
				}
			}
			if equal {
				keys[k] = c
			}
		}

		if keys[0] != 0 && keys[1] != 0 && keys[2] != 0 {
			break
		}
	}

	fmt.Printf("Orbits %v\n", keys)
	fmt.Printf("Star 2: %12.0d\n", common.LCM(keys...))
}
