package main

import (
	"fmt"

	"github.com/Mirsait/advent-of-code-2025/common"
)

type State uint64
type Machine struct {
	state  State
	pushes []State
	ampers []int
}

func main() {
	common.Clear()
	fmt.Println(common.Hello())

	lines, err := common.ReadFileByLines("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	machines := parseLines(lines)

	count1 := puzzle1(machines)
	fmt.Println("Puzzle I. Count [494]: ", count1)

}

func puzzle1(machines []Machine) int {
	count := 0
	for _, m := range machines {
		combs := bfsXOR(m.state, m.pushes)
		count += len(combs)
	}
	return count
}

type Sample struct {
	state State
	comb  []State
}

func onPush(state State, push State) State {
	return state ^ push
}

func bfsXOR(target State, numbers []State) []State {
	queue := []Sample{{state: 0, comb: []State{}}}

	best := make(map[State]int)
	best[0] = 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.state == target {
			return current.comb
		}

		for _, push := range numbers {
			newVal := onPush(current.state, push)
			newLen := len(current.comb) + 1

			// если ещё не было или есть короче путь
			if prevLen, ok := best[newVal]; !ok || newLen < prevLen {
				best[newVal] = newLen
				newComb := append([]State{}, current.comb...)
				newComb = append(newComb, push)
				queue = append(queue, Sample{state: newVal, comb: newComb})
			}
		}
	}
	return nil
}
