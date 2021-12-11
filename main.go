package main

import (
	"fmt"
	"github.com/allexysleeps/aoc2021/day11"
	"time"
)

func main() {
	start := time.Now()
	day11.Part1()
	day11.Part2()
	elapsed := time.Since(start)
	fmt.Printf("time took: %s\n", elapsed)
}
