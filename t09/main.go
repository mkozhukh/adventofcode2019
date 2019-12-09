package t09

import (
	"advent/common"
	"fmt"
)

func Star1(filename string) {
	m := common.ReadInt64Lines(filename, ",")

	vm := NewVM(m, 0, func(i int64) int64 {
		return 1
	}, func(i int64) {
		fmt.Println(i)
	})

	for vm.Run() != -2 {
	}

	fmt.Printf("\nStar 1: %12.0d\n", 0)
}

func Star2(filename string) {
	m := common.ReadInt64Lines(filename, ",")

	vm := NewVM(m, 0, func(i int64) int64 {
		return 2
	}, func(i int64) {
		fmt.Println(i)
	})

	for vm.Run() != -2 {
	}

	fmt.Printf("\nStar 1: %12.0d\n", 0)
}
