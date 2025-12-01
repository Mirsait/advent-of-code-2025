package common

import "testing"

func TestReadFileByLines(t *testing.T) {
	lines, _ := ReadFileByLines("files.txt")
	actual := 4 // lines count
	expected := len(lines)
	if actual != expected {
		t.Errorf("Hello() returned %d; expected - %d", actual, expected)
	}
}
