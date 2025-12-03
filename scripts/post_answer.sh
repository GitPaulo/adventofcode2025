#!/bin/bash
#
# Post answer to Advent of Code
# Usage: ./post_answer.sh <day> <part> <answer>

set -euo pipefail

readonly AOC_YEAR=2025

main() {
  if [[ -z "${AOC_SESSION:-}" ]]; then
    echo "Error: AOC_SESSION not set" >&2
    echo "Set with: export AOC_SESSION=your_session_cookie" >&2
    exit 1
  fi

  local -r day="${1:-${DAY:-01}}"
  local -r part="${2:-${PART:-1}}"
  local -r answer="${3:-${ANSWER:--1}}"

  echo "Submitting answer for day ${day}, part ${part} with answer ${answer}"

  # Remove leading zeros from day
  local -r day_int=$((10#${day}))
  
  local response
  response=$(curl -sS -b "session=${AOC_SESSION}" \
    -X POST \
    -d "level=${part}&answer=${answer}" \
    "https://adventofcode.com/${AOC_YEAR}/day/${day_int}/answer")

  # Extract and display article content
  echo "${response}" | grep "<article>" | sed 's/.*<article>//; s/<\/article>.*//' | sed 's/<[^>]*>//g' | sed 's/\[.*\]//g' || echo "Failed to extract article"
}

main "$@"
