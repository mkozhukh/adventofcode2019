package t07

import "log"

const HALT = 99
const ADD = 1
const MUL = 2
const IN = 3
const OUT = 4
const JIT = 5
const JIF = 6
const LS = 7
const EQ = 8

type VM struct {
	p      int
	i      int
	m      []int
	input  inputCallback
	output ouputCallback
}

type inputCallback func(int) int
type ouputCallback func(int)

func NewVM(state []int, p int, input inputCallback, output ouputCallback) *VM {
	m := make([]int, len(state))
	copy(m, state)

	return &VM{p: p, m: m, input: input, output: output, i: 0}
}

func (v *VM) Param(n int) int {
	return v.m[v.p+n]
}

func (v *VM) DataParam(n int, mode int) int {
	t := v.Param(n)
	if mode > 0 {
		return t
	}
	return v.m[t]
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
	out := 0
	for {
		out = v.Next()
		if out < 0 {
			break
		}
	}

	return out
}

func (v *VM) Reset(state []int, p int, input inputCallback, output ouputCallback) {
	copy(v.m, state)
	v.input = input
	v.output = output
	v.p = p
	v.i = 0
}

func (v *VM) SoftReset(state []int) {
	copy(v.m, state)
	v.p = 0
	v.i = 0
}

func (v *VM) Next() int {
	code := v.m[v.p]

	op := code % 100
	code = code / 100
	p1 := code % 10
	code = code / 10
	p2 := code % 10

	switch op {
	case HALT:
		return -2
	case ADD:
		v.SetDataParam(3, v.DataParam(1, p1)+v.DataParam(2, p2))
		v.p += 4
	case MUL:
		v.SetDataParam(3, v.DataParam(1, p1)*v.DataParam(2, p2))
		v.p += 4
	case IN:
		v.SetDataParam(1, v.input(v.i))
		v.i += 1
		v.p += 2
	case OUT:
		v.output(v.DataParam(1, p1))
		v.p += 2
		return -1

	case EQ:
		if v.DataParam(1, p1) == v.DataParam(2, p2) {
			v.SetDataParam(3, 1)
		} else {
			v.SetDataParam(3, 0)
		}
		v.p += 4

	case LS:
		if v.DataParam(1, p1) < v.DataParam(2, p2) {
			v.SetDataParam(3, 1)
		} else {
			v.SetDataParam(3, 0)
		}
		v.p += 4

	case JIT:
		if v.DataParam(1, p1) != 0 {
			v.p = v.DataParam(2, p2)
		} else {
			v.p += 3
		}

	case JIF:
		if v.DataParam(1, p1) == 0 {
			v.p = v.DataParam(2, p2)
		} else {
			v.p += 3
		}

	default:
		log.Fatalf("Not supported operation: %d", v.m[v.p])
	}

	return 0
}
