package main

import (
	"fmt"

	"github.com/Mirsait/advent-of-code-2025/common"
)

func main() {
	common.Clear()
	fmt.Println(common.Hello())

	lines, _ := common.ReadFileByLines("testdata")

	// start1 := time.Now()
	answer1 := puzzle1(lines)
	// elapsed1 := time.Since(start1)
	fmt.Println("Puzzle I. Answer: []", answer1)
	// fmt.Println("Time: ", elapsed1.Milliseconds())
}

func puzzle1(lines []string) int {
	return 0
}
