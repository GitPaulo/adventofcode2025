#!/bin/bash
#
# Fetch and parse Advent of Code personal leaderboard to markdown table.

set -euo pipefail

readonly AOC_YEAR=2025

main() {
  if [[ -z "${AOC_SESSION:-}" ]]; then
    echo "Error: AOC_SESSION not set" >&2
    exit 1
  fi

  local response status_code html
  response=$(curl -s -w "\n%{http_code}" -H "Cookie: session=${AOC_SESSION}" \
    "https://adventofcode.com/${AOC_YEAR}/leaderboard/self")

  html=$(echo "${response}" | head -n -1)
  status_code=$(echo "${response}" | tail -n 1)

  if [[ "${status_code}" == "500" ]]; then
    echo "Error: 500 Internal Server Error" >&2
    exit 1
  fi
  if [[ "${status_code}" == "303" ]]; then
    echo "Error: 303 Redirect - Authentication failed" >&2
    exit 1
  fi
  if [[ "${status_code}" == "403" ]]; then
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

  # Parse and format times (handles complete, incomplete days, and >24h times)
  echo "${html}" | grep -oP '\s+\d+\s+(?:\d+:\d+:\d+|&gt;24h)(?:\s+(?:\d+:\d+:\d+|&gt;24h)|\s+-)' | \
  while read -r line; do
    local day part1 part2
    read -r day part1 part2 <<< "${line}"
    # Convert HTML entity back to >
    part1="${part1//&gt;24h/>24h}"
    part2="${part2//&gt;24h/>24h}"
    printf "| %2s | %8s | %8s |\n" "${day}" "${part1}" "${part2}"
  done
}

main "$@"
