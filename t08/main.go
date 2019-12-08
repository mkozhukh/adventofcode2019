package t08

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Layer struct {
	data []byte
}

func (l Layer) Count(n byte) int {
	count := 0
	for _, c := range l.data {
		if c == n {
			count++
		}
	}

	return count
}

var WIDTH = 25
var HEIGHT = 6

func Star1(filename string) {
	data, _ := ioutil.ReadFile(filename)

	size := WIDTH * HEIGHT
	lCount := len(data) / size
	layers := make([]Layer, lCount)

	zeroCount := size + 1
	checkSum := 0
	for i := 0; i < lCount; i++ {
		layers[i].data = make([]byte, size)
		copy(layers[i].data, data[i*size:])

		testCount := layers[i].Count(0x30)
		if testCount < zeroCount {
			zeroCount = testCount
			checkSum = layers[i].Count(0x31) * layers[i].Count(0x32)
		}
	}

	fmt.Printf("Star 1: %12.0d\n", checkSum)
}

func Star2(filename string) {
	data, _ := ioutil.ReadFile(filename)

	size := WIDTH * HEIGHT
	lCount := len(data) / size
	canvas := make([]byte, size)
	copy(canvas, data)

	for i := 1; i < lCount; i++ {
		layer := make([]byte, size)
		copy(layer, data[i*size:])

		for i := range canvas {
			if canvas[i] == 0x32 {
				canvas[i] = layer[i]
			}
		}
	}

	fmt.Printf("Star 2: \n")
	for i := 0; i < HEIGHT; i++ {
		fmt.Println(strings.Replace(string(canvas[i*WIDTH:(i+1)*WIDTH]), "0", ".", -1))
	}
}
