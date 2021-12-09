package main

import (
	"fmt"
	"github.com/allexysleeps/aoc2021/day9"
	"time"
)

func main() {
	start := time.Now()
	day9.Part1()
	day9.Part2()
	elapsed := time.Since(start)
	fmt.Printf("time took: %s\n", elapsed)
}
