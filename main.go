package main

import (
	"advent/t13"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	t13.Star1("t13/data.txt")
	elapsed := time.Since(start)
	fmt.Printf("took %s\n\n", elapsed)

	start = time.Now()
	t13.Star2("t13/data.txt")
	elapsed = time.Since(start)
	fmt.Printf("took %s\n\n", elapsed)
}
