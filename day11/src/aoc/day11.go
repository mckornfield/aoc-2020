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

func IterateSeatsPt2(grid Grid, chairVisibilityMapping map[Pair][]Pair) Grid {
	newGrid := Grid{}
	for y := 0; y < grid.yMax(); y++ {
		newGrid = append(newGrid, make([]string, grid.xMax()))
		for x := 0; x < grid.xMax(); x++ {
			adjacentSeatCount := getVisibleOccupiedSeats(grid, Pair{x: x, y: y}, chairVisibilityMapping)
			currentSeat := grid.get(x, y)
			if currentSeat == "#" && adjacentSeatCount > 4 {
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

func IterateUntilRepeatPt2(input string) int {
	grid := ParseInput(input)
	gridsNotEqual := true
	visibleSeatMapping := getAdjacentVisibleSeats(grid)
	for gridsNotEqual {
		newGrid := IterateSeatsPt2(grid, visibleSeatMapping)
		gridsNotEqual = !newGrid.equal(grid)
		grid = newGrid
		// fmt.Println("----------------")
		// grid.Print()
	}
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

func getVisibleOccupiedSeats(grid Grid, currentSeat Pair, chairVisibilityMapping map[Pair][]Pair) int {
	counter := 0
	for _, pair := range chairVisibilityMapping[currentSeat] {
		if grid.getByPair(pair) == "#" {
			counter++
		}
	}
	// s := fmt.Sprintf("%d,%d %s, count: %d", x, y, grid.get(x, y), counter)
	// fmt.Println(s)
	return counter
}

// This only needs to be done once, and then be a map of Pair to []int

func getAdjacentVisibleSeats(grid Grid) map[Pair][]Pair {
	chairMapping := make(map[Pair][]Pair)
	for y := 0; y < grid.yMax(); y++ {
		for x := 0; x < grid.xMax(); x++ {
			if currentPoint := grid.get(x, y); currentPoint == "L" || currentPoint == "#" {
				centerPoint := Pair{x: x, y: y}
				findAndPopulateAdjacentPoints(grid, chairMapping, centerPoint)
			}
		}
	}
	return chairMapping
}

func findAndPopulateAdjacentPoints(grid Grid, chairMapping map[Pair][]Pair, currentPoint Pair) {
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}
			xCurrent := currentPoint.x + i
			yCurrent := currentPoint.y + j
			point := grid.get(xCurrent, yCurrent)
			for point != "X" {
				if point == "#" || point == "L" {
					foundChairLocation := Pair{x: xCurrent, y: yCurrent}
					if v, ok := chairMapping[currentPoint]; ok {
						chairMapping[currentPoint] = append(v, foundChairLocation)
					} else {
						chairMapping[currentPoint] = []Pair{foundChairLocation}
					}
					break
				}
				// Somewhat hacky, keep moving away in that direction
				yCurrent = yCurrent + j
				xCurrent = xCurrent + i
				point = grid.get(xCurrent, yCurrent)
			}
		}
	}
}

type Pair struct {
	x, y int
}

type Grid [][]string

func (grid Grid) get(x int, y int) string {
	// If out of range, pretend it's a wall, which is an X
	if y > len(grid)-1 || y < 0 || x > len(grid[0])-1 || x < 0 {
		return "X"
	}
	return grid[y][x]
}

func (grid Grid) getByPair(p Pair) string {
	return grid.get(p.x, p.y)
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
