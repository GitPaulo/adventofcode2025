package main

import (
	"fmt"

	"adventofcode2025/utils"
)

func solve(lines []string) (int, int) {
	// Part 1
	ans := 0
	var firstLargestIdx, secondLargestIdx int

	for _, line := range lines {
		firstLargestIdx = 0
		secondLargestIdx = -1

		for i := range line {
			// Note: ASCII digit ('0'–'9') comparison works correctly here
			// lexicographically because '0' < '1' < ... < '9' (single digit only)
			if line[i] > line[firstLargestIdx] {
				firstLargestIdx = i
			}
		}

		// find next largest to the RIGHT
		for i := firstLargestIdx + 1; i < len(line); i++ {
			if secondLargestIdx == -1 || line[i] > line[secondLargestIdx] {
				secondLargestIdx = i
			}
		}

		// if none found to the right, search from beginning
		if secondLargestIdx == -1 {
			for i := 0; i < firstLargestIdx; i++ {
				if secondLargestIdx == -1 || line[i] > line[secondLargestIdx] {
					secondLargestIdx = i
				}
			}
		}

		// order by position: earlier index is tens, later is ones
		if firstLargestIdx > secondLargestIdx {
			firstLargestIdx, secondLargestIdx = secondLargestIdx, firstLargestIdx
		}
		tens := int(line[firstLargestIdx] - '0')
		ones := int(line[secondLargestIdx] - '0')
		ans += tens*10 + ones
	}

	return ans, 0
}

func main() {
	lines := utils.ReadLines("input")
	p1, p2 := solve(lines)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
