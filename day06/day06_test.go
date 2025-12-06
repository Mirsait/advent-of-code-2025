package main

import "testing"

func TestPuzzle1(t *testing.T) {
	lines := []string{
		"123 328  51 64 ",
		" 45 64  387 23 ",
		"  6 98  215 314",
		"*   +   *   +  "}
	var actual int64 = puzzle1(lines)
	var expected int64 = 4277556
	if actual != expected {
		t.Errorf("Puzzle2(ranges, codes) returned %v; expected %v", actual, expected)
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
		t.Errorf("Puzzle2(ranges, codes) returned %v; expected %v", actual, expected)
	}
}
