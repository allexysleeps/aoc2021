package day10

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

var corruptionPoints = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var completionPoints = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

var openClose = map[rune]rune{
	'{': '}',
	'[': ']',
	'<': '>',
	'(': ')',
}

func getInput() []string {
	input, _ := ioutil.ReadFile("day10/input.txt")
	return strings.Split(string(input), "\n")
}

func fixLine(s string) (rune, bool, []rune) {
	expected := make([]rune, 0)
	for _, r := range s {
		v, ok := openClose[r]
		if ok {
			expected = append(expected, v)
			continue
		}
		if r != expected[len(expected)-1] {
			return r, false, nil
		}
		expected = expected[:len(expected)-1]
	}

	return 0, true, expected
}

func Part1() {
	lines := getInput()
	sum := 0
	for _, l := range lines {
		r, ok, _ := fixLine(l)
		if !ok {
			sum += corruptionPoints[r]
		}
	}

	fmt.Println(sum)
}

func Part2() {
	lines := getInput()
	scores := make([]int, 0)
	for _, l := range lines {
		_, ok, exp := fixLine(l)
		if !ok || len(exp) == 0 {
			continue
		}
		sum := 0
		for i := range exp {
			sum = sum*5 + completionPoints[exp[len(exp)-1-i]]
		}
		scores = append(scores, sum)
	}
	sort.Ints(scores)
	fmt.Println(scores[len(scores)/2])
}
