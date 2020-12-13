#!/bin/bash

dayNumber=$1
if [[ -z "$dayNumber" ]]; then
    echo "Day number is required"
    exit 1
fi
mkdir -p day${dayNumber}/src/aoc
touch day${dayNumber}/src/aoc/day${dayNumber}.go
touch day${dayNumber}/src/aoc/day${dayNumber}_test.go
touch day${dayNumber}/src/aoc/day${dayNumber}_puzzle.txt
