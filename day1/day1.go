package day1

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getNums() []int {
	lines, err := ioutil.ReadFile("day1/input.txt")
	if err != nil {
		panic(err)
	}
	items := strings.Split(string(lines), "\n")
	nums := make([]int, 0, len(items))
	for _, item := range items {
		n, _ := strconv.Atoi(item)
		nums = append(nums, n)
	}
	return nums
}

func Part1() {
	nums := getNums()
	var changes int
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			changes++
		}
	}
	fmt.Printf("day 1, part 1: %d\n", changes)
}

func Part2() {
	nums := getNums()
	var changes int
	for i := 1; i < len(nums)-2; i++ {
		cur := nums[i-1] + nums[i] + nums[i+1]
		nxt := nums[i] + nums[i+1] + nums[i+2]
		if nxt > cur {
			changes++
		}
	}
	fmt.Printf("day 1, part 2: %d\n", changes)
}
