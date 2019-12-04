package t04

import (
	"fmt"
	"math"
)

type callback func([]int)

func explode(n int) []int {
	size := int(math.Ceil(math.Log10(float64(n))))
	out := make([]int, size)

	for i := 0; i < size; i++ {
		out[size-i-1] = n % 10
		n = n / 10
	}

	return out
}

func each(a, b []int, pos int, top bool, cb callback) {
	if pos > 0 && a[pos-1] > a[pos] {
		for i := pos; i < len(a); i++ {
			a[i] = a[pos-1]
		}
	}

	notLast := pos < len(a)-1

	for {
		limit := false
		if top && a[pos] >= b[pos] {
			return
		}
		if top && a[pos] == b[pos] {
			limit = true
		}

		if notLast {
			each(a, b, pos+1, limit, cb)
		} else {
			cb(a)
		}

		if a[pos] == 9 || limit {
			a[pos] = 0
			return
		} else {
			a[pos] = a[pos] + 1
		}
	}
}

func Star1(from, to int) {
	s := explode(from)
	e := explode(to)

	count := 0
	each(s, e, 0, true, func(a []int) {
		fail := true
		for i := range a {
			if i > 0 && a[i-1] == a[i] {
				fail = false
			}
		}

		if !fail {
			count++
		}
	})

	fmt.Printf("Star 1: %12.0d\n", count)
}

func Star2(from, to int) {
	s := explode(from)
	e := explode(to)

	count := 0
	each(s, e, 0, true, func(a []int) {
		fail := true
		for i := range a {
			if i > 0 && a[i-1] == a[i] && (i == 1 || a[i-2] != a[i]) && (i == len(a)-1 || a[i] != a[i+1]) {
				fail = false
			}
		}

		if !fail {
			count++
		}
	})

	fmt.Printf("Star 1: %12.0d\n", count)
}
