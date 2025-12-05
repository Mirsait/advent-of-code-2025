package main

import (
	"fmt"
	"testing"

	"github.com/Mirsait/advent-of-code-2025/common"
)

func TestPuzzle1(t *testing.T) {
	lines := []string{"3-5", "10-14", "16-20", "12-18", "", "1", "5", "8", "11", "17", "32"}
	strRanges, strCodes := splitByEmpty(lines)
	ranges := parseRanges(strRanges)
	codes := parseCodes(strCodes)

	actual := puzzle1(ranges, codes)
	expected := 3
	if actual != expected {
		t.Errorf("Puzzle1(ranges, codes) returned %v; expected %v", actual, expected)
	}
}
func TestPuzzle2(t *testing.T) {
	lines := []string{"3-5", "10-14", "16-20", "12-18", "", "1", "5", "8", "11", "17", "32"}
	strRanges, _ := splitByEmpty(lines)
	ranges := parseRanges(strRanges)

	actual := puzzle2(ranges)
	var expected Code = 14
	if actual != expected {
		t.Errorf("Puzzle2(ranges, codes) returned %v; expected %v", actual, expected)
	}
}

func BenchmarkPuzzle1(b *testing.B) {
	lines, err := common.ReadFileByLines("input.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	strRanges, strCodes := splitByEmpty(lines)
	ranges := parseRanges(strRanges)
	codes := parseCodes(strCodes)

	b.ReportAllocs()
	for b.Loop() {
		_ = puzzle1(ranges, codes)
	}
}

func BenchmarkPuzzle2(b *testing.B) {
	lines, err := common.ReadFileByLines("input.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	strRanges, _ := splitByEmpty(lines)
	ranges := parseRanges(strRanges)

	b.ReportAllocs()
	for b.Loop() {
		_ = puzzle2(ranges)
	}
}
