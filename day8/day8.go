package day8

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

const inpSize = 10

func getInput() [][]string {
	input, _ := ioutil.ReadFile("day8/input.txt")
	lines := strings.Split(string(input), "\n")
	vals := make([][]string, 0, len(lines))

	for _, l := range lines {
		s := strings.Replace(l, "|", "", 1)
		vals = append(vals, strings.Fields(s))
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
	fmt.Printf("day 8, part 1: %d\n", sum)
}

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

func decodeNum(code string, decoder map[string]string) string {
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

func decodeNumWithSize(comp string, encoded []string, size int) string {
	for _, enc := range encoded {
		if len(enc) == size {
			u := uniqChars(comp, enc)
			if len(u) == 1 {
				return u
			}
		}
	}
	return ""
}

func getDecoder(line []string) map[string]string {
	var one, four, seven, eight string
	encoded := make([]string, 0, len(line)-4)
	for _, s := range line {
		switch len(s) {
		case 2:
			one = s
		case 4:
			four = s
		case 3:
			seven = s
		case 7:
			eight = s
		default:
			encoded = append(encoded, s)
		}
	}

	swp := make(map[string]string)

	swp["a"] = uniqChars(one, seven)
	swp["g"] = decodeNumWithSize(four+swp["a"], encoded, 6)
	swp["d"] = decodeNumWithSize(seven+swp["g"], encoded, 5)
	swp["b"] = uniqChars(one+swp["d"], four)
	zero := strings.Replace(eight, swp["d"], "", 1)
	swp["e"] = uniqChars(seven+swp["b"]+swp["g"], zero)
	swp["f"] = decodeNumWithSize(swp["a"]+swp["b"]+swp["d"]+swp["g"]+swp["e"], encoded, 6)
	swp["c"] = strings.Replace(one, swp["f"], "", 1)

	decoder := make(map[string]string)
	for k, v := range swp {
		decoder[v] = k
	}

	return decoder
}

func Part2() {
	lines := getInput()
	sum := 0
	for _, l := range lines {
		decoder := getDecoder(l)
		strNum := ""
		for _, code := range l[inpSize:] {
			strNum += decodeNum(code, decoder)
		}
		n, _ := strconv.Atoi(strNum)
		sum += n
	}
	fmt.Printf("day 8, part 2: %d\n", sum)
}
