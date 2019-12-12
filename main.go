package main

import (
	"advent/t12"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	t12.Star1("t12/data.txt")
	elapsed := time.Since(start)
	fmt.Printf("took %s\n\n", elapsed)

	start = time.Now()
	t12.Star2("t12/data.txt")
	elapsed = time.Since(start)
	fmt.Printf("took %s\n\n", elapsed)
}
