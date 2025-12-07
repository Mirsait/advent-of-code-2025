package main

import "fmt"

type Cell = int

const (
	Empty Cell = iota
	Splitter
	Ray
	Source
)

type State = [][]Cell

func transformToState(lines []string) State {
	result := make([][]Cell, len(lines))
	for j, s := range lines {
		line := make([]Cell, len(s))
		for k, v := range s {
			line[k] = parseCell(v)
		}
		result[j] = line
	}
	return result
}

func parseCell(value rune) Cell {
	switch value {
	case '|':
		return Ray
	case '^':
		return Splitter
	case 'S':
		return Source
	default:
		return Empty
	}
}

func printState(state State) {
	for _, line := range state {
		for _, item := range line {
			switch item {
			case Empty:
				fmt.Print(" ")
			case Source:
				fmt.Print("S")
			case Ray:
				fmt.Print("|")
			case Splitter:
				fmt.Print("^")
			}
		}
		fmt.Println()
	}
}
