package common

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func ReadLines(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(data), "\n")
	fmt.Printf("Input parsed, %d lines\n", len(lines))

	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}

	return lines
}

func ReadIntLines(filename string) []int {
	lines := ReadLines(filename)
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


func ReadFloatLines(filename string) []float64 {
	lines := ReadLines(filename)
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