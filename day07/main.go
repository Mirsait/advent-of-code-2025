package main

import (
	"fmt"
	"time"

	"github.com/Mirsait/advent-of-code-2025/common"
)

func main() {
	common.Clear()
	fmt.Println(common.Hello())
	fmt.Println("Day 7: Laboratories")

	lines, _ := common.ReadFileByLines("input.txt")
	// lines, _ := common.ReadFileByLines("testdata")
	data := transformToState(lines)

	start1 := time.Now()
	code1 := puzzle1(data)
	elapsed1 := time.Since(start1)
	fmt.Println("Puzzle I. Code: [1573]", code1)
	fmt.Println("Time [mcs]:", elapsed1.Microseconds())

	start2 := time.Now()
	code2 := puzzle2(data)
	elapsed2 := time.Since(start2)
	fmt.Println("Puzzle II. Code: [15093663987272]", code2)
	fmt.Println("Time [mcs]:", elapsed2.Microseconds())

}

func puzzle1(state State) int {
	n := len(state)
	count := 0
	for j, line := range state[:n-1] {
		for k, item := range line {
			if item == Source {
				state[j+1][k] = Ray
			}
			if item == Ray && state[j+1][k] == Splitter {
				state[j+1][k-1] = Ray
				state[j+1][k+1] = Ray
				count++
			} else if item == Ray {
				state[j+1][k] = Ray
			}
		}
	}
	return count
}

// PART II

type Point struct{ X, Y int }

func createPoint(x, y int) Point {
	return Point{X: x, Y: y}
}

func puzzle2(state State) int {
	graph := createGraph(state)
	count := countPathsWithMemoize(graph.Nodes[0])
	return count
}

func countPathsWithMemoize(start *Node) int {
	memo := make(map[*Node]int)

	var dfs func(cur *Node) int
	dfs = func(cur *Node) int {
		if cur == nil {
			return 0
		}
		if val, ok := memo[cur]; ok {
			return val
		}
		if len(cur.Edges) == 0 {
			memo[cur] = 1
			return 1
		}
		total := 0
		for _, next := range cur.Edges {
			total += dfs(next)
		}
		memo[cur] = total
		return total
	}
	return dfs(start)
}

func createGraph(state State) Graph {
	var graph Graph
	n := len(state)
	for j, line := range state[:n-1] {
		for k, item := range line {
			if item == Source {
				state[j+1][k] = Ray
				n1 := graph.AddNode(createPoint(j, k))
				n2 := graph.AddNode(createPoint(j+1, k))
				graph.AddEdge(n1, n2)
			}
			if item == Ray && state[j+1][k] == Splitter {
				state[j+1][k-1] = Ray
				state[j+1][k+1] = Ray
				n1 := graph.AddNode(createPoint(j, k))
				n2 := graph.AddNode(createPoint(j+1, k-1))
				n3 := graph.AddNode(createPoint(j+1, k+1))
				graph.AddEdge(n1, n2)
				graph.AddEdge(n1, n3)
			} else if item == Ray {
				state[j+1][k] = Ray
				n1 := graph.AddNode(createPoint(j, k))
				n2 := graph.AddNode(createPoint(j+1, k))
				graph.AddEdge(n1, n2)
			}
		}
	}
	return graph
}
