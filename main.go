package main

import (
	"fmt"
	"github.com/allexysleeps/aoc2021/day8"
	"time"
)

func main() {
	start := time.Now()
	//day8.Part1()
	day8.Part2()
	elapsed := time.Since(start)
	fmt.Printf("time took: %s\n", elapsed)
}
