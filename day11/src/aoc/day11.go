package aoc

import (
	"fmt"
	"strings"
)

func ParseInput(input string) Grid {
	lines := strings.Split(input, "\n")
	grid := Grid{}
	for y, line := range lines {
		grid = append(grid, make([]string, len(line)))
		for x, elem := range line {
			grid[y][x] = string(elem)
		}
	}
	return grid
}

func IterateSeats(grid Grid) Grid {
	newGrid := Grid{}
	for y := 0; y < grid.yMax(); y++ {
		newGrid = append(newGrid, make([]string, grid.xMax()))
		for x := 0; x < grid.xMax(); x++ {
			adjacentSeatCount := getAdjacentOccupiedSeats(grid, x, y)
			currentSeat := grid.get(x, y)
			if currentSeat == "#" && adjacentSeatCount > 3 {
				newGrid[y][x] = "L"
			} else if currentSeat == "L" && adjacentSeatCount == 0 {
				newGrid[y][x] = "#"
			} else {
				newGrid[y][x] = grid.get(x, y)
			}
		}
	}
	return newGrid
}

func IterateUntilRepeat(input string) int {
	grid := ParseInput(input)
	gridsNotEqual := true
	for gridsNotEqual {
		newGrid := IterateSeats(grid)
		if newGrid.equal(grid) {
			gridsNotEqual = false
		}
		grid = newGrid
	}
	// grid.Print()
	return CountOccupiedSeats(grid)
}

func CountOccupiedSeats(grid Grid) int {
	count := 0
	for y := 0; y < grid.yMax(); y++ {
		for x := 0; x < grid.xMax(); x++ {
			if grid.get(x, y) == "#" {
				count++
			}
		}
	}
	return count
}

// This function is dumb but :shrug:
func getAdjacentOccupiedSeats(grid Grid, x int, y int) int {
	counter := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if grid.get(x+i, y+j) == "#" {
				counter++
			}
		}
	}
	// s := fmt.Sprintf("%d,%d %s, count: %d", x, y, grid.get(x, y), counter)
	// fmt.Println(s)
	return counter
}

type Grid [][]string

func (grid Grid) get(x, y int) string {
	// If out of range, pretend it's the floor :)
	if y > len(grid)-1 || y < 0 || x > len(grid[0])-1 || x < 0 {
		return "."
	}
	return grid[y][x]
}

func (grid Grid) yMax() int {
	return len(grid)
}

func (grid Grid) xMax() int {
	return len(grid[0])
}

func (grid Grid) Print() {
	for y := 0; y < grid.yMax(); y++ {
		for x := 0; x < grid.xMax(); x++ {
			fmt.Print(grid.get(x, y))
		}
		fmt.Println("")
	}
}

func (grid Grid) equal(otherGrid Grid) bool {
	if grid.yMax() != otherGrid.yMax() || grid.xMax() != otherGrid.xMax() {
		return false
	}
	for y := 0; y < grid.yMax(); y++ {
		for x := 0; x < grid.xMax(); x++ {
			if grid.get(x, y) != otherGrid.get(x, y) {
				return false
			}
		}
	}
	return true
}
