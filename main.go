package main

import (
	"fmt"
	"github.com/allexysleeps/aoc2021/day12"
	"time"
)

func main() {
	start := time.Now()
	day12.Part1()
	day12.Part2()
	elapsed := time.Since(start)
	fmt.Printf("time took: %s\n", elapsed)
}
