package main

import (
	"fmt"
	"strings"
	"time"

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
	// lines := []string{
	// 	"123 328  51 64 ",
	// 	" 45 64  387 23 ",
	// 	"  6 98  215 314",
	// 	"*   +   *   +  "}

	start1 := time.Now()
	code1 := puzzle1(lines)
	elapsed1 := time.Since(start1)
	fmt.Println("Puzzle I. Code [5060053676136]: ", code1)
	fmt.Println("Time (microsecond): ", elapsed1.Microseconds())

	start2 := time.Now()
	code2 := puzzle2(lines)
	elapsed2 := time.Since(start2)
	fmt.Println("Puzzle II. Code [9695042567249]: ", code2)
	fmt.Println("Time (microsecond): ", elapsed2.Microseconds())
}

func puzzle1(lines []string) int64 {
	strNumbers, strOperations := splitData(lines)
	numbers := parseNumberLines(strNumbers)
	rotated := transpose(numbers)
	operations := parseOperations(strOperations)
	var result int64 = 0
	for j, op := range operations {
		result += op(rotated[j])
	}
	return result
}

func puzzle2(lines []string) int64 {
	strNumbers, strOperations := splitData(lines)
	rotated, operations := combineLines(strNumbers, strOperations)
	var result int64 = 0
	for j, op := range operations {
		rv := op(rotated[j])
		result += rv
	}
	return result
}

// 98  12
// 4   345
// +   *
// => [[[5, 24, 13], [8, 94] ] [Mul, Sum]
func combineLines(strNumbers []string, strOperations string) ([][]int64, []Operation) {
	opCount := len(strOperations)
	var (
		operations []Operation
		result     [][]int64
		line       []int64
	)
	for j := opCount - 1; j >= 0; j-- {
		num, ok := getNumber(j, strNumbers)
		if ok {
			line = append(line, num)
			op := strOperations[j]
			if op == '+' || op == '*' {
				operations = append(operations, parseOperation(rune(op)))
				result = append(result, line)
				line = nil
			}
		}
	}
	return result, operations
}

func getNumber(index int, lines []string) (int64, bool) {
	var b strings.Builder
	for _, line := range lines {
		if index < len(line) {
			b.WriteByte(line[index])
		}
	}
	joined := strings.TrimSpace(b.String())
	parsed, ok := common.ParseToInt64(joined)
	return parsed, ok
}

func splitData(lines []string) ([]string, string) {
	return lines[:len(lines)-1], lines[len(lines)-1]
}

func parseNumberLines(str []string) [][]int64 {
	result := make([][]int64, len(str))
	for j, line := range str {
		fields := strings.Fields(line)
		nums := parseNumbers(fields)
		result[j] = nums
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

func transpose(data [][]int64) [][]int64 {
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

type Action = func(x, y int64) int64
type Operation = func(values []int64) int64

func Sum(values []int64) int64 {
	return common.Reduce(func(x, y int64) int64 { return x + y }, 0, values)
}

func Mul(values []int64) int64 {
	return common.Reduce(func(x, y int64) int64 { return x * y }, 1, values)
}

func parseOperations(str string) []Operation {
	result := make([]Operation, 0, len(str))
	for _, ch := range str {
		if ch == '+' || ch == '*' {
			result = append(result, parseOperation(ch))
		}
	}
	return result
}

func parseOperation(op rune) Operation {
	switch op {
	case '+':
		return Sum
	case '*':
		return Mul
	default:
		panic("Undefinded operation")
	}
}
