package main

import (
	"fmt"
	"github.com/allexysleeps/aoc2021/day4"
	"time"
)

func main() {
	start := time.Now()
	day4.Part1()
	day4.Part2()
	elapsed := time.Since(start)
	fmt.Printf("strings time took: %s\n", elapsed)
}
