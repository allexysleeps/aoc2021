package day6

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const fullCycle = 7
const maturing = 2

func getInput() []int {
	input, _ := ioutil.ReadFile("day6/input.txt")
	cyclesS := strings.Split(string(input), ",")
	cyclesI := make([]int, 0, len(cyclesS))
	for _, cs := range cyclesS {
		ci, _ := strconv.Atoi(cs)
		cyclesI = append(cyclesI, ci)
	}
	return cyclesI
}

func countOffSpring(cycle, days int, cache map[int]int) int {
	v, ok := cache[cycle*1000+days]
	if ok {
		return v
	}
	fish := 1
	if days <= cycle {
		cache[cycle*1000+days] = fish
		return fish
	}

	remDays := days - cycle
	fish += countOffSpring(fullCycle+maturing, remDays, cache)
	if remDays <= fullCycle {
		cache[cycle*1000+days] = fish
		return fish
	}

	if remDays == fullCycle {
		cache[cycle*1000+days] = fish
		return fish
	}

	rem := remDays % fullCycle

	offsprings := (remDays) / (fullCycle)
	if rem == 0 {
		offsprings--
	}
	for i := 1; i <= offsprings; i++ {
		fish += countOffSpring(fullCycle+maturing, remDays-fullCycle*i, cache)
	}
	cache[cycle*1000+days] = fish
	return fish
}

func process(n int) int {
	cycles := getInput()
	cache := make(map[int]int)
	var fish int
	for _, c := range cycles {
		fish += countOffSpring(c, n, cache)
	}
	return fish
}

func Part1() {
	fmt.Printf("day 4, part 1: %d\n", process(80))
}
func Part2() {
	fmt.Printf("day 4, part 1: %d\n", process(256))
}
