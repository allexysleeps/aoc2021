package main

import (
	"fmt"
	"github.com/allexysleeps/aoc2021/day14"
	"time"
)

func main() {
	start := time.Now()
	day14.Part1()
	day14.Part2()
	elapsed := time.Since(start)
	fmt.Printf("time took: %s\n", elapsed)
}
