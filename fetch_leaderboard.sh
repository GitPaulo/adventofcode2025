#!/bin/bash
#
# Fetch and parse Advent of Code personal leaderboard to markdown table.

set -euo pipefail

if [[ -z "${AOC_SESSION:-}" ]]; then
  echo "Error: AOC_SESSION environment variable not set" >&2
  exit 1
fi

# Fetch leaderboard
html=$(curl -s -H "Cookie: session=${AOC_SESSION}" \
  "https://adventofcode.com/2025/leaderboard/self")

if [[ "${html}" == *"403 Forbidden"* ]]; then
  echo "Error: 403 Forbidden" >&2
  exit 1
fi
if [[ "${html}" == *"auth/login"* ]]; then
  echo "Error: Authentication failed" >&2
  exit 1
fi

# Output markdown table
echo "| Day | Part 1 | Part 2 |"
echo "|-----|--------|--------|"

# Parse and format times
echo "${html}" | grep -oP '\s+\d+\s+\d+:\d+:\d+\s+\d+:\d+:\d+' | \
while read -r line; do
  read -r day part1 part2 <<< "${line}"
  printf "| %2s | %8s | %8s |\n" "${day}" "${part1}" "${part2}"
done
