package main

import (
	"fmt"
	"time"

	"github.com/Mirsait/advent-of-code-2025/common"
)

type Indexes = [][2]int

type State = [][]byte

func main() {
	fmt.Println(common.Hello())
	fmt.Println("Day 4: Printing Department")

	lines, _ := common.ReadFileByLines("input.txt")
	state := toState(lines)

	common.Clear()
	start1 := time.Now()
	_, code1 := puzzle1(state)
	elapsed1 := time.Since(start1)
	fmt.Println("Puzzle I. Code [1433]: ", code1)
	fmt.Println("Puzzle I. Time (mcs): ", elapsed1.Microseconds())

	_, code2 := puzzle2(state)
	fmt.Println("Puzzle II. Code [8616]: ", code2)

	// runRender(state)
}

func puzzle1(state State) (State, int) {
	nextState := cloneState(state)
	count := 0
	rows := len(state)
	cols := len(state[0])

	for j := range rows {
		for k := range cols {
			if isRoll(state[j][k]) {
				if checkNeighborhood(state, j, k, rows, cols) {
					count++
					nextState[j][k] = '.'
				}
			}
		}
	}
	return nextState, count
}

func puzzle2(state State) (State, int) {
	sum := 0
	nextState := state
	for {
		var count int
		nextState, count = puzzle1(nextState)
		if count == 0 {
			return nextState, sum
		}
		sum += count
	}
}

func checkNeighborhood(state State, rj int, cj int, rows, cols int) bool {
	indexes := getMoore(rj, cj, rows, cols)
	count := 0
	for _, pp := range indexes {
		pj, pk := pp[0], pp[1]
		if isRoll(state[pj][pk]) {
			count++
		}
	}
	return count < 4
}

func isRoll(p byte) bool {
	return p == '@'
}

func getMoore(i, j, rows, cols int) Indexes {
	var res Indexes
	directions := Indexes{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	for _, d := range directions {
		ni, nj := i+d[0], j+d[1]
		if ni >= 0 && ni < rows && nj >= 0 && nj < cols {
			res = append(res, [2]int{ni, nj})
		}
	}
	return res
}

func toState(lines []string) State {
	var state State
	for i, line := range lines {
		state = append(state, []byte{})
		for _, v := range line {
			state[i] = append(state[i], byte(v))
		}
	}
	return state
}
func cloneState(state State) State {
	next := make(State, len(state))
	for i := range state {
		next[i] = make([]byte, len(state[i]))
		copy(next[i], state[i])
	}
	return next
}
