package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// reads a file and returns all lines.
func ReadLines(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return lines
}

// converts a string to an integer, panicking on error
func Atoi(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}

// parses comma-separated values from a string
func ParseCSV(s string) []string {
	parts := strings.Split(s, ",")
	result := make([]string, 0, len(parts))
	for _, part := range parts {
		result = append(result, strings.TrimSpace(part))
	}
	return result
}

// parses a range string like "1-10" into two integers
func ParseRange(r string) (int, int) {
	L, R, ok := strings.Cut(r, "-")
	if !ok {
		panic("invalid range")
	}
	return Atoi(L), Atoi(R)
}

// counts the number of digits in an integer
func CountDigits(n int) int {
	return len(fmt.Sprintf("%d", n))
}

// reads a grid of bytes from a file, ensuring all rows have the same length
func ReadGrid(path string) [][]byte {
	lines := ReadLines(path)
	if len(lines) == 0 {
		panic("empty file")
	}

	width := len(lines[0])
	grid := make([][]byte, len(lines))

	for i, line := range lines {
		if len(line) != width {
			panic(fmt.Sprintf("bad grid (line %d has length %d, expected %d)", i, len(line), width))
		}
		grid[i] = []byte(line)
	}

	return grid
}
