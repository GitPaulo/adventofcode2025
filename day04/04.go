package main

import (
	"fmt"

	"adventofcode2025/utils"
)

func solve(lines [][]byte) (int, int) {
	rows := len(lines)
	cols := len(lines[0])

	// avoid loops in neighbor checks
	neighbors := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	// Part 1
	p1 := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {

			// skip '.'
			if lines[r][c] == '.' {
				continue
			}

			count := 0
			// 8 neighbor positions
			for _, d := range neighbors {
				nr := r + d[0]
				nc := c + d[1]

				if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
					continue
				}

				// count '@' neighbors
				if lines[nr][nc] == '@' {
					count++
				}

				if count >= 4 {
					break
				}
			}

			if count < 4 {
				p1++
			}
		}
	}

	// Part 2
	p2 := 0
	// inline copy loop
	grid := make([][]byte, rows)
	for i := range lines {
		row := make([]byte, cols)
		copy(row, lines[i])
		grid[i] = row
	}

	for {
		removedThisRound := 0

		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {

				if grid[r][c] != '@' {
					continue
				}

				count := 0
				for _, d := range neighbors {
					nr := r + d[0]
					nc := c + d[1]

					if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
						continue
					}

					if grid[nr][nc] == '@' {
						count++
					}

					if count >= 4 {
						break
					}
				}

				if count < 4 {
					// remove it
					grid[r][c] = '.'
					removedThisRound++
					p2++
				}
			}
		}

		if removedThisRound == 0 {
			break
		}
	}

	return p1, p2
}

func main() {
	lines := utils.ReadGrid("input")
	p1, p2 := solve(lines)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
