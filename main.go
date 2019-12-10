package main

import (
	"advent/t10"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	t10.Star1("t10/data.txt")
	elapsed := time.Since(start)
	fmt.Printf("took %s\n\n", elapsed)

	start = time.Now()
	t10.Star2("t10/data.txt")
	elapsed = time.Since(start)
	fmt.Printf("took %s\n\n", elapsed)
}
