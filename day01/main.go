package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Mirsait/advent-of-code-2025/common"
)

func main() {
	fmt.Println(common.Hello())
	fmt.Println("Day 1: Secret Entrance")

	var lines, err = common.ReadFileByLines("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Lines count: ", len(lines))

	var code int = puzzle01(50, lines)
	fmt.Println("Puzzle 1. Code: ", code) // Answer: 1007
}

func puzzle01(startedValue int, lines []string) int {
	var zeroCount int = 0
	for _, line := range lines {
		var step int = parseLine(line)
		var nextValue, z = rotate(startedValue, step)
		zeroCount += z
		startedValue = nextValue
	}
	return zeroCount
}

func parseLine(input string) int {
	var sign, value int
	if strings.HasPrefix(input, "L") {
		sign = -1
		var number string = strings.TrimLeft(input, "L")
		value, _ = strconv.Atoi(number)
	} else {
		sign = 1
		var number string = strings.TrimLeft(input, "R")
		value, _ = strconv.Atoi(number)
	}
	return sign * value
}

func rotate(current, step int) (int, int) {
	var next int = current + step
	if next < 0 {
		next = 100 - (abs(next) % 100)
	}
	if next > 100 {
		next = next % 100
	}
	if next == 0 || next == 100 {
		return 0, 1
	}
	return next, 0
}

func abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}
