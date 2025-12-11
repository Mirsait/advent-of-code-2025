package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/Mirsait/advent-of-code-2025/common"
)

type Graph = map[string][]string

func main() {
	common.Clear()
	fmt.Println(common.Hello())

	lines, err := common.ReadFileByLines("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	graph := parseLines(lines)

	count1 := puzzle1(graph)
	fmt.Println("Puzzle I. Count [511]: ", count1)

	count2 := puzzle2(graph)
	fmt.Println("Puzzle II. Count [458618114529380]: ", count2)

}

func puzzle1(graph Graph) int {
	memo := make(map[string]int)
	var dfsCount func(current, target string) int
	dfsCount = func(current, target string) int {
		if val, ok := memo[current]; ok {
			return val
		}
		if current == target {
			return 1
		}
		count := 0
		for _, next := range graph[current] {
			count += dfsCount(next, target)
		}
		memo[current] = count
		return count
	}

	start, target := "you", "out"
	count := dfsCount(start, target)
	return count
}

func parseLines(lines []string) Graph {
	var graph = make(Graph)
	for _, line := range lines {
		kv := strings.Split(line, ":")
		from := kv[0]
		values := strings.Fields(kv[1])
		vals := make([]string, len(values))
		copy(vals, values)
		graph[from] = vals
	}
	return graph
}

// PART II

func puzzle2(graph Graph) int {
	dp0 := make(map[string]int) // no dac, no fft
	dp1 := make(map[string]int) // dac
	dp2 := make(map[string]int) // fft
	dp3 := make(map[string]int) // dac and fft

	start, out, m1, m2 := "svr", "out", "dac", "fft"
	dp0[start] = 1

	order := topologicalSort(graph)

	for _, u := range order {
		// переходы из состояния 0 (не посетили ни dac, ни fft)
		if val, ok := dp0[u]; ok && val > 0 {
			for _, v := range graph[u] {
				switch v {
				case m1:
					dp1[v] += val // теперь посетили только dac
				case m2:
					dp2[v] += val // теперь посетили только fft
				default:
					dp0[v] += val // остаемся в состоянии 0
				}
			}
		}

		// переходы из состояния 1 (посетили только dac)
		if val, ok := dp1[u]; ok && val > 0 {
			for _, v := range graph[u] {
				if v == m2 {
					dp3[v] += val // теперь посетили обе
				} else if v != m1 { // избегаем повторного посещения dac
					dp1[v] += val // остаемся в состоянии 1
				}
			}
		}

		// переходы из состояния 2 (посетили только fft)
		if val, ok := dp2[u]; ok && val > 0 {
			for _, v := range graph[u] {
				if v == m1 {
					dp3[v] += val // теперь посетили обе
				} else if v != m2 { // избегаем повторного посещения fft
					dp2[v] += val // остаемся в состоянии 2
				}
			}
		}

		// переходы из состояния 3 (посетили обе)
		if val, ok := dp3[u]; ok && val > 0 {
			for _, v := range graph[u] {
				if v != m1 && v != m2 { // избегаем повторного посещения dxc и fft
					dp3[v] += val // остаемся в состоянии 3
				}
			}
		}
	}
	return dp3[out]
}

func topologicalSort(graph Graph) []string {
	visited := make(map[string]bool)
	order := []string{}

	var dfs func(string)
	dfs = func(u string) {
		visited[u] = true
		for _, v := range graph[u] {
			if !visited[v] {
				dfs(v)
			}
		}
		order = append(order, u)
	}

	for u := range graph {
		if !visited[u] {
			dfs(u)
		}
	}

	slices.Reverse(order)
	return order
}
