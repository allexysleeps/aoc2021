package main

import (
	"fmt"
	"github.com/allexysleeps/aoc2021/day6"
	"time"
)

func main() {
	start := time.Now()
	day6.Part1()
	day6.Part2()
	elapsed := time.Since(start)
	fmt.Printf("time took: %s\n", elapsed)
}
