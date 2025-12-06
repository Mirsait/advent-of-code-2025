package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Mirsait/advent-of-code-2025/common"
)

func main() {
	common.Clear()
	fmt.Println(common.Hello())
	lines, err := common.ReadFileByLines("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// lines := []string{"123 328  51 64", " 45 64  387 23", "  6 98  215 314", "*   +   *   +  "}
	code1 := puzzle1(lines)
	fmt.Println("Puzzle I. Code [5060053676136]: ", code1)
}

func puzzle1(lines []string) int64 {
	strNumbers, strOperations := splitData(lines)
	numbers := parseNumberLines(strNumbers)

	rotated := rotate(numbers)
	operations := parseOperations(strOperations)
	var result int64 = 0
	for j, op := range operations {
		rv := op(rotated[j])
		result += rv
	}
	return result

}

func splitData(lines []string) ([]string, string) {
	return lines[:len(lines)-1], lines[len(lines)-1]
}

func parseNumberLines(str []string) [][]int64 {
	result := make([][]int64, len(str))
	re := regexp.MustCompile(`\d+`)
	for j, line := range str {
		matches := re.FindAllString(line, -1)
		result[j] = parseNumbers(matches)
	}
	return result
}

func rotate(data [][]int64) [][]int64 {
	rows := len(data)
	cols := len(data[0])
	result := make([][]int64, cols)
	for cj := range cols {
		result[cj] = make([]int64, rows)
		for rj := range rows {
			result[cj][rj] = data[rj][cj]
		}
	}
	return result
}

func parseNumbers(str []string) []int64 {
	result := make([]int64, len(str))
	for j, v := range str {
		numb, _ := common.ParseToInt64(v)
		result[j] = numb
	}
	return result
}

type Action = func(x, y int64) int64
type Operation = func(values []int64) int64

func Sum(values []int64) int64 {
	return reduce(func(x, y int64) int64 { return x + y }, 0, values)
}

func Mul(values []int64) int64 {
	return reduce(func(x, y int64) int64 { return x * y }, 1, values)
}

func reduce(op Action, initial int64, values []int64) int64 {
	for _, value := range values {
		initial = op(initial, value)
	}
	return initial
}

func parseOperations(str string) []Operation {
	re := regexp.MustCompile(`[+|*]`)
	m := re.FindAllString(str, -1)
	result := make([]Operation, len(m))
	for j, v := range m {
		trimd := strings.Trim(v, " ")
		result[j] = parseOperation(trimd)
	}
	return result
}

func parseOperation(op string) Operation {
	switch op {
	case "+":
		return Sum
	case "*":
		return Mul
	default:
		panic("Undefinded operation")
	}
}
