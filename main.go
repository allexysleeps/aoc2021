package main

import (
	"fmt"
	"github.com/allexysleeps/aoc2021/day7"
	"time"
)

func main() {
	start := time.Now()
	day7.Part1()
	day7.Part2()
	elapsed := time.Since(start)
	fmt.Printf("time took: %s\n", elapsed)
}
