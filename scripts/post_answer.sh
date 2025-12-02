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

  if [[ $# -ne 3 ]]; then
    echo "Usage: $0 <day> <part> <answer>" >&2
    exit 1
  fi

  local -r day="${1}"
  local -r part="${2}"
  local -r answer="${3}"

  echo "Submitting answer for Day ${day}, Part ${part}: ${answer}"

  local response
  response=$(curl -sS -b "session=${AOC_SESSION}" \
    -X POST \
    -d "level=${part}&answer=${answer}" \
    "https://adventofcode.com/${AOC_YEAR}/day/${day}/answer")

  # Parse response for feedback
  if [[ "${response}" == *"That's the right answer"* ]]; then
    echo "✓ Correct!"
  elif [[ "${response}" == *"That's not the right answer"* ]]; then
    echo "✗ Incorrect"
    if [[ "${response}" == *"too low"* ]]; then
      echo "  (Answer is too low)"
    elif [[ "${response}" == *"too high"* ]]; then
      echo "  (Answer is too high)"
    fi
  elif [[ "${response}" == *"You gave an answer too recently"* ]]; then
    echo "⏱ Rate limited - wait before submitting again"
  elif [[ "${response}" == *"Did you already complete it"* ]] || [[ "${response}" == *"don't seem to be solving the right level"* ]]; then
    echo "✓ Already completed"
  elif [[ "${response}" == *"You don't seem to be solving the right level"* ]]; then
    echo "✓ Already completed this part"
  else
    # Extract and display article content
    echo "${response}" | sed -n 's/.*<article>\(.*\)<\/article>.*/\1/p' | sed 's/<[^>]*>//g'
  fi
}

main "$@"
