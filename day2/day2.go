package day2

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getLines() []string {
	lines, err := ioutil.ReadFile("day2/input.txt")
	if err != nil {
		panic(err)
	}
	items := strings.Split(string(lines), "\n")
	return items
}

func Part1() {
	var depth int
	var distance int
	lines := getLines()
	for _, l := range lines {
		tuple := strings.Split(l, " ")
		x, _ := strconv.Atoi(tuple[1])
		switch tuple[0] {
		case "forward":
			distance += x
		case "up":
			depth -= x
		case "down":
			depth += x
		}
	}
	fmt.Printf("day 2, part 2: %d\n", depth*distance)
}

func Part2() {
	var depth int
	var aim int
	var distance int
	lines := getLines()
	for _, l := range lines {
		tuple := strings.Split(l, " ")
		x, _ := strconv.Atoi(tuple[1])
		switch tuple[0] {
		case "forward":
			if aim != 0 {
				depth += x * aim
			}
			distance += x
		case "up":
			aim -= x
		case "down":
			aim += x
		}
	}
	fmt.Printf("day 2, part 2: %d\n", depth*distance)
}
