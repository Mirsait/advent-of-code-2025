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

	var code int = puzzle(50, lines, rotate)
	fmt.Println("Puzzle 1. Code: ", code) // Answer: 1007

	var code2 int = puzzle(50, lines, m0x434C49434B)
	fmt.Println("Puzzle 2. Code: ", code2) // Answer: 5820
}

func puzzle(startedValue int, lines []string, fn func(int, int) (int, int)) int {
	var zeroCount int = 0
	for _, line := range lines {
		var step int = parseLine(line)
		var nextValue, z = fn(startedValue, step)
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

const (
	Min = 0
	Max = 100
)

func reset100(value int) int {
	if value == 100 {
		return 0
	}
	return value
}

func m0x434C49434B(current, step int) (int, int) {
	var offset = current + step

	if step == 0 {
		return current, 0
	}
	if step > 0 {
		var next = offset % Max
		var count = offset / Max
		return next, count
	}

	if offset > 0 {
		return offset, 0
	}
	if offset == 0 {
		return 0, 1
	}

	if current == 0 {
		var next = Max - (abs(offset) % Max)
		var count = abs(offset) / Max
		return next, count
	}

	var norm = abs(offset)
	var next = reset100(Max - (norm % Max))
	var count = (norm / Max) + 1
	return next, count
}
