package main

import (
	"fmt"
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
