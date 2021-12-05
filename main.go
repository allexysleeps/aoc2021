package main

import (
	"fmt"
	"github.com/allexysleeps/aoc2021/day5"
	"time"
)

func main() {
	start := time.Now()
	day5.Part1()
	day5.Part2()
	elapsed := time.Since(start)
	fmt.Printf("time took: %s\n", elapsed)
}
