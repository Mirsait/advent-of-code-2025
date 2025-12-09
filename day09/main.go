package main

import (
	"fmt"
	"strings"

	"github.com/Mirsait/advent-of-code-2025/common"
)

type Point2 struct {
	X, Y int
	C    string
}

func main() {
	common.Clear()
	fmt.Println(common.Hello())

	lines, err := common.ReadFileByLines("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	points := parsePositions(lines)

	answer1 := puzzle1(points)
	fmt.Println("Puzzle I. Answer [4773451098]:", answer1)
}

func puzzle1(points []Point2) int {
	hull := convexHullJarvis(points)
	area := findMaxArea(hull)
	return area
}

func findMaxArea(points []Point2) int {
	n := len(points)
	max := 0
	for j := 0; j < n-2; j++ {
		for k := j + 1; k < n-1; k++ {
			s := getArea(points[j], points[k])
			if s > max {
				max = s
			}
		}
	}
	return max
}

func parsePositions(lines []string) []Point2 {
	result := make([]Point2, len(lines))
	for j, line := range lines {
		str := strings.Split(line, ",")
		x, _ := common.ParseToInt(str[0])
		y, _ := common.ParseToInt(str[1])
		result[j] = Point2{x, y, "red"}
	}
	return result
}

func getArea(p, s Point2) int {
	dx := Abs(p.X-s.X) + 1
	dy := Abs(p.Y-s.Y) + 1
	return dx * dy
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
