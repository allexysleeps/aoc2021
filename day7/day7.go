package day7

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func getInput() ([]int, int, int) {
	input, _ := ioutil.ReadFile("day7/input.txt")
	posS := strings.Split(string(input), ",")
	posI := make([]int, 0, len(posS))
	var min, max int
	for i, ps := range posS {
		pi, _ := strconv.Atoi(ps)
		switch {
		case i == 0:
			min, max = pi, pi
		case pi < min:
			min = pi
		case pi > max:
			max = pi
		}
		posI = append(posI, pi)
	}
	return posI, min, max
}

func minFuel(fuelCons func(target, curr int) int) int {
	positions, min, max := getInput()
	minSum := 0
	for i := min; i <= max; i++ {
		sum := 0
		for _, p := range positions {
			sum += fuelCons(i, p)
		}
		if sum < minSum || i == min {
			minSum = sum
		}
	}
	return minSum
}

func Part1() {
	fuel := minFuel(func(target, curr int) int {
		return int(math.Abs(float64(target - curr)))
	})
	fmt.Printf("day 7, part 1: %d\n", fuel)
}

func Part2() {
	fuel := minFuel(func(target, curr int) int {
		steps := int(math.Abs(float64(target - curr)))
		return steps * (steps + 1) / 2
	})
	fmt.Printf("day 7, part 2: %d\n", fuel)
}
