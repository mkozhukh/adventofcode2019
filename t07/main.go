package t07

import (
	"advent/common"
	"fmt"
)

type NumMark struct {
	value int
	used  bool
}

type Amplifier struct {
	flash  []int
	vm     *VM
	output int
	input  int
	mode   int
}

func (a *Amplifier) SoftReset() {
	a.vm.SoftReset(a.flash)
}

func runAmps(amps []*Amplifier, loop bool) int {
	p := 0
	for i := range amps {
		amps[i].SoftReset()
		amps[i].input = 0
		amps[i].output = 0
	}

	for {
		out := amps[p].vm.Run()
		if out == -2 && loop {
			break
		}

		n := p + 1
		if n > len(amps)-1 {
			if loop {
				n = 0
			} else {
				break
			}
		}

		amps[n].input = amps[p].output
		p = n
	}

	return amps[len(amps)-1].output
}

var maxValue int

func NewAmplifier(m []int) *Amplifier {
	out := &Amplifier{}

	out.flash = make([]int, len(m))
	copy(out.flash, m)

	out.vm = NewVM(m, 0, func(c int) int {
		if c == 0 {
			return out.mode
		} else {
			return out.input
		}
	}, func(o int) {
		out.output = o
	})

	return out
}

func vmStage(nums []NumMark, amps []*Amplifier, stage int, loop bool) {
	for i := 0; i < 5; i++ {
		if nums[i].used {
			continue
		}

		nums[i].used = true
		amps[stage].mode = nums[i].value

		if stage == 4 {

			nvalue := runAmps(amps, loop)
			if nvalue > maxValue {
				maxValue = nvalue
			}
		} else {
			vmStage(nums, amps, stage+1, loop)
		}
		nums[i].used = false
	}

}

func Star1(filename string) {
	m := common.ReadIntLines(filename, ",")

	nums := []NumMark{{value: 4}, {value: 3}, {value: 2}, {value: 1}, {value: 0}}
	amps := make([]*Amplifier, 5)
	for i := range amps {
		amps[i] = NewAmplifier(m)
	}

	maxValue = 0
	vmStage(nums, amps, 0, false)
	fmt.Printf("\nStar 1: %12.0d\n", maxValue)
}

func Star2(filename string) {
	m := common.ReadIntLines(filename, ",")

	nums := []NumMark{{value: 9}, {value: 8}, {value: 7}, {value: 6}, {value: 5}}
	amps := make([]*Amplifier, 5)
	for i := range amps {
		amps[i] = NewAmplifier(m)
	}

	maxValue = 0
	vmStage(nums, amps, 0, true)
	fmt.Printf("\nStar 2: %12.0d\n", maxValue)
}
