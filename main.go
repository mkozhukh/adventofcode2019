package main

import (
	"advent/t15"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	t15.Star1("t15/data.txt")
	elapsed := time.Since(start)
	fmt.Printf("took %s\n\n", elapsed)

	start = time.Now()
	t15.Star2("t15/data.txt")
	elapsed = time.Since(start)
	fmt.Printf("took %s\n\n", elapsed)
}
