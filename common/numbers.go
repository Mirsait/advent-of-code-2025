package common

import (
	"strconv"
	"strings"
)

func ParseToInt64(value string) (int64, bool) {
	trim := strings.TrimSpace(value)
	num, err := strconv.ParseInt(trim, 10, 64)
	if err != nil {
		return 0, false
	}
	return num, true
}

func ParseToInt(value string) (int, bool) {
	trim := strings.TrimSpace(value)
	num, err := strconv.Atoi(trim)
	if err != nil {
		return 0, false
	}
	return num, true
}

func Next(start, end int64) func() (int64, bool) {
	j := start
	return func() (int64, bool) {
		if j > end {
			return 0, false
		}
		val := j
		j++
		return val, true
	}
}

func StringToDigits(str string) []int {
	nums := make([]int, len(str))
	for j, ch := range str {
		nums[j] = int(ch - '0')
	}
	return nums
}

func MaxWithIndex(nums []int) (int, int) {
	max := 0
	index := -1
	for j, v := range nums {
		if v > max {
			max = v
			index = j
		}
	}
	return max, index
}
