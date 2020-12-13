package aoc

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstExample(t *testing.T) {
	assert := assert.New(t)
	input := `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`
	grid := ParseInput(input)
	initialGrid := grid
	assert.Equal(grid.equal(grid), true)
	assert.Equal("L", grid.get(0, 0))
	assert.Equal("L", grid.get(0, 9))
	assert.Equal("L", grid.get(9, 9))
	assert.Equal(".", grid.get(7, 9))

	grid = IterateSeats(grid)
	assert.Equal(initialGrid.equal(grid), false)
	assert.Equal("#", grid.get(0, 0))
	assert.Equal("#", grid.get(0, 9))
	assert.Equal("#", grid.get(9, 9))
	assert.Equal(".", grid.get(7, 9))

	grid = IterateSeats(grid)
	assert.Equal("#", grid.get(0, 0))
	assert.Equal("#", grid.get(0, 9))
	assert.Equal("#", grid.get(9, 9))
	assert.Equal("L", grid.get(3, 9))
	assert.Equal(".", grid.get(7, 9))
}

func TestFirstExampleGetIterations(t *testing.T) {
	assert := assert.New(t)
	input := `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`
	assert.Equal(37, IterateUntilRepeat(input))
}

func TestFirstPart(t *testing.T) {
	assert := assert.New(t)
	content, err := ioutil.ReadFile("day11_puzzle.txt")
	if err != nil {
		t.Fatal(err.Error())
	}
	input := string(content)
	input = strings.TrimRight(input, "\n")
	assert.Equal(2194, IterateUntilRepeat(input))
	// fmt.Println(text)
}

func TestGeneratingSeatMappings(t *testing.T) {
	assert := assert.New(t)
	input := `.............
.L.L.#.#.#.#.
.............`
	g := ParseInput(input)
	pairMapping := getAdjacentVisibleSeats(g)
	assert.Equal([]Pair{Pair{x: 3, y: 1}},
		pairMapping[Pair{x: 1, y: 1}])
	assert.Equal([]Pair{Pair{x: 1, y: 1}, Pair{x: 5, y: 1}},
		pairMapping[Pair{x: 3, y: 1}])
	assert.Equal([]Pair{Pair{x: 9, y: 1}},
		pairMapping[Pair{x: 11, y: 1}])
}

func TestViewAdjacentOccupiedSeats(t *testing.T) {
	assert := assert.New(t)
	input := `.##.##.
#.#.#.#
##...##
...L...
##...##
#.#.#.#
.##.##.`
	g := ParseInput(input)
	pairMapping := getAdjacentVisibleSeats(g)
	occupiedSeats := getVisibleOccupiedSeats(g, Pair{x: 3, y: 3}, pairMapping)
	assert.Equal(0, occupiedSeats)
}

func TestPt2SampleInput(t *testing.T) {
	assert := assert.New(t)
	input := `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`
	count := IterateUntilRepeatPt2(input)
	assert.Equal(26, count)
}

func TestPt2PuzzleInput(t *testing.T) {
	assert := assert.New(t)
	content, err := ioutil.ReadFile("day11_puzzle.txt")
	if err != nil {
		t.Fatal(err.Error())
	}
	input := string(content)
	input = strings.TrimRight(input, "\n")
	count := IterateUntilRepeatPt2(input)
	assert.Equal(1944, count)
}
