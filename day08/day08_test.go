package main

import (
	"testing"

	"github.com/Mirsait/advent-of-code-2025/common"
)

func TestTestData1(t *testing.T) {
	lines, _ := common.ReadFileByLines("testdata")
	points := parsePoints(lines)
	actual := puzzle1(points, 10)
	expected := 40
	if actual != expected {
		t.Errorf("Puzzle1(testdata) returned %v; expected %v", actual, expected)
	}
}

func TestTestData2(t *testing.T) {
	lines, _ := common.ReadFileByLines("testdata")
	points := parsePoints(lines)
	actual := puzzle2(points)
	expected := 25272
	if actual != expected {
		t.Errorf("Puzzle2(testdata) returned %v; expected %v", actual, expected)
	}
}
