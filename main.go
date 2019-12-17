package main

import (
	"advent/t16"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	t16.Star1("t16/data.txt")
	elapsed := time.Since(start)
	fmt.Printf("took %s\n\n", elapsed)

	start = time.Now()
	t16.Star2("t16/data.txt")
	elapsed = time.Since(start)
	fmt.Printf("took %s\n\n", elapsed)
}
