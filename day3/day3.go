package day3

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getInputMatrix() ([][]int, int) {
	input, err := ioutil.ReadFile("day3/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\n")
	lineSize := len(lines[0])
	mtrx := make([][]int, 0, len(lines))

	for _, l := range lines {
		row := make([]int, 0, lineSize)
		for i := 0; i < lineSize; i++ {
			bit, _ := strconv.Atoi(string(l[i]))
			row = append(row, bit)
		}
		mtrx = append(mtrx, row)
	}
	return mtrx, lineSize
}

func Part1() {
	mtrx, lineSize := getInputMatrix()
	size := len(mtrx)
	sum := make([]int, lineSize)
	for _, m := range mtrx {
		for j := range sum {
			sum[j] += m[j]
		}
	}
	var gamma, epsilon string
	for i := range sum {
		if sum[i] > size/2 {
			gamma += "1"
			epsilon += "0"
			continue
		}
		gamma += "0"
		epsilon += "1"
	}
	g, _ := strconv.ParseInt(gamma, 2, 64)
	e, _ := strconv.ParseInt(epsilon, 2, 64)
	fmt.Printf("day 3, part 1: %d\n", g*e)
}

func filter(mtrx [][]int, maj bool, pos int) []int {
	size := len(mtrx)
	if size == 1 {
		return mtrx[0]
	}
	ones := make([][]int, 0, size/2)
	zeros := make([][]int, 0, size/2)
	for _, m := range mtrx {
		if m[pos] == 1 {
			ones = append(ones, m)
			continue
		}
		zeros = append(zeros, m)
	}

	if maj {
		if len(ones) >= len(zeros) {
			return filter(ones, maj, pos+1)
		}
		return filter(zeros, maj, pos+1)
	}

	if len(ones) < len(zeros) {
		return filter(ones, maj, pos+1)
	}
	return filter(zeros, maj, pos+1)
}

func Part2() {
	mtrx, _ := getInputMatrix()
	ogr := filter(mtrx, true, 0)
	co2 := filter(mtrx, false, 0)

	var ogrb, co2b string

	for i := range ogr {
		ogrb += fmt.Sprintf("%d", ogr[i])
		co2b += fmt.Sprintf("%d", co2[i])
	}

	ogrd, _ := strconv.ParseInt(ogrb, 2, 64)
	co2d, _ := strconv.ParseInt(co2b, 2, 64)

	fmt.Printf("day 3, part 2: %d\n", ogrd*co2d)
}
