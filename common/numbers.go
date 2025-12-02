package common

import (
	"strconv"
	"strings"
)

func ParseToInt64(value string) (int64, bool) {
	var trim = strings.TrimSpace(value)
	var num, err = strconv.ParseInt(trim, 10, 64)
	if err != nil {
		return 0, false
	}
	return num, true
}

func Next(start, end int64) func() (int64, bool) {
	var j = start
	return func() (int64, bool) {
		if j > end {
			return 0, false
		}
		var val = j
		j++
		return val, true
	}
}
