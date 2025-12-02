#!/bin/bash
#
# Parse Advent of Code personal leaderboard HTML to markdown table
# Usage: ./fetch_leaderboard.sh [file.html] or cat file.html | ./fetch_leaderboard.sh

set -euo pipefail

readonly INPUT="${1:--}"

if [[ "${INPUT}" = "-" ]]; then
  HTML=$(cat)
else
  HTML=$(cat "${INPUT}")
fi

if grep -q "403 Forbidden" <<< "${HTML}"; then
  echo "Error: 403 Forbidden. Fetch from host machine." >&2
  exit 1
fi
if grep -q "auth/login" <<< "${HTML}"; then
  echo "Error: Authentication failed." >&2
  exit 1
fi

# Output markdown table
echo "| Day | Part 1 | Part 2 |"
echo "|-----|--------|--------|"

grep -oP '^\s+\d+\s+\d+:\d+:\d+\s+\d+:\d+:\d+' <<< "${HTML}" | while read -r line; do
  read -r day part1 part2 <<< "${line}"
  printf "| %2s | %8s | %8s |\n" "${day}" "${part1}" "${part2}"
done
