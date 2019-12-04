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
