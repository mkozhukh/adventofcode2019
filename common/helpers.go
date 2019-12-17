package common

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func ReadLines(filename, sep string) []string {
	data, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(data), sep)
	fmt.Printf("Input parsed, %d lines\n", len(lines))

	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}

	return lines
}

func JoinInts(ints []int) string {
	str := make([]string, len(ints))
	for i := range ints {
		str[i] = strconv.Itoa(ints[i])
	}

	return strings.Join(str, "")
}

func BuildInt(ints []int) int {
	summ := 0
	base := 1

	for i := len(ints) - 1; i >= 0; i-- {
		summ += ints[i] * base
		base = base * 10
	}

	return summ
}

func ReadIntLines(filename, sep string) []int {
	lines := ReadLines(filename, sep)
	out := make([]int, len(lines))

	var err error
	for i := range lines {
		out[i], err = strconv.Atoi(lines[i])
		if err != nil {
			log.Fatalf("Can't parse string as an int '%s', at line %d", lines[i], i)
		}
	}

	return out
}

func ReadInt64Lines(filename, sep string) []int64 {
	lines := ReadLines(filename, sep)
	out := make([]int64, len(lines))

	var err error
	for i := range lines {
		out[i], err = strconv.ParseInt(lines[i], 10, 64)
		if err != nil {
			log.Fatalf("Can't parse string as an int '%s', at line %d", lines[i], i)
		}
	}

	return out
}

func ReadFloatLines(filename, sep string) []float64 {
	lines := ReadLines(filename, sep)
	out := make([]float64, len(lines))

	var err error
	for i := range lines {
		out[i], err = strconv.ParseFloat(lines[i], 64)
		if err != nil {
			log.Fatalf("Can't parse string as a float '%s', at line %d", lines[i], i)
		}
	}

	return out
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func GCD(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func LCM(v ...int) int {
	x := v[0]
	for i := 1; i < len(v); i++ {
		x = x * v[i] / GCD(x, v[i])
	}

	return x
}
