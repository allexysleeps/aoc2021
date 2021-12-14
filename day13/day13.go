package day13

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

type fold struct {
	axis  string
	value int
}

func (p point) fold(f fold) point {
	if f.axis == "x" {
		if p.x < f.value {
			return p
		}
		return point{x: 2*f.value - p.x, y: p.y}
	}

	if p.y < f.value {
		return p
	}
	return point{y: 2*f.value - p.y, x: p.x}
}

func getInput() ([]point, []fold) {
	input, _ := ioutil.ReadFile("day13/input.txt")
	lines := strings.Split(string(input), "\n")
	points := make([]point, 0)
	folds := make([]fold, 0)
	parseFolds := false
	for _, l := range lines {
		if l == "" {
			parseFolds = true
			continue
		}
		if !parseFolds {
			nums := strings.Split(l, ",")
			x, _ := strconv.Atoi(nums[0])
			y, _ := strconv.Atoi(nums[1])
			points = append(
				points,
				point{x: x, y: y},
			)
			continue
		}
		chunks := strings.Split(l, " ")
		f := strings.Split(chunks[2], "=")
		num, _ := strconv.Atoi(f[1])
		folds = append(
			folds,
			fold{axis: f[0], value: num},
		)

	}
	return points, folds
}

func Part1() {
	points, folds := getInput()
	foldedPoints := make(map[int]struct{})
	for _, pt := range points {
		p := pt.fold(folds[0])
		idx := p.x*1000 + p.y
		foldedPoints[idx] = struct{}{}
	}
	fmt.Printf("day 13, part 1: %d\n", len(foldedPoints))
}

func Part2() {
	points, folds := getInput()
	var foldCache map[int]struct{}
	for _, f := range folds {
		folded := make([]point, 0, len(points)/2)
		foldCache = make(map[int]struct{})
		for _, p := range points {
			fp := p.fold(f)
			_, ok := foldCache[p.x*10000+p.y]
			if ok {
				continue
			}
			foldCache[fp.x*10000+fp.y] = struct{}{}
			folded = append(folded, fp)
		}
		points = folded
	}

	fmt.Printf("day 13, part 2:\n")
	for i := 0; i < 6; i++ {
		for j := 0; j < 39; j++ {
			idx := j*10000 + i
			_, ok := foldCache[idx]
			if ok {
				fmt.Printf("#")
				continue
			}
			fmt.Printf(".")
		}
		fmt.Printf("\n")
	}
}
