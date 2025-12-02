package main

import (
	"testing"
)

func TestParseLine_L(t *testing.T) {
	var actual int = parseLine("L45")
	var expected int = -45
	if actual != expected {
		t.Errorf("parseLine() returned %d; expected %d", actual, expected)
	}
}
func TestParseLine_R(t *testing.T) {
	var actual int = parseLine("R89")
	var expected int = 89
	if actual != expected {
		t.Errorf("parseLine() returned %d; expected %d", actual, expected)
	}
}

func TestPuzzle01(t *testing.T) {
	var lines = []string{"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"}
	var expected int = 3
	var actual = puzzle(50, lines, rotate)
	if actual != expected {
		t.Errorf("puzzle01() returned %d; expected %d", actual, expected)
	}
}

func testRotate(current, rot, expected, zeros int, t *testing.T) {
	var av, zv = rotate(current, rot)
	if av != expected || zv != zeros {
		t.Errorf(
			"rotate(%d, %d) returned (%d, %d); expected (%d, %d)",
			current, rot, av, zv, expected, zeros)
	}
}

func TestRotate(t *testing.T) {
	testRotate(50, -68, 82, 0, t)
	testRotate(82, -30, 52, 0, t)
	testRotate(52, 48, 0, 1, t)
	testRotate(0, -5, 95, 0, t)
	testRotate(95, 60, 55, 0, t)
	testRotate(55, -55, 0, 1, t)
	testRotate(0, -1, 99, 0, t)
	testRotate(99, -99, 0, 1, t)
	testRotate(0, 14, 14, 0, t)
	testRotate(14, -82, 32, 0, t)
}

// PART II

func testM0x434C49434B(current, step, expectedValue int, expectedZeros int, t *testing.T) {
	var actualValue, actualZeros = m0x434C49434B(current, step)
	if actualValue != expectedValue {
		t.Errorf(
			"Check value: testM0x434C49434B(%d, %d) returned (%d, _); expected (%d, _)",
			current, step, actualValue, expectedValue)
	}
	if actualZeros != expectedZeros {
		t.Errorf(
			"Check zero: testM0x434C49434B(%d, %d) returned (_, %d); expected (_, %d)",
			current, step, actualZeros, expectedZeros)
	}
}

func TestM0x434C49434B(t *testing.T) {
	testM0x434C49434B(50, 1000, 50, 10, t)
	testM0x434C49434B(50, -1000, 50, 10, t)
	testM0x434C49434B(50, -68, 82, 1, t)
	testM0x434C49434B(82, -30, 52, 0, t)
	testM0x434C49434B(52, 48, 0, 1, t)
	testM0x434C49434B(0, -5, 95, 0, t)
	testM0x434C49434B(95, 60, 55, 1, t)
	testM0x434C49434B(55, -55, 0, 1, t)
	testM0x434C49434B(0, -1, 99, 0, t)
	testM0x434C49434B(99, -99, 0, 1, t)
	testM0x434C49434B(0, 14, 14, 0, t)
	testM0x434C49434B(14, -82, 32, 1, t)
}

func TestPuzzle02(t *testing.T) {
	var lines = []string{"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"}
	var expected int = 6
	var actual = puzzle(50, lines, m0x434C49434B)
	if actual != expected {
		t.Errorf("puzzle02() returned %d; expected %d", actual, expected)
	}
}
