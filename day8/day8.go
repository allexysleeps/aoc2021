package day8

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

const outSize = 4
const inpSize = 10

func getInput() [][]string {
	input, _ := ioutil.ReadFile("day8/input.txt")
	lines := strings.Split(string(input), "\n")
	vals := make([][]string, 0, len(lines))

	for _, l := range lines {
		inOut := strings.Split(l, "|")
		vals = append(vals, append(strings.Fields(inOut[0]), strings.Fields(inOut[1])...))
	}
	return vals
}

func Part1() {
	lines := getInput()
	sum := 0
	for _, l := range lines {
		output := l[inpSize:]
		for _, s := range output {
			if len(s) <= 4 || len(s) == 7 {
				sum++
			}
		}
	}
	fmt.Println(sum)
}

func decodeNum(code string, decoder map[string]string) string {
	var numCodes = map[string]string{
		"abcefg":  "0",
		"cf":      "1",
		"acdeg":   "2",
		"acdfg":   "3",
		"bcdf":    "4",
		"abdfg":   "5",
		"abdefg":  "6",
		"acf":     "7",
		"abcdefg": "8",
		"abcdfg":  "9",
	}

	decoded := ""
	for _, r := range code {
		s := string(r)
		decoded += decoder[s]
	}

	return numCodes[sortString(decoded)]
}

func sortString(s string) string {
	slc := strings.Split(s, "")
	sort.Strings(slc)
	return strings.Join(slc, "")
}

func uniqChars(s1, s2 string) string {
	uniq := ""
	for _, r := range s2 {
		if !strings.Contains(s1, string(r)) {
			uniq += string(r)
		}
	}
	return uniq
}

func decodeLine(line []string) map[string]string {
	swaps := make(map[string]string)
	numCodes := make([]string, 10)
	encoded := make([]string, 0)
	for _, s := range line {
		switch len(s) {
		case 2:
			numCodes[1] = s
		case 4:
			numCodes[4] = s
		case 3:
			numCodes[7] = s
		case 7:
			numCodes[8] = s
		default:
			encoded = append(encoded, s)
		}
	}

	swaps["a"] = uniqChars(numCodes[1], numCodes[7])

	for _, enc := range encoded {
		if len(enc) == 6 {
			u := uniqChars(numCodes[4]+swaps["a"], enc)
			if len(u) == 1 {
				swaps["g"] = u
				numCodes[9] = enc
				break
			}
		}
	}

	for _, enc := range encoded {
		if len(enc) == 5 {
			u := uniqChars(numCodes[7]+swaps["g"], enc)
			if len(u) == 1 {
				swaps["d"] = u
				numCodes[3] = enc
				break
			}
		}
	}

	swaps["b"] = uniqChars(numCodes[1]+swaps["d"], numCodes[4])

	zeroCode := strings.Replace(numCodes[8], swaps["d"], "", 1)
	swaps["e"] = uniqChars(numCodes[7]+swaps["b"]+swaps["g"], zeroCode)

	for _, enc := range encoded {
		if len(enc) == 6 {
			u := uniqChars(swaps["a"]+swaps["b"]+swaps["d"]+swaps["g"]+swaps["e"], enc)
			if len(u) == 1 {
				swaps["f"] = u
				numCodes[6] = enc
				break
			}
		}
	}

	swaps["c"] = strings.Replace(numCodes[1], swaps["f"], "", 1)

	decoder := make(map[string]string)

	for k, v := range swaps {
		decoder[v] = k
	}

	return decoder
}

func Part2() {
	lines := getInput()
	sum := 0
	for _, l := range lines {
		decoder := decodeLine(l)
		numS := ""
		for _, code := range l[inpSize:] {
			numS += decodeNum(code, decoder)
		}
		n, _ := strconv.Atoi(numS)
		sum += n
	}
	fmt.Println(sum)
}
