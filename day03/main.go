package main

import (
	"fmt"
	"time"

	"github.com/Mirsait/advent-of-code-2025/common"
)

func main() {
	fmt.Println(common.Hello())
	fmt.Println("Day 3: Lobby")
	var lines, _ = common.ReadFileByLines("input.txt")

	fmt.Println("Part I:", 17403)
	var start1 = time.Now()
	var code1 int = puzzle(lines, 2)
	var elapsed1 = time.Since(start1)
	fmt.Printf("Puzzle 1. Code: %d, time: %d microseconds\n", code1, elapsed1.Microseconds())

	fmt.Println("Part II:", 173416889848394)
	var start3 = time.Now()
	var code3 int = puzzle(lines, 12)
	var elapsed3 = time.Since(start3)
	fmt.Printf("Puzzle 2. Code: %d, time: %d microseconds\n", code3, elapsed3.Microseconds())
}

func puzzle(lines []string, count int) int {
	var sum = 0
	for _, line := range lines {
		var digits = common.StringToDigits(line)
		sum += getMax(digits, count)
	}
	return sum
}

func getMax(line []int, length int) int {
	var result = 0
	var ln = len(line)
	start := 0
	end := ln - length + 1
	for range length {
		var slc = line[start:end]
		var max, index = common.MaxWithIndex(slc)
		result = result*10 + max
		start += index + 1
		end += 1
	}
	return result
}
