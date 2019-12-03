package t01

import (
	"advent/common"
	"fmt"
	"math"
)

func Star1(filename string) float64 {
	mass := common.ReadFloatLines(filename)

	total := 0.0
	for _, v := range mass {
		total += math.Floor(v/3) - 2
	}

	fmt.Printf("Star 1: %12.0f\n", total)
	return total
}

func Star2(filename string) float64 {
	mass := common.ReadFloatLines(filename)

	total := 0.0
	for _, v := range mass {
		fuel := math.Floor(v/3) - 2
		total += fuel
		for {
			fuel = math.Floor(fuel/3) - 2
			if fuel < 0 {
				break
			}

			total += fuel
		}
	}

	fmt.Printf("Star 2: %12.0f\n", total)
	return total
}
