package main

import (
	"advent/t11"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	t11.Star1("t11/data.txt")
	elapsed := time.Since(start)
	fmt.Printf("took %s\n\n", elapsed)

	start = time.Now()
	t11.Star2("t11/data.txt")
	elapsed = time.Since(start)
	fmt.Printf("took %s\n\n", elapsed)
}
