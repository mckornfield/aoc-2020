#!/bin/bash

dayNumber=$1
if [[ -z "$dayNumber" ]]; then
    echo "Day number is required"
    exit 1
fi
dir_to_make=day${dayNumber}/src/aoc${dayNumber}
mkdir -p ${dir_to_make}
touch ${dir_to_make}/day${dayNumber}.go
touch ${dir_to_make}/day${dayNumber}_test.go
touch ${dir_to_make}/day${dayNumber}_puzzle.txt
