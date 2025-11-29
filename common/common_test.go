package common

import "testing"

func TestHello(t *testing.T) {
	actual := Hello()
	expected := "Welcome to Advent Of Code 2025!"
	if actual != expected {
		t.Errorf("Hello() returned %s; expected - %s", actual, expected)
	}
}
