package main

import (
	"fmt"
	"github.com/allexysleeps/aoc2021/day10"
	"time"
)

func main() {
	start := time.Now()
	day10.Part1()
	day10.Part2()
	elapsed := time.Since(start)
	fmt.Printf("time took: %s\n", elapsed)
}
