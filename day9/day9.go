package day9

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type matrix struct {
	vals [][]int
	size int
}

func (m *matrix) ptVal(i, j int) int {
	if i < 0 || i >= m.size || j < 0 || j >= m.size {
		return 10
	}
	return m.vals[i][j]
}

func (m *matrix) isLowPoint(i, j int) bool {
	if m.vals[i][j] < m.ptVal(i-1, j) && m.vals[i][j] < m.ptVal(i, j-1) &&
		m.vals[i][j] < m.ptVal(i+1, j) && m.vals[i][j] < m.ptVal(i, j+1) {
		return true
	}
	return false
}

func (m *matrix) basinSize(i, j int, cache map[int]struct{}) int {
	if m.ptVal(i, j) > 8 {
		return 0
	}
	_, ok := cache[(i+1)*10000+(j+1)]
	if ok {
		return 0
	}
	cache[(i+1)*10000+(j+1)] = struct{}{}

	sum := 1
	sum += m.basinSize(i-1, j, cache)
	sum += m.basinSize(i, j-1, cache)
	sum += m.basinSize(i+1, j, cache)
	sum += m.basinSize(i, j+1, cache)
	return sum
}

func createMatrix() matrix {
	vals := getInput()
	m := matrix{
		vals: getInput(),
		size: len(vals),
	}
	return m
}

func getInput() [][]int {
	input, _ := ioutil.ReadFile("day9/input.txt")
	lines := strings.Split(string(input), "\n")
	pts := make([][]int, 0, len(lines))
	for _, l := range lines {
		rowS := strings.Split(l, "")
		rowI := make([]int, 0, len(rowS))
		for _, s := range rowS {
			n, _ := strconv.Atoi(s)
			rowI = append(rowI, n)
		}
		pts = append(pts, rowI)
	}
	return pts
}

func replaceIfBigger(maxes []int, num int) []int {
	if len(maxes) < 3 {
		res := append(maxes, num)
		sort.Ints(res)
		return res
	}
	for i, n := range maxes {
		if num > n {
			if i == 2 {
				res := append(maxes[:i], num)
				sort.Ints(res)
				return res
			}
			res := append(append(maxes[:i], num), maxes[i+1:]...)
			sort.Ints(res)
			return res
		}
	}
	return maxes
}

func Part1() {
	pts := createMatrix()
	var sum int
	for i := 0; i < pts.size; i++ {
		for j := 0; j < pts.size; j++ {
			if pts.isLowPoint(i, j) {
				sum += pts.vals[i][j] + 1
			}
		}
	}
	fmt.Printf("day 9, part 1: %d\n", sum)
}

func Part2() {
	pts := createMatrix()
	maxes := make([]int, 0, 3)
	for i := 0; i < pts.size; i++ {
		for j := 0; j < pts.size; j++ {
			if pts.isLowPoint(i, j) {
				cache := make(map[int]struct{})
				bs := pts.basinSize(i, j, cache)
				maxes = replaceIfBigger(maxes, bs)
			}
		}
	}
	fmt.Printf("day 9, part 2: %d\n", maxes[0]*maxes[1]*maxes[2])
}
