package main

import (
	"testing"

	"github.com/Mirsait/advent-of-code-2025/common"
)

func TestXxx(t *testing.T) {
	lines, _ := common.ReadFileByLines("textdata")
	actual := puzzle1(lines)
	expected := 40
	if actual != expected {
		t.Errorf("Puzzle1(data) returned %v; expected %v", actual, expected)
	}
}
