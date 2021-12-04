package day4

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const size = 5

type board struct {
	rowSums, colSums []int
	nums             [][]int
}

func (b *board) mark(num int) bool {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if b.nums[i][j] == num {
				b.nums[i][j] = -1
				b.rowSums[i]++
				b.colSums[j]++
				if b.rowSums[i] == size || b.colSums[j] == size {
					return true
				}
				return false
			}
		}
	}
	return false
}

func (b *board) sum() int {
	var sum int
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if b.nums[i][j] > 0 {
				sum += b.nums[i][j]
			}
		}
	}
	return sum
}

func createBoard() *board {
	b := board{
		nums:    make([][]int, 0, size),
		rowSums: make([]int, size),
		colSums: make([]int, size),
	}
	return &b
}

func getNumLine(strNums []string) []int {
	nums := make([]int, 0, len(strNums))
	for _, n := range strNums {
		num, _ := strconv.Atoi(n)
		nums = append(nums, num)
	}
	return nums
}

func getInput() ([]int, []*board) {
	input, _ := ioutil.ReadFile("day4/input.txt")
	lines := strings.Split(string(input), "\n")
	nums := getNumLine(strings.Split(lines[0], ","))

	boards := make([]*board, 0)
	k := 0
	b := createBoard()
	for i := 2; i < len(lines); i++ {
		if k == size {
			k = 0
			boards = append(boards, b)
			b = createBoard()
			continue
		}
		n := getNumLine(strings.Fields(lines[i]))
		b.nums = append(b.nums, n)
		k++
	}
	boards = append(boards, b)

	return nums, boards
}

func getWinner(nums []int, boards []*board, lastWinner bool) (*board, int) {
	playingBoards := boards[:]
	for _, num := range nums {
		for i := 0; i < len(playingBoards); i++ {
			done := playingBoards[i].mark(num)
			if done {
				if !lastWinner || len(playingBoards) == 1 {
					return playingBoards[i], num
				}
				playingBoards = append(playingBoards[:i], playingBoards[i+1:]...)
				i--
			}
		}
	}
	return nil, 0
}

func Part1() {
	nums, boards := getInput()
	win, last := getWinner(nums, boards, false)
	fmt.Printf("day 4, part 1: %d\n", win.sum()*last)
}

func Part2() {
	nums, boards := getInput()
	win, last := getWinner(nums, boards, true)
	fmt.Printf("day 4, part 2: %d\n", win.sum()*last)
}
