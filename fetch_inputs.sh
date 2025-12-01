#!/bin/bash
#
# Download Advent of Code 2025 inputs
# Usage: ./fetch_inputs.sh [day_number]

set -euo pipefail

readonly AOC_YEAR=2025

download_day() {
  local -r day="${1}"
  local -r day_padded="$(printf "%02d" "${day}")"
  local -r output_dir="day${day_padded}"
  local -r output_file="${output_dir}/input"

  [[ -d "${output_dir}" ]] || return 0

  local status
  status=$(curl -sS -b "session=${AOC_SESSION}" \
    -w "%{http_code}" \
    -o "${output_file}" \
    "https://adventofcode.com/${AOC_YEAR}/day/${day}/input")

  if [[ "${status}" == "200" ]]; then
    echo "Day ${day}: Downloaded"
  elif [[ "${status}" == "404" ]]; then
    rm -f "${output_file}"
    echo "Day ${day}: Not available yet"
    return 1
  else
    rm -f "${output_file}"
    echo "Day ${day}: Error (HTTP ${status})"
  fi
}

main() {
  if [[ -z "${AOC_SESSION:-}" ]]; then
    echo "Error: AOC_SESSION not set" >&2
    echo "Set with: export AOC_SESSION=your_session_cookie" >&2
    exit 1
  fi

  if [[ $# -eq 0 ]]; then
    for day in {1..25}; do
      download_day "${day}" || break
    done
  else
    download_day "${1}"
  fi
}

main "$@"
