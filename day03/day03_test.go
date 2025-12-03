package main

import (
	"testing"
)

func TestPuzzle1(t *testing.T) {
	lines := [4]string{"987654321111111", "811111111111119", "234234234234278", "818181911112111"}
	expected := 357
	actual := puzzle(lines[:], 2)
	if actual != expected {
		t.Errorf("Puzzle1(lines) returned %v; expected %v", actual, expected)
	}
}

func TestPuzzle2(t *testing.T) {
	lines := [4]string{"987654321111111", "811111111111119", "234234234234278", "818181911112111"}
	expected := 3121910778619
	actual := puzzle(lines[:], 12)
	if actual != expected {
		t.Errorf("Puzzle1(lines) returned %v; expected %v", actual, expected)
	}
}

func TestGetMax_2(t *testing.T) {
	actual := getMax([]int{1, 2, 3, 4}, 2)
	expected := 34
	if actual != expected {
		t.Errorf("getMax(1234) returned %v; expected %v", actual, expected)
	}
}
func TestGetMax_3(t *testing.T) {
	actual := getMax([]int{1, 2, 3, 4}, 3)
	expected := 234
	if actual != expected {
		t.Errorf("getMax(1234) returned %v; expected %v", actual, expected)
	}
}
func TestGetMax_12(t *testing.T) {
	actual := getMax([]int{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8}, 12)
	expected := 434234234278
	if actual != expected {
		t.Errorf("getMax(1234) returned %v; expected %v", actual, expected)
	}
}

// go test -bench=.

func BenchmarkPuzzle2(b *testing.B) {
	lines := [4]string{"987654321111111", "811111111111119", "234234234234278", "818181911112111"}
	b.ReportAllocs()
	for b.Loop() {
		_ = puzzle(lines[:], 2)
	}
}

func BenchmarkPuzzle12(b *testing.B) {
	lines := [4]string{"987654321111111", "811111111111119", "234234234234278", "818181911112111"}
	b.ReportAllocs()
	for b.Loop() {
		_ = puzzle(lines[:], 12)
	}
}
