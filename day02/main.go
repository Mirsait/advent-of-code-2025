package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Mirsait/advent-of-code-2025/common"
)

type Range struct {
	Start, End int64
}

func main() {
	fmt.Println(common.Hello())
	fmt.Println("Day 2: Gift Shop")
	var data, err = common.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var splitted = strings.Split(data, ",")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("String Ranges count: ", len(splitted))

	var ranges []Range = convert(splitted)
	fmt.Println("Ranges count: ", len(ranges))

	fmt.Println("Part I:", 52316131093)
	var start1 = time.Now()
	var code1 int64 = puzzleParallel(ranges, isValid2)
	var elapsed1 = time.Since(start1)
	fmt.Printf("Puzzle P1. Code: %d, time: %dms\n", code1, elapsed1.Milliseconds())

	fmt.Println("Part II:", 69564213293)
	var start2 = time.Now()
	var code2 int64 = puzzleParallel(ranges, isValidN)
	var elapsed2 = time.Since(start2)
	fmt.Printf("Puzzle P2. Code: %d, time: %dms\n", code2, elapsed2.Milliseconds())
}

func puzzleParallel(ranges []Range, pred func(int64) bool) int64 {
	var sum int64
	var results = make(chan int64, len(ranges))
	var wg sync.WaitGroup
	wg.Add(len(ranges))

	for _, rng := range ranges {
		go func(r Range) {
			defer wg.Done()
			var localSum = int64(0)
			var next = common.Next(r.Start, r.End)
			for {
				var v, ok = next()
				if !ok {
					break
				}
				if !pred(v) {
					localSum += v
				}
			}
			results <- localSum
		}(rng)
	}

	// a separate goroutine for closing a channel
	go func() {
		wg.Wait()
		close(results)
	}()

	for partial := range results {
		sum += partial
	}
	return sum
}

func convert(str []string) []Range {
	var numbers []Range
	for _, s := range str {
		var nums = strings.Split(s, "-")
		var start, _ = common.ParseToInt64(nums[0])
		var end, _ = common.ParseToInt64(nums[1])
		numbers = append(numbers, Range{Start: start, End: end})
	}
	return numbers
}

func isValid2(value int64) bool {
	var s = strconv.FormatInt(value, 10)
	var n = len(s)
	if n%2 != 0 {
		return true
	}
	var left, right = s[:n/2], s[n/2:]
	return left != right
}

func isValidN(value int64) bool {
	var s string = strconv.FormatInt(value, 10)
	var doubled = s + s
	var cut = doubled[1 : len(doubled)-1]
	return !strings.Contains(cut, s)
}
