package main

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/Mirsait/advent-of-code-2025/common"
)

type Code = int64
type Range struct{ Start, End Code }

func main() {
	common.Clear()
	fmt.Println(common.Hello())

	lines, err := common.ReadFileByLines("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	strRanges, strCodes := splitByEmpty(lines)
	ranges := parseRanges(strRanges)
	codes := parseCodes(strCodes)

	start1 := time.Now()
	code1 := puzzle1(ranges, codes)
	elapsed1 := time.Since(start1)
	fmt.Println("Puzzle I. Code [617]: ", code1)
	fmt.Println("Time (microsecond): ", elapsed1.Microseconds())

	start2 := time.Now()
	code2 := puzzle2(ranges)
	elapsed2 := time.Since(start2)
	fmt.Println("Puzzle II. Code [338258295736104]: ", code2)
	fmt.Println("Time (microsecond): ", elapsed2.Microseconds())
}

func puzzle1(ranges []Range, codes []Code) int {
	slices.Sort(codes)
	slices.SortFunc(ranges, rangeCompare)
	count := 0
	for _, code := range codes {
		for _, rng := range ranges {
			if inRange(code, rng.Start, rng.End) {
				count++
				break
			}
		}
	}
	return count
}

func puzzle2(ranges []Range) Code {
	var count Code = 0
	var numbers []Code
	for _, rng := range ranges {
		numbers = append(numbers, rng.Start, rng.End+1)
	}
	slices.Sort(numbers)

	for j := 0; j < len(numbers)-1; j++ {
		inner := Range{Start: numbers[j], End: numbers[j+1] - 1}
		for _, rng := range ranges {
			if contains(rng, inner) {
				count += inner.End - inner.Start + 1
				break
			}
		}
	}
	return count
}

func contains(a, b Range) bool {
	return a.End >= b.End && a.Start <= b.Start
}

func inRange(x, a, b Code) bool {
	return x >= a && x <= b
}

func rangeCompare(a, b Range) int {
	if a.Start > b.Start {
		return 1
	}
	if a.Start < b.Start {
		return -1
	}
	if a.End > b.End {
		return -1
	}
	if a.End < b.End {
		return 1
	}
	return 0
}

func splitByEmpty(lines []string) ([]string, []string) {
	index := slices.Index(lines, "")
	return lines[:index], lines[index:]
}

func parseCodes(lines []string) []Code {
	result := make([]Code, len(lines))
	for j, line := range lines {
		code, _ := common.ParseToInt64(line)
		result[j] = code
	}
	return result
}

func parseRanges(lines []string) []Range {
	var numbers []Range
	for _, line := range lines {
		var nums = strings.Split(line, "-")
		var start, _ = common.ParseToInt64(nums[0])
		var end, _ = common.ParseToInt64(nums[1])
		numbers = append(numbers, Range{Start: start, End: end})
	}
	return numbers
}
