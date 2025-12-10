package main

import (
	"strconv"
	"strings"
)

func parseLines(lines []string) []Machine {
	machines := make([]Machine, len(lines))
	for j, line := range lines {
		machines[j] = parseLine(line)
	}
	return machines
}

func parseLine(line string) Machine {
	parts := strings.Split(line, " ")
	machine := Machine{}
	count := 0
	machine.state, count = parseState(parts[0])
	machine.ampers = parseAmpers(parts[len(parts)-1])
	machine.pushes = parsePushes(parts[1:len(parts)-1], count)
	return machine
}

func parsePushes(str []string, length int) []State {
	var states []State
	for _, sv := range str {
		si := strings.TrimSpace(sv)
		si = strings.TrimPrefix(si, "(")
		si = strings.TrimSuffix(si, ")")
		ns := strings.Split(si, ",")
		bools := make([]bool, length)
		for _, n := range ns {
			parsed, _ := strconv.Atoi(n)
			bools[parsed] = true
		}
		states = append(states, BoolsToBitmask(bools))
	}
	return states
}

func parseAmpers(str string) []int {
	s := strings.TrimSpace(str)
	s = strings.TrimPrefix(s, "{")
	s = strings.TrimSuffix(s, "}")
	ns := strings.Split(s, ",")
	nums := make([]int, len(ns))
	for j, n := range ns {
		nums[j], _ = strconv.Atoi(n)
	}
	return nums
}

func parseState(str string) (State, int) {
	s := strings.TrimSpace(str)
	s = strings.TrimPrefix(s, "[")
	s = strings.TrimSuffix(s, "]")
	var bools []bool
	for _, b := range s {
		if b == '#' {
			bools = append(bools, true)
		} else {
			bools = append(bools, false)
		}
	}
	return BoolsToBitmask(bools), len(s)
}

func BoolsToBitmask(bools []bool) State {

	var mask State
	n := len(bools)
	for j, b := range bools {
		if j >= 64 {
			break
		}
		if b {
			mask |= 1 << uint(n-1-j)
		}
	}
	return State(mask)
}
