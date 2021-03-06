package day3

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"sync"
)

func getInput() ([]int, int) {
	input, err := ioutil.ReadFile("day3/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\n")
	ls := len(lines[0])
	nums := make([]int, 0, len(lines))
	for _, l := range lines {
		n, _ := strconv.ParseInt(l, 2, 64)
		nums = append(nums, int(n))
	}
	return nums, ls
}

func hasBit(num, pos int) bool {
	return num&int(math.Pow(2, float64(pos))) != 0
}

func Part1() {
	nums, ls := getInput()
	sums := make([]int, ls)
	for _, n := range nums {
		for i := range sums {
			if hasBit(n, i) {
				sums[i] += 1
			}
		}
	}

	var gamma int
	var epsilon int

	for i := range sums {
		if sums[i] > len(nums)/2 {
			gamma |= 1 << i
			epsilon |= 0 << i
			continue
		}
		gamma |= 0 << i
		epsilon |= 1 << i
	}
	fmt.Printf("day 3, part 1: %d\n", gamma*epsilon)
}

func filter(nums []int, maj bool, pos int) int {
	size := len(nums)
	if size == 1 {
		return nums[0]
	}
	ones := make([]int, 0, size/2)
	zeros := make([]int, 0, size/2)
	for _, n := range nums {
		if hasBit(n, pos) {
			ones = append(ones, n)
			continue
		}
		zeros = append(zeros, n)
	}
	switch {
	case maj && (len(ones) >= len(zeros)):
		return filter(ones, maj, pos-1)
	case maj:
		return filter(zeros, maj, pos-1)
	case !maj && (len(ones) < len(zeros)):
		return filter(ones, maj, pos-1)
	default:
		return filter(zeros, maj, pos-1)
	}
}

func Part2() {
	nums, ls := getInput()
	var wg sync.WaitGroup
	var ogr, co2 int
	wg.Add(2)
	go func() {
		ogr = filter(nums, true, ls-1)
		wg.Done()
	}()
	go func() {
		co2 = filter(nums, false, ls-1)
		wg.Done()
	}()
	wg.Wait()
	fmt.Printf("day 3, part 2: %d\n", ogr*co2)
}
