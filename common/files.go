package common

import (
	"bufio"
	"os"
)

func ReadFileByLines(filename string) ([]string, error) {
	var file, err = os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var scanner = bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		var line = scanner.Text()
		lines = append(lines, line)
	}
	return lines, nil
}
