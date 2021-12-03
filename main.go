package main

import (
	"fmt"
	"github.com/allexysleeps/aoc2021/day3"
	"time"
)

func main() {
	start := time.Now()
	day3.Part1()
	day3.Part2()
	elapsed := time.Since(start)
	fmt.Printf("strings time took: %s\n", elapsed)

	start = time.Now()
	day3.Part1bits()
	day3.Part2bits()
	elapsed = time.Since(start)
	fmt.Printf("bitwise time took: %s\n", elapsed)
}
