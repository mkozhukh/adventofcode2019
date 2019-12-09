package main

import (
	"advent/t09"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	t09.Star1("t09/data.txt")
	elapsed := time.Since(start)
	fmt.Printf("took %s\n\n", elapsed)

	start = time.Now()
	t09.Star2("t09/data.txt")
	elapsed = time.Since(start)
	fmt.Printf("took %s\n\n", elapsed)
}
