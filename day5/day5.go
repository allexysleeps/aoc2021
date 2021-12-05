package day5

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"unicode"
)

const offset = 1000

func getInput() [][]int {
	input, _ := ioutil.ReadFile("day5/input.txt")
	lines := strings.Split(string(input), "\n")
	points := make([][]int, 0, len(lines))
	for _, l := range lines {
		ps := strings.FieldsFunc(l, func(r rune) bool {
			return !unicode.IsNumber(r)
		})
		pi := make([]int, 0, 4)
		for _, p := range ps {
			num, _ := strconv.Atoi(p)
			pi = append(pi, num)
		}
		points = append(points, pi)
	}
	return points
}

func minMax(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}

func traverse(line []int, hsh map[int]int, incDiagonal bool) {
	x1, y1, x2, y2 := line[0], line[1], line[2], line[3]

	if x1 == x2 {
		min, max := minMax(y1, y2)
		for y := min; y <= max; y++ {
			hsh[x1*offset+y]++
		}
		return
	}

	if y1 == y2 {
		min, max := minMax(x1, x2)
		for x := min; x <= max; x++ {
			hsh[x*offset+y1]++
		}
		return
	}

	if !incDiagonal {
		return
	}

	if math.Abs(float64(x1-x2)) == math.Abs(float64(y1-y2)) {
		minX, maxX := minMax(x1, x2)
		m := (y2 - y1) / (x2 - x1)
		b := y1 - (m * x1)
		for x := minX; x <= maxX; x++ {
			y := m*x + b
			hsh[x*offset+y]++
		}
	}
}

func process(incDiagonal bool) int {
	lines := getInput()
	hsh := make(map[int]int)

	for _, l := range lines {
		traverse(l, hsh, incDiagonal)
	}

	var dPoints int
	for _, v := range hsh {
		if v > 1 {
			dPoints++
		}
	}
	return dPoints
}

func Part1() {
	res := process(false)
	fmt.Printf("day 5, part 1: %d\n", res)
}

func Part2() {
	res := process(true)
	fmt.Printf("day 5, part 2: %d\n", res)
}

// 6397
// 22335
