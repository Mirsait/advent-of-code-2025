package main

import (
	"fmt"
	"testing"

	"github.com/Mirsait/advent-of-code-2025/common"
)

func TestTestData1(t *testing.T) {
	lines, err := common.ReadFileByLines("testdata")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	graph := parseLines(lines)
	actual := puzzle1(graph)
	expected := 5
	if actual != expected {
		t.Errorf("Test with TestData failed: actual: %d, expected: %v", actual, expected)
	}
}

func BenchmarkTestData1(b *testing.B) {
	lines, err := common.ReadFileByLines("testdata")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	graph := parseLines(lines)
	b.ReportAllocs()
	for b.Loop() {
		_ = puzzle1(graph)
	}
}
func BenchmarkPuzzle1(b *testing.B) {
	lines, err := common.ReadFileByLines("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	graph := parseLines(lines)
	b.ReportAllocs()
	for b.Loop() {
		_ = puzzle1(graph)
	}
}
