package day11

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getInput() [][]int {
	input, _ := ioutil.ReadFile("day11/input.txt")
	lines := strings.Split(string(input), "\n")
	mtrx := make([][]int, 0, len(lines))
	for _, l := range lines {
		numsS := strings.Split(l, "")
		numsI := make([]int, 0, len(numsS))
		for _, ns := range numsS {
			ni, _ := strconv.Atoi(ns)
			numsI = append(numsI, ni)
		}
		mtrx = append(mtrx, numsI)
	}
	return mtrx
}

func flash(i, j int, mtrx [][]int, flashed map[int]struct{}) int {
	if j < 0 || i < 0 || i == len(mtrx) || j == len(mtrx) {
		return 0
	}
	idx := i*100 + j
	_, ok := flashed[idx]
	if ok {
		return 0
	}
	if mtrx[i][j] < 9 {
		mtrx[i][j]++
		return 0
	}
	sum := 1
	mtrx[i][j] = 0
	flashed[idx] = struct{}{}
	sum += flash(i, j+1, mtrx, flashed)
	sum += flash(i, j-1, mtrx, flashed)
	sum += flash(i+1, j, mtrx, flashed)
	sum += flash(i+1, j+1, mtrx, flashed)
	sum += flash(i+1, j-1, mtrx, flashed)
	sum += flash(i-1, j, mtrx, flashed)
	sum += flash(i-1, j+1, mtrx, flashed)
	sum += flash(i-1, j-1, mtrx, flashed)
	return sum
}

func countStep(mtrx [][]int, flashed map[int]struct{}) int {
	sum := 0
	for i := 0; i < len(mtrx); i++ {
		for j := 0; j < len(mtrx); j++ {
			sum += flash(i, j, mtrx, flashed)
		}
	}
	return sum
}

func Part1() {
	mtrx := getInput()
	sum := 0
	for i := 0; i < 100; i++ {
		flashed := make(map[int]struct{})
		sum += countStep(mtrx, flashed)
	}
	fmt.Println(sum)
}

func Part2() {
	mtrx := getInput()
	step := 0
	for i := 0; ; i++ {
		flashed := make(map[int]struct{})
		stepFlashes := countStep(mtrx, flashed)
		if stepFlashes == 100 {
			step = i + 1
			break
		}
	}
	fmt.Println(step)
}
