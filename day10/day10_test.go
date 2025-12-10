package main

import (
	"testing"

	"github.com/Mirsait/advent-of-code-2025/common"
)

func TestPush1(t *testing.T) {
	target := BoolsToBitmask([]bool{false, true, true, false})
	current := BoolsToBitmask([]bool{false, false, false, false})

	current = onPush(current, BoolsToBitmask([]bool{false, false, false, true}))
	current = onPush(current, BoolsToBitmask([]bool{false, true, false, true}))
	current = onPush(current, BoolsToBitmask([]bool{false, false, true, false}))
	if current != target {
		t.Errorf("onPush failed: current: %v, target: %v", current, target)
	}
}

func TestPush2(t *testing.T) {
	target := BoolsToBitmask([]bool{false, false, false, true, false})
	current := BoolsToBitmask([]bool{false, false, false, false, false})

	current = onPush(current, BoolsToBitmask([]bool{true, false, false, false, true}))
	current = onPush(current, BoolsToBitmask([]bool{true, true, true, false, false}))
	current = onPush(current, BoolsToBitmask([]bool{false, true, true, true, true}))
	if current != target {
		t.Errorf("onPush failed: current: %v, target: %v", current, target)
	}
}

func TestParsing(t *testing.T) {
	lines, _ := common.ReadFileByLines("testdata")
	machines := parseLines(lines)
	m1 := machines[0]
	m2 := machines[1]
	m3 := machines[2]
	if m1.state != State(uint64(0b0110)) {
		t.Errorf("Machine state is invalid: %04b", m1.state)
	}
	if m2.state != State(uint64(0b00010)) {
		t.Errorf("Machine state is invalid: %04b", m2.state)
	}
	if m3.state != State(uint64(0b011101)) {
		t.Errorf("Machine state is invalid: %04b", m3.state)
	}
}

func TestTestData(t *testing.T) {
	lines, _ := common.ReadFileByLines("testdata")
	machines := parseLines(lines)
	actual := 0
	for _, m := range machines {
		combs := bfsXOR(m.state, m.pushes)
		actual += len(combs)
	}
	expected := 2 + 3 + 2
	if actual != expected {
		t.Errorf("Test with TestData failed: actual: %d, expected: %v", actual, expected)
	}
}
