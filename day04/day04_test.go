package main

import (
	"strings"
	"testing"
)

func TestPuzzle3x3(t *testing.T) {
	lines := []string{"@@@", "@@@", "@@@"}
	state := toState(lines)
	_, actual := puzzle1(state)
	expected := 4
	if actual != expected {
		t.Errorf("checkNeighborhood() returned %v; expected %v", actual, expected)
	}
}

func TestPuzzle4x5(t *testing.T) {
	lines := []string{"..@.", "@@.@", "@@@.", "@@..", "@@@."}
	state := toState(lines)
	_, actual := puzzle1(state)
	expected := 5
	if actual != expected {
		t.Errorf("checkNeighborhood() returned %v; expected %v", actual, expected)
	}
}

func TestPuzzle01(t *testing.T) {
	multiLine := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

	lines := strings.Split(multiLine, "\n")
	state := toState(lines)
	_, actual := puzzle1(state)
	expected := 13
	if actual != expected {
		t.Errorf("Puzzle1(lines) returned %v; expected %v", actual, expected)
	}
}

func TestPuzzle02(t *testing.T) {
	multiLine := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

	lines := strings.Split(multiLine, "\n")
	state := toState(lines)
	var _, actual = puzzle2(state)
	expected := 43
	if actual != expected {
		t.Errorf("Puzzle1(lines) returned %v; expected %v", actual, expected)
	}
}

// go test -bench . -benchmem

func BenchmarkPuzzle12(b *testing.B) {
	multiLine := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

	lines := strings.Split(multiLine, "\n")
	state := toState(lines)
	b.ReportAllocs()
	for b.Loop() {
		_, _ = puzzle1(state)
	}
}

func BenchmarkGetMoore(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		_ = getMoore(0, 0, 10, 10)
	}
}
