package t16

import (
	"advent/common"
	"fmt"
)

var pattern = []int{0, 1, 0, -1}

func fastFinishFFT(data []int, offset int) {
	l := len(data)
	summ := 0
	for i := offset; i < l; i++ {
		summ += data[i]
	}

	for i := offset; i < l; i++ {
		data[i], summ = (summ % 10), summ-data[i]
	}
}
func fastFFT(data []int, offset int) {
	l := len(data)
	for i := offset; i < l; i++ {
		if i > l/2 {
			fastFinishFFT(data, i)
			return
		}
		step := (i + 1) * 4
		summ := 0
		// 0, 1, 2
	plus:
		for j := i; ; j += step {
			for k := 0; k < i+1; k++ {
				pos := j + k
				if pos >= l {
					break plus
				}
				summ += data[pos]
			}
		}
		// 2, 5, 8
	minus:
		for j := 3*i + 2; ; j += step {
			for k := 0; k < i+1; k++ {
				pos := j + k
				if pos >= l {
					break minus
				}
				summ -= data[pos]
			}
		}

		data[i] = common.AbsInt(summ) % 10
	}
}

func Star1(filename string) {
	data := common.ReadIntLines(filename, "")

	for i := 0; i < 100; i++ {
		fastFFT(data, 0)
	}

	fmt.Printf("Star 1: %v\n", common.JoinInts(data[0:8]))
}

func Star2(filename string) {
	base := common.ReadIntLines(filename, "")
	offset := common.BuildInt(base[0:7])
	data := make([]int, 10000*len(base))
	for i := 0; i < 10000; i++ {
		copy(data[i*len(base):(i+1)*len(base)], base)
	}

	for i := 0; i < 100; i++ {
		fastFFT(data, offset)
	}

	fmt.Printf("Star 1: %v\n", common.JoinInts(data[offset:offset+8]))

}
