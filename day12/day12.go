package day12

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

type visitCache struct {
	cache map[string]struct{}
	twice bool
}

func addEdge(key, value string, graph map[string][]string) {
	v, ok := graph[key]
	if ok {
		graph[key] = append(v, value)
		return
	}
	graph[key] = []string{value}
}

func isUpper(s string) bool {
	return unicode.IsUpper(rune(s[0]))
}

func copyCache(cache map[string]struct{}) map[string]struct{} {
	nc := make(map[string]struct{})
	for k, v := range cache {
		nc[k] = v
	}
	return nc
}

func getInput() map[string][]string {
	input, _ := ioutil.ReadFile("day12/input.txt")
	lines := strings.Split(string(input), "\n")
	c := make(map[string][]string)
	for _, l := range lines {
		cvs := strings.Split(l, "-")
		addEdge(cvs[0], cvs[1], c)
		addEdge(cvs[1], cvs[0], c)
	}
	return c
}

func traverse(pos, dest string, path []string,
	visited visitCache, graph map[string][]string, paths *[]string, twiceSmall bool) {
	if pos == dest {
		p := append(path, pos)
		*paths = append(*paths, strings.Join(p, ""))
		return
	}
	_, ok := visited.cache[pos]
	if ok {
		if pos == "start" || pos == "end" || visited.twice || !twiceSmall {
			return
		}
		visited.twice = true
	}

	if !isUpper(pos) {
		visited.cache[pos] = struct{}{}
	}

	p := append(path, pos)
	for i := 0; i < len(graph[pos]); i++ {
		cache := visitCache{
			cache: copyCache(visited.cache),
			twice: visited.twice,
		}
		traverse(graph[pos][i], dest, p, cache, graph, paths, twiceSmall)
	}
}

func allPaths(graph map[string][]string, start, end string, twiceSmall bool) int {
	paths := make([]string, 0)
	visited := visitCache{
		cache: make(map[string]struct{}),
	}
	var path []string
	traverse(start, end, path, visited, graph, &paths, twiceSmall)
	return len(paths)
}

func Part1() {
	caves := getInput()
	res := allPaths(caves, "start", "end", false)
	fmt.Printf("day 12, part 1: %d\n", res)
}

func Part2() {
	caves := getInput()
	res := allPaths(caves, "start", "end", true)
	fmt.Printf("day 12, part 2: %d\n", res)
}

// 3708
// 93858
