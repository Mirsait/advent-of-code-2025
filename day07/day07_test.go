package main

import (
	"testing"

	"github.com/Mirsait/advent-of-code-2025/common"
)

func TestPuzzle1(t *testing.T) {
	lines, _ := common.ReadFileByLines("testdata")
	state := transformToState(lines)
	actual := puzzle1(state)
	expected := 21
	if actual != expected {
		t.Errorf("Puzzle1(data) returned %v; expected %v", actual, expected)
	}
}

func TestPuzzle2(t *testing.T) {
	lines, _ := common.ReadFileByLines("testdata")
	state := transformToState(lines)
	actual := puzzle2(state)
	expected := 40
	if actual != expected {
		t.Errorf("Puzzle2(data) returned %v; expected %v", actual, expected)
	}
}

func BenchmarkTest1(b *testing.B) {
	lines, _ := common.ReadFileByLines("testdata")
	state := transformToState(lines)
	b.ReportAllocs()
	for b.Loop() {
		_ = puzzle1(state)
	}
}

func BenchmarkTest2(b *testing.B) {
	lines, _ := common.ReadFileByLines("testdata")
	state := transformToState(lines)
	b.ReportAllocs()
	for b.Loop() {
		_ = puzzle2(state)
	}
}
