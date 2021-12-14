package day14

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func getInput() (string, map[string]string) {
	input, _ := ioutil.ReadFile("day14/input.txt")
	lines := strings.Split(string(input), "\n")
	rules := make(map[string]string)
	for i := 2; i < len(lines); i++ {
		rule := strings.Split(lines[i], " -> ")
		rules[rule[0]] = rule[1]
	}
	return lines[0], rules
}

func insert(poly string, rules map[string]string)string {
	var inserted string
	for i := 0; i < len(poly) - 1; i++ {
		pair := poly[i:i+2]
		v, ok := rules[pair]
		if !ok {
			continue
		}
		if i == 0 {
			inserted += string(pair[0]) + v + string(pair[1])
			continue
		}
		inserted += v + string(pair[1])
	}
	return inserted
}

func solve(n int)int {
	polymer, rules := getInput()
	chars := make(map[string]int)
	pairs := make(map[string]int)
	chars[string(polymer[len(polymer) - 1])]++
	for i := 0; i < len(polymer) - 1; i++ {
		pair := polymer[i:i+2]
		pairs[pair]++
		chars[string(polymer[i])]++
	}

	for i := 0; i < n; i++ {
		nextPairs := make(map[string]int)
		for pair, count := range pairs {
			rule, ok := rules[pair]
			if !ok {
				continue
			}
			p1 := string(pair[0]) + rule
			p2 := rule + string(pair[1])
			nextPairs[p1] += count
			nextPairs[p2] += count
			chars[rule] += count
		}
		pairs = nextPairs
	}

	max, min := chars["C"], chars["C"]
	for _, v := range chars {
		if v > max {
			max = v
			continue
		}
		if v < min {
			min = v
		}
	}
	return max - min
}


func Part1() {
	result := solve(10)
	fmt.Printf("day 14, part 1: %d\n", result)
}

func Part2() {
	result := solve(40)
	fmt.Printf("day 14, part 2: %d\n", result)
}