package t13

import (
	"log"
)

const HALT = 99
const ADD = 1
const MUL = 2
const IN = 3
const OUT = 4
const JIT = 5
const JIF = 6
const LS = 7
const EQ = 8
const BASE = 9

type VM struct {
	p    int64
	m    []int64
	base int64

	input        chan int64
	inputRequest chan int64
	output       chan int64
	exit         chan int64
}

func NewVM(state []int64, p int64) *VM {
	m := make([]int64, len(state)+100000)
	copy(m, state)

	return &VM{p: p, m: m, input: make(chan int64, 10), output: make(chan int64), exit: make(chan int64), inputRequest: make(chan int64)}
}

func (v *VM) Param(n int64) int64 {
	return v.m[v.p+n]
}

func (v *VM) DataParam(n int64, mode int64) int64 {
	t := v.Param(n)
	if mode == 1 {
		return t
	}

	if mode == 2 {
		return v.m[v.base+t]
	}

	return v.m[t]
}
func (v *VM) SetDataParam(n, mode, val int64) {
	if mode == 2 {
		v.m[v.base+v.m[v.p+n]] = val
		return
	}
	v.m[v.m[v.p+n]] = val
}

func (v *VM) SetData(n, val int64) {
	v.m[n] = val
}

func (v *VM) GetData(n int64) int64 {
	return v.m[n]
}

func (v *VM) Run() {
	var out int64
	for {
		out = v.Next()
		if out == -2 {
			v.exit <- out
			break
		}
	}
}

func (v *VM) Reset(state []int64, p int64) {
	copy(v.m, state)
	v.p = p
}

func (v *VM) SoftReset(state []int64) {
	copy(v.m, state)
	v.p = 0
}

func (v *VM) Next() int64 {
	code := v.m[v.p]

	op := code % 100
	code = code / 100
	p1 := code % 10
	code = code / 10
	p2 := code % 10
	code = code / 10
	p3 := code % 10

	switch op {
	case HALT:
		return -2
	case ADD:
		v.SetDataParam(3, p3, v.DataParam(1, p1)+v.DataParam(2, p2))
		v.p += 4
	case MUL:
		v.SetDataParam(3, p3, v.DataParam(1, p1)*v.DataParam(2, p2))
		v.p += 4
	case IN:
		v.inputRequest <- 1
		t := <-v.input
		v.SetDataParam(1, p1, t)
		v.p += 2
	case OUT:
		v.output <- v.DataParam(1, p1)
		v.p += 2
	case EQ:
		if v.DataParam(1, p1) == v.DataParam(2, p2) {
			v.SetDataParam(3, p3, 1)
		} else {
			v.SetDataParam(3, p3, 0)
		}
		v.p += 4
	case LS:
		if v.DataParam(1, p1) < v.DataParam(2, p2) {
			v.SetDataParam(3, p3, 1)
		} else {
			v.SetDataParam(3, p3, 0)
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
	case BASE:
		v.base += v.DataParam(1, p1)
		v.p += 2

	default:
		log.Fatalf("Not supported operation: %d", v.m[v.p])
	}

	return 0
}
