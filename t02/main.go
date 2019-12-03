package t02

import (
	"advent/common"
	"fmt"
)

func Star1(filename string) {
	m := common.ReadIntLines(filename, ",")

	vm := NewVM(m, 0)
	vm.SetData(1, 12)
	vm.SetData(2, 2)
	out := vm.Run()

	fmt.Printf("Star 1: %12.0d\n", out)
}

func Star2(filename string) {
	m := common.ReadIntLines(filename, ",")
	vm := NewVM(m, 0)

	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			vm.Reset(m, 0)
			vm.SetData(1, i)
			vm.SetData(2, j)
			out := vm.Run()

			if out == 19690720 {
				fmt.Printf("Star 2: %12.0d\n", 100*i+j)
				return
			}
		}
	}
}
