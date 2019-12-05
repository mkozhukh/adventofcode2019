package t05

import (
	"advent/common"
	"fmt"
)

func Star1(filename string) {
	m := common.ReadIntLines(filename, ",")

	var out int
	vm := NewVM(m, 0, func() int {
		return 1
	}, func(o int) {
		fmt.Printf("%d, ", o)
		out = o
	})
	vm.Run()

	fmt.Printf("\nStar 1: %12.0d\n", out)
}

func Star2(filename string) {
	m := common.ReadIntLines(filename, ",")

	var out int
	vm := NewVM(m, 0, func() int {
		return 5
	}, func(o int) {
		fmt.Printf("%d, ", o)
		out = o
	})
	vm.Run()

	fmt.Printf("\nStar 2: %12.0d\n", out)
}
