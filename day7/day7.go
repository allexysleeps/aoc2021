package day7

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func getInput() []int {
	input, _ := ioutil.ReadFile("day7/input.txt")
	posS := strings.Split(string(input), ",")
	posI := make([]int, 0, len(posS))
	for _, ps := range posS {
		pi, _ := strconv.Atoi(ps)
		posI = append(posI, pi)
	}
	return posI
}

func minMax(pos []int) (int, int) {
	min, max := pos[0], pos[0]
	for i := 1; i < len(pos); i++ {
		if pos[i] > max {
			max = pos[i]
			continue
		}
		if pos[i] < min {
			min = pos[i]
		}
	}
	return min, max
}

func Part1() {
	positions := getInput()
	min, max := minMax(positions)

	minSum := 0
	for i := min; i <= max; i++ {
		sum := 0
		for _, p := range positions {
			sum += int(math.Abs(float64(i - p)))
		}
		if sum < minSum || i == min {
			minSum = sum
		}
	}

	fmt.Printf("day 7, part 1: %d\n", minSum)
}

func Part2() {
	positions := getInput()
	min, max := minMax(positions)

	minSum := 0
	for i := min; i <= max; i++ {
		sum := 0
		for _, p := range positions {
			steps := int(math.Abs(float64(i - p)))
			sum += steps*(steps+1)/2
		}
		if sum < minSum || i == min {
			minSum = sum
		}
	}

	fmt.Printf("day 7, part 2: %d\n", minSum)
}
