package t02

import "log"

const HALT = 99
const ADD = 1
const MUL = 2

type VM struct {
	p int
	m []int
}

func NewVM(state []int, p int) *VM {
	m := make([]int, len(state))
	copy(m, state)
	return &VM{p: p, m: m}
}

func (v *VM) Param(n int) int {
	return v.m[v.p+n]
}

func (v *VM) DataParam(n int) int {
	return v.m[v.m[v.p+n]]
}
func (v *VM) SetDataParam(n, val int) {
	v.m[v.m[v.p+n]] = val
}

func (v *VM) SetData(n, val int) {
	v.m[n] = val
}

func (v *VM) GetData(n int) int {
	return v.m[n]
}

func (v *VM) Run() int {
	for {
		if !v.Next() {
			break
		}
	}

	return v.GetData(0)
}

func (v *VM) Reset(state []int, p int) {
	copy(v.m, state)
	v.p = p
}

func (v *VM) Next() bool {
	switch v.m[v.p] {
	case HALT:
		return false
	case ADD:
		v.SetDataParam(3, v.DataParam(1)+v.DataParam(2))
		v.p += 4
	case MUL:
		v.SetDataParam(3, v.DataParam(1)*v.DataParam(2))
		v.p += 4
	default:
		log.Fatalf("Not supported operation: %d", v.m[v.p])
	}

	return true
}
