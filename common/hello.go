package common

import "fmt"

func Hello() string {
	return "Welcome to Advent Of Code 2025!"
}

func Clear() {
	fmt.Print("\033[H\033[2J")
}
