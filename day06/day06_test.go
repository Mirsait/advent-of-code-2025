package main

import (
	"fmt"
	"testing"

	"github.com/Mirsait/advent-of-code-2025/common"
)

func TestPuzzle1(t *testing.T) {
	lines := []string{
		"123 328  51 64 ",
		" 45 64  387 23 ",
		"  6 98  215 314",
		"*   +   *   +  "}
	var actual int64 = puzzle1(lines)
	var expected int64 = 4277556
	if actual != expected {
		t.Errorf("Puzzle1(lines) returned %v; expected %v", actual, expected)
	}
}

func TestPuzzle2(t *testing.T) {
	lines := []string{
		"123 328  51 64 ",
		" 45 64  387 23 ",
		"  6 98  215 314",
		"*   +   *   +  "}
	var actual int64 = puzzle2(lines)
	var expected int64 = 3263827
	if actual != expected {
		t.Errorf("Puzzle2(lines) returned %v; expected %v", actual, expected)
	}
}

func BenchmarkPuzzle1(b *testing.B) {
	lines, err := common.ReadFileByLines("input.txt")
	if err != nil {
		fmt.Println(err.Error())
	}

	b.ReportAllocs()
	for b.Loop() {
		_ = puzzle1(lines)
	}
}

func BenchmarkPuzzle2(b *testing.B) {
	lines, err := common.ReadFileByLines("input.txt")
	if err != nil {
		fmt.Println(err.Error())
	}

	b.ReportAllocs()
	for b.Loop() {
		_ = puzzle2(lines)
	}
}
