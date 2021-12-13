package main

import (
	"fmt"
	"github.com/allexysleeps/aoc2021/day13"
	"time"
)

func main() {
	start := time.Now()
	day13.Part1()
	day13.Part2()
	elapsed := time.Since(start)
	fmt.Printf("time took: %s\n", elapsed)
}
