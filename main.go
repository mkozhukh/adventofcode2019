package main

import (
	"advent/t14"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	t14.Star1("t14/data.txt")
	elapsed := time.Since(start)
	fmt.Printf("took %s\n\n", elapsed)

	start = time.Now()
	t14.Star2("t14/data.txt")
	elapsed = time.Since(start)
	fmt.Printf("took %s\n\n", elapsed)
}
