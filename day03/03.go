package main

import (
	"fmt"

	"adventofcode2025/utils"
)

func solve(lines []string) (int, int64) {
	// Part 1
	ans1 := 0
	for _, line := range lines {
		firstLargestIdx := 0
		secondLargestIdx := -1

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
		ans1 += tens*10 + ones
	}

	// Part 2
	ans2 := int64(0)
	const pick = 12
	for _, line := range lines {
		res := make([]byte, 0, pick)
		remain := pick
		// start window at 0
		start := 0

		// choose remain digits
		for remain > 0 {
			// allowed search window ends so enough digits remain
			end := len(line) - remain

			// find the largest digit in [start, end]
			bestIdx := start
			for i := start; i <= end; i++ {
				if line[i] > line[bestIdx] {
					bestIdx = i
				}
			}

			res = append(res, line[bestIdx])
			remain--
			start = bestIdx + 1
		}

		// Note: max value is 999999999999 (12 nines) ≈ 10^12, needs int64
		value := int64(0)
		for _, c := range res {
			value = value*10 + int64(c-'0')
		}
		ans2 += value
	}

	return ans1, ans2
}

func main() {
	lines := utils.ReadLines("input")
	p1, p2 := solve(lines)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
