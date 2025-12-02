package main

import (
	"fmt"
	"math"
	"strings"

	"adventofcode2025/utils"
)

func parseRange(r string) (int, int) {
	L, R, ok := strings.Cut(r, "-")
	if !ok {
		panic("invalid range")
	}
	return utils.Atoi(L), utils.Atoi(R)
}

func countDigits(n int) int {
	return len(fmt.Sprintf("%d", n))
}

func solve(ranges []string) (int, int) {
	sum := 0

	for _, r := range ranges {
		L, R := parseRange(r)
		minDigits := countDigits(L)
		maxDigits := countDigits(R)

		for d := minDigits; d <= maxDigits; d++ {
			if d%2 != 0 {
				continue // skip odd
			}

			k := d / 2
			pow10k := int(math.Pow10(k))
			start := int(math.Pow10(k - 1))

			// e.g
			// A = k digits. Example: k=3 → A = 100..999,
			// so doubled IDs look like: 100100, 101101, …, 999999
			// we just need to iterate A from 10^(k-1) to 10^k - 1 and check bounds
			// if we find AA in [L, R], we add it to the sum
			for A := start; A < pow10k; A++ {
				AA := A*pow10k + A
				if AA < L {
					continue
				}
				if AA > R {
					break
				}
				sum += AA
			}
		}
	}

	return sum, 0
}

func main() {
	lines := utils.ReadLines("input")
	ranges := utils.ParseCSV(lines[0])
	p1, p2 := solve(ranges)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
