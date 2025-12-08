package main

import (
	"fmt"
	"strings"

	"github.com/Mirsait/advent-of-code-2025/common"
)

func main() {
	common.Clear()
	fmt.Println(common.Hello())

	lines, err := common.ReadFileByLines("input.txt")
	// limit := 1000

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	points := parsePoints(lines)

	// start1 := time.Now()
	// answer1 := puzzle1(points, limit)
	// elapsed1 := time.Since(start1)
	// fmt.Println("Puzzle I. Answer: [90036]", answer1)
	// fmt.Println("Time: ", elapsed1.Milliseconds())

	// start2 := time.Now()
	answer2 := puzzle2(points)
	// elapsed2 := time.Since(start1)
	fmt.Println("Puzzle II. Answer: []", answer2)
	// fmt.Println("Time: ", elapsed2.Milliseconds())
}

func parsePoints(lines []string) []Point3 {
	result := make([]Point3, len(lines))
	for j, line := range lines {
		str := strings.Split(line, ",")
		x, _ := common.ParseToInt(str[0])
		y, _ := common.ParseToInt(str[1])
		z, _ := common.ParseToInt(str[2])
		result[j] = Point3{X: x, Y: y, Z: z}
	}
	return result
}
