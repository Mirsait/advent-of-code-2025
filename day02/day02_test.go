package main

import "testing"

func testingIsValid(value int64, expected bool, t *testing.T) {
	var actual = isValid2(value)
	if actual != expected {
		t.Errorf("InValid(%d) returned %v; expected %v", value, actual, expected)
	}
}

func testingIsValidN(value int64, expected bool, t *testing.T) {
	var actual = isValidN(value)
	if actual != expected {
		t.Errorf("InValid(%d) returned %v; expected %v", value, actual, expected)
	}
}

func TestIsValid(t *testing.T) {
	var notValids = [...]int64{11, 22, 99, 1010, 1188511885, 222222, 446446, 38593859}
	for _, value := range notValids {
		testingIsValid(value, false, t)
	}
	var valids = [...]int64{12, 23, 90, 101, 1011, 1188611885, 222223, 446546, 38594859}
	for _, value := range valids {
		testingIsValid(value, true, t)
	}
}

func TestPuzzle01(t *testing.T) {
	var data = [...]string{"11-22", "95-115", "998-1012", "1188511880-1188511890", "222220-222224",
		"1698522-1698528", "446443-446449", "38593856-38593862"}
	var ranges []Range = convert(data[:])
	var actual int64 = puzzleParallel(ranges, isValid2)
	var expected int64 = 1227775554
	if actual != expected {
		t.Errorf("puzzle01(numbers) returned %v; expected %v", actual, expected)
	}
}

func TestIsValidN(t *testing.T) {
	var notValids = [...]int64{11, 22, 99, 111, 1010, 1188511885, 222222, 446446, 38593859, 565656, 824824824, 2121212121}
	for _, value := range notValids {
		testingIsValidN(value, false, t)
	}
	var valids = [...]int64{12, 23, 90, 101, 1011, 1188611885, 222223, 446546, 38594859}
	for _, value := range valids {
		testingIsValidN(value, true, t)
	}
}

func TestPuzzle02(t *testing.T) {
	var data = [...]string{"11-22", "95-115", "998-1012", "1188511880-1188511890", "222220-222224",
		"1698522-1698528", "446443-446449", "38593856-38593862", "565653-565659", "824824821-824824827", "2121212118-2121212124"}
	var ranges []Range = convert(data[:])
	var actual int64 = puzzleParallel(ranges, isValidN)
	var expected int64 = 4174379265
	if actual != expected {
		t.Errorf("puzzle01(numbers) returned %v; expected %v", actual, expected)
	}
}
