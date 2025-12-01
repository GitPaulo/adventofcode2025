package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readInput(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func solve(lines []string) (int, int) {
	part1, part2 := 0, 0

	// Part 1
	dial := 50
	password := 0
	for _, line := range lines {
		direction := line[0] // L or R
		value, _ := strconv.Atoi(line[1:]) // number of steps

		if direction == 'L' {
			dial -= value
		} else if direction == 'R' {
			dial += value
		}

		if dial % 100 == 0 {
			password++
		}
	}
	part1 = password

	// Part 2
	dial = 50
	password = 0
	for _, line := range lines {
		direction := line[0] // L or R
		value, _ := strconv.Atoi(line[1:]) // number of steps
		remainder := value % 100

		// count full rotations
		password += value / 100

		if direction == 'L' {
			// count partial rotation if we cross/land on 0
			// skip if dial is already at 0
			if dial > 0 && dial <= remainder {
				password++
			}
			dial -= value
		} else if direction == 'R' {
			// count partial rotation if we wrap past 99 to 0
			if dial + remainder > 99 {
				password++
			}
			dial += value
		}

		// wrap to 0-99
		dial = ((dial % 100) + 100) % 100
	}
	part2 = password

	return part1, part2
}

func main() {
	lines, err := readInput("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	p1, p2 := solve(lines)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
